// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1ContentPurgeService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1ContentPurgeService] method instead.
type MatterV1ContentPurgeService struct {
	Options []option.RequestOption
}

// NewMatterV1ContentPurgeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatterV1ContentPurgeService(opts ...option.RequestOption) (r *MatterV1ContentPurgeService) {
	r = &MatterV1ContentPurgeService{}
	r.Options = opts
	return
}

// Queues an idempotent hard deletion of explicitly owned content while preserving
// the Matter, Vault, and unrelated content. Use unified matter.content_purge
// webhooks for completion, not polling.
func (r *MatterV1ContentPurgeService) New(ctx context.Context, id string, body MatterV1ContentPurgeNewParams, opts ...option.RequestOption) (res *MatterV1ContentPurgeNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("matters/v1/%s/content-purges", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Owner-only receipt for operator diagnostics. Integrations must use unified
// content-purge webhooks rather than polling.
func (r *MatterV1ContentPurgeService) Get(ctx context.Context, purgeID string, opts ...option.RequestOption) (res *MatterV1ContentPurgeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if purgeID == "" {
		err = errors.New("missing required purgeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("matters/v1/content-purges/%s", purgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MatterV1ContentPurgeNewResponse struct {
	Attempt         int64                                 `json:"attempt" api:"required"`
	MatterID        string                                `json:"matter_id" api:"required"`
	PurgeID         string                                `json:"purge_id" api:"required"`
	RequestID       string                                `json:"request_id" api:"required"`
	Status          MatterV1ContentPurgeNewResponseStatus `json:"status" api:"required"`
	VaultID         string                                `json:"vault_id" api:"required"`
	WorkflowID      string                                `json:"workflow_id" api:"required,nullable"`
	CompletedAt     time.Time                             `json:"completed_at" api:"nullable" format:"date-time"`
	Counts          map[string]int64                      `json:"counts"`
	FailedAt        time.Time                             `json:"failed_at" api:"nullable" format:"date-time"`
	FailureCode     string                                `json:"failure_code" api:"nullable"`
	RequestedAt     time.Time                             `json:"requested_at" format:"date-time"`
	StartedAt       time.Time                             `json:"started_at" api:"nullable" format:"date-time"`
	TerminalEventID string                                `json:"terminal_event_id" api:"nullable"`
	JSON            matterV1ContentPurgeNewResponseJSON   `json:"-"`
}

// matterV1ContentPurgeNewResponseJSON contains the JSON metadata for the struct
// [MatterV1ContentPurgeNewResponse]
type matterV1ContentPurgeNewResponseJSON struct {
	Attempt         apijson.Field
	MatterID        apijson.Field
	PurgeID         apijson.Field
	RequestID       apijson.Field
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

func (r *MatterV1ContentPurgeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1ContentPurgeNewResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1ContentPurgeNewResponseStatus string

const (
	MatterV1ContentPurgeNewResponseStatusQueued     MatterV1ContentPurgeNewResponseStatus = "queued"
	MatterV1ContentPurgeNewResponseStatusInProgress MatterV1ContentPurgeNewResponseStatus = "in_progress"
	MatterV1ContentPurgeNewResponseStatusFailed     MatterV1ContentPurgeNewResponseStatus = "failed"
	MatterV1ContentPurgeNewResponseStatusCompleted  MatterV1ContentPurgeNewResponseStatus = "completed"
)

func (r MatterV1ContentPurgeNewResponseStatus) IsKnown() bool {
	switch r {
	case MatterV1ContentPurgeNewResponseStatusQueued, MatterV1ContentPurgeNewResponseStatusInProgress, MatterV1ContentPurgeNewResponseStatusFailed, MatterV1ContentPurgeNewResponseStatusCompleted:
		return true
	}
	return false
}

type MatterV1ContentPurgeGetResponse struct {
	Attempt         int64                                 `json:"attempt" api:"required"`
	MatterID        string                                `json:"matter_id" api:"required"`
	PurgeID         string                                `json:"purge_id" api:"required"`
	RequestID       string                                `json:"request_id" api:"required"`
	Status          MatterV1ContentPurgeGetResponseStatus `json:"status" api:"required"`
	VaultID         string                                `json:"vault_id" api:"required"`
	WorkflowID      string                                `json:"workflow_id" api:"required,nullable"`
	CompletedAt     time.Time                             `json:"completed_at" api:"nullable" format:"date-time"`
	Counts          map[string]int64                      `json:"counts"`
	FailedAt        time.Time                             `json:"failed_at" api:"nullable" format:"date-time"`
	FailureCode     string                                `json:"failure_code" api:"nullable"`
	RequestedAt     time.Time                             `json:"requested_at" format:"date-time"`
	StartedAt       time.Time                             `json:"started_at" api:"nullable" format:"date-time"`
	TerminalEventID string                                `json:"terminal_event_id" api:"nullable"`
	JSON            matterV1ContentPurgeGetResponseJSON   `json:"-"`
}

// matterV1ContentPurgeGetResponseJSON contains the JSON metadata for the struct
// [MatterV1ContentPurgeGetResponse]
type matterV1ContentPurgeGetResponseJSON struct {
	Attempt         apijson.Field
	MatterID        apijson.Field
	PurgeID         apijson.Field
	RequestID       apijson.Field
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

func (r *MatterV1ContentPurgeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1ContentPurgeGetResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1ContentPurgeGetResponseStatus string

const (
	MatterV1ContentPurgeGetResponseStatusQueued     MatterV1ContentPurgeGetResponseStatus = "queued"
	MatterV1ContentPurgeGetResponseStatusInProgress MatterV1ContentPurgeGetResponseStatus = "in_progress"
	MatterV1ContentPurgeGetResponseStatusFailed     MatterV1ContentPurgeGetResponseStatus = "failed"
	MatterV1ContentPurgeGetResponseStatusCompleted  MatterV1ContentPurgeGetResponseStatus = "completed"
)

func (r MatterV1ContentPurgeGetResponseStatus) IsKnown() bool {
	switch r {
	case MatterV1ContentPurgeGetResponseStatusQueued, MatterV1ContentPurgeGetResponseStatusInProgress, MatterV1ContentPurgeGetResponseStatusFailed, MatterV1ContentPurgeGetResponseStatusCompleted:
		return true
	}
	return false
}

type MatterV1ContentPurgeNewParams struct {
	// Stable caller idempotency ID; cannot be reused with different targets.
	RequestID        param.Field[string]   `json:"request_id" api:"required"`
	ObjectIDs        param.Field[[]string] `json:"object_ids"`
	SessionIDs       param.Field[[]string] `json:"session_ids"`
	TranscriptionIDs param.Field[[]string] `json:"transcription_ids"`
	WorkItemIDs      param.Field[[]string] `json:"work_item_ids"`
}

func (r MatterV1ContentPurgeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
