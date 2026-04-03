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
	"github.com/CaseMark/casedev-go/packages/ssestream"
)

// Create, manage, and execute AI agents with tool access, sandbox environments,
// and async run workflows
//
// AgentV2ChatService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV2ChatService] method instead.
type AgentV2ChatService struct {
	Options []option.RequestOption
	// Create, manage, and execute AI agents with tool access, sandbox environments,
	// and async run workflows
	Files *AgentV2ChatFileService
}

// NewAgentV2ChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV2ChatService(opts ...option.RequestOption) (r *AgentV2ChatService) {
	r = &AgentV2ChatService{}
	r.Options = opts
	r.Files = NewAgentV2ChatFileService(opts...)
	return
}

// Creates a persistent OpenCode chat session backed by a Daytona runtime. Session
// state is retained and can be resumed or recovered across requests.
func (r *AgentV2ChatService) New(ctx context.Context, body AgentV2ChatNewParams, opts ...option.RequestOption) (res *AgentV2ChatNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v2/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Terminates the active Daytona runtime (if any), then marks the chat as ended.
func (r *AgentV2ChatService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV2ChatDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/chat/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Aborts the active OpenCode generation for this Daytona-backed chat session.
func (r *AgentV2ChatService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV2ChatCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/chat/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Answers a pending OpenCode question for the Daytona-backed chat session and
// resumes or recovers the runtime if needed.
func (r *AgentV2ChatService) ReplyToQuestion(ctx context.Context, id string, requestID string, body AgentV2ChatReplyToQuestionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return err
	}
	path := fmt.Sprintf("agent/v2/chat/%s/question/%s/reply", id, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Streams a single assistant turn from a Daytona-backed chat runtime as normalized
// SSE events with stable turn, message, and part IDs. Emits events:
// `turn.started`, `turn.status`, `message.created`, `message.part.updated`,
// `message.completed`, `session.usage`, `turn.completed`.
//
// **When to use this endpoint:** Recommended for building custom chat UIs that
// need real-time streaming progress. This is the primary streaming endpoint for
// new integrations.
//
// **Alternatives:**
//
//   - `POST /chat/:id/message` — synchronous, returns complete response as JSON
//     (best for server-to-server)
func (r *AgentV2ChatService) RespondStreaming(ctx context.Context, id string, body AgentV2ChatRespondParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return ssestream.NewStream[string](nil, err)
	}
	path := fmt.Sprintf("agent/v2/chat/%s/respond", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

// Sends a message to a Daytona-backed chat runtime and returns the complete
// response as a single JSON body. Blocks until the assistant turn completes.
//
// **When to use this endpoint:** Best for server-to-server integrations,
// background processing, or any context where you want the full response in one
// call without managing an SSE stream.
//
// **Alternatives:**
//
//   - `POST /chat/:id/respond` — streaming SSE with normalized events (recommended
//     for custom chat UIs)
func (r *AgentV2ChatService) SendMessage(ctx context.Context, id string, body AgentV2ChatSendMessageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/v2/chat/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Relays OpenCode SSE events for this Daytona-backed chat runtime. Supports replay
// from buffered events using Last-Event-ID and transparently reconnects stopped or
// archived runtimes.
func (r *AgentV2ChatService) StreamStreaming(ctx context.Context, id string, query AgentV2ChatStreamParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return ssestream.NewStream[string](nil, err)
	}
	path := fmt.Sprintf("agent/v2/chat/%s/stream", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

type AgentV2ChatNewResponse struct {
	ID            string                         `json:"id"`
	CreatedAt     time.Time                      `json:"createdAt" format:"date-time"`
	IdleTimeoutMs int64                          `json:"idleTimeoutMs"`
	Provider      AgentV2ChatNewResponseProvider `json:"provider"`
	RuntimeState  string                         `json:"runtimeState"`
	Status        string                         `json:"status"`
	JSON          agentV2ChatNewResponseJSON     `json:"-"`
}

// agentV2ChatNewResponseJSON contains the JSON metadata for the struct
// [AgentV2ChatNewResponse]
type agentV2ChatNewResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IdleTimeoutMs apijson.Field
	Provider      apijson.Field
	RuntimeState  apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentV2ChatNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ChatNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2ChatNewResponseProvider string

const (
	AgentV2ChatNewResponseProviderDaytona AgentV2ChatNewResponseProvider = "daytona"
)

func (r AgentV2ChatNewResponseProvider) IsKnown() bool {
	switch r {
	case AgentV2ChatNewResponseProviderDaytona:
		return true
	}
	return false
}

type AgentV2ChatDeleteResponse struct {
	ID        string                        `json:"id"`
	Cost      float64                       `json:"cost"`
	Provider  string                        `json:"provider" api:"nullable"`
	RuntimeID string                        `json:"runtimeId" api:"nullable"`
	RuntimeMs int64                         `json:"runtimeMs"`
	Status    string                        `json:"status"`
	JSON      agentV2ChatDeleteResponseJSON `json:"-"`
}

// agentV2ChatDeleteResponseJSON contains the JSON metadata for the struct
// [AgentV2ChatDeleteResponse]
type agentV2ChatDeleteResponseJSON struct {
	ID          apijson.Field
	Cost        apijson.Field
	Provider    apijson.Field
	RuntimeID   apijson.Field
	RuntimeMs   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV2ChatDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ChatDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2ChatCancelResponse struct {
	ID   string                        `json:"id"`
	Ok   bool                          `json:"ok"`
	JSON agentV2ChatCancelResponseJSON `json:"-"`
}

// agentV2ChatCancelResponseJSON contains the JSON metadata for the struct
// [AgentV2ChatCancelResponse]
type agentV2ChatCancelResponseJSON struct {
	ID          apijson.Field
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV2ChatCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ChatCancelResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2ChatNewParams struct {
	// Idle timeout before the runtime is eligible to stop. Defaults to 15 minutes.
	IdleTimeoutMs param.Field[int64] `json:"idleTimeoutMs"`
	// Optional model override for the OpenCode session
	Model param.Field[string] `json:"model"`
	// Optional human-readable session title
	Title param.Field[string] `json:"title"`
	// Restrict the chat session to specific vault IDs
	VaultIDs param.Field[[]string] `json:"vaultIds"`
}

func (r AgentV2ChatNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV2ChatReplyToQuestionParams struct {
	// Answer selections for each prompt element in the pending question
	Answers param.Field[[][]string] `json:"answers" api:"required"`
}

func (r AgentV2ChatReplyToQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV2ChatRespondParams struct {
	// Optional model override. When provided, the runtime bootstrap config is updated
	// so subsequent turns use this model. Conversation history is preserved.
	Model param.Field[string] `json:"model"`
	// Message content parts. Currently only "text" type is supported. Additional types
	// (e.g. file, image) may be added in future versions.
	Parts param.Field[[]AgentV2ChatRespondParamsPart] `json:"parts"`
}

func (r AgentV2ChatRespondParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV2ChatRespondParamsPart struct {
	// The message text content
	Text param.Field[string] `json:"text" api:"required"`
	// Part type. Currently only "text" is supported.
	Type param.Field[AgentV2ChatRespondParamsPartsType] `json:"type" api:"required"`
}

func (r AgentV2ChatRespondParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Part type. Currently only "text" is supported.
type AgentV2ChatRespondParamsPartsType string

const (
	AgentV2ChatRespondParamsPartsTypeText AgentV2ChatRespondParamsPartsType = "text"
)

func (r AgentV2ChatRespondParamsPartsType) IsKnown() bool {
	switch r {
	case AgentV2ChatRespondParamsPartsTypeText:
		return true
	}
	return false
}

type AgentV2ChatSendMessageParams struct {
	// Optional model override. When provided, the runtime bootstrap config is updated
	// so subsequent turns use this model. Conversation history is preserved.
	Model param.Field[string] `json:"model"`
	// Message content parts. Currently only "text" type is supported. Additional types
	// (e.g. file, image) may be added in future versions.
	Parts param.Field[[]AgentV2ChatSendMessageParamsPart] `json:"parts"`
}

func (r AgentV2ChatSendMessageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV2ChatSendMessageParamsPart struct {
	// The message text content
	Text param.Field[string] `json:"text" api:"required"`
	// Part type. Currently only "text" is supported.
	Type param.Field[AgentV2ChatSendMessageParamsPartsType] `json:"type" api:"required"`
}

func (r AgentV2ChatSendMessageParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Part type. Currently only "text" is supported.
type AgentV2ChatSendMessageParamsPartsType string

const (
	AgentV2ChatSendMessageParamsPartsTypeText AgentV2ChatSendMessageParamsPartsType = "text"
)

func (r AgentV2ChatSendMessageParamsPartsType) IsKnown() bool {
	switch r {
	case AgentV2ChatSendMessageParamsPartsTypeText:
		return true
	}
	return false
}

type AgentV2ChatStreamParams struct {
	// Replay events after this sequence number
	LastEventID param.Field[int64] `query:"lastEventId"`
}

// URLQuery serializes [AgentV2ChatStreamParams]'s query parameters as
// `url.Values`.
func (r AgentV2ChatStreamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
