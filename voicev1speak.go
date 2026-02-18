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

// VoiceV1SpeakService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceV1SpeakService] method instead.
type VoiceV1SpeakService struct {
	Options []option.RequestOption
}

// NewVoiceV1SpeakService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceV1SpeakService(opts ...option.RequestOption) (r *VoiceV1SpeakService) {
	r = &VoiceV1SpeakService{}
	r.Options = opts
	return
}

// Convert text to natural-sounding audio using ElevenLabs voices. Ideal for
// creating audio summaries of legal documents, client presentations, or
// accessibility features. Supports multiple languages and voice customization.
func (r *VoiceV1SpeakService) New(ctx context.Context, body VoiceV1SpeakNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "voice/v1/speak"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VoiceV1SpeakNewParams struct {
	// Text to convert to speech
	Text param.Field[string] `json:"text,required"`
	// Apply automatic text normalization
	ApplyTextNormalization param.Field[bool] `json:"apply_text_normalization"`
	// Enable request logging
	EnableLogging param.Field[bool] `json:"enable_logging"`
	// Language code for multilingual models
	LanguageCode param.Field[string] `json:"language_code"`
	// ElevenLabs model ID
	ModelID param.Field[VoiceV1SpeakNewParamsModelID] `json:"model_id"`
	// Next context for better pronunciation
	NextText param.Field[string] `json:"next_text"`
	// Optimize for streaming latency (0-4)
	OptimizeStreamingLatency param.Field[int64] `json:"optimize_streaming_latency"`
	// Audio output format
	OutputFormat param.Field[VoiceV1SpeakNewParamsOutputFormat] `json:"output_format"`
	// Previous context for better pronunciation
	PreviousText param.Field[string] `json:"previous_text"`
	// Seed for reproducible generation
	Seed param.Field[int64] `json:"seed"`
	// ElevenLabs voice ID (defaults to Rachel - professional, clear)
	VoiceID param.Field[string] `json:"voice_id"`
	// Voice customization settings
	VoiceSettings param.Field[VoiceV1SpeakNewParamsVoiceSettings] `json:"voice_settings"`
}

func (r VoiceV1SpeakNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ElevenLabs model ID
type VoiceV1SpeakNewParamsModelID string

const (
	VoiceV1SpeakNewParamsModelIDElevenMultilingualV2 VoiceV1SpeakNewParamsModelID = "eleven_multilingual_v2"
	VoiceV1SpeakNewParamsModelIDElevenTurboV2        VoiceV1SpeakNewParamsModelID = "eleven_turbo_v2"
	VoiceV1SpeakNewParamsModelIDElevenMonolingualV1  VoiceV1SpeakNewParamsModelID = "eleven_monolingual_v1"
)

func (r VoiceV1SpeakNewParamsModelID) IsKnown() bool {
	switch r {
	case VoiceV1SpeakNewParamsModelIDElevenMultilingualV2, VoiceV1SpeakNewParamsModelIDElevenTurboV2, VoiceV1SpeakNewParamsModelIDElevenMonolingualV1:
		return true
	}
	return false
}

// Audio output format
type VoiceV1SpeakNewParamsOutputFormat string

const (
	VoiceV1SpeakNewParamsOutputFormatMP3_44100_128 VoiceV1SpeakNewParamsOutputFormat = "mp3_44100_128"
	VoiceV1SpeakNewParamsOutputFormatMP3_44100_192 VoiceV1SpeakNewParamsOutputFormat = "mp3_44100_192"
	VoiceV1SpeakNewParamsOutputFormatPcm16000      VoiceV1SpeakNewParamsOutputFormat = "pcm_16000"
	VoiceV1SpeakNewParamsOutputFormatPcm22050      VoiceV1SpeakNewParamsOutputFormat = "pcm_22050"
	VoiceV1SpeakNewParamsOutputFormatPcm24000      VoiceV1SpeakNewParamsOutputFormat = "pcm_24000"
	VoiceV1SpeakNewParamsOutputFormatPcm44100      VoiceV1SpeakNewParamsOutputFormat = "pcm_44100"
)

func (r VoiceV1SpeakNewParamsOutputFormat) IsKnown() bool {
	switch r {
	case VoiceV1SpeakNewParamsOutputFormatMP3_44100_128, VoiceV1SpeakNewParamsOutputFormatMP3_44100_192, VoiceV1SpeakNewParamsOutputFormatPcm16000, VoiceV1SpeakNewParamsOutputFormatPcm22050, VoiceV1SpeakNewParamsOutputFormatPcm24000, VoiceV1SpeakNewParamsOutputFormatPcm44100:
		return true
	}
	return false
}

// Voice customization settings
type VoiceV1SpeakNewParamsVoiceSettings struct {
	// Similarity boost (0-1)
	SimilarityBoost param.Field[float64] `json:"similarity_boost"`
	// Voice stability (0-1)
	Stability param.Field[float64] `json:"stability"`
	// Style exaggeration (0-1)
	Style param.Field[float64] `json:"style"`
	// Enable speaker boost
	UseSpeakerBoost param.Field[bool] `json:"use_speaker_boost"`
}

func (r VoiceV1SpeakNewParamsVoiceSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
