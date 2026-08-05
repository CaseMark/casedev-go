// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// LincService contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLincService] method instead.
type LincService struct {
	Options []option.RequestOption
	V1      *LincV1Service
}

// NewLincService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLincService(opts ...option.RequestOption) (r *LincService) {
	r = &LincService{}
	r.Options = opts
	r.V1 = NewLincV1Service(opts...)
	return
}
