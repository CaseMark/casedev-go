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

// FormatV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFormatV1Service] method instead.
type FormatV1Service struct {
	Options   []option.RequestOption
	Templates *FormatV1TemplateService
}

// NewFormatV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFormatV1Service(opts ...option.RequestOption) (r *FormatV1Service) {
	r = &FormatV1Service{}
	r.Options = opts
	r.Templates = NewFormatV1TemplateService(opts...)
	return
}

// Convert Markdown, JSON, or text content to professionally formatted PDF, DOCX,
// or HTML documents. Supports template components with variable interpolation for
// creating consistent legal documents like contracts, briefs, and reports.
func (r *FormatV1Service) NewDocument(ctx context.Context, body FormatV1NewDocumentParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	path := "format/v1/document"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FormatV1NewDocumentParams struct {
	// The source content to format
	Content param.Field[string] `json:"content,required"`
	// Desired output format
	OutputFormat param.Field[FormatV1NewDocumentParamsOutputFormat] `json:"output_format,required"`
	// Format of the input content
	InputFormat param.Field[FormatV1NewDocumentParamsInputFormat] `json:"input_format"`
	Options     param.Field[FormatV1NewDocumentParamsOptions]     `json:"options"`
}

func (r FormatV1NewDocumentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Desired output format
type FormatV1NewDocumentParamsOutputFormat string

const (
	FormatV1NewDocumentParamsOutputFormatPdf         FormatV1NewDocumentParamsOutputFormat = "pdf"
	FormatV1NewDocumentParamsOutputFormatDocx        FormatV1NewDocumentParamsOutputFormat = "docx"
	FormatV1NewDocumentParamsOutputFormatHTMLPreview FormatV1NewDocumentParamsOutputFormat = "html_preview"
)

func (r FormatV1NewDocumentParamsOutputFormat) IsKnown() bool {
	switch r {
	case FormatV1NewDocumentParamsOutputFormatPdf, FormatV1NewDocumentParamsOutputFormatDocx, FormatV1NewDocumentParamsOutputFormatHTMLPreview:
		return true
	}
	return false
}

// Format of the input content
type FormatV1NewDocumentParamsInputFormat string

const (
	FormatV1NewDocumentParamsInputFormatMd   FormatV1NewDocumentParamsInputFormat = "md"
	FormatV1NewDocumentParamsInputFormatJson FormatV1NewDocumentParamsInputFormat = "json"
	FormatV1NewDocumentParamsInputFormatText FormatV1NewDocumentParamsInputFormat = "text"
)

func (r FormatV1NewDocumentParamsInputFormat) IsKnown() bool {
	switch r {
	case FormatV1NewDocumentParamsInputFormatMd, FormatV1NewDocumentParamsInputFormatJson, FormatV1NewDocumentParamsInputFormatText:
		return true
	}
	return false
}

type FormatV1NewDocumentParamsOptions struct {
	// Template components with variables
	Components param.Field[[]FormatV1NewDocumentParamsOptionsComponent] `json:"components"`
}

func (r FormatV1NewDocumentParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FormatV1NewDocumentParamsOptionsComponent struct {
	// Inline template content
	Content param.Field[string] `json:"content"`
	// Custom styling options
	Styles param.Field[interface{}] `json:"styles"`
	// ID of saved template component
	TemplateID param.Field[string] `json:"templateId"`
	// Variables for template interpolation
	Variables param.Field[interface{}] `json:"variables"`
}

func (r FormatV1NewDocumentParamsOptionsComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
