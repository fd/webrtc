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
using talk_base::RefCountedObject;
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


class RTCCreateSessionDescriptionObserver : public CreateSessionDescriptionObserver {
public:

  static scoped_refptr<RTCCreateSessionDescriptionObserver> Create(Ref ref) {
    RefCountedObject<RTCCreateSessionDescriptionObserver>* implementation =
         new RefCountedObject<RTCCreateSessionDescriptionObserver>(ref);
    return implementation;
  }

  RTCCreateSessionDescriptionObserver(Ref ref) {
    _ref = ref;
  }

  void OnSuccess(SessionDescriptionInterface* desc) {

  }

  void OnFailure(const std::string& error) {
    c_OnFailure(_ref, strdup(error.c_str()));
  }

protected:
  ~RTCCreateSessionDescriptionObserver() {
    c_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
};

class RTCSetSessionDescriptionObserver : public SetSessionDescriptionObserver {
public:

  static scoped_refptr<RTCSetSessionDescriptionObserver> Create(Ref ref) {
    RefCountedObject<RTCSetSessionDescriptionObserver>* implementation =
         new RefCountedObject<RTCSetSessionDescriptionObserver>(ref);
    return implementation;
  }

  RTCSetSessionDescriptionObserver(Ref ref) {
    _ref = ref;
  }

  void OnSuccess() {

  }

  void OnFailure(const std::string& error) {
    c_OnFailure(_ref, strdup(error.c_str()));
  }

protected:
  ~RTCSetSessionDescriptionObserver() {
    c_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
};


class RTCPeerConnectionObserver : public PeerConnectionObserver {

public:
  RTCPeerConnectionObserver(Ref ref) {
    _ref = ref;
  }

  void OnError() {
    std::cout << "OnError\n";
    c_RTCPeerConnectionObserver_OnError(_ref);
  }

  // Triggered when the SignalingState changed.
  void OnSignalingChange(
     PeerConnectionInterface::SignalingState new_state)
  {
    std::cout << "OnSignalingChange\n";
    c_RTCPeerConnectionObserver_OnSignalingChange(_ref, int(new_state));
  }

  // Triggered when SignalingState or IceState have changed.
  // TODO(bemasc): Remove once callers transition to OnSignalingChange.
  void OnStateChange(PeerConnectionObserver::StateType state_changed)
  {
    std::cout << "OnStateChange\n";
    c_RTCPeerConnectionObserver_OnStateChange(_ref, int(state_changed));
  }

  // Triggered when media is received on a new stream from remote peer.
  void OnAddStream(MediaStreamInterface* stream)
  {
    std::cout << "OnAddStream\n";
    c_RTCPeerConnectionObserver_OnAddStream(_ref);
  }


  // Triggered when a remote peer close a stream.
  void OnRemoveStream(MediaStreamInterface* stream)
  {
    std::cout << "OnRemoveStream\n";
    c_RTCPeerConnectionObserver_OnRemoveStream(_ref);
  }


  // Triggered when a remote peer open a data channel.
  // TODO(perkj): Make pure
  void OnDataChannel(DataChannelInterface* data_channel)
  {
    std::cout << "OnDataChannel\n";
    c_RTCPeerConnectionObserver_OnDataChannel(_ref);
  }

  // Triggered when renegotiation is needed, for example the ICE has restarted.
  void OnRenegotiationNeeded()
  {
    std::cout << "OnRenegotiationNeeded\n";
    c_RTCPeerConnectionObserver_OnRenegotiationNeeded(_ref);
  }

  // Called any time the IceConnectionState changes
  void OnIceConnectionChange(
      PeerConnectionInterface::IceConnectionState new_state)
  {
    std::cout << "OnIceConnectionChange\n";
    c_RTCPeerConnectionObserver_OnIceConnectionChange(_ref);
  }

  // Called any time the IceGatheringState changes
  void OnIceGatheringChange(
      PeerConnectionInterface::IceGatheringState new_state)
  {
    std::cout << "OnIceGatheringChange\n";
    c_RTCPeerConnectionObserver_OnIceGatheringChange(_ref);
  }

  // New Ice candidate have been found.
  void OnIceCandidate(const IceCandidateInterface* candidate)
  {
    std::cout << "OnIceCandidate\n";
    c_RTCPeerConnectionObserver_OnIceCandidate(_ref);
  }

  // TODO(bemasc): Remove this once callers transition to OnIceGatheringChange.
  // All Ice candidates have been found.
  void OnIceComplete()
  {
    std::cout << "OnIceComplete\n";
    c_RTCPeerConnectionObserver_OnIceComplete(_ref);
  }

  ~RTCPeerConnectionObserver() {
    c_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
};

class MediaConstraints : public MediaConstraintsInterface {
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

  const MediaConstraintsInterface::Constraints& GetMandatory() const {
    return mandatory;
  }

  const MediaConstraintsInterface::Constraints& GetOptional() const {
    return optional;
  }

  ~MediaConstraints() {

  }

private:
  MediaConstraintsInterface::Constraints mandatory;
  MediaConstraintsInterface::Constraints optional;
};


// PeerConnection
PeerConnection WebRTC_PeerConnection_Create(
  Factory factory,
  void* servers,     int nservers,
  void* constraints, int nconstraints,
  Ref observerRef)
{
  PeerConnectionInterface::IceServers iceServers;
  void** rservers = (void**)servers;
  void** rconstraints = (void**)constraints;

  // collect servers
  for (int i = 0; i < nservers; i++) {
    PeerConnectionInterface::IceServer iceServer;
    iceServer.uri = c_ICEServer_URL(rservers[i]);
    iceServer.username = c_ICEServer_Username(rservers[i]);
    iceServer.password = c_ICEServer_Password(rservers[i]);
    std::cout << "server: " << iceServer.uri << "\n";
    iceServers.push_back(iceServer);
  }

  // setup observer
  RTCPeerConnectionObserver* observer     = new RTCPeerConnectionObserver(observerRef);
  MediaConstraintsInterface* mconstraints = new MediaConstraints(rconstraints, nconstraints);

  CreateRaw(
    PeerConnectionInterface,
    CastPtr(PeerConnectionFactoryInterface, factory)->CreatePeerConnection(
      iceServers,
      mconstraints,
      NULL,
      NULL,
      observer
    )
  );
}

void WebRTC_PeerConnection_Free(PeerConnection ptr) {
  if (ptr == NULL) return;
  CastPtr(PeerConnectionInterface, ptr)->Release();
}

void WebRTC_PeerConnection_CreateOffer(
  PeerConnection ptr,
  Ref observerRef,
  void* constraints, int nconstraints)
{
  if (ptr == NULL) return;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);
  void** rconstraints = (void**)constraints;

  scoped_refptr<RTCCreateSessionDescriptionObserver> observer = RTCCreateSessionDescriptionObserver::Create(observerRef);

  pc->CreateOffer(observer,
    new MediaConstraints(rconstraints, nconstraints));
}

} // extern "C"
