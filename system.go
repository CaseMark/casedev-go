// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// SystemService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemService] method instead.
type SystemService struct {
	Options []option.RequestOption
}

// NewSystemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSystemService(opts ...option.RequestOption) (r *SystemService) {
	r = &SystemService{}
	r.Options = opts
	return
}

// Returns the public Case.dev services catalog derived from
// docs.case.dev/services. This endpoint is unauthenticated and intended for
// discovery surfaces such as the case.dev homepage.
func (r *SystemService) ListServices(ctx context.Context, opts ...option.RequestOption) (res *SystemListServicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SystemListServicesResponse struct {
	Services []SystemListServicesResponseService `json:"services" api:"required"`
	JSON     systemListServicesResponseJSON      `json:"-"`
}

// systemListServicesResponseJSON contains the JSON metadata for the struct
// [SystemListServicesResponse]
type systemListServicesResponseJSON struct {
	Services    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemListServicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemListServicesResponseJSON) RawJSON() string {
	return r.raw
}

type SystemListServicesResponseService struct {
	ID          string                                `json:"id" api:"required"`
	Description string                                `json:"description" api:"required"`
	Href        string                                `json:"href" api:"required"`
	Icon        string                                `json:"icon" api:"required"`
	Name        string                                `json:"name" api:"required"`
	Order       int64                                 `json:"order" api:"required"`
	JSON        systemListServicesResponseServiceJSON `json:"-"`
}

// systemListServicesResponseServiceJSON contains the JSON metadata for the struct
// [SystemListServicesResponseService]
type systemListServicesResponseServiceJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Href        apijson.Field
	Icon        apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SystemListServicesResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r systemListServicesResponseServiceJSON) RawJSON() string {
	return r.raw
}
