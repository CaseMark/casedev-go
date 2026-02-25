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

// VaultMultipartService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultMultipartService] method instead.
type VaultMultipartService struct {
	Options []option.RequestOption
}

// NewVaultMultipartService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultMultipartService(opts ...option.RequestOption) (r *VaultMultipartService) {
	r = &VaultMultipartService{}
	r.Options = opts
	return
}

// Abort a multipart upload and discard uploaded parts (live).
func (r *VaultMultipartService) Abort(ctx context.Context, id string, body VaultMultipartAbortParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/multipart/abort", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Generate presigned URLs for individual multipart upload parts (live).
func (r *VaultMultipartService) GetPartURLs(ctx context.Context, id string, body VaultMultipartGetPartURLsParams, opts ...option.RequestOption) (res *VaultMultipartGetPartURLsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/multipart/part-urls", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VaultMultipartGetPartURLsResponse struct {
	URLs []VaultMultipartGetPartURLsResponseURL `json:"urls"`
	JSON vaultMultipartGetPartURLsResponseJSON  `json:"-"`
}

// vaultMultipartGetPartURLsResponseJSON contains the JSON metadata for the struct
// [VaultMultipartGetPartURLsResponse]
type vaultMultipartGetPartURLsResponseJSON struct {
	URLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMultipartGetPartURLsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMultipartGetPartURLsResponseJSON) RawJSON() string {
	return r.raw
}

type VaultMultipartGetPartURLsResponseURL struct {
	PartNumber int64                                    `json:"partNumber"`
	URL        string                                   `json:"url"`
	JSON       vaultMultipartGetPartURLsResponseURLJSON `json:"-"`
}

// vaultMultipartGetPartURLsResponseURLJSON contains the JSON metadata for the
// struct [VaultMultipartGetPartURLsResponseURL]
type vaultMultipartGetPartURLsResponseURLJSON struct {
	PartNumber  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultMultipartGetPartURLsResponseURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMultipartGetPartURLsResponseURLJSON) RawJSON() string {
	return r.raw
}

type VaultMultipartAbortParams struct {
	ObjectID param.Field[string] `json:"objectId" api:"required"`
	UploadID param.Field[string] `json:"uploadId" api:"required"`
}

func (r VaultMultipartAbortParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartGetPartURLsParams struct {
	ObjectID param.Field[string]                                `json:"objectId" api:"required"`
	Parts    param.Field[[]VaultMultipartGetPartURLsParamsPart] `json:"parts" api:"required"`
	UploadID param.Field[string]                                `json:"uploadId" api:"required"`
}

func (r VaultMultipartGetPartURLsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartGetPartURLsParamsPart struct {
	PartNumber param.Field[int64] `json:"partNumber" api:"required"`
	// Part size in bytes (min 5MB except final part, max 5GB).
	SizeBytes param.Field[int64] `json:"sizeBytes" api:"required"`
}

func (r VaultMultipartGetPartURLsParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
