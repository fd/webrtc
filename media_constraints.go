package webrtc

import (
	"C"
	"fmt"
	"unsafe"
)

type MediaConstraint struct {
	Key      MediaConstraintKey
	Value    interface{}
	Optional bool
}

func (c *MediaConstraint) StringValue() string {
	switch v := c.Value.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return v
	default:
		return fmt.Sprintf("%s", v)
	}
}

//export c_MediaConstraint_Key
func c_MediaConstraint_Key(ptr unsafe.Pointer) *C.char {
	if ptr == nil {
		return nil
	}
	return C.CString(string((*MediaConstraint)(ptr).Key))
}

//export c_MediaConstraint_Value
func c_MediaConstraint_Value(ptr unsafe.Pointer) *C.char {
	if ptr == nil {
		return nil
	}
	return C.CString((*MediaConstraint)(ptr).StringValue())
}

//export c_MediaConstraint_Optional
func c_MediaConstraint_Optional(ptr unsafe.Pointer) C.int {
	if ptr == nil {
		return 1
	}
	if (*MediaConstraint)(ptr).Optional {
		return 1
	}
	return 0
}

type MediaConstraintKey string

const (
	MinAspectRatioConstraint = MediaConstraintKey("minAspectRatio")
	MaxAspectRatioConstraint = MediaConstraintKey("maxAspectRatio")
	MaxWidthConstraint       = MediaConstraintKey("maxWidth")
	MinWidthConstraint       = MediaConstraintKey("minWidth")
	MaxHeightConstraint      = MediaConstraintKey("maxHeight")
	MinHeightConstraint      = MediaConstraintKey("minHeight")
	MaxFrameRateConstraint   = MediaConstraintKey("maxFrameRate")
	MinFrameRateConstraint   = MediaConstraintKey("minFrameRate")

	// Constraint keys used by a local audio source.
	// These keys are google specific.
	EchoCancellationConstraint             = MediaConstraintKey("googEchoCancellation")
	ExperimentalEchoCancellationConstraint = MediaConstraintKey("googEchoCancellation2")
	AutoGainControlConstraint              = MediaConstraintKey("googAutoGainControl")
	ExperimentalAutoGainControlConstraint  = MediaConstraintKey("googAutoGainControl2")
	NoiseSuppressionConstraint             = MediaConstraintKey("googNoiseSuppression")
	ExperimentalNoiseSuppressionConstraint = MediaConstraintKey("googNoiseSuppression2")
	HighpassFilterConstraint               = MediaConstraintKey("googHighpassFilter")
	TypingNoiseDetectionConstraint         = MediaConstraintKey("googTypingNoiseDetection")
	AudioMirroringConstraint               = MediaConstraintKey("googAudioMirroring")

	// Google-specific constraint keys for a local video source
	NoiseReductionConstraint            = MediaConstraintKey("googNoiseReduction")
	LeakyBucketConstraint               = MediaConstraintKey("googLeakyBucket")
	TemporalLayeredScreencastConstraint = MediaConstraintKey("googTemporalLayeredScreencast")

	// Constraint keys for CreateOffer / CreateAnswer
	// Specified by the W3C PeerConnection spec
	OfferToReceiveVideoConstraint    = MediaConstraintKey("OfferToReceiveVideo")
	OfferToReceiveAudioConstraint    = MediaConstraintKey("OfferToReceiveAudio")
	VoiceActivityDetectionConstraint = MediaConstraintKey("VoiceActivityDetection")
	IceRestartConstraint             = MediaConstraintKey("IceRestart")
	// These keys are google specific.
	UseRtpMuxConstraint = MediaConstraintKey("googUseRtpMUX")

	// PeerConnection constraint keys.
	// Temporary pseudo-constraints used to enable DTLS-SRTP
	EnableDtlsSrtpConstraint = MediaConstraintKey("Enable DTLS-SRTP")
	// Temporary pseudo-constraints used to enable DataChannels
	EnableRtpDataChannelsConstraint = MediaConstraintKey("Enable RTP DataChannels")
	// Google-specific constraint keys.
	// Temporary pseudo-constraint for enabling DSCP through JS.
	EnableDscpConstraint = MediaConstraintKey("googDscp")
	// Constraint to enable IPv6 through JS.
	EnableIPv6Constraint = MediaConstraintKey("googIPv6")
	// Temporary constraint to enable suspend below min bitrate feature.
	EnableVideoSuspendBelowMinBitrateConstraint = MediaConstraintKey("googSuspendBelowMinBitrate")
	ImprovedWifiBweConstraint                   = MediaConstraintKey("googImprovedWifiBwe")
	ScreencastMinBitrateConstraint              = MediaConstraintKey("googScreencastMinBitrate")
	SkipEncodingUnusedStreamsConstraint         = MediaConstraintKey("googSkipEncodingUnusedStreams")
	CpuOveruseDetectionConstraint               = MediaConstraintKey("googCpuOveruseDetection")
	CpuUnderuseThresholdConstraint              = MediaConstraintKey("googCpuUnderuseThreshold")
	CpuOveruseThresholdConstraint               = MediaConstraintKey("googCpuOveruseThreshold")
	CpuOveruseEncodeUsageConstraint             = MediaConstraintKey("googCpuOveruseEncodeUsage")
	HighStartBitrateConstraint                  = MediaConstraintKey("googHighStartBitrate")
	HighBitrateConstraint                       = MediaConstraintKey("googHighBitrate")
	VeryHighBitrateConstraint                   = MediaConstraintKey("googVeryHighBitrate")
)
