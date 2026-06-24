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

// Secure document storage with semantic search and GraphRAG
//
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
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a document's filename, path, or metadata. Use this to rename files or
// organize them into virtual folders. The path is stored in metadata.path and can
// be used to build folder hierarchies in your application.
func (r *VaultObjectService) Update(ctx context.Context, id string, objectID string, body VaultObjectUpdateParams, opts ...option.RequestOption) (res *VaultObjectUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve all objects stored in a specific vault, including document metadata,
// ingestion status, and processing statistics.
func (r *VaultObjectService) List(ctx context.Context, id string, query VaultObjectListParams, opts ...option.RequestOption) (res *VaultObjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a document from the vault including all associated vectors,
// chunks, graph data, and the original file. This operation cannot be undone.
func (r *VaultObjectService) Delete(ctx context.Context, id string, objectID string, body VaultObjectDeleteParams, opts ...option.RequestOption) (res *VaultObjectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Merges one or more PDF vault objects onto the end of an existing PDF vault
// object, overwriting the target in place before returning. Optionally rewrites
// citation links in the original target into internal PDF jumps and adds back
// links on appended pages. The target object’s ingestion state is not affected;
// appended pages are not searchable.
func (r *VaultObjectService) Append(ctx context.Context, id string, objectID string, body VaultObjectAppendParams, opts ...option.RequestOption) (res *VaultObjectAppendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/append", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generate presigned URLs for direct S3 operations (GET, PUT, DELETE, HEAD) on
// vault objects. This allows secure, time-limited access to files without proxying
// through the API. Essential for large document uploads/downloads in legal
// workflows.
func (r *VaultObjectService) NewPresignedURL(ctx context.Context, id string, objectID string, body VaultObjectNewPresignedURLParams, opts ...option.RequestOption) (res *VaultObjectNewPresignedURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/presigned-url", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Downloads a file from a vault by redirecting to a short-lived presigned S3 URL.
// Useful for retrieving contracts, depositions, case files, and other legal
// documents stored in your vault.
func (r *VaultObjectService) Download(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/download", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves full extracted chunk text for a processed vault object. Use this after
// search when a truncated preview is not enough and you need the exact chunk text
// or adjacent chunks for surrounding context such as tables, exhibit lists, or
// multi-part passages.
func (r *VaultObjectService) GetChunks(ctx context.Context, id string, objectID string, query VaultObjectGetChunksParams, opts ...option.RequestOption) (res *VaultObjectGetChunksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/chunks", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves word-level OCR bounding box data for a processed PDF document. Each
// word includes its text, normalized bounding box coordinates (0-1 range),
// confidence score, and global word index. Use this data to highlight specific
// text ranges in a PDF viewer based on word indices from search results.
func (r *VaultObjectService) GetOcrWords(ctx context.Context, id string, objectID string, query VaultObjectGetOcrWordsParams, opts ...option.RequestOption) (res *VaultObjectGetOcrWordsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/ocr-words", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves the raw text of a processed vault object split by page. The object
// must have completed ingestion before pages can be retrieved — for PDFs this
// requires the OCR pipeline to have finished writing the per-page sidecar, so
// freshly uploaded PDFs return 400 with the current `ingestionStatus` until
// processing completes. For PDFs this returns the per-page OCR text. For plain
// text files (txt, md, source code, court reporter transcripts) the text is split
// using right-aligned page-number markers when present (preserving the original
// document numbering, including continuations like Volume 2 starting at page 234),
// falling back to form-feed (\f) page-break characters, and finally a single page
// if neither signal is present. Use the optional `start` and `end` query
// parameters to fetch a specific inclusive page range. Pages with no text are
// omitted.
func (r *VaultObjectService) GetPages(ctx context.Context, id string, objectID string, query VaultObjectGetPagesParams, opts ...option.RequestOption) (res *VaultObjectGetPagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/pages", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the status of a CaseMark summary workflow job.
func (r *VaultObjectService) GetSummarizeJob(ctx context.Context, id string, objectID string, jobID string, opts ...option.RequestOption) (res *VaultObjectGetSummarizeJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/summarize/%s", id, objectID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the full extracted text content from a processed vault object. Returns
// the concatenated text from all chunks, useful for document review, analysis, or
// export. The object must have completed processing before text can be retrieved.
func (r *VaultObjectService) GetText(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *VaultObjectGetTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/text", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Triggers a CaseMark AI workflow to summarize or analyze a document stored in the
// vault. The workflow processes the document asynchronously and stores the result
// as a new object in the same vault, linked to the original document.
func (r *VaultObjectService) Summarize(ctx context.Context, id string, objectID string, body VaultObjectSummarizeParams, opts ...option.RequestOption) (res *VaultObjectSummarizeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/objects/%s/summarize", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type VaultObjectGetResponse struct {
	// Object ID
	ID string `json:"id" api:"required"`
	// MIME type
	ContentType string `json:"contentType" api:"required"`
	// Upload timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Presigned S3 download URL
	DownloadURL string `json:"downloadUrl" api:"required"`
	// URL expiration time in seconds
	ExpiresIn int64 `json:"expiresIn" api:"required"`
	// Original filename
	Filename string `json:"filename" api:"required"`
	// Processing status (pending, processing, completed, failed)
	IngestionStatus string `json:"ingestionStatus" api:"required"`
	// Vault ID
	VaultID string `json:"vaultId" api:"required"`
	// Number of text chunks created
	ChunkCount int64 `json:"chunkCount"`
	// Error details when ingestion fails
	IngestionError string `json:"ingestionError" api:"nullable"`
	// Whether the file was marked as AI-generated work product at upload time
	IsAIGenerated bool `json:"is_ai_generated"`
	// Additional metadata
	Metadata interface{} `json:"metadata"`
	// Number of pages (for documents)
	PageCount int64 `json:"pageCount"`
	// Optional folder path for hierarchy preservation
	Path string `json:"path" api:"nullable"`
	// File size in bytes
	SizeBytes int64 `json:"sizeBytes"`
	// Length of extracted text
	TextLength int64 `json:"textLength"`
	// Object ID of the completed transcript (if available)
	TranscriptObjectID string `json:"transcript_object_id" api:"nullable"`
	// Number of embedding vectors generated
	VectorCount int64                      `json:"vectorCount"`
	JSON        vaultObjectGetResponseJSON `json:"-"`
}

// vaultObjectGetResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetResponse]
type vaultObjectGetResponseJSON struct {
	ID                 apijson.Field
	ContentType        apijson.Field
	CreatedAt          apijson.Field
	DownloadURL        apijson.Field
	ExpiresIn          apijson.Field
	Filename           apijson.Field
	IngestionStatus    apijson.Field
	VaultID            apijson.Field
	ChunkCount         apijson.Field
	IngestionError     apijson.Field
	IsAIGenerated      apijson.Field
	Metadata           apijson.Field
	PageCount          apijson.Field
	Path               apijson.Field
	SizeBytes          apijson.Field
	TextLength         apijson.Field
	TranscriptObjectID apijson.Field
	VectorCount        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	Path string `json:"path" api:"nullable"`
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
	Count   float64                         `json:"count" api:"required"`
	Objects []VaultObjectListResponseObject `json:"objects" api:"required"`
	// The ID of the vault
	VaultID string                      `json:"vaultId" api:"required"`
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
	ID string `json:"id" api:"required"`
	// MIME type of the document
	ContentType string `json:"contentType" api:"required"`
	// Document upload timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Original filename of the uploaded document
	Filename string `json:"filename" api:"required"`
	// Processing status of the document
	IngestionStatus string `json:"ingestionStatus" api:"required"`
	// Number of text chunks created for vectorization
	ChunkCount float64 `json:"chunkCount"`
	// Processing completion timestamp
	IngestionCompletedAt time.Time `json:"ingestionCompletedAt" format:"date-time"`
	// Failure reason when ingestion status is a failed state
	IngestionError string `json:"ingestionError" api:"nullable"`
	// When ingestion processing began
	IngestionStartedAt time.Time `json:"ingestionStartedAt" api:"nullable" format:"date-time"`
	// Durable workflow run ID for the active or last ingestion attempt
	IngestionWorkflowID string `json:"ingestionWorkflowId" api:"nullable"`
	// Whether the file was marked as AI-generated work product at upload time
	IsAIGenerated bool `json:"is_ai_generated"`
	// Custom metadata associated with the document
	Metadata interface{} `json:"metadata"`
	// Number of pages in the document
	PageCount float64 `json:"pageCount"`
	// Optional folder path for hierarchy preservation from source systems
	Path string `json:"path" api:"nullable"`
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
	IngestionError       apijson.Field
	IngestionStartedAt   apijson.Field
	IngestionWorkflowID  apijson.Field
	IsAIGenerated        apijson.Field
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

type VaultObjectAppendResponse struct {
	ID              string                        `json:"id"`
	Checksum        string                        `json:"checksum"`
	ContentType     string                        `json:"contentType"`
	CreatedAt       time.Time                     `json:"createdAt" format:"date-time"`
	DownloadURL     string                        `json:"downloadUrl"`
	ExpiresIn       int64                         `json:"expiresIn"`
	Filename        string                        `json:"filename"`
	IngestionStatus string                        `json:"ingestionStatus"`
	Metadata        interface{}                   `json:"metadata"`
	PageCount       int64                         `json:"pageCount"`
	SizeBytes       int64                         `json:"sizeBytes"`
	VaultID         string                        `json:"vaultId"`
	JSON            vaultObjectAppendResponseJSON `json:"-"`
}

// vaultObjectAppendResponseJSON contains the JSON metadata for the struct
// [VaultObjectAppendResponse]
type vaultObjectAppendResponseJSON struct {
	ID              apijson.Field
	Checksum        apijson.Field
	ContentType     apijson.Field
	CreatedAt       apijson.Field
	DownloadURL     apijson.Field
	ExpiresIn       apijson.Field
	Filename        apijson.Field
	IngestionStatus apijson.Field
	Metadata        apijson.Field
	PageCount       apijson.Field
	SizeBytes       apijson.Field
	VaultID         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaultObjectAppendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectAppendResponseJSON) RawJSON() string {
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

type VaultObjectGetChunksResponse struct {
	// Full chunk objects for the requested range
	Chunks []VaultObjectGetChunksResponseChunk `json:"chunks" api:"required"`
	// The object ID
	ObjectID string `json:"object_id" api:"required"`
	// Total number of chunks stored for the object
	TotalChunks int64 `json:"total_chunks" api:"required"`
	// The vault ID
	VaultID string                           `json:"vault_id" api:"required"`
	JSON    vaultObjectGetChunksResponseJSON `json:"-"`
}

// vaultObjectGetChunksResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetChunksResponse]
type vaultObjectGetChunksResponseJSON struct {
	Chunks      apijson.Field
	ObjectID    apijson.Field
	TotalChunks apijson.Field
	VaultID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetChunksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetChunksResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetChunksResponseChunk struct {
	// Chunk index within the document
	Index int64 `json:"index" api:"required"`
	// Last page covered by the chunk, if page mapping is available
	PageEnd int64 `json:"page_end" api:"required,nullable"`
	// First page covered by the chunk, if page mapping is available
	PageStart int64 `json:"page_start" api:"required,nullable"`
	// Full text for the chunk
	Text string `json:"text" api:"required"`
	// Last OCR word index covered by the chunk, if available
	WordEndIndex int64 `json:"word_end_index" api:"required,nullable"`
	// First OCR word index covered by the chunk, if available
	WordStartIndex int64 `json:"word_start_index" api:"required,nullable"`
	// Source media timestamp for the last word in the chunk. Present only for
	// media-backed transcripts with real word timing.
	EndMs int64 `json:"end_ms"`
	// Source media timestamp for the first word in the chunk. Present only for
	// media-backed transcripts with real word timing.
	StartMs int64                                 `json:"start_ms"`
	JSON    vaultObjectGetChunksResponseChunkJSON `json:"-"`
}

// vaultObjectGetChunksResponseChunkJSON contains the JSON metadata for the struct
// [VaultObjectGetChunksResponseChunk]
type vaultObjectGetChunksResponseChunkJSON struct {
	Index          apijson.Field
	PageEnd        apijson.Field
	PageStart      apijson.Field
	Text           apijson.Field
	WordEndIndex   apijson.Field
	WordStartIndex apijson.Field
	EndMs          apijson.Field
	StartMs        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultObjectGetChunksResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetChunksResponseChunkJSON) RawJSON() string {
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
	Confidence float64 `json:"confidence" api:"nullable"`
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

type VaultObjectGetPagesResponse struct {
	Metadata VaultObjectGetPagesResponseMetadata `json:"metadata" api:"required"`
	// Per-page OCR text in ascending page order
	Pages []VaultObjectGetPagesResponsePage `json:"pages" api:"required"`
	JSON  vaultObjectGetPagesResponseJSON   `json:"-"`
}

// vaultObjectGetPagesResponseJSON contains the JSON metadata for the struct
// [VaultObjectGetPagesResponse]
type vaultObjectGetPagesResponseJSON struct {
	Metadata    apijson.Field
	Pages       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetPagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetPagesResponseJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetPagesResponseMetadata struct {
	Filename string `json:"filename" api:"required"`
	ObjectID string `json:"object_id" api:"required"`
	// Total number of pages with extracted text in the document
	PageCount int64 `json:"page_count" api:"required"`
	// Number of pages returned after applying the range filter
	ReturnedPages int64 `json:"returned_pages" api:"required"`
	// Where the page text came from. `ocr` for PDFs (per-page OCR sidecar). `txt` for
	// plain-text files split on form-feed (\f) characters.
	Source  VaultObjectGetPagesResponseMetadataSource `json:"source" api:"required"`
	VaultID string                                    `json:"vault_id" api:"required"`
	// Echoes the end query param if provided
	End int64 `json:"end" api:"nullable"`
	// Echoes the start query param if provided
	Start int64                                   `json:"start" api:"nullable"`
	JSON  vaultObjectGetPagesResponseMetadataJSON `json:"-"`
}

// vaultObjectGetPagesResponseMetadataJSON contains the JSON metadata for the
// struct [VaultObjectGetPagesResponseMetadata]
type vaultObjectGetPagesResponseMetadataJSON struct {
	Filename      apijson.Field
	ObjectID      apijson.Field
	PageCount     apijson.Field
	ReturnedPages apijson.Field
	Source        apijson.Field
	VaultID       apijson.Field
	End           apijson.Field
	Start         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultObjectGetPagesResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetPagesResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// Where the page text came from. `ocr` for PDFs (per-page OCR sidecar). `txt` for
// plain-text files split on form-feed (\f) characters.
type VaultObjectGetPagesResponseMetadataSource string

const (
	VaultObjectGetPagesResponseMetadataSourceOcr VaultObjectGetPagesResponseMetadataSource = "ocr"
	VaultObjectGetPagesResponseMetadataSourceTxt VaultObjectGetPagesResponseMetadataSource = "txt"
)

func (r VaultObjectGetPagesResponseMetadataSource) IsKnown() bool {
	switch r {
	case VaultObjectGetPagesResponseMetadataSourceOcr, VaultObjectGetPagesResponseMetadataSourceTxt:
		return true
	}
	return false
}

type VaultObjectGetPagesResponsePage struct {
	// Page number (1-indexed)
	Page int64 `json:"page" api:"required"`
	// OCR text for this page
	Text string                              `json:"text" api:"required"`
	JSON vaultObjectGetPagesResponsePageJSON `json:"-"`
}

// vaultObjectGetPagesResponsePageJSON contains the JSON metadata for the struct
// [VaultObjectGetPagesResponsePage]
type vaultObjectGetPagesResponsePageJSON struct {
	Page        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultObjectGetPagesResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectGetPagesResponsePageJSON) RawJSON() string {
	return r.raw
}

type VaultObjectGetSummarizeJobResponse struct {
	// When the job completed
	CompletedAt time.Time `json:"completedAt" api:"nullable" format:"date-time"`
	// When the job was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Error message (if failed)
	Error string `json:"error" api:"nullable"`
	// Case.dev job ID
	JobID string `json:"jobId"`
	// Filename of the result document (if completed)
	ResultFilename string `json:"resultFilename" api:"nullable"`
	// ID of the result document (if completed)
	ResultObjectID string `json:"resultObjectId" api:"nullable"`
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
	Metadata VaultObjectGetTextResponseMetadata `json:"metadata" api:"required"`
	// Full concatenated text content from all chunks
	Text string                         `json:"text" api:"required"`
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
	ChunkCount int64 `json:"chunk_count" api:"required"`
	// Original filename of the document
	Filename string `json:"filename" api:"required"`
	// Total character count of the extracted text
	Length int64 `json:"length" api:"required"`
	// The object ID
	ObjectID string `json:"object_id" api:"required"`
	// The vault ID
	VaultID string `json:"vault_id" api:"required"`
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

type VaultObjectSummarizeResponse struct {
	// CaseMark workflow ID
	CasemarkWorkflowID string `json:"casemarkWorkflowId"`
	// Case.dev job ID for tracking
	JobID string `json:"jobId"`
	// Current job status
	Status VaultObjectSummarizeResponseStatus `json:"status"`
	// URL to check job status
	StatusURL string `json:"statusUrl"`
	// Type of workflow being executed
	WorkflowType string                           `json:"workflowType"`
	JSON         vaultObjectSummarizeResponseJSON `json:"-"`
}

// vaultObjectSummarizeResponseJSON contains the JSON metadata for the struct
// [VaultObjectSummarizeResponse]
type vaultObjectSummarizeResponseJSON struct {
	CasemarkWorkflowID apijson.Field
	JobID              apijson.Field
	Status             apijson.Field
	StatusURL          apijson.Field
	WorkflowType       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VaultObjectSummarizeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultObjectSummarizeResponseJSON) RawJSON() string {
	return r.raw
}

// Current job status
type VaultObjectSummarizeResponseStatus string

const (
	VaultObjectSummarizeResponseStatusPending    VaultObjectSummarizeResponseStatus = "pending"
	VaultObjectSummarizeResponseStatusProcessing VaultObjectSummarizeResponseStatus = "processing"
	VaultObjectSummarizeResponseStatusCompleted  VaultObjectSummarizeResponseStatus = "completed"
	VaultObjectSummarizeResponseStatusFailed     VaultObjectSummarizeResponseStatus = "failed"
)

func (r VaultObjectSummarizeResponseStatus) IsKnown() bool {
	switch r {
	case VaultObjectSummarizeResponseStatusPending, VaultObjectSummarizeResponseStatusProcessing, VaultObjectSummarizeResponseStatusCompleted, VaultObjectSummarizeResponseStatusFailed:
		return true
	}
	return false
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

type VaultObjectListParams struct {
	// Include placeholders for uploads that were never completed (awaiting_upload) or
	// were cancelled (aborted). Excluded by default.
	IncludeUnconfirmed param.Field[bool] `query:"includeUnconfirmed"`
}

// URLQuery serializes [VaultObjectListParams]'s query parameters as `url.Values`.
func (r VaultObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type VaultObjectAppendParams struct {
	// Vault object IDs whose pages will be appended onto the target object, in order.
	// Must not include the target object itself.
	AppendObjectIDs param.Field[[]string] `json:"appendObjectIds" api:"required"`
	// Adds back links on appended pages
	BackLinks param.Field[bool] `json:"backLinks"`
	// Label text for the back link. Used only when backLinks is true and rendered
	// centered at the bottom of each appended page.
	BackLinksText param.Field[string] `json:"backLinksText"`
	// When true, rewrites links in the target object to internal PDF jumps when the
	// URL contains exactly one appended object ID as a standalone query parameter
	// value or decoded path segment.
	RewriteLinks param.Field[bool] `json:"rewriteLinks"`
}

func (r VaultObjectAppendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type VaultObjectGetChunksParams struct {
	// The last chunk index to return (inclusive). If omitted, only the `start` chunk
	// is returned. Ranges are limited to 10 chunks.
	End param.Field[int64] `query:"end"`
	// The first chunk index to return (0-based). Defaults to 0.
	Start param.Field[int64] `query:"start"`
}

// URLQuery serializes [VaultObjectGetChunksParams]'s query parameters as
// `url.Values`.
func (r VaultObjectGetChunksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type VaultObjectGetPagesParams struct {
	// Last page to return (inclusive, 1-indexed). If omitted, returns through the last
	// page with text.
	End param.Field[int64] `query:"end"`
	// First page to return (inclusive, 1-indexed). If omitted, starts at the first
	// page with text.
	Start param.Field[int64] `query:"start"`
}

// URLQuery serializes [VaultObjectGetPagesParams]'s query parameters as
// `url.Values`.
func (r VaultObjectGetPagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VaultObjectSummarizeParams struct {
	// Output format for the summary document
	OutputFormat param.Field[VaultObjectSummarizeParamsOutputFormat] `json:"outputFormat"`
	// Type of CaseMark workflow to run
	WorkflowType param.Field[string] `json:"workflowType"`
}

func (r VaultObjectSummarizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output format for the summary document
type VaultObjectSummarizeParamsOutputFormat string

const (
	VaultObjectSummarizeParamsOutputFormatPdf  VaultObjectSummarizeParamsOutputFormat = "PDF"
	VaultObjectSummarizeParamsOutputFormatWord VaultObjectSummarizeParamsOutputFormat = "WORD"
)

func (r VaultObjectSummarizeParamsOutputFormat) IsKnown() bool {
	switch r {
	case VaultObjectSummarizeParamsOutputFormatPdf, VaultObjectSummarizeParamsOutputFormatWord:
		return true
	}
	return false
}
