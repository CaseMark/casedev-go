// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// MediaV1TranscriptService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaV1TranscriptService] method instead.
type MediaV1TranscriptService struct {
	Options []option.RequestOption
	// Transcript retrieval and captioned media clip generation
	Search *MediaV1TranscriptSearchService
	// Transcript retrieval and captioned media clip generation
	Retrieve *MediaV1TranscriptRetrieveService
}

// NewMediaV1TranscriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaV1TranscriptService(opts ...option.RequestOption) (r *MediaV1TranscriptService) {
	r = &MediaV1TranscriptService{}
	r.Options = opts
	r.Search = NewMediaV1TranscriptSearchService(opts...)
	r.Retrieve = NewMediaV1TranscriptRetrieveService(opts...)
	return
}
