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

// Web application deployment management
//
// ApplicationV1DeploymentService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationV1DeploymentService] method instead.
type ApplicationV1DeploymentService struct {
	Options []option.RequestOption
}

// NewApplicationV1DeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationV1DeploymentService(opts ...option.RequestOption) (r *ApplicationV1DeploymentService) {
	r = &ApplicationV1DeploymentService{}
	r.Options = opts
	return
}

// Creates a deployment for an existing project by fetching repository files from
// GitHub and uploading them to the hosting provider. Use ref to deploy a branch,
// tag, or commit other than the project default branch.
func (r *ApplicationV1DeploymentService) New(ctx context.Context, body ApplicationV1DeploymentNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "applications/v1/deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Returns deployment details for one project in the authenticated organization.
// Set includeLogs=true to include recent build output in the response.
func (r *ApplicationV1DeploymentService) Get(ctx context.Context, id string, query ApplicationV1DeploymentGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("applications/v1/deployments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Lists recent deployments for one project in the authenticated organization. Use
// the optional filters to narrow results by target or deployment state.
func (r *ApplicationV1DeploymentService) List(ctx context.Context, query ApplicationV1DeploymentListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "applications/v1/deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Cancels a running deployment after verifying that the referenced project belongs
// to the authenticated organization. Use this when a build is stuck,
// misconfigured, or no longer needed.
func (r *ApplicationV1DeploymentService) Cancel(ctx context.Context, id string, body ApplicationV1DeploymentCancelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("applications/v1/deployments/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Create a deployment from raw file contents (for Thurgood sandbox deployments)
func (r *ApplicationV1DeploymentService) NewFromFiles(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "applications/v1/deployments/from-files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Returns build and runtime log events for a deployment after verifying access to
// the owning project. Use this when you need detailed output for a failed or
// in-progress build.
func (r *ApplicationV1DeploymentService) GetLogs(ctx context.Context, id string, query ApplicationV1DeploymentGetLogsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("applications/v1/deployments/%s/logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Returns the current status of a deployment without fetching full build logs. Use
// this endpoint for lightweight polling while a deployment is building or waiting
// to become ready.
func (r *ApplicationV1DeploymentService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("applications/v1/deployments/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Stream real-time deployment progress events via Server-Sent Events
func (r *ApplicationV1DeploymentService) Stream(ctx context.Context, id string, query ApplicationV1DeploymentStreamParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("applications/v1/deployments/%s/stream", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type ApplicationV1DeploymentNewParams struct {
	// Project ID to deploy
	ProjectID param.Field[string] `json:"projectId" api:"required"`
	// Git branch, tag, or commit to deploy. Defaults to the project branch.
	Ref param.Field[string] `json:"ref"`
	// Deployment target environment
	Target param.Field[ApplicationV1DeploymentNewParamsTarget] `json:"target"`
}

func (r ApplicationV1DeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Deployment target environment
type ApplicationV1DeploymentNewParamsTarget string

const (
	ApplicationV1DeploymentNewParamsTargetProduction ApplicationV1DeploymentNewParamsTarget = "production"
	ApplicationV1DeploymentNewParamsTargetPreview    ApplicationV1DeploymentNewParamsTarget = "preview"
)

func (r ApplicationV1DeploymentNewParamsTarget) IsKnown() bool {
	switch r {
	case ApplicationV1DeploymentNewParamsTargetProduction, ApplicationV1DeploymentNewParamsTargetPreview:
		return true
	}
	return false
}

type ApplicationV1DeploymentGetParams struct {
	// Project ID used to verify access to the deployment
	ProjectID param.Field[string] `query:"projectId" api:"required"`
	// Whether to include build logs in the response
	IncludeLogs param.Field[bool] `query:"includeLogs"`
}

// URLQuery serializes [ApplicationV1DeploymentGetParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1DeploymentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ApplicationV1DeploymentListParams struct {
	// Project ID to list deployments for
	ProjectID param.Field[string] `query:"projectId" api:"required"`
	// Maximum number of deployments to return
	Limit param.Field[float64] `query:"limit"`
	// Deployment state to filter by
	State param.Field[string] `query:"state"`
	// Deployment target to filter by
	Target param.Field[ApplicationV1DeploymentListParamsTarget] `query:"target"`
}

// URLQuery serializes [ApplicationV1DeploymentListParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1DeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Deployment target to filter by
type ApplicationV1DeploymentListParamsTarget string

const (
	ApplicationV1DeploymentListParamsTargetProduction ApplicationV1DeploymentListParamsTarget = "production"
	ApplicationV1DeploymentListParamsTargetStaging    ApplicationV1DeploymentListParamsTarget = "staging"
)

func (r ApplicationV1DeploymentListParamsTarget) IsKnown() bool {
	switch r {
	case ApplicationV1DeploymentListParamsTargetProduction, ApplicationV1DeploymentListParamsTargetStaging:
		return true
	}
	return false
}

type ApplicationV1DeploymentCancelParams struct {
	// Project ID used to verify access to the deployment
	ProjectID param.Field[string] `json:"projectId" api:"required"`
}

func (r ApplicationV1DeploymentCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1DeploymentGetLogsParams struct {
	// Project ID used to verify access to the deployment
	ProjectID param.Field[string] `query:"projectId" api:"required"`
}

// URLQuery serializes [ApplicationV1DeploymentGetLogsParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1DeploymentGetLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ApplicationV1DeploymentStreamParams struct {
	// Project ID (for authorization)
	ProjectID param.Field[string] `query:"projectId" api:"required"`
	// Resume stream from this index (for reconnection)
	StartIndex param.Field[float64] `query:"startIndex"`
}

// URLQuery serializes [ApplicationV1DeploymentStreamParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1DeploymentStreamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
