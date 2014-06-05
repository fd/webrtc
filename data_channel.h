#ifndef _DATA_CHANNEL
#define _DATA_CHANNEL
#ifdef __cplusplus
extern "C" {
#endif

#include "ref.h"

extern void*  c_DataChannel_Create(void* _pc, char* label, Ref _ref);
extern void   c_DataChannel_Accept(void* _dc, Ref _ref);
extern void   c_DataChannel_Free(void* _dc);
extern int    c_DataChannel_State(void* _dc);
extern int    c_DataChannel_Send(void* _dc, void* bytes, int nbytes);
extern void   c_DataChannel_Close(void* _dc);

#ifdef __cplusplus
} // extern "C"
#endif
#endif
