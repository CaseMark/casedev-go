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

// AgentV1ExecuteService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV1ExecuteService] method instead.
type AgentV1ExecuteService struct {
	Options []option.RequestOption
}

// NewAgentV1ExecuteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV1ExecuteService(opts ...option.RequestOption) (r *AgentV1ExecuteService) {
	r = &AgentV1ExecuteService{}
	r.Options = opts
	return
}

// Creates an ephemeral agent and immediately executes a run. Returns the run ID
// for polling status and results. This is the fastest way to run an agent without
// managing agent lifecycle.
func (r *AgentV1ExecuteService) New(ctx context.Context, body AgentV1ExecuteNewParams, opts ...option.RequestOption) (res *AgentV1ExecuteNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentV1ExecuteNewResponse struct {
	// Ephemeral agent ID (auto-created)
	AgentID string `json:"agentId"`
	Message string `json:"message"`
	// Run ID — poll /agent/v1/run/:id/status
	RunID  string                          `json:"runId"`
	Status AgentV1ExecuteNewResponseStatus `json:"status"`
	JSON   agentV1ExecuteNewResponseJSON   `json:"-"`
}

// agentV1ExecuteNewResponseJSON contains the JSON metadata for the struct
// [AgentV1ExecuteNewResponse]
type agentV1ExecuteNewResponseJSON struct {
	AgentID     apijson.Field
	Message     apijson.Field
	RunID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1ExecuteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1ExecuteNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1ExecuteNewResponseStatus string

const (
	AgentV1ExecuteNewResponseStatusRunning AgentV1ExecuteNewResponseStatus = "running"
)

func (r AgentV1ExecuteNewResponseStatus) IsKnown() bool {
	switch r {
	case AgentV1ExecuteNewResponseStatusRunning:
		return true
	}
	return false
}

type AgentV1ExecuteNewParams struct {
	// Task prompt for the agent
	Prompt param.Field[string] `json:"prompt" api:"required"`
	// Denylist of tools the agent cannot use
	DisabledTools param.Field[[]string] `json:"disabledTools"`
	// Allowlist of tools the agent can use
	EnabledTools param.Field[[]string] `json:"enabledTools"`
	// Additional context or constraints for this run
	Guidance param.Field[string] `json:"guidance"`
	// System instructions. Defaults to a general-purpose legal assistant prompt if not
	// provided.
	Instructions param.Field[string] `json:"instructions"`
	// LLM model identifier. Defaults to anthropic/claude-sonnet-4.6
	Model param.Field[string] `json:"model"`
	// Scope this run to specific vault object IDs. The agent will only access these
	// objects.
	ObjectIDs param.Field[[]string] `json:"objectIds"`
	// Custom sandbox resources (cpu, memoryMiB)
	Sandbox param.Field[AgentV1ExecuteNewParamsSandbox] `json:"sandbox"`
	// Restrict agent to specific vault IDs
	VaultIDs param.Field[[]string] `json:"vaultIds"`
}

func (r AgentV1ExecuteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom sandbox resources (cpu, memoryMiB)
type AgentV1ExecuteNewParamsSandbox struct {
	// Number of CPUs
	CPU param.Field[int64] `json:"cpu"`
	// Memory in MiB
	MemoryMiB param.Field[int64] `json:"memoryMiB"`
}

func (r AgentV1ExecuteNewParamsSandbox) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
