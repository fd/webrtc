package webrtc

import (
	"encoding/json"
	"errors"
	"runtime"
	"unsafe"
)

/*
#include "wrapper.h"
#include "ref.h"
*/
import "C"

func InitializeSSL() {
	C.WebRTC_InitializeSSL()
}

func CleanupSSL() {
	C.WebRTC_CleanupSSL()
}

type (
	Factory        struct{ ptr C.Factory }
	MediaStream    struct{ ptr C.MediaStream }
	PeerConnection struct{ ptr C.PeerConnection }
	IceCandidate   struct{ ptr C.IceCandidate }
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

func (f *Factory) CreatePeerConnection(servers []*ICEServer, constraints []*MediaConstraint, observer PeerConnectionObserver) *PeerConnection {
	var (
		pservers     unsafe.Pointer
		pconstraints unsafe.Pointer
	)

	if len(servers) > 0 {
		pservers = unsafe.Pointer(&servers[0])
	}

	if len(constraints) > 0 {
		pconstraints = unsafe.Pointer(&constraints[0])
	}

	inner := C.WebRTC_PeerConnection_Create(f.ptr,
		pservers, C.int(len(servers)),
		pconstraints, C.int(len(constraints)),
		register(observer))
	outer := &PeerConnection{inner}
	runtime.SetFinalizer(outer, (*PeerConnection).free)
	return outer
}

func (s *PeerConnection) free() {
	if s != nil {
		C.WebRTC_PeerConnection_Free(s.ptr)
	}
}

func (s *PeerConnection) AddIceCandidate(candidate *IceCandidate) bool {
	if s == nil || candidate == nil {
		return false
	}

	ok := C.WebRTC_PeerConnection_AddIceCandidate(s.ptr, candidate.ptr)
	return ok == 1
}

func (s *PeerConnection) UpdateIce(servers []*ICEServer, constraints []*MediaConstraint) bool {
	var (
		pservers     unsafe.Pointer
		pconstraints unsafe.Pointer
	)

	if len(servers) > 0 {
		pservers = unsafe.Pointer(&servers[0])
	}

	if len(constraints) > 0 {
		pconstraints = unsafe.Pointer(&constraints[0])
	}

	if s != nil {
		ok := C.WebRTC_PeerConnection_UpdateICE(s.ptr,
			pservers, C.int(len(servers)),
			pconstraints, C.int(len(constraints)))

		if ok == 1 {
			return true
		}
	}

	return false
}

func (f *Factory) CreateMediaStream(label string) *MediaStream {
	inner := C.WebRTC_CreateMediaStreamWithLabel(f.ptr, C.CString(label))
	outer := &MediaStream{inner}
	runtime.SetFinalizer(outer, (*MediaStream).free)
	return outer
}

func (s *MediaStream) free() {
	if s != nil {
		C.WebRTC_MediaStream_Free(s.ptr)
	}
}

func wrap_IceCandidate(cptr C.IceCandidate) *IceCandidate {
	outer := &IceCandidate{cptr}
	runtime.SetFinalizer(outer, (*IceCandidate).free)
	return outer
}

func (c *IceCandidate) free() {
	if c != nil {
		C.WebRTC_IceCandidate_Free(c.ptr)
		c.ptr = nil
	}
}

func (c *IceCandidate) SDP() string {
	if c != nil {
		cstr := C.WebRTC_IceCandidate_SDP(c.ptr)
		defer C.free(unsafe.Pointer(cstr))
		return C.GoString(cstr)
	}
	return ""
}

func (c *IceCandidate) Id() string {
	if c != nil {
		cstr := C.WebRTC_IceCandidate_ID(c.ptr)
		defer C.free(unsafe.Pointer(cstr))
		return C.GoString(cstr)
	}
	return ""
}

func (c *IceCandidate) Index() int {
	if c != nil {
		return int(C.WebRTC_IceCandidate_Index(c.ptr))
	}
	return -1
}

func (c *IceCandidate) MarshalJSON() ([]byte, error) {
	var inner = struct {
		Type      string `json:"type,omitempty"`
		Label     int    `json:"label,omitempty"`
		Id        string `json:"id,omitempty"`
		Candidate string `json:"candidate,omitempty"`
	}{
		Type:      "candidate",
		Label:     c.Index(),
		Id:        c.Id(),
		Candidate: c.SDP(),
	}
	return json.Marshal(&inner)
}

func (c *IceCandidate) UnmarshalJSON(p []byte) error {
	var inner struct {
		Type      string `json:"type,omitempty"`
		Label     int    `json:"label,omitempty"`
		Id        string `json:"id,omitempty"`
		Candidate string `json:"candidate,omitempty"`
	}
	err := json.Unmarshal(p, &inner)
	if err != nil {
		return err
	}

	candidate := C.CString(inner.Candidate)
	defer C.free(unsafe.Pointer(candidate))

	id := C.CString(inner.Id)
	defer C.free(unsafe.Pointer(id))

	ptr := C.WebRTC_IceCandidate_Parse(id, C.int(inner.Label), candidate)
	if ptr == nil {
		return errors.New("webrtc: unable to parse ice candidate")
	}

	c.ptr = ptr
	runtime.SetFinalizer(c, (*IceCandidate).free)

	return nil
}
