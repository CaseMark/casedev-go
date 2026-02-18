// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// VoiceService contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceService] method instead.
type VoiceService struct {
	Options       []option.RequestOption
	Streaming     *VoiceStreamingService
	Transcription *VoiceTranscriptionService
	V1            *VoiceV1Service
}

// NewVoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVoiceService(opts ...option.RequestOption) (r *VoiceService) {
	r = &VoiceService{}
	r.Options = opts
	r.Streaming = NewVoiceStreamingService(opts...)
	r.Transcription = NewVoiceTranscriptionService(opts...)
	r.V1 = NewVoiceV1Service(opts...)
	return
}
