// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Import and export between provider folders and vaults
//
// ConnectorV1Service contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1Service] method instead.
type ConnectorV1Service struct {
	Options      []option.RequestOption
	Applications *ConnectorV1ApplicationService
	// Import and export between provider folders and vaults
	Installations *ConnectorV1InstallationService
	// Import and export between provider folders and vaults
	Connections *ConnectorV1ConnectionService
	// Import and export between provider folders and vaults
	Links *ConnectorV1LinkService
}

// NewConnectorV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorV1Service(opts ...option.RequestOption) (r *ConnectorV1Service) {
	r = &ConnectorV1Service{}
	r.Options = opts
	r.Applications = NewConnectorV1ApplicationService(opts...)
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
func (r *ConnectorV1Service) SyncLink(ctx context.Context, params ConnectorV1SyncLinkParams, opts ...option.RequestOption) (res *ConnectorV1SyncLinkResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/sync-link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// One-shot import (provider folder → vault), export (vault → provider folder), or
// both. Direction both creates paired import/export links and defaults export to a
// CaseMark Output subfolder. Upserts links by (connection_id, direction, remote,
// vault_id): first call backfills, later calls move only new/changed files via the
// ledger. Poll GET /links/{id} → active_run for progress.
func (r *ConnectorV1Service) Transfer(ctx context.Context, params ConnectorV1TransferParams, opts ...option.RequestOption) (res *ConnectorV1TransferResponse, err error) {
	if params.XCaseConnectorSubject.Present {
		opts = append(opts, option.WithHeader("x-case-connector-subject", fmt.Sprintf("%v", params.XCaseConnectorSubject)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectors/v1/transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type SyncLinkLinkRun struct {
	LinkID  string                `json:"link_id" api:"required"`
	Started bool                  `json:"started" api:"required"`
	Backlog int64                 `json:"backlog"`
	Reason  SyncLinkLinkRunReason `json:"reason"`
	JSON    syncLinkLinkRunJSON   `json:"-"`
}

// syncLinkLinkRunJSON contains the JSON metadata for the struct [SyncLinkLinkRun]
type syncLinkLinkRunJSON struct {
	LinkID      apijson.Field
	Started     apijson.Field
	Backlog     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncLinkLinkRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncLinkLinkRunJSON) RawJSON() string {
	return r.raw
}

type SyncLinkLinkRunReason string

const (
	SyncLinkLinkRunReasonStaleRunRecovery SyncLinkLinkRunReason = "stale_run_recovery"
	SyncLinkLinkRunReasonAlreadyRunning   SyncLinkLinkRunReason = "already_running"
	SyncLinkLinkRunReasonIngestionBacklog SyncLinkLinkRunReason = "ingestion_backlog"
)

func (r SyncLinkLinkRunReason) IsKnown() bool {
	switch r {
	case SyncLinkLinkRunReasonStaleRunRecovery, SyncLinkLinkRunReasonAlreadyRunning, SyncLinkLinkRunReasonIngestionBacklog:
		return true
	}
	return false
}

// Whether a provider scan started or was deferred.
type SyncLinkRun struct {
	Started bool              `json:"started" api:"required"`
	Backlog int64             `json:"backlog"`
	Reason  SyncLinkRunReason `json:"reason"`
	JSON    syncLinkRunJSON   `json:"-"`
}

// syncLinkRunJSON contains the JSON metadata for the struct [SyncLinkRun]
type syncLinkRunJSON struct {
	Started     apijson.Field
	Backlog     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncLinkRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncLinkRunJSON) RawJSON() string {
	return r.raw
}

type SyncLinkRunReason string

const (
	SyncLinkRunReasonStaleRunRecovery SyncLinkRunReason = "stale_run_recovery"
	SyncLinkRunReasonAlreadyRunning   SyncLinkRunReason = "already_running"
	SyncLinkRunReasonIngestionBacklog SyncLinkRunReason = "ingestion_backlog"
)

func (r SyncLinkRunReason) IsKnown() bool {
	switch r {
	case SyncLinkRunReasonStaleRunRecovery, SyncLinkRunReasonAlreadyRunning, SyncLinkRunReasonIngestionBacklog:
		return true
	}
	return false
}

type TransferLinkRun struct {
	LinkID  string                `json:"link_id" api:"required"`
	Started bool                  `json:"started" api:"required"`
	Backlog int64                 `json:"backlog"`
	Reason  TransferLinkRunReason `json:"reason"`
	JSON    transferLinkRunJSON   `json:"-"`
}

// transferLinkRunJSON contains the JSON metadata for the struct [TransferLinkRun]
type transferLinkRunJSON struct {
	LinkID      apijson.Field
	Started     apijson.Field
	Backlog     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferLinkRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferLinkRunJSON) RawJSON() string {
	return r.raw
}

type TransferLinkRunReason string

const (
	TransferLinkRunReasonStaleRunRecovery TransferLinkRunReason = "stale_run_recovery"
	TransferLinkRunReasonAlreadyRunning   TransferLinkRunReason = "already_running"
	TransferLinkRunReasonIngestionBacklog TransferLinkRunReason = "ingestion_backlog"
)

func (r TransferLinkRunReason) IsKnown() bool {
	switch r {
	case TransferLinkRunReasonStaleRunRecovery, TransferLinkRunReasonAlreadyRunning, TransferLinkRunReasonIngestionBacklog:
		return true
	}
	return false
}

// Whether a provider scan started or was deferred.
type TransferRun struct {
	Started bool              `json:"started" api:"required"`
	Backlog int64             `json:"backlog"`
	Reason  TransferRunReason `json:"reason"`
	JSON    transferRunJSON   `json:"-"`
}

// transferRunJSON contains the JSON metadata for the struct [TransferRun]
type transferRunJSON struct {
	Started     apijson.Field
	Backlog     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransferRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transferRunJSON) RawJSON() string {
	return r.raw
}

type TransferRunReason string

const (
	TransferRunReasonStaleRunRecovery TransferRunReason = "stale_run_recovery"
	TransferRunReasonAlreadyRunning   TransferRunReason = "already_running"
	TransferRunReasonIngestionBacklog TransferRunReason = "ingestion_backlog"
)

func (r TransferRunReason) IsKnown() bool {
	switch r {
	case TransferRunReasonStaleRunRecovery, TransferRunReasonAlreadyRunning, TransferRunReasonIngestionBacklog:
		return true
	}
	return false
}

type ConnectorV1SyncLinkResponse struct {
	Links []interface{} `json:"links"`
	// Whether a provider scan started or was deferred.
	Run SyncLinkRun `json:"run"`
	// Start result for each link when direction is both.
	Runs []SyncLinkLinkRun               `json:"runs"`
	JSON connectorV1SyncLinkResponseJSON `json:"-"`
}

// connectorV1SyncLinkResponseJSON contains the JSON metadata for the struct
// [ConnectorV1SyncLinkResponse]
type connectorV1SyncLinkResponseJSON struct {
	Links       apijson.Field
	Run         apijson.Field
	Runs        apijson.Field
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
	Links []interface{} `json:"links"`
	// Whether a provider scan started or was deferred.
	Run TransferRun `json:"run"`
	// Start result for each link when direction is both.
	Runs []TransferLinkRun               `json:"runs"`
	JSON connectorV1TransferResponseJSON `json:"-"`
}

// connectorV1TransferResponseJSON contains the JSON metadata for the struct
// [ConnectorV1TransferResponse]
type connectorV1TransferResponseJSON struct {
	Links       apijson.Field
	Run         apijson.Field
	Runs        apijson.Field
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
	ExportDestination     param.Field[ConnectorV1SyncLinkParamsExportDestination] `json:"export_destination"`
	MatterID              param.Field[string]                                     `json:"matter_id"`
	Policy                param.Field[ConnectorV1SyncLinkParamsPolicy]            `json:"policy"`
	XCaseConnectorSubject param.Field[string]                                     `header:"x-case-connector-subject"`
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
	FolderID     param.Field[string] `json:"folder_id" api:"required"`
	ContainerID  param.Field[string] `json:"container_id"`
	Path         param.Field[string] `json:"path"`
	ResourceType param.Field[string] `json:"resource_type"`
	SiteID       param.Field[string] `json:"site_id"`
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
	// Skip these stable document ids before download, including future versions.
	ExcludeFileIDs param.Field[[]string] `json:"exclude_file_ids"`
	// Skip these folders and all descendants during document imports.
	ExcludeFolderIDs param.Field[[]string] `json:"exclude_folder_ids"`
	ExcludeMime      param.Field[[]string] `json:"exclude_mime"`
	MaxSizeBytes     param.Field[int64]    `json:"max_size_bytes"`
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
	ExportDestination     param.Field[ConnectorV1TransferParamsExportDestination] `json:"export_destination"`
	MatterID              param.Field[string]                                     `json:"matter_id"`
	Policy                param.Field[ConnectorV1TransferParamsPolicy]            `json:"policy"`
	RunMode               param.Field[ConnectorV1TransferParamsRunMode]           `json:"run_mode"`
	XCaseConnectorSubject param.Field[string]                                     `header:"x-case-connector-subject"`
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
	FolderID     param.Field[string] `json:"folder_id" api:"required"`
	ContainerID  param.Field[string] `json:"container_id"`
	Path         param.Field[string] `json:"path"`
	ResourceType param.Field[string] `json:"resource_type"`
	SiteID       param.Field[string] `json:"site_id"`
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
	// Skip these stable document ids before download, including future versions.
	ExcludeFileIDs param.Field[[]string] `json:"exclude_file_ids"`
	// Skip these folders and all descendants during document imports.
	ExcludeFolderIDs param.Field[[]string] `json:"exclude_folder_ids"`
	ExcludeMime      param.Field[[]string] `json:"exclude_mime"`
	MaxSizeBytes     param.Field[int64]    `json:"max_size_bytes"`
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
