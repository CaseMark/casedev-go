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

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1TypeService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1TypeService] method instead.
type MatterV1TypeService struct {
	Options []option.RequestOption
}

// NewMatterV1TypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1TypeService(opts ...option.RequestOption) (r *MatterV1TypeService) {
	r = &MatterV1TypeService{}
	r.Options = opts
	return
}

// Create a matter type with plain-English operating instructions and seeded work.
func (r *MatterV1TypeService) New(ctx context.Context, body MatterV1TypeNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get a single matter type.
func (r *MatterV1TypeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/types/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Update a matter type.
func (r *MatterV1TypeService) Update(ctx context.Context, id string, body MatterV1TypeUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/types/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List matter types for the authenticated organization.
func (r *MatterV1TypeService) List(ctx context.Context, query MatterV1TypeListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type MatterV1TypeNewParams struct {
	Name               param.Field[string]                                 `json:"name" api:"required"`
	DefaultAgentTypeID param.Field[string]                                 `json:"default_agent_type_id"`
	DefaultMetadata    param.Field[map[string]interface{}]                 `json:"default_metadata"`
	DefaultWorkItems   param.Field[[]map[string]interface{}]               `json:"default_work_items"`
	Description        param.Field[string]                                 `json:"description"`
	ExitCriteria       param.Field[[]string]                               `json:"exit_criteria"`
	Instructions       param.Field[string]                                 `json:"instructions"`
	IntakeRequirements param.Field[[]string]                               `json:"intake_requirements"`
	IsActive           param.Field[bool]                                   `json:"is_active"`
	OrchestrationMode  param.Field[MatterV1TypeNewParamsOrchestrationMode] `json:"orchestration_mode"`
	ReviewAgentTypeID  param.Field[string]                                 `json:"review_agent_type_id"`
	ReviewCriteria     param.Field[[]string]                               `json:"review_criteria"`
	SkillRefs          param.Field[[]string]                               `json:"skill_refs"`
	Slug               param.Field[string]                                 `json:"slug"`
}

func (r MatterV1TypeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1TypeNewParamsOrchestrationMode string

const (
	MatterV1TypeNewParamsOrchestrationModeAuto  MatterV1TypeNewParamsOrchestrationMode = "auto"
	MatterV1TypeNewParamsOrchestrationModeHuman MatterV1TypeNewParamsOrchestrationMode = "human"
)

func (r MatterV1TypeNewParamsOrchestrationMode) IsKnown() bool {
	switch r {
	case MatterV1TypeNewParamsOrchestrationModeAuto, MatterV1TypeNewParamsOrchestrationModeHuman:
		return true
	}
	return false
}

type MatterV1TypeUpdateParams struct {
	DefaultAgentTypeID param.Field[string]                                    `json:"default_agent_type_id"`
	DefaultMetadata    param.Field[map[string]interface{}]                    `json:"default_metadata"`
	DefaultWorkItems   param.Field[[]map[string]interface{}]                  `json:"default_work_items"`
	Description        param.Field[string]                                    `json:"description"`
	ExitCriteria       param.Field[[]string]                                  `json:"exit_criteria"`
	Instructions       param.Field[string]                                    `json:"instructions"`
	IntakeRequirements param.Field[[]string]                                  `json:"intake_requirements"`
	IsActive           param.Field[bool]                                      `json:"is_active"`
	Name               param.Field[string]                                    `json:"name"`
	OrchestrationMode  param.Field[MatterV1TypeUpdateParamsOrchestrationMode] `json:"orchestration_mode"`
	ReviewAgentTypeID  param.Field[string]                                    `json:"review_agent_type_id"`
	ReviewCriteria     param.Field[[]string]                                  `json:"review_criteria"`
	SkillRefs          param.Field[[]string]                                  `json:"skill_refs"`
	Slug               param.Field[string]                                    `json:"slug"`
}

func (r MatterV1TypeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1TypeUpdateParamsOrchestrationMode string

const (
	MatterV1TypeUpdateParamsOrchestrationModeAuto  MatterV1TypeUpdateParamsOrchestrationMode = "auto"
	MatterV1TypeUpdateParamsOrchestrationModeHuman MatterV1TypeUpdateParamsOrchestrationMode = "human"
)

func (r MatterV1TypeUpdateParamsOrchestrationMode) IsKnown() bool {
	switch r {
	case MatterV1TypeUpdateParamsOrchestrationModeAuto, MatterV1TypeUpdateParamsOrchestrationModeHuman:
		return true
	}
	return false
}

type MatterV1TypeListParams struct {
	Active param.Field[bool] `query:"active"`
}

// URLQuery serializes [MatterV1TypeListParams]'s query parameters as `url.Values`.
func (r MatterV1TypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
