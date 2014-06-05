package webrtc

/*
#include "data_channel.h"
#include "ref.h"
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
	ptr      unsafe.Pointer
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

var dataChannelStateStrings = map[DataChannelState]string{
	DataChannelConnecting: "connecting",
	DataChannelOpen:       "open",
	DataChannelClosing:    "closing",
	DataChannelClosed:     "closed",
}

func (d DataChannelState) String() string {
	return dataChannelStateStrings[d]
}

type DataChannelObserver interface {
	OnStateChange()
	OnMessage(msg []byte)
}

func (pc *PeerConnection) CreateDataChannel(label string, options DataChannelOptions) (*DataChannel, error) {
	outer := &DataChannel{options: options}

	outer.ptr = C.c_DataChannel_Create(unsafe.Pointer(pc.ptr), C.CString(label), register(outer))
	runtime.SetFinalizer(outer, (*DataChannel).free)

	if outer.ptr == nil {
		return nil, errors.New("webrtc: failed to create DataChannel")
	}

	return outer, nil
}

func (dc *DataChannel) free() {
	if dc != nil {
		C.c_DataChannel_Free(dc.ptr)
	}
}

func (dc *DataChannel) State() DataChannelState {
	return DataChannelState(C.c_DataChannel_State(dc.ptr))
}

func (dc *DataChannel) Close() error {
	C.c_DataChannel_Close(dc.ptr)
	return nil
}

func (dc *DataChannel) Send(p []byte) bool {
	var (
		ptr unsafe.Pointer
	)

	if len(p) > 0 {
		ptr = unsafe.Pointer(&p[0])
	}

	ok := C.c_DataChannel_Send(dc.ptr, ptr, C.int(len(p)))
	return ok > 0
}

//export go_DataChannelOptions_Ordered
func go_DataChannelOptions_Ordered(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Unordered {
			return 0
		}
	}
	return 1
}

//export go_DataChannelOptions_MaxRetransmitTime
func go_DataChannelOptions_MaxRetransmitTime(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.MaxRetransmitTime > 0 {
			return C.int(dc.options.MaxRetransmitTime / time.Millisecond)
		}
	}
	return -1
}

//export go_DataChannelOptions_MaxRetransmits
func go_DataChannelOptions_MaxRetransmits(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.MaxRetransmits > 0 {
			return C.int(dc.options.MaxRetransmits)
		}
	}
	return -1
}

//export go_DataChannelOptions_Protocol
func go_DataChannelOptions_Protocol(ref C.Ref) *C.char {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Protocol != "" {
			return C.CString(dc.options.Protocol)
		}
	}
	return C.CString("")
}

//export go_DataChannelOptions_Negotiated
func go_DataChannelOptions_Negotiated(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Negotiated {
			return 1
		}
	}
	return 0
}

//export go_DataChannelOptions_Id
func go_DataChannelOptions_Id(ref C.Ref) C.int {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.options.Id > 0 {
			return C.int(dc.options.Id)
		}
	}
	return -1
}

//export go_DataChannel_OnStateChange
func go_DataChannel_OnStateChange(ref C.Ref) {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.State() == DataChannelClosed {
			unregister(uint64(ref))
		}

		if dc.Observer != nil {
			dc.Observer.OnStateChange()
		}
	}
}

//export go_DataChannel_OnMessage
func go_DataChannel_OnMessage(ref C.Ref, buf unsafe.Pointer, bufn C.int) {
	if dc, ok := resolve(ref).(*DataChannel); ok && dc != nil {
		if dc.Observer != nil {
			dc.Observer.OnMessage(C.GoBytes(buf, bufn))
		}
	}
}
