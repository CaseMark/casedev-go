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
// AgentV2RunService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV2RunService] method instead.
type AgentV2RunService struct {
	Options []option.RequestOption
}

// NewAgentV2RunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV2RunService(opts ...option.RequestOption) (r *AgentV2RunService) {
	r = &AgentV2RunService{}
	r.Options = opts
	return
}

// Creates a v2 run in queued state. Call POST /agent/v2/run/:id/exec to start
// execution on the Daytona runtime.
func (r *AgentV2RunService) New(ctx context.Context, body AgentV2RunNewParams, opts ...option.RequestOption) (res *AgentV2RunNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v2/run"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Streams real-time v2 run events over SSE with replay support.
func (r *AgentV2RunService) EventsStreaming(ctx context.Context, id string, query AgentV2RunEventsParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
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
	path := fmt.Sprintf("agent/v2/run/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

// Starts execution of a queued v2 run. The agent runs in a durable workflow on a
// Daytona runtime.
func (r *AgentV2RunService) Exec(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV2RunExecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/run/%s/exec", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Full audit trail for a v2 run, with provider-neutral runtime metadata.
func (r *AgentV2RunService) GetDetails(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV2RunGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/run/%s/details", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lightweight status poll for a v2 run including neutral runtime metadata.
func (r *AgentV2RunService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV2RunGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/run/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentV2RunNewResponse struct {
	ID        string                      `json:"id"`
	AgentID   string                      `json:"agentId"`
	CreatedAt time.Time                   `json:"createdAt" format:"date-time"`
	ObjectIDs []string                    `json:"objectIds" api:"nullable"`
	Status    AgentV2RunNewResponseStatus `json:"status"`
	JSON      agentV2RunNewResponseJSON   `json:"-"`
}

// agentV2RunNewResponseJSON contains the JSON metadata for the struct
// [AgentV2RunNewResponse]
type agentV2RunNewResponseJSON struct {
	ID          apijson.Field
	AgentID     apijson.Field
	CreatedAt   apijson.Field
	ObjectIDs   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV2RunNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2RunNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2RunNewResponseStatus string

const (
	AgentV2RunNewResponseStatusQueued AgentV2RunNewResponseStatus = "queued"
)

func (r AgentV2RunNewResponseStatus) IsKnown() bool {
	switch r {
	case AgentV2RunNewResponseStatusQueued:
		return true
	}
	return false
}

type AgentV2RunExecResponse struct {
	ID           string                             `json:"id"`
	Message      string                             `json:"message"`
	Provider     AgentV2RunExecResponseProvider     `json:"provider"`
	RuntimeState AgentV2RunExecResponseRuntimeState `json:"runtimeState"`
	Status       AgentV2RunExecResponseStatus       `json:"status"`
	WorkflowID   string                             `json:"workflowId"`
	JSON         agentV2RunExecResponseJSON         `json:"-"`
}

// agentV2RunExecResponseJSON contains the JSON metadata for the struct
// [AgentV2RunExecResponse]
type agentV2RunExecResponseJSON struct {
	ID           apijson.Field
	Message      apijson.Field
	Provider     apijson.Field
	RuntimeState apijson.Field
	Status       apijson.Field
	WorkflowID   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentV2RunExecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2RunExecResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2RunExecResponseProvider string

const (
	AgentV2RunExecResponseProviderDaytona AgentV2RunExecResponseProvider = "daytona"
)

func (r AgentV2RunExecResponseProvider) IsKnown() bool {
	switch r {
	case AgentV2RunExecResponseProviderDaytona:
		return true
	}
	return false
}

type AgentV2RunExecResponseRuntimeState string

const (
	AgentV2RunExecResponseRuntimeStateRunning AgentV2RunExecResponseRuntimeState = "running"
)

func (r AgentV2RunExecResponseRuntimeState) IsKnown() bool {
	switch r {
	case AgentV2RunExecResponseRuntimeStateRunning:
		return true
	}
	return false
}

type AgentV2RunExecResponseStatus string

const (
	AgentV2RunExecResponseStatusRunning AgentV2RunExecResponseStatus = "running"
)

func (r AgentV2RunExecResponseStatus) IsKnown() bool {
	switch r {
	case AgentV2RunExecResponseStatusRunning:
		return true
	}
	return false
}

type AgentV2RunGetDetailsResponse = interface{}

type AgentV2RunGetStatusResponse struct {
	ID           string                            `json:"id"`
	CompletedAt  time.Time                         `json:"completedAt" api:"nullable" format:"date-time"`
	DurationMs   int64                             `json:"durationMs" api:"nullable"`
	Provider     string                            `json:"provider" api:"nullable"`
	RuntimeID    string                            `json:"runtimeId" api:"nullable"`
	RuntimeState string                            `json:"runtimeState" api:"nullable"`
	StartedAt    time.Time                         `json:"startedAt" api:"nullable" format:"date-time"`
	Status       AgentV2RunGetStatusResponseStatus `json:"status"`
	JSON         agentV2RunGetStatusResponseJSON   `json:"-"`
}

// agentV2RunGetStatusResponseJSON contains the JSON metadata for the struct
// [AgentV2RunGetStatusResponse]
type agentV2RunGetStatusResponseJSON struct {
	ID           apijson.Field
	CompletedAt  apijson.Field
	DurationMs   apijson.Field
	Provider     apijson.Field
	RuntimeID    apijson.Field
	RuntimeState apijson.Field
	StartedAt    apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentV2RunGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2RunGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2RunGetStatusResponseStatus string

const (
	AgentV2RunGetStatusResponseStatusQueued    AgentV2RunGetStatusResponseStatus = "queued"
	AgentV2RunGetStatusResponseStatusRunning   AgentV2RunGetStatusResponseStatus = "running"
	AgentV2RunGetStatusResponseStatusCompleted AgentV2RunGetStatusResponseStatus = "completed"
	AgentV2RunGetStatusResponseStatusFailed    AgentV2RunGetStatusResponseStatus = "failed"
	AgentV2RunGetStatusResponseStatusCancelled AgentV2RunGetStatusResponseStatus = "cancelled"
)

func (r AgentV2RunGetStatusResponseStatus) IsKnown() bool {
	switch r {
	case AgentV2RunGetStatusResponseStatusQueued, AgentV2RunGetStatusResponseStatusRunning, AgentV2RunGetStatusResponseStatusCompleted, AgentV2RunGetStatusResponseStatusFailed, AgentV2RunGetStatusResponseStatusCancelled:
		return true
	}
	return false
}

type AgentV2RunNewParams struct {
	AgentID     param.Field[string]   `json:"agentId" api:"required"`
	Prompt      param.Field[string]   `json:"prompt" api:"required"`
	CallbackURL param.Field[string]   `json:"callbackUrl" format:"uri"`
	Guidance    param.Field[string]   `json:"guidance"`
	Model       param.Field[string]   `json:"model"`
	ObjectIDs   param.Field[[]string] `json:"objectIds"`
}

func (r AgentV2RunNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV2RunEventsParams struct {
	LastEventID param.Field[int64] `query:"lastEventId"`
}

// URLQuery serializes [AgentV2RunEventsParams]'s query parameters as `url.Values`.
func (r AgentV2RunEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
