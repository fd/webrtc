package webrtc

// #include "ref.h"
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

func resolve(id C.Ref) interface{} {
	refMtx.RLock()
	defer refMtx.RUnlock()

	return refMap[uint64(id)]
}

//export go_Ref_Unregister
func go_Ref_Unregister(id C.Ref) {
	unregister(uint64(id))
}
