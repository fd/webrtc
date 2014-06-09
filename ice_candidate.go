package webrtc

/*
#include "wrapper.h"
#include "ice_candidate.h"
*/
import "C"

import (
	"errors"
	"os"
	"unsafe"
)

type IceCandidate struct {
	SdpMid        string `json:"sdp_mid"`
	SdpMlineIndex int    `json:"sdp_mline_index"`
	Candidate     string `json:"candidate"`
}

func (s *PeerConnection) AddIceCandidate(candidate *IceCandidate) error {
	if s == nil || candidate == nil {
		return os.ErrInvalid
	}

	cerr := C.c_PeerConnection_AddIceCandidate(unsafe.Pointer(s.ptr), unsafe.Pointer(candidate))
	if cerr != nil {
		defer C.free(unsafe.Pointer(cerr))
		return errors.New(C.GoString(cerr))
	}

	return nil
}

//export go_IceCandidate_ToGo
func go_IceCandidate_ToGo(sdp_mid *C.char, sdp_mline_index C.int, sdp *C.char) unsafe.Pointer {
	candidate := &IceCandidate{
		SdpMid:        C.GoString(sdp_mid),
		SdpMlineIndex: int(sdp_mline_index),
		Candidate:     C.GoString(sdp),
	}
	return unsafe.Pointer(candidate)
}

//export go_IceCandidate_GetSdpMid
func go_IceCandidate_GetSdpMid(ptr unsafe.Pointer) *C.char {
	if ptr == nil {
		return C.CString("")
	}

	candidate := (*IceCandidate)(ptr)
	return C.CString(candidate.SdpMid)
}

//export go_IceCandidate_GetSdpMlineIndex
func go_IceCandidate_GetSdpMlineIndex(ptr unsafe.Pointer) C.int {
	if ptr == nil {
		return -1
	}

	candidate := (*IceCandidate)(ptr)
	return C.int(candidate.SdpMlineIndex)
}

//export go_IceCandidate_GetCandidate
func go_IceCandidate_GetCandidate(ptr unsafe.Pointer) *C.char {
	if ptr == nil {
		return C.CString("")
	}

	candidate := (*IceCandidate)(ptr)
	return C.CString(candidate.Candidate)
}
