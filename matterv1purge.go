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
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1PurgeService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1PurgeService] method instead.
type MatterV1PurgeService struct {
	Options []option.RequestOption
}

// NewMatterV1PurgeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1PurgeService(opts ...option.RequestOption) (r *MatterV1PurgeService) {
	r = &MatterV1PurgeService{}
	r.Options = opts
	return
}

// Returns a Matter purge receipt for manual diagnostics. Integrations should
// subscribe to Matter purge webhooks rather than polling this route.
func (r *MatterV1PurgeService) Get(ctx context.Context, purgeID string, opts ...option.RequestOption) (res *MatterV1PurgeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if purgeID == "" {
		err = errors.New("missing required purgeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("matters/v1/purges/%s", purgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MatterV1PurgeGetResponse struct {
	Attempt     int64                          `json:"attempt" api:"required"`
	MatterID    string                         `json:"matter_id" api:"required"`
	PurgeID     string                         `json:"purge_id" api:"required"`
	Status      MatterV1PurgeGetResponseStatus `json:"status" api:"required"`
	VaultID     string                         `json:"vault_id" api:"required"`
	WorkflowID  string                         `json:"workflow_id" api:"required,nullable"`
	CompletedAt time.Time                      `json:"completed_at" api:"nullable" format:"date-time"`
	Counts      MatterV1PurgeGetResponseCounts `json:"counts"`
	FailedAt    time.Time                      `json:"failed_at" api:"nullable" format:"date-time"`
	FailureCode string                         `json:"failure_code" api:"nullable"`
	RequestedAt time.Time                      `json:"requested_at" format:"date-time"`
	StartedAt   time.Time                      `json:"started_at" api:"nullable" format:"date-time"`
	// Stable ID of the failed or completed terminal webhook event
	TerminalEventID string                       `json:"terminal_event_id" api:"nullable"`
	JSON            matterV1PurgeGetResponseJSON `json:"-"`
}

// matterV1PurgeGetResponseJSON contains the JSON metadata for the struct
// [MatterV1PurgeGetResponse]
type matterV1PurgeGetResponseJSON struct {
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

func (r *MatterV1PurgeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1PurgeGetResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1PurgeGetResponseStatus string

const (
	MatterV1PurgeGetResponseStatusQueued     MatterV1PurgeGetResponseStatus = "queued"
	MatterV1PurgeGetResponseStatusInProgress MatterV1PurgeGetResponseStatus = "in_progress"
	MatterV1PurgeGetResponseStatusFailed     MatterV1PurgeGetResponseStatus = "failed"
	MatterV1PurgeGetResponseStatusCompleted  MatterV1PurgeGetResponseStatus = "completed"
)

func (r MatterV1PurgeGetResponseStatus) IsKnown() bool {
	switch r {
	case MatterV1PurgeGetResponseStatusQueued, MatterV1PurgeGetResponseStatusInProgress, MatterV1PurgeGetResponseStatusFailed, MatterV1PurgeGetResponseStatusCompleted:
		return true
	}
	return false
}

type MatterV1PurgeGetResponseCounts struct {
	Chats          int64                              `json:"chats"`
	Objects        int64                              `json:"objects"`
	Sessions       int64                              `json:"sessions"`
	Transcriptions int64                              `json:"transcriptions"`
	JSON           matterV1PurgeGetResponseCountsJSON `json:"-"`
}

// matterV1PurgeGetResponseCountsJSON contains the JSON metadata for the struct
// [MatterV1PurgeGetResponseCounts]
type matterV1PurgeGetResponseCountsJSON struct {
	Chats          apijson.Field
	Objects        apijson.Field
	Sessions       apijson.Field
	Transcriptions apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MatterV1PurgeGetResponseCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1PurgeGetResponseCountsJSON) RawJSON() string {
	return r.raw
}
