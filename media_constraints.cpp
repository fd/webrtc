
#include "media_constraints_prv.h"

extern "C" {
#include "_cgo_export.h"

MediaConstraints::MediaConstraints(void** ptr, int len) {
  for (int i=0; i<len; i++) {
    if (go_MediaConstraint_Optional(ptr[i]) > 0) {
      optional.push_back(webrtc::MediaConstraintsInterface::Constraint(
        go_MediaConstraint_Key(ptr[i]),
        go_MediaConstraint_Value(ptr[i])
      ));
    } else {
      mandatory.push_back(webrtc::MediaConstraintsInterface::Constraint(
        go_MediaConstraint_Key(ptr[i]),
        go_MediaConstraint_Value(ptr[i])
      ));
    }
  }
}

const webrtc::MediaConstraintsInterface::Constraints& MediaConstraints::GetMandatory() const {
  return mandatory;
}

const webrtc::MediaConstraintsInterface::Constraints& MediaConstraints::GetOptional() const {
  return optional;
}

MediaConstraints::~MediaConstraints() {

}

} // extern "C"
