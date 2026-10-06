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

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1Service] method instead.
type MatterV1Service struct {
	Options []option.RequestOption
	// Matter-native legal workspaces and orchestration primitives
	Purges *MatterV1PurgeService
	// Matter-native legal workspaces and orchestration primitives
	ContentPurges *MatterV1ContentPurgeService
	// Matter-native legal workspaces and orchestration primitives
	AgentTypes *MatterV1AgentTypeService
	// Matter-native legal workspaces and orchestration primitives
	Parties *MatterV1PartyService
	// Matter-native legal workspaces and orchestration primitives
	Types  *MatterV1TypeService
	Events *MatterV1EventService
	// Matter-native legal workspaces and orchestration primitives
	Log *MatterV1LogService
	// Matter-native legal workspaces and orchestration primitives
	MatterParties *MatterV1MatterPartyService
	// Matter-native legal workspaces and orchestration primitives
	Shares *MatterV1ShareService
	// Matter-native legal workspaces and orchestration primitives
	WorkItems *MatterV1WorkItemService
}

// NewMatterV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1Service(opts ...option.RequestOption) (r *MatterV1Service) {
	r = &MatterV1Service{}
	r.Options = opts
	r.Purges = NewMatterV1PurgeService(opts...)
	r.ContentPurges = NewMatterV1ContentPurgeService(opts...)
	r.AgentTypes = NewMatterV1AgentTypeService(opts...)
	r.Parties = NewMatterV1PartyService(opts...)
	r.Types = NewMatterV1TypeService(opts...)
	r.Events = NewMatterV1EventService(opts...)
	r.Log = NewMatterV1LogService(opts...)
	r.MatterParties = NewMatterV1MatterPartyService(opts...)
	r.Shares = NewMatterV1ShareService(opts...)
	r.WorkItems = NewMatterV1WorkItemService(opts...)
	return
}

// Create a new legal matter and optionally link an existing primary vault.
func (r *MatterV1Service) New(ctx context.Context, body MatterV1NewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get a single matter by ID.
func (r *MatterV1Service) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Update mutable matter fields.
func (r *MatterV1Service) Update(ctx context.Context, id string, body MatterV1UpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List matters for the authenticated organization, newest update first. Pagination
// is opt-in: pass `limit` (1-200) to receive a bounded page, then replay
// `pagination.next_cursor` as `?cursor=` while `pagination.has_more` is true.
// Cursors are opaque and are only valid for the exact filter set they were issued
// under. A request with neither `limit` nor `cursor` still returns every matter,
// and `pagination.limit` is null. That default will become a bounded page in a
// future release — paginate now to avoid the change.
func (r *MatterV1Service) List(ctx context.Context, query MatterV1ListParams, opts ...option.RequestOption) (res *MatterV1ListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "matters/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Queues a durable, idempotent purge of a Matter and all linked live content. Use
// matter purge webhooks for status changes; the inspection route is intended for
// manual diagnostics only.
func (r *MatterV1Service) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MatterV1DeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matters/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type MatterV1ListResponse struct {
	Data       []interface{}                  `json:"data"`
	Pagination MatterV1ListResponsePagination `json:"pagination"`
	JSON       matterV1ListResponseJSON       `json:"-"`
}

// matterV1ListResponseJSON contains the JSON metadata for the struct
// [MatterV1ListResponse]
type matterV1ListResponseJSON struct {
	Data        apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatterV1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1ListResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1ListResponsePagination struct {
	HasMore    bool                               `json:"has_more"`
	Limit      int64                              `json:"limit" api:"nullable"`
	NextCursor string                             `json:"next_cursor" api:"nullable"`
	JSON       matterV1ListResponsePaginationJSON `json:"-"`
}

// matterV1ListResponsePaginationJSON contains the JSON metadata for the struct
// [MatterV1ListResponsePagination]
type matterV1ListResponsePaginationJSON struct {
	HasMore     apijson.Field
	Limit       apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatterV1ListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1ListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type MatterV1DeleteResponse struct {
	Attempt     int64                        `json:"attempt" api:"required"`
	MatterID    string                       `json:"matter_id" api:"required"`
	PurgeID     string                       `json:"purge_id" api:"required"`
	Status      MatterV1DeleteResponseStatus `json:"status" api:"required"`
	VaultID     string                       `json:"vault_id" api:"required"`
	WorkflowID  string                       `json:"workflow_id" api:"required,nullable"`
	CompletedAt time.Time                    `json:"completed_at" api:"nullable" format:"date-time"`
	Counts      MatterV1DeleteResponseCounts `json:"counts"`
	FailedAt    time.Time                    `json:"failed_at" api:"nullable" format:"date-time"`
	FailureCode string                       `json:"failure_code" api:"nullable"`
	RequestedAt time.Time                    `json:"requested_at" format:"date-time"`
	StartedAt   time.Time                    `json:"started_at" api:"nullable" format:"date-time"`
	// Stable ID of the failed or completed terminal webhook event
	TerminalEventID string                     `json:"terminal_event_id" api:"nullable"`
	JSON            matterV1DeleteResponseJSON `json:"-"`
}

// matterV1DeleteResponseJSON contains the JSON metadata for the struct
// [MatterV1DeleteResponse]
type matterV1DeleteResponseJSON struct {
	Attempt         apijson.Field
	MatterID        apijson.Field
	PurgeID         apijson.Field
	Status          apijson.Field
	VaultID         apijson.Field
	WorkflowID      apijson.Field
	CompletedAt     apijson.Field
	Counts          apijson.Field
	FailedAt        apijson.Field
	FailureCode     apijson.Field
	RequestedAt     apijson.Field
	StartedAt       apijson.Field
	TerminalEventID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MatterV1DeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1DeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1DeleteResponseStatus string

const (
	MatterV1DeleteResponseStatusQueued     MatterV1DeleteResponseStatus = "queued"
	MatterV1DeleteResponseStatusInProgress MatterV1DeleteResponseStatus = "in_progress"
	MatterV1DeleteResponseStatusFailed     MatterV1DeleteResponseStatus = "failed"
	MatterV1DeleteResponseStatusCompleted  MatterV1DeleteResponseStatus = "completed"
)

func (r MatterV1DeleteResponseStatus) IsKnown() bool {
	switch r {
	case MatterV1DeleteResponseStatusQueued, MatterV1DeleteResponseStatusInProgress, MatterV1DeleteResponseStatusFailed, MatterV1DeleteResponseStatusCompleted:
		return true
	}
	return false
}

type MatterV1DeleteResponseCounts struct {
	Chats          int64                            `json:"chats"`
	Objects        int64                            `json:"objects"`
	Sessions       int64                            `json:"sessions"`
	Transcriptions int64                            `json:"transcriptions"`
	JSON           matterV1DeleteResponseCountsJSON `json:"-"`
}

// matterV1DeleteResponseCountsJSON contains the JSON metadata for the struct
// [MatterV1DeleteResponseCounts]
type matterV1DeleteResponseCountsJSON struct {
	Chats          apijson.Field
	Objects        apijson.Field
	Sessions       apijson.Field
	Transcriptions apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MatterV1DeleteResponseCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1DeleteResponseCountsJSON) RawJSON() string {
	return r.raw
}

type MatterV1NewParams struct {
	Title                 param.Field[string]                  `json:"title" api:"required"`
	Billing               param.Field[map[string]interface{}]  `json:"billing"`
	ClientName            param.Field[string]                  `json:"client_name"`
	ClientPartyID         param.Field[string]                  `json:"client_party_id"`
	CustomFields          param.Field[map[string]interface{}]  `json:"custom_fields"`
	Description           param.Field[string]                  `json:"description"`
	DisplayID             param.Field[string]                  `json:"display_id"`
	ImportantDates        param.Field[map[string]interface{}]  `json:"important_dates"`
	Jurisdiction          param.Field[map[string]interface{}]  `json:"jurisdiction"`
	MatterType            param.Field[string]                  `json:"matter_type"`
	Metadata              param.Field[map[string]interface{}]  `json:"metadata"`
	PracticeArea          param.Field[string]                  `json:"practice_area"`
	ResponsibleAttorneyID param.Field[string]                  `json:"responsible_attorney_id"`
	Status                param.Field[MatterV1NewParamsStatus] `json:"status"`
	Subtype               param.Field[string]                  `json:"subtype"`
	Vault                 param.Field[MatterV1NewParamsVault]  `json:"vault"`
	VaultID               param.Field[string]                  `json:"vault_id"`
}

func (r MatterV1NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1NewParamsStatus string

const (
	MatterV1NewParamsStatusIntake   MatterV1NewParamsStatus = "intake"
	MatterV1NewParamsStatusOpen     MatterV1NewParamsStatus = "open"
	MatterV1NewParamsStatusPending  MatterV1NewParamsStatus = "pending"
	MatterV1NewParamsStatusClosed   MatterV1NewParamsStatus = "closed"
	MatterV1NewParamsStatusArchived MatterV1NewParamsStatus = "archived"
)

func (r MatterV1NewParamsStatus) IsKnown() bool {
	switch r {
	case MatterV1NewParamsStatusIntake, MatterV1NewParamsStatusOpen, MatterV1NewParamsStatusPending, MatterV1NewParamsStatusClosed, MatterV1NewParamsStatusArchived:
		return true
	}
	return false
}

type MatterV1NewParamsVault struct {
	Description    param.Field[string]                 `json:"description"`
	EnableIndexing param.Field[bool]                   `json:"enableIndexing"`
	Metadata       param.Field[map[string]interface{}] `json:"metadata"`
}

func (r MatterV1NewParamsVault) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1UpdateParams struct {
	ArchivedAt            param.Field[time.Time]                  `json:"archived_at" format:"date-time"`
	Billing               param.Field[map[string]interface{}]     `json:"billing"`
	ClientName            param.Field[string]                     `json:"client_name"`
	ClientPartyID         param.Field[string]                     `json:"client_party_id"`
	CustomFields          param.Field[map[string]interface{}]     `json:"custom_fields"`
	Description           param.Field[string]                     `json:"description"`
	DisplayID             param.Field[string]                     `json:"display_id"`
	ImportantDates        param.Field[map[string]interface{}]     `json:"important_dates"`
	Jurisdiction          param.Field[map[string]interface{}]     `json:"jurisdiction"`
	MatterType            param.Field[string]                     `json:"matter_type"`
	Metadata              param.Field[map[string]interface{}]     `json:"metadata"`
	PracticeArea          param.Field[string]                     `json:"practice_area"`
	ResponsibleAttorneyID param.Field[string]                     `json:"responsible_attorney_id"`
	Status                param.Field[MatterV1UpdateParamsStatus] `json:"status"`
	Subtype               param.Field[string]                     `json:"subtype"`
	Title                 param.Field[string]                     `json:"title"`
}

func (r MatterV1UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1UpdateParamsStatus string

const (
	MatterV1UpdateParamsStatusIntake   MatterV1UpdateParamsStatus = "intake"
	MatterV1UpdateParamsStatusOpen     MatterV1UpdateParamsStatus = "open"
	MatterV1UpdateParamsStatusPending  MatterV1UpdateParamsStatus = "pending"
	MatterV1UpdateParamsStatusClosed   MatterV1UpdateParamsStatus = "closed"
	MatterV1UpdateParamsStatusArchived MatterV1UpdateParamsStatus = "archived"
)

func (r MatterV1UpdateParamsStatus) IsKnown() bool {
	switch r {
	case MatterV1UpdateParamsStatusIntake, MatterV1UpdateParamsStatusOpen, MatterV1UpdateParamsStatusPending, MatterV1UpdateParamsStatusClosed, MatterV1UpdateParamsStatusArchived:
		return true
	}
	return false
}

type MatterV1ListParams struct {
	// Opaque continuation cursor from `pagination.next_cursor` of the previous page.
	// Must be replayed with the same filters that produced it.
	Cursor param.Field[string] `query:"cursor"`
	// Matters per page (1-200). Omit to receive every matter. Supplying a cursor
	// without a limit uses 50.
	Limit        param.Field[int64]  `query:"limit"`
	MatterType   param.Field[string] `query:"matter_type"`
	PracticeArea param.Field[string] `query:"practice_area"`
	Query        param.Field[string] `query:"query"`
	Status       param.Field[string] `query:"status"`
}

// URLQuery serializes [MatterV1ListParams]'s query parameters as `url.Values`.
func (r MatterV1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
