#ifndef _MEDIA_CONSTRAINTS
#define _MEDIA_CONSTRAINTS

#include "talk/app/webrtc/mediaconstraintsinterface.h"

class MediaConstraints : public webrtc::MediaConstraintsInterface {
public:
  MediaConstraints(void** ptr, int len);
  const webrtc::MediaConstraintsInterface::Constraints& GetMandatory() const;
  const webrtc::MediaConstraintsInterface::Constraints& GetOptional() const;
  ~MediaConstraints();

private:
  webrtc::MediaConstraintsInterface::Constraints mandatory;
  webrtc::MediaConstraintsInterface::Constraints optional;
};

#endif
