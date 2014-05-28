#ifndef _WEBRTC_WARP_H_
#define _WEBRTC_WARP_H_

#include <stdint.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

extern const char* WebRTC_Version();
extern int WebRTC_InitializeSSL();
extern int WebRTC_CleanupSSL();

typedef uint64_t Ref;
typedef void* Factory;
typedef void* MediaStream;
typedef void* PeerConnection;
typedef void* SessionDescription;

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
extern void WebRTC_PeerConnection_CreateOffer(
  PeerConnection ptr,
  Ref observerRef,
  void* constraints, int nconstraints);
extern void WebRTC_PeerConnection_SetLocalDescription(
  PeerConnection ptr,
  Ref observerRef,
  SessionDescription desc);
extern void WebRTC_PeerConnection_SetRemoteDescription(
  PeerConnection ptr,
  Ref observerRef,
  SessionDescription desc);

extern char* WebRTC_SessionDescription_String(SessionDescription ptr);
extern void WebRTC_SessionDescription_Free(SessionDescription ptr);


#ifdef __cplusplus
} // extern "C"
#endif

#endif
