package webrtc

/*
#include "wrapper.h"
*/
import "C"

import (
	"errors"
	"runtime"
	"time"
	"unsafe"
)

type DataChannel struct {
	Observer DataChannelObserver
	ptr      C.DataChannel
	options  DataChannelOptions
}

type DataChannelOptions struct {
	Unordered         bool
	MaxRetransmitTime time.Duration
	MaxRetransmits    int
	Protocol          string
	Negotiated        bool
	Id                int
}

type DataChannelState uint8

const (
	DataChannelConnecting DataChannelState = iota
	DataChannelOpen
	DataChannelClosing
	DataChannelClosed
)

type DataChannelObserver interface {
	OnStateChange()
	OnMessage(msg []byte)
}

func (pc *PeerConnection) CreateDataChannel(label string, options DataChannelOptions) (*DataChannel, error) {
	outer := &DataChannel{options: options}

	outer.ptr = C.WebRTC_DataChannel_Create(pc.ptr, C.CString(label), register(outer))
	runtime.SetFinalizer(outer, (*DataChannel).free)

	if outer.ptr == nil {
		return nil, errors.New("webrtc: failed to create DataChannel")
	}

	return outer, nil
}

func (dc *DataChannel) free() {
	if dc != nil {
		C.WebRTC_DataChannel_Free(dc.ptr)
	}
}

func (dc *DataChannel) State() DataChannelState {
	return DataChannelState(C.WebRTC_DataChannel_State(dc.ptr))
}

func (dc *DataChannel) Close() error {
	C.WebRTC_DataChannel_Close(dc.ptr)
	return nil
}

func (dc *DataChannel) Send(p []byte) bool {
	var (
		ptr unsafe.Pointer
	)

	if len(p) > 0 {
		ptr = unsafe.Pointer(&p[0])
	}

	ok := C.WebRTC_DataChannel_Send(dc.ptr, ptr, C.int(len(p)))
	return ok > 0
}

//export c_DataChannelOptions_Ordered
func c_DataChannelOptions_Ordered(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Unordered {
			return 0
		}
	}
	return 1
}

//export c_DataChannelOptions_MaxRetransmitTime
func c_DataChannelOptions_MaxRetransmitTime(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.MaxRetransmitTime > 0 {
			return C.int(dc.options.MaxRetransmitTime / time.Millisecond)
		}
	}
	return -1
}

//export c_DataChannelOptions_MaxRetransmits
func c_DataChannelOptions_MaxRetransmits(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.MaxRetransmits > 0 {
			return C.int(dc.options.MaxRetransmits)
		}
	}
	return -1
}

//export c_DataChannelOptions_Protocol
func c_DataChannelOptions_Protocol(ref C.Ref) *C.char {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Protocol != "" {
			return C.CString(dc.options.Protocol)
		}
	}
	return C.CString("")
}

//export c_DataChannelOptions_Negotiated
func c_DataChannelOptions_Negotiated(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Negotiated {
			return 1
		}
	}
	return 0
}

//export c_DataChannelOptions_Id
func c_DataChannelOptions_Id(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Id > 0 {
			return C.int(dc.options.Id)
		}
	}
	return -1
}

//export c_DataChannel_OnStateChange
func c_DataChannel_OnStateChange(ref C.Ref) {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.State() == DataChannelClosed {
			unregister(uint64(ref))
		}

		if dc.Observer != nil {
			dc.Observer.OnStateChange()
		}
	}
}

//export c_DataChannel_OnMessage
func c_DataChannel_OnMessage(ref C.Ref, buf unsafe.Pointer, bufn C.int) {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.Observer != nil {
			dc.Observer.OnMessage(C.GoBytes(buf, bufn))
		}
	}
}
