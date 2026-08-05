// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Import and export between provider folders (Google Drive) and vaults
//
// ConnectorV1InstallationService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1InstallationService] method instead.
type ConnectorV1InstallationService struct {
	Options []option.RequestOption
	// Import and export between provider folders (Google Drive) and vaults
	Vaults *ConnectorV1InstallationVaultService
}

// NewConnectorV1InstallationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorV1InstallationService(opts ...option.RequestOption) (r *ConnectorV1InstallationService) {
	r = &ConnectorV1InstallationService{}
	r.Options = opts
	r.Vaults = NewConnectorV1InstallationVaultService(opts...)
	return
}

// List application installations (tenants) in this organization.
func (r *ConnectorV1InstallationService) List(ctx context.Context, query ConnectorV1InstallationListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "connectors/v1/installations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Idempotently create (or return) the installation for (application,
// external_tenant_id) in this organization. Send the returned installation id as
// X-Case-Installation-Id on connector requests to scope them to this tenant.
func (r *ConnectorV1InstallationService) Ensure(ctx context.Context, body ConnectorV1InstallationEnsureParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "connectors/v1/installations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type ConnectorV1InstallationListParams struct {
	Application      param.Field[string] `query:"application"`
	ExternalTenantID param.Field[string] `query:"external_tenant_id"`
}

// URLQuery serializes [ConnectorV1InstallationListParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1InstallationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1InstallationEnsureParams struct {
	// Consuming application key (e.g. "p3").
	Application param.Field[string] `json:"application" api:"required"`
	// The application's own tenant identifier (e.g. a P3 organization id).
	ExternalTenantID param.Field[string] `json:"external_tenant_id" api:"required"`
}

func (r ConnectorV1InstallationEnsureParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
