// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// SuperdocService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperdocService] method instead.
type SuperdocService struct {
	Options []option.RequestOption
	V1      *SuperdocV1Service
}

// NewSuperdocService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuperdocService(opts ...option.RequestOption) (r *SuperdocService) {
	r = &SuperdocService{}
	r.Options = opts
	r.V1 = NewSuperdocV1Service(opts...)
	return
}
