// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Import and export between provider folders and vaults
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
func (r *ConnectorV1ConnectionService) New(ctx context.Context, params ConnectorV1ConnectionNewParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionNewResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve one provider connection, including account identity and health.
func (r *ConnectorV1ConnectionService) Get(ctx context.Context, id string, query ConnectorV1ConnectionGetParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionGetResponse, err error) {
	if query.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", query.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List provider connections for the organization, with health status. Returns at
// most `limit` connections (default 200, maximum 200). When `pagination.has_more`
// is true, replay `pagination.next_cursor` as `?cursor=` to fetch the following
// page. Cursors are opaque and are only valid for the exact filter set and
// installation/subject scope they were issued under.
func (r *ConnectorV1ConnectionService) List(ctx context.Context, params ConnectorV1ConnectionListParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionListResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Unlink a provider account: revoke tokens at the provider and delete them.
// purge=true additionally deletes the vault documents its import links brought in.
func (r *ConnectorV1ConnectionService) Delete(ctx context.Context, id string, params ConnectorV1ConnectionDeleteParams, opts ...option.RequestOption) (err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Browse the provider one level at a time. Without a site, container, or parent,
// returns top-level resources. Pass the stable browse_ref fields returned by one
// response to navigate into the next level. Returns 403
// provider_scope_insufficient when the connection scope cannot browse server-side.
// Clio browsing shares request capacity with background runs; throttled responses
// include Retry-After when a retry deadline is known.
func (r *ConnectorV1ConnectionService) Browse(ctx context.Context, id string, params ConnectorV1ConnectionBrowseParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionBrowseResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/browse", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Create one folder at the provider directly under a browse location, then export
// into it with POST /connectors/v1/transfer. Name the parent with the same site,
// container, and parent fields browse accepts (for Clio and Smokeball, a matter
// root is the browse_ref the matter returned). A folder of the same name already
// under the parent is returned with 200 instead of a duplicate; a new folder
// returns 201. Smokeball creates folders asynchronously: when the new folder is
// not listed yet the response is 202 with code folder_creation_pending and no
// folder. Repeat the request after a short wait to get it with 200; a repeat sent
// while it is still not listed asks Smokeball again and can leave two folders with
// that name, so space repeats out. Returns 403 provider_scope_insufficient when
// the connection was granted read access only; reconnect with write access.
func (r *ConnectorV1ConnectionService) NewFolder(ctx context.Context, id string, params ConnectorV1ConnectionNewFolderParams, opts ...option.RequestOption) (err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/folders", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Create a seven-day, replaceable IT approval link for a pending Microsoft
// connection. Approval does not connect an account. Only delegated document read
// scopes are requested.
func (r *ConnectorV1ConnectionService) NewMicrosoftAdminApproval(ctx context.Context, id string, body ConnectorV1ConnectionNewMicrosoftAdminApprovalParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionNewMicrosoftAdminApprovalResponse, err error) {
	if body.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", body.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/microsoft-admin-approval", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// List every file under one folder, at any depth, as a flat paginated list, to
// preview what importing that folder would bring in. Pass the browse_ref fields of
// a folder or container. Each file carries its path relative to that folder. A
// page returns as soon as it holds files, so it may be shorter than page_size
// while cursor is still set; keep requesting until cursor is null. The listing
// stops after 10,000 files and sets truncated. Returns 422 contents_not_supported
// when the provider cannot list folder contents, and shares the provider request
// budget with browsing and background runs.
func (r *ConnectorV1ConnectionService) ListContents(ctx context.Context, id string, params ConnectorV1ConnectionListContentsParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionListContentsResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/contents", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Resolve a pasted provider folder link to that folder's browse item, including
// its browse_ref, name, and path when the provider reports it cheaply. Supports
// Google Drive folder and open?id= links, Box folder and shared links, Dropbox
// shared links, and OneDrive/SharePoint sharing links. Returns 400
// link_not_recognized for links of another provider or form, 422 link_not_folder
// for file links, 422 link_not_in_account when a Dropbox or Box shared folder is
// not in the connected account, 404 link_not_accessible when the account cannot
// see it, and 422 link_resolve_not_supported for providers without folder links.
// Shares the provider request budget with browsing.
func (r *ConnectorV1ConnectionService) ResolveLink(ctx context.Context, id string, params ConnectorV1ConnectionResolveLinkParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionResolveLinkResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/resolve", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Search the provider's files by name, and by content where the provider supports
// it; with kind=folder, search folders by name. Without a scope, searches
// everything the connection can see; pass the browse_ref fields of a folder or
// container to search inside it. Folder results carry a browse_ref to open them
// with browse. Each response states whether its folder scope covers the whole
// subtree. Returns 422 search_not_supported when the provider has no search and
// search_kind_not_supported when it cannot search folders, and shares the provider
// request budget with browsing and background runs.
func (r *ConnectorV1ConnectionService) Search(ctx context.Context, id string, params ConnectorV1ConnectionSearchParams, opts ...option.RequestOption) (res *ConnectorV1ConnectionSearchResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/connections/%s/search", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Enable or disable new runs and scheduled syncs for one provider across the
// authenticated organization or installation. This organization-wide operation
// requires explicit confirmation. Existing credentials, links, and imported files
// are preserved; active runs are not interrupted.
func (r *ConnectorV1ConnectionService) UpdateAll(ctx context.Context, body ConnectorV1ConnectionUpdateAllParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "connectors/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
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

type ConnectorV1ConnectionGetResponse struct {
	// Authenticated organization owning this connection; server-side approval binding
	// uses this identifier.
	OrganizationID string                               `json:"organization_id"`
	ExtraFields    map[string]interface{}               `json:"-" api:"extrafields"`
	JSON           connectorV1ConnectionGetResponseJSON `json:"-"`
}

// connectorV1ConnectionGetResponseJSON contains the JSON metadata for the struct
// [ConnectorV1ConnectionGetResponse]
type connectorV1ConnectionGetResponseJSON struct {
	OrganizationID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectorV1ConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponse struct {
	Capabilities ConnectorV1ConnectionListResponseCapabilities `json:"capabilities"`
	Connections  []ConnectorV1ConnectionListResponseConnection `json:"connections"`
	// Mirror of `pagination.next_cursor`. Prefer `pagination`.
	//
	// Deprecated: deprecated
	Cursor     string                                      `json:"cursor" api:"nullable"`
	Pagination ConnectorV1ConnectionListResponsePagination `json:"pagination"`
	// Available import providers and their adapter-declared capabilities.
	Providers []ConnectorV1ConnectionListResponseProvider `json:"providers"`
	JSON      connectorV1ConnectionListResponseJSON       `json:"-"`
}

// connectorV1ConnectionListResponseJSON contains the JSON metadata for the struct
// [ConnectorV1ConnectionListResponse]
type connectorV1ConnectionListResponseJSON struct {
	Capabilities apijson.Field
	Connections  apijson.Field
	Cursor       apijson.Field
	Pagination   apijson.Field
	Providers    apijson.Field
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

type ConnectorV1ConnectionListResponseConnection struct {
	// Authenticated organization owning this connection; server-side approval binding
	// uses this identifier.
	OrganizationID string                                          `json:"organization_id"`
	ExtraFields    map[string]interface{}                          `json:"-" api:"extrafields"`
	JSON           connectorV1ConnectionListResponseConnectionJSON `json:"-"`
}

// connectorV1ConnectionListResponseConnectionJSON contains the JSON metadata for
// the struct [ConnectorV1ConnectionListResponseConnection]
type connectorV1ConnectionListResponseConnectionJSON struct {
	OrganizationID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponseConnectionJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponsePagination struct {
	HasMore    bool                                            `json:"has_more"`
	Limit      int64                                           `json:"limit"`
	NextCursor string                                          `json:"next_cursor" api:"nullable"`
	JSON       connectorV1ConnectionListResponsePaginationJSON `json:"-"`
}

// connectorV1ConnectionListResponsePaginationJSON contains the JSON metadata for
// the struct [ConnectorV1ConnectionListResponsePagination]
type connectorV1ConnectionListResponsePaginationJSON struct {
	HasMore     apijson.Field
	Limit       apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponseProvider struct {
	ID string `json:"id"`
	// Adapter-declared consent choices, filtered by deployment and organization
	// policy.
	AuthorizationOptions []ConnectorV1ConnectionListResponseProvidersAuthorizationOption `json:"authorization_options"`
	// Exact enabled tiers accepted by the adapter browse policy.
	BrowseScopeTiers []string `json:"browse_scope_tiers"`
	ResourceTypes    []string `json:"resource_types"`
	ScopeTier        string   `json:"scope_tier"`
	// Present when GET /connections/{id}/search works for this provider.
	Search ConnectorV1ConnectionListResponseProvidersSearch `json:"search" api:"nullable"`
	// True when POST /connections/{id}/folders can create folders for this provider.
	SupportsCreateFolder bool                                          `json:"supports_create_folder"`
	SupportsExport       bool                                          `json:"supports_export"`
	JSON                 connectorV1ConnectionListResponseProviderJSON `json:"-"`
}

// connectorV1ConnectionListResponseProviderJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionListResponseProvider]
type connectorV1ConnectionListResponseProviderJSON struct {
	ID                   apijson.Field
	AuthorizationOptions apijson.Field
	BrowseScopeTiers     apijson.Field
	ResourceTypes        apijson.Field
	ScopeTier            apijson.Field
	Search               apijson.Field
	SupportsCreateFolder apijson.Field
	SupportsExport       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponseProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponseProviderJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponseProvidersAuthorizationOption struct {
	Label     string `json:"label"`
	ScopeTier string `json:"scope_tier"`
	// Explicit write-authorization upgrade for this account choice. Null means
	// unavailable by policy; omitted when no upgrade is required.
	WriteScopeTier string                                                            `json:"write_scope_tier" api:"nullable"`
	JSON           connectorV1ConnectionListResponseProvidersAuthorizationOptionJSON `json:"-"`
}

// connectorV1ConnectionListResponseProvidersAuthorizationOptionJSON contains the
// JSON metadata for the struct
// [ConnectorV1ConnectionListResponseProvidersAuthorizationOption]
type connectorV1ConnectionListResponseProvidersAuthorizationOptionJSON struct {
	Label          apijson.Field
	ScopeTier      apijson.Field
	WriteScopeTier apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponseProvidersAuthorizationOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponseProvidersAuthorizationOptionJSON) RawJSON() string {
	return r.raw
}

// Present when GET /connections/{id}/search works for this provider.
type ConnectorV1ConnectionListResponseProvidersSearch struct {
	FolderScopeRecursive bool `json:"folder_scope_recursive"`
	FullText             bool `json:"full_text"`
	// Values the search kind parameter accepts.
	Kinds       []ConnectorV1ConnectionListResponseProvidersSearchKind  `json:"kinds"`
	MaxPageSize int64                                                   `json:"max_page_size"`
	Scopes      []ConnectorV1ConnectionListResponseProvidersSearchScope `json:"scopes"`
	JSON        connectorV1ConnectionListResponseProvidersSearchJSON    `json:"-"`
}

// connectorV1ConnectionListResponseProvidersSearchJSON contains the JSON metadata
// for the struct [ConnectorV1ConnectionListResponseProvidersSearch]
type connectorV1ConnectionListResponseProvidersSearchJSON struct {
	FolderScopeRecursive apijson.Field
	FullText             apijson.Field
	Kinds                apijson.Field
	MaxPageSize          apijson.Field
	Scopes               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListResponseProvidersSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListResponseProvidersSearchJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListResponseProvidersSearchKind string

const (
	ConnectorV1ConnectionListResponseProvidersSearchKindFile   ConnectorV1ConnectionListResponseProvidersSearchKind = "file"
	ConnectorV1ConnectionListResponseProvidersSearchKindFolder ConnectorV1ConnectionListResponseProvidersSearchKind = "folder"
)

func (r ConnectorV1ConnectionListResponseProvidersSearchKind) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionListResponseProvidersSearchKindFile, ConnectorV1ConnectionListResponseProvidersSearchKindFolder:
		return true
	}
	return false
}

type ConnectorV1ConnectionListResponseProvidersSearchScope string

const (
	ConnectorV1ConnectionListResponseProvidersSearchScopeAccount ConnectorV1ConnectionListResponseProvidersSearchScope = "account"
	ConnectorV1ConnectionListResponseProvidersSearchScopeFolder  ConnectorV1ConnectionListResponseProvidersSearchScope = "folder"
)

func (r ConnectorV1ConnectionListResponseProvidersSearchScope) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionListResponseProvidersSearchScopeAccount, ConnectorV1ConnectionListResponseProvidersSearchScopeFolder:
		return true
	}
	return false
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
	ID          string      `json:"id"`
	BrowseRef   interface{} `json:"browse_ref" api:"nullable"`
	ContainerID string      `json:"container_id" api:"nullable"`
	Description string      `json:"description" api:"nullable"`
	// Importable container root, when different from its browse reference.
	ImportRef  interface{}                                  `json:"import_ref" api:"nullable"`
	Kind       ConnectorV1ConnectionBrowseResponseItemsKind `json:"kind"`
	MimeType   string                                       `json:"mime_type" api:"nullable"`
	ModifiedAt string                                       `json:"modified_at" api:"nullable"`
	Name       string                                       `json:"name"`
	// Signed reference to this file for POST /connectors/v1/fetch. Null for folders
	// and providers without single-file fetch. Pass it back exactly as received.
	OpenRef   string                                      `json:"open_ref" api:"nullable"`
	ParentIDs []string                                    `json:"parent_ids"`
	Path      string                                      `json:"path" api:"nullable"`
	SizeBytes int64                                       `json:"size_bytes" api:"nullable"`
	JSON      connectorV1ConnectionBrowseResponseItemJSON `json:"-"`
}

// connectorV1ConnectionBrowseResponseItemJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionBrowseResponseItem]
type connectorV1ConnectionBrowseResponseItemJSON struct {
	ID          apijson.Field
	BrowseRef   apijson.Field
	ContainerID apijson.Field
	Description apijson.Field
	ImportRef   apijson.Field
	Kind        apijson.Field
	MimeType    apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	OpenRef     apijson.Field
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

type ConnectorV1ConnectionNewMicrosoftAdminApprovalResponse struct {
	ApprovalURL string                                                     `json:"approval_url" api:"required" format:"uri"`
	ExpiresAt   time.Time                                                  `json:"expires_at" api:"required" format:"date-time"`
	JSON        connectorV1ConnectionNewMicrosoftAdminApprovalResponseJSON `json:"-"`
}

// connectorV1ConnectionNewMicrosoftAdminApprovalResponseJSON contains the JSON
// metadata for the struct [ConnectorV1ConnectionNewMicrosoftAdminApprovalResponse]
type connectorV1ConnectionNewMicrosoftAdminApprovalResponseJSON struct {
	ApprovalURL apijson.Field
	ExpiresAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1ConnectionNewMicrosoftAdminApprovalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionNewMicrosoftAdminApprovalResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListContentsResponse struct {
	Cursor string                                          `json:"cursor" api:"nullable"`
	Items  []ConnectorV1ConnectionListContentsResponseItem `json:"items"`
	// The listed folder, as a browse_ref
	ScopeRef interface{} `json:"scope_ref"`
	// True when the listing stopped before every file was listed, at the 10,000-file
	// limit or because the folder tree was too wide to page through
	Truncated bool                                          `json:"truncated"`
	JSON      connectorV1ConnectionListContentsResponseJSON `json:"-"`
}

// connectorV1ConnectionListContentsResponseJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionListContentsResponse]
type connectorV1ConnectionListContentsResponseJSON struct {
	Cursor      apijson.Field
	Items       apijson.Field
	ScopeRef    apijson.Field
	Truncated   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListContentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListContentsResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListContentsResponseItem struct {
	ID          string                                             `json:"id"`
	ContainerID string                                             `json:"container_id" api:"nullable"`
	Kind        ConnectorV1ConnectionListContentsResponseItemsKind `json:"kind"`
	MimeType    string                                             `json:"mime_type" api:"nullable"`
	ModifiedAt  string                                             `json:"modified_at" api:"nullable"`
	Name        string                                             `json:"name"`
	// Signed reference to this file for POST /connectors/v1/fetch. Null for providers
	// without single-file fetch. Pass it back exactly as received.
	OpenRef   string   `json:"open_ref" api:"nullable"`
	ParentIDs []string `json:"parent_ids"`
	// Path relative to the listed folder, ending in the file's name
	Path      string `json:"path"`
	SizeBytes int64  `json:"size_bytes" api:"nullable"`
	// Set when an import would skip the file (for example a Box note)
	UnsupportedReason string                                            `json:"unsupported_reason" api:"nullable"`
	JSON              connectorV1ConnectionListContentsResponseItemJSON `json:"-"`
}

// connectorV1ConnectionListContentsResponseItemJSON contains the JSON metadata for
// the struct [ConnectorV1ConnectionListContentsResponseItem]
type connectorV1ConnectionListContentsResponseItemJSON struct {
	ID                apijson.Field
	ContainerID       apijson.Field
	Kind              apijson.Field
	MimeType          apijson.Field
	ModifiedAt        apijson.Field
	Name              apijson.Field
	OpenRef           apijson.Field
	ParentIDs         apijson.Field
	Path              apijson.Field
	SizeBytes         apijson.Field
	UnsupportedReason apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConnectorV1ConnectionListContentsResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionListContentsResponseItemJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionListContentsResponseItemsKind string

const (
	ConnectorV1ConnectionListContentsResponseItemsKindFile ConnectorV1ConnectionListContentsResponseItemsKind = "file"
)

func (r ConnectorV1ConnectionListContentsResponseItemsKind) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionListContentsResponseItemsKindFile:
		return true
	}
	return false
}

type ConnectorV1ConnectionResolveLinkResponse struct {
	ID string `json:"id"`
	// Pass these fields to browse to open the folder
	BrowseRef   interface{}                                  `json:"browse_ref"`
	ContainerID string                                       `json:"container_id" api:"nullable"`
	Description string                                       `json:"description" api:"nullable"`
	ImportRef   interface{}                                  `json:"import_ref" api:"nullable"`
	Kind        ConnectorV1ConnectionResolveLinkResponseKind `json:"kind"`
	MimeType    string                                       `json:"mime_type" api:"nullable"`
	ModifiedAt  string                                       `json:"modified_at" api:"nullable"`
	Name        string                                       `json:"name"`
	// Always null for folders
	OpenRef   string   `json:"open_ref" api:"nullable"`
	ParentIDs []string `json:"parent_ids"`
	// Display path in the account, when the provider reports it
	Path      string                                       `json:"path" api:"nullable"`
	SizeBytes int64                                        `json:"size_bytes" api:"nullable"`
	JSON      connectorV1ConnectionResolveLinkResponseJSON `json:"-"`
}

// connectorV1ConnectionResolveLinkResponseJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionResolveLinkResponse]
type connectorV1ConnectionResolveLinkResponseJSON struct {
	ID          apijson.Field
	BrowseRef   apijson.Field
	ContainerID apijson.Field
	Description apijson.Field
	ImportRef   apijson.Field
	Kind        apijson.Field
	MimeType    apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	OpenRef     apijson.Field
	ParentIDs   apijson.Field
	Path        apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1ConnectionResolveLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionResolveLinkResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionResolveLinkResponseKind string

const (
	ConnectorV1ConnectionResolveLinkResponseKindFolder ConnectorV1ConnectionResolveLinkResponseKind = "folder"
)

func (r ConnectorV1ConnectionResolveLinkResponseKind) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionResolveLinkResponseKindFolder:
		return true
	}
	return false
}

type ConnectorV1ConnectionSearchResponse struct {
	Cursor string `json:"cursor" api:"nullable"`
	// False when a folder scope covers its direct children only
	FolderScopeRecursive bool `json:"folder_scope_recursive"`
	// Whether the provider matches file contents, not only names. Always false for
	// kind=folder.
	FullText bool                                      `json:"full_text"`
	Items    []ConnectorV1ConnectionSearchResponseItem `json:"items"`
	Scope    ConnectorV1ConnectionSearchResponseScope  `json:"scope"`
	ScopeRef interface{}                               `json:"scope_ref" api:"nullable"`
	JSON     connectorV1ConnectionSearchResponseJSON   `json:"-"`
}

// connectorV1ConnectionSearchResponseJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionSearchResponse]
type connectorV1ConnectionSearchResponseJSON struct {
	Cursor               apijson.Field
	FolderScopeRecursive apijson.Field
	FullText             apijson.Field
	Items                apijson.Field
	Scope                apijson.Field
	ScopeRef             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConnectorV1ConnectionSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionSearchResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionSearchResponseItem struct {
	ID string `json:"id"`
	// Reference for GET /connectors/v1/connections/{id}/browse that opens this folder.
	// Null for files.
	BrowseRef   interface{}                                  `json:"browse_ref" api:"nullable"`
	ContainerID string                                       `json:"container_id" api:"nullable"`
	Kind        ConnectorV1ConnectionSearchResponseItemsKind `json:"kind"`
	MimeType    string                                       `json:"mime_type" api:"nullable"`
	ModifiedAt  string                                       `json:"modified_at" api:"nullable"`
	Name        string                                       `json:"name"`
	// Signed reference to this file for POST /connectors/v1/fetch. Null for providers
	// without single-file fetch. Pass it back exactly as received.
	OpenRef   string   `json:"open_ref" api:"nullable"`
	ParentIDs []string `json:"parent_ids"`
	Path      string   `json:"path" api:"nullable"`
	SizeBytes int64    `json:"size_bytes" api:"nullable"`
	// Set when the file cannot be imported (for example a Box note)
	UnsupportedReason string                                      `json:"unsupported_reason" api:"nullable"`
	JSON              connectorV1ConnectionSearchResponseItemJSON `json:"-"`
}

// connectorV1ConnectionSearchResponseItemJSON contains the JSON metadata for the
// struct [ConnectorV1ConnectionSearchResponseItem]
type connectorV1ConnectionSearchResponseItemJSON struct {
	ID                apijson.Field
	BrowseRef         apijson.Field
	ContainerID       apijson.Field
	Kind              apijson.Field
	MimeType          apijson.Field
	ModifiedAt        apijson.Field
	Name              apijson.Field
	OpenRef           apijson.Field
	ParentIDs         apijson.Field
	Path              apijson.Field
	SizeBytes         apijson.Field
	UnsupportedReason apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConnectorV1ConnectionSearchResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ConnectionSearchResponseItemJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1ConnectionSearchResponseItemsKind string

const (
	ConnectorV1ConnectionSearchResponseItemsKindFile   ConnectorV1ConnectionSearchResponseItemsKind = "file"
	ConnectorV1ConnectionSearchResponseItemsKindFolder ConnectorV1ConnectionSearchResponseItemsKind = "folder"
)

func (r ConnectorV1ConnectionSearchResponseItemsKind) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionSearchResponseItemsKindFile, ConnectorV1ConnectionSearchResponseItemsKindFolder:
		return true
	}
	return false
}

type ConnectorV1ConnectionSearchResponseScope string

const (
	ConnectorV1ConnectionSearchResponseScopeAccount ConnectorV1ConnectionSearchResponseScope = "account"
	ConnectorV1ConnectionSearchResponseScopeFolder  ConnectorV1ConnectionSearchResponseScope = "folder"
)

func (r ConnectorV1ConnectionSearchResponseScope) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionSearchResponseScopeAccount, ConnectorV1ConnectionSearchResponseScopeFolder:
		return true
	}
	return false
}

type ConnectorV1ConnectionNewParams struct {
	Provider param.Field[ConnectorV1ConnectionNewParamsProvider] `json:"provider" api:"required"`
	// HTTPS URL the user is sent back to after consent.
	ReturnURL param.Field[string] `json:"return_url" api:"required"`
	// Provider-specific OAuth permission tier. Omit to use the provider's default. Use
	// microsoft.account.read to let Microsoft choose a personal or work/school
	// account. Existing explicit microsoft.read and microsoft.personal.read tiers
	// remain supported. Microsoft write tiers require the separate write switch and
	// explicit authorization for bounded exports and paired sync.
	ScopeTier             param.Field[ConnectorV1ConnectionNewParamsScopeTier] `json:"scope_tier"`
	XCaseConnectorSubject param.Field[string]                                  `header:"x-case-connector-subject"`
}

func (r ConnectorV1ConnectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1ConnectionNewParamsProvider string

const (
	ConnectorV1ConnectionNewParamsProviderBox       ConnectorV1ConnectionNewParamsProvider = "box"
	ConnectorV1ConnectionNewParamsProviderClio      ConnectorV1ConnectionNewParamsProvider = "clio"
	ConnectorV1ConnectionNewParamsProviderDropbox   ConnectorV1ConnectionNewParamsProvider = "dropbox"
	ConnectorV1ConnectionNewParamsProviderGdrive    ConnectorV1ConnectionNewParamsProvider = "gdrive"
	ConnectorV1ConnectionNewParamsProviderMicrosoft ConnectorV1ConnectionNewParamsProvider = "microsoft"
	ConnectorV1ConnectionNewParamsProviderSharefile ConnectorV1ConnectionNewParamsProvider = "sharefile"
	ConnectorV1ConnectionNewParamsProviderSmokeball ConnectorV1ConnectionNewParamsProvider = "smokeball"
)

func (r ConnectorV1ConnectionNewParamsProvider) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionNewParamsProviderBox, ConnectorV1ConnectionNewParamsProviderClio, ConnectorV1ConnectionNewParamsProviderDropbox, ConnectorV1ConnectionNewParamsProviderGdrive, ConnectorV1ConnectionNewParamsProviderMicrosoft, ConnectorV1ConnectionNewParamsProviderSharefile, ConnectorV1ConnectionNewParamsProviderSmokeball:
		return true
	}
	return false
}

// Provider-specific OAuth permission tier. Omit to use the provider's default. Use
// microsoft.account.read to let Microsoft choose a personal or work/school
// account. Existing explicit microsoft.read and microsoft.personal.read tiers
// remain supported. Microsoft write tiers require the separate write switch and
// explicit authorization for bounded exports and paired sync.
type ConnectorV1ConnectionNewParamsScopeTier string

const (
	ConnectorV1ConnectionNewParamsScopeTierBoxReadwrite            ConnectorV1ConnectionNewParamsScopeTier = "box.readwrite"
	ConnectorV1ConnectionNewParamsScopeTierBoxReadwriteWebhooks    ConnectorV1ConnectionNewParamsScopeTier = "box.readwrite.webhooks"
	ConnectorV1ConnectionNewParamsScopeTierClioUs                  ConnectorV1ConnectionNewParamsScopeTier = "clio.us"
	ConnectorV1ConnectionNewParamsScopeTierDropboxReadwrite        ConnectorV1ConnectionNewParamsScopeTier = "dropbox.readwrite"
	ConnectorV1ConnectionNewParamsScopeTierDrive                   ConnectorV1ConnectionNewParamsScopeTier = "drive"
	ConnectorV1ConnectionNewParamsScopeTierMicrosoftAccountRead    ConnectorV1ConnectionNewParamsScopeTier = "microsoft.account.read"
	ConnectorV1ConnectionNewParamsScopeTierMicrosoftRead           ConnectorV1ConnectionNewParamsScopeTier = "microsoft.read"
	ConnectorV1ConnectionNewParamsScopeTierMicrosoftPersonalRead   ConnectorV1ConnectionNewParamsScopeTier = "microsoft.personal.read"
	ConnectorV1ConnectionNewParamsScopeTierMicrosoftWrite          ConnectorV1ConnectionNewParamsScopeTier = "microsoft.write"
	ConnectorV1ConnectionNewParamsScopeTierMicrosoftPersonalWrite  ConnectorV1ConnectionNewParamsScopeTier = "microsoft.personal.write"
	ConnectorV1ConnectionNewParamsScopeTierSharefile               ConnectorV1ConnectionNewParamsScopeTier = "sharefile"
	ConnectorV1ConnectionNewParamsScopeTierSmokeballUs             ConnectorV1ConnectionNewParamsScopeTier = "smokeball.us"
	ConnectorV1ConnectionNewParamsScopeTierSmokeballUsStaging      ConnectorV1ConnectionNewParamsScopeTier = "smokeball.us.staging"
	ConnectorV1ConnectionNewParamsScopeTierSmokeballUsStagingWrite ConnectorV1ConnectionNewParamsScopeTier = "smokeball.us.staging.write"
	ConnectorV1ConnectionNewParamsScopeTierSmokeballUsWrite        ConnectorV1ConnectionNewParamsScopeTier = "smokeball.us.write"
)

func (r ConnectorV1ConnectionNewParamsScopeTier) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionNewParamsScopeTierBoxReadwrite, ConnectorV1ConnectionNewParamsScopeTierBoxReadwriteWebhooks, ConnectorV1ConnectionNewParamsScopeTierClioUs, ConnectorV1ConnectionNewParamsScopeTierDropboxReadwrite, ConnectorV1ConnectionNewParamsScopeTierDrive, ConnectorV1ConnectionNewParamsScopeTierMicrosoftAccountRead, ConnectorV1ConnectionNewParamsScopeTierMicrosoftRead, ConnectorV1ConnectionNewParamsScopeTierMicrosoftPersonalRead, ConnectorV1ConnectionNewParamsScopeTierMicrosoftWrite, ConnectorV1ConnectionNewParamsScopeTierMicrosoftPersonalWrite, ConnectorV1ConnectionNewParamsScopeTierSharefile, ConnectorV1ConnectionNewParamsScopeTierSmokeballUs, ConnectorV1ConnectionNewParamsScopeTierSmokeballUsStaging, ConnectorV1ConnectionNewParamsScopeTierSmokeballUsStagingWrite, ConnectorV1ConnectionNewParamsScopeTierSmokeballUsWrite:
		return true
	}
	return false
}

type ConnectorV1ConnectionGetParams struct {
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

type ConnectorV1ConnectionListParams struct {
	// Opaque continuation cursor from `pagination.next_cursor` of the previous page.
	// Must be replayed with the same filters and scope that produced it.
	Cursor param.Field[string] `query:"cursor"`
	// Connections per page (1-200). Defaults to 200.
	Limit                 param.Field[int64]                                 `query:"limit"`
	Provider              param.Field[string]                                `query:"provider"`
	Status                param.Field[ConnectorV1ConnectionListParamsStatus] `query:"status"`
	XCaseConnectorSubject param.Field[string]                                `header:"x-case-connector-subject"`
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
	Purge                 param.Field[bool]   `query:"purge"`
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
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
	Site                  param.Field[string] `query:"site"`
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

// URLQuery serializes [ConnectorV1ConnectionBrowseParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1ConnectionBrowseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1ConnectionNewFolderParams struct {
	// Folder name. Trimmed; 1-255 characters; no '/', '\', or control characters; not
	// '.' or '..'. Providers may refuse further characters.
	Name param.Field[string] `json:"name" api:"required"`
	// Container id of the parent, from its browse_ref
	Container param.Field[string] `json:"container"`
	// Folder id of the parent, from its browse_ref. Omit for a container root.
	Parent param.Field[string] `json:"parent"`
	// Site id of the parent, from its browse_ref
	Site                  param.Field[string] `json:"site"`
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

func (r ConnectorV1ConnectionNewFolderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1ConnectionNewMicrosoftAdminApprovalParams struct {
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

type ConnectorV1ConnectionListContentsParams struct {
	// Container to list, or the container holding parent
	Container param.Field[string] `query:"container"`
	Cursor    param.Field[string] `query:"cursor"`
	PageSize  param.Field[int64]  `query:"page_size"`
	// Folder to list
	Parent param.Field[string] `query:"parent"`
	// Site of the folder
	Site                  param.Field[string] `query:"site"`
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

// URLQuery serializes [ConnectorV1ConnectionListContentsParams]'s query parameters
// as `url.Values`.
func (r ConnectorV1ConnectionListContentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1ConnectionResolveLinkParams struct {
	// Folder link copied from the provider, including any query string
	URL                   param.Field[string] `query:"url" api:"required"`
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

// URLQuery serializes [ConnectorV1ConnectionResolveLinkParams]'s query parameters
// as `url.Values`.
func (r ConnectorV1ConnectionResolveLinkParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1ConnectionSearchParams struct {
	// Search text
	Query param.Field[string] `query:"query" api:"required"`
	// Container to search in, or the container holding parent
	Container param.Field[string] `query:"container"`
	Cursor    param.Field[string] `query:"cursor"`
	// What to search for: files, or folders by name
	Kind     param.Field[ConnectorV1ConnectionSearchParamsKind] `query:"kind"`
	PageSize param.Field[int64]                                 `query:"page_size"`
	// Folder to search in
	Parent param.Field[string] `query:"parent"`
	// Site to search in
	Site                  param.Field[string] `query:"site"`
	XCaseConnectorSubject param.Field[string] `header:"x-case-connector-subject"`
}

// URLQuery serializes [ConnectorV1ConnectionSearchParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1ConnectionSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// What to search for: files, or folders by name
type ConnectorV1ConnectionSearchParamsKind string

const (
	ConnectorV1ConnectionSearchParamsKindFile   ConnectorV1ConnectionSearchParamsKind = "file"
	ConnectorV1ConnectionSearchParamsKindFolder ConnectorV1ConnectionSearchParamsKind = "folder"
)

func (r ConnectorV1ConnectionSearchParamsKind) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionSearchParamsKindFile, ConnectorV1ConnectionSearchParamsKindFolder:
		return true
	}
	return false
}

type ConnectorV1ConnectionUpdateAllParams struct {
	// Confirms that this change applies to every user connection in scope.
	ConfirmOrganizationWide param.Field[ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWide] `json:"confirm_organization_wide" api:"required"`
	Enabled                 param.Field[bool]                                                        `json:"enabled" api:"required"`
	Provider                param.Field[string]                                                      `json:"provider" api:"required"`
}

func (r ConnectorV1ConnectionUpdateAllParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Confirms that this change applies to every user connection in scope.
type ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWide bool

const (
	ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWideTrue ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWide = true
)

func (r ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWide) IsKnown() bool {
	switch r {
	case ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWideTrue:
		return true
	}
	return false
}
