// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// TranslateService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTranslateService] method instead.
type TranslateService struct {
	Options []option.RequestOption
	V1      *TranslateV1Service
}

// NewTranslateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTranslateService(opts ...option.RequestOption) (r *TranslateService) {
	r = &TranslateService{}
	r.Options = opts
	r.V1 = NewTranslateV1Service(opts...)
	return
}
