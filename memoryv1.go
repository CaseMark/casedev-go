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

// MemoryV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemoryV1Service] method instead.
type MemoryV1Service struct {
	Options []option.RequestOption
}

// NewMemoryV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMemoryV1Service(opts ...option.RequestOption) (r *MemoryV1Service) {
	r = &MemoryV1Service{}
	r.Options = opts
	return
}

// Store memories from conversation messages. Automatically extracts facts and
// handles deduplication.
//
// Use tag_1 through tag_12 for filtering - these are generic indexed fields you
// can use for any purpose:
//
// - Legal app: tag_1=client_id, tag_2=matter_id
// - Healthcare: tag_1=patient_id, tag_2=encounter_id
// - E-commerce: tag_1=customer_id, tag_2=order_id
func (r *MemoryV1Service) New(ctx context.Context, body MemoryV1NewParams, opts ...option.RequestOption) (res *MemoryV1NewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "memory/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a single memory by its ID.
func (r *MemoryV1Service) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MemoryV1GetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("memory/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all memories with optional filtering by tags and category.
func (r *MemoryV1Service) List(ctx context.Context, query MemoryV1ListParams, opts ...option.RequestOption) (res *MemoryV1ListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "memory/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a single memory by its ID.
func (r *MemoryV1Service) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MemoryV1DeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("memory/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete multiple memories matching tag filter criteria. CAUTION: This will delete
// all matching memories for your organization.
func (r *MemoryV1Service) DeleteAll(ctx context.Context, body MemoryV1DeleteAllParams, opts ...option.RequestOption) (res *MemoryV1DeleteAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "memory/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Search memories using semantic similarity. Filter by tag fields to narrow
// results.
//
// Use tag_1 through tag_12 for filtering - these are generic indexed fields you
// define:
//
// - Legal app: tag_1=client_id, tag_2=matter_id
// - Healthcare: tag_1=patient_id, tag_2=encounter_id
func (r *MemoryV1Service) Search(ctx context.Context, body MemoryV1SearchParams, opts ...option.RequestOption) (res *MemoryV1SearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "memory/v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MemoryV1NewResponse struct {
	Results []MemoryV1NewResponseResult `json:"results"`
	JSON    memoryV1NewResponseJSON     `json:"-"`
}

// memoryV1NewResponseJSON contains the JSON metadata for the struct
// [MemoryV1NewResponse]
type memoryV1NewResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1NewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1NewResponseJSON) RawJSON() string {
	return r.raw
}

type MemoryV1NewResponseResult struct {
	// Memory ID
	ID string `json:"id"`
	// What happened to this memory
	Event MemoryV1NewResponseResultsEvent `json:"event"`
	// Extracted memory text
	Memory string                        `json:"memory"`
	JSON   memoryV1NewResponseResultJSON `json:"-"`
}

// memoryV1NewResponseResultJSON contains the JSON metadata for the struct
// [MemoryV1NewResponseResult]
type memoryV1NewResponseResultJSON struct {
	ID          apijson.Field
	Event       apijson.Field
	Memory      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1NewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1NewResponseResultJSON) RawJSON() string {
	return r.raw
}

// What happened to this memory
type MemoryV1NewResponseResultsEvent string

const (
	MemoryV1NewResponseResultsEventAdd    MemoryV1NewResponseResultsEvent = "ADD"
	MemoryV1NewResponseResultsEventUpdate MemoryV1NewResponseResultsEvent = "UPDATE"
	MemoryV1NewResponseResultsEventDelete MemoryV1NewResponseResultsEvent = "DELETE"
	MemoryV1NewResponseResultsEventNone   MemoryV1NewResponseResultsEvent = "NONE"
)

func (r MemoryV1NewResponseResultsEvent) IsKnown() bool {
	switch r {
	case MemoryV1NewResponseResultsEventAdd, MemoryV1NewResponseResultsEventUpdate, MemoryV1NewResponseResultsEventDelete, MemoryV1NewResponseResultsEventNone:
		return true
	}
	return false
}

type MemoryV1GetResponse struct {
	// Memory ID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Memory content
	Memory string `json:"memory"`
	// Memory metadata
	Metadata  interface{}             `json:"metadata"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	JSON      memoryV1GetResponseJSON `json:"-"`
}

// memoryV1GetResponseJSON contains the JSON metadata for the struct
// [MemoryV1GetResponse]
type memoryV1GetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Memory      apijson.Field
	Metadata    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1GetResponseJSON) RawJSON() string {
	return r.raw
}

type MemoryV1ListResponse struct {
	// Total count matching filters
	Count   int64                        `json:"count"`
	Results []MemoryV1ListResponseResult `json:"results"`
	JSON    memoryV1ListResponseJSON     `json:"-"`
}

// memoryV1ListResponseJSON contains the JSON metadata for the struct
// [MemoryV1ListResponse]
type memoryV1ListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1ListResponseJSON) RawJSON() string {
	return r.raw
}

type MemoryV1ListResponseResult struct {
	ID        string                         `json:"id"`
	CreatedAt time.Time                      `json:"created_at" format:"date-time"`
	Memory    string                         `json:"memory"`
	Metadata  interface{}                    `json:"metadata"`
	Tags      interface{}                    `json:"tags"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      memoryV1ListResponseResultJSON `json:"-"`
}

// memoryV1ListResponseResultJSON contains the JSON metadata for the struct
// [MemoryV1ListResponseResult]
type memoryV1ListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Memory      apijson.Field
	Metadata    apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1ListResponseResultJSON) RawJSON() string {
	return r.raw
}

type MemoryV1DeleteResponse struct {
	Message string                     `json:"message"`
	Success bool                       `json:"success"`
	JSON    memoryV1DeleteResponseJSON `json:"-"`
}

// memoryV1DeleteResponseJSON contains the JSON metadata for the struct
// [MemoryV1DeleteResponse]
type memoryV1DeleteResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1DeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1DeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MemoryV1DeleteAllResponse struct {
	// Number of memories deleted
	Deleted int64                         `json:"deleted"`
	JSON    memoryV1DeleteAllResponseJSON `json:"-"`
}

// memoryV1DeleteAllResponseJSON contains the JSON metadata for the struct
// [MemoryV1DeleteAllResponse]
type memoryV1DeleteAllResponseJSON struct {
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1DeleteAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1DeleteAllResponseJSON) RawJSON() string {
	return r.raw
}

type MemoryV1SearchResponse struct {
	Results []MemoryV1SearchResponseResult `json:"results"`
	JSON    memoryV1SearchResponseJSON     `json:"-"`
}

// memoryV1SearchResponseJSON contains the JSON metadata for the struct
// [MemoryV1SearchResponse]
type memoryV1SearchResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1SearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1SearchResponseJSON) RawJSON() string {
	return r.raw
}

type MemoryV1SearchResponseResult struct {
	// Memory ID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Memory content
	Memory string `json:"memory"`
	// Additional metadata
	Metadata interface{} `json:"metadata"`
	// Similarity score (0-1)
	Score float64 `json:"score"`
	// Tag values for this memory
	Tags      MemoryV1SearchResponseResultsTags `json:"tags"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      memoryV1SearchResponseResultJSON  `json:"-"`
}

// memoryV1SearchResponseResultJSON contains the JSON metadata for the struct
// [MemoryV1SearchResponseResult]
type memoryV1SearchResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Memory      apijson.Field
	Metadata    apijson.Field
	Score       apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1SearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1SearchResponseResultJSON) RawJSON() string {
	return r.raw
}

// Tag values for this memory
type MemoryV1SearchResponseResultsTags struct {
	Tag1  string                                `json:"tag_1"`
	Tag10 string                                `json:"tag_10"`
	Tag11 string                                `json:"tag_11"`
	Tag12 string                                `json:"tag_12"`
	Tag2  string                                `json:"tag_2"`
	Tag3  string                                `json:"tag_3"`
	Tag4  string                                `json:"tag_4"`
	Tag5  string                                `json:"tag_5"`
	Tag6  string                                `json:"tag_6"`
	Tag7  string                                `json:"tag_7"`
	Tag8  string                                `json:"tag_8"`
	Tag9  string                                `json:"tag_9"`
	JSON  memoryV1SearchResponseResultsTagsJSON `json:"-"`
}

// memoryV1SearchResponseResultsTagsJSON contains the JSON metadata for the struct
// [MemoryV1SearchResponseResultsTags]
type memoryV1SearchResponseResultsTagsJSON struct {
	Tag1        apijson.Field
	Tag10       apijson.Field
	Tag11       apijson.Field
	Tag12       apijson.Field
	Tag2        apijson.Field
	Tag3        apijson.Field
	Tag4        apijson.Field
	Tag5        apijson.Field
	Tag6        apijson.Field
	Tag7        apijson.Field
	Tag8        apijson.Field
	Tag9        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryV1SearchResponseResultsTags) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryV1SearchResponseResultsTagsJSON) RawJSON() string {
	return r.raw
}

type MemoryV1NewParams struct {
	// Conversation messages to extract memories from
	Messages param.Field[[]MemoryV1NewParamsMessage] `json:"messages,required"`
	// Custom category (e.g., "fact", "preference", "deadline")
	Category param.Field[string] `json:"category"`
	// Optional custom prompt for fact extraction
	ExtractionPrompt param.Field[string] `json:"extraction_prompt"`
	// Whether to extract facts from messages (default: true)
	Infer param.Field[bool] `json:"infer"`
	// Additional metadata (not indexed)
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Generic indexed filter field 1 (you decide what it means)
	Tag1 param.Field[string] `json:"tag_1"`
	// Generic indexed filter field 10
	Tag10 param.Field[string] `json:"tag_10"`
	// Generic indexed filter field 11
	Tag11 param.Field[string] `json:"tag_11"`
	// Generic indexed filter field 12
	Tag12 param.Field[string] `json:"tag_12"`
	// Generic indexed filter field 2
	Tag2 param.Field[string] `json:"tag_2"`
	// Generic indexed filter field 3
	Tag3 param.Field[string] `json:"tag_3"`
	// Generic indexed filter field 4
	Tag4 param.Field[string] `json:"tag_4"`
	// Generic indexed filter field 5
	Tag5 param.Field[string] `json:"tag_5"`
	// Generic indexed filter field 6
	Tag6 param.Field[string] `json:"tag_6"`
	// Generic indexed filter field 7
	Tag7 param.Field[string] `json:"tag_7"`
	// Generic indexed filter field 8
	Tag8 param.Field[string] `json:"tag_8"`
	// Generic indexed filter field 9
	Tag9 param.Field[string] `json:"tag_9"`
}

func (r MemoryV1NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryV1NewParamsMessage struct {
	// Message content
	Content param.Field[string] `json:"content,required"`
	// Message role
	Role param.Field[MemoryV1NewParamsMessagesRole] `json:"role,required"`
}

func (r MemoryV1NewParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Message role
type MemoryV1NewParamsMessagesRole string

const (
	MemoryV1NewParamsMessagesRoleUser      MemoryV1NewParamsMessagesRole = "user"
	MemoryV1NewParamsMessagesRoleAssistant MemoryV1NewParamsMessagesRole = "assistant"
	MemoryV1NewParamsMessagesRoleSystem    MemoryV1NewParamsMessagesRole = "system"
)

func (r MemoryV1NewParamsMessagesRole) IsKnown() bool {
	switch r {
	case MemoryV1NewParamsMessagesRoleUser, MemoryV1NewParamsMessagesRoleAssistant, MemoryV1NewParamsMessagesRoleSystem:
		return true
	}
	return false
}

type MemoryV1ListParams struct {
	// Filter by category
	Category param.Field[string] `query:"category"`
	// Number of results
	Limit param.Field[int64] `query:"limit"`
	// Pagination offset
	Offset param.Field[int64] `query:"offset"`
	// Filter by tag_1
	Tag1 param.Field[string] `query:"tag_1"`
	// Filter by tag_10
	Tag10 param.Field[string] `query:"tag_10"`
	// Filter by tag_11
	Tag11 param.Field[string] `query:"tag_11"`
	// Filter by tag_12
	Tag12 param.Field[string] `query:"tag_12"`
	// Filter by tag_2
	Tag2 param.Field[string] `query:"tag_2"`
	// Filter by tag_3
	Tag3 param.Field[string] `query:"tag_3"`
	// Filter by tag_4
	Tag4 param.Field[string] `query:"tag_4"`
	// Filter by tag_5
	Tag5 param.Field[string] `query:"tag_5"`
	// Filter by tag_6
	Tag6 param.Field[string] `query:"tag_6"`
	// Filter by tag_7
	Tag7 param.Field[string] `query:"tag_7"`
	// Filter by tag_8
	Tag8 param.Field[string] `query:"tag_8"`
	// Filter by tag_9
	Tag9 param.Field[string] `query:"tag_9"`
}

// URLQuery serializes [MemoryV1ListParams]'s query parameters as `url.Values`.
func (r MemoryV1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MemoryV1DeleteAllParams struct {
	// Filter by tag_1
	Tag1 param.Field[string] `query:"tag_1"`
	// Filter by tag_10
	Tag10 param.Field[string] `query:"tag_10"`
	// Filter by tag_11
	Tag11 param.Field[string] `query:"tag_11"`
	// Filter by tag_12
	Tag12 param.Field[string] `query:"tag_12"`
	// Filter by tag_2
	Tag2 param.Field[string] `query:"tag_2"`
	// Filter by tag_3
	Tag3 param.Field[string] `query:"tag_3"`
	// Filter by tag_4
	Tag4 param.Field[string] `query:"tag_4"`
	// Filter by tag_5
	Tag5 param.Field[string] `query:"tag_5"`
	// Filter by tag_6
	Tag6 param.Field[string] `query:"tag_6"`
	// Filter by tag_7
	Tag7 param.Field[string] `query:"tag_7"`
	// Filter by tag_8
	Tag8 param.Field[string] `query:"tag_8"`
	// Filter by tag_9
	Tag9 param.Field[string] `query:"tag_9"`
}

// URLQuery serializes [MemoryV1DeleteAllParams]'s query parameters as
// `url.Values`.
func (r MemoryV1DeleteAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MemoryV1SearchParams struct {
	// Search query for semantic matching
	Query param.Field[string] `json:"query,required"`
	// Filter by category
	Category param.Field[string] `json:"category"`
	// Filter by tag_1
	Tag1 param.Field[string] `json:"tag_1"`
	// Filter by tag_10
	Tag10 param.Field[string] `json:"tag_10"`
	// Filter by tag_11
	Tag11 param.Field[string] `json:"tag_11"`
	// Filter by tag_12
	Tag12 param.Field[string] `json:"tag_12"`
	// Filter by tag_2
	Tag2 param.Field[string] `json:"tag_2"`
	// Filter by tag_3
	Tag3 param.Field[string] `json:"tag_3"`
	// Filter by tag_4
	Tag4 param.Field[string] `json:"tag_4"`
	// Filter by tag_5
	Tag5 param.Field[string] `json:"tag_5"`
	// Filter by tag_6
	Tag6 param.Field[string] `json:"tag_6"`
	// Filter by tag_7
	Tag7 param.Field[string] `json:"tag_7"`
	// Filter by tag_8
	Tag8 param.Field[string] `json:"tag_8"`
	// Filter by tag_9
	Tag9 param.Field[string] `json:"tag_9"`
	// Maximum number of results to return
	TopK param.Field[int64] `json:"top_k"`
}

func (r MemoryV1SearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
