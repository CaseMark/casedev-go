// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// MediaV1TranscriptSearchService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaV1TranscriptSearchService] method instead.
type MediaV1TranscriptSearchService struct {
	Options []option.RequestOption
}

// NewMediaV1TranscriptSearchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaV1TranscriptSearchService(opts ...option.RequestOption) (r *MediaV1TranscriptSearchService) {
	r = &MediaV1TranscriptSearchService{}
	r.Options = opts
	return
}

// Search transcript words
func (r *MediaV1TranscriptSearchService) New(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "media/v1/transcripts/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}
