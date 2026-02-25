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

// SearchV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchV1Service] method instead.
type SearchV1Service struct {
	Options []option.RequestOption
}

// NewSearchV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSearchV1Service(opts ...option.RequestOption) (r *SearchV1Service) {
	r = &SearchV1Service{}
	r.Options = opts
	return
}

// Generate comprehensive answers to questions using web search results. Supports
// two modes: native provider answers or custom LLM-powered answers using
// Case.dev's AI gateway. Perfect for legal research, fact-checking, and gathering
// supporting evidence for cases.
func (r *SearchV1Service) Answer(ctx context.Context, body SearchV1AnswerParams, opts ...option.RequestOption) (res *SearchV1AnswerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "search/v1/answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Scrapes and extracts text content from web pages, PDFs, and documents. Useful
// for legal research, evidence collection, and document analysis. Supports live
// crawling, subpage extraction, and content summarization.
func (r *SearchV1Service) Contents(ctx context.Context, body SearchV1ContentsParams, opts ...option.RequestOption) (res *SearchV1ContentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "search/v1/contents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Performs deep research by conducting multi-step analysis, gathering information
// from multiple sources, and providing comprehensive insights. Ideal for legal
// research, case analysis, and due diligence investigations.
func (r *SearchV1Service) Research(ctx context.Context, body SearchV1ResearchParams, opts ...option.RequestOption) (res *SearchV1ResearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "search/v1/research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the status and results of a deep research task by ID. Supports both
// standard JSON responses and streaming for real-time updates as the research
// progresses. Research tasks analyze topics comprehensively using web search and
// AI synthesis.
func (r *SearchV1Service) GetResearch(ctx context.Context, id string, query SearchV1GetResearchParams, opts ...option.RequestOption) (res *SearchV1GetResearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("search/v1/research/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Executes intelligent web search queries with advanced filtering and
// customization options. Ideal for legal research, case law discovery, and
// gathering supporting documentation for litigation or compliance matters.
func (r *SearchV1Service) Search(ctx context.Context, body SearchV1SearchParams, opts ...option.RequestOption) (res *SearchV1SearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "search/v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find web pages and documents similar to a given URL. Useful for legal research
// to discover related case law, statutes, or legal commentary that shares similar
// themes or content structure.
func (r *SearchV1Service) Similar(ctx context.Context, body SearchV1SimilarParams, opts ...option.RequestOption) (res *SearchV1SimilarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "search/v1/similar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SearchV1AnswerResponse struct {
	// The generated answer with citations
	Answer string `json:"answer"`
	// Sources used to generate the answer
	Citations []SearchV1AnswerResponseCitation `json:"citations"`
	// Model used for answer generation
	Model string `json:"model"`
	// Type of search performed
	SearchType string                     `json:"searchType"`
	JSON       searchV1AnswerResponseJSON `json:"-"`
}

// searchV1AnswerResponseJSON contains the JSON metadata for the struct
// [SearchV1AnswerResponse]
type searchV1AnswerResponseJSON struct {
	Answer      apijson.Field
	Citations   apijson.Field
	Model       apijson.Field
	SearchType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1AnswerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1AnswerResponseJSON) RawJSON() string {
	return r.raw
}

type SearchV1AnswerResponseCitation struct {
	ID            string                             `json:"id"`
	PublishedDate string                             `json:"publishedDate"`
	Text          string                             `json:"text"`
	Title         string                             `json:"title"`
	URL           string                             `json:"url"`
	JSON          searchV1AnswerResponseCitationJSON `json:"-"`
}

// searchV1AnswerResponseCitationJSON contains the JSON metadata for the struct
// [SearchV1AnswerResponseCitation]
type searchV1AnswerResponseCitationJSON struct {
	ID            apijson.Field
	PublishedDate apijson.Field
	Text          apijson.Field
	Title         apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SearchV1AnswerResponseCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1AnswerResponseCitationJSON) RawJSON() string {
	return r.raw
}

type SearchV1ContentsResponse struct {
	Results []SearchV1ContentsResponseResult `json:"results"`
	JSON    searchV1ContentsResponseJSON     `json:"-"`
}

// searchV1ContentsResponseJSON contains the JSON metadata for the struct
// [SearchV1ContentsResponse]
type searchV1ContentsResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1ContentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1ContentsResponseJSON) RawJSON() string {
	return r.raw
}

type SearchV1ContentsResponseResult struct {
	// Content highlights if requested
	Highlights []string `json:"highlights"`
	// Additional metadata about the content
	Metadata interface{} `json:"metadata"`
	// Content summary if requested
	Summary string `json:"summary"`
	// Extracted text content
	Text string `json:"text"`
	// Page title
	Title string `json:"title"`
	// Source URL
	URL  string                             `json:"url"`
	JSON searchV1ContentsResponseResultJSON `json:"-"`
}

// searchV1ContentsResponseResultJSON contains the JSON metadata for the struct
// [SearchV1ContentsResponseResult]
type searchV1ContentsResponseResultJSON struct {
	Highlights  apijson.Field
	Metadata    apijson.Field
	Summary     apijson.Field
	Text        apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1ContentsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1ContentsResponseResultJSON) RawJSON() string {
	return r.raw
}

type SearchV1ResearchResponse struct {
	// Model used for research
	Model string `json:"model"`
	// Unique identifier for this research
	ResearchID string `json:"researchId"`
	// Research findings and analysis
	Results interface{}                  `json:"results"`
	JSON    searchV1ResearchResponseJSON `json:"-"`
}

// searchV1ResearchResponseJSON contains the JSON metadata for the struct
// [SearchV1ResearchResponse]
type searchV1ResearchResponseJSON struct {
	Model       apijson.Field
	ResearchID  apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1ResearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1ResearchResponseJSON) RawJSON() string {
	return r.raw
}

type SearchV1GetResearchResponse struct {
	// Research task ID
	ID string `json:"id"`
	// Task completion timestamp
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Task creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Research model used
	Model SearchV1GetResearchResponseModel `json:"model"`
	// Completion percentage (0-100)
	Progress float64 `json:"progress"`
	// Original research query
	Query string `json:"query"`
	// Research findings and analysis
	Results SearchV1GetResearchResponseResults `json:"results"`
	// Current status of the research task
	Status SearchV1GetResearchResponseStatus `json:"status"`
	JSON   searchV1GetResearchResponseJSON   `json:"-"`
}

// searchV1GetResearchResponseJSON contains the JSON metadata for the struct
// [SearchV1GetResearchResponse]
type searchV1GetResearchResponseJSON struct {
	ID          apijson.Field
	CompletedAt apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	Progress    apijson.Field
	Query       apijson.Field
	Results     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1GetResearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1GetResearchResponseJSON) RawJSON() string {
	return r.raw
}

// Research model used
type SearchV1GetResearchResponseModel string

const (
	SearchV1GetResearchResponseModelFast   SearchV1GetResearchResponseModel = "fast"
	SearchV1GetResearchResponseModelNormal SearchV1GetResearchResponseModel = "normal"
	SearchV1GetResearchResponseModelPro    SearchV1GetResearchResponseModel = "pro"
)

func (r SearchV1GetResearchResponseModel) IsKnown() bool {
	switch r {
	case SearchV1GetResearchResponseModelFast, SearchV1GetResearchResponseModelNormal, SearchV1GetResearchResponseModelPro:
		return true
	}
	return false
}

// Research findings and analysis
type SearchV1GetResearchResponseResults struct {
	// Detailed research sections
	Sections []SearchV1GetResearchResponseResultsSection `json:"sections"`
	// All sources referenced in research
	Sources []SearchV1GetResearchResponseResultsSource `json:"sources"`
	// Executive summary of research findings
	Summary string                                 `json:"summary"`
	JSON    searchV1GetResearchResponseResultsJSON `json:"-"`
}

// searchV1GetResearchResponseResultsJSON contains the JSON metadata for the struct
// [SearchV1GetResearchResponseResults]
type searchV1GetResearchResponseResultsJSON struct {
	Sections    apijson.Field
	Sources     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1GetResearchResponseResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1GetResearchResponseResultsJSON) RawJSON() string {
	return r.raw
}

type SearchV1GetResearchResponseResultsSection struct {
	Content string                                             `json:"content"`
	Sources []SearchV1GetResearchResponseResultsSectionsSource `json:"sources"`
	Title   string                                             `json:"title"`
	JSON    searchV1GetResearchResponseResultsSectionJSON      `json:"-"`
}

// searchV1GetResearchResponseResultsSectionJSON contains the JSON metadata for the
// struct [SearchV1GetResearchResponseResultsSection]
type searchV1GetResearchResponseResultsSectionJSON struct {
	Content     apijson.Field
	Sources     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1GetResearchResponseResultsSection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1GetResearchResponseResultsSectionJSON) RawJSON() string {
	return r.raw
}

type SearchV1GetResearchResponseResultsSectionsSource struct {
	Snippet string                                               `json:"snippet"`
	Title   string                                               `json:"title"`
	URL     string                                               `json:"url"`
	JSON    searchV1GetResearchResponseResultsSectionsSourceJSON `json:"-"`
}

// searchV1GetResearchResponseResultsSectionsSourceJSON contains the JSON metadata
// for the struct [SearchV1GetResearchResponseResultsSectionsSource]
type searchV1GetResearchResponseResultsSectionsSourceJSON struct {
	Snippet     apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1GetResearchResponseResultsSectionsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1GetResearchResponseResultsSectionsSourceJSON) RawJSON() string {
	return r.raw
}

type SearchV1GetResearchResponseResultsSource struct {
	Snippet string                                       `json:"snippet"`
	Title   string                                       `json:"title"`
	URL     string                                       `json:"url"`
	JSON    searchV1GetResearchResponseResultsSourceJSON `json:"-"`
}

// searchV1GetResearchResponseResultsSourceJSON contains the JSON metadata for the
// struct [SearchV1GetResearchResponseResultsSource]
type searchV1GetResearchResponseResultsSourceJSON struct {
	Snippet     apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchV1GetResearchResponseResultsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1GetResearchResponseResultsSourceJSON) RawJSON() string {
	return r.raw
}

// Current status of the research task
type SearchV1GetResearchResponseStatus string

const (
	SearchV1GetResearchResponseStatusPending   SearchV1GetResearchResponseStatus = "pending"
	SearchV1GetResearchResponseStatusRunning   SearchV1GetResearchResponseStatus = "running"
	SearchV1GetResearchResponseStatusCompleted SearchV1GetResearchResponseStatus = "completed"
	SearchV1GetResearchResponseStatusFailed    SearchV1GetResearchResponseStatus = "failed"
)

func (r SearchV1GetResearchResponseStatus) IsKnown() bool {
	switch r {
	case SearchV1GetResearchResponseStatusPending, SearchV1GetResearchResponseStatusRunning, SearchV1GetResearchResponseStatusCompleted, SearchV1GetResearchResponseStatusFailed:
		return true
	}
	return false
}

type SearchV1SearchResponse struct {
	// Original search query
	Query string `json:"query"`
	// Array of search results
	Results []SearchV1SearchResponseResult `json:"results"`
	// Total number of results found
	TotalResults int64                      `json:"totalResults"`
	JSON         searchV1SearchResponseJSON `json:"-"`
}

// searchV1SearchResponseJSON contains the JSON metadata for the struct
// [SearchV1SearchResponse]
type searchV1SearchResponseJSON struct {
	Query        apijson.Field
	Results      apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SearchV1SearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1SearchResponseJSON) RawJSON() string {
	return r.raw
}

type SearchV1SearchResponseResult struct {
	// Domain of the source
	Domain string `json:"domain"`
	// Publication date of the content
	PublishedDate time.Time `json:"publishedDate" format:"date-time"`
	// Brief excerpt from the content
	Snippet string `json:"snippet"`
	// Title of the search result
	Title string `json:"title"`
	// URL of the search result
	URL  string                           `json:"url"`
	JSON searchV1SearchResponseResultJSON `json:"-"`
}

// searchV1SearchResponseResultJSON contains the JSON metadata for the struct
// [SearchV1SearchResponseResult]
type searchV1SearchResponseResultJSON struct {
	Domain        apijson.Field
	PublishedDate apijson.Field
	Snippet       apijson.Field
	Title         apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SearchV1SearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1SearchResponseResultJSON) RawJSON() string {
	return r.raw
}

type SearchV1SimilarResponse struct {
	ProcessingTime float64                         `json:"processingTime"`
	Results        []SearchV1SimilarResponseResult `json:"results"`
	TotalResults   int64                           `json:"totalResults"`
	JSON           searchV1SimilarResponseJSON     `json:"-"`
}

// searchV1SimilarResponseJSON contains the JSON metadata for the struct
// [SearchV1SimilarResponse]
type searchV1SimilarResponseJSON struct {
	ProcessingTime apijson.Field
	Results        apijson.Field
	TotalResults   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SearchV1SimilarResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1SimilarResponseJSON) RawJSON() string {
	return r.raw
}

type SearchV1SimilarResponseResult struct {
	Domain          string                            `json:"domain"`
	PublishedDate   string                            `json:"publishedDate"`
	SimilarityScore float64                           `json:"similarityScore"`
	Snippet         string                            `json:"snippet"`
	Text            string                            `json:"text"`
	Title           string                            `json:"title"`
	URL             string                            `json:"url"`
	JSON            searchV1SimilarResponseResultJSON `json:"-"`
}

// searchV1SimilarResponseResultJSON contains the JSON metadata for the struct
// [SearchV1SimilarResponseResult]
type searchV1SimilarResponseResultJSON struct {
	Domain          apijson.Field
	PublishedDate   apijson.Field
	SimilarityScore apijson.Field
	Snippet         apijson.Field
	Text            apijson.Field
	Title           apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SearchV1SimilarResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchV1SimilarResponseResultJSON) RawJSON() string {
	return r.raw
}

type SearchV1AnswerParams struct {
	// The question or topic to research and answer
	Query param.Field[string] `json:"query" api:"required"`
	// Exclude these domains from search
	ExcludeDomains param.Field[[]string] `json:"excludeDomains"`
	// Only search within these domains
	IncludeDomains param.Field[[]string] `json:"includeDomains"`
	// Maximum tokens for LLM response
	MaxTokens param.Field[int64] `json:"maxTokens"`
	// LLM model to use when useCustomLLM is true
	Model param.Field[string] `json:"model"`
	// Number of search results to consider
	NumResults param.Field[int64] `json:"numResults"`
	// Type of search to perform
	SearchType param.Field[SearchV1AnswerParamsSearchType] `json:"searchType"`
	// Stream the response (only for native provider answers)
	Stream param.Field[bool] `json:"stream"`
	// LLM temperature for answer generation
	Temperature param.Field[float64] `json:"temperature"`
	// Include text content in response
	Text param.Field[bool] `json:"text"`
	// Use Case.dev LLM for answer generation instead of provider's native answer
	UseCustomLlm param.Field[bool] `json:"useCustomLLM"`
}

func (r SearchV1AnswerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of search to perform
type SearchV1AnswerParamsSearchType string

const (
	SearchV1AnswerParamsSearchTypeAuto     SearchV1AnswerParamsSearchType = "auto"
	SearchV1AnswerParamsSearchTypeWeb      SearchV1AnswerParamsSearchType = "web"
	SearchV1AnswerParamsSearchTypeNews     SearchV1AnswerParamsSearchType = "news"
	SearchV1AnswerParamsSearchTypeAcademic SearchV1AnswerParamsSearchType = "academic"
)

func (r SearchV1AnswerParamsSearchType) IsKnown() bool {
	switch r {
	case SearchV1AnswerParamsSearchTypeAuto, SearchV1AnswerParamsSearchTypeWeb, SearchV1AnswerParamsSearchTypeNews, SearchV1AnswerParamsSearchTypeAcademic:
		return true
	}
	return false
}

type SearchV1ContentsParams struct {
	// Array of URLs to scrape and extract content from
	URLs param.Field[[]string] `json:"urls" api:"required" format:"uri"`
	// Context to guide content extraction and summarization
	Context param.Field[string] `json:"context"`
	// Additional extraction options
	Extras param.Field[interface{}] `json:"extras"`
	// Whether to include content highlights
	Highlights param.Field[bool] `json:"highlights"`
	// Whether to perform live crawling for dynamic content
	Livecrawl param.Field[bool] `json:"livecrawl"`
	// Timeout in seconds for live crawling
	LivecrawlTimeout param.Field[int64] `json:"livecrawlTimeout"`
	// Whether to extract content from linked subpages
	Subpages param.Field[bool] `json:"subpages"`
	// Maximum number of subpages to crawl
	SubpageTarget param.Field[int64] `json:"subpageTarget"`
	// Whether to generate content summaries
	Summary param.Field[bool] `json:"summary"`
	// Whether to extract text content
	Text param.Field[bool] `json:"text"`
}

func (r SearchV1ContentsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SearchV1ResearchParams struct {
	// Research instructions or query
	Instructions param.Field[string] `json:"instructions" api:"required"`
	// Research quality level - fast (quick), normal (balanced), pro (comprehensive)
	Model param.Field[SearchV1ResearchParamsModel] `json:"model"`
	// Optional JSON schema to structure the research output
	OutputSchema param.Field[interface{}] `json:"outputSchema"`
	// Alias for instructions (for convenience)
	Query param.Field[string] `json:"query"`
}

func (r SearchV1ResearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Research quality level - fast (quick), normal (balanced), pro (comprehensive)
type SearchV1ResearchParamsModel string

const (
	SearchV1ResearchParamsModelFast   SearchV1ResearchParamsModel = "fast"
	SearchV1ResearchParamsModelNormal SearchV1ResearchParamsModel = "normal"
	SearchV1ResearchParamsModelPro    SearchV1ResearchParamsModel = "pro"
)

func (r SearchV1ResearchParamsModel) IsKnown() bool {
	switch r {
	case SearchV1ResearchParamsModelFast, SearchV1ResearchParamsModelNormal, SearchV1ResearchParamsModelPro:
		return true
	}
	return false
}

type SearchV1GetResearchParams struct {
	// Filter specific event types for streaming
	Events param.Field[string] `query:"events"`
	// Enable streaming for real-time updates
	Stream param.Field[bool] `query:"stream"`
}

// URLQuery serializes [SearchV1GetResearchParams]'s query parameters as
// `url.Values`.
func (r SearchV1GetResearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchV1SearchParams struct {
	// Primary search query
	Query param.Field[string] `json:"query" api:"required"`
	// Additional related search queries to enhance results
	AdditionalQueries param.Field[[]string] `json:"additionalQueries"`
	// Category filter for search results
	Category param.Field[string] `json:"category"`
	// Specific content type to search for
	Contents param.Field[string] `json:"contents"`
	// End date for crawl date filtering
	EndCrawlDate param.Field[time.Time] `json:"endCrawlDate" format:"date"`
	// End date for published date filtering
	EndPublishedDate param.Field[time.Time] `json:"endPublishedDate" format:"date"`
	// Domains to exclude from search results
	ExcludeDomains param.Field[[]string] `json:"excludeDomains"`
	// Domains to include in search results
	IncludeDomains param.Field[[]string] `json:"includeDomains"`
	// Whether to include full text content in results
	IncludeText param.Field[bool] `json:"includeText"`
	// Number of search results to return
	NumResults param.Field[int64] `json:"numResults"`
	// Start date for crawl date filtering
	StartCrawlDate param.Field[time.Time] `json:"startCrawlDate" format:"date"`
	// Start date for published date filtering
	StartPublishedDate param.Field[time.Time] `json:"startPublishedDate" format:"date"`
	// Type of search to perform
	Type param.Field[SearchV1SearchParamsType] `json:"type"`
	// Geographic location for localized results
	UserLocation param.Field[string] `json:"userLocation"`
}

func (r SearchV1SearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of search to perform
type SearchV1SearchParamsType string

const (
	SearchV1SearchParamsTypeAuto   SearchV1SearchParamsType = "auto"
	SearchV1SearchParamsTypeSearch SearchV1SearchParamsType = "search"
	SearchV1SearchParamsTypeNews   SearchV1SearchParamsType = "news"
)

func (r SearchV1SearchParamsType) IsKnown() bool {
	switch r {
	case SearchV1SearchParamsTypeAuto, SearchV1SearchParamsTypeSearch, SearchV1SearchParamsTypeNews:
		return true
	}
	return false
}

type SearchV1SimilarParams struct {
	// The URL to find similar content for
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Additional content to consider for similarity matching
	Contents param.Field[string] `json:"contents"`
	// Only include pages crawled before this date
	EndCrawlDate param.Field[time.Time] `json:"endCrawlDate" format:"date"`
	// Only include pages published before this date
	EndPublishedDate param.Field[time.Time] `json:"endPublishedDate" format:"date"`
	// Exclude results from these domains
	ExcludeDomains param.Field[[]string] `json:"excludeDomains"`
	// Only search within these domains
	IncludeDomains param.Field[[]string] `json:"includeDomains"`
	// Whether to include extracted text content in results
	IncludeText param.Field[bool] `json:"includeText"`
	// Number of similar results to return
	NumResults param.Field[int64] `json:"numResults"`
	// Only include pages crawled after this date
	StartCrawlDate param.Field[time.Time] `json:"startCrawlDate" format:"date"`
	// Only include pages published after this date
	StartPublishedDate param.Field[time.Time] `json:"startPublishedDate" format:"date"`
}

func (r SearchV1SimilarParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
