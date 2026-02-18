// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// VoiceTranscriptionService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceTranscriptionService] method instead.
type VoiceTranscriptionService struct {
	Options []option.RequestOption
}

// NewVoiceTranscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVoiceTranscriptionService(opts ...option.RequestOption) (r *VoiceTranscriptionService) {
	r = &VoiceTranscriptionService{}
	r.Options = opts
	return
}

// Creates an asynchronous transcription job for audio files. Supports two modes:
//
// **Vault-based (recommended)**: Pass `vault_id` and `object_id` to transcribe
// audio from your vault. The transcript will automatically be saved back to the
// vault when complete.
//
// **Direct URL (legacy)**: Pass `audio_url` for direct transcription without
// automatic storage.
func (r *VoiceTranscriptionService) New(ctx context.Context, body VoiceTranscriptionNewParams, opts ...option.RequestOption) (res *VoiceTranscriptionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice/transcription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the status and result of an audio transcription job. For vault-based
// jobs, returns status and result_object_id when complete. For legacy direct URL
// jobs, returns the full transcription data.
func (r *VoiceTranscriptionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VoiceTranscriptionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("voice/transcription/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a transcription job. For managed vault jobs (tr\_\*), also removes local
// job records and managed transcript result objects. Idempotent: returns success
// if already deleted.
func (r *VoiceTranscriptionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("voice/transcription/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type VoiceTranscriptionNewResponse struct {
	// Unique transcription job ID
	ID string `json:"id"`
	// Source audio object ID (only for vault-based transcription)
	SourceObjectID string `json:"source_object_id"`
	// Current status of the transcription job
	Status VoiceTranscriptionNewResponseStatus `json:"status"`
	// Vault ID (only for vault-based transcription)
	VaultID string                            `json:"vault_id"`
	JSON    voiceTranscriptionNewResponseJSON `json:"-"`
}

// voiceTranscriptionNewResponseJSON contains the JSON metadata for the struct
// [VoiceTranscriptionNewResponse]
type voiceTranscriptionNewResponseJSON struct {
	ID             apijson.Field
	SourceObjectID apijson.Field
	Status         apijson.Field
	VaultID        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VoiceTranscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceTranscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the transcription job
type VoiceTranscriptionNewResponseStatus string

const (
	VoiceTranscriptionNewResponseStatusQueued     VoiceTranscriptionNewResponseStatus = "queued"
	VoiceTranscriptionNewResponseStatusProcessing VoiceTranscriptionNewResponseStatus = "processing"
	VoiceTranscriptionNewResponseStatusCompleted  VoiceTranscriptionNewResponseStatus = "completed"
	VoiceTranscriptionNewResponseStatusError      VoiceTranscriptionNewResponseStatus = "error"
)

func (r VoiceTranscriptionNewResponseStatus) IsKnown() bool {
	switch r {
	case VoiceTranscriptionNewResponseStatusQueued, VoiceTranscriptionNewResponseStatusProcessing, VoiceTranscriptionNewResponseStatusCompleted, VoiceTranscriptionNewResponseStatusError:
		return true
	}
	return false
}

type VoiceTranscriptionGetResponse struct {
	// Unique transcription job ID
	ID string `json:"id,required"`
	// Current status of the transcription job
	Status VoiceTranscriptionGetResponseStatus `json:"status,required"`
	// Duration of the audio file in seconds
	AudioDuration float64 `json:"audio_duration"`
	// Overall confidence score (0-100)
	Confidence float64 `json:"confidence"`
	// Error message (only present when status is failed)
	Error string `json:"error"`
	// Result transcript object ID (vault-based jobs, when completed)
	ResultObjectID string `json:"result_object_id"`
	// Source audio object ID (vault-based jobs only)
	SourceObjectID string `json:"source_object_id"`
	// Full transcription text (legacy direct URL jobs only)
	Text string `json:"text"`
	// Vault ID (vault-based jobs only)
	VaultID string `json:"vault_id"`
	// Number of words in the transcript
	WordCount int64 `json:"word_count"`
	// Word-level timestamps (legacy direct URL jobs only)
	Words []interface{}                     `json:"words"`
	JSON  voiceTranscriptionGetResponseJSON `json:"-"`
}

// voiceTranscriptionGetResponseJSON contains the JSON metadata for the struct
// [VoiceTranscriptionGetResponse]
type voiceTranscriptionGetResponseJSON struct {
	ID             apijson.Field
	Status         apijson.Field
	AudioDuration  apijson.Field
	Confidence     apijson.Field
	Error          apijson.Field
	ResultObjectID apijson.Field
	SourceObjectID apijson.Field
	Text           apijson.Field
	VaultID        apijson.Field
	WordCount      apijson.Field
	Words          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VoiceTranscriptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceTranscriptionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the transcription job
type VoiceTranscriptionGetResponseStatus string

const (
	VoiceTranscriptionGetResponseStatusQueued     VoiceTranscriptionGetResponseStatus = "queued"
	VoiceTranscriptionGetResponseStatusProcessing VoiceTranscriptionGetResponseStatus = "processing"
	VoiceTranscriptionGetResponseStatusCompleted  VoiceTranscriptionGetResponseStatus = "completed"
	VoiceTranscriptionGetResponseStatusFailed     VoiceTranscriptionGetResponseStatus = "failed"
)

func (r VoiceTranscriptionGetResponseStatus) IsKnown() bool {
	switch r {
	case VoiceTranscriptionGetResponseStatusQueued, VoiceTranscriptionGetResponseStatusProcessing, VoiceTranscriptionGetResponseStatusCompleted, VoiceTranscriptionGetResponseStatusFailed:
		return true
	}
	return false
}

type VoiceTranscriptionNewParams struct {
	// URL of the audio file to transcribe (legacy mode, no auto-storage)
	AudioURL param.Field[string] `json:"audio_url"`
	// Automatically extract key phrases and topics
	AutoHighlights param.Field[bool] `json:"auto_highlights"`
	// How much to boost custom vocabulary
	BoostParam param.Field[VoiceTranscriptionNewParamsBoostParam] `json:"boost_param"`
	// Enable content moderation and safety labeling
	ContentSafetyLabels param.Field[bool] `json:"content_safety_labels"`
	// Output format for the transcript when using vault mode
	Format param.Field[VoiceTranscriptionNewParamsFormat] `json:"format"`
	// Format text with proper capitalization
	FormatText param.Field[bool] `json:"format_text"`
	// Language code (e.g., 'en_us', 'es', 'fr'). If not specified, language will be
	// auto-detected
	LanguageCode param.Field[string] `json:"language_code"`
	// Enable automatic language detection
	LanguageDetection param.Field[bool] `json:"language_detection"`
	// Object ID of the audio file in the vault (use with vault_id)
	ObjectID param.Field[string] `json:"object_id"`
	// Add punctuation to the transcript
	Punctuate param.Field[bool] `json:"punctuate"`
	// Enable speaker identification and labeling
	SpeakerLabels param.Field[bool] `json:"speaker_labels"`
	// Expected number of speakers (improves accuracy when known)
	SpeakersExpected param.Field[int64] `json:"speakers_expected"`
	// Priority-ordered speech models to use
	SpeechModels param.Field[[]string] `json:"speech_models"`
	// Vault ID containing the audio file (use with object_id)
	VaultID param.Field[string] `json:"vault_id"`
	// Custom vocabulary words to boost (e.g., legal terms)
	WordBoost param.Field[[]string] `json:"word_boost"`
}

func (r VoiceTranscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much to boost custom vocabulary
type VoiceTranscriptionNewParamsBoostParam string

const (
	VoiceTranscriptionNewParamsBoostParamLow     VoiceTranscriptionNewParamsBoostParam = "low"
	VoiceTranscriptionNewParamsBoostParamDefault VoiceTranscriptionNewParamsBoostParam = "default"
	VoiceTranscriptionNewParamsBoostParamHigh    VoiceTranscriptionNewParamsBoostParam = "high"
)

func (r VoiceTranscriptionNewParamsBoostParam) IsKnown() bool {
	switch r {
	case VoiceTranscriptionNewParamsBoostParamLow, VoiceTranscriptionNewParamsBoostParamDefault, VoiceTranscriptionNewParamsBoostParamHigh:
		return true
	}
	return false
}

// Output format for the transcript when using vault mode
type VoiceTranscriptionNewParamsFormat string

const (
	VoiceTranscriptionNewParamsFormatJson VoiceTranscriptionNewParamsFormat = "json"
	VoiceTranscriptionNewParamsFormatText VoiceTranscriptionNewParamsFormat = "text"
)

func (r VoiceTranscriptionNewParamsFormat) IsKnown() bool {
	switch r {
	case VoiceTranscriptionNewParamsFormatJson, VoiceTranscriptionNewParamsFormatText:
		return true
	}
	return false
}
