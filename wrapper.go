package webrtc

/*
#include "wrapper.h"
*/
import "C"

import (
  "runtime"
  "unsafe"
)

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
