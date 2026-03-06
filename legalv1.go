// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Legal research tools including citation verification
//
// LegalV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegalV1Service] method instead.
type LegalV1Service struct {
	Options []option.RequestOption
}

// NewLegalV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLegalV1Service(opts ...option.RequestOption) (r *LegalV1Service) {
	r = &LegalV1Service{}
	r.Options = opts
	return
}

// Search federal court dockets or retrieve a specific docket with optional filing
// entries via CourtListener RECAP data.
func (r *LegalV1Service) Docket(ctx context.Context, body LegalV1DocketParams, opts ...option.RequestOption) (res *LegalV1DocketResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/docket"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a legal document with structured inputs. Powered by an agent that
// handles research, formatting, citation verification, and vault upload. Returns a
// run ID for polling.
func (r *LegalV1Service) Draft(ctx context.Context, body LegalV1DraftParams, opts ...option.RequestOption) (res *LegalV1DraftResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/draft"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search for legal sources including cases, statutes, and regulations from
// authoritative legal databases. Returns ranked candidates. Always verify with
// legal.verify() before citing.
func (r *LegalV1Service) Find(ctx context.Context, body LegalV1FindParams, opts ...option.RequestOption) (res *LegalV1FindResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/find"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Parses legal citations from text and returns structured Bluebook components
// (case name, reporter, volume, page, year, court). Accepts either a single
// citation or a full text block.
func (r *LegalV1Service) GetCitations(ctx context.Context, body LegalV1GetCitationsParams, opts ...option.RequestOption) (res *LegalV1GetCitationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/citations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extract all legal citations and references from a document URL. Returns
// structured citation data including case citations, statute references, and
// regulatory citations.
func (r *LegalV1Service) GetCitationsFromURL(ctx context.Context, body LegalV1GetCitationsFromURLParams, opts ...option.RequestOption) (res *LegalV1GetCitationsFromURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/citations-from-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the full text content of a legal document. Use after verifying the
// source with legal.verify(). Returns complete text with optional highlights and
// AI summary.
func (r *LegalV1Service) GetFullText(ctx context.Context, body LegalV1GetFullTextParams, opts ...option.RequestOption) (res *LegalV1GetFullTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/full-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns CourtListener court IDs and names for docket filtering. Use these IDs in
// legal.docket() as the court parameter.
func (r *LegalV1Service) ListCourts(ctx context.Context, body LegalV1ListCourtsParams, opts ...option.RequestOption) (res *LegalV1ListCourtsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/courts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search for a jurisdiction by name. Returns matching jurisdictions with their IDs
// for use in legal.find() and other legal research endpoints.
func (r *LegalV1Service) ListJurisdictions(ctx context.Context, body LegalV1ListJurisdictionsParams, opts ...option.RequestOption) (res *LegalV1ListJurisdictionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/jurisdictions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search the USPTO Open Data Portal for US patent applications and granted
// patents. Supports free-text queries, field-specific search, filters by
// assignee/inventor/status/type, date ranges, and pagination. Covers applications
// filed on or after January 1, 2001. Data is refreshed daily.
func (r *LegalV1Service) PatentSearch(ctx context.Context, body LegalV1PatentSearchParams, opts ...option.RequestOption) (res *LegalV1PatentSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/patent-search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform comprehensive legal research with multiple query variations. Uses
// advanced deep search to find relevant sources across different phrasings of the
// legal issue.
func (r *LegalV1Service) Research(ctx context.Context, body LegalV1ResearchParams, opts ...option.RequestOption) (res *LegalV1ResearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find cases and documents similar to a given legal source. Useful for finding
// citing cases, related precedents, or similar statutes.
func (r *LegalV1Service) Similar(ctx context.Context, body LegalV1SimilarParams, opts ...option.RequestOption) (res *LegalV1SimilarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/similar"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Look up trademark status and details from the USPTO Trademark Status & Document
// Retrieval (TSDR) system. Supports lookup by serial number or registration
// number. Returns mark text, status, owner, goods/services, Nice classification,
// filing/registration dates, and more.
func (r *LegalV1Service) TrademarkSearch(ctx context.Context, body LegalV1TrademarkSearchParams, opts ...option.RequestOption) (res *LegalV1TrademarkSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/trademark-search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validates legal citations against authoritative case law sources (CourtListener
// database of ~10M cases). Returns verification status and case metadata for each
// citation found in the input text. Accepts either a single citation or a full
// text block containing multiple citations.
func (r *LegalV1Service) Verify(ctx context.Context, body LegalV1VerifyParams, opts ...option.RequestOption) (res *LegalV1VerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legal/v1/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Full docket record (lookup mode)
type DocketDetail struct {
	ID             string           `json:"id"`
	AssignedTo     string           `json:"assignedTo" api:"nullable"`
	CaseName       string           `json:"caseName" api:"nullable"`
	Cause          string           `json:"cause" api:"nullable"`
	Court          string           `json:"court" api:"nullable"`
	CourtID        string           `json:"courtId" api:"nullable"`
	DateFiled      time.Time        `json:"dateFiled" api:"nullable" format:"date"`
	DateTerminated time.Time        `json:"dateTerminated" api:"nullable" format:"date"`
	DocketNumber   string           `json:"docketNumber" api:"nullable"`
	NatureOfSuit   string           `json:"natureOfSuit" api:"nullable"`
	PacerCaseID    string           `json:"pacerCaseId" api:"nullable"`
	Parties        []string         `json:"parties"`
	URL            string           `json:"url"`
	JSON           docketDetailJSON `json:"-"`
}

// docketDetailJSON contains the JSON metadata for the struct [DocketDetail]
type docketDetailJSON struct {
	ID             apijson.Field
	AssignedTo     apijson.Field
	CaseName       apijson.Field
	Cause          apijson.Field
	Court          apijson.Field
	CourtID        apijson.Field
	DateFiled      apijson.Field
	DateTerminated apijson.Field
	DocketNumber   apijson.Field
	NatureOfSuit   apijson.Field
	PacerCaseID    apijson.Field
	Parties        apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DocketDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r docketDetailJSON) RawJSON() string {
	return r.raw
}

type DocketSearchResult struct {
	ID             string                 `json:"id"`
	AssignedTo     string                 `json:"assignedTo" api:"nullable"`
	CaseName       string                 `json:"caseName" api:"nullable"`
	Cause          string                 `json:"cause" api:"nullable"`
	Court          string                 `json:"court" api:"nullable"`
	CourtID        string                 `json:"courtId" api:"nullable"`
	DateFiled      time.Time              `json:"dateFiled" api:"nullable" format:"date"`
	DateTerminated time.Time              `json:"dateTerminated" api:"nullable" format:"date"`
	DocketNumber   string                 `json:"docketNumber" api:"nullable"`
	NatureOfSuit   string                 `json:"natureOfSuit" api:"nullable"`
	PacerCaseID    string                 `json:"pacerCaseId" api:"nullable"`
	Parties        []string               `json:"parties"`
	URL            string                 `json:"url"`
	JSON           docketSearchResultJSON `json:"-"`
}

// docketSearchResultJSON contains the JSON metadata for the struct
// [DocketSearchResult]
type docketSearchResultJSON struct {
	ID             apijson.Field
	AssignedTo     apijson.Field
	CaseName       apijson.Field
	Cause          apijson.Field
	Court          apijson.Field
	CourtID        apijson.Field
	DateFiled      apijson.Field
	DateTerminated apijson.Field
	DocketNumber   apijson.Field
	NatureOfSuit   apijson.Field
	PacerCaseID    apijson.Field
	Parties        apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DocketSearchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r docketSearchResultJSON) RawJSON() string {
	return r.raw
}

type LegalV1DocketResponse struct {
	// Echo of court filter (search mode only)
	Court string `json:"court" api:"nullable"`
	// Echo of date filter
	DateFiledAfter time.Time `json:"dateFiledAfter" api:"nullable" format:"date"`
	// Echo of date filter
	DateFiledBefore time.Time `json:"dateFiledBefore" api:"nullable" format:"date"`
	// Full docket record (lookup mode)
	Docket DocketDetail `json:"docket" api:"nullable"`
	// Search results (search mode)
	Dockets []DocketSearchResult `json:"dockets"`
	// Docket entries/filings (lookup mode with includeEntries)
	Entries []LegalV1DocketResponseEntry `json:"entries" api:"nullable"`
	Found   int64                        `json:"found"`
	// Whether entries were requested (lookup mode only)
	IncludeEntries bool `json:"includeEntries"`
	// Pagination info for entry list (lookup mode with includeEntries)
	Pagination LegalV1DocketResponsePagination `json:"pagination" api:"nullable"`
	// Echo of search query (search mode only)
	Query string                    `json:"query" api:"nullable"`
	Type  LegalV1DocketResponseType `json:"type"`
	JSON  legalV1DocketResponseJSON `json:"-"`
}

// legalV1DocketResponseJSON contains the JSON metadata for the struct
// [LegalV1DocketResponse]
type legalV1DocketResponseJSON struct {
	Court           apijson.Field
	DateFiledAfter  apijson.Field
	DateFiledBefore apijson.Field
	Docket          apijson.Field
	Dockets         apijson.Field
	Entries         apijson.Field
	Found           apijson.Field
	IncludeEntries  apijson.Field
	Pagination      apijson.Field
	Query           apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LegalV1DocketResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1DocketResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1DocketResponseEntry struct {
	Date        time.Time                              `json:"date" api:"nullable" format:"date"`
	Description string                                 `json:"description" api:"nullable"`
	Documents   []LegalV1DocketResponseEntriesDocument `json:"documents"`
	EntryNumber int64                                  `json:"entryNumber" api:"nullable"`
	JSON        legalV1DocketResponseEntryJSON         `json:"-"`
}

// legalV1DocketResponseEntryJSON contains the JSON metadata for the struct
// [LegalV1DocketResponseEntry]
type legalV1DocketResponseEntryJSON struct {
	Date        apijson.Field
	Description apijson.Field
	Documents   apijson.Field
	EntryNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1DocketResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1DocketResponseEntryJSON) RawJSON() string {
	return r.raw
}

type LegalV1DocketResponseEntriesDocument struct {
	ID               string                                   `json:"id"`
	AttachmentNumber int64                                    `json:"attachmentNumber" api:"nullable"`
	Description      string                                   `json:"description" api:"nullable"`
	DocumentNumber   string                                   `json:"documentNumber" api:"nullable"`
	IsAvailable      bool                                     `json:"isAvailable"`
	PageCount        int64                                    `json:"pageCount" api:"nullable"`
	PdfURL           string                                   `json:"pdfUrl" api:"nullable"`
	JSON             legalV1DocketResponseEntriesDocumentJSON `json:"-"`
}

// legalV1DocketResponseEntriesDocumentJSON contains the JSON metadata for the
// struct [LegalV1DocketResponseEntriesDocument]
type legalV1DocketResponseEntriesDocumentJSON struct {
	ID               apijson.Field
	AttachmentNumber apijson.Field
	Description      apijson.Field
	DocumentNumber   apijson.Field
	IsAvailable      apijson.Field
	PageCount        apijson.Field
	PdfURL           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LegalV1DocketResponseEntriesDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1DocketResponseEntriesDocumentJSON) RawJSON() string {
	return r.raw
}

// Pagination info for entry list (lookup mode with includeEntries)
type LegalV1DocketResponsePagination struct {
	Limit    int64                               `json:"limit"`
	Offset   int64                               `json:"offset"`
	Returned int64                               `json:"returned"`
	JSON     legalV1DocketResponsePaginationJSON `json:"-"`
}

// legalV1DocketResponsePaginationJSON contains the JSON metadata for the struct
// [LegalV1DocketResponsePagination]
type legalV1DocketResponsePaginationJSON struct {
	Limit       apijson.Field
	Offset      apijson.Field
	Returned    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1DocketResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1DocketResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type LegalV1DocketResponseType string

const (
	LegalV1DocketResponseTypeSearch LegalV1DocketResponseType = "search"
	LegalV1DocketResponseTypeLookup LegalV1DocketResponseType = "lookup"
)

func (r LegalV1DocketResponseType) IsKnown() bool {
	switch r {
	case LegalV1DocketResponseTypeSearch, LegalV1DocketResponseTypeLookup:
		return true
	}
	return false
}

type LegalV1DraftResponse struct {
	// Ephemeral agent ID
	AgentID string `json:"agent_id"`
	Message string `json:"message"`
	// Run ID — poll /agent/v1/run/:id/status for progress
	RunID  string                     `json:"run_id"`
	Status LegalV1DraftResponseStatus `json:"status"`
	JSON   legalV1DraftResponseJSON   `json:"-"`
}

// legalV1DraftResponseJSON contains the JSON metadata for the struct
// [LegalV1DraftResponse]
type legalV1DraftResponseJSON struct {
	AgentID     apijson.Field
	Message     apijson.Field
	RunID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1DraftResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1DraftResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1DraftResponseStatus string

const (
	LegalV1DraftResponseStatusRunning LegalV1DraftResponseStatus = "running"
)

func (r LegalV1DraftResponseStatus) IsKnown() bool {
	switch r {
	case LegalV1DraftResponseStatusRunning:
		return true
	}
	return false
}

type LegalV1FindResponse struct {
	Candidates []LegalV1FindResponseCandidate `json:"candidates"`
	// Number of candidates found
	Found int64 `json:"found"`
	// Usage guidance
	Hint string `json:"hint"`
	// Jurisdiction filter applied
	Jurisdiction string `json:"jurisdiction"`
	// Original search query
	Query string                  `json:"query"`
	JSON  legalV1FindResponseJSON `json:"-"`
}

// legalV1FindResponseJSON contains the JSON metadata for the struct
// [LegalV1FindResponse]
type legalV1FindResponseJSON struct {
	Candidates   apijson.Field
	Found        apijson.Field
	Hint         apijson.Field
	Jurisdiction apijson.Field
	Query        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LegalV1FindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1FindResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1FindResponseCandidate struct {
	// Text excerpt from the document
	Snippet string `json:"snippet"`
	// Domain of the source
	Source string `json:"source"`
	// Title of the document
	Title string `json:"title"`
	// URL of the legal source
	URL  string                           `json:"url"`
	JSON legalV1FindResponseCandidateJSON `json:"-"`
}

// legalV1FindResponseCandidateJSON contains the JSON metadata for the struct
// [LegalV1FindResponseCandidate]
type legalV1FindResponseCandidateJSON struct {
	Snippet     apijson.Field
	Source      apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1FindResponseCandidate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1FindResponseCandidateJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsResponse struct {
	Citations []LegalV1GetCitationsResponseCitation `json:"citations"`
	JSON      legalV1GetCitationsResponseJSON       `json:"-"`
}

// legalV1GetCitationsResponseJSON contains the JSON metadata for the struct
// [LegalV1GetCitationsResponse]
type legalV1GetCitationsResponseJSON struct {
	Citations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsResponseCitation struct {
	// Structured Bluebook components. Null if citation format is not recognized.
	Components LegalV1GetCitationsResponseCitationsComponents `json:"components" api:"nullable"`
	// Whether citation was found in CourtListener database
	Found bool `json:"found"`
	// Normalized citation string
	Normalized string `json:"normalized"`
	// Original citation as found in text
	Original string                                   `json:"original"`
	Span     LegalV1GetCitationsResponseCitationsSpan `json:"span"`
	JSON     legalV1GetCitationsResponseCitationJSON  `json:"-"`
}

// legalV1GetCitationsResponseCitationJSON contains the JSON metadata for the
// struct [LegalV1GetCitationsResponseCitation]
type legalV1GetCitationsResponseCitationJSON struct {
	Components  apijson.Field
	Found       apijson.Field
	Normalized  apijson.Field
	Original    apijson.Field
	Span        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsResponseCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsResponseCitationJSON) RawJSON() string {
	return r.raw
}

// Structured Bluebook components. Null if citation format is not recognized.
type LegalV1GetCitationsResponseCitationsComponents struct {
	// Case name, e.g., "Bush v. Gore"
	CaseName string `json:"caseName"`
	// Court identifier
	Court string `json:"court"`
	// Starting page number
	Page int64 `json:"page"`
	// Pin cite (specific page)
	PinCite int64 `json:"pinCite"`
	// Reporter abbreviation, e.g., "U.S."
	Reporter string `json:"reporter"`
	// Volume number
	Volume int64 `json:"volume"`
	// Decision year
	Year int64                                              `json:"year"`
	JSON legalV1GetCitationsResponseCitationsComponentsJSON `json:"-"`
}

// legalV1GetCitationsResponseCitationsComponentsJSON contains the JSON metadata
// for the struct [LegalV1GetCitationsResponseCitationsComponents]
type legalV1GetCitationsResponseCitationsComponentsJSON struct {
	CaseName    apijson.Field
	Court       apijson.Field
	Page        apijson.Field
	PinCite     apijson.Field
	Reporter    apijson.Field
	Volume      apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsResponseCitationsComponents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsResponseCitationsComponentsJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsResponseCitationsSpan struct {
	End   int64                                        `json:"end"`
	Start int64                                        `json:"start"`
	JSON  legalV1GetCitationsResponseCitationsSpanJSON `json:"-"`
}

// legalV1GetCitationsResponseCitationsSpanJSON contains the JSON metadata for the
// struct [LegalV1GetCitationsResponseCitationsSpan]
type legalV1GetCitationsResponseCitationsSpanJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsResponseCitationsSpan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsResponseCitationsSpanJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsFromURLResponse struct {
	Citations LegalV1GetCitationsFromURLResponseCitations `json:"citations"`
	// External links found in the document
	ExternalLinks []string `json:"externalLinks"`
	// Usage guidance
	Hint string `json:"hint"`
	// Document title
	Title string `json:"title"`
	// Total citations found
	TotalCitations int64 `json:"totalCitations"`
	// Source document URL
	URL  string                                 `json:"url"`
	JSON legalV1GetCitationsFromURLResponseJSON `json:"-"`
}

// legalV1GetCitationsFromURLResponseJSON contains the JSON metadata for the struct
// [LegalV1GetCitationsFromURLResponse]
type legalV1GetCitationsFromURLResponseJSON struct {
	Citations      apijson.Field
	ExternalLinks  apijson.Field
	Hint           apijson.Field
	Title          apijson.Field
	TotalCitations apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LegalV1GetCitationsFromURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsFromURLResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsFromURLResponseCitations struct {
	Cases       []LegalV1GetCitationsFromURLResponseCitationsCase       `json:"cases"`
	Regulations []LegalV1GetCitationsFromURLResponseCitationsRegulation `json:"regulations"`
	Statutes    []LegalV1GetCitationsFromURLResponseCitationsStatute    `json:"statutes"`
	JSON        legalV1GetCitationsFromURLResponseCitationsJSON         `json:"-"`
}

// legalV1GetCitationsFromURLResponseCitationsJSON contains the JSON metadata for
// the struct [LegalV1GetCitationsFromURLResponseCitations]
type legalV1GetCitationsFromURLResponseCitationsJSON struct {
	Cases       apijson.Field
	Regulations apijson.Field
	Statutes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsFromURLResponseCitations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsFromURLResponseCitationsJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsFromURLResponseCitationsCase struct {
	// The citation string
	Citation string `json:"citation"`
	// Number of occurrences
	Count int64 `json:"count"`
	// Citation type (usReporter, federalReporter, etc.)
	Type string                                              `json:"type"`
	JSON legalV1GetCitationsFromURLResponseCitationsCaseJSON `json:"-"`
}

// legalV1GetCitationsFromURLResponseCitationsCaseJSON contains the JSON metadata
// for the struct [LegalV1GetCitationsFromURLResponseCitationsCase]
type legalV1GetCitationsFromURLResponseCitationsCaseJSON struct {
	Citation    apijson.Field
	Count       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsFromURLResponseCitationsCase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsFromURLResponseCitationsCaseJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsFromURLResponseCitationsRegulation struct {
	// The citation string
	Citation string `json:"citation"`
	// Number of occurrences
	Count int64 `json:"count"`
	// Citation type (cfr)
	Type string                                                    `json:"type"`
	JSON legalV1GetCitationsFromURLResponseCitationsRegulationJSON `json:"-"`
}

// legalV1GetCitationsFromURLResponseCitationsRegulationJSON contains the JSON
// metadata for the struct [LegalV1GetCitationsFromURLResponseCitationsRegulation]
type legalV1GetCitationsFromURLResponseCitationsRegulationJSON struct {
	Citation    apijson.Field
	Count       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsFromURLResponseCitationsRegulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsFromURLResponseCitationsRegulationJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetCitationsFromURLResponseCitationsStatute struct {
	// The citation string
	Citation string `json:"citation"`
	// Number of occurrences
	Count int64 `json:"count"`
	// Citation type (usc)
	Type string                                                 `json:"type"`
	JSON legalV1GetCitationsFromURLResponseCitationsStatuteJSON `json:"-"`
}

// legalV1GetCitationsFromURLResponseCitationsStatuteJSON contains the JSON
// metadata for the struct [LegalV1GetCitationsFromURLResponseCitationsStatute]
type legalV1GetCitationsFromURLResponseCitationsStatuteJSON struct {
	Citation    apijson.Field
	Count       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1GetCitationsFromURLResponseCitationsStatute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetCitationsFromURLResponseCitationsStatuteJSON) RawJSON() string {
	return r.raw
}

type LegalV1GetFullTextResponse struct {
	// Author or court
	Author string `json:"author" api:"nullable"`
	// Total characters in text
	CharacterCount int64 `json:"characterCount"`
	// Highlighted relevant passages
	Highlights []string `json:"highlights"`
	// Publication date
	PublishedDate string `json:"publishedDate" api:"nullable"`
	// AI-generated summary
	Summary string `json:"summary" api:"nullable"`
	// Full document text
	Text string `json:"text"`
	// Document title
	Title string `json:"title"`
	// Document URL
	URL  string                         `json:"url"`
	JSON legalV1GetFullTextResponseJSON `json:"-"`
}

// legalV1GetFullTextResponseJSON contains the JSON metadata for the struct
// [LegalV1GetFullTextResponse]
type legalV1GetFullTextResponseJSON struct {
	Author         apijson.Field
	CharacterCount apijson.Field
	Highlights     apijson.Field
	PublishedDate  apijson.Field
	Summary        apijson.Field
	Text           apijson.Field
	Title          apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LegalV1GetFullTextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1GetFullTextResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1ListCourtsResponse struct {
	Courts []LegalV1ListCourtsResponseCourt `json:"courts"`
	Found  int64                            `json:"found"`
	// Whether results are filtered to in-use courts only
	InUseOnly    bool                          `json:"inUseOnly"`
	Jurisdiction string                        `json:"jurisdiction" api:"nullable"`
	Query        string                        `json:"query" api:"nullable"`
	JSON         legalV1ListCourtsResponseJSON `json:"-"`
}

// legalV1ListCourtsResponseJSON contains the JSON metadata for the struct
// [LegalV1ListCourtsResponse]
type legalV1ListCourtsResponseJSON struct {
	Courts       apijson.Field
	Found        apijson.Field
	InUseOnly    apijson.Field
	Jurisdiction apijson.Field
	Query        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LegalV1ListCourtsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1ListCourtsResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1ListCourtsResponseCourt struct {
	// CourtListener court slug
	ID           string                             `json:"id"`
	FullName     string                             `json:"fullName" api:"nullable"`
	Jurisdiction string                             `json:"jurisdiction" api:"nullable"`
	PacerCourtID int64                              `json:"pacerCourtId" api:"nullable"`
	ShortName    string                             `json:"shortName" api:"nullable"`
	JSON         legalV1ListCourtsResponseCourtJSON `json:"-"`
}

// legalV1ListCourtsResponseCourtJSON contains the JSON metadata for the struct
// [LegalV1ListCourtsResponseCourt]
type legalV1ListCourtsResponseCourtJSON struct {
	ID           apijson.Field
	FullName     apijson.Field
	Jurisdiction apijson.Field
	PacerCourtID apijson.Field
	ShortName    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LegalV1ListCourtsResponseCourt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1ListCourtsResponseCourtJSON) RawJSON() string {
	return r.raw
}

type LegalV1ListJurisdictionsResponse struct {
	// Number of matching jurisdictions
	Found int64 `json:"found"`
	// Usage guidance
	Hint          string                                         `json:"hint"`
	Jurisdictions []LegalV1ListJurisdictionsResponseJurisdiction `json:"jurisdictions"`
	// Original search query
	Query string                               `json:"query"`
	JSON  legalV1ListJurisdictionsResponseJSON `json:"-"`
}

// legalV1ListJurisdictionsResponseJSON contains the JSON metadata for the struct
// [LegalV1ListJurisdictionsResponse]
type legalV1ListJurisdictionsResponseJSON struct {
	Found         apijson.Field
	Hint          apijson.Field
	Jurisdictions apijson.Field
	Query         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LegalV1ListJurisdictionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1ListJurisdictionsResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1ListJurisdictionsResponseJurisdiction struct {
	// Jurisdiction ID to use in other endpoints
	ID string `json:"id"`
	// Jurisdiction level
	Level LegalV1ListJurisdictionsResponseJurisdictionsLevel `json:"level"`
	// Full jurisdiction name
	Name string `json:"name"`
	// State abbreviation (if applicable)
	State string                                           `json:"state" api:"nullable"`
	JSON  legalV1ListJurisdictionsResponseJurisdictionJSON `json:"-"`
}

// legalV1ListJurisdictionsResponseJurisdictionJSON contains the JSON metadata for
// the struct [LegalV1ListJurisdictionsResponseJurisdiction]
type legalV1ListJurisdictionsResponseJurisdictionJSON struct {
	ID          apijson.Field
	Level       apijson.Field
	Name        apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1ListJurisdictionsResponseJurisdiction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1ListJurisdictionsResponseJurisdictionJSON) RawJSON() string {
	return r.raw
}

// Jurisdiction level
type LegalV1ListJurisdictionsResponseJurisdictionsLevel string

const (
	LegalV1ListJurisdictionsResponseJurisdictionsLevelFederal   LegalV1ListJurisdictionsResponseJurisdictionsLevel = "federal"
	LegalV1ListJurisdictionsResponseJurisdictionsLevelState     LegalV1ListJurisdictionsResponseJurisdictionsLevel = "state"
	LegalV1ListJurisdictionsResponseJurisdictionsLevelCounty    LegalV1ListJurisdictionsResponseJurisdictionsLevel = "county"
	LegalV1ListJurisdictionsResponseJurisdictionsLevelMunicipal LegalV1ListJurisdictionsResponseJurisdictionsLevel = "municipal"
)

func (r LegalV1ListJurisdictionsResponseJurisdictionsLevel) IsKnown() bool {
	switch r {
	case LegalV1ListJurisdictionsResponseJurisdictionsLevelFederal, LegalV1ListJurisdictionsResponseJurisdictionsLevelState, LegalV1ListJurisdictionsResponseJurisdictionsLevelCounty, LegalV1ListJurisdictionsResponseJurisdictionsLevelMunicipal:
		return true
	}
	return false
}

type LegalV1PatentSearchResponse struct {
	// Number of results returned
	Limit int64 `json:"limit"`
	// Current pagination offset
	Offset int64 `json:"offset"`
	// Original search query
	Query string `json:"query"`
	// Array of matching patent applications
	Results []LegalV1PatentSearchResponseResult `json:"results"`
	// Total number of matching patent applications
	TotalResults int64                           `json:"totalResults"`
	JSON         legalV1PatentSearchResponseJSON `json:"-"`
}

// legalV1PatentSearchResponseJSON contains the JSON metadata for the struct
// [LegalV1PatentSearchResponse]
type legalV1PatentSearchResponseJSON struct {
	Limit        apijson.Field
	Offset       apijson.Field
	Query        apijson.Field
	Results      apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LegalV1PatentSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1PatentSearchResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1PatentSearchResponseResult struct {
	// Patent application serial number
	ApplicationNumber string `json:"applicationNumber"`
	// Application type (Utility, Design, Plant, etc.)
	ApplicationType string `json:"applicationType"`
	// List of assignee/owner names
	Assignees []string `json:"assignees"`
	// Entity status (e.g. "Small Entity", "Micro Entity")
	EntityStatus string `json:"entityStatus" api:"nullable"`
	// Date the application was filed
	FilingDate time.Time `json:"filingDate" api:"nullable" format:"date"`
	// Date the patent was granted
	GrantDate time.Time `json:"grantDate" api:"nullable" format:"date"`
	// List of inventor names
	Inventors []string `json:"inventors"`
	// Granted patent number (if granted)
	PatentNumber string `json:"patentNumber" api:"nullable"`
	// Current application status (e.g. "Patented Case", "Pending")
	Status string `json:"status"`
	// Invention title
	Title string                                `json:"title"`
	JSON  legalV1PatentSearchResponseResultJSON `json:"-"`
}

// legalV1PatentSearchResponseResultJSON contains the JSON metadata for the struct
// [LegalV1PatentSearchResponseResult]
type legalV1PatentSearchResponseResultJSON struct {
	ApplicationNumber apijson.Field
	ApplicationType   apijson.Field
	Assignees         apijson.Field
	EntityStatus      apijson.Field
	FilingDate        apijson.Field
	GrantDate         apijson.Field
	Inventors         apijson.Field
	PatentNumber      apijson.Field
	Status            apijson.Field
	Title             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LegalV1PatentSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1PatentSearchResponseResultJSON) RawJSON() string {
	return r.raw
}

type LegalV1ResearchResponse struct {
	// Additional queries used
	AdditionalQueries []string                           `json:"additionalQueries" api:"nullable"`
	Candidates        []LegalV1ResearchResponseCandidate `json:"candidates"`
	// Number of candidates found
	Found int64 `json:"found"`
	// Usage guidance
	Hint string `json:"hint"`
	// Jurisdiction filter applied
	Jurisdiction string `json:"jurisdiction"`
	// Primary search query
	Query string `json:"query"`
	// Search type used (deep)
	SearchType string                      `json:"searchType"`
	JSON       legalV1ResearchResponseJSON `json:"-"`
}

// legalV1ResearchResponseJSON contains the JSON metadata for the struct
// [LegalV1ResearchResponse]
type legalV1ResearchResponseJSON struct {
	AdditionalQueries apijson.Field
	Candidates        apijson.Field
	Found             apijson.Field
	Hint              apijson.Field
	Jurisdiction      apijson.Field
	Query             apijson.Field
	SearchType        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LegalV1ResearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1ResearchResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1ResearchResponseCandidate struct {
	// Highlighted relevant passages
	Highlights []string `json:"highlights"`
	// Publication date
	PublishedDate string `json:"publishedDate" api:"nullable"`
	// Text excerpt from the document
	Snippet string `json:"snippet"`
	// Domain of the source
	Source string `json:"source"`
	// Title of the document
	Title string `json:"title"`
	// URL of the legal source
	URL  string                               `json:"url"`
	JSON legalV1ResearchResponseCandidateJSON `json:"-"`
}

// legalV1ResearchResponseCandidateJSON contains the JSON metadata for the struct
// [LegalV1ResearchResponseCandidate]
type legalV1ResearchResponseCandidateJSON struct {
	Highlights    apijson.Field
	PublishedDate apijson.Field
	Snippet       apijson.Field
	Source        apijson.Field
	Title         apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LegalV1ResearchResponseCandidate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1ResearchResponseCandidateJSON) RawJSON() string {
	return r.raw
}

type LegalV1SimilarResponse struct {
	// Number of similar sources found
	Found int64 `json:"found"`
	// Usage guidance
	Hint string `json:"hint"`
	// Jurisdiction filter applied
	Jurisdiction   string                                `json:"jurisdiction"`
	SimilarSources []LegalV1SimilarResponseSimilarSource `json:"similarSources"`
	// Original source URL
	SourceURL string                     `json:"sourceUrl"`
	JSON      legalV1SimilarResponseJSON `json:"-"`
}

// legalV1SimilarResponseJSON contains the JSON metadata for the struct
// [LegalV1SimilarResponse]
type legalV1SimilarResponseJSON struct {
	Found          apijson.Field
	Hint           apijson.Field
	Jurisdiction   apijson.Field
	SimilarSources apijson.Field
	SourceURL      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LegalV1SimilarResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1SimilarResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1SimilarResponseSimilarSource struct {
	// Publication date
	PublishedDate string `json:"publishedDate" api:"nullable"`
	// Text excerpt from the document
	Snippet string `json:"snippet"`
	// Domain of the source
	Source string `json:"source"`
	// Title of the document
	Title string `json:"title"`
	// URL of the similar source
	URL  string                                  `json:"url"`
	JSON legalV1SimilarResponseSimilarSourceJSON `json:"-"`
}

// legalV1SimilarResponseSimilarSourceJSON contains the JSON metadata for the
// struct [LegalV1SimilarResponseSimilarSource]
type legalV1SimilarResponseSimilarSourceJSON struct {
	PublishedDate apijson.Field
	Snippet       apijson.Field
	Source        apijson.Field
	Title         apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LegalV1SimilarResponseSimilarSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1SimilarResponseSimilarSourceJSON) RawJSON() string {
	return r.raw
}

type LegalV1TrademarkSearchResponse struct {
	// Attorney of record
	Attorney string `json:"attorney" api:"nullable"`
	// Date the application was filed
	FilingDate time.Time `json:"filingDate" api:"nullable" format:"date"`
	// Goods and services descriptions with class numbers
	GoodsAndServices []LegalV1TrademarkSearchResponseGoodsAndService `json:"goodsAndServices"`
	// URL to the mark image on USPTO CDN
	ImageURL string `json:"imageUrl" api:"nullable"`
	// The text of the trademark
	MarkText string `json:"markText" api:"nullable"`
	// Type of mark (e.g. "Standard Character Mark", "Design Mark")
	MarkType string `json:"markType" api:"nullable"`
	// Nice classification class numbers
	NiceClasses []int64 `json:"niceClasses"`
	// Current owner/applicant information
	Owner LegalV1TrademarkSearchResponseOwner `json:"owner" api:"nullable"`
	// Date the mark was registered
	RegistrationDate time.Time `json:"registrationDate" api:"nullable" format:"date"`
	// USPTO registration number (if registered)
	RegistrationNumber string `json:"registrationNumber" api:"nullable"`
	// USPTO serial number
	SerialNumber string `json:"serialNumber"`
	// Current status (e.g. "Registered", "Pending", "Abandoned", "Cancelled")
	Status string `json:"status" api:"nullable"`
	// Date of most recent status update
	StatusDate time.Time `json:"statusDate" api:"nullable" format:"date"`
	// Canonical TSDR link for this mark
	UsptoURL string                             `json:"usptoUrl"`
	JSON     legalV1TrademarkSearchResponseJSON `json:"-"`
}

// legalV1TrademarkSearchResponseJSON contains the JSON metadata for the struct
// [LegalV1TrademarkSearchResponse]
type legalV1TrademarkSearchResponseJSON struct {
	Attorney           apijson.Field
	FilingDate         apijson.Field
	GoodsAndServices   apijson.Field
	ImageURL           apijson.Field
	MarkText           apijson.Field
	MarkType           apijson.Field
	NiceClasses        apijson.Field
	Owner              apijson.Field
	RegistrationDate   apijson.Field
	RegistrationNumber apijson.Field
	SerialNumber       apijson.Field
	Status             apijson.Field
	StatusDate         apijson.Field
	UsptoURL           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LegalV1TrademarkSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1TrademarkSearchResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1TrademarkSearchResponseGoodsAndService struct {
	ClassNumber string                                            `json:"classNumber" api:"nullable"`
	Description string                                            `json:"description" api:"nullable"`
	JSON        legalV1TrademarkSearchResponseGoodsAndServiceJSON `json:"-"`
}

// legalV1TrademarkSearchResponseGoodsAndServiceJSON contains the JSON metadata for
// the struct [LegalV1TrademarkSearchResponseGoodsAndService]
type legalV1TrademarkSearchResponseGoodsAndServiceJSON struct {
	ClassNumber apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1TrademarkSearchResponseGoodsAndService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1TrademarkSearchResponseGoodsAndServiceJSON) RawJSON() string {
	return r.raw
}

// Current owner/applicant information
type LegalV1TrademarkSearchResponseOwner struct {
	Address    string                                  `json:"address" api:"nullable"`
	EntityType string                                  `json:"entityType" api:"nullable"`
	Name       string                                  `json:"name" api:"nullable"`
	JSON       legalV1TrademarkSearchResponseOwnerJSON `json:"-"`
}

// legalV1TrademarkSearchResponseOwnerJSON contains the JSON metadata for the
// struct [LegalV1TrademarkSearchResponseOwner]
type legalV1TrademarkSearchResponseOwnerJSON struct {
	Address     apijson.Field
	EntityType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1TrademarkSearchResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1TrademarkSearchResponseOwnerJSON) RawJSON() string {
	return r.raw
}

type LegalV1VerifyResponse struct {
	Citations []LegalV1VerifyResponseCitation `json:"citations"`
	Summary   LegalV1VerifyResponseSummary    `json:"summary"`
	JSON      legalV1VerifyResponseJSON       `json:"-"`
}

// legalV1VerifyResponseJSON contains the JSON metadata for the struct
// [LegalV1VerifyResponse]
type legalV1VerifyResponseJSON struct {
	Citations   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1VerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1VerifyResponseJSON) RawJSON() string {
	return r.raw
}

type LegalV1VerifyResponseCitation struct {
	// Multiple candidates (when multiple_matches or heuristic verification)
	Candidates []LegalV1VerifyResponseCitationsCandidate `json:"candidates"`
	// Case metadata (when verified)
	Case LegalV1VerifyResponseCitationsCase `json:"case"`
	// Confidence score (1.0 for CourtListener, heuristic score for fallback).
	Confidence float64 `json:"confidence"`
	// Normalized citation string
	Normalized string `json:"normalized"`
	// Original citation as found in text
	Original string                               `json:"original"`
	Span     LegalV1VerifyResponseCitationsSpan   `json:"span"`
	Status   LegalV1VerifyResponseCitationsStatus `json:"status"`
	// Source of verification result (heuristic for fallback matches).
	VerificationSource LegalV1VerifyResponseCitationsVerificationSource `json:"verificationSource"`
	JSON               legalV1VerifyResponseCitationJSON                `json:"-"`
}

// legalV1VerifyResponseCitationJSON contains the JSON metadata for the struct
// [LegalV1VerifyResponseCitation]
type legalV1VerifyResponseCitationJSON struct {
	Candidates         apijson.Field
	Case               apijson.Field
	Confidence         apijson.Field
	Normalized         apijson.Field
	Original           apijson.Field
	Span               apijson.Field
	Status             apijson.Field
	VerificationSource apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LegalV1VerifyResponseCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1VerifyResponseCitationJSON) RawJSON() string {
	return r.raw
}

type LegalV1VerifyResponseCitationsCandidate struct {
	Court       string                                      `json:"court"`
	DateDecided string                                      `json:"dateDecided"`
	Name        string                                      `json:"name"`
	URL         string                                      `json:"url"`
	JSON        legalV1VerifyResponseCitationsCandidateJSON `json:"-"`
}

// legalV1VerifyResponseCitationsCandidateJSON contains the JSON metadata for the
// struct [LegalV1VerifyResponseCitationsCandidate]
type legalV1VerifyResponseCitationsCandidateJSON struct {
	Court       apijson.Field
	DateDecided apijson.Field
	Name        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1VerifyResponseCitationsCandidate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1VerifyResponseCitationsCandidateJSON) RawJSON() string {
	return r.raw
}

// Case metadata (when verified)
type LegalV1VerifyResponseCitationsCase struct {
	ID                int64                                  `json:"id"`
	Court             string                                 `json:"court"`
	DateDecided       string                                 `json:"dateDecided"`
	DocketNumber      string                                 `json:"docketNumber"`
	Name              string                                 `json:"name"`
	ParallelCitations []string                               `json:"parallelCitations"`
	ShortName         string                                 `json:"shortName"`
	URL               string                                 `json:"url"`
	JSON              legalV1VerifyResponseCitationsCaseJSON `json:"-"`
}

// legalV1VerifyResponseCitationsCaseJSON contains the JSON metadata for the struct
// [LegalV1VerifyResponseCitationsCase]
type legalV1VerifyResponseCitationsCaseJSON struct {
	ID                apijson.Field
	Court             apijson.Field
	DateDecided       apijson.Field
	DocketNumber      apijson.Field
	Name              apijson.Field
	ParallelCitations apijson.Field
	ShortName         apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LegalV1VerifyResponseCitationsCase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1VerifyResponseCitationsCaseJSON) RawJSON() string {
	return r.raw
}

type LegalV1VerifyResponseCitationsSpan struct {
	End   int64                                  `json:"end"`
	Start int64                                  `json:"start"`
	JSON  legalV1VerifyResponseCitationsSpanJSON `json:"-"`
}

// legalV1VerifyResponseCitationsSpanJSON contains the JSON metadata for the struct
// [LegalV1VerifyResponseCitationsSpan]
type legalV1VerifyResponseCitationsSpanJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalV1VerifyResponseCitationsSpan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1VerifyResponseCitationsSpanJSON) RawJSON() string {
	return r.raw
}

type LegalV1VerifyResponseCitationsStatus string

const (
	LegalV1VerifyResponseCitationsStatusVerified        LegalV1VerifyResponseCitationsStatus = "verified"
	LegalV1VerifyResponseCitationsStatusNotFound        LegalV1VerifyResponseCitationsStatus = "not_found"
	LegalV1VerifyResponseCitationsStatusMultipleMatches LegalV1VerifyResponseCitationsStatus = "multiple_matches"
)

func (r LegalV1VerifyResponseCitationsStatus) IsKnown() bool {
	switch r {
	case LegalV1VerifyResponseCitationsStatusVerified, LegalV1VerifyResponseCitationsStatusNotFound, LegalV1VerifyResponseCitationsStatusMultipleMatches:
		return true
	}
	return false
}

// Source of verification result (heuristic for fallback matches).
type LegalV1VerifyResponseCitationsVerificationSource string

const (
	LegalV1VerifyResponseCitationsVerificationSourceCourtlistener LegalV1VerifyResponseCitationsVerificationSource = "courtlistener"
	LegalV1VerifyResponseCitationsVerificationSourceHeuristic     LegalV1VerifyResponseCitationsVerificationSource = "heuristic"
)

func (r LegalV1VerifyResponseCitationsVerificationSource) IsKnown() bool {
	switch r {
	case LegalV1VerifyResponseCitationsVerificationSourceCourtlistener, LegalV1VerifyResponseCitationsVerificationSourceHeuristic:
		return true
	}
	return false
}

type LegalV1VerifyResponseSummary struct {
	// Citations with multiple possible matches
	MultipleMatches int64 `json:"multipleMatches"`
	// Citations not found in database
	NotFound int64 `json:"notFound"`
	// Total citations found
	Total int64 `json:"total"`
	// Citations verified against real cases
	Verified int64                            `json:"verified"`
	JSON     legalV1VerifyResponseSummaryJSON `json:"-"`
}

// legalV1VerifyResponseSummaryJSON contains the JSON metadata for the struct
// [LegalV1VerifyResponseSummary]
type legalV1VerifyResponseSummaryJSON struct {
	MultipleMatches apijson.Field
	NotFound        apijson.Field
	Total           apijson.Field
	Verified        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LegalV1VerifyResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalV1VerifyResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type LegalV1DocketParams struct {
	// Search dockets or look up a docket by ID
	Type param.Field[LegalV1DocketParamsType] `json:"type" api:"required"`
	// Optional CourtListener court slug (e.g. "nysd", "ca9", "cafc")
	Court param.Field[string] `json:"court"`
	// Optional lower bound for filing date (YYYY-MM-DD)
	DateFiledAfter param.Field[time.Time] `json:"dateFiledAfter" format:"date"`
	// Optional upper bound for filing date (YYYY-MM-DD)
	DateFiledBefore param.Field[time.Time] `json:"dateFiledBefore" format:"date"`
	// CourtListener docket ID (required for lookup)
	DocketID param.Field[string] `json:"docketId"`
	// Include docket entries/filings in lookup responses
	IncludeEntries param.Field[bool] `json:"includeEntries"`
	// Page size for search results or entry list (default 25 for search, 50 for
	// lookup)
	Limit param.Field[int64] `json:"limit"`
	// Reserved for future PACER live fetch support. Setting true currently
	// returns 400.
	Live param.Field[bool] `json:"live"`
	// Offset for search results or entry list
	Offset param.Field[int64] `json:"offset"`
	// Case name or party name search query (required for search)
	Query param.Field[string] `json:"query"`
}

func (r LegalV1DocketParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Search dockets or look up a docket by ID
type LegalV1DocketParamsType string

const (
	LegalV1DocketParamsTypeSearch LegalV1DocketParamsType = "search"
	LegalV1DocketParamsTypeLookup LegalV1DocketParamsType = "lookup"
)

func (r LegalV1DocketParamsType) IsKnown() bool {
	switch r {
	case LegalV1DocketParamsTypeSearch, LegalV1DocketParamsTypeLookup:
		return true
	}
	return false
}

type LegalV1DraftParams struct {
	// What to draft — the core task. E.g., "Motion to compel defendant to produce
	// discovery responses"
	Instructions param.Field[string] `json:"instructions" api:"required"`
	// Vault ID where the final document will be uploaded
	VaultID param.Field[string] `json:"vault_id" api:"required"`
	// Research and include legal citations
	Citations param.Field[bool] `json:"citations"`
	// Court or jurisdiction formatting hint. Triggers a legal-skills search. E.g.,
	// "California Superior Court", "SDNY", "federal pleading"
	Format param.Field[string] `json:"format"`
	// Target document length
	Length param.Field[LegalV1DraftParamsLength] `json:"length"`
	// LLM model override. Defaults to anthropic/claude-sonnet-4.6
	Model param.Field[string] `json:"model"`
	// Vault object IDs to use as source/reference documents
	ObjectIDs param.Field[[]string] `json:"object_ids"`
	// Filename for the output document. Auto-generated if omitted.
	OutputName param.Field[string] `json:"output_name"`
	// Output file format
	OutputType param.Field[LegalV1DraftParamsOutputType] `json:"output_type"`
	// Verify all citations in a loop — re-run verification and repair bad citations
	// until they pass
	Verified param.Field[bool] `json:"verified"`
}

func (r LegalV1DraftParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Target document length
type LegalV1DraftParamsLength struct {
	// Target value (e.g., 2000 words or 5 pages)
	Target param.Field[float64]                      `json:"target"`
	Unit   param.Field[LegalV1DraftParamsLengthUnit] `json:"unit"`
}

func (r LegalV1DraftParamsLength) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1DraftParamsLengthUnit string

const (
	LegalV1DraftParamsLengthUnitWords LegalV1DraftParamsLengthUnit = "words"
	LegalV1DraftParamsLengthUnitPages LegalV1DraftParamsLengthUnit = "pages"
)

func (r LegalV1DraftParamsLengthUnit) IsKnown() bool {
	switch r {
	case LegalV1DraftParamsLengthUnitWords, LegalV1DraftParamsLengthUnitPages:
		return true
	}
	return false
}

// Output file format
type LegalV1DraftParamsOutputType string

const (
	LegalV1DraftParamsOutputTypePdf  LegalV1DraftParamsOutputType = "pdf"
	LegalV1DraftParamsOutputTypeDocx LegalV1DraftParamsOutputType = "docx"
	LegalV1DraftParamsOutputTypeXlsx LegalV1DraftParamsOutputType = "xlsx"
	LegalV1DraftParamsOutputTypePptx LegalV1DraftParamsOutputType = "pptx"
	LegalV1DraftParamsOutputTypeMd   LegalV1DraftParamsOutputType = "md"
)

func (r LegalV1DraftParamsOutputType) IsKnown() bool {
	switch r {
	case LegalV1DraftParamsOutputTypePdf, LegalV1DraftParamsOutputTypeDocx, LegalV1DraftParamsOutputTypeXlsx, LegalV1DraftParamsOutputTypePptx, LegalV1DraftParamsOutputTypeMd:
		return true
	}
	return false
}

type LegalV1FindParams struct {
	// Search query (e.g., "fair use copyright", "Miranda rights")
	Query param.Field[string] `json:"query" api:"required"`
	// Optional jurisdiction ID from resolveJurisdiction (e.g., "california",
	// "us-federal")
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Number of results 1-25 (default: 10)
	NumResults param.Field[int64] `json:"numResults"`
}

func (r LegalV1FindParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1GetCitationsParams struct {
	// Text containing citations to extract. Can be a single citation (e.g., "531 U.S.
	// 98") or a full document with multiple citations.
	Text param.Field[string] `json:"text" api:"required"`
}

func (r LegalV1GetCitationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1GetCitationsFromURLParams struct {
	// URL of the legal document to extract citations from
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
}

func (r LegalV1GetCitationsFromURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1GetFullTextParams struct {
	// URL of the verified legal document
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Optional query to extract relevant highlights (e.g., "What is the holding?")
	HighlightQuery param.Field[string] `json:"highlightQuery"`
	// Maximum characters to return (default: 10000, max: 50000)
	MaxCharacters param.Field[int64] `json:"maxCharacters"`
	// Optional query for generating a summary (e.g., "Summarize the key ruling")
	SummaryQuery param.Field[string] `json:"summaryQuery"`
}

func (r LegalV1GetFullTextParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1ListCourtsParams struct {
	// Only return courts currently in use by CourtListener
	InUseOnly param.Field[bool] `json:"inUseOnly"`
	// Optional CourtListener jurisdiction code filter (e.g. FD, F, S)
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	Limit        param.Field[int64]  `json:"limit"`
	Offset       param.Field[int64]  `json:"offset"`
	// Search by court name or slug (e.g. "Northern District", "nysd", "ca9")
	Query param.Field[string] `json:"query"`
}

func (r LegalV1ListCourtsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1ListJurisdictionsParams struct {
	// Jurisdiction name (e.g., "California", "US Federal", "NY")
	Name param.Field[string] `json:"name" api:"required"`
}

func (r LegalV1ListJurisdictionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1PatentSearchParams struct {
	// Free-text search across all patent fields, or field-specific query (e.g.
	// "applicationMetaData.patentNumber:11234567"). Supports AND, OR, NOT operators.
	Query param.Field[string] `json:"query" api:"required"`
	// Filter by application status (e.g. "Patented Case", "Abandoned", "Pending")
	ApplicationStatus param.Field[string] `json:"applicationStatus"`
	// Filter by application type
	ApplicationType param.Field[LegalV1PatentSearchParamsApplicationType] `json:"applicationType"`
	// Filter by assignee/owner name (e.g. "Google LLC")
	Assignee param.Field[string] `json:"assignee"`
	// Start of filing date range (YYYY-MM-DD)
	FilingDateFrom param.Field[time.Time] `json:"filingDateFrom" format:"date"`
	// End of filing date range (YYYY-MM-DD)
	FilingDateTo param.Field[time.Time] `json:"filingDateTo" format:"date"`
	// Start of grant date range (YYYY-MM-DD)
	GrantDateFrom param.Field[time.Time] `json:"grantDateFrom" format:"date"`
	// End of grant date range (YYYY-MM-DD)
	GrantDateTo param.Field[time.Time] `json:"grantDateTo" format:"date"`
	// Filter by inventor name
	Inventor param.Field[string] `json:"inventor"`
	// Number of results to return (default 25, max 100)
	Limit param.Field[int64] `json:"limit"`
	// Starting position for pagination
	Offset param.Field[int64] `json:"offset"`
	// Field to sort results by
	SortBy param.Field[LegalV1PatentSearchParamsSortBy] `json:"sortBy"`
	// Sort order (default desc, newest first)
	SortOrder param.Field[LegalV1PatentSearchParamsSortOrder] `json:"sortOrder"`
}

func (r LegalV1PatentSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter by application type
type LegalV1PatentSearchParamsApplicationType string

const (
	LegalV1PatentSearchParamsApplicationTypeUtility     LegalV1PatentSearchParamsApplicationType = "Utility"
	LegalV1PatentSearchParamsApplicationTypeDesign      LegalV1PatentSearchParamsApplicationType = "Design"
	LegalV1PatentSearchParamsApplicationTypePlant       LegalV1PatentSearchParamsApplicationType = "Plant"
	LegalV1PatentSearchParamsApplicationTypeProvisional LegalV1PatentSearchParamsApplicationType = "Provisional"
	LegalV1PatentSearchParamsApplicationTypeReissue     LegalV1PatentSearchParamsApplicationType = "Reissue"
)

func (r LegalV1PatentSearchParamsApplicationType) IsKnown() bool {
	switch r {
	case LegalV1PatentSearchParamsApplicationTypeUtility, LegalV1PatentSearchParamsApplicationTypeDesign, LegalV1PatentSearchParamsApplicationTypePlant, LegalV1PatentSearchParamsApplicationTypeProvisional, LegalV1PatentSearchParamsApplicationTypeReissue:
		return true
	}
	return false
}

// Field to sort results by
type LegalV1PatentSearchParamsSortBy string

const (
	LegalV1PatentSearchParamsSortByFilingDate LegalV1PatentSearchParamsSortBy = "filingDate"
	LegalV1PatentSearchParamsSortByGrantDate  LegalV1PatentSearchParamsSortBy = "grantDate"
)

func (r LegalV1PatentSearchParamsSortBy) IsKnown() bool {
	switch r {
	case LegalV1PatentSearchParamsSortByFilingDate, LegalV1PatentSearchParamsSortByGrantDate:
		return true
	}
	return false
}

// Sort order (default desc, newest first)
type LegalV1PatentSearchParamsSortOrder string

const (
	LegalV1PatentSearchParamsSortOrderAsc  LegalV1PatentSearchParamsSortOrder = "asc"
	LegalV1PatentSearchParamsSortOrderDesc LegalV1PatentSearchParamsSortOrder = "desc"
)

func (r LegalV1PatentSearchParamsSortOrder) IsKnown() bool {
	switch r {
	case LegalV1PatentSearchParamsSortOrderAsc, LegalV1PatentSearchParamsSortOrderDesc:
		return true
	}
	return false
}

type LegalV1ResearchParams struct {
	// Primary search query
	Query param.Field[string] `json:"query" api:"required"`
	// Additional query variations to search (e.g., different phrasings of the legal
	// issue)
	AdditionalQueries param.Field[[]string] `json:"additionalQueries"`
	// Optional jurisdiction ID from resolveJurisdiction
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Number of results 1-25 (default: 10)
	NumResults param.Field[int64] `json:"numResults"`
}

func (r LegalV1ResearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1SimilarParams struct {
	// URL of a legal document to find similar sources for
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Optional jurisdiction ID to filter results
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Number of results 1-25 (default: 10)
	NumResults param.Field[int64] `json:"numResults"`
	// Optional ISO date to find only newer documents (e.g., "2020-01-01")
	StartPublishedDate param.Field[time.Time] `json:"startPublishedDate" format:"date"`
}

func (r LegalV1SimilarParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1TrademarkSearchParams struct {
	// USPTO registration number (e.g. "6123456"). Provide either serialNumber or
	// registrationNumber.
	RegistrationNumber param.Field[string] `json:"registrationNumber"`
	// USPTO serial number (e.g. "97123456"). Provide either serialNumber or
	// registrationNumber.
	SerialNumber param.Field[string] `json:"serialNumber"`
}

func (r LegalV1TrademarkSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalV1VerifyParams struct {
	// Text containing citations to verify. Can be a single citation (e.g., "531 U.S.
	// 98") or a full document with multiple citations.
	Text param.Field[string] `json:"text" api:"required"`
}

func (r LegalV1VerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
