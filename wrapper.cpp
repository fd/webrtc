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
#include "talk/app/webrtc/peerconnectioninterface.h"
#include "talk/app/webrtc/videosourceinterface.h"
#include "talk/app/webrtc/videotrack.h"
#include "talk/base/logging.h"
#include "talk/base/ssladapter.h"

#include "wrapper.h"
#include "session_description.h"
#include "media_constraints_prv.h"

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
    c_RTCPeerConnectionObserver_OnIceConnectionChange(_ref, (int)new_state);
  }

  // Called any time the IceGatheringState changes
  void OnIceGatheringChange(
      PeerConnectionInterface::IceGatheringState new_state)
  {
    c_RTCPeerConnectionObserver_OnIceGatheringChange(_ref, (int)new_state);
  }

  // New Ice candidate have been found.
  void OnIceCandidate(const IceCandidateInterface* candidate)
  {
    std::string sdp_mid         = candidate->sdp_mid();
    int         sdp_mline_index = candidate->sdp_mline_index();
    std::string sdp;

    if (!candidate->ToString(&sdp)) {
      return;
    }

    void* _candidate = go_IceCandidate_ToGo(
      (char*)sdp_mid.c_str(),
      sdp_mline_index,
      (char*)sdp.c_str()
    );

    c_RTCPeerConnectionObserver_OnIceCandidate(_ref, _candidate);
  }

  ~RTCPeerConnectionObserver() {
    go_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
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

} // extern "C"
