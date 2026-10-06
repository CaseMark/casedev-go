// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// ConnectorV1ApplicationService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1ApplicationService] method instead.
type ConnectorV1ApplicationService struct {
	Options []option.RequestOption
	// Import and export between provider folders and vaults
	Keys *ConnectorV1ApplicationKeyService
}

// NewConnectorV1ApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorV1ApplicationService(opts ...option.RequestOption) (r *ConnectorV1ApplicationService) {
	r = &ConnectorV1ApplicationService{}
	r.Options = opts
	r.Keys = NewConnectorV1ApplicationKeyService(opts...)
	return
}
