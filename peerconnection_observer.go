package webrtc

/*
#include "wrapper.h"
#include "data_channel.h"
#include "ref.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

type PeerConnectionObserver interface {
	OnError()
	OnSignalingChange(state SignalingState)
	OnAddStream()
	OnRemoveStream()
	OnDataChannel(dataChannel *DataChannel)
	OnRenegotiationNeeded()
	OnIceConnectionChange(state IceConnectionState)
	OnIceGatheringChange(state IceGatheringState)
	OnIceCandidate(candidate *IceCandidate)
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

var signalingStateStrings = map[SignalingState]string{
	SignalingStable:             "stable",
	SignalingHaveLocalOffer:     "have-local-offer",
	SignalingHaveLocalPrAnswer:  "have-local-pranswer",
	SignalingHaveRemoteOffer:    "have-remote-offer",
	SignalingHaveRemotePrAnswer: "have-remote-pranswer",
	SignalingClosed:             "closed",
}

func (s SignalingState) String() string {
	return signalingStateStrings[s]
}

type IceGatheringState uint8

const (
	IceGatheringNew IceGatheringState = iota
	IceGatheringGathering
	IceGatheringComplete
)

var iceGatheringStateStrings = map[IceGatheringState]string{
	IceGatheringNew:       "new",
	IceGatheringGathering: "gathering",
	IceGatheringComplete:  "complete",
}

func (s IceGatheringState) String() string {
	return iceGatheringStateStrings[s]
}

type IceConnectionState uint8

const (
	IceConnectionNew IceConnectionState = iota
	IceConnectionChecking
	IceConnectionConnected
	IceConnectionCompleted
	IceConnectionFailed
	IceConnectionDisconnected
	IceConnectionClosed
)

var iceConnectionStateStrings = map[IceConnectionState]string{
	IceConnectionNew:          "new",
	IceConnectionChecking:     "checking",
	IceConnectionConnected:    "connected",
	IceConnectionCompleted:    "completed",
	IceConnectionFailed:       "failed",
	IceConnectionDisconnected: "disconnected",
	IceConnectionClosed:       "closed",
}

func (s IceConnectionState) String() string {
	return iceConnectionStateStrings[s]
}

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
func c_RTCPeerConnectionObserver_OnDataChannel(ref C.Ref, ptr unsafe.Pointer) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		outer := &DataChannel{ptr: ptr}
		C.c_DataChannel_Accept(ptr, register(outer))
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
func c_RTCPeerConnectionObserver_OnIceConnectionChange(ref C.Ref, state C.int) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceConnectionChange(IceConnectionState(state))
	}
}

//export c_RTCPeerConnectionObserver_OnIceGatheringChange
func c_RTCPeerConnectionObserver_OnIceGatheringChange(ref C.Ref, state C.int) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceGatheringChange(IceGatheringState(state))
	}
}

//export c_RTCPeerConnectionObserver_OnIceCandidate
func c_RTCPeerConnectionObserver_OnIceCandidate(ref C.Ref, candidate C.IceCandidate) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceCandidate(&IceCandidate{candidate})
	}
}
