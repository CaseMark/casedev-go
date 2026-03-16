// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
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

// Search and read legal AI skills for agents
//
// SkillCustomService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkillCustomService] method instead.
type SkillCustomService struct {
	Options []option.RequestOption
}

// NewSkillCustomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkillCustomService(opts ...option.RequestOption) (r *SkillCustomService) {
	r = &SkillCustomService{}
	r.Options = opts
	return
}

// List all custom skills for the authenticated organization. Supports cursor-based
// pagination.
func (r *SkillCustomService) List(ctx context.Context, query SkillCustomListParams, opts ...option.RequestOption) (res *SkillCustomListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skills/custom"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SkillCustomListResponse struct {
	HasMore    bool                           `json:"has_more"`
	NextCursor string                         `json:"next_cursor" api:"nullable"`
	Skills     []SkillCustomListResponseSkill `json:"skills"`
	JSON       skillCustomListResponseJSON    `json:"-"`
}

// skillCustomListResponseJSON contains the JSON metadata for the struct
// [SkillCustomListResponse]
type skillCustomListResponseJSON struct {
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Skills      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillCustomListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillCustomListResponseJSON) RawJSON() string {
	return r.raw
}

type SkillCustomListResponseSkill struct {
	CreatedAt time.Time                        `json:"created_at" format:"date-time"`
	Metadata  interface{}                      `json:"metadata"`
	Name      string                           `json:"name"`
	Slug      string                           `json:"slug"`
	Summary   string                           `json:"summary" api:"nullable"`
	Tags      []string                         `json:"tags"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	Version   int64                            `json:"version"`
	JSON      skillCustomListResponseSkillJSON `json:"-"`
}

// skillCustomListResponseSkillJSON contains the JSON metadata for the struct
// [SkillCustomListResponseSkill]
type skillCustomListResponseSkillJSON struct {
	CreatedAt   apijson.Field
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

func (r *SkillCustomListResponseSkill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillCustomListResponseSkillJSON) RawJSON() string {
	return r.raw
}

type SkillCustomListParams struct {
	// Cursor for pagination (skill ID from previous page)
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results (1-100)
	Limit param.Field[int64] `query:"limit"`
	// Filter by tag
	Tag param.Field[string] `query:"tag"`
}

// URLQuery serializes [SkillCustomListParams]'s query parameters as `url.Values`.
func (r SkillCustomListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
