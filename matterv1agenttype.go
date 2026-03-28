// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1AgentTypeService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1AgentTypeService] method instead.
type MatterV1AgentTypeService struct {
	Options []option.RequestOption
}

// NewMatterV1AgentTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatterV1AgentTypeService(opts ...option.RequestOption) (r *MatterV1AgentTypeService) {
	r = &MatterV1AgentTypeService{}
	r.Options = opts
	return
}

// Create a reusable agent role for legal matter orchestration.
func (r *MatterV1AgentTypeService) New(ctx context.Context, body MatterV1AgentTypeNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1/agent-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// List reusable agent roles for the authenticated organization.
func (r *MatterV1AgentTypeService) List(ctx context.Context, query MatterV1AgentTypeListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1/agent-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type MatterV1AgentTypeNewParams struct {
	Instructions  param.Field[string]                 `json:"instructions" api:"required"`
	Name          param.Field[string]                 `json:"name" api:"required"`
	Description   param.Field[string]                 `json:"description"`
	DisabledTools param.Field[[]string]               `json:"disabled_tools"`
	EnabledTools  param.Field[[]string]               `json:"enabled_tools"`
	IsActive      param.Field[bool]                   `json:"is_active"`
	IsDefault     param.Field[bool]                   `json:"is_default"`
	Metadata      param.Field[map[string]interface{}] `json:"metadata"`
	Model         param.Field[string]                 `json:"model"`
	SkillRefs     param.Field[[]string]               `json:"skill_refs"`
	Slug          param.Field[string]                 `json:"slug"`
}

func (r MatterV1AgentTypeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1AgentTypeListParams struct {
	Active param.Field[bool] `query:"active"`
}

// URLQuery serializes [MatterV1AgentTypeListParams]'s query parameters as
// `url.Values`.
func (r MatterV1AgentTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
