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
// MatterV1ShareService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1ShareService] method instead.
type MatterV1ShareService struct {
	Options []option.RequestOption
}

// NewMatterV1ShareService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1ShareService(opts ...option.RequestOption) (r *MatterV1ShareService) {
	r = &MatterV1ShareService{}
	r.Options = opts
	return
}

// Grant another organization scoped access to this matter and its primary vault.
func (r *MatterV1ShareService) New(ctx context.Context, id string, body MatterV1ShareNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/shares", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// List cross-org shares for a matter. Owner only.
func (r *MatterV1ShareService) List(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/shares", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Revoke a matter share and its linked vault share.
func (r *MatterV1ShareService) Delete(ctx context.Context, id string, shareID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if shareID == "" {
		err = errors.New("missing required shareId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/shares/%s", id, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type MatterV1ShareNewParams struct {
	TargetOrgID param.Field[string]                           `json:"target_org_id" api:"required"`
	ExpiresAt   param.Field[time.Time]                        `json:"expires_at" format:"date-time"`
	Permission  param.Field[MatterV1ShareNewParamsPermission] `json:"permission"`
}

func (r MatterV1ShareNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1ShareNewParamsPermission string

const (
	MatterV1ShareNewParamsPermissionRead MatterV1ShareNewParamsPermission = "read"
	MatterV1ShareNewParamsPermissionEdit MatterV1ShareNewParamsPermission = "edit"
)

func (r MatterV1ShareNewParamsPermission) IsKnown() bool {
	switch r {
	case MatterV1ShareNewParamsPermissionRead, MatterV1ShareNewParamsPermissionEdit:
		return true
	}
	return false
}
