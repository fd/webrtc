package webrtc

import (
	"runtime"
	"unsafe"
)

/*
#include "wrapper.h"
*/
import "C"

func InitializeSSL() {
	C.WebRTC_InitializeSSL()
}

func CleanupSSL() {
	C.WebRTC_CleanupSSL()
}

type (
	Factory     struct{ ptr unsafe.Pointer }
	MediaStream struct{ ptr unsafe.Pointer }
)

func New() *Factory {
	inner := C.WebRTC_PeerConnectionFactory_Create()
	outer := &Factory{inner}
	runtime.SetFinalizer(outer, (*Factory).free)
	return outer
}

func (f *Factory) free() {
	if f != nil {
		C.WebRTC_PeerConnectionFactory_Free(f.ptr)
	}
}

func (f *Factory) CreateMediaStream(label string) *MediaStream {
	inner := C.WebRTC_PeerConnectionFactory_CreateMediaStreamWithLabel(f.ptr, C.CString(label))
	outer := &MediaStream{inner}
	runtime.SetFinalizer(outer, (*MediaStream).free)
	return outer
}

func (s *MediaStream) free() {
	if s != nil {
		C.WebRTC_MediaStream_Free(s.ptr)
	}
}
