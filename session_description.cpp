#include <cstdlib>
#include <cstring>
#include <string>

#include "talk/app/webrtc/jsep.h"
#include "talk/app/webrtc/peerconnectioninterface.h"

#include "ref.h"
#include "session_description.h"
#include "media_constraints_prv.h"

const char* errNillNotAllowed = "nil arguments are not allowed";

extern "C" {
#include "_cgo_export.h"

using talk_base::scoped_refptr;
using talk_base::RefCountedObject;
using webrtc::PeerConnectionInterface;

void* c_SessionDescription_ToGo(const webrtc::SessionDescriptionInterface* desc);
webrtc::SessionDescriptionInterface* c_SessionDescription_ToWebRTC(
  void* desc, webrtc::SdpParseError* error);


class setSessionDescriptionObserver : public webrtc::SetSessionDescriptionObserver {
public:

  static scoped_refptr<setSessionDescriptionObserver> Create(Ref ref) {
    RefCountedObject<setSessionDescriptionObserver>* implementation =
         new RefCountedObject<setSessionDescriptionObserver>(ref);
    return implementation;
  }

  setSessionDescriptionObserver(Ref ref) {
    _ref = ref;
  }

  void OnSuccess() {
    go_SetSessionDescription_OnSuccess(_ref);
  }

  void OnFailure(const std::string& error) {
    go_SetSessionDescription_OnFailure(_ref, (char*)error.c_str());
  }

protected:
  ~setSessionDescriptionObserver() {
    go_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
};

class createSessionDescriptionObserver : public webrtc::CreateSessionDescriptionObserver {
public:

  static scoped_refptr<createSessionDescriptionObserver> Create(Ref ref) {
    RefCountedObject<createSessionDescriptionObserver>* implementation =
         new RefCountedObject<createSessionDescriptionObserver>(ref);
    return implementation;
  }

  createSessionDescriptionObserver(Ref ref) {
    _ref = ref;
  }

  void OnSuccess(webrtc::SessionDescriptionInterface* desc) {
    go_CreateSessionDescription_OnSuccess(_ref, c_SessionDescription_ToGo(desc));
  }

  void OnFailure(const std::string& error) {
    go_CreateSessionDescription_OnFailure(_ref, (char*)error.c_str());
  }

protected:
  ~createSessionDescriptionObserver() {
    go_Ref_Unregister(_ref);
  }

private:
  Ref _ref;
};


void* c_PeerConnection_GetLocalDescription(void* _pc)
{
  if (_pc == NULL) return NULL;
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;
  return c_SessionDescription_ToGo(pc->local_description());
}

void* c_PeerConnection_GetRemoteDescription(void* _pc)
{
  if (_pc == NULL) return NULL;
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;
  return c_SessionDescription_ToGo(pc->remote_description());
}


char* c_PeerConnection_SetLocalDescription(
  void* _pc, Ref _ref, void* _sd)
{
  if (_pc == NULL || _sd == NULL) return strdup(errNillNotAllowed);

  webrtc::SdpParseError error;
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;
  webrtc::SessionDescriptionInterface* sd;

  sd = c_SessionDescription_ToWebRTC(_sd, &error);
  if (sd == NULL) {
    return strdup(error.description.c_str());
  }

  scoped_refptr<setSessionDescriptionObserver> observer =
    setSessionDescriptionObserver::Create(_ref);

  pc->SetLocalDescription(observer, sd);
  return NULL;
}

char* c_PeerConnection_SetRemoteDescription(
  void* _pc, Ref _ref, void* _sd)
{
  if (_pc == NULL || _sd == NULL) return strdup(errNillNotAllowed);

  webrtc::SdpParseError error;
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;
  webrtc::SessionDescriptionInterface* sd;

  sd = c_SessionDescription_ToWebRTC(_sd, &error);
  if (sd == NULL) {
    return strdup(error.description.c_str());
  }

  scoped_refptr<setSessionDescriptionObserver> observer =
    setSessionDescriptionObserver::Create(_ref);

  pc->SetRemoteDescription(observer, sd);
  return NULL;
}

void c_PeerConnection_CreateOffer(
  void* _pc, Ref _ref,
  void* constraints, int nconstraints)
{
  if (_pc == NULL) return;
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;
  void** rconstraints = (void**)constraints;

  scoped_refptr<createSessionDescriptionObserver> observer =
    createSessionDescriptionObserver::Create(_ref);

  pc->CreateOffer(observer,
    new MediaConstraints(rconstraints, nconstraints));
}

void c_PeerConnection_CreateAnswer(
  void* _pc, Ref _ref,
  void* constraints, int nconstraints)
{
  if (_pc == NULL) return;
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;
  void** rconstraints = (void**)constraints;

  scoped_refptr<createSessionDescriptionObserver> observer =
    createSessionDescriptionObserver::Create(_ref);

  pc->CreateAnswer(observer,
    new MediaConstraints(rconstraints, nconstraints));
}

void* c_SessionDescription_ToGo(const webrtc::SessionDescriptionInterface* desc)
{
  if (desc == NULL) return NULL;

  std::string sdp;
  std::string type = desc->type();

  if (!desc->ToString(&sdp)) return NULL;

  return go_SessionDescription_New((char*)type.c_str(), (char*)sdp.c_str());
}

webrtc::SessionDescriptionInterface* c_SessionDescription_ToWebRTC(
  void* desc, webrtc::SdpParseError* error)
{
  std::string type = go_SessionDescription_GetType(desc);
  std::string sdp  = go_SessionDescription_GetSdp(desc);
  return webrtc::CreateSessionDescription(type, sdp, error);
}


} // extern "C"
