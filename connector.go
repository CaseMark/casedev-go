// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// ConnectorService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorService] method instead.
type ConnectorService struct {
	Options []option.RequestOption
	// Import and export between provider folders (Google Drive) and vaults
	V1 *ConnectorV1Service
}

// NewConnectorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorService(opts ...option.RequestOption) (r *ConnectorService) {
	r = &ConnectorService{}
	r.Options = opts
	r.V1 = NewConnectorV1Service(opts...)
	return
}
