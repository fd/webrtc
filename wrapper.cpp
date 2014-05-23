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
using namespace webrtc;

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


Factory WebRTC_PeerConnectionFactory_Create() {
  CreateRaw(PeerConnectionFactoryInterface, webrtc::CreatePeerConnectionFactory());
}

void WebRTC_PeerConnectionFactory_Free(Factory ptr) {
  if (ptr == NULL) return;
  CastPtr(PeerConnectionFactoryInterface, ptr)->Release();
}

// RTCMediaStream
MediaStream WebRTC_CreateMediaStreamWithLabel(Factory ptr, char* clabel) {
  PeerConnectionFactoryInterface* factory = CastPtr(PeerConnectionFactoryInterface, ptr);
  std::string label = clabel;

  CreateRaw(MediaStreamInterface, factory->CreateLocalMediaStream(label));
}

void WebRTC_MediaStream_Free(MediaStream ptr) {
  if (ptr == NULL) return;
  CastPtr(MediaStreamInterface, ptr)->Release();
}

// PeerConnection
PeerConnection WebRTC_PeerConnection_Create(
  Factory factory, int nservers, ICEServer* servers,
  MediaConstraints constraints, PeerConnectionObserver observerImpl)
{
  // collect servers
  PeerConnectionInterface::IceServers iceServers;
  for (int i = 0; i < nservers; i++) {
    PeerConnectionInterface::IceServer iceServer;
    iceServer.uri = [[self.URI absoluteString] UTF8String];
    iceServer.username = [self.username UTF8String];
    iceServer.password = [self.password UTF8String];
    iceServers.push_back(iceServer);
  }

  // setup observer
  RTCPeerConnectionObserver *observer = new RTCPeerConnectionObserver(observerImpl);

  DTLSIdentityServiceInterface* dummy_dtls_identity_service = NULL;

  CreateRaw(
    PeerConnectionInterface,
    factory->CreatePeerConnection(
      iceServers,
      CastPtr(RTCMediaConstraints, constraints),
      dummy_dtls_identity_service,
      observer
    )
  );
}

void WebRTC_PeerConnection_Free(PeerConnection ptr) {
  if (ptr == NULL) return;
  CastPtr(PeerConnectionInterface, ptr)->Release();
}

} // extern "C"
