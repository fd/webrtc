#include "talk/app/webrtc/datachannelinterface.h"
#include "talk/app/webrtc/peerconnectioninterface.h"

#include "ref.h"

extern "C" {
#include "_cgo_export.h"

using talk_base::scoped_refptr;
using talk_base::RefCountedObject;

class dataChannelObserver : public webrtc::DataChannelObserver {

public:
  dataChannelObserver(Ref ref) {
    _ref = ref;
  }

  ~dataChannelObserver() {
    go_Ref_Unregister(_ref);
  }

  void OnStateChange()
  {
    go_DataChannel_OnStateChange(_ref);
  }

  void OnMessage(const webrtc::DataBuffer& buffer)
  {
    go_DataChannel_OnMessage(_ref, (void*)buffer.data.data(), buffer.size());
  }

private:
  Ref _ref;
};


void* c_DataChannel_Create(void* _pc, char* label, Ref _ref) {
  if (_pc == NULL || label == NULL) return NULL;
  webrtc::PeerConnectionInterface* pc = (webrtc::PeerConnectionInterface*)_pc;

  std::string _label = label;
  if (label != NULL) free(label);

  webrtc::DataChannelInit* options = new webrtc::DataChannelInit();
  options->ordered = go_DataChannelOptions_Ordered(_ref) == 1 ? true : false;
  options->maxRetransmitTime = go_DataChannelOptions_MaxRetransmitTime(_ref);
  options->maxRetransmits = go_DataChannelOptions_MaxRetransmits(_ref);
  options->protocol = go_DataChannelOptions_Protocol(_ref);
  options->negotiated = go_DataChannelOptions_Negotiated(_ref) == 1 ? true : false;
  options->id = go_DataChannelOptions_Id(_ref);

  scoped_refptr<webrtc::DataChannelInterface> dc =
    pc->CreateDataChannel(_label, options);
  if (dc == NULL) {
    return NULL;
  }

  dataChannelObserver* observer = new dataChannelObserver(_ref);
  dc->RegisterObserver(observer);

  webrtc::DataChannelInterface* _dc = dc.get();
  _dc->AddRef();
  return _dc;
}

void c_DataChannel_Accept(void* _dc, Ref _ref)
{
  if (_dc == NULL) return;
  webrtc::DataChannelInterface* dc = (webrtc::DataChannelInterface*)_dc;

  dataChannelObserver* observer = new dataChannelObserver(_ref);
  dc->RegisterObserver(observer);
}

void c_DataChannel_Free(void* _dc) {
  if (_dc == NULL) return;
  webrtc::DataChannelInterface* dc = (webrtc::DataChannelInterface*)_dc;
  dc->Release();
}

int c_DataChannel_State(void* _dc)
{
  if (_dc == NULL) return webrtc::DataChannelInterface::kClosed;
  webrtc::DataChannelInterface* dc = (webrtc::DataChannelInterface*)_dc;
  return dc->state();
}

int c_DataChannel_Send(void* _dc, void* bytes, int nbytes)
{
  if (_dc == NULL) return 0;
  webrtc::DataChannelInterface* dc = (webrtc::DataChannelInterface*)_dc;

  talk_base::Buffer data(bytes, size_t(nbytes));
  webrtc::DataBuffer buffer(data, false);
  return dc->Send(buffer);
}

void c_DataChannel_Close(void* _dc)
{
  if (_dc == NULL) return;
  webrtc::DataChannelInterface* dc = (webrtc::DataChannelInterface*)_dc;
  dc->Close();
}

} // extern "C"
