#ifndef _WRAPPER
#define _WRAPPER

#include <stdint.h>
#include <stdlib.h>
#include "ref.h"

#ifdef __cplusplus
extern "C" {
#endif

extern const char* WebRTC_Version();
extern int WebRTC_InitializeSSL();
extern int WebRTC_CleanupSSL();

typedef void* Factory;
typedef void* MediaStream;
typedef void* PeerConnection;
typedef void* IceCandidate;

extern Factory WebRTC_PeerConnectionFactory_Create();
extern void WebRTC_PeerConnectionFactory_Free(Factory ptr);

extern MediaStream WebRTC_CreateMediaStreamWithLabel(Factory ptr, char* clabel);
extern void WebRTC_MediaStream_Free(MediaStream ptr);

extern PeerConnection WebRTC_PeerConnection_Create(
  Factory factory,
  void* servers,     int nservers,
  void* constraints, int nconstraints,
  Ref observerRef);
extern void WebRTC_PeerConnection_Free(PeerConnection ptr);
extern int WebRTC_PeerConnection_UpdateICE(
  PeerConnection ptr,
  void* servers,     int nservers,
  void* constraints, int nconstraints);

#ifdef __cplusplus
} // extern "C"
#endif

#endif
