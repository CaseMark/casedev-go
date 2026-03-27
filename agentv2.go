// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// AgentV2Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV2Service] method instead.
type AgentV2Service struct {
	Options []option.RequestOption
	// Create, manage, and execute AI agents with tool access, sandbox environments,
	// and async run workflows
	Run *AgentV2RunService
	// Create, manage, and execute AI agents with tool access, sandbox environments,
	// and async run workflows
	Execute *AgentV2ExecuteService
	// Create, manage, and execute AI agents with tool access, sandbox environments,
	// and async run workflows
	Chat *AgentV2ChatService
}

// NewAgentV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentV2Service(opts ...option.RequestOption) (r *AgentV2Service) {
	r = &AgentV2Service{}
	r.Options = opts
	r.Run = NewAgentV2RunService(opts...)
	r.Execute = NewAgentV2ExecuteService(opts...)
	r.Chat = NewAgentV2ChatService(opts...)
	return
}
