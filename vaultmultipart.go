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
		return err
	}
	path := fmt.Sprintf("vault/%s/multipart/abort", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Complete a multipart upload by providing the list of part numbers and ETags
// (live). Single PUT uploads are capped at 5GB; multipart default max is 16GB
// (configurable).
func (r *VaultMultipartService) Complete(ctx context.Context, id string, body VaultMultipartCompleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("vault/%s/multipart/complete", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Generate presigned URLs for individual multipart upload parts (live).
func (r *VaultMultipartService) GetPartURLs(ctx context.Context, id string, body VaultMultipartGetPartURLsParams, opts ...option.RequestOption) (res *VaultMultipartGetPartURLsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/multipart/part-urls", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Initiate a multipart upload for large files (>5GB). Single PUT uploads are
// capped at 5GB; multipart default max is 16GB (configurable). Multipart uploads
// are supported in production. Returns an uploadId and object metadata. Use part
// URLs endpoint to upload parts and complete endpoint to finalize.
func (r *VaultMultipartService) Init(ctx context.Context, id string, body VaultMultipartInitParams, opts ...option.RequestOption) (res *VaultMultipartInitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/multipart/init", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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

type VaultMultipartInitResponse struct {
	NextStep      string                         `json:"next_step"`
	ObjectID      string                         `json:"objectId"`
	PartCount     int64                          `json:"partCount"`
	PartSizeBytes int64                          `json:"partSizeBytes"`
	S3Key         string                         `json:"s3Key"`
	UploadID      string                         `json:"uploadId"`
	JSON          vaultMultipartInitResponseJSON `json:"-"`
}

// vaultMultipartInitResponseJSON contains the JSON metadata for the struct
// [VaultMultipartInitResponse]
type vaultMultipartInitResponseJSON struct {
	NextStep      apijson.Field
	ObjectID      apijson.Field
	PartCount     apijson.Field
	PartSizeBytes apijson.Field
	S3Key         apijson.Field
	UploadID      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultMultipartInitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultMultipartInitResponseJSON) RawJSON() string {
	return r.raw
}

type VaultMultipartAbortParams struct {
	// Vault object ID associated with the multipart upload
	ObjectID param.Field[string] `json:"objectId" api:"required"`
	// Multipart upload ID returned when the upload was initialized
	UploadID param.Field[string] `json:"uploadId" api:"required"`
}

func (r VaultMultipartAbortParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartCompleteParams struct {
	ObjectID param.Field[string]                             `json:"objectId" api:"required"`
	Parts    param.Field[[]VaultMultipartCompleteParamsPart] `json:"parts" api:"required"`
	// File size in bytes (default max 16GB). Configure via
	// VAULT_MULTIPART_MAX_FILE_SIZE_BYTES.
	SizeBytes param.Field[int64]  `json:"sizeBytes" api:"required"`
	UploadID  param.Field[string] `json:"uploadId" api:"required"`
}

func (r VaultMultipartCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartCompleteParamsPart struct {
	Etag       param.Field[string] `json:"etag" api:"required"`
	PartNumber param.Field[int64]  `json:"partNumber" api:"required"`
}

func (r VaultMultipartCompleteParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartGetPartURLsParams struct {
	// Vault object ID associated with the multipart upload
	ObjectID param.Field[string] `json:"objectId" api:"required"`
	// Multipart parts that need presigned upload URLs
	Parts param.Field[[]VaultMultipartGetPartURLsParamsPart] `json:"parts" api:"required"`
	// Multipart upload ID returned when the upload was initialized
	UploadID param.Field[string] `json:"uploadId" api:"required"`
}

func (r VaultMultipartGetPartURLsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartGetPartURLsParamsPart struct {
	// 1-based multipart part number
	PartNumber param.Field[int64] `json:"partNumber" api:"required"`
	// Part size in bytes (min 5MB except final part, max 5GB).
	SizeBytes param.Field[int64] `json:"sizeBytes" api:"required"`
}

func (r VaultMultipartGetPartURLsParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultMultipartInitParams struct {
	// MIME type of the file
	ContentType param.Field[string] `json:"contentType" api:"required"`
	// Name of the file to upload
	Filename param.Field[string] `json:"filename" api:"required"`
	// File size in bytes (required, default max 16GB). Configure via
	// VAULT_MULTIPART_MAX_FILE_SIZE_BYTES.
	SizeBytes param.Field[int64] `json:"sizeBytes" api:"required"`
	// Whether to automatically process and index the file for search
	AutoIndex param.Field[bool] `json:"auto_index"`
	// Marks the file as AI-generated work product (e.g. uploaded by an agent) rather
	// than a user-provided source document. Persisted on the object and returned by
	// object listings so clients can distinguish provenance.
	IsAIGenerated param.Field[bool] `json:"is_ai_generated"`
	// Additional metadata to associate with the file
	Metadata param.Field[interface{}] `json:"metadata"`
	// Multipart part size in bytes (min 5MB, max 5GB). Defaults to 64MB.
	PartSizeBytes param.Field[int64] `json:"partSizeBytes"`
	// Optional folder path for hierarchy preservation
	Path param.Field[string] `json:"path"`
}

func (r VaultMultipartInitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
