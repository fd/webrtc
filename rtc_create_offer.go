package webrtc

/*
#include "wrapper.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

func (c *PeerConnection) CreateOffer(constraints []*MediaConstraint) error {
	var (
		pconstraints unsafe.Pointer
		observer     = &createOfferObserver{c: make(chan struct{})}
	)

	if len(constraints) > 0 {
		pconstraints = unsafe.Pointer(&constraints[0])
	}

	C.WebRTC_PeerConnection_CreateOffer(c.ptr, register(observer), pconstraints, C.int(len(constraints)))

	<-observer.c
	return observer.err
}

type createOfferObserver struct {
	c   chan struct{}
	err error
}

func (o *createOfferObserver) OnFailure(err string) {
	o.err = errors.New(err)
	o.c <- struct{}{}
}
