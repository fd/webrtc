#ifndef _V8_WARP_H_
#define _V8_WARP_H_

#include <stdint.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

extern const char* WebRTC_Version();
extern int WebRTC_InitializeSSL();
extern int WebRTC_CleanupSSL();

extern void* WebRTC_PeerConnectionFactory_Create();
extern void WebRTC_PeerConnectionFactory_Free(void* ptr);

extern void* WebRTC_PeerConnectionFactory_CreateMediaStreamWithLabel(void* ptr, char* clabel);
extern void WebRTC_MediaStream_Free(void* ptr);

#ifdef __cplusplus
} // extern "C"
#endif

#endif
