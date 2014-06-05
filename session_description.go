package webrtc

// #include "session_description.h"
// #include "ref.h"
// #include "stdlib.h"
import "C"

import (
	"errors"
	"os"
	"unsafe"
)

type SessionDescriptionType string

const (
	Offer             = SessionDescriptionType("offer")
	Answer            = SessionDescriptionType("answer")
	ProvisionalAnswer = SessionDescriptionType("pranswer")
)

type SessionDescription struct {
	Type SessionDescriptionType `json:"type,omitempty"`
	Sdp  string                 `json:"sdp,omitempty"`
}

func (pc *PeerConnection) LocalDescription() *SessionDescription {
	if pc == nil {
		return nil
	}

	ptr := C.c_PeerConnection_GetLocalDescription(unsafe.Pointer(pc.ptr))
	return (*SessionDescription)(ptr)
}

func (pc *PeerConnection) RemoteDescription() *SessionDescription {
	if pc == nil {
		return nil
	}

	ptr := C.c_PeerConnection_GetRemoteDescription(unsafe.Pointer(pc.ptr))
	return (*SessionDescription)(ptr)
}

func (pc *PeerConnection) SetLocalDescription(desc *SessionDescription) error {
	if pc == nil || desc == nil {
		return os.ErrInvalid
	}

	var (
		observer = &setDescriptionObserver{c: make(chan struct{})}
	)

	cerr := C.c_PeerConnection_SetLocalDescription(
		unsafe.Pointer(pc.ptr), register(observer), unsafe.Pointer(desc))
	if cerr != nil {
		defer C.free(unsafe.Pointer(cerr))
		return errors.New(C.GoString(cerr))
	}

	<-observer.c
	return observer.err
}

func (pc *PeerConnection) SetRemoteDescription(desc *SessionDescription) error {
	if pc == nil || desc == nil {
		return os.ErrInvalid
	}

	var (
		observer = &setDescriptionObserver{c: make(chan struct{})}
	)

	cerr := C.c_PeerConnection_SetRemoteDescription(
		unsafe.Pointer(pc.ptr), register(observer), unsafe.Pointer(desc))
	if cerr != nil {
		defer C.free(unsafe.Pointer(cerr))
		return errors.New(C.GoString(cerr))
	}

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

//export go_SetSessionDescription_OnSuccess
func go_SetSessionDescription_OnSuccess(ref C.Ref) {
	observer, ok := resolve(ref).(*setDescriptionObserver)
	if ok && observer != nil {
		observer.OnSuccess()
	}
}

//export go_SetSessionDescription_OnFailure
func go_SetSessionDescription_OnFailure(ref C.Ref, err *C.char) {
	observer, ok := resolve(ref).(*setDescriptionObserver)
	if ok && observer != nil {
		observer.OnFailure(C.GoString(err))
	}
}

func (c *PeerConnection) CreateOffer(constraints []*MediaConstraint) (*SessionDescription, error) {
	var (
		pconstraints unsafe.Pointer
		observer     = &createOfferObserver{c: make(chan struct{})}
	)

	if len(constraints) > 0 {
		pconstraints = unsafe.Pointer(&constraints[0])
	}

	C.c_PeerConnection_CreateOffer(
		unsafe.Pointer(c.ptr), register(observer),
		pconstraints, C.int(len(constraints)))

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

	C.c_PeerConnection_CreateAnswer(
		unsafe.Pointer(c.ptr), register(observer),
		pconstraints, C.int(len(constraints)))

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

//export go_CreateSessionDescription_OnSuccess
func go_CreateSessionDescription_OnSuccess(ref C.Ref, ptr unsafe.Pointer) {
	observer, ok := resolve(ref).(*createOfferObserver)
	if ok && observer != nil {
		observer.OnSuccess((*SessionDescription)(ptr))
	}
}

//export go_CreateSessionDescription_OnFailure
func go_CreateSessionDescription_OnFailure(ref C.Ref, err *C.char) {
	observer, ok := resolve(ref).(*createOfferObserver)
	if ok && observer != nil {
		observer.OnFailure(C.GoString(err))
	}
}

//export go_SessionDescription_New
func go_SessionDescription_New(typ, sdp *C.char) unsafe.Pointer {
	ptr := &SessionDescription{
		Type: SessionDescriptionType(C.GoString(typ)),
		Sdp:  C.GoString(sdp),
	}
	return unsafe.Pointer(ptr)
}

//export go_SessionDescription_GetType
func go_SessionDescription_GetType(ptr unsafe.Pointer) *C.char {
	if ptr == nil {
		return nil
	}
	desc := (*SessionDescription)(ptr)
	str := string(desc.Type)
	return C.CString(str)
}

//export go_SessionDescription_GetSdp
func go_SessionDescription_GetSdp(ptr unsafe.Pointer) *C.char {
	if ptr == nil {
		return nil
	}
	desc := (*SessionDescription)(ptr)
	return C.CString(desc.Sdp)
}
