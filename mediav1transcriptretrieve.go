// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// MediaV1TranscriptRetrieveService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaV1TranscriptRetrieveService] method instead.
type MediaV1TranscriptRetrieveService struct {
	Options []option.RequestOption
}

// NewMediaV1TranscriptRetrieveService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMediaV1TranscriptRetrieveService(opts ...option.RequestOption) (r *MediaV1TranscriptRetrieveService) {
	r = &MediaV1TranscriptRetrieveService{}
	r.Options = opts
	return
}

// Retrieves the full transcript text for a vault transcript object or an
// audio/video source object with a completed transcription job. When object_id is
// a source media object, access to that source object grants access to its
// generated transcript artifact.
func (r *MediaV1TranscriptRetrieveService) New(ctx context.Context, body MediaV1TranscriptRetrieveNewParams, opts ...option.RequestOption) (res *MediaV1TranscriptRetrieveNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media/v1/transcripts/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MediaV1TranscriptRetrieveNewResponse struct {
	// Requested object ID.
	ObjectID string                                     `json:"object_id" api:"required"`
	Status   MediaV1TranscriptRetrieveNewResponseStatus `json:"status" api:"required"`
	// Full transcript text.
	Text          string `json:"text" api:"required"`
	VaultID       string `json:"vault_id" api:"required"`
	AudioDuration int64  `json:"audio_duration"`
	Confidence    int64  `json:"confidence"`
	Filename      string `json:"filename"`
	// Source media object ID when known.
	SourceObjectID string `json:"source_object_id"`
	// Transcript object ID when known.
	TranscriptObjectID string `json:"transcript_object_id"`
	// Transcription job ID when known.
	TranscriptionJobID string                                   `json:"transcription_job_id"`
	WordCount          int64                                    `json:"word_count"`
	JSON               mediaV1TranscriptRetrieveNewResponseJSON `json:"-"`
}

// mediaV1TranscriptRetrieveNewResponseJSON contains the JSON metadata for the
// struct [MediaV1TranscriptRetrieveNewResponse]
type mediaV1TranscriptRetrieveNewResponseJSON struct {
	ObjectID           apijson.Field
	Status             apijson.Field
	Text               apijson.Field
	VaultID            apijson.Field
	AudioDuration      apijson.Field
	Confidence         apijson.Field
	Filename           apijson.Field
	SourceObjectID     apijson.Field
	TranscriptObjectID apijson.Field
	TranscriptionJobID apijson.Field
	WordCount          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MediaV1TranscriptRetrieveNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mediaV1TranscriptRetrieveNewResponseJSON) RawJSON() string {
	return r.raw
}

type MediaV1TranscriptRetrieveNewResponseStatus string

const (
	MediaV1TranscriptRetrieveNewResponseStatusCompleted MediaV1TranscriptRetrieveNewResponseStatus = "completed"
)

func (r MediaV1TranscriptRetrieveNewResponseStatus) IsKnown() bool {
	switch r {
	case MediaV1TranscriptRetrieveNewResponseStatusCompleted:
		return true
	}
	return false
}

type MediaV1TranscriptRetrieveNewParams struct {
	// Object ID for either the source audio/video file or transcript artifact.
	ObjectID param.Field[string] `json:"object_id" api:"required"`
	// Vault ID containing the source media or transcript object.
	VaultID param.Field[string] `json:"vault_id" api:"required"`
	// Alternative nested transcript object reference.
	Transcript param.Field[MediaV1TranscriptRetrieveNewParamsTranscript] `json:"transcript"`
}

func (r MediaV1TranscriptRetrieveNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Alternative nested transcript object reference.
type MediaV1TranscriptRetrieveNewParamsTranscript struct {
	ObjectID param.Field[string] `json:"object_id"`
	VaultID  param.Field[string] `json:"vault_id"`
}

func (r MediaV1TranscriptRetrieveNewParamsTranscript) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
