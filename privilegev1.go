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

// PrivilegeV1Service contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrivilegeV1Service] method instead.
type PrivilegeV1Service struct {
	Options []option.RequestOption
}

// NewPrivilegeV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrivilegeV1Service(opts ...option.RequestOption) (r *PrivilegeV1Service) {
	r = &PrivilegeV1Service{}
	r.Options = opts
	return
}

// Analyzes text or vault documents for legal privilege. Detects attorney-client
// privilege, work product doctrine, common interest privilege, and litigation hold
// materials.
//
// Returns structured privilege flags with confidence scores and policy-friendly
// rationale suitable for discovery workflows and privilege logs.
//
// **Size Limit:** Maximum 200,000 characters (larger documents rejected).
//
// **Permissions:** Requires `chat` permission. When using `document_id`, also
// requires `vault` permission.
//
// **Note:** When analyzing vault documents, results are automatically stored in
// the document's `privilege_analysis` metadata field.
func (r *PrivilegeV1Service) Detect(ctx context.Context, body PrivilegeV1DetectParams, opts ...option.RequestOption) (res *PrivilegeV1DetectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "privilege/v1/detect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PrivilegeV1DetectResponse struct {
	Categories []PrivilegeV1DetectResponseCategory `json:"categories" api:"required"`
	// Overall confidence score (0-1)
	Confidence float64 `json:"confidence" api:"required"`
	// Policy-friendly explanation for privilege log
	PolicyRationale string `json:"policy_rationale" api:"required"`
	// Whether any privilege was detected
	Privileged bool `json:"privileged" api:"required"`
	// Recommended action for discovery
	Recommendation PrivilegeV1DetectResponseRecommendation `json:"recommendation" api:"required"`
	JSON           privilegeV1DetectResponseJSON           `json:"-"`
}

// privilegeV1DetectResponseJSON contains the JSON metadata for the struct
// [PrivilegeV1DetectResponse]
type privilegeV1DetectResponseJSON struct {
	Categories      apijson.Field
	Confidence      apijson.Field
	PolicyRationale apijson.Field
	Privileged      apijson.Field
	Recommendation  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PrivilegeV1DetectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r privilegeV1DetectResponseJSON) RawJSON() string {
	return r.raw
}

type PrivilegeV1DetectResponseCategory struct {
	// Confidence for this category (0-1)
	Confidence float64 `json:"confidence"`
	// Whether this privilege type was detected
	Detected bool `json:"detected"`
	// Specific phrases or patterns found
	Indicators []string `json:"indicators"`
	// Explanation of detection result
	Rationale string `json:"rationale"`
	// Privilege category
	Type string                                `json:"type"`
	JSON privilegeV1DetectResponseCategoryJSON `json:"-"`
}

// privilegeV1DetectResponseCategoryJSON contains the JSON metadata for the struct
// [PrivilegeV1DetectResponseCategory]
type privilegeV1DetectResponseCategoryJSON struct {
	Confidence  apijson.Field
	Detected    apijson.Field
	Indicators  apijson.Field
	Rationale   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrivilegeV1DetectResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r privilegeV1DetectResponseCategoryJSON) RawJSON() string {
	return r.raw
}

// Recommended action for discovery
type PrivilegeV1DetectResponseRecommendation string

const (
	PrivilegeV1DetectResponseRecommendationWithhold PrivilegeV1DetectResponseRecommendation = "withhold"
	PrivilegeV1DetectResponseRecommendationRedact   PrivilegeV1DetectResponseRecommendation = "redact"
	PrivilegeV1DetectResponseRecommendationProduce  PrivilegeV1DetectResponseRecommendation = "produce"
	PrivilegeV1DetectResponseRecommendationReview   PrivilegeV1DetectResponseRecommendation = "review"
)

func (r PrivilegeV1DetectResponseRecommendation) IsKnown() bool {
	switch r {
	case PrivilegeV1DetectResponseRecommendationWithhold, PrivilegeV1DetectResponseRecommendationRedact, PrivilegeV1DetectResponseRecommendationProduce, PrivilegeV1DetectResponseRecommendationReview:
		return true
	}
	return false
}

type PrivilegeV1DetectParams struct {
	// Privilege categories to check. Defaults to all: attorney_client, work_product,
	// common_interest, litigation_hold
	Categories param.Field[[]PrivilegeV1DetectParamsCategory] `json:"categories"`
	// Text content to analyze (required if document_id not provided)
	Content param.Field[string] `json:"content"`
	// Vault object ID to analyze (required if content not provided)
	DocumentID param.Field[string] `json:"document_id"`
	// Include detailed rationale for each category
	IncludeRationale param.Field[bool] `json:"include_rationale"`
	// Jurisdiction for privilege rules
	Jurisdiction param.Field[PrivilegeV1DetectParamsJurisdiction] `json:"jurisdiction"`
	// LLM model to use for analysis
	Model param.Field[string] `json:"model"`
	// Vault ID (required when using document_id)
	VaultID param.Field[string] `json:"vault_id"`
}

func (r PrivilegeV1DetectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrivilegeV1DetectParamsCategory string

const (
	PrivilegeV1DetectParamsCategoryAttorneyClient PrivilegeV1DetectParamsCategory = "attorney_client"
	PrivilegeV1DetectParamsCategoryWorkProduct    PrivilegeV1DetectParamsCategory = "work_product"
	PrivilegeV1DetectParamsCategoryCommonInterest PrivilegeV1DetectParamsCategory = "common_interest"
	PrivilegeV1DetectParamsCategoryLitigationHold PrivilegeV1DetectParamsCategory = "litigation_hold"
)

func (r PrivilegeV1DetectParamsCategory) IsKnown() bool {
	switch r {
	case PrivilegeV1DetectParamsCategoryAttorneyClient, PrivilegeV1DetectParamsCategoryWorkProduct, PrivilegeV1DetectParamsCategoryCommonInterest, PrivilegeV1DetectParamsCategoryLitigationHold:
		return true
	}
	return false
}

// Jurisdiction for privilege rules
type PrivilegeV1DetectParamsJurisdiction string

const (
	PrivilegeV1DetectParamsJurisdictionUsFederal PrivilegeV1DetectParamsJurisdiction = "US-Federal"
)

func (r PrivilegeV1DetectParamsJurisdiction) IsKnown() bool {
	switch r {
	case PrivilegeV1DetectParamsJurisdictionUsFederal:
		return true
	}
	return false
}
