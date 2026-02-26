// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// AgentV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV1Service] method instead.
type AgentV1Service struct {
	Options []option.RequestOption
	Agents  *AgentV1AgentService
	Run     *AgentV1RunService
	Execute *AgentV1ExecuteService
}

// NewAgentV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentV1Service(opts ...option.RequestOption) (r *AgentV1Service) {
	r = &AgentV1Service{}
	r.Options = opts
	r.Agents = NewAgentV1AgentService(opts...)
	r.Run = NewAgentV1RunService(opts...)
	r.Execute = NewAgentV1ExecuteService(opts...)
	return
}
