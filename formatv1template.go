// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// FormatV1TemplateService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFormatV1TemplateService] method instead.
type FormatV1TemplateService struct {
	Options []option.RequestOption
}

// NewFormatV1TemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFormatV1TemplateService(opts ...option.RequestOption) (r *FormatV1TemplateService) {
	r = &FormatV1TemplateService{}
	r.Options = opts
	return
}

// Create a new format template for document formatting. Templates support
// variables using `{{variable}}` syntax and can be used for captions, signatures,
// letterheads, certificates, footers, or custom formatting needs.
func (r *FormatV1TemplateService) New(ctx context.Context, body FormatV1TemplateNewParams, opts ...option.RequestOption) (res *FormatV1TemplateNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "format/v1/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific document format template by ID. Format templates define how
// documents should be structured and formatted for specific legal use cases such
// as contracts, briefs, or pleadings.
func (r *FormatV1TemplateService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FormatV1TemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("format/v1/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all format templates for the organization. Templates define reusable
// document formatting patterns with customizable variables for consistent legal
// document generation.
//
// Filter by type to get specific template categories like contracts, pleadings, or
// correspondence.
func (r *FormatV1TemplateService) List(ctx context.Context, query FormatV1TemplateListParams, opts ...option.RequestOption) (res *FormatV1TemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "format/v1/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FormatV1TemplateNewResponse struct {
	// Template ID
	ID string `json:"id"`
	// Creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Template name
	Name string `json:"name"`
	// Template type
	Type string `json:"type"`
	// Detected template variables
	Variables []string                        `json:"variables"`
	JSON      formatV1TemplateNewResponseJSON `json:"-"`
}

// formatV1TemplateNewResponseJSON contains the JSON metadata for the struct
// [FormatV1TemplateNewResponse]
type formatV1TemplateNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Variables   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FormatV1TemplateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r formatV1TemplateNewResponseJSON) RawJSON() string {
	return r.raw
}

type FormatV1TemplateGetResponse struct {
	// Unique template identifier
	ID string `json:"id"`
	// Template formatting rules and structure
	Content interface{} `json:"content"`
	// Template creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Template description
	Description string `json:"description"`
	// Template name
	Name string `json:"name"`
	// Organization ID that owns the template
	OrganizationID string `json:"organizationId"`
	// Template last modification timestamp
	UpdatedAt time.Time                       `json:"updatedAt" format:"date-time"`
	JSON      formatV1TemplateGetResponseJSON `json:"-"`
}

// formatV1TemplateGetResponseJSON contains the JSON metadata for the struct
// [FormatV1TemplateGetResponse]
type formatV1TemplateGetResponseJSON struct {
	ID             apijson.Field
	Content        apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FormatV1TemplateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r formatV1TemplateGetResponseJSON) RawJSON() string {
	return r.raw
}

type FormatV1TemplateListResponse struct {
	Templates []FormatV1TemplateListResponseTemplate `json:"templates"`
	JSON      formatV1TemplateListResponseJSON       `json:"-"`
}

// formatV1TemplateListResponseJSON contains the JSON metadata for the struct
// [FormatV1TemplateListResponse]
type formatV1TemplateListResponseJSON struct {
	Templates   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FormatV1TemplateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r formatV1TemplateListResponseJSON) RawJSON() string {
	return r.raw
}

type FormatV1TemplateListResponseTemplate struct {
	// Unique template identifier
	ID string `json:"id"`
	// Template creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Template description
	Description string `json:"description"`
	// Template name
	Name string `json:"name"`
	// Template tags for organization
	Tags []interface{} `json:"tags"`
	// Template type/category
	Type string `json:"type"`
	// Number of times template has been used
	UsageCount int64 `json:"usageCount"`
	// Template variables for customization
	Variables []interface{}                            `json:"variables"`
	JSON      formatV1TemplateListResponseTemplateJSON `json:"-"`
}

// formatV1TemplateListResponseTemplateJSON contains the JSON metadata for the
// struct [FormatV1TemplateListResponseTemplate]
type formatV1TemplateListResponseTemplateJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	UsageCount  apijson.Field
	Variables   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FormatV1TemplateListResponseTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r formatV1TemplateListResponseTemplateJSON) RawJSON() string {
	return r.raw
}

type FormatV1TemplateNewParams struct {
	// Template content with {{variable}} placeholders
	Content param.Field[string] `json:"content" api:"required"`
	// Template name
	Name param.Field[string] `json:"name" api:"required"`
	// Template type
	Type param.Field[FormatV1TemplateNewParamsType] `json:"type" api:"required"`
	// Template description
	Description param.Field[string] `json:"description"`
	// CSS styles for the template
	Styles param.Field[interface{}] `json:"styles"`
	// Template tags for organization
	Tags param.Field[[]string] `json:"tags"`
	// Template variables (auto-detected if not provided)
	Variables param.Field[[]string] `json:"variables"`
}

func (r FormatV1TemplateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Template type
type FormatV1TemplateNewParamsType string

const (
	FormatV1TemplateNewParamsTypeCaption     FormatV1TemplateNewParamsType = "caption"
	FormatV1TemplateNewParamsTypeSignature   FormatV1TemplateNewParamsType = "signature"
	FormatV1TemplateNewParamsTypeLetterhead  FormatV1TemplateNewParamsType = "letterhead"
	FormatV1TemplateNewParamsTypeCertificate FormatV1TemplateNewParamsType = "certificate"
	FormatV1TemplateNewParamsTypeFooter      FormatV1TemplateNewParamsType = "footer"
	FormatV1TemplateNewParamsTypeCustom      FormatV1TemplateNewParamsType = "custom"
)

func (r FormatV1TemplateNewParamsType) IsKnown() bool {
	switch r {
	case FormatV1TemplateNewParamsTypeCaption, FormatV1TemplateNewParamsTypeSignature, FormatV1TemplateNewParamsTypeLetterhead, FormatV1TemplateNewParamsTypeCertificate, FormatV1TemplateNewParamsTypeFooter, FormatV1TemplateNewParamsTypeCustom:
		return true
	}
	return false
}

type FormatV1TemplateListParams struct {
	// Filter templates by type (e.g., contract, pleading, letter)
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [FormatV1TemplateListParams]'s query parameters as
// `url.Values`.
func (r FormatV1TemplateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
