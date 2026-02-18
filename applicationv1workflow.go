// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// ApplicationV1WorkflowService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationV1WorkflowService] method instead.
type ApplicationV1WorkflowService struct {
	Options []option.RequestOption
}

// NewApplicationV1WorkflowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationV1WorkflowService(opts ...option.RequestOption) (r *ApplicationV1WorkflowService) {
	r = &ApplicationV1WorkflowService{}
	r.Options = opts
	return
}

// Get current deployment workflow status and accumulated events
func (r *ApplicationV1WorkflowService) GetStatus(ctx context.Context, id string, query ApplicationV1WorkflowGetStatusParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/workflows/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ApplicationV1WorkflowGetStatusParams struct {
	// Project ID (for authorization)
	ProjectID param.Field[string] `query:"projectId,required"`
}

// URLQuery serializes [ApplicationV1WorkflowGetStatusParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1WorkflowGetStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
