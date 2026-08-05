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
// ConnectorV1LinkService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1LinkService] method instead.
type ConnectorV1LinkService struct {
	Options []option.RequestOption
}

// NewConnectorV1LinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorV1LinkService(opts ...option.RequestOption) (r *ConnectorV1LinkService) {
	r = &ConnectorV1LinkService{}
	r.Options = opts
	return
}

// Retrieve one link: state, counts, and embedded active_run/last_run. Poll this
// after POST /transfer.
func (r *ConnectorV1LinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Pause/resume a link (state "paused" | "ready"), change its mode (synced -> once
// is the sync downgrade), or edit its policy in place.
func (r *ConnectorV1LinkService) Update(ctx context.Context, id string, body ConnectorV1LinkUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List transfer links, filterable by vault, connection, direction, mode, and
// state.
func (r *ConnectorV1LinkService) List(ctx context.Context, query ConnectorV1LinkListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "connectors/v1/links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Delete a link and its ledger. vault_docs=delete additionally removes the vault
// documents an import link brought in (default: keep).
func (r *ConnectorV1LinkService) Delete(ctx context.Context, id string, body ConnectorV1LinkDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Per-file transfer ledger for a link: provider item, vault object, path, content
// version, state, and error.
func (r *ConnectorV1LinkService) ListObjects(ctx context.Context, id string, query ConnectorV1LinkListObjectsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/links/%s/objects", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type ConnectorV1LinkUpdateParams struct {
	Mode   param.Field[ConnectorV1LinkUpdateParamsMode]  `json:"mode"`
	Policy param.Field[interface{}]                      `json:"policy"`
	State  param.Field[ConnectorV1LinkUpdateParamsState] `json:"state"`
}

func (r ConnectorV1LinkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1LinkUpdateParamsMode string

const (
	ConnectorV1LinkUpdateParamsModeOnce   ConnectorV1LinkUpdateParamsMode = "once"
	ConnectorV1LinkUpdateParamsModeSynced ConnectorV1LinkUpdateParamsMode = "synced"
)

func (r ConnectorV1LinkUpdateParamsMode) IsKnown() bool {
	switch r {
	case ConnectorV1LinkUpdateParamsModeOnce, ConnectorV1LinkUpdateParamsModeSynced:
		return true
	}
	return false
}

type ConnectorV1LinkUpdateParamsState string

const (
	ConnectorV1LinkUpdateParamsStatePaused ConnectorV1LinkUpdateParamsState = "paused"
	ConnectorV1LinkUpdateParamsStateReady  ConnectorV1LinkUpdateParamsState = "ready"
)

func (r ConnectorV1LinkUpdateParamsState) IsKnown() bool {
	switch r {
	case ConnectorV1LinkUpdateParamsStatePaused, ConnectorV1LinkUpdateParamsStateReady:
		return true
	}
	return false
}

type ConnectorV1LinkListParams struct {
	ConnectionID param.Field[string]                             `query:"connection_id"`
	Direction    param.Field[ConnectorV1LinkListParamsDirection] `query:"direction"`
	Mode         param.Field[ConnectorV1LinkListParamsMode]      `query:"mode"`
	PairID       param.Field[string]                             `query:"pair_id"`
	State        param.Field[ConnectorV1LinkListParamsState]     `query:"state"`
	VaultID      param.Field[string]                             `query:"vault_id"`
}

// URLQuery serializes [ConnectorV1LinkListParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1LinkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1LinkListParamsDirection string

const (
	ConnectorV1LinkListParamsDirectionImport ConnectorV1LinkListParamsDirection = "import"
	ConnectorV1LinkListParamsDirectionExport ConnectorV1LinkListParamsDirection = "export"
)

func (r ConnectorV1LinkListParamsDirection) IsKnown() bool {
	switch r {
	case ConnectorV1LinkListParamsDirectionImport, ConnectorV1LinkListParamsDirectionExport:
		return true
	}
	return false
}

type ConnectorV1LinkListParamsMode string

const (
	ConnectorV1LinkListParamsModeOnce   ConnectorV1LinkListParamsMode = "once"
	ConnectorV1LinkListParamsModeSynced ConnectorV1LinkListParamsMode = "synced"
)

func (r ConnectorV1LinkListParamsMode) IsKnown() bool {
	switch r {
	case ConnectorV1LinkListParamsModeOnce, ConnectorV1LinkListParamsModeSynced:
		return true
	}
	return false
}

type ConnectorV1LinkListParamsState string

const (
	ConnectorV1LinkListParamsStateReady    ConnectorV1LinkListParamsState = "ready"
	ConnectorV1LinkListParamsStateRunning  ConnectorV1LinkListParamsState = "running"
	ConnectorV1LinkListParamsStateActive   ConnectorV1LinkListParamsState = "active"
	ConnectorV1LinkListParamsStatePaused   ConnectorV1LinkListParamsState = "paused"
	ConnectorV1LinkListParamsStateOrphaned ConnectorV1LinkListParamsState = "orphaned"
	ConnectorV1LinkListParamsStateError    ConnectorV1LinkListParamsState = "error"
)

func (r ConnectorV1LinkListParamsState) IsKnown() bool {
	switch r {
	case ConnectorV1LinkListParamsStateReady, ConnectorV1LinkListParamsStateRunning, ConnectorV1LinkListParamsStateActive, ConnectorV1LinkListParamsStatePaused, ConnectorV1LinkListParamsStateOrphaned, ConnectorV1LinkListParamsStateError:
		return true
	}
	return false
}

type ConnectorV1LinkDeleteParams struct {
	VaultDocs param.Field[ConnectorV1LinkDeleteParamsVaultDocs] `query:"vault_docs"`
}

// URLQuery serializes [ConnectorV1LinkDeleteParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1LinkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1LinkDeleteParamsVaultDocs string

const (
	ConnectorV1LinkDeleteParamsVaultDocsKeep   ConnectorV1LinkDeleteParamsVaultDocs = "keep"
	ConnectorV1LinkDeleteParamsVaultDocsDelete ConnectorV1LinkDeleteParamsVaultDocs = "delete"
)

func (r ConnectorV1LinkDeleteParamsVaultDocs) IsKnown() bool {
	switch r {
	case ConnectorV1LinkDeleteParamsVaultDocsKeep, ConnectorV1LinkDeleteParamsVaultDocsDelete:
		return true
	}
	return false
}

type ConnectorV1LinkListObjectsParams struct {
	Cursor param.Field[string]                                `query:"cursor"`
	State  param.Field[ConnectorV1LinkListObjectsParamsState] `query:"state"`
}

// URLQuery serializes [ConnectorV1LinkListObjectsParams]'s query parameters as
// `url.Values`.
func (r ConnectorV1LinkListObjectsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectorV1LinkListObjectsParamsState string

const (
	ConnectorV1LinkListObjectsParamsStatePending      ConnectorV1LinkListObjectsParamsState = "pending"
	ConnectorV1LinkListObjectsParamsStateTransferring ConnectorV1LinkListObjectsParamsState = "transferring"
	ConnectorV1LinkListObjectsParamsStateIngesting    ConnectorV1LinkListObjectsParamsState = "ingesting"
	ConnectorV1LinkListObjectsParamsStateSynced       ConnectorV1LinkListObjectsParamsState = "synced"
	ConnectorV1LinkListObjectsParamsStateSkipped      ConnectorV1LinkListObjectsParamsState = "skipped"
	ConnectorV1LinkListObjectsParamsStateFailed       ConnectorV1LinkListObjectsParamsState = "failed"
	ConnectorV1LinkListObjectsParamsStateTombstoned   ConnectorV1LinkListObjectsParamsState = "tombstoned"
)

func (r ConnectorV1LinkListObjectsParamsState) IsKnown() bool {
	switch r {
	case ConnectorV1LinkListObjectsParamsStatePending, ConnectorV1LinkListObjectsParamsStateTransferring, ConnectorV1LinkListObjectsParamsStateIngesting, ConnectorV1LinkListObjectsParamsStateSynced, ConnectorV1LinkListObjectsParamsStateSkipped, ConnectorV1LinkListObjectsParamsStateFailed, ConnectorV1LinkListObjectsParamsStateTombstoned:
		return true
	}
	return false
}
