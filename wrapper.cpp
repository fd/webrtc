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

void WebRTC_SessionDescription_Free(SessionDescription ptr) {
  if (ptr == NULL) return;
  delete (SessionDescriptionInterface*)ptr;
}

char* WebRTC_SessionDescription_String(SessionDescription ptr) {
  if (ptr == NULL) return NULL;
  std::string sd;
  ((SessionDescriptionInterface*)ptr)->ToString(&sd);
  return strdup(sd.c_str());
}

int WebRTC_SessionDescription_AddCandidate(SessionDescription ptr, IceCandidate c)
{
  if (ptr == NULL || c == NULL) return 0;
  SessionDescriptionInterface* _ptr = CastPtr(SessionDescriptionInterface, ptr);
  IceCandidateInterface* _c = CastPtr(IceCandidateInterface, c);
  return _ptr->AddCandidate(_c) ? 1 : 0;
}

char* WebRTC_SessionDescription_Type(SessionDescription ptr) {
  if (ptr == NULL) return NULL;
  std::string type = ((SessionDescriptionInterface*)ptr)->type();
  return strdup(type.c_str());
}

SessionDescription WebRTC_SessionDescription_Parse(char* type, char* raw)
{
  std::string _type = type;
  std::string sdp = raw;
  return CreateSessionDescription(_type, sdp, NULL);
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
    c_CreateSessionDescription_OnSuccess(_ref, desc);
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
    c_SetSessionDescription_OnSuccess(_ref);
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
    c_RTCPeerConnectionObserver_OnError(_ref);
  }

  // Triggered when the SignalingState changed.
  void OnSignalingChange(
     PeerConnectionInterface::SignalingState new_state)
  {
    c_RTCPeerConnectionObserver_OnSignalingChange(_ref, int(new_state));
  }

  // Triggered when SignalingState or IceState have changed.
  // TODO(bemasc): Remove once callers transition to OnSignalingChange.
  void OnStateChange(PeerConnectionObserver::StateType state_changed)
  {
    c_RTCPeerConnectionObserver_OnStateChange(_ref, int(state_changed));
  }

  // Triggered when media is received on a new stream from remote peer.
  void OnAddStream(MediaStreamInterface* stream)
  {
    c_RTCPeerConnectionObserver_OnAddStream(_ref);
  }


  // Triggered when a remote peer close a stream.
  void OnRemoveStream(MediaStreamInterface* stream)
  {
    c_RTCPeerConnectionObserver_OnRemoveStream(_ref);
  }


  // Triggered when a remote peer open a data channel.
  // TODO(perkj): Make pure
  void OnDataChannel(DataChannelInterface* data_channel)
  {
    // keep the channel
    data_channel->AddRef();
    c_RTCPeerConnectionObserver_OnDataChannel(_ref, data_channel);
  }

  // Triggered when renegotiation is needed, for example the ICE has restarted.
  void OnRenegotiationNeeded()
  {
    c_RTCPeerConnectionObserver_OnRenegotiationNeeded(_ref);
  }

  // Called any time the IceConnectionState changes
  void OnIceConnectionChange(
      PeerConnectionInterface::IceConnectionState new_state)
  {
    c_RTCPeerConnectionObserver_OnIceConnectionChange(_ref);
  }

  // Called any time the IceGatheringState changes
  void OnIceGatheringChange(
      PeerConnectionInterface::IceGatheringState new_state)
  {
    c_RTCPeerConnectionObserver_OnIceGatheringChange(_ref);
  }

  // New Ice candidate have been found.
  void OnIceCandidate(const IceCandidateInterface* candidate)
  {
    c_RTCPeerConnectionObserver_OnIceCandidate(_ref, IceCandidate(candidate));
  }

  // TODO(bemasc): Remove this once callers transition to OnIceGatheringChange.
  // All Ice candidates have been found.
  void OnIceComplete()
  {
    c_RTCPeerConnectionObserver_OnIceComplete(_ref);
  }

  ~RTCPeerConnectionObserver() {
    c_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
};


class RTCDataChannelObserver : public DataChannelObserver {

public:
  RTCDataChannelObserver(Ref ref) {
    _ref = ref;
  }

  ~RTCDataChannelObserver() {
    c_Ref_Unregister(_ref);
  }

  void OnStateChange()
  {
    c_DataChannel_OnStateChange(_ref);
  }

  void OnMessage(const DataBuffer& buffer)
  {
    c_DataChannel_OnMessage(_ref, (void*)buffer.data.data(), buffer.size());
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

int WebRTC_PeerConnection_UpdateICE(
  PeerConnection ptr,
  void* servers,     int nservers,
  void* constraints, int nconstraints)
{
  if (ptr == NULL) return 0;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);

  PeerConnectionInterface::IceServers iceServers;
  void** rservers = (void**)servers;
  void** rconstraints = (void**)constraints;

  // collect servers
  for (int i = 0; i < nservers; i++) {
    PeerConnectionInterface::IceServer iceServer;
    iceServer.uri = c_ICEServer_URL(rservers[i]);
    iceServer.username = c_ICEServer_Username(rservers[i]);
    iceServer.password = c_ICEServer_Password(rservers[i]);
    iceServers.push_back(iceServer);
  }

  MediaConstraintsInterface* mconstraints = new MediaConstraints(rconstraints, nconstraints);

  return pc->UpdateIce(iceServers, mconstraints);
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

void WebRTC_PeerConnection_CreateAnswer(
  PeerConnection ptr,
  Ref observerRef,
  void* constraints, int nconstraints)
{
  if (ptr == NULL) return;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);
  void** rconstraints = (void**)constraints;

  scoped_refptr<RTCCreateSessionDescriptionObserver> observer = RTCCreateSessionDescriptionObserver::Create(observerRef);

  pc->CreateAnswer(observer,
    new MediaConstraints(rconstraints, nconstraints));
}

void WebRTC_PeerConnection_SetLocalDescription(
  PeerConnection ptr,
  Ref observerRef,
  SessionDescription desc)
{
  if (ptr == NULL || desc == NULL) return;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);

  scoped_refptr<RTCSetSessionDescriptionObserver> observer =
    RTCSetSessionDescriptionObserver::Create(observerRef);

  pc->SetLocalDescription(observer,
    (SessionDescriptionInterface*)desc);
}

void WebRTC_PeerConnection_SetRemoteDescription(
  PeerConnection ptr,
  Ref observerRef,
  SessionDescription desc)
{
  if (ptr == NULL || desc == NULL) return;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);

  scoped_refptr<RTCSetSessionDescriptionObserver> observer =
    RTCSetSessionDescriptionObserver::Create(observerRef);

  pc->SetRemoteDescription(observer,
    (SessionDescriptionInterface*)desc);
}

SessionDescription WebRTC_PeerConnection_GetLocalDescription(
  PeerConnection ptr)
{
  if (ptr == NULL) return NULL;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);

  return (SessionDescription)pc->local_description();
}

SessionDescription WebRTC_PeerConnection_GetRemoteDescription(
  PeerConnection ptr)
{
  if (ptr == NULL) return NULL;
  PeerConnectionInterface* pc = CastPtr(PeerConnectionInterface, ptr);

  return (SessionDescription)pc->remote_description();
}


int WebRTC_PeerConnection_AddIceCandidate(
  PeerConnection ptr,
  IceCandidate c)
{
  if (ptr == NULL || c == NULL) return 0;
  PeerConnectionInterface* _ptr = CastPtr(PeerConnectionInterface, ptr);
  IceCandidateInterface* _c = CastPtr(IceCandidateInterface, c);
  return _ptr->AddIceCandidate(_c) ? 1 : 0;
}


DataChannel WebRTC_DataChannel_Create(PeerConnection pc, char* label, Ref dc) {
  if (pc == NULL || label == NULL) return NULL;
  PeerConnectionInterface* _pc = CastPtr(PeerConnectionInterface, pc);

  std::string _label = label;
  if (label != NULL) free(label);

  DataChannelInit* options = new DataChannelInit();
  options->ordered = c_DataChannelOptions_Ordered(dc) == 1 ? true : false;
  options->maxRetransmitTime = c_DataChannelOptions_MaxRetransmitTime(dc);
  options->maxRetransmits = c_DataChannelOptions_MaxRetransmits(dc);
  options->protocol = c_DataChannelOptions_Protocol(dc);
  options->negotiated = c_DataChannelOptions_Negotiated(dc) == 1 ? true : false;
  options->id = c_DataChannelOptions_Id(dc);

  scoped_refptr<DataChannelInterface> _dc =
    _pc->CreateDataChannel(_label, options);
  if (_dc == NULL) {
    return NULL;
  }

  RTCDataChannelObserver* observer = new RTCDataChannelObserver(dc);
  _dc->RegisterObserver(observer);

  CreateRaw(DataChannelInterface, _dc)
}

void WebRTC_DataChannel_Accept(DataChannel ptr, Ref ref)
{
  if (ptr == NULL) return;
  DataChannelInterface* dc = CastPtr(DataChannelInterface, ptr);

  RTCDataChannelObserver* observer = new RTCDataChannelObserver(ref);
  dc->RegisterObserver(observer);
}

void WebRTC_DataChannel_Free(DataChannel ptr) {
  if (ptr == NULL) return;
  CastPtr(DataChannelInterface, ptr)->Release();
}

int WebRTC_DataChannel_State(
  DataChannel ptr)
{
  if (ptr == NULL) return DataChannelInterface::kClosed;
  DataChannelInterface* dc = CastPtr(DataChannelInterface, ptr);
  return dc->state();
}

int WebRTC_DataChannel_Send(
  DataChannel ptr,
  void* bytes, int nbytes)
{
  if (ptr == NULL) return DataChannelInterface::kClosed;
  DataChannelInterface* dc = CastPtr(DataChannelInterface, ptr);

  talk_base::Buffer data(bytes, size_t(nbytes));
  DataBuffer buffer(data, false);
  return dc->Send(buffer);
}

void WebRTC_DataChannel_Close(
  DataChannel ptr)
{
  if (ptr == NULL) return;
  DataChannelInterface* dc = CastPtr(DataChannelInterface, ptr);
  return dc->Close();
}

IceCandidate WebRTC_IceCandidate_Parse(char* id, int label, char* candidate)
{
  std::string _candidate = candidate;
  std::string _id = id;
  return CreateIceCandidate(_id, label, _candidate, NULL);
}

void WebRTC_IceCandidate_Free(IceCandidate ptr)
{
  if (ptr == NULL) return;
  IceCandidateInterface* c = CastPtr(IceCandidateInterface, ptr);
  delete c;
}

char* WebRTC_IceCandidate_SDP(IceCandidate ptr)
{
  std::string out = "";

  if (ptr != NULL) {
    IceCandidateInterface* c = CastPtr(IceCandidateInterface, ptr);
    c->ToString(&out);
  }

  return strdup(out.c_str());
}

char* WebRTC_IceCandidate_ID(IceCandidate ptr)
{
  std::string out = "";

  if (ptr != NULL) {
    IceCandidateInterface* c = CastPtr(IceCandidateInterface, ptr);
    out = c->sdp_mid();
  }

  return strdup(out.c_str());
}

int WebRTC_IceCandidate_Index(IceCandidate ptr)
{
  if (ptr != NULL) {
    IceCandidateInterface* c = CastPtr(IceCandidateInterface, ptr);
    return c->sdp_mline_index();
  }

  return -1;
}

} // extern "C"
