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

// Extract text from PDFs, images, and scanned documents
//
// OcrV1Service contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOcrV1Service] method instead.
type OcrV1Service struct {
	Options []option.RequestOption
}

// NewOcrV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOcrV1Service(opts ...option.RequestOption) (r *OcrV1Service) {
	r = &OcrV1Service{}
	r.Options = opts
	return
}

// Retrieve the status and results of an OCR job. Returns job progress and
// metadata; full extracted text is included only when include_text=true.
func (r *OcrV1Service) Get(ctx context.Context, id string, query OcrV1GetParams, opts ...option.RequestOption) (res *OcrV1GetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ocr/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Download OCR processing results in various formats by redirecting to the OCR
// service download URL.
func (r *OcrV1Service) Download(ctx context.Context, id string, type_ OcrV1DownloadParamsType, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ocr/v1/%s/download/%v", id, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submit a document for OCR processing to extract text, detect tables, forms, and
// other features. Supports PDFs, images, and scanned documents. Returns a job ID
// that can be used to track processing status.
func (r *OcrV1Service) Process(ctx context.Context, body OcrV1ProcessParams, opts ...option.RequestOption) (res *OcrV1ProcessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ocr/v1/process"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type OcrV1GetResponse struct {
	// OCR job ID
	ID string `json:"id" api:"required"`
	// Job creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Current job status
	Status OcrV1GetResponseStatus `json:"status" api:"required"`
	// Total number of OCR chunks.
	ChunkCount int64 `json:"chunk_count"`
	// Exact number of failed OCR chunks.
	ChunksFailed int64 `json:"chunks_failed"`
	// Job completion timestamp
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// Safe human-readable failure message; null when no failure is recorded.
	Error string `json:"error" api:"nullable"`
	// Stable failure identifier, such as source_access_denied, invalid_pdf,
	// inference_failed, chunks_failed, or result_upload_failed. Historic failures may
	// be unclassified.
	ErrorCode string `json:"error_code" api:"nullable"`
	// Operation that failed.
	ErrorStage OcrV1GetResponseErrorStage `json:"error_stage" api:"nullable"`
	// Bounded failure samples. Page ranges are one-based and inclusive; chunks_failed
	// and chunk_count retain exact totals.
	FailedChunks []OcrV1GetResponseFailedChunk `json:"failed_chunks" api:"nullable"`
	// Additional processing metadata
	Metadata interface{} `json:"metadata"`
	// Number of pages processed
	PageCount int64 `json:"page_count"`
	// Whether retrying the same document may help. Unknown failures retain bounded
	// retry behavior.
	Retryable bool `json:"retryable" api:"nullable"`
	// Extracted text content when completed and include_text=true
	Text string               `json:"text"`
	JSON ocrV1GetResponseJSON `json:"-"`
}

// ocrV1GetResponseJSON contains the JSON metadata for the struct
// [OcrV1GetResponse]
type ocrV1GetResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Status       apijson.Field
	ChunkCount   apijson.Field
	ChunksFailed apijson.Field
	CompletedAt  apijson.Field
	Error        apijson.Field
	ErrorCode    apijson.Field
	ErrorStage   apijson.Field
	FailedChunks apijson.Field
	Metadata     apijson.Field
	PageCount    apijson.Field
	Retryable    apijson.Field
	Text         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OcrV1GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ocrV1GetResponseJSON) RawJSON() string {
	return r.raw
}

// Current job status
type OcrV1GetResponseStatus string

const (
	OcrV1GetResponseStatusPending    OcrV1GetResponseStatus = "pending"
	OcrV1GetResponseStatusProcessing OcrV1GetResponseStatus = "processing"
	OcrV1GetResponseStatusCompleted  OcrV1GetResponseStatus = "completed"
	OcrV1GetResponseStatusFailed     OcrV1GetResponseStatus = "failed"
)

func (r OcrV1GetResponseStatus) IsKnown() bool {
	switch r {
	case OcrV1GetResponseStatusPending, OcrV1GetResponseStatusProcessing, OcrV1GetResponseStatusCompleted, OcrV1GetResponseStatusFailed:
		return true
	}
	return false
}

// Operation that failed.
type OcrV1GetResponseErrorStage string

const (
	OcrV1GetResponseErrorStageDispatch           OcrV1GetResponseErrorStage = "dispatch"
	OcrV1GetResponseErrorStageDownload           OcrV1GetResponseErrorStage = "download"
	OcrV1GetResponseErrorStageDocumentValidation OcrV1GetResponseErrorStage = "document_validation"
	OcrV1GetResponseErrorStageChunking           OcrV1GetResponseErrorStage = "chunking"
	OcrV1GetResponseErrorStageEngineInit         OcrV1GetResponseErrorStage = "engine_init"
	OcrV1GetResponseErrorStageInference          OcrV1GetResponseErrorStage = "inference"
	OcrV1GetResponseErrorStageFeatureProcessing  OcrV1GetResponseErrorStage = "feature_processing"
	OcrV1GetResponseErrorStageResultWrite        OcrV1GetResponseErrorStage = "result_write"
	OcrV1GetResponseErrorStageChunkProcessing    OcrV1GetResponseErrorStage = "chunk_processing"
	OcrV1GetResponseErrorStageResultAssembly     OcrV1GetResponseErrorStage = "result_assembly"
	OcrV1GetResponseErrorStageResultUpload       OcrV1GetResponseErrorStage = "result_upload"
)

func (r OcrV1GetResponseErrorStage) IsKnown() bool {
	switch r {
	case OcrV1GetResponseErrorStageDispatch, OcrV1GetResponseErrorStageDownload, OcrV1GetResponseErrorStageDocumentValidation, OcrV1GetResponseErrorStageChunking, OcrV1GetResponseErrorStageEngineInit, OcrV1GetResponseErrorStageInference, OcrV1GetResponseErrorStageFeatureProcessing, OcrV1GetResponseErrorStageResultWrite, OcrV1GetResponseErrorStageChunkProcessing, OcrV1GetResponseErrorStageResultAssembly, OcrV1GetResponseErrorStageResultUpload:
		return true
	}
	return false
}

type OcrV1GetResponseFailedChunk struct {
	ChunkID    string                          `json:"chunk_id"`
	Error      string                          `json:"error"`
	ErrorCode  string                          `json:"error_code"`
	ErrorStage string                          `json:"error_stage"`
	PageEnd    int64                           `json:"page_end"`
	PageStart  int64                           `json:"page_start"`
	Retryable  bool                            `json:"retryable"`
	JSON       ocrV1GetResponseFailedChunkJSON `json:"-"`
}

// ocrV1GetResponseFailedChunkJSON contains the JSON metadata for the struct
// [OcrV1GetResponseFailedChunk]
type ocrV1GetResponseFailedChunkJSON struct {
	ChunkID     apijson.Field
	Error       apijson.Field
	ErrorCode   apijson.Field
	ErrorStage  apijson.Field
	PageEnd     apijson.Field
	PageStart   apijson.Field
	Retryable   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OcrV1GetResponseFailedChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ocrV1GetResponseFailedChunkJSON) RawJSON() string {
	return r.raw
}

type OcrV1ProcessResponse struct {
	// Unique job identifier
	ID string `json:"id"`
	// Job creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Document identifier
	DocumentID string `json:"document_id"`
	// OCR engine used
	Engine string `json:"engine"`
	// Estimated completion time
	EstimatedCompletion time.Time `json:"estimated_completion" format:"date-time"`
	// Number of pages detected
	PageCount int64 `json:"page_count"`
	// Current job status
	Status OcrV1ProcessResponseStatus `json:"status"`
	JSON   ocrV1ProcessResponseJSON   `json:"-"`
}

// ocrV1ProcessResponseJSON contains the JSON metadata for the struct
// [OcrV1ProcessResponse]
type ocrV1ProcessResponseJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	DocumentID          apijson.Field
	Engine              apijson.Field
	EstimatedCompletion apijson.Field
	PageCount           apijson.Field
	Status              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OcrV1ProcessResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ocrV1ProcessResponseJSON) RawJSON() string {
	return r.raw
}

// Current job status
type OcrV1ProcessResponseStatus string

const (
	OcrV1ProcessResponseStatusQueued     OcrV1ProcessResponseStatus = "queued"
	OcrV1ProcessResponseStatusProcessing OcrV1ProcessResponseStatus = "processing"
	OcrV1ProcessResponseStatusCompleted  OcrV1ProcessResponseStatus = "completed"
	OcrV1ProcessResponseStatusFailed     OcrV1ProcessResponseStatus = "failed"
)

func (r OcrV1ProcessResponseStatus) IsKnown() bool {
	switch r {
	case OcrV1ProcessResponseStatusQueued, OcrV1ProcessResponseStatusProcessing, OcrV1ProcessResponseStatusCompleted, OcrV1ProcessResponseStatusFailed:
		return true
	}
	return false
}

type OcrV1GetParams struct {
	// Include full OCR text in completed responses (default: false)
	IncludeText param.Field[OcrV1GetParamsIncludeText] `query:"include_text"`
}

// URLQuery serializes [OcrV1GetParams]'s query parameters as `url.Values`.
func (r OcrV1GetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Include full OCR text in completed responses (default: false)
type OcrV1GetParamsIncludeText string

const (
	OcrV1GetParamsIncludeTextTrue  OcrV1GetParamsIncludeText = "true"
	OcrV1GetParamsIncludeTextFalse OcrV1GetParamsIncludeText = "false"
)

func (r OcrV1GetParamsIncludeText) IsKnown() bool {
	switch r {
	case OcrV1GetParamsIncludeTextTrue, OcrV1GetParamsIncludeTextFalse:
		return true
	}
	return false
}

type OcrV1DownloadParamsType string

const (
	OcrV1DownloadParamsTypeText     OcrV1DownloadParamsType = "text"
	OcrV1DownloadParamsTypeJson     OcrV1DownloadParamsType = "json"
	OcrV1DownloadParamsTypePdf      OcrV1DownloadParamsType = "pdf"
	OcrV1DownloadParamsTypeOriginal OcrV1DownloadParamsType = "original"
)

func (r OcrV1DownloadParamsType) IsKnown() bool {
	switch r {
	case OcrV1DownloadParamsTypeText, OcrV1DownloadParamsTypeJson, OcrV1DownloadParamsTypePdf, OcrV1DownloadParamsTypeOriginal:
		return true
	}
	return false
}

type OcrV1ProcessParams struct {
	// URL or S3 path to the document to process
	DocumentURL param.Field[string] `json:"document_url" api:"required"`
	// URL to receive completion webhook
	CallbackURL param.Field[string] `json:"callback_url"`
	// Optional custom document identifier
	DocumentID param.Field[string] `json:"document_id"`
	// OCR engine to use
	Engine param.Field[OcrV1ProcessParamsEngine] `json:"engine"`
	// Additional processing options
	Features param.Field[OcrV1ProcessParamsFeatures] `json:"features"`
	// S3 bucket to store results
	ResultBucket param.Field[string] `json:"result_bucket"`
	// S3 key prefix for results
	ResultPrefix param.Field[string] `json:"result_prefix"`
}

func (r OcrV1ProcessParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OCR engine to use
type OcrV1ProcessParamsEngine string

const (
	OcrV1ProcessParamsEngineDoctr     OcrV1ProcessParamsEngine = "doctr"
	OcrV1ProcessParamsEnginePaddleocr OcrV1ProcessParamsEngine = "paddleocr"
)

func (r OcrV1ProcessParamsEngine) IsKnown() bool {
	switch r {
	case OcrV1ProcessParamsEngineDoctr, OcrV1ProcessParamsEnginePaddleocr:
		return true
	}
	return false
}

// Additional processing options
type OcrV1ProcessParamsFeatures struct {
	// Generate searchable PDF with text layer
	Embed param.Field[map[string]interface{}] `json:"embed"`
	// Detect and extract form fields
	Forms param.Field[map[string]interface{}] `json:"forms"`
	// Extract tables as structured data
	Tables param.Field[OcrV1ProcessParamsFeaturesTables] `json:"tables"`
}

func (r OcrV1ProcessParamsFeatures) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extract tables as structured data
type OcrV1ProcessParamsFeaturesTables struct {
	// Output format for extracted tables
	Format      param.Field[OcrV1ProcessParamsFeaturesTablesFormat] `json:"format"`
	ExtraFields map[string]interface{}                              `json:"-,extras"`
}

func (r OcrV1ProcessParamsFeaturesTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output format for extracted tables
type OcrV1ProcessParamsFeaturesTablesFormat string

const (
	OcrV1ProcessParamsFeaturesTablesFormatCsv  OcrV1ProcessParamsFeaturesTablesFormat = "csv"
	OcrV1ProcessParamsFeaturesTablesFormatJson OcrV1ProcessParamsFeaturesTablesFormat = "json"
)

func (r OcrV1ProcessParamsFeaturesTablesFormat) IsKnown() bool {
	switch r {
	case OcrV1ProcessParamsFeaturesTablesFormatCsv, OcrV1ProcessParamsFeaturesTablesFormatJson:
		return true
	}
	return false
}
