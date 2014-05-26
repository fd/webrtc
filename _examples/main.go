package main

import (
	"fmt"

	"github.com/fd/webrtc"
)

func main() {
	webrtc.InitializeSSL()

	factory := webrtc.New()

	servers := []*webrtc.ICEServer{
		{URL: "stun:stun.l.google.com:19302"},
		{URL: "stun:stun1.l.google.com:19302"},
		{URL: "stun:stun2.l.google.com:19302"},
		{URL: "stun:stun3.l.google.com:19302"},
		{URL: "stun:stun4.l.google.com:19302"},
	}

	constraints := []*webrtc.MediaConstraint{
		{Key: webrtc.OfferToReceiveAudioConstraint, Value: false},
		{Key: webrtc.OfferToReceiveVideoConstraint, Value: false},
	}

	fmt.Println("==> create peer connection")
	pc := factory.CreatePeerConnection(servers, constraints, &Observer{})

	fmt.Println("==> create offer")
	pc.CreateOffer(&OfferObserver{}, constraints)

	select {}
}

type OfferObserver struct{}

func (o *OfferObserver) OnFailure(err string) {
	fmt.Printf("EVENT: %s => %q\n", "OnFailure", err)
}

type Observer struct{}

func (o *Observer) OnError() {
	fmt.Printf("EVENT: %s\n", "OnError")
}
func (o *Observer) OnSignalingChange() {
	fmt.Printf("EVENT: %s\n", "OnSignalingChange")
}
func (o *Observer) OnStateChange() {
	fmt.Printf("EVENT: %s\n", "OnStateChange")
}
func (o *Observer) OnAddStream() {
	fmt.Printf("EVENT: %s\n", "OnAddStream")
}
func (o *Observer) OnRemoveStream() {
	fmt.Printf("EVENT: %s\n", "OnRemoveStream")
}
func (o *Observer) OnDataChannel() {
	fmt.Printf("EVENT: %s\n", "OnDataChannel")
}
func (o *Observer) OnRenegotiationNeeded() {
	fmt.Printf("EVENT: %s\n", "OnRenegotiationNeeded")
}
func (o *Observer) OnIceConnectionChange() {
	fmt.Printf("EVENT: %s\n", "OnIceConnectionChange")
}
func (o *Observer) OnIceGatheringChange() {
	fmt.Printf("EVENT: %s\n", "OnIceGatheringChange")
}
func (o *Observer) OnIceCandidate() {
	fmt.Printf("EVENT: %s\n", "OnIceCandidate")
}
func (o *Observer) OnIceComplete() {
	fmt.Printf("EVENT: %s\n", "OnIceComplete")
}
