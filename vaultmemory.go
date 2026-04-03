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

// Secure document storage with semantic search and GraphRAG
//
// VaultMemoryService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultMemoryService] method instead.
type VaultMemoryService struct {
	Options []option.RequestOption
}

// NewVaultMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultMemoryService(opts ...option.RequestOption) (r *VaultMemoryService) {
	r = &VaultMemoryService{}
	r.Options = opts
	return
}

// Append a new file-backed memory entry to a vault.
func (r *VaultMemoryService) New(ctx context.Context, id string, body VaultMemoryNewParams, opts ...option.RequestOption) (res *VaultMemoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/memory", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Rewrite a file-backed vault memory entry with updated content, source, or tags.
func (r *VaultMemoryService) Update(ctx context.Context, id string, entryID string, body VaultMemoryUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if entryID == "" {
		err = errors.New("missing required entryId parameter")
		return err
	}
	path := fmt.Sprintf("vault/%s/memory/%s", id, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Retrieve file-backed memory entries stored in a vault.
func (r *VaultMemoryService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *VaultMemoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/memory", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove a file-backed memory entry from a vault.
func (r *VaultMemoryService) Delete(ctx context.Context, id string, entryID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if entryID == "" {
		err = errors.New("missing required entryId parameter")
		return err
	}
	path := fmt.Sprintf("vault/%s/memory/%s", id, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Search file-backed vault memory using simple full-text matching over content and
// tags.
func (r *VaultMemoryService) Search(ctx context.Context, id string, body VaultMemorySearchParams, opts ...option.RequestOption) (res *VaultMemorySearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/memory/search", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type VaultMemoryNewResponse struct {
	Entry VaultMemoryNewResponseEntry `json:"entry"`
	JSON  vaultMemoryNewResponseJSON  `json:"-"`
}

// vaultMemoryNewResponseJSON contains the JSON metadata for the struct
// [VaultMemoryNewResponse]
type vaultMemoryNewResponseJSON struct {
	Entry       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemoryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemoryNewResponseJSON) RawJSON() string {
	return r.raw
}

type VaultMemoryNewResponseEntry struct {
	ID        string                          `json:"id"`
	Content   string                          `json:"content"`
	CreatedAt time.Time                       `json:"created_at" format:"date-time"`
	CreatedBy string                          `json:"created_by" api:"nullable"`
	Source    string                          `json:"source" api:"nullable"`
	Tags      []string                        `json:"tags"`
	Type      string                          `json:"type"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      vaultMemoryNewResponseEntryJSON `json:"-"`
}

// vaultMemoryNewResponseEntryJSON contains the JSON metadata for the struct
// [VaultMemoryNewResponseEntry]
type vaultMemoryNewResponseEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedAt   apijson.Field
	CreatedBy   apijson.Field
	Source      apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemoryNewResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemoryNewResponseEntryJSON) RawJSON() string {
	return r.raw
}

type VaultMemoryListResponse struct {
	Entries []VaultMemoryListResponseEntry `json:"entries"`
	Meta    VaultMemoryListResponseMeta    `json:"meta"`
	JSON    vaultMemoryListResponseJSON    `json:"-"`
}

// vaultMemoryListResponseJSON contains the JSON metadata for the struct
// [VaultMemoryListResponse]
type vaultMemoryListResponseJSON struct {
	Entries     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemoryListResponseJSON) RawJSON() string {
	return r.raw
}

type VaultMemoryListResponseEntry struct {
	ID        string                           `json:"id"`
	Content   string                           `json:"content"`
	CreatedAt time.Time                        `json:"created_at" format:"date-time"`
	CreatedBy string                           `json:"created_by" api:"nullable"`
	Source    string                           `json:"source" api:"nullable"`
	Tags      []string                         `json:"tags"`
	Type      string                           `json:"type"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      vaultMemoryListResponseEntryJSON `json:"-"`
}

// vaultMemoryListResponseEntryJSON contains the JSON metadata for the struct
// [VaultMemoryListResponseEntry]
type vaultMemoryListResponseEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedAt   apijson.Field
	CreatedBy   apijson.Field
	Source      apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemoryListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemoryListResponseEntryJSON) RawJSON() string {
	return r.raw
}

type VaultMemoryListResponseMeta struct {
	Chars     int64                           `json:"chars"`
	Count     int64                           `json:"count"`
	MaxChars  int64                           `json:"max_chars"`
	UpdatedAt time.Time                       `json:"updated_at" api:"nullable" format:"date-time"`
	JSON      vaultMemoryListResponseMetaJSON `json:"-"`
}

// vaultMemoryListResponseMetaJSON contains the JSON metadata for the struct
// [VaultMemoryListResponseMeta]
type vaultMemoryListResponseMetaJSON struct {
	Chars       apijson.Field
	Count       apijson.Field
	MaxChars    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemoryListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemoryListResponseMetaJSON) RawJSON() string {
	return r.raw
}

type VaultMemorySearchResponse struct {
	Results []VaultMemorySearchResponseResult `json:"results"`
	JSON    vaultMemorySearchResponseJSON     `json:"-"`
}

// vaultMemorySearchResponseJSON contains the JSON metadata for the struct
// [VaultMemorySearchResponse]
type vaultMemorySearchResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemorySearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemorySearchResponseJSON) RawJSON() string {
	return r.raw
}

type VaultMemorySearchResponseResult struct {
	ID        string                              `json:"id"`
	Content   string                              `json:"content"`
	CreatedAt time.Time                           `json:"created_at" format:"date-time"`
	CreatedBy string                              `json:"created_by" api:"nullable"`
	Source    string                              `json:"source" api:"nullable"`
	Tags      []string                            `json:"tags"`
	Type      string                              `json:"type"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      vaultMemorySearchResponseResultJSON `json:"-"`
}

// vaultMemorySearchResponseResultJSON contains the JSON metadata for the struct
// [VaultMemorySearchResponseResult]
type vaultMemorySearchResponseResultJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedAt   apijson.Field
	CreatedBy   apijson.Field
	Source      apijson.Field
	Tags        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMemorySearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMemorySearchResponseResultJSON) RawJSON() string {
	return r.raw
}

type VaultMemoryNewParams struct {
	Content param.Field[string]                   `json:"content" api:"required"`
	Type    param.Field[VaultMemoryNewParamsType] `json:"type" api:"required"`
	Source  param.Field[string]                   `json:"source"`
	Tags    param.Field[[]string]                 `json:"tags"`
}

func (r VaultMemoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMemoryNewParamsType string

const (
	VaultMemoryNewParamsTypeFact       VaultMemoryNewParamsType = "fact"
	VaultMemoryNewParamsTypeParty      VaultMemoryNewParamsType = "party"
	VaultMemoryNewParamsTypeIssue      VaultMemoryNewParamsType = "issue"
	VaultMemoryNewParamsTypeDeadline   VaultMemoryNewParamsType = "deadline"
	VaultMemoryNewParamsTypeDiscovery  VaultMemoryNewParamsType = "discovery"
	VaultMemoryNewParamsTypeCorrection VaultMemoryNewParamsType = "correction"
	VaultMemoryNewParamsTypePreference VaultMemoryNewParamsType = "preference"
)

func (r VaultMemoryNewParamsType) IsKnown() bool {
	switch r {
	case VaultMemoryNewParamsTypeFact, VaultMemoryNewParamsTypeParty, VaultMemoryNewParamsTypeIssue, VaultMemoryNewParamsTypeDeadline, VaultMemoryNewParamsTypeDiscovery, VaultMemoryNewParamsTypeCorrection, VaultMemoryNewParamsTypePreference:
		return true
	}
	return false
}

type VaultMemoryUpdateParams struct {
	Content param.Field[string]   `json:"content"`
	Source  param.Field[string]   `json:"source"`
	Tags    param.Field[[]string] `json:"tags"`
}

func (r VaultMemoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMemorySearchParams struct {
	Query param.Field[string]   `json:"query" api:"required"`
	Limit param.Field[int64]    `json:"limit"`
	Tags  param.Field[[]string] `json:"tags"`
	Types param.Field[[]string] `json:"types"`
}

func (r VaultMemorySearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
