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

// Streams a single assistant turn as normalized state events with stable turn,
// message, and part ids. Emits session.usage before turn.completed when token data
// is available.
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

// Proxies a message to the OpenCode session bound to this chat.
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

// Streams a single assistant turn as AI SDK UIMessageChunk SSE events for direct
// client rendering.
func (r *AgentV1ChatService) UiStreamStreaming(ctx context.Context, id string, body AgentV1ChatUiStreamParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
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
	path := fmt.Sprintf("agent/v1/chat/%s/ui-stream", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
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
	Answers param.Field[[][]string] `json:"answers" api:"required"`
}

func (r AgentV1ChatReplyToQuestionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1ChatRespondParams struct {
	// OpenCode message payload. Passed through 1:1.
	Body interface{} `json:"body" api:"required"`
}

func (r AgentV1ChatRespondParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AgentV1ChatSendMessageParams struct {
	// OpenCode message payload. Passed through 1:1.
	Body interface{} `json:"body" api:"required"`
}

func (r AgentV1ChatSendMessageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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

type AgentV1ChatUiStreamParams struct {
	// OpenCode message payload. Passed through 1:1.
	Body interface{} `json:"body" api:"required"`
}

func (r AgentV1ChatUiStreamParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
