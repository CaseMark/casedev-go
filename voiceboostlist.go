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

// Audio transcription and text-to-speech
//
// VoiceBoostListService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceBoostListService] method instead.
type VoiceBoostListService struct {
	Options []option.RequestOption
}

// NewVoiceBoostListService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceBoostListService(opts ...option.RequestOption) (r *VoiceBoostListService) {
	r = &VoiceBoostListService{}
	r.Options = opts
	return
}

// Extracts a categorized word boost list from vault documents or raw text using
// LLM entity extraction. The resulting list can be passed as `word_boost` to the
// transcription endpoint for improved accuracy.
func (r *VoiceBoostListService) Extract(ctx context.Context, body VoiceBoostListExtractParams, opts ...option.RequestOption) (res *VoiceBoostListExtractResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice/boost-list/extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generates a categorized word boost list from a completed transcription job.
// Extracts entities from the pass-1 transcript for use as `word_boost` in a second
// transcription pass.
func (r *VoiceBoostListService) Generate(ctx context.Context, body VoiceBoostListGenerateParams, opts ...option.RequestOption) (res *VoiceBoostListGenerateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "voice/boost-list/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type VoiceBoostListExtractResponse struct {
	Items     []VoiceBoostListExtractResponseItem `json:"items"`
	Source    VoiceBoostListExtractResponseSource `json:"source"`
	SourceIDs []string                            `json:"source_ids"`
	JSON      voiceBoostListExtractResponseJSON   `json:"-"`
}

// voiceBoostListExtractResponseJSON contains the JSON metadata for the struct
// [VoiceBoostListExtractResponse]
type voiceBoostListExtractResponseJSON struct {
	Items       apijson.Field
	Source      apijson.Field
	SourceIDs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceBoostListExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceBoostListExtractResponseJSON) RawJSON() string {
	return r.raw
}

type VoiceBoostListExtractResponseItem struct {
	BoostParam VoiceBoostListExtractResponseItemsBoostParam `json:"boost_param"`
	Category   string                                       `json:"category"`
	Word       string                                       `json:"word"`
	JSON       voiceBoostListExtractResponseItemJSON        `json:"-"`
}

// voiceBoostListExtractResponseItemJSON contains the JSON metadata for the struct
// [VoiceBoostListExtractResponseItem]
type voiceBoostListExtractResponseItemJSON struct {
	BoostParam  apijson.Field
	Category    apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceBoostListExtractResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceBoostListExtractResponseItemJSON) RawJSON() string {
	return r.raw
}

type VoiceBoostListExtractResponseItemsBoostParam string

const (
	VoiceBoostListExtractResponseItemsBoostParamLow     VoiceBoostListExtractResponseItemsBoostParam = "low"
	VoiceBoostListExtractResponseItemsBoostParamDefault VoiceBoostListExtractResponseItemsBoostParam = "default"
	VoiceBoostListExtractResponseItemsBoostParamHigh    VoiceBoostListExtractResponseItemsBoostParam = "high"
)

func (r VoiceBoostListExtractResponseItemsBoostParam) IsKnown() bool {
	switch r {
	case VoiceBoostListExtractResponseItemsBoostParamLow, VoiceBoostListExtractResponseItemsBoostParamDefault, VoiceBoostListExtractResponseItemsBoostParamHigh:
		return true
	}
	return false
}

type VoiceBoostListExtractResponseSource string

const (
	VoiceBoostListExtractResponseSourceDocument VoiceBoostListExtractResponseSource = "document"
	VoiceBoostListExtractResponseSourceText     VoiceBoostListExtractResponseSource = "text"
)

func (r VoiceBoostListExtractResponseSource) IsKnown() bool {
	switch r {
	case VoiceBoostListExtractResponseSourceDocument, VoiceBoostListExtractResponseSourceText:
		return true
	}
	return false
}

type VoiceBoostListGenerateResponse struct {
	Items     []VoiceBoostListGenerateResponseItem `json:"items"`
	Source    VoiceBoostListGenerateResponseSource `json:"source"`
	SourceIDs []string                             `json:"source_ids"`
	JSON      voiceBoostListGenerateResponseJSON   `json:"-"`
}

// voiceBoostListGenerateResponseJSON contains the JSON metadata for the struct
// [VoiceBoostListGenerateResponse]
type voiceBoostListGenerateResponseJSON struct {
	Items       apijson.Field
	Source      apijson.Field
	SourceIDs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceBoostListGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceBoostListGenerateResponseJSON) RawJSON() string {
	return r.raw
}

type VoiceBoostListGenerateResponseItem struct {
	BoostParam VoiceBoostListGenerateResponseItemsBoostParam `json:"boost_param"`
	Category   string                                        `json:"category"`
	Word       string                                        `json:"word"`
	JSON       voiceBoostListGenerateResponseItemJSON        `json:"-"`
}

// voiceBoostListGenerateResponseItemJSON contains the JSON metadata for the struct
// [VoiceBoostListGenerateResponseItem]
type voiceBoostListGenerateResponseItemJSON struct {
	BoostParam  apijson.Field
	Category    apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VoiceBoostListGenerateResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r voiceBoostListGenerateResponseItemJSON) RawJSON() string {
	return r.raw
}

type VoiceBoostListGenerateResponseItemsBoostParam string

const (
	VoiceBoostListGenerateResponseItemsBoostParamLow     VoiceBoostListGenerateResponseItemsBoostParam = "low"
	VoiceBoostListGenerateResponseItemsBoostParamDefault VoiceBoostListGenerateResponseItemsBoostParam = "default"
	VoiceBoostListGenerateResponseItemsBoostParamHigh    VoiceBoostListGenerateResponseItemsBoostParam = "high"
)

func (r VoiceBoostListGenerateResponseItemsBoostParam) IsKnown() bool {
	switch r {
	case VoiceBoostListGenerateResponseItemsBoostParamLow, VoiceBoostListGenerateResponseItemsBoostParamDefault, VoiceBoostListGenerateResponseItemsBoostParamHigh:
		return true
	}
	return false
}

type VoiceBoostListGenerateResponseSource string

const (
	VoiceBoostListGenerateResponseSourceTranscript VoiceBoostListGenerateResponseSource = "transcript"
)

func (r VoiceBoostListGenerateResponseSource) IsKnown() bool {
	switch r {
	case VoiceBoostListGenerateResponseSourceTranscript:
		return true
	}
	return false
}

type VoiceBoostListExtractParams struct {
	// Optional filter for entity categories to extract
	Categories param.Field[[]VoiceBoostListExtractParamsCategory] `json:"categories"`
	// Object IDs of documents to extract entities from (PDFs, text files)
	ObjectIDs param.Field[[]string] `json:"object_ids"`
	// Raw text input for entity extraction (alternative to vault documents)
	Text param.Field[string] `json:"text"`
	// Vault ID containing the source documents (use with object_ids)
	VaultID param.Field[string] `json:"vault_id"`
}

func (r VoiceBoostListExtractParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VoiceBoostListExtractParamsCategory string

const (
	VoiceBoostListExtractParamsCategoryPerson       VoiceBoostListExtractParamsCategory = "person"
	VoiceBoostListExtractParamsCategoryOrganization VoiceBoostListExtractParamsCategory = "organization"
	VoiceBoostListExtractParamsCategoryLegalTerm    VoiceBoostListExtractParamsCategory = "legal_term"
	VoiceBoostListExtractParamsCategoryMedical      VoiceBoostListExtractParamsCategory = "medical"
	VoiceBoostListExtractParamsCategoryCitation     VoiceBoostListExtractParamsCategory = "citation"
	VoiceBoostListExtractParamsCategoryEmail        VoiceBoostListExtractParamsCategory = "email"
)

func (r VoiceBoostListExtractParamsCategory) IsKnown() bool {
	switch r {
	case VoiceBoostListExtractParamsCategoryPerson, VoiceBoostListExtractParamsCategoryOrganization, VoiceBoostListExtractParamsCategoryLegalTerm, VoiceBoostListExtractParamsCategoryMedical, VoiceBoostListExtractParamsCategoryCitation, VoiceBoostListExtractParamsCategoryEmail:
		return true
	}
	return false
}

type VoiceBoostListGenerateParams struct {
	// Completed pass-1 transcription job ID (tr\_...)
	TranscriptionJobID param.Field[string] `json:"transcription_job_id" api:"required"`
	// Optional filter for entity categories to extract
	Categories param.Field[[]VoiceBoostListGenerateParamsCategory] `json:"categories"`
}

func (r VoiceBoostListGenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VoiceBoostListGenerateParamsCategory string

const (
	VoiceBoostListGenerateParamsCategoryPerson       VoiceBoostListGenerateParamsCategory = "person"
	VoiceBoostListGenerateParamsCategoryOrganization VoiceBoostListGenerateParamsCategory = "organization"
	VoiceBoostListGenerateParamsCategoryLegalTerm    VoiceBoostListGenerateParamsCategory = "legal_term"
	VoiceBoostListGenerateParamsCategoryMedical      VoiceBoostListGenerateParamsCategory = "medical"
	VoiceBoostListGenerateParamsCategoryCitation     VoiceBoostListGenerateParamsCategory = "citation"
	VoiceBoostListGenerateParamsCategoryEmail        VoiceBoostListGenerateParamsCategory = "email"
)

func (r VoiceBoostListGenerateParamsCategory) IsKnown() bool {
	switch r {
	case VoiceBoostListGenerateParamsCategoryPerson, VoiceBoostListGenerateParamsCategoryOrganization, VoiceBoostListGenerateParamsCategoryLegalTerm, VoiceBoostListGenerateParamsCategoryMedical, VoiceBoostListGenerateParamsCategoryCitation, VoiceBoostListGenerateParamsCategoryEmail:
		return true
	}
	return false
}
