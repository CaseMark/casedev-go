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

// VaultService contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultService] method instead.
type VaultService struct {
	Options   []option.RequestOption
	Events    *VaultEventService
	Graphrag  *VaultGraphragService
	Groups    *VaultGroupService
	Multipart *VaultMultipartService
	Objects   *VaultObjectService
}

// NewVaultService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVaultService(opts ...option.RequestOption) (r *VaultService) {
	r = &VaultService{}
	r.Options = opts
	r.Events = NewVaultEventService(opts...)
	r.Graphrag = NewVaultGraphragService(opts...)
	r.Groups = NewVaultGroupService(opts...)
	r.Multipart = NewVaultMultipartService(opts...)
	r.Objects = NewVaultObjectService(opts...)
	return
}

// Creates a new secure vault with dedicated S3 storage and vector search
// capabilities. Each vault provides isolated document storage with semantic
// search, OCR processing, and optional GraphRAG knowledge graph features for legal
// document analysis and discovery.
func (r *VaultService) New(ctx context.Context, body VaultNewParams, opts ...option.RequestOption) (res *VaultNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vault"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific vault, including storage
// configuration, chunking strategy, and usage statistics. Returns vault metadata,
// bucket information, and vector storage details.
func (r *VaultService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VaultGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update vault settings including name, description, and enableGraph. Changing
// enableGraph only affects future document uploads - existing documents retain
// their current graph/non-graph state.
func (r *VaultService) Update(ctx context.Context, id string, body VaultUpdateParams, opts ...option.RequestOption) (res *VaultUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all vaults for the authenticated organization. Returns vault metadata
// including name, description, storage configuration, and usage statistics.
func (r *VaultService) List(ctx context.Context, opts ...option.RequestOption) (res *VaultListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vault"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a vault and all its contents including documents, vectors,
// graph data, and S3 buckets. This operation cannot be undone. For large vaults,
// use the async=true query parameter to queue deletion in the background.
func (r *VaultService) Delete(ctx context.Context, id string, body VaultDeleteParams, opts ...option.RequestOption) (res *VaultDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Confirm whether a direct-to-S3 vault upload succeeded or failed. This endpoint
// emits vault.upload.completed or vault.upload.failed events and is idempotent for
// repeated confirmations.
func (r *VaultService) ConfirmUpload(ctx context.Context, id string, objectID string, body VaultConfirmUploadParams, opts ...option.RequestOption) (res *VaultConfirmUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/upload/%s/confirm", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Triggers ingestion workflow for a vault object to extract text, generate chunks,
// and create embeddings. For supported file types (PDF, DOCX, TXT, RTF, XML,
// audio, video), processing happens asynchronously. For unsupported types (images,
// archives, etc.), the file is marked as completed immediately without text
// extraction. GraphRAG indexing must be triggered separately via POST
// /vault/:id/graphrag/:objectId.
func (r *VaultService) Ingest(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *VaultIngestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/ingest/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Search across vault documents using multiple methods including hybrid vector +
// graph search, GraphRAG global search, entity-based search, and fast similarity
// search. Returns relevant documents and contextual answers based on the search
// method.
func (r *VaultService) Search(ctx context.Context, id string, body VaultSearchParams, opts ...option.RequestOption) (res *VaultSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/search", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a presigned URL for uploading files directly to a vault's S3 storage.
// After uploading to S3, confirm the upload result via POST
// /vault/:vaultId/upload/:objectId/confirm before triggering ingestion.
func (r *VaultService) Upload(ctx context.Context, id string, body VaultUploadParams, opts ...option.RequestOption) (res *VaultUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/upload", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VaultNewResponse struct {
	// Unique vault identifier
	ID string `json:"id"`
	// Vault creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Vault description
	Description string `json:"description"`
	// Whether vector indexing is enabled for this vault
	EnableIndexing bool `json:"enableIndexing"`
	// S3 bucket name for document storage
	FilesBucket string `json:"filesBucket"`
	// Vector search index name. Null for storage-only vaults.
	IndexName string `json:"indexName" api:"nullable"`
	// Vault display name
	Name string `json:"name"`
	// AWS region for storage
	Region string `json:"region"`
	// S3 bucket name for vector embeddings. Null for storage-only vaults.
	VectorBucket string               `json:"vectorBucket" api:"nullable"`
	JSON         vaultNewResponseJSON `json:"-"`
}

// vaultNewResponseJSON contains the JSON metadata for the struct
// [VaultNewResponse]
type vaultNewResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	EnableIndexing apijson.Field
	FilesBucket    apijson.Field
	IndexName      apijson.Field
	Name           apijson.Field
	Region         apijson.Field
	VectorBucket   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultNewResponseJSON) RawJSON() string {
	return r.raw
}

type VaultGetResponse struct {
	// Vault identifier
	ID string `json:"id" api:"required"`
	// Vault creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// S3 bucket for document storage
	FilesBucket string `json:"filesBucket" api:"required"`
	// Vault name
	Name string `json:"name" api:"required"`
	// AWS region
	Region string `json:"region" api:"required"`
	// Document chunking strategy configuration
	ChunkStrategy VaultGetResponseChunkStrategy `json:"chunkStrategy"`
	// Vault description
	Description string `json:"description"`
	// Whether GraphRAG is enabled
	EnableGraph bool `json:"enableGraph"`
	// Search index name
	IndexName string `json:"indexName"`
	// KMS key for encryption
	KmsKeyID string `json:"kmsKeyId"`
	// Additional vault metadata
	Metadata interface{} `json:"metadata"`
	// Total storage size in bytes
	TotalBytes int64 `json:"totalBytes"`
	// Number of stored documents
	TotalObjects int64 `json:"totalObjects"`
	// Number of vector embeddings
	TotalVectors int64 `json:"totalVectors"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// S3 bucket for vector embeddings
	VectorBucket string               `json:"vectorBucket" api:"nullable"`
	JSON         vaultGetResponseJSON `json:"-"`
}

// vaultGetResponseJSON contains the JSON metadata for the struct
// [VaultGetResponse]
type vaultGetResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	FilesBucket   apijson.Field
	Name          apijson.Field
	Region        apijson.Field
	ChunkStrategy apijson.Field
	Description   apijson.Field
	EnableGraph   apijson.Field
	IndexName     apijson.Field
	KmsKeyID      apijson.Field
	Metadata      apijson.Field
	TotalBytes    apijson.Field
	TotalObjects  apijson.Field
	TotalVectors  apijson.Field
	UpdatedAt     apijson.Field
	VectorBucket  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultGetResponseJSON) RawJSON() string {
	return r.raw
}

// Document chunking strategy configuration
type VaultGetResponseChunkStrategy struct {
	// Target size for each chunk in tokens
	ChunkSize int64 `json:"chunkSize"`
	// Chunking method (e.g., 'semantic', 'fixed')
	Method string `json:"method"`
	// Minimum chunk size in tokens
	MinChunkSize int64 `json:"minChunkSize"`
	// Number of overlapping tokens between chunks
	Overlap int64                             `json:"overlap"`
	JSON    vaultGetResponseChunkStrategyJSON `json:"-"`
}

// vaultGetResponseChunkStrategyJSON contains the JSON metadata for the struct
// [VaultGetResponseChunkStrategy]
type vaultGetResponseChunkStrategyJSON struct {
	ChunkSize    apijson.Field
	Method       apijson.Field
	MinChunkSize apijson.Field
	Overlap      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VaultGetResponseChunkStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultGetResponseChunkStrategyJSON) RawJSON() string {
	return r.raw
}

type VaultUpdateResponse struct {
	// Vault identifier
	ID string `json:"id"`
	// Document chunking strategy configuration
	ChunkStrategy interface{} `json:"chunkStrategy"`
	// Vault creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Vault description
	Description string `json:"description" api:"nullable"`
	// Whether GraphRAG is enabled for future uploads
	EnableGraph bool `json:"enableGraph"`
	// S3 bucket for document storage
	FilesBucket string `json:"filesBucket"`
	// Search index name
	IndexName string `json:"indexName"`
	// KMS key for encryption
	KmsKeyID string `json:"kmsKeyId"`
	// Additional vault metadata
	Metadata interface{} `json:"metadata"`
	// Vault name
	Name string `json:"name"`
	// AWS region
	Region string `json:"region"`
	// Total storage size in bytes
	TotalBytes int64 `json:"totalBytes"`
	// Number of stored documents
	TotalObjects int64 `json:"totalObjects"`
	// Number of vector embeddings
	TotalVectors int64 `json:"totalVectors"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// S3 bucket for vector embeddings
	VectorBucket string                  `json:"vectorBucket" api:"nullable"`
	JSON         vaultUpdateResponseJSON `json:"-"`
}

// vaultUpdateResponseJSON contains the JSON metadata for the struct
// [VaultUpdateResponse]
type vaultUpdateResponseJSON struct {
	ID            apijson.Field
	ChunkStrategy apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	EnableGraph   apijson.Field
	FilesBucket   apijson.Field
	IndexName     apijson.Field
	KmsKeyID      apijson.Field
	Metadata      apijson.Field
	Name          apijson.Field
	Region        apijson.Field
	TotalBytes    apijson.Field
	TotalObjects  apijson.Field
	TotalVectors  apijson.Field
	UpdatedAt     apijson.Field
	VectorBucket  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type VaultListResponse struct {
	// Total number of vaults
	Total  int64                    `json:"total"`
	Vaults []VaultListResponseVault `json:"vaults"`
	JSON   vaultListResponseJSON    `json:"-"`
}

// vaultListResponseJSON contains the JSON metadata for the struct
// [VaultListResponse]
type vaultListResponseJSON struct {
	Total       apijson.Field
	Vaults      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultListResponseJSON) RawJSON() string {
	return r.raw
}

type VaultListResponseVault struct {
	// Vault identifier
	ID string `json:"id"`
	// Vault creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Vault description
	Description string `json:"description"`
	// Whether GraphRAG is enabled
	EnableGraph bool `json:"enableGraph"`
	// Vault name
	Name string `json:"name"`
	// Total storage size in bytes
	TotalBytes int64 `json:"totalBytes"`
	// Number of stored documents
	TotalObjects int64                      `json:"totalObjects"`
	JSON         vaultListResponseVaultJSON `json:"-"`
}

// vaultListResponseVaultJSON contains the JSON metadata for the struct
// [VaultListResponseVault]
type vaultListResponseVaultJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Description  apijson.Field
	EnableGraph  apijson.Field
	Name         apijson.Field
	TotalBytes   apijson.Field
	TotalObjects apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VaultListResponseVault) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultListResponseVaultJSON) RawJSON() string {
	return r.raw
}

type VaultDeleteResponse struct {
	DeletedVault VaultDeleteResponseDeletedVault `json:"deletedVault"`
	// Either 'deleted' or 'deletion_queued'
	Status  string                  `json:"status"`
	Success bool                    `json:"success"`
	JSON    vaultDeleteResponseJSON `json:"-"`
}

// vaultDeleteResponseJSON contains the JSON metadata for the struct
// [VaultDeleteResponse]
type vaultDeleteResponseJSON struct {
	DeletedVault apijson.Field
	Status       apijson.Field
	Success      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VaultDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type VaultDeleteResponseDeletedVault struct {
	ID             string                              `json:"id"`
	BytesFreed     int64                               `json:"bytesFreed"`
	Name           string                              `json:"name"`
	ObjectsDeleted int64                               `json:"objectsDeleted"`
	VectorsDeleted int64                               `json:"vectorsDeleted"`
	JSON           vaultDeleteResponseDeletedVaultJSON `json:"-"`
}

// vaultDeleteResponseDeletedVaultJSON contains the JSON metadata for the struct
// [VaultDeleteResponseDeletedVault]
type vaultDeleteResponseDeletedVaultJSON struct {
	ID             apijson.Field
	BytesFreed     apijson.Field
	Name           apijson.Field
	ObjectsDeleted apijson.Field
	VectorsDeleted apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultDeleteResponseDeletedVault) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultDeleteResponseDeletedVaultJSON) RawJSON() string {
	return r.raw
}

type VaultConfirmUploadResponse struct {
	AlreadyConfirmed bool                             `json:"alreadyConfirmed"`
	ObjectID         string                           `json:"objectId"`
	Status           VaultConfirmUploadResponseStatus `json:"status"`
	VaultID          string                           `json:"vaultId"`
	JSON             vaultConfirmUploadResponseJSON   `json:"-"`
}

// vaultConfirmUploadResponseJSON contains the JSON metadata for the struct
// [VaultConfirmUploadResponse]
type vaultConfirmUploadResponseJSON struct {
	AlreadyConfirmed apijson.Field
	ObjectID         apijson.Field
	Status           apijson.Field
	VaultID          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VaultConfirmUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultConfirmUploadResponseJSON) RawJSON() string {
	return r.raw
}

type VaultConfirmUploadResponseStatus string

const (
	VaultConfirmUploadResponseStatusCompleted VaultConfirmUploadResponseStatus = "completed"
	VaultConfirmUploadResponseStatusFailed    VaultConfirmUploadResponseStatus = "failed"
)

func (r VaultConfirmUploadResponseStatus) IsKnown() bool {
	switch r {
	case VaultConfirmUploadResponseStatusCompleted, VaultConfirmUploadResponseStatusFailed:
		return true
	}
	return false
}

type VaultIngestResponse struct {
	// Always false - GraphRAG must be triggered separately via POST
	// /vault/:id/graphrag/:objectId
	EnableGraphRag bool `json:"enableGraphRAG" api:"required"`
	// Human-readable status message
	Message string `json:"message" api:"required"`
	// ID of the vault object being processed
	ObjectID string `json:"objectId" api:"required"`
	// Current ingestion status. 'stored' for file types without text extraction (no
	// chunks/vectors created).
	Status VaultIngestResponseStatus `json:"status" api:"required"`
	// Workflow run ID for tracking progress. Null for file types that skip processing.
	WorkflowID string                  `json:"workflowId" api:"required,nullable"`
	JSON       vaultIngestResponseJSON `json:"-"`
}

// vaultIngestResponseJSON contains the JSON metadata for the struct
// [VaultIngestResponse]
type vaultIngestResponseJSON struct {
	EnableGraphRag apijson.Field
	Message        apijson.Field
	ObjectID       apijson.Field
	Status         apijson.Field
	WorkflowID     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultIngestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultIngestResponseJSON) RawJSON() string {
	return r.raw
}

// Current ingestion status. 'stored' for file types without text extraction (no
// chunks/vectors created).
type VaultIngestResponseStatus string

const (
	VaultIngestResponseStatusProcessing VaultIngestResponseStatus = "processing"
	VaultIngestResponseStatusStored     VaultIngestResponseStatus = "stored"
)

func (r VaultIngestResponseStatus) IsKnown() bool {
	switch r {
	case VaultIngestResponseStatusProcessing, VaultIngestResponseStatusStored:
		return true
	}
	return false
}

type VaultSearchResponse struct {
	// Relevant text chunks with similarity scores and page locations
	Chunks []VaultSearchResponseChunk `json:"chunks"`
	// Search method used
	Method string `json:"method"`
	// Original search query
	Query string `json:"query"`
	// AI-generated answer based on search results (for global/entity methods)
	Response string                      `json:"response"`
	Sources  []VaultSearchResponseSource `json:"sources"`
	// ID of the searched vault
	VaultID string                  `json:"vault_id"`
	JSON    vaultSearchResponseJSON `json:"-"`
}

// vaultSearchResponseJSON contains the JSON metadata for the struct
// [VaultSearchResponse]
type vaultSearchResponseJSON struct {
	Chunks      apijson.Field
	Method      apijson.Field
	Query       apijson.Field
	Response    apijson.Field
	Sources     apijson.Field
	VaultID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultSearchResponseJSON) RawJSON() string {
	return r.raw
}

type VaultSearchResponseChunk struct {
	// Index of the chunk within the document (0-based)
	ChunkIndex int64 `json:"chunk_index"`
	// Vector similarity distance (lower is more similar)
	Distance float64 `json:"distance"`
	// ID of the source document
	ObjectID string `json:"object_id"`
	// PDF page number where the chunk ends (1-indexed). Null for non-PDF documents or
	// documents ingested before page tracking was added.
	PageEnd int64 `json:"page_end" api:"nullable"`
	// PDF page number where the chunk begins (1-indexed). Null for non-PDF documents
	// or documents ingested before page tracking was added.
	PageStart int64 `json:"page_start" api:"nullable"`
	// Relevance score (deprecated, use distance or hybridScore)
	Score float64 `json:"score"`
	// Source identifier (deprecated, use object_id)
	Source string `json:"source"`
	// Preview of the chunk text (up to 500 characters)
	Text string `json:"text"`
	// Ending word index (0-based) in the OCR word list. Use with GET
	// /vault/:id/objects/:objectId/ocr-words to retrieve bounding boxes for
	// highlighting.
	WordEndIndex int64 `json:"word_end_index" api:"nullable"`
	// Starting word index (0-based) in the OCR word list. Use with GET
	// /vault/:id/objects/:objectId/ocr-words to retrieve bounding boxes for
	// highlighting.
	WordStartIndex int64                        `json:"word_start_index" api:"nullable"`
	JSON           vaultSearchResponseChunkJSON `json:"-"`
}

// vaultSearchResponseChunkJSON contains the JSON metadata for the struct
// [VaultSearchResponseChunk]
type vaultSearchResponseChunkJSON struct {
	ChunkIndex     apijson.Field
	Distance       apijson.Field
	ObjectID       apijson.Field
	PageEnd        apijson.Field
	PageStart      apijson.Field
	Score          apijson.Field
	Source         apijson.Field
	Text           apijson.Field
	WordEndIndex   apijson.Field
	WordStartIndex apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultSearchResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultSearchResponseChunkJSON) RawJSON() string {
	return r.raw
}

type VaultSearchResponseSource struct {
	ID                   string                        `json:"id"`
	ChunkCount           int64                         `json:"chunkCount"`
	CreatedAt            time.Time                     `json:"createdAt" format:"date-time"`
	Filename             string                        `json:"filename"`
	IngestionCompletedAt time.Time                     `json:"ingestionCompletedAt" format:"date-time"`
	PageCount            int64                         `json:"pageCount"`
	TextLength           int64                         `json:"textLength"`
	JSON                 vaultSearchResponseSourceJSON `json:"-"`
}

// vaultSearchResponseSourceJSON contains the JSON metadata for the struct
// [VaultSearchResponseSource]
type vaultSearchResponseSourceJSON struct {
	ID                   apijson.Field
	ChunkCount           apijson.Field
	CreatedAt            apijson.Field
	Filename             apijson.Field
	IngestionCompletedAt apijson.Field
	PageCount            apijson.Field
	TextLength           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *VaultSearchResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultSearchResponseSourceJSON) RawJSON() string {
	return r.raw
}

type VaultUploadResponse struct {
	// Whether the file will be automatically indexed
	AutoIndex bool `json:"auto_index"`
	// Whether the vault supports indexing. False for storage-only vaults.
	EnableIndexing bool `json:"enableIndexing"`
	// URL expiration time in seconds
	ExpiresIn    float64                         `json:"expiresIn"`
	Instructions VaultUploadResponseInstructions `json:"instructions"`
	// Next API endpoint to call for processing
	NextStep string `json:"next_step" api:"nullable"`
	// Unique identifier for the uploaded object
	ObjectID string `json:"objectId"`
	// Folder path for hierarchy if provided
	Path string `json:"path" api:"nullable"`
	// S3 object key for the file
	S3Key string `json:"s3Key"`
	// Presigned URL for uploading the file
	UploadURL string                  `json:"uploadUrl"`
	JSON      vaultUploadResponseJSON `json:"-"`
}

// vaultUploadResponseJSON contains the JSON metadata for the struct
// [VaultUploadResponse]
type vaultUploadResponseJSON struct {
	AutoIndex      apijson.Field
	EnableIndexing apijson.Field
	ExpiresIn      apijson.Field
	Instructions   apijson.Field
	NextStep       apijson.Field
	ObjectID       apijson.Field
	Path           apijson.Field
	S3Key          apijson.Field
	UploadURL      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultUploadResponseJSON) RawJSON() string {
	return r.raw
}

type VaultUploadResponseInstructions struct {
	Headers interface{}                         `json:"headers"`
	Method  string                              `json:"method"`
	Note    string                              `json:"note"`
	JSON    vaultUploadResponseInstructionsJSON `json:"-"`
}

// vaultUploadResponseInstructionsJSON contains the JSON metadata for the struct
// [VaultUploadResponseInstructions]
type vaultUploadResponseInstructionsJSON struct {
	Headers     apijson.Field
	Method      apijson.Field
	Note        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultUploadResponseInstructions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultUploadResponseInstructionsJSON) RawJSON() string {
	return r.raw
}

type VaultNewParams struct {
	// Display name for the vault
	Name param.Field[string] `json:"name" api:"required"`
	// Optional description of the vault's purpose
	Description param.Field[string] `json:"description"`
	// Enable knowledge graph for entity relationship mapping. Only applies when
	// enableIndexing is true.
	EnableGraph param.Field[bool] `json:"enableGraph"`
	// Enable vector indexing and search capabilities. Set to false for storage-only
	// vaults.
	EnableIndexing param.Field[bool] `json:"enableIndexing"`
	// Assign the vault to a vault group for access control. Required when using a
	// group-scoped API key.
	GroupID param.Field[string] `json:"groupId"`
	// Optional metadata to attach to the vault (e.g., { containsPHI: true } for HIPAA
	// compliance tracking)
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r VaultNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultUpdateParams struct {
	// New description for the vault. Set to null to remove.
	Description param.Field[string] `json:"description"`
	// Whether to enable GraphRAG for future document uploads
	EnableGraph param.Field[bool] `json:"enableGraph"`
	// Move the vault to a different group, or set to null to remove from its current
	// group.
	GroupID param.Field[string] `json:"groupId"`
	// New name for the vault
	Name param.Field[string] `json:"name"`
}

func (r VaultUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultDeleteParams struct {
	// If true and vault has many objects, queue deletion in background and return
	// immediately
	Async param.Field[bool] `query:"async"`
}

// URLQuery serializes [VaultDeleteParams]'s query parameters as `url.Values`.
func (r VaultDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VaultConfirmUploadParams struct {
	Body VaultConfirmUploadParamsBodyUnion `json:"body" api:"required"`
}

func (r VaultConfirmUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type VaultConfirmUploadParamsBody struct {
	// Whether the upload succeeded
	Success param.Field[VaultConfirmUploadParamsBodySuccess] `json:"success" api:"required"`
	// Client-side error code
	ErrorCode param.Field[string] `json:"errorCode"`
	// Client-side error message
	ErrorMessage param.Field[string] `json:"errorMessage"`
	// S3 ETag for the uploaded object (optional if client cannot access ETag header)
	Etag param.Field[string] `json:"etag"`
	// Uploaded file size in bytes
	SizeBytes param.Field[int64] `json:"sizeBytes"`
}

func (r VaultConfirmUploadParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VaultConfirmUploadParamsBody) implementsVaultConfirmUploadParamsBodyUnion() {}

// Satisfied by [VaultConfirmUploadParamsBodyVaultConfirmUploadSuccess],
// [VaultConfirmUploadParamsBodyVaultConfirmUploadFailure],
// [VaultConfirmUploadParamsBody].
type VaultConfirmUploadParamsBodyUnion interface {
	implementsVaultConfirmUploadParamsBodyUnion()
}

type VaultConfirmUploadParamsBodyVaultConfirmUploadSuccess struct {
	// Uploaded file size in bytes
	SizeBytes param.Field[int64] `json:"sizeBytes" api:"required"`
	// Whether the upload succeeded
	Success param.Field[VaultConfirmUploadParamsBodyVaultConfirmUploadSuccessSuccess] `json:"success" api:"required"`
	// S3 ETag for the uploaded object (optional if client cannot access ETag header)
	Etag param.Field[string] `json:"etag"`
}

func (r VaultConfirmUploadParamsBodyVaultConfirmUploadSuccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VaultConfirmUploadParamsBodyVaultConfirmUploadSuccess) implementsVaultConfirmUploadParamsBodyUnion() {
}

// Whether the upload succeeded
type VaultConfirmUploadParamsBodyVaultConfirmUploadSuccessSuccess bool

const (
	VaultConfirmUploadParamsBodyVaultConfirmUploadSuccessSuccessTrue VaultConfirmUploadParamsBodyVaultConfirmUploadSuccessSuccess = true
)

func (r VaultConfirmUploadParamsBodyVaultConfirmUploadSuccessSuccess) IsKnown() bool {
	switch r {
	case VaultConfirmUploadParamsBodyVaultConfirmUploadSuccessSuccessTrue:
		return true
	}
	return false
}

type VaultConfirmUploadParamsBodyVaultConfirmUploadFailure struct {
	// Client-side error code
	ErrorCode param.Field[string] `json:"errorCode" api:"required"`
	// Client-side error message
	ErrorMessage param.Field[string] `json:"errorMessage" api:"required"`
	// Whether the upload succeeded
	Success param.Field[VaultConfirmUploadParamsBodyVaultConfirmUploadFailureSuccess] `json:"success" api:"required"`
}

func (r VaultConfirmUploadParamsBodyVaultConfirmUploadFailure) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VaultConfirmUploadParamsBodyVaultConfirmUploadFailure) implementsVaultConfirmUploadParamsBodyUnion() {
}

// Whether the upload succeeded
type VaultConfirmUploadParamsBodyVaultConfirmUploadFailureSuccess bool

const (
	VaultConfirmUploadParamsBodyVaultConfirmUploadFailureSuccessFalse VaultConfirmUploadParamsBodyVaultConfirmUploadFailureSuccess = false
)

func (r VaultConfirmUploadParamsBodyVaultConfirmUploadFailureSuccess) IsKnown() bool {
	switch r {
	case VaultConfirmUploadParamsBodyVaultConfirmUploadFailureSuccessFalse:
		return true
	}
	return false
}

// Whether the upload succeeded
type VaultConfirmUploadParamsBodySuccess bool

const (
	VaultConfirmUploadParamsBodySuccessTrue  VaultConfirmUploadParamsBodySuccess = true
	VaultConfirmUploadParamsBodySuccessFalse VaultConfirmUploadParamsBodySuccess = false
)

func (r VaultConfirmUploadParamsBodySuccess) IsKnown() bool {
	switch r {
	case VaultConfirmUploadParamsBodySuccessTrue, VaultConfirmUploadParamsBodySuccessFalse:
		return true
	}
	return false
}

type VaultSearchParams struct {
	// Search query or question to find relevant documents
	Query param.Field[string] `json:"query" api:"required"`
	// Filters to narrow search results to specific documents
	Filters param.Field[VaultSearchParamsFilters] `json:"filters"`
	// Search method: 'global' for comprehensive questions, 'entity' for specific
	// entities, 'fast' for quick similarity search, 'hybrid' for combined approach
	Method param.Field[VaultSearchParamsMethod] `json:"method"`
	// Maximum number of results to return
	TopK param.Field[int64] `json:"topK"`
}

func (r VaultSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filters to narrow search results to specific documents
type VaultSearchParamsFilters struct {
	// Filter to specific document(s) by object ID. Accepts a single ID or array of
	// IDs.
	ObjectID    param.Field[VaultSearchParamsFiltersObjectIDUnion] `json:"object_id"`
	ExtraFields map[string]interface{}                             `json:"-,extras"`
}

func (r VaultSearchParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter to specific document(s) by object ID. Accepts a single ID or array of
// IDs.
//
// Satisfied by [shared.UnionString], [VaultSearchParamsFiltersObjectIDArray].
type VaultSearchParamsFiltersObjectIDUnion interface {
	ImplementsVaultSearchParamsFiltersObjectIDUnion()
}

type VaultSearchParamsFiltersObjectIDArray []string

func (r VaultSearchParamsFiltersObjectIDArray) ImplementsVaultSearchParamsFiltersObjectIDUnion() {}

// Search method: 'global' for comprehensive questions, 'entity' for specific
// entities, 'fast' for quick similarity search, 'hybrid' for combined approach
type VaultSearchParamsMethod string

const (
	VaultSearchParamsMethodVector VaultSearchParamsMethod = "vector"
	VaultSearchParamsMethodGraph  VaultSearchParamsMethod = "graph"
	VaultSearchParamsMethodHybrid VaultSearchParamsMethod = "hybrid"
	VaultSearchParamsMethodGlobal VaultSearchParamsMethod = "global"
	VaultSearchParamsMethodLocal  VaultSearchParamsMethod = "local"
	VaultSearchParamsMethodFast   VaultSearchParamsMethod = "fast"
	VaultSearchParamsMethodEntity VaultSearchParamsMethod = "entity"
)

func (r VaultSearchParamsMethod) IsKnown() bool {
	switch r {
	case VaultSearchParamsMethodVector, VaultSearchParamsMethodGraph, VaultSearchParamsMethodHybrid, VaultSearchParamsMethodGlobal, VaultSearchParamsMethodLocal, VaultSearchParamsMethodFast, VaultSearchParamsMethodEntity:
		return true
	}
	return false
}

type VaultUploadParams struct {
	// MIME type of the file (e.g., application/pdf, image/jpeg)
	ContentType param.Field[string] `json:"contentType" api:"required"`
	// Name of the file to upload
	Filename param.Field[string] `json:"filename" api:"required"`
	// Whether to automatically process and index the file for search
	AutoIndex param.Field[bool] `json:"auto_index"`
	// Additional metadata to associate with the file
	Metadata param.Field[interface{}] `json:"metadata"`
	// Optional folder path for hierarchy preservation. Allows integrations to maintain
	// source folder structure from systems like NetDocs, Clio, or Smokeball. Example:
	// '/Discovery/Depositions/2024'
	Path param.Field[string] `json:"path"`
	// File size in bytes (optional, max 5GB for single PUT uploads). When provided,
	// enforces exact file size at S3 level.
	SizeBytes param.Field[int64] `json:"sizeBytes"`
}

func (r VaultUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
