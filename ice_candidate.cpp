#include <cstdlib>
#include <cstring>
#include <string>

#include "talk/app/webrtc/jsep.h"
#include "talk/app/webrtc/peerconnectioninterface.h"

#include "constants.h"
#include "ice_candidate.h"

extern "C" {
#include "_cgo_export.h"

using talk_base::scoped_refptr;
using talk_base::RefCountedObject;
using webrtc::PeerConnectionInterface;

char* c_PeerConnection_AddIceCandidate(
  void* _pc, void* _candidate)
{
  if (_pc == NULL || _candidate == NULL) return strdup(errNillNotAllowed);
  PeerConnectionInterface* pc = (PeerConnectionInterface*)_pc;

  std::string                    sdp_mid;
  int                            sdp_mline_index;
  std::string                    sdp;
  webrtc::SdpParseError          error;
  webrtc::IceCandidateInterface* candidate;

  sdp_mid         = go_IceCandidate_GetSdpMid(_candidate);
  sdp_mline_index = go_IceCandidate_GetSdpMlineIndex(_candidate);
  sdp             = go_IceCandidate_GetCandidate(_candidate);

  candidate = webrtc::CreateIceCandidate(sdp_mid, sdp_mline_index, sdp, &error);
  if (candidate == NULL) {
    return strdup(error.description.c_str());
  }

  bool ok = pc->AddIceCandidate(candidate);
  if (!ok) {
    delete candidate;
    return strdup(errFailedToAddIceCandidate);
  }

  delete candidate;
  return NULL;
}

} // extern "C"
