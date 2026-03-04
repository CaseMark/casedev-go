// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Search and read legal AI skills for agents
//
// SkillService contains methods and other services that help with interacting with
// the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkillService] method instead.
type SkillService struct {
	Options []option.RequestOption
}

// NewSkillService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSkillService(opts ...option.RequestOption) (r *SkillService) {
	r = &SkillService{}
	r.Options = opts
	return
}

// Read the full content of a legal skill by its slug. Returns markdown content,
// tags, and metadata.
func (r *SkillService) Read(ctx context.Context, slug string, opts ...option.RequestOption) (res *SkillReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("skills/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Search the Legal Skills Store using hybrid search (text + tag + semantic).
// Returns ranked results with relevance scores.
func (r *SkillService) Resolve(ctx context.Context, query SkillResolveParams, opts ...option.RequestOption) (res *SkillResolveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skills/resolve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SkillReadResponse struct {
	// Skill author
	AuthorName string `json:"author_name"`
	// Full skill content in markdown
	Content string `json:"content"`
	// Skill license
	License string `json:"license"`
	// Skill name
	Name string `json:"name"`
	// Unique skill identifier
	Slug string `json:"slug"`
	// Brief skill description
	Summary string `json:"summary"`
	// Skill tags
	Tags []string `json:"tags"`
	// Skill version
	Version string                `json:"version"`
	JSON    skillReadResponseJSON `json:"-"`
}

// skillReadResponseJSON contains the JSON metadata for the struct
// [SkillReadResponse]
type skillReadResponseJSON struct {
	AuthorName  apijson.Field
	Content     apijson.Field
	License     apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Summary     apijson.Field
	Tags        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillReadResponseJSON) RawJSON() string {
	return r.raw
}

type SkillResolveResponse struct {
	// Search methods used (text, tag, semantic)
	MethodsUsed []string                     `json:"methods_used"`
	Results     []SkillResolveResponseResult `json:"results"`
	JSON        skillResolveResponseJSON     `json:"-"`
}

// skillResolveResponseJSON contains the JSON metadata for the struct
// [SkillResolveResponse]
type skillResolveResponseJSON struct {
	MethodsUsed apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillResolveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillResolveResponseJSON) RawJSON() string {
	return r.raw
}

type SkillResolveResponseResult struct {
	// Skill name
	Name string `json:"name"`
	// Relevance score
	Score float64 `json:"score"`
	// Unique skill identifier
	Slug string `json:"slug"`
	// Brief skill description
	Summary string `json:"summary"`
	// Skill tags
	Tags []string                       `json:"tags"`
	JSON skillResolveResponseResultJSON `json:"-"`
}

// skillResolveResponseResultJSON contains the JSON metadata for the struct
// [SkillResolveResponseResult]
type skillResolveResponseResultJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	Slug        apijson.Field
	Summary     apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillResolveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillResolveResponseResultJSON) RawJSON() string {
	return r.raw
}

type SkillResolveParams struct {
	// Search query string
	Q param.Field[string] `query:"q" api:"required"`
	// Maximum number of results to return (1-20)
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [SkillResolveParams]'s query parameters as `url.Values`.
func (r SkillResolveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
