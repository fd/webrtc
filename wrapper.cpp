#include <cstdlib>
#include <cstring>
#include <sstream>
#include <iostream>
#include <string>
#include <stdio.h>
#include <netinet/in.h>

#include "talk/app/webrtc/audiotrack.h"
#include "talk/app/webrtc/mediaconstraintsinterface.h"
#include "talk/app/webrtc/mediastreaminterface.h"
#include "talk/app/webrtc/peerconnectionfactory.h"
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

class RTCPeerConnectionObserver;

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



class RTCPeerConnectionObserver {

public:
  RTCPeerConnectionObserver(void* ptr) {
    _go = ptr;
  }

  void OnError() {
    c_RTCPeerConnectionObserver_OnError(_go);
  }

  // Triggered when the SignalingState changed.
  void OnSignalingChange(
     PeerConnectionInterface::SignalingState new_state)
  {
    c_RTCPeerConnectionObserver_OnSignalingChange(_go, int(new_state));
  }

  // Triggered when SignalingState or IceState have changed.
  // TODO(bemasc): Remove once callers transition to OnSignalingChange.
  void OnStateChange(PeerConnectionObserver::StateType state_changed)
  {
    c_RTCPeerConnectionObserver_OnStateChange(_go, int(state_changed));
  }

  // Triggered when media is received on a new stream from remote peer.
  void OnAddStream(MediaStreamInterface* stream)
  {
    c_RTCPeerConnectionObserver_OnAddStream(_go);
  }


  // Triggered when a remote peer close a stream.
  void OnRemoveStream(MediaStreamInterface* stream)
  {
    c_RTCPeerConnectionObserver_OnRemoveStream(_go);
  }


  // Triggered when a remote peer open a data channel.
  // TODO(perkj): Make pure
  void OnDataChannel(DataChannelInterface* data_channel)
  {
    c_RTCPeerConnectionObserver_OnDataChannel(_go);
  }

  // Triggered when renegotiation is needed, for example the ICE has restarted.
  void OnRenegotiationNeeded()
  {
    c_RTCPeerConnectionObserver_OnRenegotiationNeeded(_go);
  }

  // Called any time the IceConnectionState changes
  void OnIceConnectionChange(
      PeerConnectionInterface::IceConnectionState new_state)
  {
    c_RTCPeerConnectionObserver_OnIceConnectionChange(_go);
  }

  // Called any time the IceGatheringState changes
  void OnIceGatheringChange(
      PeerConnectionInterface::IceGatheringState new_state)
  {
    c_RTCPeerConnectionObserver_OnIceGatheringChange(_go);
  }

  // New Ice candidate have been found.
  void OnIceCandidate(const IceCandidateInterface* candidate)
  {
    c_RTCPeerConnectionObserver_OnIceCandidate(_go);
  }

  // TODO(bemasc): Remove this once callers transition to OnIceGatheringChange.
  // All Ice candidates have been found.
  void OnIceComplete()
  {
    c_RTCPeerConnectionObserver_OnIceComplete(_go);
  }

private:
  void* _go;
};

class MediaConstraints {
public:
  MediaConstraints(void** ptr, int len) {
    for (int i=0; i<len; i++) {
      if (c_MediaConstraint_Optional(ptr[i]) > 0) {
        optional.push_back(MediaConstraintsInterface::Constraint(
          c_MediaConstraint_Key(ptr[i]),
          c_MediaConstraint_Value(ptr[i])
        ));
      } else {
        mandatory.push_back(MediaConstraintsInterface::Constraint(
          c_MediaConstraint_Key(ptr[i]),
          c_MediaConstraint_Value(ptr[i])
        ));
      }
    }
  }

  const MediaConstraintsInterface::Constraints& GetMandatory() {
    return mandatory;
  }

  const MediaConstraintsInterface::Constraints& GetOptional() {
    return optional;
  }

  ~MediaConstraints() {

  }

private:
  MediaConstraintsInterface::Constraints mandatory;
  MediaConstraintsInterface::Constraints optional;
};

void* Make_MediaConstraints(void** ptr, int len) {
  return new MediaConstraints(ptr, len);
}

// PeerConnection
PeerConnection WebRTC_PeerConnection_Create(
  Factory factory, int nservers, void* servers,
  void** constraints, int nconstraints,
  void* observerImpl)
{
  PeerConnectionInterface::RTCConfiguration configuration;

  // collect servers
  for (int i = 0; i < c_ICEServers_Len(servers); i++) {
    PeerConnectionInterface::IceServer iceServer;
    iceServer.uri = c_ICEServer_URL(servers, i);
    iceServer.username = c_ICEServer_Username(servers, i);
    iceServer.password = c_ICEServer_Password(servers, i);
    configuration.servers.push_back(iceServer);
  }

  // setup observer
  RTCPeerConnectionObserver *observer = new RTCPeerConnectionObserver(observerImpl);

  DTLSIdentityServiceInterface* dummy_dtls_identity_service = NULL;

  CreateRaw(
    PeerConnectionInterface,
    CastPtr(PeerConnectionFactoryInterface, factory)->CreatePeerConnection(
      configuration,
      (MediaConstraintsInterface*)Make_MediaConstraints(constraints, nconstraints),
      NULL,
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
