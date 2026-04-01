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
// MatterV1WorkItemService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1WorkItemService] method instead.
type MatterV1WorkItemService struct {
	Options []option.RequestOption
}

// NewMatterV1WorkItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatterV1WorkItemService(opts ...option.RequestOption) (r *MatterV1WorkItemService) {
	r = &MatterV1WorkItemService{}
	r.Options = opts
	return
}

// Create an active work item on a matter.
func (r *MatterV1WorkItemService) New(ctx context.Context, id string, body MatterV1WorkItemNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/work-items", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get a single work item for a matter.
func (r *MatterV1WorkItemService) Get(ctx context.Context, id string, workItemID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workItemID == "" {
		err = errors.New("missing required workItemId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/work-items/%s", id, workItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Update a matter work item.
func (r *MatterV1WorkItemService) Update(ctx context.Context, id string, workItemID string, body MatterV1WorkItemUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workItemID == "" {
		err = errors.New("missing required workItemId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/work-items/%s", id, workItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List active work items for a matter.
func (r *MatterV1WorkItemService) List(ctx context.Context, id string, query MatterV1WorkItemListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/work-items", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Approve, revise, block, or reassign a work item. Used by humans or agents to
// move work items through their lifecycle.
func (r *MatterV1WorkItemService) Decide(ctx context.Context, id string, workItemID string, body MatterV1WorkItemDecideParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workItemID == "" {
		err = errors.New("missing required workItemId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/work-items/%s/decision", id, workItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// List execution attempts for a work item, including agent and run linkage.
func (r *MatterV1WorkItemService) ListExecutions(ctx context.Context, id string, workItemID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workItemID == "" {
		err = errors.New("missing required workItemId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/work-items/%s/executions", id, workItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type MatterV1WorkItemNewParams struct {
	Title        param.Field[string]                            `json:"title" api:"required"`
	AssigneeID   param.Field[string]                            `json:"assignee_id"`
	Description  param.Field[string]                            `json:"description"`
	DueAt        param.Field[time.Time]                         `json:"due_at" format:"date-time"`
	ExitCriteria param.Field[[]string]                          `json:"exit_criteria"`
	Instructions param.Field[string]                            `json:"instructions"`
	Metadata     param.Field[map[string]interface{}]            `json:"metadata"`
	Priority     param.Field[MatterV1WorkItemNewParamsPriority] `json:"priority"`
	Type         param.Field[MatterV1WorkItemNewParamsType]     `json:"type"`
}

func (r MatterV1WorkItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1WorkItemNewParamsPriority string

const (
	MatterV1WorkItemNewParamsPriorityLow    MatterV1WorkItemNewParamsPriority = "low"
	MatterV1WorkItemNewParamsPriorityNormal MatterV1WorkItemNewParamsPriority = "normal"
	MatterV1WorkItemNewParamsPriorityHigh   MatterV1WorkItemNewParamsPriority = "high"
	MatterV1WorkItemNewParamsPriorityUrgent MatterV1WorkItemNewParamsPriority = "urgent"
)

func (r MatterV1WorkItemNewParamsPriority) IsKnown() bool {
	switch r {
	case MatterV1WorkItemNewParamsPriorityLow, MatterV1WorkItemNewParamsPriorityNormal, MatterV1WorkItemNewParamsPriorityHigh, MatterV1WorkItemNewParamsPriorityUrgent:
		return true
	}
	return false
}

type MatterV1WorkItemNewParamsType string

const (
	MatterV1WorkItemNewParamsTypeTask          MatterV1WorkItemNewParamsType = "task"
	MatterV1WorkItemNewParamsTypeDeadline      MatterV1WorkItemNewParamsType = "deadline"
	MatterV1WorkItemNewParamsTypeReview        MatterV1WorkItemNewParamsType = "review"
	MatterV1WorkItemNewParamsTypeFiling        MatterV1WorkItemNewParamsType = "filing"
	MatterV1WorkItemNewParamsTypeCommunication MatterV1WorkItemNewParamsType = "communication"
	MatterV1WorkItemNewParamsTypeResearch      MatterV1WorkItemNewParamsType = "research"
	MatterV1WorkItemNewParamsTypeDrafting      MatterV1WorkItemNewParamsType = "drafting"
	MatterV1WorkItemNewParamsTypeCollection    MatterV1WorkItemNewParamsType = "collection"
	MatterV1WorkItemNewParamsTypeIntake        MatterV1WorkItemNewParamsType = "intake"
)

func (r MatterV1WorkItemNewParamsType) IsKnown() bool {
	switch r {
	case MatterV1WorkItemNewParamsTypeTask, MatterV1WorkItemNewParamsTypeDeadline, MatterV1WorkItemNewParamsTypeReview, MatterV1WorkItemNewParamsTypeFiling, MatterV1WorkItemNewParamsTypeCommunication, MatterV1WorkItemNewParamsTypeResearch, MatterV1WorkItemNewParamsTypeDrafting, MatterV1WorkItemNewParamsTypeCollection, MatterV1WorkItemNewParamsTypeIntake:
		return true
	}
	return false
}

type MatterV1WorkItemUpdateParams struct {
	AssigneeID   param.Field[string]                               `json:"assignee_id"`
	CompletedAt  param.Field[time.Time]                            `json:"completed_at" format:"date-time"`
	Description  param.Field[string]                               `json:"description"`
	DueAt        param.Field[time.Time]                            `json:"due_at" format:"date-time"`
	ExitCriteria param.Field[[]string]                             `json:"exit_criteria"`
	Instructions param.Field[string]                               `json:"instructions"`
	Metadata     param.Field[map[string]interface{}]               `json:"metadata"`
	Priority     param.Field[MatterV1WorkItemUpdateParamsPriority] `json:"priority"`
	StartedAt    param.Field[time.Time]                            `json:"started_at" format:"date-time"`
	Status       param.Field[MatterV1WorkItemUpdateParamsStatus]   `json:"status"`
	Title        param.Field[string]                               `json:"title"`
	Type         param.Field[MatterV1WorkItemUpdateParamsType]     `json:"type"`
}

func (r MatterV1WorkItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1WorkItemUpdateParamsPriority string

const (
	MatterV1WorkItemUpdateParamsPriorityLow    MatterV1WorkItemUpdateParamsPriority = "low"
	MatterV1WorkItemUpdateParamsPriorityNormal MatterV1WorkItemUpdateParamsPriority = "normal"
	MatterV1WorkItemUpdateParamsPriorityHigh   MatterV1WorkItemUpdateParamsPriority = "high"
	MatterV1WorkItemUpdateParamsPriorityUrgent MatterV1WorkItemUpdateParamsPriority = "urgent"
)

func (r MatterV1WorkItemUpdateParamsPriority) IsKnown() bool {
	switch r {
	case MatterV1WorkItemUpdateParamsPriorityLow, MatterV1WorkItemUpdateParamsPriorityNormal, MatterV1WorkItemUpdateParamsPriorityHigh, MatterV1WorkItemUpdateParamsPriorityUrgent:
		return true
	}
	return false
}

type MatterV1WorkItemUpdateParamsStatus string

const (
	MatterV1WorkItemUpdateParamsStatusDraft         MatterV1WorkItemUpdateParamsStatus = "draft"
	MatterV1WorkItemUpdateParamsStatusQueued        MatterV1WorkItemUpdateParamsStatus = "queued"
	MatterV1WorkItemUpdateParamsStatusInProgress    MatterV1WorkItemUpdateParamsStatus = "in_progress"
	MatterV1WorkItemUpdateParamsStatusBlocked       MatterV1WorkItemUpdateParamsStatus = "blocked"
	MatterV1WorkItemUpdateParamsStatusInReview      MatterV1WorkItemUpdateParamsStatus = "in_review"
	MatterV1WorkItemUpdateParamsStatusAwaitingHuman MatterV1WorkItemUpdateParamsStatus = "awaiting_human"
	MatterV1WorkItemUpdateParamsStatusDone          MatterV1WorkItemUpdateParamsStatus = "done"
	MatterV1WorkItemUpdateParamsStatusCanceled      MatterV1WorkItemUpdateParamsStatus = "canceled"
)

func (r MatterV1WorkItemUpdateParamsStatus) IsKnown() bool {
	switch r {
	case MatterV1WorkItemUpdateParamsStatusDraft, MatterV1WorkItemUpdateParamsStatusQueued, MatterV1WorkItemUpdateParamsStatusInProgress, MatterV1WorkItemUpdateParamsStatusBlocked, MatterV1WorkItemUpdateParamsStatusInReview, MatterV1WorkItemUpdateParamsStatusAwaitingHuman, MatterV1WorkItemUpdateParamsStatusDone, MatterV1WorkItemUpdateParamsStatusCanceled:
		return true
	}
	return false
}

type MatterV1WorkItemUpdateParamsType string

const (
	MatterV1WorkItemUpdateParamsTypeTask          MatterV1WorkItemUpdateParamsType = "task"
	MatterV1WorkItemUpdateParamsTypeDeadline      MatterV1WorkItemUpdateParamsType = "deadline"
	MatterV1WorkItemUpdateParamsTypeReview        MatterV1WorkItemUpdateParamsType = "review"
	MatterV1WorkItemUpdateParamsTypeFiling        MatterV1WorkItemUpdateParamsType = "filing"
	MatterV1WorkItemUpdateParamsTypeCommunication MatterV1WorkItemUpdateParamsType = "communication"
	MatterV1WorkItemUpdateParamsTypeResearch      MatterV1WorkItemUpdateParamsType = "research"
	MatterV1WorkItemUpdateParamsTypeDrafting      MatterV1WorkItemUpdateParamsType = "drafting"
	MatterV1WorkItemUpdateParamsTypeCollection    MatterV1WorkItemUpdateParamsType = "collection"
	MatterV1WorkItemUpdateParamsTypeIntake        MatterV1WorkItemUpdateParamsType = "intake"
)

func (r MatterV1WorkItemUpdateParamsType) IsKnown() bool {
	switch r {
	case MatterV1WorkItemUpdateParamsTypeTask, MatterV1WorkItemUpdateParamsTypeDeadline, MatterV1WorkItemUpdateParamsTypeReview, MatterV1WorkItemUpdateParamsTypeFiling, MatterV1WorkItemUpdateParamsTypeCommunication, MatterV1WorkItemUpdateParamsTypeResearch, MatterV1WorkItemUpdateParamsTypeDrafting, MatterV1WorkItemUpdateParamsTypeCollection, MatterV1WorkItemUpdateParamsTypeIntake:
		return true
	}
	return false
}

type MatterV1WorkItemListParams struct {
	AssigneeID param.Field[string] `query:"assignee_id"`
	Status     param.Field[string] `query:"status"`
}

// URLQuery serializes [MatterV1WorkItemListParams]'s query parameters as
// `url.Values`.
func (r MatterV1WorkItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MatterV1WorkItemDecideParams struct {
	Decision    param.Field[MatterV1WorkItemDecideParamsDecision] `json:"decision" api:"required"`
	AgentTypeID param.Field[string]                               `json:"agent_type_id"`
	Metadata    param.Field[map[string]interface{}]               `json:"metadata"`
	Reason      param.Field[string]                               `json:"reason"`
}

func (r MatterV1WorkItemDecideParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1WorkItemDecideParamsDecision string

const (
	MatterV1WorkItemDecideParamsDecisionApprove  MatterV1WorkItemDecideParamsDecision = "approve"
	MatterV1WorkItemDecideParamsDecisionRevise   MatterV1WorkItemDecideParamsDecision = "revise"
	MatterV1WorkItemDecideParamsDecisionBlock    MatterV1WorkItemDecideParamsDecision = "block"
	MatterV1WorkItemDecideParamsDecisionReassign MatterV1WorkItemDecideParamsDecision = "reassign"
)

func (r MatterV1WorkItemDecideParamsDecision) IsKnown() bool {
	switch r {
	case MatterV1WorkItemDecideParamsDecisionApprove, MatterV1WorkItemDecideParamsDecisionRevise, MatterV1WorkItemDecideParamsDecisionBlock, MatterV1WorkItemDecideParamsDecisionReassign:
		return true
	}
	return false
}
