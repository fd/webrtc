package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fd/webrtc"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	webrtc.InitializeSSL()
	defer webrtc.CleanupSSL()

	factory := webrtc.New()

	a := &Observer{"A", make(chan []byte, 50), make(chan []byte, 50), nil, nil}
	b := &Observer{"B", a.sig_snd, a.sig_rcv, nil, nil}

	go a.offer(factory)
	go b.anwser(factory)

	select {}
}

type Observer struct {
	Name                string
	sig_rcv             chan []byte
	sig_snd             chan []byte
	pc                  *webrtc.PeerConnection
	ice_candidate_queue []json.RawMessage
}

func (o *Observer) offer(factory *webrtc.Factory) {
	servers := []*webrtc.ICEServer{
		{URL: "stun:stun.l.google.com:19302"},
	}

	pc_constraints := []*webrtc.MediaConstraint{
		{Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
		{Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
		{Key: webrtc.EnableRtpDataChannelsConstraint, Value: true},
	}

	constraints := []*webrtc.MediaConstraint{
		{Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
		{Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
	}

	var (
		localDesc  *webrtc.SessionDescription
		remoteDesc *webrtc.SessionDescription
		err        error
	)

	log.Printf("[%s] ==> create peer connection", o.Name)
	pc := factory.CreatePeerConnection(servers, pc_constraints, o)
	o.pc = pc

	log.Printf("[%s] ==> create data channel", o.Name)
	dc, err := pc.CreateDataChannel("sig", webrtc.DataChannelOptions{})
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}
	dc.Observer = &Pinger{o, dc, nil}
	defer dc.Close()

	log.Printf("[%s] ==> create offer", o.Name)
	localDesc, err = pc.CreateOffer(constraints)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}

	log.Printf("[%s] ==> set local session description", o.Name)
	err = pc.SetLocalDescription(localDesc)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}

	log.Printf("[%s] ==> send sdp offer (%s)", o.Name, localDesc.Type)
	{
		data, err := json.Marshal(localDesc)
		if err != nil {
			log.Fatalf("[%s] error: %s", o.Name, err)
		}
		o.sig_snd <- data
	}

	log.Printf("[%s] ==> receive sdp answer", o.Name)
	{
		var (
			obj  map[string]interface{}
			data = <-o.sig_rcv
		)
		err := json.Unmarshal(data, &obj)
		if err != nil {
			log.Fatalf("[%s] error: %s", o.Name, err)
		}
		if obj["sdp"] != nil {
			err := json.Unmarshal(data, &remoteDesc)
			if err != nil {
				log.Fatalf("[%s] error: %s", o.Name, err)
			}
		}
	}

	log.Printf("[%s] ==> set remote session description", o.Name)
	err = pc.SetRemoteDescription(remoteDesc)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}

	for sig := range o.sig_rcv {
		log.Printf("[%s] ==> receive signal", o.Name)
		{
			var (
				obj struct {
					Type string
					Sdp  interface{}
				}
			)
			err := json.Unmarshal(sig, &obj)
			if err != nil {
				log.Fatalf("[%s] error: %s", o.Name, err)
			}

			if obj.Type == "candidate" {
				var candidate *webrtc.IceCandidate
				err := json.Unmarshal(sig, &candidate)
				if err != nil {
					log.Fatalf("[%s] error: %s", o.Name, err)
				}

				if pc.AddIceCandidate(candidate) {
					log.Printf("[%s] Added candidate %q", o.Name, candidate)
				} else {
					log.Printf("[%s] Failed to add candidate %q", o.Name, candidate)
				}
			}

			if obj.Sdp != nil {
				var desc *webrtc.SessionDescription
				err := json.Unmarshal(sig, &desc)
				if err != nil {
					log.Fatalf("[%s] error: %s", o.Name, err)
				}

				pc.SetRemoteDescription(desc)
				log.Printf("[%s] set remote sdp\n%s", o.Name, obj.Sdp)
			}
		}
	}
}

func (o *Observer) anwser(factory *webrtc.Factory) {
	servers := []*webrtc.ICEServer{
		{URL: "stun:stun.l.google.com:19302"},
	}

	pc_constraints := []*webrtc.MediaConstraint{
		{Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
		{Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
		{Key: webrtc.EnableRtpDataChannelsConstraint, Value: true},
	}

	constraints := []*webrtc.MediaConstraint{
		{Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
		{Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
	}

	var (
		localDesc  *webrtc.SessionDescription
		remoteDesc *webrtc.SessionDescription
	)

	log.Printf("[%s] ==> receive sdp offer", o.Name)
	{
		var (
			obj  map[string]interface{}
			data = <-o.sig_rcv
		)
		err := json.Unmarshal(data, &obj)
		if err != nil {
			log.Fatalf("[%s] error: %s", o.Name, err)
		}
		if obj["sdp"] != nil {
			err = json.Unmarshal(data, &remoteDesc)
			if err != nil {
				log.Fatalf("[%s] error: %s", o.Name, err)
			}
		}
	}

	log.Printf("[%s] ==> create peer connection", o.Name)
	pc := factory.CreatePeerConnection(servers, pc_constraints, o)
	o.pc = pc

	log.Printf("[%s] ==> set remote session description", o.Name)
	err := pc.SetRemoteDescription(remoteDesc)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}

	log.Printf("[%s] ==> create answer", o.Name)
	localDesc, err = pc.CreateAnswer(constraints)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}

	log.Printf("[%s] ==> set local session description", o.Name)
	err = pc.SetLocalDescription(localDesc)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
		return
	}

	log.Printf("[%s] ==> send sdp answer", o.Name)
	{
		data, err := json.Marshal(localDesc)
		if err != nil {
			log.Fatalf("[%s] error: %s", o.Name, err)
		}
		o.sig_snd <- data
	}

	for sig := range o.sig_rcv {
		log.Printf("[%s] ==> receive signal", o.Name)
		{
			var (
				obj struct {
					Type string
					Sdp  interface{}
				}
			)
			err := json.Unmarshal(sig, &obj)
			if err != nil {
				log.Fatalf("[%s] error: %s", o.Name, err)
			}

			if obj.Type == "candidate" {
				var candidate *webrtc.IceCandidate
				err := json.Unmarshal(sig, &candidate)
				if err != nil {
					log.Fatalf("[%s] error: %s", o.Name, err)
				}

				if pc.AddIceCandidate(candidate) {
					log.Printf("[%s] Added candidate %q", o.Name, candidate)
				} else {
					log.Printf("[%s] Failed to add candidate %q", o.Name, candidate)
				}
			}

			if obj.Sdp != nil {
				var desc *webrtc.SessionDescription
				err := json.Unmarshal(sig, &desc)
				if err != nil {
					log.Fatalf("[%s] error: %s", o.Name, err)
				}

				pc.SetRemoteDescription(desc)
				log.Printf("[%s] set remote sdp\n%s", o.Name, obj.Sdp)
			}
		}
	}
}

func (o *Observer) OnError() {
	log.Printf("[%s] EVENT: %s", o.Name, "OnError")
}
func (o *Observer) OnSignalingChange(s webrtc.SignalingState) {
	log.Printf("[%s] EVENT: %s => %s", o.Name, "OnSignalingChange", s)
}
func (o *Observer) OnAddStream() {
	log.Printf("[%s] EVENT: %s", o.Name, "OnAddStream")
}
func (o *Observer) OnRemoveStream() {
	log.Printf("[%s] EVENT: %s", o.Name, "OnRemoveStream")
}
func (o *Observer) OnDataChannel(dc *webrtc.DataChannel) {
	log.Printf("[%s] EVENT: %s", o.Name, "OnDataChannel")
	dc.Observer = &Pinger{o, dc, nil}
}
func (o *Observer) OnRenegotiationNeeded() {
	log.Printf("[%s] EVENT: %s", o.Name, "OnRenegotiationNeeded")
}
func (o *Observer) OnIceConnectionChange(state webrtc.IceConnectionState) {
	log.Printf("[%s] EVENT: %s => %s", o.Name, "OnIceConnectionChange", state)
}
func (o *Observer) OnIceGatheringChange(state webrtc.IceGatheringState) {
	log.Printf("[%s] EVENT: %s => %s", o.Name, "OnIceGatheringChange", state)
	if state == webrtc.IceGatheringComplete {
		for _, candidate := range o.ice_candidate_queue {
			log.Printf("[%s] send signal %q", o.Name, candidate)
			o.sig_snd <- candidate
		}

		o.ice_candidate_queue = nil
	}
}
func (o *Observer) OnIceCandidate(candidate *webrtc.IceCandidate) {
	data, err := json.Marshal(candidate)
	if err != nil {
		log.Fatalf("[%s] error: %s", o.Name, err)
	}

	log.Printf("[%s] EVENT: %s (%s)", o.Name, "OnIceCandidate", data)
	o.ice_candidate_queue = append(o.ice_candidate_queue, data)
}

type Pinger struct {
	o     *Observer
	dc    *webrtc.DataChannel
	timer *time.Ticker
}

func (p *Pinger) OnStateChange() {
	state := p.dc.State()

	switch state {
	case webrtc.DataChannelOpen:
		log.Printf("=> [%s] dc state: %s", p.o.Name, "open")
	case webrtc.DataChannelClosing:
		log.Printf("=> [%s] dc state: %s", p.o.Name, "closing")
	case webrtc.DataChannelClosed:
		log.Printf("=> [%s] dc state: %s", p.o.Name, "closed")
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
	log.Printf("[%s] rcv: %q", p.o.Name, string(buf))
}

func (p *Pinger) pinger() {
	i := 0
	for _ = range p.timer.C {
		i++
		msg := fmt.Sprintf("ping %d from %s", i, p.o.Name)
		p.dc.Send([]byte(msg))
	}
}
