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

// Secure document storage with semantic search
//
// VaultFolderService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultFolderService] method instead.
type VaultFolderService struct {
	Options []option.RequestOption
}

// NewVaultFolderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultFolderService(opts ...option.RequestOption) (r *VaultFolderService) {
	r = &VaultFolderService{}
	r.Options = opts
	return
}

// Creates a durable empty folder and missing parents. Metadata only; no S3 objects
// are created. Requires vault-wide write access and vault ownership.
func (r *VaultFolderService) New(ctx context.Context, id string, body VaultFolderNewParams, opts ...option.RequestOption) (res *VaultFolderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/folders", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists durable folder paths, including empty folders. Combine with paginated
// object paths to display imported folders. Requires vault-wide read access;
// shared vault reads are supported.
func (r *VaultFolderService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *VaultFolderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/folders", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Atomically renames or moves a folder, its empty descendants and all descendant
// object paths. Rejects an existing destination, root and self-descendant moves,
// and busy objects. Supports vaults with at most 25000 objects. Uses a 5-second
// lock timeout and 30-second statement timeout. Does not transfer file bytes.
// Requires vault-wide write access and vault ownership.
func (r *VaultFolderService) Move(ctx context.Context, id string, body VaultFolderMoveParams, opts ...option.RequestOption) (res *VaultFolderMoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/folders", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type VaultFolderNewResponse struct {
	FolderPaths []string                   `json:"folderPaths" api:"required"`
	Path        string                     `json:"path" api:"required"`
	JSON        vaultFolderNewResponseJSON `json:"-"`
}

// vaultFolderNewResponseJSON contains the JSON metadata for the struct
// [VaultFolderNewResponse]
type vaultFolderNewResponseJSON struct {
	FolderPaths apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultFolderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultFolderNewResponseJSON) RawJSON() string {
	return r.raw
}

type VaultFolderListResponse struct {
	FolderPaths []string                    `json:"folderPaths" api:"required"`
	JSON        vaultFolderListResponseJSON `json:"-"`
}

// vaultFolderListResponseJSON contains the JSON metadata for the struct
// [VaultFolderListResponse]
type vaultFolderListResponseJSON struct {
	FolderPaths apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultFolderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultFolderListResponseJSON) RawJSON() string {
	return r.raw
}

type VaultFolderMoveResponse struct {
	FolderPaths []string                    `json:"folderPaths"`
	MovedFiles  int64                       `json:"movedFiles"`
	NewPath     string                      `json:"newPath"`
	Path        string                      `json:"path"`
	Revision    string                      `json:"revision"`
	JSON        vaultFolderMoveResponseJSON `json:"-"`
}

// vaultFolderMoveResponseJSON contains the JSON metadata for the struct
// [VaultFolderMoveResponse]
type vaultFolderMoveResponseJSON struct {
	FolderPaths apijson.Field
	MovedFiles  apijson.Field
	NewPath     apijson.Field
	Path        apijson.Field
	Revision    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultFolderMoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultFolderMoveResponseJSON) RawJSON() string {
	return r.raw
}

type VaultFolderNewParams struct {
	Path param.Field[string] `json:"path" api:"required"`
}

func (r VaultFolderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultFolderMoveParams struct {
	NewPath param.Field[string] `json:"newPath" api:"required"`
	Path    param.Field[string] `json:"path" api:"required"`
	// Revision from a previous move, for safe Undo.
	ExpectedRevision param.Field[string] `json:"expectedRevision"`
}

func (r VaultFolderMoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
