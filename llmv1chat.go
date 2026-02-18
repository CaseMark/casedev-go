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

// LlmV1ChatService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmV1ChatService] method instead.
type LlmV1ChatService struct {
	Options []option.RequestOption
}

// NewLlmV1ChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLlmV1ChatService(opts ...option.RequestOption) (r *LlmV1ChatService) {
	r = &LlmV1ChatService{}
	r.Options = opts
	return
}

// Create a completion for the provided prompt and parameters. Compatible with
// OpenAI's chat completions API. Supports 40+ models including GPT-4, Claude,
// Gemini, and CaseMark legal AI models. Includes streaming support, token
// counting, and usage tracking.
func (r *LlmV1ChatService) NewCompletion(ctx context.Context, body LlmV1ChatNewCompletionParams, opts ...option.RequestOption) (res *LlmV1ChatNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm/v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LlmV1ChatNewCompletionResponse struct {
	// Unique identifier for the completion
	ID      string                                 `json:"id"`
	Choices []LlmV1ChatNewCompletionResponseChoice `json:"choices"`
	// Unix timestamp of completion creation
	Created int64 `json:"created"`
	// Model used for completion
	Model  string                              `json:"model"`
	Object string                              `json:"object"`
	Usage  LlmV1ChatNewCompletionResponseUsage `json:"usage"`
	JSON   llmV1ChatNewCompletionResponseJSON  `json:"-"`
}

// llmV1ChatNewCompletionResponseJSON contains the JSON metadata for the struct
// [LlmV1ChatNewCompletionResponse]
type llmV1ChatNewCompletionResponseJSON struct {
	ID          apijson.Field
	Choices     apijson.Field
	Created     apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmV1ChatNewCompletionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ChatNewCompletionResponseJSON) RawJSON() string {
	return r.raw
}

type LlmV1ChatNewCompletionResponseChoice struct {
	FinishReason string                                       `json:"finish_reason"`
	Index        int64                                        `json:"index"`
	Message      LlmV1ChatNewCompletionResponseChoicesMessage `json:"message"`
	JSON         llmV1ChatNewCompletionResponseChoiceJSON     `json:"-"`
}

// llmV1ChatNewCompletionResponseChoiceJSON contains the JSON metadata for the
// struct [LlmV1ChatNewCompletionResponseChoice]
type llmV1ChatNewCompletionResponseChoiceJSON struct {
	FinishReason apijson.Field
	Index        apijson.Field
	Message      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LlmV1ChatNewCompletionResponseChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ChatNewCompletionResponseChoiceJSON) RawJSON() string {
	return r.raw
}

type LlmV1ChatNewCompletionResponseChoicesMessage struct {
	Content string                                           `json:"content"`
	Role    string                                           `json:"role"`
	JSON    llmV1ChatNewCompletionResponseChoicesMessageJSON `json:"-"`
}

// llmV1ChatNewCompletionResponseChoicesMessageJSON contains the JSON metadata for
// the struct [LlmV1ChatNewCompletionResponseChoicesMessage]
type llmV1ChatNewCompletionResponseChoicesMessageJSON struct {
	Content     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LlmV1ChatNewCompletionResponseChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ChatNewCompletionResponseChoicesMessageJSON) RawJSON() string {
	return r.raw
}

type LlmV1ChatNewCompletionResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	// Cost in USD
	Cost         float64                                 `json:"cost"`
	PromptTokens int64                                   `json:"prompt_tokens"`
	TotalTokens  int64                                   `json:"total_tokens"`
	JSON         llmV1ChatNewCompletionResponseUsageJSON `json:"-"`
}

// llmV1ChatNewCompletionResponseUsageJSON contains the JSON metadata for the
// struct [LlmV1ChatNewCompletionResponseUsage]
type llmV1ChatNewCompletionResponseUsageJSON struct {
	CompletionTokens apijson.Field
	Cost             apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LlmV1ChatNewCompletionResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r llmV1ChatNewCompletionResponseUsageJSON) RawJSON() string {
	return r.raw
}

type LlmV1ChatNewCompletionParams struct {
	// List of messages comprising the conversation
	Messages param.Field[[]LlmV1ChatNewCompletionParamsMessage] `json:"messages,required"`
	// CaseMark-only: when true, allows reasoning fields in responses. Defaults to
	// false (reasoning is suppressed).
	CasemarkShowReasoning param.Field[bool] `json:"casemark_show_reasoning"`
	// Frequency penalty parameter
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Maximum number of tokens to generate
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Model to use for completion. Defaults to casemark/casemark-core-3 if not
	// specified
	Model param.Field[string] `json:"model"`
	// Presence penalty parameter
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Whether to stream back partial progress
	Stream param.Field[bool] `json:"stream"`
	// Sampling temperature between 0 and 2
	Temperature param.Field[float64] `json:"temperature"`
	// Nucleus sampling parameter
	TopP param.Field[float64] `json:"top_p"`
}

func (r LlmV1ChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LlmV1ChatNewCompletionParamsMessage struct {
	// The contents of the message
	Content param.Field[string] `json:"content"`
	// The role of the message author
	Role param.Field[LlmV1ChatNewCompletionParamsMessagesRole] `json:"role"`
}

func (r LlmV1ChatNewCompletionParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The role of the message author
type LlmV1ChatNewCompletionParamsMessagesRole string

const (
	LlmV1ChatNewCompletionParamsMessagesRoleSystem    LlmV1ChatNewCompletionParamsMessagesRole = "system"
	LlmV1ChatNewCompletionParamsMessagesRoleUser      LlmV1ChatNewCompletionParamsMessagesRole = "user"
	LlmV1ChatNewCompletionParamsMessagesRoleAssistant LlmV1ChatNewCompletionParamsMessagesRole = "assistant"
)

func (r LlmV1ChatNewCompletionParamsMessagesRole) IsKnown() bool {
	switch r {
	case LlmV1ChatNewCompletionParamsMessagesRoleSystem, LlmV1ChatNewCompletionParamsMessagesRoleUser, LlmV1ChatNewCompletionParamsMessagesRoleAssistant:
		return true
	}
	return false
}
