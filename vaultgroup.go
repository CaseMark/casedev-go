// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Secure document storage with semantic search and GraphRAG
//
// VaultGroupService contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultGroupService] method instead.
type VaultGroupService struct {
	Options []option.RequestOption
}

// NewVaultGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultGroupService(opts ...option.RequestOption) (r *VaultGroupService) {
	r = &VaultGroupService{}
	r.Options = opts
	return
}

// Creates a vault group for organizing vaults and applying group-scoped access
// controls. Group-scoped API keys cannot create or manage vault groups.
func (r *VaultGroupService) New(ctx context.Context, body VaultGroupNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "vault/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Updates a vault group for the authenticated organization. Only provided fields
// are changed, and setting description to null removes the current description.
func (r *VaultGroupService) Update(ctx context.Context, groupID string, body VaultGroupUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return err
	}
	path := fmt.Sprintf("vault/groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Lists vault groups visible to the authenticated organization. Group-scoped API
// keys only receive groups within their allowed scope.
func (r *VaultGroupService) List(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "vault/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Soft-deletes a vault group that no longer has any active vaults assigned. This
// operation is blocked when the group still contains vaults.
func (r *VaultGroupService) Delete(ctx context.Context, groupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return err
	}
	path := fmt.Sprintf("vault/groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type VaultGroupNewParams struct {
	// Human-readable name for the vault group
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description of the vault group purpose
	Description param.Field[string] `json:"description"`
}

func (r VaultGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultGroupUpdateParams struct {
	// Updated vault group description. Pass null to remove the current description.
	Description param.Field[string] `json:"description"`
	// New human-readable name for the vault group
	Name param.Field[string] `json:"name"`
}

func (r VaultGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
