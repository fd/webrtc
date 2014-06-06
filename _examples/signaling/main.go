package main

import (
  "bytes"
  "encoding/json"
  "flag"
  "fmt"
  "log"
  "net/http"
  "time"

  "github.com/fd/webrtc"
)

var (
  joinPtr       = flag.String("join", "", "Join another server")
  signalAddrPtr = flag.String("signal-addr", ":0", "Addr of signaling server")
)

func main() {
  flag.Parse()
  webrtc.InitializeSSL()

  peers := NewPeers()

  if *joinPtr != "" {
    _, err := peers.Join(*joinPtr)
    if err != nil {
      log.Fatalln(err)
    }
  }

  http.ListenAndServe(*signalAddrPtr, peers)
}

type Peers struct {
  factory *webrtc.Factory
}

func NewPeers() *Peers {
  factory := webrtc.New()
  return &Peers{factory}
}

func (p *Peers) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  if req.Header.Get("Content-Type") != "text/jsep+offer" {
    http.NotFound(rw, req)
    return
  }

  var (
    offer  *webrtc.SessionDescription
    answer *webrtc.SessionDescription
  )

  err := json.NewDecoder(req.Body).Decode(&offer)
  if err != nil {
    panic(err) //internal
  }

  peer, err := p.Accept(offer)
  if err != nil {
    panic(err) //internal
  }

  answer = peer.conn.LocalDescription()

  rw.WriteHeader(200)
  json.NewEncoder(rw).Encode(answer)
}

type Peer struct {
  conn     *webrtc.PeerConnection
  signals  *webrtc.DataChannel
  join_url string
  setup    bool
  ready    chan struct{}
}

func (p *Peers) Accept(offer *webrtc.SessionDescription) (*Peer, error) {
  peer := &Peer{
    ready: make(chan struct{}),
  }

  servers := []*webrtc.ICEServer{
    {URL: "stun:stun.l.google.com:19302"},
  }

  constraints := []*webrtc.MediaConstraint{
    {Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
    {Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
    {Key: webrtc.EnableRtpDataChannelsConstraint, Value: true},
  }

  peer.conn = p.factory.CreatePeerConnection(servers, constraints, peer)

  err := peer.conn.SetRemoteDescription(offer)
  if err != nil {
    return nil, err
  }

  desc, err := peer.conn.CreateAnswer([]*webrtc.MediaConstraint{
    {Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
    {Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
  })
  if err != nil {
    return nil, err
  }

  peer.conn.SetLocalDescription(desc)

  <-peer.ready

  return peer, nil
}

func (p *Peers) Join(rawurl string) (*Peer, error) {
  peer := &Peer{join_url: rawurl}

  servers := []*webrtc.ICEServer{
    {URL: "stun:stun.l.google.com:19302"},
  }

  constraints := []*webrtc.MediaConstraint{
    {Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
    {Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
    {Key: webrtc.EnableRtpDataChannelsConstraint, Value: true},
  }

  peer.conn = p.factory.CreatePeerConnection(servers, constraints, peer)

  options := webrtc.DataChannelOptions{}
  // options.MaxRetransmitTime = 1 * time.Minute
  // options.MaxRetransmits = 15
  dc, err := peer.conn.CreateDataChannel("signals", options)
  if err != nil {
    return nil, err
  }

  peer.signals = dc
  dc.Observer = &Pinger{dc, nil}

  offer, err := peer.conn.CreateOffer([]*webrtc.MediaConstraint{
    {Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
    {Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
  })
  if err != nil {
    return nil, err
  }

  log.Printf("OFFER: %+v", offer)

  err = peer.conn.SetLocalDescription(offer)
  if err != nil {
    return nil, err
  }

  return peer, nil
}

func (p *Peer) OnAddStream() {

}

func (p *Peer) OnRemoveStream() {

}

func (p *Peer) OnRenegotiationNeeded() {
  log.Printf("OnRenegotiationNeeded")
}

func (p *Peer) OnError() {
  log.Printf("OnError")
}

func (p *Peer) OnSignalingChange(state webrtc.SignalingState) {
  log.Printf("OnSignalingChange: %v", state)
}

func (p *Peer) OnIceCandidate(candidate *webrtc.IceCandidate) {
  log.Printf("OnIceCandidate: %v", candidate)
}

func (p *Peer) OnIceConnectionChange(state webrtc.IceConnectionState) {
  log.Printf("OnIceConnectionChange: %s", state)
}

func (p *Peer) OnIceGatheringChange(state webrtc.IceGatheringState) {
  log.Printf("OnIceGatheringChange: %s", state)

  if state != webrtc.IceGatheringComplete {
    return
  }

  if p.setup {
    return
  }

  go func() {
    // joiner
    if p.join_url != "" {
      var (
        offer     *webrtc.SessionDescription
        answer    *webrtc.SessionDescription
        offerData bytes.Buffer
      )

      offer = p.conn.LocalDescription()
      err := json.NewEncoder(&offerData).Encode(offer)
      if err != nil {
        log.Fatalln(err)
      }

      resp, err := http.Post(p.join_url, "text/jsep+offer", &offerData)
      if err != nil {
        log.Fatalln(err)
      }
      defer resp.Body.Close()

      err = json.NewDecoder(resp.Body).Decode(&answer)
      if err != nil {
        log.Fatalln(err)
      }

      err = p.conn.SetRemoteDescription(answer)
      if err != nil {
        log.Fatalln(err)
      }

      return
    }

    if p.ready != nil {
      p.ready <- struct{}{}
      p.ready = nil
      return
    }
  }()
}

func (p *Peer) OnDataChannel(c *webrtc.DataChannel) {
  c.Observer = &Pinger{c, nil}
}

type Pinger struct {
  dc    *webrtc.DataChannel
  timer *time.Ticker
}

func (p *Pinger) OnStateChange() {
  state := p.dc.State()

  switch state {
  case webrtc.DataChannelOpen:
    log.Printf("=> [%s] dc state: %s", "pinger", "open")
  case webrtc.DataChannelClosing:
    log.Printf("=> [%s] dc state: %s", "pinger", "closing")
  case webrtc.DataChannelClosed:
    log.Printf("=> [%s] dc state: %s", "pinger", "closed")
  }

  switch state {
  case webrtc.DataChannelOpen:
    p.timer = time.NewTicker(5 * time.Second)
    go p.pinger()
  case webrtc.DataChannelClosing, webrtc.DataChannelClosed:
    if p.timer != nil {
      p.timer.Stop()
      p.timer = nil
    }
  }
}

func (p *Pinger) OnMessage(buf []byte) {
  log.Printf("[%s] rcv: %q", "pinger", string(buf))
}

func (p *Pinger) pinger() {
  i := 0
  for _ = range p.timer.C {
    i++
    msg := fmt.Sprintf("ping %d from %s", i, "pinger")
    p.dc.Send([]byte(msg))
  }
}
