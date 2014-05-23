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

typedef void* Factory;
typedef void* MediaStream;
typedef void* PeerConnection;
typedef void* ICEServer;
typedef void* MediaConstraints;
typedef void* PeerConnectionObserver;

extern Factory WebRTC_PeerConnectionFactory_Create();
extern void WebRTC_PeerConnectionFactory_Free(Factory ptr);

extern MediaStream WebRTC_CreateMediaStreamWithLabel(Factory ptr, char* clabel);
extern void WebRTC_MediaStream_Free(MediaStream ptr);

extern PeerConnection WebRTC_PeerConnection_Create(
  Factory factory, int nservers, ICEServer* servers,
  MediaConstraints constraints, PeerConnectionObserver observerImpl);
extern void WebRTC_PeerConnection_Free(PeerConnection ptr);


#ifdef __cplusplus
} // extern "C"
#endif

#endif
