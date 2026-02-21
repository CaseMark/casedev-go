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

// VaultObjectService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultObjectService] method instead.
type VaultObjectService struct {
	Options []option.RequestOption
}

// NewVaultObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultObjectService(opts ...option.RequestOption) (r *VaultObjectService) {
	r = &VaultObjectService{}
	r.Options = opts
	return
}

// Retrieves metadata for a specific document in a vault and generates a temporary
// download URL. The download URL expires after 1 hour for security. This endpoint
// also updates the file size if it wasn't previously calculated.
func (r *VaultObjectService) Get(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *VaultObjectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a document's filename, path, or metadata. Use this to rename files or
// organize them into virtual folders. The path is stored in metadata.path and can
// be used to build folder hierarchies in your application.
func (r *VaultObjectService) Update(ctx context.Context, id string, objectID string, body VaultObjectUpdateParams, opts ...option.RequestOption) (res *VaultObjectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all objects stored in a specific vault, including document metadata,
// ingestion status, and processing statistics.
func (r *VaultObjectService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *VaultObjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a document from the vault including all associated vectors,
// chunks, graph data, and the original file. This operation cannot be undone.
func (r *VaultObjectService) Delete(ctx context.Context, id string, objectID string, body VaultObjectDeleteParams, opts ...option.RequestOption) (res *VaultObjectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Generate presigned URLs for direct S3 operations (GET, PUT, DELETE, HEAD) on
// vault objects. This allows secure, time-limited access to files without proxying
// through the API. Essential for large document uploads/downloads in legal
// workflows.
func (r *VaultObjectService) NewPresignedURL(ctx context.Context, id string, objectID string, body VaultObjectNewPresignedURLParams, opts ...option.RequestOption) (res *VaultObjectNewPresignedURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s/presigned-url", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Downloads a file from a vault. Returns the actual file content as a binary
// stream with appropriate headers for file download. Useful for retrieving
// contracts, depositions, case files, and other legal documents stored in your
// vault.
func (r *VaultObjectService) Download(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s/download", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves word-level OCR bounding box data for a processed PDF document. Each
// word includes its text, normalized bounding box coordinates (0-1 range),
// confidence score, and global word index. Use this data to highlight specific
// text ranges in a PDF viewer based on word indices from search results.
func (r *VaultObjectService) GetOcrWords(ctx context.Context, id string, objectID string, query VaultObjectGetOcrWordsParams, opts ...option.RequestOption) (res *VaultObjectGetOcrWordsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s/ocr-words", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the status of a CaseMark summary workflow job.
func (r *VaultObjectService) GetSummarizeJob(ctx context.Context, id string, objectID string, jobID string, opts ...option.RequestOption) (res *VaultObjectGetSummarizeJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s/summarize/%s", id, objectID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the full extracted text content from a processed vault object. Returns
// the concatenated text from all chunks, useful for document review, analysis, or
// export. The object must have completed processing before text can be retrieved.
func (r *VaultObjectService) GetText(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *VaultObjectGetTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/objects/%s/text", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VaultObjectGetResponse struct {
	// Object ID
	ID string `json:"id,required"`
	// MIME type
	ContentType string `json:"contentType,required"`
	// Upload timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Presigned S3 download URL
	DownloadURL string `json:"downloadUrl,required"`
	// URL expiration time in seconds
	ExpiresIn int64 `json:"expiresIn,required"`
	// Original filename
	Filename string `json:"filename,required"`
	// Processing status (pending, processing, completed, failed)
	IngestionStatus string `json:"ingestionStatus,required"`
	// Vault ID
	VaultID string `json:"vaultId,required"`
	// Number of text chunks created
	ChunkCount int64 `json:"chunkCount"`
	// Additional metadata
	Metadata interface{} `json:"metadata"`
	// Number of pages (for documents)
	PageCount int64 `json:"pageCount"`
	// Optional folder path for hierarchy preservation
	Path string `json:"path,nullable"`
	// File size in bytes
	SizeBytes int64 `json:"sizeBytes"`
	// Length of extracted text
	TextLength int64 `json:"textLength"`
	// Number of embedding vectors generated
	VectorCount int64                      `json:"vectorCount"`
	JSON        vaultObjectGetResponseJSON `json:"-"`
}

// vaultObjectGetResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetResponse]
type vaultObjectGetResponseJSON struct {
	ID              apijson.Field
	ContentType     apijson.Field
	CreatedAt       apijson.Field
	DownloadURL     apijson.Field
	ExpiresIn       apijson.Field
	Filename        apijson.Field
	IngestionStatus apijson.Field
	VaultID         apijson.Field
	ChunkCount      apijson.Field
	Metadata        apijson.Field
	PageCount       apijson.Field
	Path            apijson.Field
	SizeBytes       apijson.Field
	TextLength      apijson.Field
	VectorCount     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaultObjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectUpdateResponse struct {
	// Object ID
	ID string `json:"id"`
	// MIME type
	ContentType string `json:"contentType"`
	// Updated filename
	Filename string `json:"filename"`
	// Processing status
	IngestionStatus string `json:"ingestionStatus"`
	// Full metadata object
	Metadata interface{} `json:"metadata"`
	// Folder path for hierarchy preservation
	Path string `json:"path,nullable"`
	// File size in bytes
	SizeBytes int64 `json:"sizeBytes"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Vault ID
	VaultID string                        `json:"vaultId"`
	JSON    vaultObjectUpdateResponseJSON `json:"-"`
}

// vaultObjectUpdateResponseJSON contains the JSON metadata for the struct
// [VaultObjectUpdateResponse]
type vaultObjectUpdateResponseJSON struct {
	ID              apijson.Field
	ContentType     apijson.Field
	Filename        apijson.Field
	IngestionStatus apijson.Field
	Metadata        apijson.Field
	Path            apijson.Field
	SizeBytes       apijson.Field
	UpdatedAt       apijson.Field
	VaultID         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaultObjectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectListResponse struct {
	// Total number of objects in the vault
	Count   float64                         `json:"count,required"`
	Objects []VaultObjectListResponseObject `json:"objects,required"`
	// The ID of the vault
	VaultID string                      `json:"vaultId,required"`
	JSON    vaultObjectListResponseJSON `json:"-"`
}

// vaultObjectListResponseJSON contains the JSON metadata for the struct
// [VaultObjectListResponse]
type vaultObjectListResponseJSON struct {
	Count       apijson.Field
	Objects     apijson.Field
	VaultID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectListResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectListResponseObject struct {
	// Unique object identifier
	ID string `json:"id,required"`
	// MIME type of the document
	ContentType string `json:"contentType,required"`
	// Document upload timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Original filename of the uploaded document
	Filename string `json:"filename,required"`
	// Processing status of the document
	IngestionStatus string `json:"ingestionStatus,required"`
	// Number of text chunks created for vectorization
	ChunkCount float64 `json:"chunkCount"`
	// Processing completion timestamp
	IngestionCompletedAt time.Time `json:"ingestionCompletedAt" format:"date-time"`
	// Custom metadata associated with the document
	Metadata interface{} `json:"metadata"`
	// Number of pages in the document
	PageCount float64 `json:"pageCount"`
	// Optional folder path for hierarchy preservation from source systems
	Path string `json:"path,nullable"`
	// File size in bytes
	SizeBytes float64 `json:"sizeBytes"`
	// Custom tags associated with the document
	Tags []string `json:"tags"`
	// Total character count of extracted text
	TextLength float64 `json:"textLength"`
	// Number of vectors generated for semantic search
	VectorCount float64                           `json:"vectorCount"`
	JSON        vaultObjectListResponseObjectJSON `json:"-"`
}

// vaultObjectListResponseObjectJSON contains the JSON metadata for the struct
// [VaultObjectListResponseObject]
type vaultObjectListResponseObjectJSON struct {
	ID                   apijson.Field
	ContentType          apijson.Field
	CreatedAt            apijson.Field
	Filename             apijson.Field
	IngestionStatus      apijson.Field
	ChunkCount           apijson.Field
	IngestionCompletedAt apijson.Field
	Metadata             apijson.Field
	PageCount            apijson.Field
	Path                 apijson.Field
	SizeBytes            apijson.Field
	Tags                 apijson.Field
	TextLength           apijson.Field
	VectorCount          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *VaultObjectListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectListResponseObjectJSON) RawJSON() string {
	return r.raw
}

type VaultObjectDeleteResponse struct {
	DeletedObject VaultObjectDeleteResponseDeletedObject `json:"deletedObject"`
	Success       bool                                   `json:"success"`
	JSON          vaultObjectDeleteResponseJSON          `json:"-"`
}

// vaultObjectDeleteResponseJSON contains the JSON metadata for the struct
// [VaultObjectDeleteResponse]
type vaultObjectDeleteResponseJSON struct {
	DeletedObject apijson.Field
	Success       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultObjectDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectDeleteResponseDeletedObject struct {
	// Deleted object ID
	ID string `json:"id"`
	// Original filename
	Filename string `json:"filename"`
	// Size of deleted file in bytes
	SizeBytes int64 `json:"sizeBytes"`
	// Number of vectors deleted
	VectorsDeleted int64                                      `json:"vectorsDeleted"`
	JSON           vaultObjectDeleteResponseDeletedObjectJSON `json:"-"`
}

// vaultObjectDeleteResponseDeletedObjectJSON contains the JSON metadata for the
// struct [VaultObjectDeleteResponseDeletedObject]
type vaultObjectDeleteResponseDeletedObjectJSON struct {
	ID             apijson.Field
	Filename       apijson.Field
	SizeBytes      apijson.Field
	VectorsDeleted apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultObjectDeleteResponseDeletedObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectDeleteResponseDeletedObjectJSON) RawJSON() string {
	return r.raw
}

type VaultObjectNewPresignedURLResponse struct {
	// URL expiration timestamp
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// URL expiration time in seconds
	ExpiresIn int64 `json:"expiresIn"`
	// Original filename
	Filename string `json:"filename"`
	// Usage instructions and examples
	Instructions interface{}                                `json:"instructions"`
	Metadata     VaultObjectNewPresignedURLResponseMetadata `json:"metadata"`
	// The object identifier
	ObjectID string `json:"objectId"`
	// The operation type
	Operation string `json:"operation"`
	// The presigned URL for direct S3 access
	PresignedURL string `json:"presignedUrl"`
	// S3 object key
	S3Key string `json:"s3Key"`
	// The vault identifier
	VaultID string                                 `json:"vaultId"`
	JSON    vaultObjectNewPresignedURLResponseJSON `json:"-"`
}

// vaultObjectNewPresignedURLResponseJSON contains the JSON metadata for the struct
// [VaultObjectNewPresignedURLResponse]
type vaultObjectNewPresignedURLResponseJSON struct {
	ExpiresAt    apijson.Field
	ExpiresIn    apijson.Field
	Filename     apijson.Field
	Instructions apijson.Field
	Metadata     apijson.Field
	ObjectID     apijson.Field
	Operation    apijson.Field
	PresignedURL apijson.Field
	S3Key        apijson.Field
	VaultID      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VaultObjectNewPresignedURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectNewPresignedURLResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectNewPresignedURLResponseMetadata struct {
	Bucket      string                                         `json:"bucket"`
	ContentType string                                         `json:"contentType"`
	Region      string                                         `json:"region"`
	SizeBytes   int64                                          `json:"sizeBytes"`
	JSON        vaultObjectNewPresignedURLResponseMetadataJSON `json:"-"`
}

// vaultObjectNewPresignedURLResponseMetadataJSON contains the JSON metadata for
// the struct [VaultObjectNewPresignedURLResponseMetadata]
type vaultObjectNewPresignedURLResponseMetadataJSON struct {
	Bucket      apijson.Field
	ContentType apijson.Field
	Region      apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectNewPresignedURLResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectNewPresignedURLResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetOcrWordsResponse struct {
	// When the OCR data was extracted
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The object ID
	ObjectID string `json:"objectId"`
	// Total number of pages in the document
	PageCount int64 `json:"pageCount"`
	// Per-page word data with bounding boxes
	Pages []VaultObjectGetOcrWordsResponsePage `json:"pages"`
	// Total number of words extracted from the document
	TotalWords int64                              `json:"totalWords"`
	JSON       vaultObjectGetOcrWordsResponseJSON `json:"-"`
}

// vaultObjectGetOcrWordsResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetOcrWordsResponse]
type vaultObjectGetOcrWordsResponseJSON struct {
	CreatedAt   apijson.Field
	ObjectID    apijson.Field
	PageCount   apijson.Field
	Pages       apijson.Field
	TotalWords  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetOcrWordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetOcrWordsResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetOcrWordsResponsePage struct {
	// Page number (1-indexed)
	Page  int64                                     `json:"page"`
	Words []VaultObjectGetOcrWordsResponsePagesWord `json:"words"`
	JSON  vaultObjectGetOcrWordsResponsePageJSON    `json:"-"`
}

// vaultObjectGetOcrWordsResponsePageJSON contains the JSON metadata for the struct
// [VaultObjectGetOcrWordsResponsePage]
type vaultObjectGetOcrWordsResponsePageJSON struct {
	Page        apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetOcrWordsResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetOcrWordsResponsePageJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetOcrWordsResponsePagesWord struct {
	// Bounding box [x0, y0, x1, y1] normalized to 0-1 range
	Bbox []float64 `json:"bbox"`
	// OCR confidence score (0-1)
	Confidence float64 `json:"confidence,nullable"`
	// The word text
	Text string `json:"text"`
	// Global word index across the entire document (0-based)
	WordIndex int64                                       `json:"wordIndex"`
	JSON      vaultObjectGetOcrWordsResponsePagesWordJSON `json:"-"`
}

// vaultObjectGetOcrWordsResponsePagesWordJSON contains the JSON metadata for the
// struct [VaultObjectGetOcrWordsResponsePagesWord]
type vaultObjectGetOcrWordsResponsePagesWordJSON struct {
	Bbox        apijson.Field
	Confidence  apijson.Field
	Text        apijson.Field
	WordIndex   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetOcrWordsResponsePagesWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetOcrWordsResponsePagesWordJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetSummarizeJobResponse struct {
	// When the job completed
	CompletedAt time.Time `json:"completedAt,nullable" format:"date-time"`
	// When the job was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Error message (if failed)
	Error string `json:"error,nullable"`
	// Case.dev job ID
	JobID string `json:"jobId"`
	// Filename of the result document (if completed)
	ResultFilename string `json:"resultFilename,nullable"`
	// ID of the result document (if completed)
	ResultObjectID string `json:"resultObjectId,nullable"`
	// ID of the source document
	SourceObjectID string `json:"sourceObjectId"`
	// Current job status
	Status VaultObjectGetSummarizeJobResponseStatus `json:"status"`
	// Type of workflow being executed
	WorkflowType string                                 `json:"workflowType"`
	JSON         vaultObjectGetSummarizeJobResponseJSON `json:"-"`
}

// vaultObjectGetSummarizeJobResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetSummarizeJobResponse]
type vaultObjectGetSummarizeJobResponseJSON struct {
	CompletedAt    apijson.Field
	CreatedAt      apijson.Field
	Error          apijson.Field
	JobID          apijson.Field
	ResultFilename apijson.Field
	ResultObjectID apijson.Field
	SourceObjectID apijson.Field
	Status         apijson.Field
	WorkflowType   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultObjectGetSummarizeJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetSummarizeJobResponseJSON) RawJSON() string {
	return r.raw
}

// Current job status
type VaultObjectGetSummarizeJobResponseStatus string

const (
	VaultObjectGetSummarizeJobResponseStatusPending    VaultObjectGetSummarizeJobResponseStatus = "pending"
	VaultObjectGetSummarizeJobResponseStatusProcessing VaultObjectGetSummarizeJobResponseStatus = "processing"
	VaultObjectGetSummarizeJobResponseStatusCompleted  VaultObjectGetSummarizeJobResponseStatus = "completed"
	VaultObjectGetSummarizeJobResponseStatusFailed     VaultObjectGetSummarizeJobResponseStatus = "failed"
)

func (r VaultObjectGetSummarizeJobResponseStatus) IsKnown() bool {
	switch r {
	case VaultObjectGetSummarizeJobResponseStatusPending, VaultObjectGetSummarizeJobResponseStatusProcessing, VaultObjectGetSummarizeJobResponseStatusCompleted, VaultObjectGetSummarizeJobResponseStatusFailed:
		return true
	}
	return false
}

type VaultObjectGetTextResponse struct {
	Metadata VaultObjectGetTextResponseMetadata `json:"metadata,required"`
	// Full concatenated text content from all chunks
	Text string                         `json:"text,required"`
	JSON vaultObjectGetTextResponseJSON `json:"-"`
}

// vaultObjectGetTextResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetTextResponse]
type vaultObjectGetTextResponseJSON struct {
	Metadata    apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetTextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetTextResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetTextResponseMetadata struct {
	// Number of text chunks the document was split into
	ChunkCount int64 `json:"chunk_count,required"`
	// Original filename of the document
	Filename string `json:"filename,required"`
	// Total character count of the extracted text
	Length int64 `json:"length,required"`
	// The object ID
	ObjectID string `json:"object_id,required"`
	// The vault ID
	VaultID string `json:"vault_id,required"`
	// When the document processing completed
	IngestionCompletedAt time.Time                              `json:"ingestion_completed_at" format:"date-time"`
	JSON                 vaultObjectGetTextResponseMetadataJSON `json:"-"`
}

// vaultObjectGetTextResponseMetadataJSON contains the JSON metadata for the struct
// [VaultObjectGetTextResponseMetadata]
type vaultObjectGetTextResponseMetadataJSON struct {
	ChunkCount           apijson.Field
	Filename             apijson.Field
	Length               apijson.Field
	ObjectID             apijson.Field
	VaultID              apijson.Field
	IngestionCompletedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *VaultObjectGetTextResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetTextResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type VaultObjectUpdateParams struct {
	// New filename for the document (affects display name and downloads)
	Filename param.Field[string] `json:"filename"`
	// Additional metadata to merge with existing metadata
	Metadata param.Field[interface{}] `json:"metadata"`
	// Folder path for hierarchy preservation (e.g., '/Discovery/Depositions'). Set to
	// null or empty string to remove.
	Path param.Field[string] `json:"path"`
}

func (r VaultObjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultObjectDeleteParams struct {
	// Force delete a stuck document that is still in 'processing' state. Use this if a
	// document got stuck during ingestion (e.g., OCR timeout).
	Force param.Field[VaultObjectDeleteParamsForce] `query:"force"`
}

// URLQuery serializes [VaultObjectDeleteParams]'s query parameters as
// `url.Values`.
func (r VaultObjectDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Force delete a stuck document that is still in 'processing' state. Use this if a
// document got stuck during ingestion (e.g., OCR timeout).
type VaultObjectDeleteParamsForce string

const (
	VaultObjectDeleteParamsForceTrue VaultObjectDeleteParamsForce = "true"
)

func (r VaultObjectDeleteParamsForce) IsKnown() bool {
	switch r {
	case VaultObjectDeleteParamsForceTrue:
		return true
	}
	return false
}

type VaultObjectNewPresignedURLParams struct {
	// Content type for PUT operations (optional, defaults to object's content type)
	ContentType param.Field[string] `json:"contentType"`
	// URL expiration time in seconds (1 minute to 7 days)
	ExpiresIn param.Field[int64] `json:"expiresIn"`
	// The S3 operation to generate URL for
	Operation param.Field[VaultObjectNewPresignedURLParamsOperation] `json:"operation"`
	// File size in bytes (optional, max 5GB for single PUT uploads). When provided for
	// PUT operations, enforces exact file size at S3 level.
	SizeBytes param.Field[int64] `json:"sizeBytes"`
}

func (r VaultObjectNewPresignedURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The S3 operation to generate URL for
type VaultObjectNewPresignedURLParamsOperation string

const (
	VaultObjectNewPresignedURLParamsOperationGet    VaultObjectNewPresignedURLParamsOperation = "GET"
	VaultObjectNewPresignedURLParamsOperationPut    VaultObjectNewPresignedURLParamsOperation = "PUT"
	VaultObjectNewPresignedURLParamsOperationDelete VaultObjectNewPresignedURLParamsOperation = "DELETE"
	VaultObjectNewPresignedURLParamsOperationHead   VaultObjectNewPresignedURLParamsOperation = "HEAD"
)

func (r VaultObjectNewPresignedURLParamsOperation) IsKnown() bool {
	switch r {
	case VaultObjectNewPresignedURLParamsOperationGet, VaultObjectNewPresignedURLParamsOperationPut, VaultObjectNewPresignedURLParamsOperationDelete, VaultObjectNewPresignedURLParamsOperationHead:
		return true
	}
	return false
}

type VaultObjectGetOcrWordsParams struct {
	// Filter to a specific page number (1-indexed). If omitted, returns all pages.
	Page param.Field[int64] `query:"page"`
	// Filter to words ending at this index (inclusive). Useful for retrieving words
	// for a specific chunk.
	WordEnd param.Field[int64] `query:"wordEnd"`
	// Filter to words starting at this index (inclusive). Useful for retrieving words
	// for a specific chunk.
	WordStart param.Field[int64] `query:"wordStart"`
}

// URLQuery serializes [VaultObjectGetOcrWordsParams]'s query parameters as
// `url.Values`.
func (r VaultObjectGetOcrWordsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
