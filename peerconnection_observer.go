package webrtc

/*
#include "wrapper.h"
*/
import "C"

type PeerConnectionObserver interface {
	OnError()
	OnSignalingChange()
	OnStateChange()
	OnAddStream()
	OnRemoveStream()
	OnDataChannel()
	OnRenegotiationNeeded()
	OnIceConnectionChange()
	OnIceGatheringChange()
	OnIceCandidate()
	OnIceComplete()
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
		observer.OnSignalingChange()
	}
}

//export c_RTCPeerConnectionObserver_OnStateChange
func c_RTCPeerConnectionObserver_OnStateChange(ref C.Ref, s C.int) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnStateChange()
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
func c_RTCPeerConnectionObserver_OnDataChannel(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnDataChannel()
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
func c_RTCPeerConnectionObserver_OnIceCandidate(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceCandidate()
	}
}

//export c_RTCPeerConnectionObserver_OnIceComplete
func c_RTCPeerConnectionObserver_OnIceComplete(ref C.Ref) {
	observer, ok := resolve(ref).(PeerConnectionObserver)
	if ok && observer != nil {
		observer.OnIceComplete()
	}
}
