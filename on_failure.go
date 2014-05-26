package webrtc

/*
#include "wrapper.h"
*/
import "C"

type FailureObserver interface {
	OnFailure(err string)
}

//export c_OnFailure
func c_OnFailure(ref C.Ref, err *C.char) {
	observer, ok := resolve(ref).(FailureObserver)
	if ok && observer != nil {
		observer.OnFailure(C.GoString(err))
	}
}
