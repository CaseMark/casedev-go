// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
	"github.com/tidwall/gjson"
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
	// Search and read legal AI skills for agents
	Custom *SkillCustomService
}

// NewSkillService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSkillService(opts ...option.RequestOption) (r *SkillService) {
	r = &SkillService{}
	r.Options = opts
	r.Custom = NewSkillCustomService(opts...)
	return
}

// Create an org-scoped custom skill. The skill will be searchable via
// /skills/resolve alongside curated skills.
func (r *SkillService) New(ctx context.Context, body SkillNewParams, opts ...option.RequestOption) (res *SkillNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skills"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an org-scoped custom skill by slug. Only provided fields are updated.
// Version is auto-incremented.
func (r *SkillService) Update(ctx context.Context, slug string, body SkillUpdateParams, opts ...option.RequestOption) (res *SkillUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("skills/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Soft-delete an org-scoped custom skill by slug. The skill will no longer appear
// in search results.
func (r *SkillService) Delete(ctx context.Context, slug string, opts ...option.RequestOption) (res *SkillDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("skills/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Read the full content of a legal skill by its slug. Returns markdown content,
// tags, and metadata.
func (r *SkillService) Read(ctx context.Context, slug string, opts ...option.RequestOption) (res *SkillReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("skills/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search the Legal Skills Store using hybrid search (text + tag + semantic).
// Returns ranked results with relevance scores.
func (r *SkillService) Resolve(ctx context.Context, query SkillResolveParams, opts ...option.RequestOption) (res *SkillResolveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skills/resolve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ReadResponseFileBundle struct {
	Path        string                     `json:"path" api:"required"`
	Role        ReadResponseFileBundleRole `json:"role" api:"required"`
	RootSlug    string                     `json:"root_slug" api:"required"`
	ContentType string                     `json:"content_type" api:"nullable"`
	JSON        readResponseFileBundleJSON `json:"-"`
}

// readResponseFileBundleJSON contains the JSON metadata for the struct
// [ReadResponseFileBundle]
type readResponseFileBundleJSON struct {
	Path        apijson.Field
	Role        apijson.Field
	RootSlug    apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReadResponseFileBundle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r readResponseFileBundleJSON) RawJSON() string {
	return r.raw
}

func (r ReadResponseFileBundle) implementsSkillReadResponseBundle() {}

type ReadResponseFileBundleRole string

const (
	ReadResponseFileBundleRoleFile ReadResponseFileBundleRole = "file"
)

func (r ReadResponseFileBundleRole) IsKnown() bool {
	switch r {
	case ReadResponseFileBundleRoleFile:
		return true
	}
	return false
}

type ReadResponseRootBundle struct {
	Files []ReadResponseRootBundleFile `json:"files" api:"required"`
	Role  ReadResponseRootBundleRole   `json:"role" api:"required"`
	JSON  readResponseRootBundleJSON   `json:"-"`
}

// readResponseRootBundleJSON contains the JSON metadata for the struct
// [ReadResponseRootBundle]
type readResponseRootBundleJSON struct {
	Files       apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReadResponseRootBundle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r readResponseRootBundleJSON) RawJSON() string {
	return r.raw
}

func (r ReadResponseRootBundle) implementsSkillReadResponseBundle() {}

type ReadResponseRootBundleFile struct {
	Path        string                         `json:"path" api:"required"`
	Slug        string                         `json:"slug" api:"required"`
	ContentType string                         `json:"content_type" api:"nullable"`
	Name        string                         `json:"name" api:"nullable"`
	JSON        readResponseRootBundleFileJSON `json:"-"`
}

// readResponseRootBundleFileJSON contains the JSON metadata for the struct
// [ReadResponseRootBundleFile]
type readResponseRootBundleFileJSON struct {
	Path        apijson.Field
	Slug        apijson.Field
	ContentType apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReadResponseRootBundleFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r readResponseRootBundleFileJSON) RawJSON() string {
	return r.raw
}

type ReadResponseRootBundleRole string

const (
	ReadResponseRootBundleRoleRoot ReadResponseRootBundleRole = "root"
)

func (r ReadResponseRootBundleRole) IsKnown() bool {
	switch r {
	case ReadResponseRootBundleRoleRoot:
		return true
	}
	return false
}

type SkillNewResponse struct {
	Content   string               `json:"content"`
	CreatedAt time.Time            `json:"created_at" format:"date-time"`
	Metadata  interface{}          `json:"metadata"`
	Name      string               `json:"name"`
	Slug      string               `json:"slug"`
	Summary   string               `json:"summary" api:"nullable"`
	Tags      []string             `json:"tags"`
	Version   int64                `json:"version"`
	JSON      skillNewResponseJSON `json:"-"`
}

// skillNewResponseJSON contains the JSON metadata for the struct
// [SkillNewResponse]
type skillNewResponseJSON struct {
	Content     apijson.Field
	CreatedAt   apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Summary     apijson.Field
	Tags        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillNewResponseJSON) RawJSON() string {
	return r.raw
}

type SkillUpdateResponse struct {
	Content   string                  `json:"content"`
	Metadata  interface{}             `json:"metadata"`
	Name      string                  `json:"name"`
	Slug      string                  `json:"slug"`
	Summary   string                  `json:"summary" api:"nullable"`
	Tags      []string                `json:"tags"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	Version   int64                   `json:"version"`
	JSON      skillUpdateResponseJSON `json:"-"`
}

// skillUpdateResponseJSON contains the JSON metadata for the struct
// [SkillUpdateResponse]
type skillUpdateResponseJSON struct {
	Content     apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Summary     apijson.Field
	Tags        apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SkillDeleteResponse struct {
	Deleted bool                    `json:"deleted"`
	Slug    string                  `json:"slug"`
	JSON    skillDeleteResponseJSON `json:"-"`
}

// skillDeleteResponseJSON contains the JSON metadata for the struct
// [SkillDeleteResponse]
type skillDeleteResponseJSON struct {
	Deleted     apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SkillReadResponse struct {
	// Skill author
	AuthorName string `json:"author_name"`
	// Skill bundle metadata for root skills and companion file rows
	Bundle SkillReadResponseBundle `json:"bundle" api:"nullable"`
	// Full skill content in markdown
	Content string `json:"content"`
	// Skill license
	License string `json:"license"`
	// Custom metadata (custom skills only)
	Metadata interface{} `json:"metadata"`
	// Skill name
	Name string `json:"name"`
	// Unique skill identifier
	Slug string `json:"slug"`
	// Skill source (authenticated requests only)
	Source SkillReadResponseSource `json:"source"`
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
	Bundle      apijson.Field
	Content     apijson.Field
	License     apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Source      apijson.Field
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

// Skill bundle metadata for root skills and companion file rows
type SkillReadResponseBundle struct {
	Role        SkillReadResponseBundleRole `json:"role" api:"required"`
	ContentType string                      `json:"content_type" api:"nullable"`
	// This field can have the runtime type of [[]ReadResponseRootBundleFile].
	Files    interface{}                 `json:"files"`
	Path     string                      `json:"path"`
	RootSlug string                      `json:"root_slug"`
	JSON     skillReadResponseBundleJSON `json:"-"`
	union    SkillReadResponseBundleUnion
}

// skillReadResponseBundleJSON contains the JSON metadata for the struct
// [SkillReadResponseBundle]
type skillReadResponseBundleJSON struct {
	Role        apijson.Field
	ContentType apijson.Field
	Files       apijson.Field
	Path        apijson.Field
	RootSlug    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r skillReadResponseBundleJSON) RawJSON() string {
	return r.raw
}

func (r *SkillReadResponseBundle) UnmarshalJSON(data []byte) (err error) {
	*r = SkillReadResponseBundle{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SkillReadResponseBundleUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ReadResponseRootBundle],
// [ReadResponseFileBundle].
func (r SkillReadResponseBundle) AsUnion() SkillReadResponseBundleUnion {
	return r.union
}

// Skill bundle metadata for root skills and companion file rows
//
// Union satisfied by [ReadResponseRootBundle] or [ReadResponseFileBundle].
type SkillReadResponseBundleUnion interface {
	implementsSkillReadResponseBundle()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SkillReadResponseBundleUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReadResponseRootBundle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReadResponseFileBundle{}),
		},
	)
}

type SkillReadResponseBundleRole string

const (
	SkillReadResponseBundleRoleRoot SkillReadResponseBundleRole = "root"
	SkillReadResponseBundleRoleFile SkillReadResponseBundleRole = "file"
)

func (r SkillReadResponseBundleRole) IsKnown() bool {
	switch r {
	case SkillReadResponseBundleRoleRoot, SkillReadResponseBundleRoleFile:
		return true
	}
	return false
}

// Skill source (authenticated requests only)
type SkillReadResponseSource string

const (
	SkillReadResponseSourceCurated SkillReadResponseSource = "curated"
	SkillReadResponseSourceCustom  SkillReadResponseSource = "custom"
)

func (r SkillReadResponseSource) IsKnown() bool {
	switch r {
	case SkillReadResponseSourceCurated, SkillReadResponseSourceCustom:
		return true
	}
	return false
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
	// Whether the skill is curated or org-custom
	Source SkillResolveResponseResultsSource `json:"source"`
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
	Source      apijson.Field
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

// Whether the skill is curated or org-custom
type SkillResolveResponseResultsSource string

const (
	SkillResolveResponseResultsSourceCurated SkillResolveResponseResultsSource = "curated"
	SkillResolveResponseResultsSourceCustom  SkillResolveResponseResultsSource = "custom"
)

func (r SkillResolveResponseResultsSource) IsKnown() bool {
	switch r {
	case SkillResolveResponseResultsSourceCurated, SkillResolveResponseResultsSourceCustom:
		return true
	}
	return false
}

type SkillNewParams struct {
	// Full skill content in markdown
	Content param.Field[string] `json:"content" api:"required"`
	// Skill name
	Name param.Field[string] `json:"name" api:"required"`
	// Arbitrary metadata (author, license, etc.)
	Metadata param.Field[interface{}] `json:"metadata"`
	// URL-safe slug. Auto-generated from name if omitted.
	Slug param.Field[string] `json:"slug"`
	// Brief description (1-2 sentences)
	Summary param.Field[string] `json:"summary"`
	// Tags for categorization and search boosting
	Tags param.Field[[]string] `json:"tags"`
}

func (r SkillNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SkillUpdateParams struct {
	Content  param.Field[string]      `json:"content"`
	Metadata param.Field[interface{}] `json:"metadata"`
	Name     param.Field[string]      `json:"name"`
	// New slug (renames the skill)
	Slug    param.Field[string]   `json:"slug"`
	Summary param.Field[string]   `json:"summary"`
	Tags    param.Field[[]string] `json:"tags"`
}

func (r SkillUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
