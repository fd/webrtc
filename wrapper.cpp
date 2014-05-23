#include <cstdlib>
#include <cstring>
#include <sstream>
#include <iostream>
#include <string>
#include <stdio.h>
#include <netinet/in.h>

#include "talk/app/webrtc/audiotrack.h"
#include "talk/app/webrtc/mediastreaminterface.h"
#include "talk/app/webrtc/peerconnectionfactory.h"
#include "talk/app/webrtc/peerconnectioninterface.h"
#include "talk/app/webrtc/videosourceinterface.h"
#include "talk/app/webrtc/videotrack.h"
#include "talk/base/logging.h"
#include "talk/base/ssladapter.h"

#include "wrapper.h"

extern "C" {

#include "_cgo_export.h"
using talk_base::scoped_refptr;
using webrtc::PeerConnectionFactoryInterface;
using webrtc::MediaStreamInterface;

#define CreateRaw(T, ctor) \
  scoped_refptr<T> _ptr = (ctor); \
  T* _raw = _ptr.get(); \
  _raw->AddRef(); \
  return _raw;
#define CastPtr(T, ptr) ((T*)ptr)


const char* WebRTC_Version() {
  return "hello";

}
int WebRTC_InitializeSSL() {
  return talk_base::InitializeSSL();
}

int WebRTC_CleanupSSL() {
  return talk_base::CleanupSSL();
}


void* WebRTC_PeerConnectionFactory_Create() {
  CreateRaw(PeerConnectionFactoryInterface, webrtc::CreatePeerConnectionFactory());
}

void WebRTC_PeerConnectionFactory_Free(void* ptr) {
  if (ptr == NULL) return;
  CastPtr(PeerConnectionFactoryInterface, ptr)->Release();
}

// // RTCMediaStream
void* WebRTC_PeerConnectionFactory_CreateMediaStreamWithLabel(void* ptr, char* clabel) {
  PeerConnectionFactoryInterface* factory = CastPtr(PeerConnectionFactoryInterface, ptr);
  std::string label = clabel;

  CreateRaw(MediaStreamInterface, factory->CreateLocalMediaStream(label));
}

void WebRTC_MediaStream_Free(void* ptr) {
  if (ptr == NULL) return;
  CastPtr(MediaStreamInterface, ptr)->Release();
}


} // extern "C"
