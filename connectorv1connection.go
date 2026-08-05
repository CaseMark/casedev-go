// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
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
// ConnectorV1ConnectionService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1ConnectionService] method instead.
type ConnectorV1ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectorV1ConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorV1ConnectionService(opts ...option.RequestOption) (r *ConnectorV1ConnectionService) {
	r = &ConnectorV1ConnectionService{}
	r.Options = opts
	return
}

// Create a pending provider connection and return a one-time connect_url for the
// hosted OAuth flow. The user completes provider consent at connect_url and is
// redirected to return_url with ?connection_id=.
func (r *ConnectorV1ConnectionService) New(ctx context.Context, body ConnectorV1ConnectionNewParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve one provider connection, including account identity and health.
func (r *ConnectorV1ConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// List provider connections for the organization, with health status.
func (r *ConnectorV1ConnectionService) List(ctx context.Context, query ConnectorV1ConnectionListParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Unlink a provider account: revoke tokens at the provider and delete them.
// purge=true additionally deletes the vault documents its import links brought in.
func (r *ConnectorV1ConnectionService) Delete(ctx context.Context, id string, body ConnectorV1ConnectionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Browse the provider one level at a time. Without a site, container, or parent,
// returns top-level resources. Pass the stable browse_ref fields returned by one
// response to navigate into the next level. Returns 403
// provider_scope_insufficient when the connection scope cannot browse server-side.
func (r *ConnectorV1ConnectionService) Browse(ctx context.Context, id string, query ConnectorV1ConnectionBrowseParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionBrowseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/browse", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConnectorV1ConnectionNewResponse struct {
	ConnectURL   string                               `json:"connect_url"`
	ConnectionID string                               `json:"connection_id"`
	ExpiresAt    string                               `json:"expires_at"`
	JSON         connectorV1ConnectionNewResponseJSON `json:"-"`
}

// connectorV1ConnectionNewResponseJSON contains the JSON metadata for the struct
// [ConnectorV1ConnectionNewResponse]
type connectorV1ConnectionNewResponseJSON struct {
	ConnectURL   apijson.Field
	ConnectionID apijson.Field
	ExpiresAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorV1ConnectionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponse struct {
	Capabilities ConnectorV1ConnectionListResponseCapabilities `json:"capabilities"`
	Connections  []interface{}                                 `json:"connections"`
	Cursor       string                                        `json:"cursor" api:"nullable"`
	JSON         connectorV1ConnectionListResponseJSON         `json:"-"`
}

// connectorV1ConnectionListResponseJSON contains the JSON metadata for the struct
// [ConnectorV1ConnectionListResponse]
type connectorV1ConnectionListResponseJSON struct {
	Capabilities apijson.Field
	Connections  apijson.Field
	Cursor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponseCapabilities struct {
	GoogleDriveFolderMirroring bool                                              `json:"google_drive_folder_mirroring"`
	JSON                       connectorV1ConnectionListResponseCapabilitiesJSON `json:"-"`
}

// connectorV1ConnectionListResponseCapabilitiesJSON contains the JSON metadata for
// the struct [ConnectorV1ConnectionListResponseCapabilities]
type connectorV1ConnectionListResponseCapabilitiesJSON struct {
	GoogleDriveFolderMirroring apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponseCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponseCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionBrowseResponse struct {
	Cursor string                                    `json:"cursor" api:"nullable"`
	Items  []ConnectorV1ConnectionBrowseResponseItem `json:"items"`
	JSON   connectorV1ConnectionBrowseResponseJSON   `json:"-"`
}

// connectorV1ConnectionBrowseResponseJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionBrowseResponse]
type connectorV1ConnectionBrowseResponseJSON struct {
	Cursor      apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1ConnectionBrowseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionBrowseResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionBrowseResponseItem struct {
	ID          string                                       `json:"id"`
	BrowseRef   interface{}                                  `json:"browse_ref" api:"nullable"`
	ContainerID string                                       `json:"container_id" api:"nullable"`
	Kind        ConnectorV1ConnectionBrowseResponseItemsKind `json:"kind"`
	MimeType    string                                       `json:"mime_type" api:"nullable"`
	ModifiedAt  string                                       `json:"modified_at" api:"nullable"`
	Name        string                                       `json:"name"`
	ParentIDs   []string                                     `json:"parent_ids"`
	Path        string                                       `json:"path" api:"nullable"`
	SizeBytes   int64                                        `json:"size_bytes" api:"nullable"`
	JSON        connectorV1ConnectionBrowseResponseItemJSON  `json:"-"`
}

// connectorV1ConnectionBrowseResponseItemJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionBrowseResponseItem]
type connectorV1ConnectionBrowseResponseItemJSON struct {
	ID          apijson.Field
	BrowseRef   apijson.Field
	ContainerID apijson.Field
	Kind        apijson.Field
	MimeType    apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	ParentIDs   apijson.Field
	Path        apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1ConnectionBrowseResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionBrowseResponseItemJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionBrowseResponseItemsKind string

const (
	ConnectorV1ConnectionBrowseResponseItemsKindMyDrive         ConnectorV1ConnectionBrowseResponseItemsKind = "my_drive"
	ConnectorV1ConnectionBrowseResponseItemsKindSharedDrive     ConnectorV1ConnectionBrowseResponseItemsKind = "shared_drive"
	ConnectorV1ConnectionBrowseResponseItemsKindMatter          ConnectorV1ConnectionBrowseResponseItemsKind = "matter"
	ConnectorV1ConnectionBrowseResponseItemsKindSite            ConnectorV1ConnectionBrowseResponseItemsKind = "site"
	ConnectorV1ConnectionBrowseResponseItemsKindDocumentLibrary ConnectorV1ConnectionBrowseResponseItemsKind = "document_library"
	ConnectorV1ConnectionBrowseResponseItemsKindFolder          ConnectorV1ConnectionBrowseResponseItemsKind = "folder"
	ConnectorV1ConnectionBrowseResponseItemsKindFile            ConnectorV1ConnectionBrowseResponseItemsKind = "file"
)

func (r ConnectorV1ConnectionBrowseResponseItemsKind) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionBrowseResponseItemsKindMyDrive, ConnectorV1ConnectionBrowseResponseItemsKindSharedDrive, ConnectorV1ConnectionBrowseResponseItemsKindMatter, ConnectorV1ConnectionBrowseResponseItemsKindSite, ConnectorV1ConnectionBrowseResponseItemsKindDocumentLibrary, ConnectorV1ConnectionBrowseResponseItemsKindFolder, ConnectorV1ConnectionBrowseResponseItemsKindFile:
		return true
	}
	return false
}

type ConnectorV1ConnectionNewParams struct {
	Provider param.Field[ConnectorV1ConnectionNewParamsProvider] `json:"provider" api:"required"`
	// HTTPS URL the user is sent back to after consent.
	ReturnURL param.Field[string] `json:"return_url" api:"required"`
	// Provider-specific OAuth permission tier. Omit to use the provider's default.
	ScopeTier param.Field[ConnectorV1ConnectionNewParamsScopeTier] `json:"scope_tier"`
}

func (r ConnectorV1ConnectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1ConnectionNewParamsProvider string

const (
	ConnectorV1ConnectionNewParamsProviderClio      ConnectorV1ConnectionNewParamsProvider = "clio"
	ConnectorV1ConnectionNewParamsProviderGdrive    ConnectorV1ConnectionNewParamsProvider = "gdrive"
	ConnectorV1ConnectionNewParamsProviderMicrosoft ConnectorV1ConnectionNewParamsProvider = "microsoft"
)

func (r ConnectorV1ConnectionNewParamsProvider) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionNewParamsProviderClio, ConnectorV1ConnectionNewParamsProviderGdrive, ConnectorV1ConnectionNewParamsProviderMicrosoft:
		return true
	}
	return false
}

// Provider-specific OAuth permission tier. Omit to use the provider's default.
type ConnectorV1ConnectionNewParamsScopeTier string

const (
	ConnectorV1ConnectionNewParamsScopeTierClioUs        ConnectorV1ConnectionNewParamsScopeTier = "clio.us"
	ConnectorV1ConnectionNewParamsScopeTierDrive         ConnectorV1ConnectionNewParamsScopeTier = "drive"
	ConnectorV1ConnectionNewParamsScopeTierMicrosoftRead ConnectorV1ConnectionNewParamsScopeTier = "microsoft.read"
)

func (r ConnectorV1ConnectionNewParamsScopeTier) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionNewParamsScopeTierClioUs, ConnectorV1ConnectionNewParamsScopeTierDrive, ConnectorV1ConnectionNewParamsScopeTierMicrosoftRead:
		return true
	}
	return false
}

type ConnectorV1ConnectionListParams struct {
	Provider param.Field[string]                                `query:"provider"`
	Status   param.Field[ConnectorV1ConnectionListParamsStatus] `query:"status"`
}

// URLQuery serializes [ConnectorV1ConnectionListParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1ConnectionListParamsStatus string

const (
	ConnectorV1ConnectionListParamsStatusPending        ConnectorV1ConnectionListParamsStatus = "pending"
	ConnectorV1ConnectionListParamsStatusHealthy        ConnectorV1ConnectionListParamsStatus = "healthy"
	ConnectorV1ConnectionListParamsStatusReauthRequired ConnectorV1ConnectionListParamsStatus = "reauth_required"
	ConnectorV1ConnectionListParamsStatusRevoked        ConnectorV1ConnectionListParamsStatus = "revoked"
	ConnectorV1ConnectionListParamsStatusThrottled      ConnectorV1ConnectionListParamsStatus = "throttled"
)

func (r ConnectorV1ConnectionListParamsStatus) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionListParamsStatusPending, ConnectorV1ConnectionListParamsStatusHealthy, ConnectorV1ConnectionListParamsStatusReauthRequired, ConnectorV1ConnectionListParamsStatusRevoked, ConnectorV1ConnectionListParamsStatusThrottled:
		return true
	}
	return false
}

type ConnectorV1ConnectionDeleteParams struct {
	Purge param.Field[bool] `query:"purge"`
}

// URLQuery serializes [ConnectorV1ConnectionDeleteParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1ConnectionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1ConnectionBrowseParams struct {
	// Container id to list, or the container containing parent
	Container param.Field[string] `query:"container"`
	Cursor    param.Field[string] `query:"cursor"`
	PageSize  param.Field[int64]  `query:"page_size"`
	// Folder id to list
	Parent param.Field[string] `query:"parent"`
	// Optional provider-supported search text
	Query param.Field[string] `query:"query"`
	// Site id to list
	Site param.Field[string] `query:"site"`
}

// URLQuery serializes [ConnectorV1ConnectionBrowseParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1ConnectionBrowseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
