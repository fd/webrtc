package webrtc

import (
	"C"
	"unsafe"
)

type observer struct {
}

//export c_RTCPeerConnectionObserver_OnError
func c_RTCPeerConnectionObserver_OnError(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnSignalingChange
func c_RTCPeerConnectionObserver_OnSignalingChange(ptr unsafe.Pointer, s C.int) {}

//export c_RTCPeerConnectionObserver_OnStateChange
func c_RTCPeerConnectionObserver_OnStateChange(ptr unsafe.Pointer, s C.int) {}

//export c_RTCPeerConnectionObserver_OnAddStream
func c_RTCPeerConnectionObserver_OnAddStream(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnRemoveStream
func c_RTCPeerConnectionObserver_OnRemoveStream(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnDataChannel
func c_RTCPeerConnectionObserver_OnDataChannel(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnRenegotiationNeeded
func c_RTCPeerConnectionObserver_OnRenegotiationNeeded(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnIceConnectionChange
func c_RTCPeerConnectionObserver_OnIceConnectionChange(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnIceGatheringChange
func c_RTCPeerConnectionObserver_OnIceGatheringChange(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnIceCandidate
func c_RTCPeerConnectionObserver_OnIceCandidate(ptr unsafe.Pointer) {}

//export c_RTCPeerConnectionObserver_OnIceComplete
func c_RTCPeerConnectionObserver_OnIceComplete(ptr unsafe.Pointer) {}
