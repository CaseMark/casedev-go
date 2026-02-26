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

// AgentV1AgentService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV1AgentService] method instead.
type AgentV1AgentService struct {
	Options []option.RequestOption
}

// NewAgentV1AgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV1AgentService(opts ...option.RequestOption) (r *AgentV1AgentService) {
	r = &AgentV1AgentService{}
	r.Options = opts
	return
}

// Creates a new agent definition with a scoped API key. The agent can then be used
// to create and execute runs.
func (r *AgentV1AgentService) New(ctx context.Context, body AgentV1AgentNewParams, opts ...option.RequestOption) (res *AgentV1AgentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a single agent definition by ID.
func (r *AgentV1AgentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1AgentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an agent definition. Only provided fields are changed.
func (r *AgentV1AgentService) Update(ctx context.Context, id string, body AgentV1AgentUpdateParams, opts ...option.RequestOption) (res *AgentV1AgentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists all active agents for the authenticated organization.
func (r *AgentV1AgentService) List(ctx context.Context, opts ...option.RequestOption) (res *AgentV1AgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent/v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Soft-deletes an agent and revokes its scoped API key.
func (r *AgentV1AgentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1AgentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("agent/v1/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AgentV1AgentNewResponse struct {
	ID            string                      `json:"id"`
	CreatedAt     time.Time                   `json:"createdAt" format:"date-time"`
	Description   string                      `json:"description" api:"nullable"`
	DisabledTools []string                    `json:"disabledTools" api:"nullable"`
	EnabledTools  []string                    `json:"enabledTools" api:"nullable"`
	Instructions  string                      `json:"instructions"`
	Model         string                      `json:"model"`
	Name          string                      `json:"name"`
	Sandbox       interface{}                 `json:"sandbox" api:"nullable"`
	UpdatedAt     time.Time                   `json:"updatedAt" format:"date-time"`
	VaultIDs      []string                    `json:"vaultIds" api:"nullable"`
	JSON          agentV1AgentNewResponseJSON `json:"-"`
}

// agentV1AgentNewResponseJSON contains the JSON metadata for the struct
// [AgentV1AgentNewResponse]
type agentV1AgentNewResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	DisabledTools apijson.Field
	EnabledTools  apijson.Field
	Instructions  apijson.Field
	Model         apijson.Field
	Name          apijson.Field
	Sandbox       apijson.Field
	UpdatedAt     apijson.Field
	VaultIDs      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentV1AgentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1AgentNewResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1AgentGetResponse struct {
	ID            string                      `json:"id"`
	CreatedAt     time.Time                   `json:"createdAt" format:"date-time"`
	Description   string                      `json:"description" api:"nullable"`
	DisabledTools []string                    `json:"disabledTools" api:"nullable"`
	EnabledTools  []string                    `json:"enabledTools" api:"nullable"`
	Instructions  string                      `json:"instructions"`
	IsActive      bool                        `json:"isActive"`
	Model         string                      `json:"model"`
	Name          string                      `json:"name"`
	Sandbox       interface{}                 `json:"sandbox" api:"nullable"`
	UpdatedAt     time.Time                   `json:"updatedAt" format:"date-time"`
	VaultIDs      []string                    `json:"vaultIds" api:"nullable"`
	JSON          agentV1AgentGetResponseJSON `json:"-"`
}

// agentV1AgentGetResponseJSON contains the JSON metadata for the struct
// [AgentV1AgentGetResponse]
type agentV1AgentGetResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	DisabledTools apijson.Field
	EnabledTools  apijson.Field
	Instructions  apijson.Field
	IsActive      apijson.Field
	Model         apijson.Field
	Name          apijson.Field
	Sandbox       apijson.Field
	UpdatedAt     apijson.Field
	VaultIDs      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentV1AgentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1AgentGetResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1AgentUpdateResponse struct {
	ID            string                         `json:"id"`
	CreatedAt     time.Time                      `json:"createdAt" format:"date-time"`
	Description   string                         `json:"description" api:"nullable"`
	DisabledTools []string                       `json:"disabledTools" api:"nullable"`
	EnabledTools  []string                       `json:"enabledTools" api:"nullable"`
	Instructions  string                         `json:"instructions"`
	IsActive      bool                           `json:"isActive"`
	Model         string                         `json:"model"`
	Name          string                         `json:"name"`
	Sandbox       interface{}                    `json:"sandbox" api:"nullable"`
	UpdatedAt     time.Time                      `json:"updatedAt" format:"date-time"`
	VaultGroups   []string                       `json:"vaultGroups" api:"nullable"`
	VaultIDs      []string                       `json:"vaultIds" api:"nullable"`
	JSON          agentV1AgentUpdateResponseJSON `json:"-"`
}

// agentV1AgentUpdateResponseJSON contains the JSON metadata for the struct
// [AgentV1AgentUpdateResponse]
type agentV1AgentUpdateResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	DisabledTools apijson.Field
	EnabledTools  apijson.Field
	Instructions  apijson.Field
	IsActive      apijson.Field
	Model         apijson.Field
	Name          apijson.Field
	Sandbox       apijson.Field
	UpdatedAt     apijson.Field
	VaultGroups   apijson.Field
	VaultIDs      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentV1AgentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1AgentUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1AgentListResponse struct {
	Agents []AgentV1AgentListResponseAgent `json:"agents"`
	JSON   agentV1AgentListResponseJSON    `json:"-"`
}

// agentV1AgentListResponseJSON contains the JSON metadata for the struct
// [AgentV1AgentListResponse]
type agentV1AgentListResponseJSON struct {
	Agents      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1AgentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1AgentListResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1AgentListResponseAgent struct {
	ID          string                            `json:"id"`
	CreatedAt   time.Time                         `json:"createdAt" format:"date-time"`
	Description string                            `json:"description" api:"nullable"`
	IsActive    bool                              `json:"isActive"`
	Model       string                            `json:"model"`
	Name        string                            `json:"name"`
	UpdatedAt   time.Time                         `json:"updatedAt" format:"date-time"`
	VaultIDs    []string                          `json:"vaultIds" api:"nullable"`
	JSON        agentV1AgentListResponseAgentJSON `json:"-"`
}

// agentV1AgentListResponseAgentJSON contains the JSON metadata for the struct
// [AgentV1AgentListResponseAgent]
type agentV1AgentListResponseAgentJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	IsActive    apijson.Field
	Model       apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	VaultIDs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1AgentListResponseAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1AgentListResponseAgentJSON) RawJSON() string {
	return r.raw
}

type AgentV1AgentDeleteResponse struct {
	Ok   bool                           `json:"ok"`
	JSON agentV1AgentDeleteResponseJSON `json:"-"`
}

// agentV1AgentDeleteResponseJSON contains the JSON metadata for the struct
// [AgentV1AgentDeleteResponse]
type agentV1AgentDeleteResponseJSON struct {
	Ok          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1AgentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1AgentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1AgentNewParams struct {
	// System instructions that define agent behavior
	Instructions param.Field[string] `json:"instructions" api:"required"`
	// Display name for the agent
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description of the agent
	Description param.Field[string] `json:"description"`
	// Denylist of tools the agent cannot use
	DisabledTools param.Field[[]string] `json:"disabledTools"`
	// Allowlist of tools the agent can use
	EnabledTools param.Field[[]string] `json:"enabledTools"`
	// LLM model identifier (e.g. anthropic/claude-sonnet-4.6). Defaults to
	// anthropic/claude-sonnet-4.6
	Model param.Field[string] `json:"model"`
	// Custom sandbox configuration (cpu, memoryMiB)
	Sandbox param.Field[AgentV1AgentNewParamsSandbox] `json:"sandbox"`
	// Restrict agent to vaults within specific vault group IDs
	VaultGroups param.Field[[]string] `json:"vaultGroups"`
	// Restrict agent to specific vault IDs
	VaultIDs param.Field[[]string] `json:"vaultIds"`
}

func (r AgentV1AgentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom sandbox configuration (cpu, memoryMiB)
type AgentV1AgentNewParamsSandbox struct {
	// Number of CPUs
	CPU param.Field[int64] `json:"cpu"`
	// Memory in MiB
	MemoryMiB param.Field[int64] `json:"memoryMiB"`
}

func (r AgentV1AgentNewParamsSandbox) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentV1AgentUpdateParams struct {
	Description   param.Field[string]      `json:"description"`
	DisabledTools param.Field[[]string]    `json:"disabledTools"`
	EnabledTools  param.Field[[]string]    `json:"enabledTools"`
	Instructions  param.Field[string]      `json:"instructions"`
	Model         param.Field[string]      `json:"model"`
	Name          param.Field[string]      `json:"name"`
	Sandbox       param.Field[interface{}] `json:"sandbox"`
	VaultGroups   param.Field[[]string]    `json:"vaultGroups"`
	VaultIDs      param.Field[[]string]    `json:"vaultIds"`
}

func (r AgentV1AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
