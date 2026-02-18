// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// ApplicationService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationService] method instead.
type ApplicationService struct {
	Options []option.RequestOption
	V1      *ApplicationV1Service
}

// NewApplicationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApplicationService(opts ...option.RequestOption) (r *ApplicationService) {
	r = &ApplicationService{}
	r.Options = opts
	r.V1 = NewApplicationV1Service(opts...)
	return
}
