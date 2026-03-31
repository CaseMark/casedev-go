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

// Create, manage, and execute AI agents with tool access, sandbox environments,
// and async run workflows
//
// AgentV2ExecuteService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV2ExecuteService] method instead.
type AgentV2ExecuteService struct {
	Options []option.RequestOption
}

// NewAgentV2ExecuteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV2ExecuteService(opts ...option.RequestOption) (r *AgentV2ExecuteService) {
	r = &AgentV2ExecuteService{}
	r.Options = opts
	return
}

// Creates an ephemeral agent and executes it immediately. By default this uses the
// lightweight synchronous linc runtime on Vercel Sandbox. Set `agentRuntime: true`
// to opt into the legacy Daytona-backed agent runtime.
func (r *AgentV2ExecuteService) New(ctx context.Context, body AgentV2ExecuteNewParams, opts ...option.RequestOption) (res *AgentV2ExecuteNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v2/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AgentV2ExecuteNewResponse struct {
	AgentID      string                                `json:"agentId"`
	Error        string                                `json:"error" api:"nullable"`
	Logs         AgentV2ExecuteNewResponseLogs         `json:"logs" api:"nullable"`
	Message      string                                `json:"message" api:"nullable"`
	Output       string                                `json:"output" api:"nullable"`
	Provider     AgentV2ExecuteNewResponseProvider     `json:"provider"`
	RunID        string                                `json:"runId"`
	RuntimeID    string                                `json:"runtimeId" api:"nullable"`
	RuntimeState AgentV2ExecuteNewResponseRuntimeState `json:"runtimeState"`
	Status       AgentV2ExecuteNewResponseStatus       `json:"status"`
	Usage        interface{}                           `json:"usage" api:"nullable"`
	JSON         agentV2ExecuteNewResponseJSON         `json:"-"`
}

// agentV2ExecuteNewResponseJSON contains the JSON metadata for the struct
// [AgentV2ExecuteNewResponse]
type agentV2ExecuteNewResponseJSON struct {
	AgentID      apijson.Field
	Error        apijson.Field
	Logs         apijson.Field
	Message      apijson.Field
	Output       apijson.Field
	Provider     apijson.Field
	RunID        apijson.Field
	RuntimeID    apijson.Field
	RuntimeState apijson.Field
	Status       apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentV2ExecuteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ExecuteNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2ExecuteNewResponseLogs struct {
	Linc   string                            `json:"linc" api:"nullable"`
	Runner string                            `json:"runner" api:"nullable"`
	JSON   agentV2ExecuteNewResponseLogsJSON `json:"-"`
}

// agentV2ExecuteNewResponseLogsJSON contains the JSON metadata for the struct
// [AgentV2ExecuteNewResponseLogs]
type agentV2ExecuteNewResponseLogsJSON struct {
	Linc        apijson.Field
	Runner      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV2ExecuteNewResponseLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ExecuteNewResponseLogsJSON) RawJSON() string {
	return r.raw
}

type AgentV2ExecuteNewResponseProvider string

const (
	AgentV2ExecuteNewResponseProviderDaytona AgentV2ExecuteNewResponseProvider = "daytona"
	AgentV2ExecuteNewResponseProviderVercel  AgentV2ExecuteNewResponseProvider = "vercel"
)

func (r AgentV2ExecuteNewResponseProvider) IsKnown() bool {
	switch r {
	case AgentV2ExecuteNewResponseProviderDaytona, AgentV2ExecuteNewResponseProviderVercel:
		return true
	}
	return false
}

type AgentV2ExecuteNewResponseRuntimeState string

const (
	AgentV2ExecuteNewResponseRuntimeStateRunning AgentV2ExecuteNewResponseRuntimeState = "running"
	AgentV2ExecuteNewResponseRuntimeStateEnded   AgentV2ExecuteNewResponseRuntimeState = "ended"
	AgentV2ExecuteNewResponseRuntimeStateError   AgentV2ExecuteNewResponseRuntimeState = "error"
)

func (r AgentV2ExecuteNewResponseRuntimeState) IsKnown() bool {
	switch r {
	case AgentV2ExecuteNewResponseRuntimeStateRunning, AgentV2ExecuteNewResponseRuntimeStateEnded, AgentV2ExecuteNewResponseRuntimeStateError:
		return true
	}
	return false
}

type AgentV2ExecuteNewResponseStatus string

const (
	AgentV2ExecuteNewResponseStatusRunning   AgentV2ExecuteNewResponseStatus = "running"
	AgentV2ExecuteNewResponseStatusCompleted AgentV2ExecuteNewResponseStatus = "completed"
	AgentV2ExecuteNewResponseStatusFailed    AgentV2ExecuteNewResponseStatus = "failed"
)

func (r AgentV2ExecuteNewResponseStatus) IsKnown() bool {
	switch r {
	case AgentV2ExecuteNewResponseStatusRunning, AgentV2ExecuteNewResponseStatusCompleted, AgentV2ExecuteNewResponseStatusFailed:
		return true
	}
	return false
}

type AgentV2ExecuteNewParams struct {
	Prompt param.Field[string] `json:"prompt" api:"required"`
	// Set to true to opt into the legacy Daytona-backed agent runtime.
	AgentRuntime  param.Field[bool]                           `json:"agentRuntime"`
	DisabledTools param.Field[[]string]                       `json:"disabledTools"`
	EnabledTools  param.Field[[]string]                       `json:"enabledTools"`
	Guidance      param.Field[string]                         `json:"guidance"`
	Instructions  param.Field[string]                         `json:"instructions"`
	Model         param.Field[string]                         `json:"model"`
	ObjectIDs     param.Field[[]string]                       `json:"objectIds"`
	Sandbox       param.Field[AgentV2ExecuteNewParamsSandbox] `json:"sandbox"`
	VaultIDs      param.Field[[]string]                       `json:"vaultIds"`
}

func (r AgentV2ExecuteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV2ExecuteNewParamsSandbox struct {
	CPU       param.Field[int64] `json:"cpu"`
	MemoryMiB param.Field[int64] `json:"memoryMiB"`
}

func (r AgentV2ExecuteNewParamsSandbox) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
