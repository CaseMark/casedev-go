// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// AgentV1RunService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV1RunService] method instead.
type AgentV1RunService struct {
	Options []option.RequestOption
}

// NewAgentV1RunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV1RunService(opts ...option.RequestOption) (r *AgentV1RunService) {
	r = &AgentV1RunService{}
	r.Options = opts
	return
}

// Creates a run in queued state. Call POST /agent/v1/run/:id/exec to start
// execution.
func (r *AgentV1RunService) New(ctx context.Context, body AgentV1RunNewParams, opts ...option.RequestOption) (res *AgentV1RunNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v1/run"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cancels a running or queued run. Idempotent — cancelling a finished run returns
// its current status.
func (r *AgentV1RunService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1RunCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/run/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Starts execution of a queued run. The agent runs in a durable workflow — poll
// /run/:id/status for progress.
func (r *AgentV1RunService) Exec(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1RunExecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/run/%s/exec", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Full audit trail for a run including output, steps (tool calls, text), and token
// usage.
func (r *AgentV1RunService) GetDetails(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1RunGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/run/%s/details", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lightweight status poll for a run. Use /run/:id/details for the full audit
// trail.
func (r *AgentV1RunService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1RunGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/run/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Register a callback URL to receive notifications when the run completes. URL
// must use https and must not point to a private network.
func (r *AgentV1RunService) Watch(ctx context.Context, id string, body AgentV1RunWatchParams, opts ...option.RequestOption) (res *AgentV1RunWatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/run/%s/watch", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentV1RunNewResponse struct {
	ID        string                      `json:"id"`
	AgentID   string                      `json:"agentId"`
	CreatedAt time.Time                   `json:"createdAt" format:"date-time"`
	Status    AgentV1RunNewResponseStatus `json:"status"`
	JSON      agentV1RunNewResponseJSON   `json:"-"`
}

// agentV1RunNewResponseJSON contains the JSON metadata for the struct
// [AgentV1RunNewResponse]
type agentV1RunNewResponseJSON struct {
	ID          apijson.Field
	AgentID     apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunNewResponseStatus string

const (
	AgentV1RunNewResponseStatusQueued AgentV1RunNewResponseStatus = "queued"
)

func (r AgentV1RunNewResponseStatus) IsKnown() bool {
	switch r {
	case AgentV1RunNewResponseStatusQueued:
		return true
	}
	return false
}

type AgentV1RunCancelResponse struct {
	ID string `json:"id"`
	// Present if run was already finished
	Message string                         `json:"message"`
	Status  AgentV1RunCancelResponseStatus `json:"status"`
	JSON    agentV1RunCancelResponseJSON   `json:"-"`
}

// agentV1RunCancelResponseJSON contains the JSON metadata for the struct
// [AgentV1RunCancelResponse]
type agentV1RunCancelResponseJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunCancelResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunCancelResponseStatus string

const (
	AgentV1RunCancelResponseStatusCancelled AgentV1RunCancelResponseStatus = "cancelled"
	AgentV1RunCancelResponseStatusCompleted AgentV1RunCancelResponseStatus = "completed"
	AgentV1RunCancelResponseStatusFailed    AgentV1RunCancelResponseStatus = "failed"
)

func (r AgentV1RunCancelResponseStatus) IsKnown() bool {
	switch r {
	case AgentV1RunCancelResponseStatusCancelled, AgentV1RunCancelResponseStatusCompleted, AgentV1RunCancelResponseStatusFailed:
		return true
	}
	return false
}

type AgentV1RunExecResponse struct {
	ID      string                       `json:"id"`
	Message string                       `json:"message"`
	Status  AgentV1RunExecResponseStatus `json:"status"`
	// Durable workflow run ID
	WorkflowID string                     `json:"workflowId"`
	JSON       agentV1RunExecResponseJSON `json:"-"`
}

// agentV1RunExecResponseJSON contains the JSON metadata for the struct
// [AgentV1RunExecResponse]
type agentV1RunExecResponseJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Status      apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunExecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunExecResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunExecResponseStatus string

const (
	AgentV1RunExecResponseStatusRunning AgentV1RunExecResponseStatus = "running"
)

func (r AgentV1RunExecResponseStatus) IsKnown() bool {
	switch r {
	case AgentV1RunExecResponseStatusRunning:
		return true
	}
	return false
}

type AgentV1RunGetDetailsResponse struct {
	ID          string    `json:"id"`
	AgentID     string    `json:"agentId"`
	CompletedAt time.Time `json:"completedAt" api:"nullable" format:"date-time"`
	CreatedAt   time.Time `json:"createdAt" format:"date-time"`
	Guidance    string    `json:"guidance" api:"nullable"`
	Model       string    `json:"model" api:"nullable"`
	Prompt      string    `json:"prompt"`
	// Final output from the agent
	Result    AgentV1RunGetDetailsResponseResult `json:"result" api:"nullable"`
	StartedAt time.Time                          `json:"startedAt" api:"nullable" format:"date-time"`
	Status    AgentV1RunGetDetailsResponseStatus `json:"status"`
	Steps     []AgentV1RunGetDetailsResponseStep `json:"steps"`
	// Token usage statistics
	Usage AgentV1RunGetDetailsResponseUsage `json:"usage" api:"nullable"`
	JSON  agentV1RunGetDetailsResponseJSON  `json:"-"`
}

// agentV1RunGetDetailsResponseJSON contains the JSON metadata for the struct
// [AgentV1RunGetDetailsResponse]
type agentV1RunGetDetailsResponseJSON struct {
	ID          apijson.Field
	AgentID     apijson.Field
	CompletedAt apijson.Field
	CreatedAt   apijson.Field
	Guidance    apijson.Field
	Model       apijson.Field
	Prompt      apijson.Field
	Result      apijson.Field
	StartedAt   apijson.Field
	Status      apijson.Field
	Steps       apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// Final output from the agent
type AgentV1RunGetDetailsResponseResult struct {
	// Sandbox execution logs (OpenCode server + runner script)
	Logs   AgentV1RunGetDetailsResponseResultLogs `json:"logs" api:"nullable"`
	Output string                                 `json:"output"`
	JSON   agentV1RunGetDetailsResponseResultJSON `json:"-"`
}

// agentV1RunGetDetailsResponseResultJSON contains the JSON metadata for the struct
// [AgentV1RunGetDetailsResponseResult]
type agentV1RunGetDetailsResponseResultJSON struct {
	Logs        apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunGetDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunGetDetailsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Sandbox execution logs (OpenCode server + runner script)
type AgentV1RunGetDetailsResponseResultLogs struct {
	// OpenCode server stdout/stderr
	Opencode string `json:"opencode"`
	// Runner script stdout/stderr
	Runner string                                     `json:"runner"`
	JSON   agentV1RunGetDetailsResponseResultLogsJSON `json:"-"`
}

// agentV1RunGetDetailsResponseResultLogsJSON contains the JSON metadata for the
// struct [AgentV1RunGetDetailsResponseResultLogs]
type agentV1RunGetDetailsResponseResultLogsJSON struct {
	Opencode    apijson.Field
	Runner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunGetDetailsResponseResultLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunGetDetailsResponseResultLogsJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunGetDetailsResponseStatus string

const (
	AgentV1RunGetDetailsResponseStatusQueued    AgentV1RunGetDetailsResponseStatus = "queued"
	AgentV1RunGetDetailsResponseStatusRunning   AgentV1RunGetDetailsResponseStatus = "running"
	AgentV1RunGetDetailsResponseStatusCompleted AgentV1RunGetDetailsResponseStatus = "completed"
	AgentV1RunGetDetailsResponseStatusFailed    AgentV1RunGetDetailsResponseStatus = "failed"
	AgentV1RunGetDetailsResponseStatusCancelled AgentV1RunGetDetailsResponseStatus = "cancelled"
)

func (r AgentV1RunGetDetailsResponseStatus) IsKnown() bool {
	switch r {
	case AgentV1RunGetDetailsResponseStatusQueued, AgentV1RunGetDetailsResponseStatusRunning, AgentV1RunGetDetailsResponseStatusCompleted, AgentV1RunGetDetailsResponseStatusFailed, AgentV1RunGetDetailsResponseStatusCancelled:
		return true
	}
	return false
}

type AgentV1RunGetDetailsResponseStep struct {
	ID         string                                `json:"id"`
	Content    string                                `json:"content" api:"nullable"`
	DurationMs int64                                 `json:"durationMs" api:"nullable"`
	Timestamp  time.Time                             `json:"timestamp" format:"date-time"`
	ToolInput  interface{}                           `json:"toolInput"`
	ToolName   string                                `json:"toolName" api:"nullable"`
	ToolOutput interface{}                           `json:"toolOutput"`
	Type       AgentV1RunGetDetailsResponseStepsType `json:"type"`
	JSON       agentV1RunGetDetailsResponseStepJSON  `json:"-"`
}

// agentV1RunGetDetailsResponseStepJSON contains the JSON metadata for the struct
// [AgentV1RunGetDetailsResponseStep]
type agentV1RunGetDetailsResponseStepJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	DurationMs  apijson.Field
	Timestamp   apijson.Field
	ToolInput   apijson.Field
	ToolName    apijson.Field
	ToolOutput  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunGetDetailsResponseStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunGetDetailsResponseStepJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunGetDetailsResponseStepsType string

const (
	AgentV1RunGetDetailsResponseStepsTypeOutput     AgentV1RunGetDetailsResponseStepsType = "output"
	AgentV1RunGetDetailsResponseStepsTypeThinking   AgentV1RunGetDetailsResponseStepsType = "thinking"
	AgentV1RunGetDetailsResponseStepsTypeToolCall   AgentV1RunGetDetailsResponseStepsType = "tool_call"
	AgentV1RunGetDetailsResponseStepsTypeToolResult AgentV1RunGetDetailsResponseStepsType = "tool_result"
)

func (r AgentV1RunGetDetailsResponseStepsType) IsKnown() bool {
	switch r {
	case AgentV1RunGetDetailsResponseStepsTypeOutput, AgentV1RunGetDetailsResponseStepsTypeThinking, AgentV1RunGetDetailsResponseStepsTypeToolCall, AgentV1RunGetDetailsResponseStepsTypeToolResult:
		return true
	}
	return false
}

// Token usage statistics
type AgentV1RunGetDetailsResponseUsage struct {
	DurationMs   int64                                 `json:"durationMs"`
	InputTokens  int64                                 `json:"inputTokens"`
	Model        string                                `json:"model"`
	OutputTokens int64                                 `json:"outputTokens"`
	ToolCalls    int64                                 `json:"toolCalls"`
	JSON         agentV1RunGetDetailsResponseUsageJSON `json:"-"`
}

// agentV1RunGetDetailsResponseUsageJSON contains the JSON metadata for the struct
// [AgentV1RunGetDetailsResponseUsage]
type agentV1RunGetDetailsResponseUsageJSON struct {
	DurationMs   apijson.Field
	InputTokens  apijson.Field
	Model        apijson.Field
	OutputTokens apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentV1RunGetDetailsResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunGetDetailsResponseUsageJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunGetStatusResponse struct {
	ID          string    `json:"id"`
	CompletedAt time.Time `json:"completedAt" api:"nullable" format:"date-time"`
	// Elapsed time in milliseconds
	DurationMs int64                             `json:"durationMs" api:"nullable"`
	StartedAt  time.Time                         `json:"startedAt" api:"nullable" format:"date-time"`
	Status     AgentV1RunGetStatusResponseStatus `json:"status"`
	JSON       agentV1RunGetStatusResponseJSON   `json:"-"`
}

// agentV1RunGetStatusResponseJSON contains the JSON metadata for the struct
// [AgentV1RunGetStatusResponse]
type agentV1RunGetStatusResponseJSON struct {
	ID          apijson.Field
	CompletedAt apijson.Field
	DurationMs  apijson.Field
	StartedAt   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunGetStatusResponseStatus string

const (
	AgentV1RunGetStatusResponseStatusQueued    AgentV1RunGetStatusResponseStatus = "queued"
	AgentV1RunGetStatusResponseStatusRunning   AgentV1RunGetStatusResponseStatus = "running"
	AgentV1RunGetStatusResponseStatusCompleted AgentV1RunGetStatusResponseStatus = "completed"
	AgentV1RunGetStatusResponseStatusFailed    AgentV1RunGetStatusResponseStatus = "failed"
	AgentV1RunGetStatusResponseStatusCancelled AgentV1RunGetStatusResponseStatus = "cancelled"
)

func (r AgentV1RunGetStatusResponseStatus) IsKnown() bool {
	switch r {
	case AgentV1RunGetStatusResponseStatusQueued, AgentV1RunGetStatusResponseStatusRunning, AgentV1RunGetStatusResponseStatusCompleted, AgentV1RunGetStatusResponseStatusFailed, AgentV1RunGetStatusResponseStatusCancelled:
		return true
	}
	return false
}

type AgentV1RunWatchResponse struct {
	CallbackURL string                      `json:"callbackUrl"`
	Ok          bool                        `json:"ok"`
	JSON        agentV1RunWatchResponseJSON `json:"-"`
}

// agentV1RunWatchResponseJSON contains the JSON metadata for the struct
// [AgentV1RunWatchResponse]
type agentV1RunWatchResponseJSON struct {
	CallbackURL apijson.Field
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1RunWatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1RunWatchResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1RunNewParams struct {
	// ID of the agent to run
	AgentID param.Field[string] `json:"agentId" api:"required"`
	// Task prompt for the agent
	Prompt param.Field[string] `json:"prompt" api:"required"`
	// Additional guidance for this run
	Guidance param.Field[string] `json:"guidance"`
	// Override the agent default model for this run
	Model param.Field[string] `json:"model"`
}

func (r AgentV1RunNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1RunWatchParams struct {
	// HTTPS URL to receive completion callback
	CallbackURL param.Field[string] `json:"callbackUrl" api:"required" format:"uri"`
}

func (r AgentV1RunWatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
