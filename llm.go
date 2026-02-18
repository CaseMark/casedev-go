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

// LlmService contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmService] method instead.
type LlmService struct {
	Options []option.RequestOption
	V1      *LlmV1Service
}

// NewLlmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLlmService(opts ...option.RequestOption) (r *LlmService) {
	r = &LlmService{}
	r.Options = opts
	r.V1 = NewLlmV1Service(opts...)
	return
}

// Retrieves the AI Gateway configuration including all available language models
// and their specifications. This endpoint returns model information compatible
// with the Vercel AI SDK Gateway format, making it easy to integrate with existing
// AI applications.
//
// Use this endpoint to:
//
// - Discover available language models
// - Get model specifications and pricing
// - Configure AI SDK clients
// - Build model selection interfaces
func (r *LlmService) GetConfig(ctx context.Context, opts ...option.RequestOption) (res *LlmGetConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LlmGetConfigResponse struct {
	Models []LlmGetConfigResponseModel `json:"models,required"`
	JSON   llmGetConfigResponseJSON    `json:"-"`
}

// llmGetConfigResponseJSON contains the JSON metadata for the struct
// [LlmGetConfigResponse]
type llmGetConfigResponseJSON struct {
	Models      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmGetConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmGetConfigResponseJSON) RawJSON() string {
	return r.raw
}

type LlmGetConfigResponseModel struct {
	// Unique model identifier
	ID string `json:"id,required"`
	// Type of model (e.g., language, embedding)
	ModelType string `json:"modelType,required"`
	// Human-readable model name
	Name string `json:"name,required"`
	// Model description and capabilities
	Description string `json:"description"`
	// Pricing information for the model
	Pricing interface{} `json:"pricing"`
	// Technical specifications and limits
	Specification interface{}                   `json:"specification"`
	JSON          llmGetConfigResponseModelJSON `json:"-"`
}

// llmGetConfigResponseModelJSON contains the JSON metadata for the struct
// [LlmGetConfigResponseModel]
type llmGetConfigResponseModelJSON struct {
	ID            apijson.Field
	ModelType     apijson.Field
	Name          apijson.Field
	Description   apijson.Field
	Pricing       apijson.Field
	Specification apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LlmGetConfigResponseModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmGetConfigResponseModelJSON) RawJSON() string {
	return r.raw
}
