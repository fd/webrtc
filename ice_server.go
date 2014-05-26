package webrtc

import "C"
import "unsafe"

type ICEServer struct {
	URL      string
	Username string
	Password string
}

//export c_ICEServers_Len
func c_ICEServers_Len(s unsafe.Pointer) int {
	if s == nil {
		return 0
	}
	return len([]ICEServer(s))
}

//export c_ICEServer_URL
func c_ICEServer_URL(s unsafe.Pointer, i int) *C.char {
	if s == nil || i < 0 || i >= c_ICEServers_Len(s) {
		return nil
	}
	return C.CString([]ICEServer(s)[i].URL)
}

//export c_ICEServer_Username
func c_ICEServer_Username(s unsafe.Pointer, i int) *C.char {
	if s == nil || i < 0 || i >= c_ICEServers_Len(s) {
		return nil
	}
	return C.CString([]ICEServer(s)[i].Username)
}

//export c_ICEServer_Password
func c_ICEServer_Password(s unsafe.Pointer, i int) *C.char {
	if s == nil || i < 0 || i >= c_ICEServers_Len(s) {
		return nil
	}
	return C.CString([]ICEServer(s)[i].Password)
}
