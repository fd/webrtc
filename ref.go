package webrtc

/*
#include "wrapper.h"
*/
import "C"

import (
	"sync"
)

var (
	refMap = map[uint64]interface{}{}
	refMtx sync.RWMutex
	refKey uint64
)

func register(i interface{}) C.Ref {
	refMtx.Lock()
	defer refMtx.Unlock()

	refKey++
	refMap[refKey] = i
	return C.Ref(refKey)
}

func unregister(id uint64) {
	refMtx.Lock()
	defer refMtx.Unlock()

	delete(refMap, id)
}

//export c_Ref_Unregister
func c_Ref_Unregister(id C.Ref) {
	unregister(uint64(id))
}

func resolve(id C.Ref) interface{} {
	refMtx.RLock()
	defer refMtx.RUnlock()

	return refMap[uint64(id)]
}
