// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Import and export between provider folders (Google Drive) and vaults
//
// ConnectorV1Service contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1Service] method instead.
type ConnectorV1Service struct {
	Options []option.RequestOption
	// Import and export between provider folders (Google Drive) and vaults
	Installations *ConnectorV1InstallationService
	// Import and export between provider folders (Google Drive) and vaults
	Connections *ConnectorV1ConnectionService
	// Import and export between provider folders (Google Drive) and vaults
	Links *ConnectorV1LinkService
}

// NewConnectorV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorV1Service(opts ...option.RequestOption) (r *ConnectorV1Service) {
	r = &ConnectorV1Service{}
	r.Options = opts
	r.Installations = NewConnectorV1InstallationService(opts...)
	r.Connections = NewConnectorV1ConnectionService(opts...)
	r.Links = NewConnectorV1LinkService(opts...)
	return
}

// Standing promise: backfill now, then stay current (the sync sweeper re-runs
// synced links on a schedule). Direction both creates paired import/export links
// and defaults export to a CaseMark Output subfolder. Same body as /transfer minus
// run_mode. Upserts links by (connection_id, direction, remote, vault_id);
// existing once-links are upgraded in place with their ledger and cursor
// preserved. Downgrade or pause via PATCH /links/{id}.
func (r *ConnectorV1Service) SyncLink(ctx context.Context, body ConnectorV1SyncLinkParams, opts ...option.RequestOption) (res *ConnectorV1SyncLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/sync-link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// One-shot import (provider folder → vault), export (vault → provider folder), or
// both. Direction both creates paired import/export links and defaults export to a
// CaseMark Output subfolder. Upserts links by (connection_id, direction, remote,
// vault_id): first call backfills, later calls move only new/changed files via the
// ledger. Poll GET /links/{id} → active_run for progress.
func (r *ConnectorV1Service) Transfer(ctx context.Context, body ConnectorV1TransferParams, opts ...option.RequestOption) (res *ConnectorV1TransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConnectorV1SyncLinkResponse struct {
	Links []interface{}                   `json:"links"`
	JSON  connectorV1SyncLinkResponseJSON `json:"-"`
}

// connectorV1SyncLinkResponseJSON contains the JSON metadata for the struct
// [ConnectorV1SyncLinkResponse]
type connectorV1SyncLinkResponseJSON struct {
	Links       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1SyncLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1SyncLinkResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1TransferResponse struct {
	Links []interface{}                   `json:"links"`
	JSON  connectorV1TransferResponseJSON `json:"-"`
}

// connectorV1TransferResponseJSON contains the JSON metadata for the struct
// [ConnectorV1TransferResponse]
type connectorV1TransferResponseJSON struct {
	Links       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorV1TransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1TransferResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1SyncLinkParams struct {
	ConnectionID param.Field[string]                             `json:"connection_id" api:"required"`
	Direction    param.Field[ConnectorV1SyncLinkParamsDirection] `json:"direction" api:"required"`
	Remote       param.Field[ConnectorV1SyncLinkParamsRemote]    `json:"remote" api:"required"`
	VaultID      param.Field[string]                             `json:"vault_id" api:"required"`
	// Optional destination for direction both. Defaults to CaseMark Output under
	// remote.
	ExportDestination param.Field[ConnectorV1SyncLinkParamsExportDestination] `json:"export_destination"`
	MatterID          param.Field[string]                                     `json:"matter_id"`
	Policy            param.Field[ConnectorV1SyncLinkParamsPolicy]            `json:"policy"`
}

func (r ConnectorV1SyncLinkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1SyncLinkParamsDirection string

const (
	ConnectorV1SyncLinkParamsDirectionImport ConnectorV1SyncLinkParamsDirection = "import"
	ConnectorV1SyncLinkParamsDirectionExport ConnectorV1SyncLinkParamsDirection = "export"
	ConnectorV1SyncLinkParamsDirectionBoth   ConnectorV1SyncLinkParamsDirection = "both"
)

func (r ConnectorV1SyncLinkParamsDirection) IsKnown() bool {
	switch r {
	case ConnectorV1SyncLinkParamsDirectionImport, ConnectorV1SyncLinkParamsDirectionExport, ConnectorV1SyncLinkParamsDirectionBoth:
		return true
	}
	return false
}

type ConnectorV1SyncLinkParamsRemote struct {
	FolderID    param.Field[string] `json:"folder_id" api:"required"`
	ContainerID param.Field[string] `json:"container_id"`
	Path        param.Field[string] `json:"path"`
	SiteID      param.Field[string] `json:"site_id"`
}

func (r ConnectorV1SyncLinkParamsRemote) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional destination for direction both. Defaults to CaseMark Output under
// remote.
type ConnectorV1SyncLinkParamsExportDestination struct {
	FolderID    param.Field[string] `json:"folder_id" api:"required"`
	ContainerID param.Field[string] `json:"container_id"`
	Path        param.Field[string] `json:"path"`
	SiteID      param.Field[string] `json:"site_id"`
}

func (r ConnectorV1SyncLinkParamsExportDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1SyncLinkParamsPolicy struct {
	Collisions param.Field[ConnectorV1SyncLinkParamsPolicyCollisions] `json:"collisions"`
	Deletes    param.Field[ConnectorV1SyncLinkParamsPolicyDeletes]    `json:"deletes"`
	Filters    param.Field[ConnectorV1SyncLinkParamsPolicyFilters]    `json:"filters"`
}

func (r ConnectorV1SyncLinkParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1SyncLinkParamsPolicyCollisions string

const (
	ConnectorV1SyncLinkParamsPolicyCollisionsVersion   ConnectorV1SyncLinkParamsPolicyCollisions = "version"
	ConnectorV1SyncLinkParamsPolicyCollisionsOverwrite ConnectorV1SyncLinkParamsPolicyCollisions = "overwrite"
	ConnectorV1SyncLinkParamsPolicyCollisionsSkip      ConnectorV1SyncLinkParamsPolicyCollisions = "skip"
)

func (r ConnectorV1SyncLinkParamsPolicyCollisions) IsKnown() bool {
	switch r {
	case ConnectorV1SyncLinkParamsPolicyCollisionsVersion, ConnectorV1SyncLinkParamsPolicyCollisionsOverwrite, ConnectorV1SyncLinkParamsPolicyCollisionsSkip:
		return true
	}
	return false
}

type ConnectorV1SyncLinkParamsPolicyDeletes string

const (
	ConnectorV1SyncLinkParamsPolicyDeletesMirror   ConnectorV1SyncLinkParamsPolicyDeletes = "mirror"
	ConnectorV1SyncLinkParamsPolicyDeletesPreserve ConnectorV1SyncLinkParamsPolicyDeletes = "preserve"
)

func (r ConnectorV1SyncLinkParamsPolicyDeletes) IsKnown() bool {
	switch r {
	case ConnectorV1SyncLinkParamsPolicyDeletesMirror, ConnectorV1SyncLinkParamsPolicyDeletesPreserve:
		return true
	}
	return false
}

type ConnectorV1SyncLinkParamsPolicyFilters struct {
	ExcludeMime  param.Field[[]string] `json:"exclude_mime"`
	MaxSizeBytes param.Field[int64]    `json:"max_size_bytes"`
}

func (r ConnectorV1SyncLinkParamsPolicyFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1TransferParams struct {
	ConnectionID param.Field[string]                             `json:"connection_id" api:"required"`
	Direction    param.Field[ConnectorV1TransferParamsDirection] `json:"direction" api:"required"`
	Remote       param.Field[ConnectorV1TransferParamsRemote]    `json:"remote" api:"required"`
	VaultID      param.Field[string]                             `json:"vault_id" api:"required"`
	// Optional destination for direction both. Defaults to CaseMark Output under
	// remote.
	ExportDestination param.Field[ConnectorV1TransferParamsExportDestination] `json:"export_destination"`
	MatterID          param.Field[string]                                     `json:"matter_id"`
	Policy            param.Field[ConnectorV1TransferParamsPolicy]            `json:"policy"`
	RunMode           param.Field[ConnectorV1TransferParamsRunMode]           `json:"run_mode"`
}

func (r ConnectorV1TransferParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1TransferParamsDirection string

const (
	ConnectorV1TransferParamsDirectionImport ConnectorV1TransferParamsDirection = "import"
	ConnectorV1TransferParamsDirectionExport ConnectorV1TransferParamsDirection = "export"
	ConnectorV1TransferParamsDirectionBoth   ConnectorV1TransferParamsDirection = "both"
)

func (r ConnectorV1TransferParamsDirection) IsKnown() bool {
	switch r {
	case ConnectorV1TransferParamsDirectionImport, ConnectorV1TransferParamsDirectionExport, ConnectorV1TransferParamsDirectionBoth:
		return true
	}
	return false
}

type ConnectorV1TransferParamsRemote struct {
	FolderID    param.Field[string] `json:"folder_id" api:"required"`
	ContainerID param.Field[string] `json:"container_id"`
	Path        param.Field[string] `json:"path"`
	SiteID      param.Field[string] `json:"site_id"`
}

func (r ConnectorV1TransferParamsRemote) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional destination for direction both. Defaults to CaseMark Output under
// remote.
type ConnectorV1TransferParamsExportDestination struct {
	FolderID    param.Field[string] `json:"folder_id" api:"required"`
	ContainerID param.Field[string] `json:"container_id"`
	Path        param.Field[string] `json:"path"`
	SiteID      param.Field[string] `json:"site_id"`
}

func (r ConnectorV1TransferParamsExportDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1TransferParamsPolicy struct {
	Collisions param.Field[ConnectorV1TransferParamsPolicyCollisions] `json:"collisions"`
	Deletes    param.Field[ConnectorV1TransferParamsPolicyDeletes]    `json:"deletes"`
	Filters    param.Field[ConnectorV1TransferParamsPolicyFilters]    `json:"filters"`
}

func (r ConnectorV1TransferParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1TransferParamsPolicyCollisions string

const (
	ConnectorV1TransferParamsPolicyCollisionsVersion   ConnectorV1TransferParamsPolicyCollisions = "version"
	ConnectorV1TransferParamsPolicyCollisionsOverwrite ConnectorV1TransferParamsPolicyCollisions = "overwrite"
	ConnectorV1TransferParamsPolicyCollisionsSkip      ConnectorV1TransferParamsPolicyCollisions = "skip"
)

func (r ConnectorV1TransferParamsPolicyCollisions) IsKnown() bool {
	switch r {
	case ConnectorV1TransferParamsPolicyCollisionsVersion, ConnectorV1TransferParamsPolicyCollisionsOverwrite, ConnectorV1TransferParamsPolicyCollisionsSkip:
		return true
	}
	return false
}

type ConnectorV1TransferParamsPolicyDeletes string

const (
	ConnectorV1TransferParamsPolicyDeletesMirror   ConnectorV1TransferParamsPolicyDeletes = "mirror"
	ConnectorV1TransferParamsPolicyDeletesPreserve ConnectorV1TransferParamsPolicyDeletes = "preserve"
)

func (r ConnectorV1TransferParamsPolicyDeletes) IsKnown() bool {
	switch r {
	case ConnectorV1TransferParamsPolicyDeletesMirror, ConnectorV1TransferParamsPolicyDeletesPreserve:
		return true
	}
	return false
}

type ConnectorV1TransferParamsPolicyFilters struct {
	ExcludeMime  param.Field[[]string] `json:"exclude_mime"`
	MaxSizeBytes param.Field[int64]    `json:"max_size_bytes"`
}

func (r ConnectorV1TransferParamsPolicyFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1TransferParamsRunMode string

const (
	ConnectorV1TransferParamsRunModeAuto          ConnectorV1TransferParamsRunMode = "auto"
	ConnectorV1TransferParamsRunModeFullReconcile ConnectorV1TransferParamsRunMode = "full_reconcile"
)

func (r ConnectorV1TransferParamsRunMode) IsKnown() bool {
	switch r {
	case ConnectorV1TransferParamsRunModeAuto, ConnectorV1TransferParamsRunModeFullReconcile:
		return true
	}
	return false
}
