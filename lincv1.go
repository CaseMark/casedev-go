// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// LincV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLincV1Service] method instead.
type LincV1Service struct {
	Options []option.RequestOption
	// Create, manage, and execute AI agents with tool access, sandbox environments,
	// and async run workflows
	Sessions *LincV1SessionService
}

// NewLincV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLincV1Service(opts ...option.RequestOption) (r *LincV1Service) {
	r = &LincV1Service{}
	r.Options = opts
	r.Sessions = NewLincV1SessionService(opts...)
	return
}
