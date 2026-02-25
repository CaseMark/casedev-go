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

// VaultGraphragService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultGraphragService] method instead.
type VaultGraphragService struct {
	Options []option.RequestOption
}

// NewVaultGraphragService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultGraphragService(opts ...option.RequestOption) (r *VaultGraphragService) {
	r = &VaultGraphragService{}
	r.Options = opts
	return
}

// Retrieve GraphRAG (Graph Retrieval-Augmented Generation) statistics for a
// specific vault. This includes metrics about the knowledge graph structure,
// entity relationships, and processing status that enable advanced semantic search
// and AI-powered document analysis.
func (r *VaultGraphragService) GetStats(ctx context.Context, id string, opts ...option.RequestOption) (res *VaultGraphragGetStatsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/graphrag/stats", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Initialize a GraphRAG workspace for a vault to enable advanced knowledge graph
// and retrieval-augmented generation capabilities. This creates the necessary
// infrastructure for semantic document analysis and graph-based querying within
// the vault.
func (r *VaultGraphragService) Init(ctx context.Context, id string, opts ...option.RequestOption) (res *VaultGraphragInitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/graphrag/init", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Manually trigger GraphRAG indexing for a vault object. The object must already
// be ingested (completed status). This extracts entities, relationships, and
// communities from the document for advanced knowledge graph queries.
func (r *VaultGraphragService) ProcessObject(ctx context.Context, id string, objectID string, opts ...option.RequestOption) (res *VaultGraphragProcessObjectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/graphrag/%s", id, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type VaultGraphragGetStatsResponse struct {
	// Number of entity communities identified
	Communities int64 `json:"communities"`
	// Number of processed documents
	Documents int64 `json:"documents"`
	// Total number of entities extracted from documents
	Entities int64 `json:"entities"`
	// Timestamp of last GraphRAG processing
	LastProcessed time.Time `json:"lastProcessed" format:"date-time"`
	// Total number of relationships between entities
	Relationships int64 `json:"relationships"`
	// Current processing status
	Status VaultGraphragGetStatsResponseStatus `json:"status"`
	JSON   vaultGraphragGetStatsResponseJSON   `json:"-"`
}

// vaultGraphragGetStatsResponseJSON contains the JSON metadata for the struct
// [VaultGraphragGetStatsResponse]
type vaultGraphragGetStatsResponseJSON struct {
	Communities   apijson.Field
	Documents     apijson.Field
	Entities      apijson.Field
	LastProcessed apijson.Field
	Relationships apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultGraphragGetStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultGraphragGetStatsResponseJSON) RawJSON() string {
	return r.raw
}

// Current processing status
type VaultGraphragGetStatsResponseStatus string

const (
	VaultGraphragGetStatsResponseStatusProcessing VaultGraphragGetStatsResponseStatus = "processing"
	VaultGraphragGetStatsResponseStatusCompleted  VaultGraphragGetStatsResponseStatus = "completed"
	VaultGraphragGetStatsResponseStatusError      VaultGraphragGetStatsResponseStatus = "error"
)

func (r VaultGraphragGetStatsResponseStatus) IsKnown() bool {
	switch r {
	case VaultGraphragGetStatsResponseStatusProcessing, VaultGraphragGetStatsResponseStatusCompleted, VaultGraphragGetStatsResponseStatusError:
		return true
	}
	return false
}

type VaultGraphragInitResponse struct {
	Message string                        `json:"message"`
	Status  string                        `json:"status"`
	Success bool                          `json:"success"`
	VaultID string                        `json:"vault_id"`
	JSON    vaultGraphragInitResponseJSON `json:"-"`
}

// vaultGraphragInitResponseJSON contains the JSON metadata for the struct
// [VaultGraphragInitResponse]
type vaultGraphragInitResponseJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	VaultID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultGraphragInitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultGraphragInitResponseJSON) RawJSON() string {
	return r.raw
}

type VaultGraphragProcessObjectResponse struct {
	// Number of communities detected
	Communities int64 `json:"communities" api:"required"`
	// Number of entities extracted
	Entities int64 `json:"entities" api:"required"`
	// ID of the indexed object
	ObjectID string `json:"objectId" api:"required"`
	// Number of relationships extracted
	Relationships int64 `json:"relationships" api:"required"`
	// Extraction statistics
	Stats VaultGraphragProcessObjectResponseStats `json:"stats" api:"required"`
	// Status from GraphRAG service
	Status string `json:"status" api:"required"`
	// Whether indexing completed successfully
	Success bool `json:"success" api:"required"`
	// ID of the vault
	VaultID string                                 `json:"vaultId" api:"required"`
	JSON    vaultGraphragProcessObjectResponseJSON `json:"-"`
}

// vaultGraphragProcessObjectResponseJSON contains the JSON metadata for the struct
// [VaultGraphragProcessObjectResponse]
type vaultGraphragProcessObjectResponseJSON struct {
	Communities   apijson.Field
	Entities      apijson.Field
	ObjectID      apijson.Field
	Relationships apijson.Field
	Stats         apijson.Field
	Status        apijson.Field
	Success       apijson.Field
	VaultID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VaultGraphragProcessObjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultGraphragProcessObjectResponseJSON) RawJSON() string {
	return r.raw
}

// Extraction statistics
type VaultGraphragProcessObjectResponseStats struct {
	CommunityCount    int64                                       `json:"community_count"`
	EntityCount       int64                                       `json:"entity_count"`
	RelationshipCount int64                                       `json:"relationship_count"`
	JSON              vaultGraphragProcessObjectResponseStatsJSON `json:"-"`
}

// vaultGraphragProcessObjectResponseStatsJSON contains the JSON metadata for the
// struct [VaultGraphragProcessObjectResponseStats]
type vaultGraphragProcessObjectResponseStatsJSON struct {
	CommunityCount    apijson.Field
	EntityCount       apijson.Field
	RelationshipCount apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *VaultGraphragProcessObjectResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultGraphragProcessObjectResponseStatsJSON) RawJSON() string {
	return r.raw
}
