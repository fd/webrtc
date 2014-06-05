#ifndef _SESSION_DESCRIPTION
#define _SESSION_DESCRIPTION
#ifdef __cplusplus
extern "C" {
#endif

#include "ref.h"

extern void* c_PeerConnection_GetLocalDescription(
  void* _pc);
extern void* c_PeerConnection_GetRemoteDescription(
  void* _pc);

extern char* c_PeerConnection_SetLocalDescription(
  void* _pc, Ref _ref, void* _sd);
extern char* c_PeerConnection_SetRemoteDescription(
  void* _pc, Ref _ref, void* _sd);

extern void c_PeerConnection_CreateOffer(
  void* _pc, Ref _ref,
  void* constraints, int nconstraints);
extern void c_PeerConnection_CreateAnswer(
  void* _pc, Ref _ref,
  void* constraints, int nconstraints);

#ifdef __cplusplus
} // extern "C"
#endif
#endif
