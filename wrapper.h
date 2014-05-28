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
typedef void* DataChannel;
typedef void* SessionDescription;
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
extern void WebRTC_PeerConnection_CreateOffer(
  PeerConnection ptr,
  Ref observerRef,
  void* constraints, int nconstraints);
extern void WebRTC_PeerConnection_CreateAnswer(
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
extern SessionDescription WebRTC_PeerConnection_GetLocalDescription(
  PeerConnection ptr);
extern SessionDescription WebRTC_PeerConnection_GetRemoteDescription(
  PeerConnection ptr);
extern int WebRTC_PeerConnection_UpdateICE(
  PeerConnection ptr,
  void* servers,     int nservers,
  void* constraints, int nconstraints);
extern void WebRTC_PeerConnection_CreateDataChannel(
  PeerConnection ptr,
  char* label,
  SessionDescription desc);
extern int WebRTC_PeerConnection_AddIceCandidate(
  PeerConnection ptr,
  IceCandidate c);

extern SessionDescription WebRTC_SessionDescription_Parse(char* type, char* raw);
extern char* WebRTC_SessionDescription_String(SessionDescription ptr);
extern int WebRTC_SessionDescription_AddCandidate(SessionDescription ptr, IceCandidate c);
extern char* WebRTC_SessionDescription_Type(SessionDescription ptr);
extern void WebRTC_SessionDescription_Free(SessionDescription ptr);

extern DataChannel WebRTC_DataChannel_Create(PeerConnection pc, char* label, Ref dc);
extern void WebRTC_DataChannel_Free(DataChannel ptr);
extern void WebRTC_DataChannel_Accept(DataChannel ptr, Ref dc);
extern int WebRTC_DataChannel_State(DataChannel ptr);
extern int WebRTC_DataChannel_Send(DataChannel ptr, void* bytes, int nbytes);
extern void WebRTC_DataChannel_Close(DataChannel ptr);

extern IceCandidate WebRTC_IceCandidate_Parse(char* id, int label, char* candidate);
extern void WebRTC_IceCandidate_Free(IceCandidate ptr);
extern char* WebRTC_IceCandidate_SDP(IceCandidate ptr);
extern char* WebRTC_IceCandidate_ID(IceCandidate ptr);
extern int WebRTC_IceCandidate_Index(IceCandidate ptr);

#ifdef __cplusplus
} // extern "C"
#endif

#endif
