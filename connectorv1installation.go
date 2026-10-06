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

// Import and export between provider folders and vaults
//
// ConnectorV1InstallationService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1InstallationService] method instead.
type ConnectorV1InstallationService struct {
	Options []option.RequestOption
	// Import and export between provider folders and vaults
	Tokens *ConnectorV1InstallationTokenService
	// Import and export between provider folders and vaults
	Vaults *ConnectorV1InstallationVaultService
}

// NewConnectorV1InstallationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorV1InstallationService(opts ...option.RequestOption) (r *ConnectorV1InstallationService) {
	r = &ConnectorV1InstallationService{}
	r.Options = opts
	r.Tokens = NewConnectorV1InstallationTokenService(opts...)
	r.Vaults = NewConnectorV1InstallationVaultService(opts...)
	return
}

// List application installations (tenants) in this organization. Returns at most
// `limit` installations (default 200, maximum 200). When `pagination.has_more` is
// true, replay `pagination.next_cursor` as `?cursor=` to fetch the following page.
// Cursors are opaque and are only valid for the exact filter set and caller scope
// they were issued under.
func (r *ConnectorV1InstallationService) List(ctx context.Context, query ConnectorV1InstallationListParams, opts ...option.RequestOption) (res *ConnectorV1InstallationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/installations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

type ConnectorV1InstallationListResponse struct {
	Installations []interface{}                                 `json:"installations"`
	Pagination    ConnectorV1InstallationListResponsePagination `json:"pagination"`
	JSON          connectorV1InstallationListResponseJSON       `json:"-"`
}

// connectorV1InstallationListResponseJSON contains the JSON metadata for the
// struct [ConnectorV1InstallationListResponse]
type connectorV1InstallationListResponseJSON struct {
	Installations apijson.Field
	Pagination    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectorV1InstallationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1InstallationListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1InstallationListResponsePagination struct {
	HasMore    bool                                              `json:"has_more"`
	Limit      int64                                             `json:"limit"`
	NextCursor string                                            `json:"next_cursor" api:"nullable"`
	JSON       connectorV1InstallationListResponsePaginationJSON `json:"-"`
}

// connectorV1InstallationListResponsePaginationJSON contains the JSON metadata for
// the struct [ConnectorV1InstallationListResponsePagination]
type connectorV1InstallationListResponsePaginationJSON struct {
	HasMore     apijson.Field
	Limit       apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1InstallationListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1InstallationListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1InstallationListParams struct {
	Application param.Field[string] `query:"application"`
	// Opaque continuation cursor from `pagination.next_cursor` of the previous page.
	// Must be replayed with the same filters and scope that produced it.
	Cursor           param.Field[string] `query:"cursor"`
	ExternalTenantID param.Field[string] `query:"external_tenant_id"`
	// Installations per page (1-200). Defaults to 200.
	Limit param.Field[int64] `query:"limit"`
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
