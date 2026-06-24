// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Create, manage, and execute AI agents with tool access, sandbox environments,
// and async run workflows
//
// LincV1SessionService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLincV1SessionService] method instead.
type LincV1SessionService struct {
	Options []option.RequestOption
}

// NewLincV1SessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLincV1SessionService(opts ...option.RequestOption) (r *LincV1SessionService) {
	r = &LincV1SessionService{}
	r.Options = opts
	return
}

// Creates a Daytona-backed native Linc session with scoped Case.dev credentials.
// This endpoint starts the sandbox actor only; messages and event replay use
// separate endpoints.
func (r *LincV1SessionService) New(ctx context.Context, body LincV1SessionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "linc/v1/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// End native Linc session
func (r *LincV1SessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Cancel native Linc session turn
func (r *LincV1SessionService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Runtime ingest endpoint for sandbox runtimes. Frames are persisted for replay;
// terminal frames emit the durable Linc session ended webhook.
func (r *LincV1SessionService) IngestEvents(ctx context.Context, id string, body LincV1SessionIngestEventsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s/events/ingest", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Returns persisted native Pi/Linc event envelopes after the requested cursor.
// Live delivery is handled by the Linc stream service.
func (r *LincV1SessionService) GetEvents(ctx context.Context, id string, query LincV1SessionGetEventsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Returns completed Pi/Linc message entries derived from durable native Linc
// events. This is the stable session-message read model for callers that need to
// persist or recover chat history without depending on a live SSE stream.
func (r *LincV1SessionService) GetMessages(ctx context.Context, id string, query LincV1SessionGetMessagesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s/messages", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Get native Linc session state
func (r *LincV1SessionService) GetState(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s/state", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Forwards a native Pi/Linc RPC command object to the sandbox-local Linc bridge
// unchanged. The route returns after Pi accepts or rejects the command; native
// events are read through the events endpoint.
func (r *LincV1SessionService) SendRpc(ctx context.Context, id string, body LincV1SessionSendRpcParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("linc/v1/sessions/%s/rpc", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type LincV1SessionNewParams struct {
	// Specific document template slugs to inject into the using-document-templates
	// skill.
	DocumentTemplateSlugs param.Field[[]string] `json:"documentTemplateSlugs"`
	IdleTimeoutMs         param.Field[int64]    `json:"idleTimeoutMs"`
	// When true, inject all active org document templates into the
	// using-document-templates skill.
	IncludeDocumentTemplates param.Field[bool] `json:"includeDocumentTemplates"`
	// Privileged C3-only hidden app instructions to append to the sandbox AGENTS.md.
	Instructions param.Field[string] `json:"instructions"`
	Model        param.Field[string] `json:"model"`
	// Optional caller-provided scoped Case.dev API key for the runtime.
	ScopedAPIKey param.Field[string] `json:"scopedApiKey"`
	// Skills API slugs to install into the runtime sandbox before the native session
	// starts.
	SkillSlugs param.Field[[]string] `json:"skillSlugs"`
	Title      param.Field[string]   `json:"title"`
	VaultIDs   param.Field[[]string] `json:"vaultIds"`
}

func (r LincV1SessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LincV1SessionIngestEventsParams struct {
	// Native Linc event frames to persist for replay.
	Frames param.Field[[]LincV1SessionIngestEventsParamsFrame] `json:"frames" api:"required"`
}

func (r LincV1SessionIngestEventsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LincV1SessionIngestEventsParamsFrame struct {
	// Native Linc event payload.
	Event param.Field[map[string]interface{}] `json:"event" api:"required"`
	// Monotonic native event sequence number.
	Seq param.Field[int64] `json:"seq" api:"required"`
	// Native Linc event type.
	Type param.Field[string] `json:"type" api:"required"`
}

func (r LincV1SessionIngestEventsParamsFrame) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LincV1SessionGetEventsParams struct {
	// Alias for cursor. Ignored when cursor is also provided.
	AfterSeq param.Field[int64] `query:"afterSeq"`
	// Replay events with a sequence number greater than this cursor.
	Cursor param.Field[int64] `query:"cursor"`
	// Comma-separated Linc event types to omit from replay.
	ExcludeEventTypes param.Field[[]string] `query:"excludeEventTypes"`
	// Maximum number of events to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [LincV1SessionGetEventsParams]'s query parameters as
// `url.Values`.
func (r LincV1SessionGetEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LincV1SessionGetMessagesParams struct {
	// Alias for cursor. Ignored when cursor is also provided.
	AfterSeq param.Field[int64] `query:"afterSeq"`
	// Replay messages with a source event sequence number greater than this cursor.
	Cursor param.Field[int64] `query:"cursor"`
	// Maximum number of source events to scan for completed messages.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [LincV1SessionGetMessagesParams]'s query parameters as
// `url.Values`.
func (r LincV1SessionGetMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LincV1SessionSendRpcParams struct {
	// Native Pi/Linc RPC command type. Prompt commands also require a string id for
	// idempotency.
	Type param.Field[string] `json:"type" api:"required"`
	// Command idempotency key. Required when type is prompt.
	ID param.Field[string] `json:"id"`
}

func (r LincV1SessionSendRpcParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
