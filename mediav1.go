// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// MediaV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaV1Service] method instead.
type MediaV1Service struct {
	Options []option.RequestOption
	// Transcript retrieval and captioned media clip generation
	Clips       *MediaV1ClipService
	Transcripts *MediaV1TranscriptService
}

// NewMediaV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMediaV1Service(opts ...option.RequestOption) (r *MediaV1Service) {
	r = &MediaV1Service{}
	r.Options = opts
	r.Clips = NewMediaV1ClipService(opts...)
	r.Transcripts = NewMediaV1TranscriptService(opts...)
	return
}
