// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// VoiceStreamingService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceStreamingService] method instead.
type VoiceStreamingService struct {
	Options []option.RequestOption
}

// NewVoiceStreamingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceStreamingService(opts ...option.RequestOption) (r *VoiceStreamingService) {
	r = &VoiceStreamingService{}
	r.Options = opts
	return
}

// Returns the WebSocket URL and connection details for real-time audio
// transcription. The returned URL can be used to establish a WebSocket connection
// for streaming audio data and receiving transcribed text in real-time.
//
// **Audio Requirements:**
//
// - Sample Rate: 16kHz
// - Encoding: PCM 16-bit little-endian
// - Channels: Mono (1 channel)
//
// **Pricing:** $0.01 per minute ($0.60 per hour)
func (r *VoiceStreamingService) GetURL(ctx context.Context, opts ...option.RequestOption) (res *VoiceStreamingGetURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice/streaming/url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VoiceStreamingGetURLResponse struct {
	AudioFormat VoiceStreamingGetURLResponseAudioFormat `json:"audio_format"`
	// Complete WebSocket URL with authentication token
	ConnectURL string                              `json:"connect_url"`
	Pricing    VoiceStreamingGetURLResponsePricing `json:"pricing"`
	// Connection protocol
	Protocol string `json:"protocol"`
	// Base WebSocket URL for streaming transcription
	URL  string                           `json:"url"`
	JSON voiceStreamingGetURLResponseJSON `json:"-"`
}

// voiceStreamingGetURLResponseJSON contains the JSON metadata for the struct
// [VoiceStreamingGetURLResponse]
type voiceStreamingGetURLResponseJSON struct {
	AudioFormat apijson.Field
	ConnectURL  apijson.Field
	Pricing     apijson.Field
	Protocol    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceStreamingGetURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceStreamingGetURLResponseJSON) RawJSON() string {
	return r.raw
}

type VoiceStreamingGetURLResponseAudioFormat struct {
	// Number of audio channels
	Channels int64 `json:"channels"`
	// Required audio encoding format
	Encoding string `json:"encoding"`
	// Required audio sample rate in Hz
	SampleRate int64                                       `json:"sample_rate"`
	JSON       voiceStreamingGetURLResponseAudioFormatJSON `json:"-"`
}

// voiceStreamingGetURLResponseAudioFormatJSON contains the JSON metadata for the
// struct [VoiceStreamingGetURLResponseAudioFormat]
type voiceStreamingGetURLResponseAudioFormatJSON struct {
	Channels    apijson.Field
	Encoding    apijson.Field
	SampleRate  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceStreamingGetURLResponseAudioFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceStreamingGetURLResponseAudioFormatJSON) RawJSON() string {
	return r.raw
}

type VoiceStreamingGetURLResponsePricing struct {
	// Currency for pricing
	Currency string `json:"currency"`
	// Cost per hour of transcription
	PerHour float64 `json:"per_hour"`
	// Cost per minute of transcription
	PerMinute float64                                 `json:"per_minute"`
	JSON      voiceStreamingGetURLResponsePricingJSON `json:"-"`
}

// voiceStreamingGetURLResponsePricingJSON contains the JSON metadata for the
// struct [VoiceStreamingGetURLResponsePricing]
type voiceStreamingGetURLResponsePricingJSON struct {
	Currency    apijson.Field
	PerHour     apijson.Field
	PerMinute   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceStreamingGetURLResponsePricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceStreamingGetURLResponsePricingJSON) RawJSON() string {
	return r.raw
}
