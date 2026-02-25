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

// SuperdocV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperdocV1Service] method instead.
type SuperdocV1Service struct {
	Options []option.RequestOption
}

// NewSuperdocV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuperdocV1Service(opts ...option.RequestOption) (r *SuperdocV1Service) {
	r = &SuperdocV1Service{}
	r.Options = opts
	return
}

// Populate fields inside a DOCX template using SuperDoc annotations. Supports
// text, images, dates, and numbers. Can target individual fields by ID or multiple
// fields by group.
func (r *SuperdocV1Service) Annotate(ctx context.Context, body SuperdocV1AnnotateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	path := "superdoc/v1/annotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Convert documents between formats using SuperDoc. Supports DOCX to PDF, Markdown
// to DOCX, and HTML to DOCX conversions.
func (r *SuperdocV1Service) Convert(ctx context.Context, body SuperdocV1ConvertParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	path := "superdoc/v1/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SuperdocV1AnnotateParams struct {
	// Document source - provide either URL or base64
	Document param.Field[SuperdocV1AnnotateParamsDocument] `json:"document" api:"required"`
	// Fields to populate in the template
	Fields param.Field[[]SuperdocV1AnnotateParamsField] `json:"fields" api:"required"`
	// Output format for the annotated document
	OutputFormat param.Field[SuperdocV1AnnotateParamsOutputFormat] `json:"output_format"`
}

func (r SuperdocV1AnnotateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Document source - provide either URL or base64
type SuperdocV1AnnotateParamsDocument struct {
	// Base64-encoded DOCX template
	Base64 param.Field[string] `json:"base64"`
	// URL to the DOCX template
	URL param.Field[string] `json:"url"`
}

func (r SuperdocV1AnnotateParamsDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperdocV1AnnotateParamsField struct {
	// Field data type
	Type param.Field[SuperdocV1AnnotateParamsFieldsType] `json:"type" api:"required"`
	// Value to populate
	Value param.Field[SuperdocV1AnnotateParamsFieldsValueUnion] `json:"value" api:"required"`
	// Target field ID (for single field)
	ID param.Field[string] `json:"id"`
	// Target field group (for multiple fields with same tag)
	Group param.Field[string] `json:"group"`
	// Optional configuration (e.g., image dimensions)
	Options param.Field[SuperdocV1AnnotateParamsFieldsOptions] `json:"options"`
}

func (r SuperdocV1AnnotateParamsField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field data type
type SuperdocV1AnnotateParamsFieldsType string

const (
	SuperdocV1AnnotateParamsFieldsTypeText   SuperdocV1AnnotateParamsFieldsType = "text"
	SuperdocV1AnnotateParamsFieldsTypeImage  SuperdocV1AnnotateParamsFieldsType = "image"
	SuperdocV1AnnotateParamsFieldsTypeDate   SuperdocV1AnnotateParamsFieldsType = "date"
	SuperdocV1AnnotateParamsFieldsTypeNumber SuperdocV1AnnotateParamsFieldsType = "number"
)

func (r SuperdocV1AnnotateParamsFieldsType) IsKnown() bool {
	switch r {
	case SuperdocV1AnnotateParamsFieldsTypeText, SuperdocV1AnnotateParamsFieldsTypeImage, SuperdocV1AnnotateParamsFieldsTypeDate, SuperdocV1AnnotateParamsFieldsTypeNumber:
		return true
	}
	return false
}

// Value to populate
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type SuperdocV1AnnotateParamsFieldsValueUnion interface {
	ImplementsSuperdocV1AnnotateParamsFieldsValueUnion()
}

// Optional configuration (e.g., image dimensions)
type SuperdocV1AnnotateParamsFieldsOptions struct {
	// Image height in pixels
	Height param.Field[float64] `json:"height"`
	// Image width in pixels
	Width param.Field[float64] `json:"width"`
}

func (r SuperdocV1AnnotateParamsFieldsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output format for the annotated document
type SuperdocV1AnnotateParamsOutputFormat string

const (
	SuperdocV1AnnotateParamsOutputFormatDocx SuperdocV1AnnotateParamsOutputFormat = "docx"
	SuperdocV1AnnotateParamsOutputFormatPdf  SuperdocV1AnnotateParamsOutputFormat = "pdf"
)

func (r SuperdocV1AnnotateParamsOutputFormat) IsKnown() bool {
	switch r {
	case SuperdocV1AnnotateParamsOutputFormatDocx, SuperdocV1AnnotateParamsOutputFormatPdf:
		return true
	}
	return false
}

type SuperdocV1ConvertParams struct {
	// Source format of the document
	From param.Field[SuperdocV1ConvertParamsFrom] `json:"from" api:"required"`
	// Base64-encoded document content
	DocumentBase64 param.Field[string] `json:"document_base64"`
	// URL to the document to convert
	DocumentURL param.Field[string] `json:"document_url"`
	// Target format for conversion
	To param.Field[SuperdocV1ConvertParamsTo] `json:"to"`
}

func (r SuperdocV1ConvertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Source format of the document
type SuperdocV1ConvertParamsFrom string

const (
	SuperdocV1ConvertParamsFromDocx SuperdocV1ConvertParamsFrom = "docx"
	SuperdocV1ConvertParamsFromMd   SuperdocV1ConvertParamsFrom = "md"
	SuperdocV1ConvertParamsFromHTML SuperdocV1ConvertParamsFrom = "html"
)

func (r SuperdocV1ConvertParamsFrom) IsKnown() bool {
	switch r {
	case SuperdocV1ConvertParamsFromDocx, SuperdocV1ConvertParamsFromMd, SuperdocV1ConvertParamsFromHTML:
		return true
	}
	return false
}

// Target format for conversion
type SuperdocV1ConvertParamsTo string

const (
	SuperdocV1ConvertParamsToPdf  SuperdocV1ConvertParamsTo = "pdf"
	SuperdocV1ConvertParamsToDocx SuperdocV1ConvertParamsTo = "docx"
)

func (r SuperdocV1ConvertParamsTo) IsKnown() bool {
	switch r {
	case SuperdocV1ConvertParamsToPdf, SuperdocV1ConvertParamsToDocx:
		return true
	}
	return false
}
