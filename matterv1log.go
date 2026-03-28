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
// MatterV1LogService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1LogService] method instead.
type MatterV1LogService struct {
	Options []option.RequestOption
}

// NewMatterV1LogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1LogService(opts ...option.RequestOption) (r *MatterV1LogService) {
	r = &MatterV1LogService{}
	r.Options = opts
	return
}

// Append a manual operational note or event to a matter log.
func (r *MatterV1LogService) New(ctx context.Context, id string, body MatterV1LogNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/log", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// List the operational history for a matter.
func (r *MatterV1LogService) List(ctx context.Context, id string, query MatterV1LogListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/log", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Bulk export matter log entries for audits, visibility, and eval pipelines.
// Supports json, csv, tsv, and jsonl. Limited to 10,000 entries per request.
func (r *MatterV1LogService) Export(ctx context.Context, id string, body MatterV1LogExportParams, opts ...option.RequestOption) (res *MatterV1LogExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matters/v1/%s/log/export", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MatterV1LogExportResponse struct {
	Data []map[string]interface{}      `json:"data"`
	JSON matterV1LogExportResponseJSON `json:"-"`
}

// matterV1LogExportResponseJSON contains the JSON metadata for the struct
// [MatterV1LogExportResponse]
type matterV1LogExportResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatterV1LogExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1LogExportResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1LogNewParams struct {
	Summary    param.Field[string]                 `json:"summary" api:"required"`
	Details    param.Field[map[string]interface{}] `json:"details"`
	EventType  param.Field[string]                 `json:"event_type"`
	WorkItemID param.Field[string]                 `json:"work_item_id"`
}

func (r MatterV1LogNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1LogListParams struct {
	// Filter by actor ID
	ActorID param.Field[string] `query:"actor_id"`
	// Filter by actor type
	ActorType param.Field[string] `query:"actor_type"`
	// End of time range (ISO 8601)
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// Filter by exact event type
	EventType param.Field[string] `query:"event_type"`
	// Maximum number of log entries to return (max 200)
	Limit param.Field[int64] `query:"limit"`
	// Number of log entries to skip for pagination
	Offset param.Field[int64] `query:"offset"`
	// Filter by scope: matter, work_item, execution, sharing, all
	Scope param.Field[MatterV1LogListParamsScopeUnion] `query:"scope"`
	// Start of time range (ISO 8601)
	StartTime param.Field[time.Time] `query:"start_time" format:"date-time"`
	// Filter by work item ID
	WorkItemID param.Field[string] `query:"work_item_id"`
}

// URLQuery serializes [MatterV1LogListParams]'s query parameters as `url.Values`.
func (r MatterV1LogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by scope: matter, work_item, execution, sharing, all
//
// Satisfied by [shared.UnionString], [MatterV1LogListParamsScopeArray].
type MatterV1LogListParamsScopeUnion interface {
	ImplementsMatterV1LogListParamsScopeUnion()
}

type MatterV1LogListParamsScopeArray []string

func (r MatterV1LogListParamsScopeArray) ImplementsMatterV1LogListParamsScopeUnion() {}

type MatterV1LogExportParams struct {
	// Filter by actor ID
	ActorID param.Field[string] `json:"actor_id"`
	// Filter by actor type
	ActorType param.Field[string] `json:"actor_type"`
	// End of time range (ISO 8601)
	EndTime param.Field[time.Time] `json:"end_time" format:"date-time"`
	// Filter by exact event type
	EventType param.Field[string] `json:"event_type"`
	// Export format. Defaults to jsonl.
	Format param.Field[MatterV1LogExportParamsFormat] `json:"format"`
	// Filter by scope: matter, work_item, execution, sharing, all
	Scope param.Field[MatterV1LogExportParamsScopeUnion] `json:"scope"`
	// Start of time range (ISO 8601)
	StartTime param.Field[time.Time] `json:"start_time" format:"date-time"`
	// Filter by work item ID
	WorkItemID param.Field[string] `json:"work_item_id"`
}

func (r MatterV1LogExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Export format. Defaults to jsonl.
type MatterV1LogExportParamsFormat string

const (
	MatterV1LogExportParamsFormatJson  MatterV1LogExportParamsFormat = "json"
	MatterV1LogExportParamsFormatJSONL MatterV1LogExportParamsFormat = "jsonl"
	MatterV1LogExportParamsFormatCsv   MatterV1LogExportParamsFormat = "csv"
	MatterV1LogExportParamsFormatTsv   MatterV1LogExportParamsFormat = "tsv"
)

func (r MatterV1LogExportParamsFormat) IsKnown() bool {
	switch r {
	case MatterV1LogExportParamsFormatJson, MatterV1LogExportParamsFormatJSONL, MatterV1LogExportParamsFormatCsv, MatterV1LogExportParamsFormatTsv:
		return true
	}
	return false
}

// Filter by scope: matter, work_item, execution, sharing, all
//
// Satisfied by [shared.UnionString], [MatterV1LogExportParamsScopeArray].
type MatterV1LogExportParamsScopeUnion interface {
	ImplementsMatterV1LogExportParamsScopeUnion()
}

type MatterV1LogExportParamsScopeArray []string

func (r MatterV1LogExportParamsScopeArray) ImplementsMatterV1LogExportParamsScopeUnion() {}
