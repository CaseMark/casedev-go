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
// AgentV1ChatService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV1ChatService] method instead.
type AgentV1ChatService struct {
	Options []option.RequestOption
}

// NewAgentV1ChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV1ChatService(opts ...option.RequestOption) (r *AgentV1ChatService) {
	r = &AgentV1ChatService{}
	r.Options = opts
	return
}

// Creates a persistent OpenCode chat session in a Modal sandbox. Session state is
// retained and can be resumed across requests.
func (r *AgentV1ChatService) New(ctx context.Context, body AgentV1ChatNewParams, opts ...option.RequestOption) (res *AgentV1ChatNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Snapshots and terminates the active sandbox (if any), then marks the chat as
// ended.
func (r *AgentV1ChatService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1ChatDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v1/chat/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Aborts the active OpenCode generation for this chat session.
func (r *AgentV1ChatService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1ChatCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v1/chat/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Answers a pending OpenCode question for the chat session bound to this agent
// chat.
func (r *AgentV1ChatService) ReplyToQuestion(ctx context.Context, id string, requestID string, body AgentV1ChatReplyToQuestionParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("agent/v1/chat/%s/question/%s/reply", id, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Streams a single assistant turn as normalized SSE events with stable turn,
// message, and part IDs. Emits events: `turn.started`, `turn.status`,
// `message.created`, `message.part.updated`, `message.completed`, `session.usage`,
// `turn.completed`.
//
// **When to use this endpoint:** Recommended for building custom chat UIs that
// need real-time streaming progress. This is the primary streaming endpoint for
// new integrations.
//
// **Alternatives:**
//
//   - `POST /chat/:id/message` — synchronous, returns complete response as JSON
//     (best for server-to-server)
func (r *AgentV1ChatService) RespondStreaming(ctx context.Context, id string, body AgentV1ChatRespondParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
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
	path := fmt.Sprintf("agent/v1/chat/%s/respond", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

// Sends a message and returns the complete response as a single JSON body. Blocks
// until the agent turn completes.
//
// **When to use this endpoint:** Best for server-to-server integrations,
// background processing, or any context where you want the full response in one
// call without managing an SSE stream.
//
// **Alternatives:**
//
//   - `POST /chat/:id/respond` — streaming SSE with normalized events (recommended
//     for custom chat UIs)
func (r *AgentV1ChatService) SendMessage(ctx context.Context, id string, body AgentV1ChatSendMessageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/v1/chat/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Relays OpenCode SSE events for this chat. Supports replay from buffered events
// using Last-Event-ID.
func (r *AgentV1ChatService) StreamStreaming(ctx context.Context, id string, query AgentV1ChatStreamParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
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
	path := fmt.Sprintf("agent/v1/chat/%s/stream", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

type AgentV1ChatNewResponse struct {
	ID            string                     `json:"id"`
	CreatedAt     time.Time                  `json:"createdAt" format:"date-time"`
	IdleTimeoutMs int64                      `json:"idleTimeoutMs"`
	Status        string                     `json:"status"`
	JSON          agentV1ChatNewResponseJSON `json:"-"`
}

// agentV1ChatNewResponseJSON contains the JSON metadata for the struct
// [AgentV1ChatNewResponse]
type agentV1ChatNewResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	IdleTimeoutMs apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentV1ChatNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1ChatNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1ChatDeleteResponse struct {
	ID              string                        `json:"id"`
	Cost            float64                       `json:"cost"`
	RuntimeMs       int64                         `json:"runtimeMs"`
	SnapshotImageID string                        `json:"snapshotImageId" api:"nullable"`
	Status          string                        `json:"status"`
	JSON            agentV1ChatDeleteResponseJSON `json:"-"`
}

// agentV1ChatDeleteResponseJSON contains the JSON metadata for the struct
// [AgentV1ChatDeleteResponse]
type agentV1ChatDeleteResponseJSON struct {
	ID              apijson.Field
	Cost            apijson.Field
	RuntimeMs       apijson.Field
	SnapshotImageID apijson.Field
	Status          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AgentV1ChatDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1ChatDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1ChatCancelResponse struct {
	ID   string                        `json:"id"`
	Ok   bool                          `json:"ok"`
	JSON agentV1ChatCancelResponseJSON `json:"-"`
}

// agentV1ChatCancelResponseJSON contains the JSON metadata for the struct
// [AgentV1ChatCancelResponse]
type agentV1ChatCancelResponseJSON struct {
	ID          apijson.Field
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1ChatCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1ChatCancelResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1ChatNewParams struct {
	// Idle timeout before session is eligible for snapshot/termination. Defaults to 15
	// minutes.
	IdleTimeoutMs param.Field[int64] `json:"idleTimeoutMs"`
	// Optional model override for the OpenCode session
	Model param.Field[string] `json:"model"`
	// Optional human-readable session title
	Title param.Field[string] `json:"title"`
	// Restrict the chat session to specific vault IDs
	VaultIDs param.Field[[]string] `json:"vaultIds"`
}

func (r AgentV1ChatNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1ChatReplyToQuestionParams struct {
	// Answer selections for each prompt element in the pending question
	Answers param.Field[[][]string] `json:"answers" api:"required"`
}

func (r AgentV1ChatReplyToQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1ChatRespondParams struct {
	// Message content parts. Currently only "text" type is supported. Additional types
	// (e.g. file, image) may be added in future versions.
	Parts param.Field[[]AgentV1ChatRespondParamsPart] `json:"parts"`
}

func (r AgentV1ChatRespondParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1ChatRespondParamsPart struct {
	// The message text content
	Text param.Field[string] `json:"text" api:"required"`
	// Part type. Currently only "text" is supported.
	Type param.Field[AgentV1ChatRespondParamsPartsType] `json:"type" api:"required"`
}

func (r AgentV1ChatRespondParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Part type. Currently only "text" is supported.
type AgentV1ChatRespondParamsPartsType string

const (
	AgentV1ChatRespondParamsPartsTypeText AgentV1ChatRespondParamsPartsType = "text"
)

func (r AgentV1ChatRespondParamsPartsType) IsKnown() bool {
	switch r {
	case AgentV1ChatRespondParamsPartsTypeText:
		return true
	}
	return false
}

type AgentV1ChatSendMessageParams struct {
	// Message content parts. Currently only "text" type is supported. Additional types
	// (e.g. file, image) may be added in future versions.
	Parts param.Field[[]AgentV1ChatSendMessageParamsPart] `json:"parts"`
}

func (r AgentV1ChatSendMessageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1ChatSendMessageParamsPart struct {
	// The message text content
	Text param.Field[string] `json:"text" api:"required"`
	// Part type. Currently only "text" is supported.
	Type param.Field[AgentV1ChatSendMessageParamsPartsType] `json:"type" api:"required"`
}

func (r AgentV1ChatSendMessageParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Part type. Currently only "text" is supported.
type AgentV1ChatSendMessageParamsPartsType string

const (
	AgentV1ChatSendMessageParamsPartsTypeText AgentV1ChatSendMessageParamsPartsType = "text"
)

func (r AgentV1ChatSendMessageParamsPartsType) IsKnown() bool {
	switch r {
	case AgentV1ChatSendMessageParamsPartsTypeText:
		return true
	}
	return false
}

type AgentV1ChatStreamParams struct {
	// Replay events after this sequence number
	LastEventID param.Field[int64] `query:"lastEventId"`
}

// URLQuery serializes [AgentV1ChatStreamParams]'s query parameters as
// `url.Values`.
func (r AgentV1ChatStreamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
