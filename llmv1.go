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

// LlmV1Service contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmV1Service] method instead.
type LlmV1Service struct {
	Options []option.RequestOption
	Chat    *LlmV1ChatService
}

// NewLlmV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLlmV1Service(opts ...option.RequestOption) (r *LlmV1Service) {
	r = &LlmV1Service{}
	r.Options = opts
	r.Chat = NewLlmV1ChatService(opts...)
	return
}

// Create vector embeddings from text using OpenAI-compatible models. Perfect for
// semantic search, document similarity, and building RAG systems for legal
// documents.
func (r *LlmV1Service) NewEmbedding(ctx context.Context, body LlmV1NewEmbeddingParams, opts ...option.RequestOption) (res *LlmV1NewEmbeddingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm/v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a list of all available language models from 40+ providers including
// OpenAI, Anthropic, Google, and Case.dev's specialized legal models. Returns
// OpenAI-compatible model metadata with pricing information.
//
// This endpoint is compatible with OpenAI's models API format, making it easy to
// integrate with existing applications.
func (r *LlmV1Service) ListModels(ctx context.Context, opts ...option.RequestOption) (res *LlmV1ListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm/v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LlmV1NewEmbeddingResponse struct {
	Data   []LlmV1NewEmbeddingResponseData `json:"data"`
	Model  string                          `json:"model"`
	Object string                          `json:"object"`
	Usage  LlmV1NewEmbeddingResponseUsage  `json:"usage"`
	JSON   llmV1NewEmbeddingResponseJSON   `json:"-"`
}

// llmV1NewEmbeddingResponseJSON contains the JSON metadata for the struct
// [LlmV1NewEmbeddingResponse]
type llmV1NewEmbeddingResponseJSON struct {
	Data        apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmV1NewEmbeddingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1NewEmbeddingResponseJSON) RawJSON() string {
	return r.raw
}

type LlmV1NewEmbeddingResponseData struct {
	Embedding []float64                         `json:"embedding"`
	Index     int64                             `json:"index"`
	Object    string                            `json:"object"`
	JSON      llmV1NewEmbeddingResponseDataJSON `json:"-"`
}

// llmV1NewEmbeddingResponseDataJSON contains the JSON metadata for the struct
// [LlmV1NewEmbeddingResponseData]
type llmV1NewEmbeddingResponseDataJSON struct {
	Embedding   apijson.Field
	Index       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmV1NewEmbeddingResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1NewEmbeddingResponseDataJSON) RawJSON() string {
	return r.raw
}

type LlmV1NewEmbeddingResponseUsage struct {
	PromptTokens int64                              `json:"prompt_tokens"`
	TotalTokens  int64                              `json:"total_tokens"`
	JSON         llmV1NewEmbeddingResponseUsageJSON `json:"-"`
}

// llmV1NewEmbeddingResponseUsageJSON contains the JSON metadata for the struct
// [LlmV1NewEmbeddingResponseUsage]
type llmV1NewEmbeddingResponseUsageJSON struct {
	PromptTokens apijson.Field
	TotalTokens  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LlmV1NewEmbeddingResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1NewEmbeddingResponseUsageJSON) RawJSON() string {
	return r.raw
}

type LlmV1ListModelsResponse struct {
	Data []LlmV1ListModelsResponseData `json:"data"`
	// Response object type, always 'list'
	Object string                      `json:"object"`
	JSON   llmV1ListModelsResponseJSON `json:"-"`
}

// llmV1ListModelsResponseJSON contains the JSON metadata for the struct
// [LlmV1ListModelsResponse]
type llmV1ListModelsResponseJSON struct {
	Data        apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmV1ListModelsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ListModelsResponseJSON) RawJSON() string {
	return r.raw
}

type LlmV1ListModelsResponseData struct {
	// Unique model identifier
	ID string `json:"id"`
	// Unix timestamp of model creation
	Created int64 `json:"created"`
	// Object type, always 'model'
	Object string `json:"object"`
	// Model provider (openai, anthropic, google, casemark, etc.)
	OwnedBy string                             `json:"owned_by"`
	Pricing LlmV1ListModelsResponseDataPricing `json:"pricing"`
	JSON    llmV1ListModelsResponseDataJSON    `json:"-"`
}

// llmV1ListModelsResponseDataJSON contains the JSON metadata for the struct
// [LlmV1ListModelsResponseData]
type llmV1ListModelsResponseDataJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Object      apijson.Field
	OwnedBy     apijson.Field
	Pricing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmV1ListModelsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ListModelsResponseDataJSON) RawJSON() string {
	return r.raw
}

type LlmV1ListModelsResponseDataPricing struct {
	// Input token price per token
	Input string `json:"input"`
	// Cache read price per token (if supported)
	InputCacheRead string `json:"input_cache_read"`
	// Output token price per token
	Output string                                 `json:"output"`
	JSON   llmV1ListModelsResponseDataPricingJSON `json:"-"`
}

// llmV1ListModelsResponseDataPricingJSON contains the JSON metadata for the struct
// [LlmV1ListModelsResponseDataPricing]
type llmV1ListModelsResponseDataPricingJSON struct {
	Input          apijson.Field
	InputCacheRead apijson.Field
	Output         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LlmV1ListModelsResponseDataPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ListModelsResponseDataPricingJSON) RawJSON() string {
	return r.raw
}

type LlmV1NewEmbeddingParams struct {
	// Text or array of texts to create embeddings for
	Input param.Field[LlmV1NewEmbeddingParamsInputUnion] `json:"input" api:"required"`
	// Embedding model to use (e.g., text-embedding-ada-002, text-embedding-3-small)
	Model param.Field[string] `json:"model" api:"required"`
	// Number of dimensions for the embeddings (model-specific)
	Dimensions param.Field[int64] `json:"dimensions"`
	// Format for returned embeddings
	EncodingFormat param.Field[LlmV1NewEmbeddingParamsEncodingFormat] `json:"encoding_format"`
	// Unique identifier for the end-user
	User param.Field[string] `json:"user"`
}

func (r LlmV1NewEmbeddingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Text or array of texts to create embeddings for
//
// Satisfied by [shared.UnionString], [LlmV1NewEmbeddingParamsInputArray].
type LlmV1NewEmbeddingParamsInputUnion interface {
	ImplementsLlmV1NewEmbeddingParamsInputUnion()
}

type LlmV1NewEmbeddingParamsInputArray []string

func (r LlmV1NewEmbeddingParamsInputArray) ImplementsLlmV1NewEmbeddingParamsInputUnion() {}

// Format for returned embeddings
type LlmV1NewEmbeddingParamsEncodingFormat string

const (
	LlmV1NewEmbeddingParamsEncodingFormatFloat  LlmV1NewEmbeddingParamsEncodingFormat = "float"
	LlmV1NewEmbeddingParamsEncodingFormatBase64 LlmV1NewEmbeddingParamsEncodingFormat = "base64"
)

func (r LlmV1NewEmbeddingParamsEncodingFormat) IsKnown() bool {
	switch r {
	case LlmV1NewEmbeddingParamsEncodingFormatFloat, LlmV1NewEmbeddingParamsEncodingFormatBase64:
		return true
	}
	return false
}
