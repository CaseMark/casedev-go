// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// VoiceV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceV1Service] method instead.
type VoiceV1Service struct {
	Options []option.RequestOption
	Speak   *VoiceV1SpeakService
}

// NewVoiceV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVoiceV1Service(opts ...option.RequestOption) (r *VoiceV1Service) {
	r = &VoiceV1Service{}
	r.Options = opts
	r.Speak = NewVoiceV1SpeakService(opts...)
	return
}

// Retrieve a list of available voices for text-to-speech synthesis. This endpoint
// provides access to a comprehensive catalog of voices with various
// characteristics, languages, and styles suitable for legal document narration,
// client presentations, and accessibility purposes.
func (r *VoiceV1Service) ListVoices(ctx context.Context, query VoiceV1ListVoicesParams, opts ...option.RequestOption) (res *VoiceV1ListVoicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice/v1/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VoiceV1ListVoicesResponse struct {
	// Token for next page of results
	NextPageToken string `json:"next_page_token"`
	// Total number of voices (if requested)
	TotalCount int64                            `json:"total_count"`
	Voices     []VoiceV1ListVoicesResponseVoice `json:"voices"`
	JSON       voiceV1ListVoicesResponseJSON    `json:"-"`
}

// voiceV1ListVoicesResponseJSON contains the JSON metadata for the struct
// [VoiceV1ListVoicesResponse]
type voiceV1ListVoicesResponseJSON struct {
	NextPageToken apijson.Field
	TotalCount    apijson.Field
	Voices        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VoiceV1ListVoicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceV1ListVoicesResponseJSON) RawJSON() string {
	return r.raw
}

type VoiceV1ListVoicesResponseVoice struct {
	// Available subscription tiers
	AvailableForTiers []string `json:"available_for_tiers"`
	// Voice category
	Category string `json:"category"`
	// Voice description
	Description string `json:"description"`
	// Voice characteristics and metadata
	Labels interface{} `json:"labels"`
	// Voice name
	Name string `json:"name"`
	// URL to preview audio sample
	PreviewURL string `json:"preview_url"`
	// Unique voice identifier
	VoiceID string                             `json:"voice_id"`
	JSON    voiceV1ListVoicesResponseVoiceJSON `json:"-"`
}

// voiceV1ListVoicesResponseVoiceJSON contains the JSON metadata for the struct
// [VoiceV1ListVoicesResponseVoice]
type voiceV1ListVoicesResponseVoiceJSON struct {
	AvailableForTiers apijson.Field
	Category          apijson.Field
	Description       apijson.Field
	Labels            apijson.Field
	Name              apijson.Field
	PreviewURL        apijson.Field
	VoiceID           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *VoiceV1ListVoicesResponseVoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceV1ListVoicesResponseVoiceJSON) RawJSON() string {
	return r.raw
}

type VoiceV1ListVoicesParams struct {
	// Filter by voice category
	Category param.Field[string] `query:"category"`
	// Filter by voice collection ID
	CollectionID param.Field[string] `query:"collection_id"`
	// Whether to include total count in response
	IncludeTotalCount param.Field[bool] `query:"include_total_count"`
	// Token for retrieving the next page of results
	NextPageToken param.Field[string] `query:"next_page_token"`
	// Number of voices to return per page (max 100)
	PageSize param.Field[int64] `query:"page_size"`
	// Search term to filter voices by name or description
	Search param.Field[string] `query:"search"`
	// Field to sort by
	Sort param.Field[VoiceV1ListVoicesParamsSort] `query:"sort"`
	// Sort direction
	SortDirection param.Field[VoiceV1ListVoicesParamsSortDirection] `query:"sort_direction"`
	// Filter by voice type
	VoiceType param.Field[VoiceV1ListVoicesParamsVoiceType] `query:"voice_type"`
}

// URLQuery serializes [VoiceV1ListVoicesParams]'s query parameters as
// `url.Values`.
func (r VoiceV1ListVoicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to sort by
type VoiceV1ListVoicesParamsSort string

const (
	VoiceV1ListVoicesParamsSortName      VoiceV1ListVoicesParamsSort = "name"
	VoiceV1ListVoicesParamsSortCreatedAt VoiceV1ListVoicesParamsSort = "created_at"
	VoiceV1ListVoicesParamsSortUpdatedAt VoiceV1ListVoicesParamsSort = "updated_at"
)

func (r VoiceV1ListVoicesParamsSort) IsKnown() bool {
	switch r {
	case VoiceV1ListVoicesParamsSortName, VoiceV1ListVoicesParamsSortCreatedAt, VoiceV1ListVoicesParamsSortUpdatedAt:
		return true
	}
	return false
}

// Sort direction
type VoiceV1ListVoicesParamsSortDirection string

const (
	VoiceV1ListVoicesParamsSortDirectionAsc  VoiceV1ListVoicesParamsSortDirection = "asc"
	VoiceV1ListVoicesParamsSortDirectionDesc VoiceV1ListVoicesParamsSortDirection = "desc"
)

func (r VoiceV1ListVoicesParamsSortDirection) IsKnown() bool {
	switch r {
	case VoiceV1ListVoicesParamsSortDirectionAsc, VoiceV1ListVoicesParamsSortDirectionDesc:
		return true
	}
	return false
}

// Filter by voice type
type VoiceV1ListVoicesParamsVoiceType string

const (
	VoiceV1ListVoicesParamsVoiceTypePremade      VoiceV1ListVoicesParamsVoiceType = "premade"
	VoiceV1ListVoicesParamsVoiceTypeCloned       VoiceV1ListVoicesParamsVoiceType = "cloned"
	VoiceV1ListVoicesParamsVoiceTypeProfessional VoiceV1ListVoicesParamsVoiceType = "professional"
)

func (r VoiceV1ListVoicesParamsVoiceType) IsKnown() bool {
	switch r {
	case VoiceV1ListVoicesParamsVoiceTypePremade, VoiceV1ListVoicesParamsVoiceTypeCloned, VoiceV1ListVoicesParamsVoiceTypeProfessional:
		return true
	}
	return false
}
