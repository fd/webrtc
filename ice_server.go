package webrtc

import "C"
import "unsafe"

type ICEServer struct {
	URL      string
	Username string
	Password string
}

//export c_ICEServer_URL
func c_ICEServer_URL(s unsafe.Pointer) *C.char {
	if s == nil {
		return nil
	}
	return C.CString((*ICEServer)(s).URL)
}

//export c_ICEServer_Username
func c_ICEServer_Username(s unsafe.Pointer) *C.char {
	if s == nil {
		return nil
	}
	return C.CString((*ICEServer)(s).Username)
}

//export c_ICEServer_Password
func c_ICEServer_Password(s unsafe.Pointer) *C.char {
	if s == nil {
		return nil
	}
	return C.CString((*ICEServer)(s).Password)
}
