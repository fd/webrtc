package webrtc

/*
#include "wrapper.h"
*/
import "C"

import (
	"runtime"
)

type PeerConnectionObserver interface {
	OnError()
	OnSignalingChange(state SignalingState)
	OnStateChange(state State)
	OnAddStream()
	OnRemoveStream()
	OnDataChannel(dataChannel *DataChannel)
	OnRenegotiationNeeded()
	OnIceConnectionChange()
	OnIceGatheringChange()
	OnIceCandidate(candidate *IceCandidate)
	OnIceComplete()
}

type SignalingState uint8

const (
	SignalingStable SignalingState = iota
	SignalingHaveLocalOffer
	SignalingHaveLocalPrAnswer
	SignalingHaveRemoteOffer
	SignalingHaveRemotePrAnswer
	SignalingClosed
)

type State uint8

const (
	StateSignaling State = iota
	StateIce
)

//export c_RTCPeerConnectionObserver_OnError
func c_RTCPeerConnectionObserver_OnError(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnError()
	}
}

//export c_RTCPeerConnectionObserver_OnSignalingChange
func c_RTCPeerConnectionObserver_OnSignalingChange(ref C.Ref, s C.int) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnSignalingChange(SignalingState(s))
	}
}

//export c_RTCPeerConnectionObserver_OnStateChange
func c_RTCPeerConnectionObserver_OnStateChange(ref C.Ref, s C.int) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnStateChange(State(s))
	}
}

//export c_RTCPeerConnectionObserver_OnAddStream
func c_RTCPeerConnectionObserver_OnAddStream(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnAddStream()
	}
}

//export c_RTCPeerConnectionObserver_OnRemoveStream
func c_RTCPeerConnectionObserver_OnRemoveStream(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnRemoveStream()
	}
}

//export c_RTCPeerConnectionObserver_OnDataChannel
func c_RTCPeerConnectionObserver_OnDataChannel(ref C.Ref, ptr C.DataChannel) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		outer := &DataChannel{ptr: ptr}
		C.WebRTC_DataChannel_Accept(ptr, register(outer))
		runtime.SetFinalizer(outer, (*DataChannel).free)

		observer.OnDataChannel(outer)
	}
}

//export c_RTCPeerConnectionObserver_OnRenegotiationNeeded
func c_RTCPeerConnectionObserver_OnRenegotiationNeeded(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnRenegotiationNeeded()
	}
}

//export c_RTCPeerConnectionObserver_OnIceConnectionChange
func c_RTCPeerConnectionObserver_OnIceConnectionChange(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceConnectionChange()
	}
}

//export c_RTCPeerConnectionObserver_OnIceGatheringChange
func c_RTCPeerConnectionObserver_OnIceGatheringChange(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceGatheringChange()
	}
}

//export c_RTCPeerConnectionObserver_OnIceCandidate
func c_RTCPeerConnectionObserver_OnIceCandidate(ref C.Ref, candidate C.IceCandidate) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceCandidate(&IceCandidate{candidate})
	}
}

//export c_RTCPeerConnectionObserver_OnIceComplete
func c_RTCPeerConnectionObserver_OnIceComplete(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceComplete()
	}
}
