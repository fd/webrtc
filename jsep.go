package webrtc

/*
#include "wrapper.h"
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

type (
	SessionDescription struct{ ptr C.SessionDescription }
)

func ParseSessionDescription(typ, desc string) (*SessionDescription, error) {
	ctyp := C.CString(typ)
	defer C.free(unsafe.Pointer(ctyp))

	cdesc := C.CString(desc)
	defer C.free(unsafe.Pointer(cdesc))

	sdp := C.WebRTC_SessionDescription_Parse(ctyp, cdesc)
	if sdp == nil {
		return nil, errors.New("webrtc: unable to parse sdp")
	}

	return wrap_SessionDescription(sdp), nil
}

func wrap_SessionDescription(ptr C.SessionDescription) *SessionDescription {
	outer := &SessionDescription{ptr}
	runtime.SetFinalizer(outer, (*SessionDescription).free)
	return outer
}

func (p *SessionDescription) free() {
	if p != nil {
		C.WebRTC_SessionDescription_Free(p.ptr)
	}
}

func (p *SessionDescription) String() string {
	if p == nil {
		return ""
	}
	cstr := C.WebRTC_SessionDescription_String(p.ptr)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func (p *SessionDescription) AddCandidate(candidate *IceCandidate) bool {
	if p == nil || candidate == nil {
		return false
	}
	ok := C.WebRTC_SessionDescription_AddCandidate(p.ptr, candidate.ptr)
	return ok == 1
}

func (p *SessionDescription) Type() string {
	if p == nil {
		return ""
	}
	cstr := C.WebRTC_SessionDescription_Type(p.ptr)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func (c *PeerConnection) CreateOffer(constraints []*MediaConstraint) (*SessionDescription, error) {
	var (
		pconstraints unsafe.Pointer
		observer     = &createOfferObserver{c: make(chan struct{})}
	)

	if len(constraints) > 0 {
		pconstraints = unsafe.Pointer(&constraints[0])
	}

	C.WebRTC_PeerConnection_CreateOffer(c.ptr, register(observer), pconstraints, C.int(len(constraints)))

	<-observer.c
	return observer.desc, observer.err
}

func (c *PeerConnection) CreateAnswer(constraints []*MediaConstraint) (*SessionDescription, error) {
	var (
		pconstraints unsafe.Pointer
		observer     = &createOfferObserver{c: make(chan struct{})}
	)

	if len(constraints) > 0 {
		pconstraints = unsafe.Pointer(&constraints[0])
	}

	C.WebRTC_PeerConnection_CreateAnswer(c.ptr, register(observer), pconstraints, C.int(len(constraints)))

	<-observer.c
	return observer.desc, observer.err
}

type createOfferObserver struct {
	c    chan struct{}
	desc *SessionDescription
	err  error
}

func (o *createOfferObserver) OnFailure(err string) {
	o.err = errors.New(err)
	o.c <- struct{}{}
}

func (o *createOfferObserver) OnSuccess(desc *SessionDescription) {
	o.desc = desc
	o.c <- struct{}{}
}

//export c_CreateSessionDescription_OnSuccess
func c_CreateSessionDescription_OnSuccess(ref C.Ref, ptr C.SessionDescription) {
	observer, ok := resolve(ref).(*createOfferObserver)
	if ok && observer != nil {
		observer.OnSuccess(wrap_SessionDescription(ptr))
	}
}

func (c *PeerConnection) LocalDescription() *SessionDescription {
	desc := C.WebRTC_PeerConnection_GetLocalDescription(c.ptr)
	return wrap_SessionDescription(desc)
}

func (c *PeerConnection) RemoteDescription() *SessionDescription {
	desc := C.WebRTC_PeerConnection_GetRemoteDescription(c.ptr)
	return wrap_SessionDescription(desc)
}

func (c *PeerConnection) SetLocalDescription(desc *SessionDescription) error {
	var (
		observer = &setDescriptionObserver{c: make(chan struct{})}
	)

	C.WebRTC_PeerConnection_SetLocalDescription(c.ptr, register(observer), desc.ptr)

	<-observer.c
	return observer.err
}

func (c *PeerConnection) SetRemoteDescription(desc *SessionDescription) error {
	var (
		observer = &setDescriptionObserver{c: make(chan struct{})}
	)

	C.WebRTC_PeerConnection_SetRemoteDescription(c.ptr, register(observer), desc.ptr)

	<-observer.c
	return observer.err
}

type setDescriptionObserver struct {
	c   chan struct{}
	err error
}

func (o *setDescriptionObserver) OnFailure(err string) {
	o.err = errors.New(err)
	o.c <- struct{}{}
}

func (o *setDescriptionObserver) OnSuccess() {
	o.c <- struct{}{}
}

//export c_SetSessionDescription_OnSuccess
func c_SetSessionDescription_OnSuccess(ref C.Ref) {
	observer, ok := resolve(ref).(*setDescriptionObserver)
	if ok && observer != nil {
		observer.OnSuccess()
	}
}
