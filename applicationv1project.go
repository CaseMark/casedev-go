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

// ApplicationV1ProjectService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationV1ProjectService] method instead.
type ApplicationV1ProjectService struct {
	Options []option.RequestOption
}

// NewApplicationV1ProjectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationV1ProjectService(opts ...option.RequestOption) (r *ApplicationV1ProjectService) {
	r = &ApplicationV1ProjectService{}
	r.Options = opts
	return
}

// Create a new web application project
func (r *ApplicationV1ProjectService) New(ctx context.Context, body ApplicationV1ProjectNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "applications/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get details of a specific web application project
func (r *ApplicationV1ProjectService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// List all web application projects
func (r *ApplicationV1ProjectService) List(ctx context.Context, opts ...option.RequestOption) (res *ApplicationV1ProjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "applications/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a web application project
func (r *ApplicationV1ProjectService) Delete(ctx context.Context, id string, body ApplicationV1ProjectDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Trigger a new deployment for a project.
func (r *ApplicationV1ProjectService) NewDeployment(ctx context.Context, id string, body ApplicationV1ProjectNewDeploymentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/deployments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Add a custom domain to a project
func (r *ApplicationV1ProjectService) NewDomain(ctx context.Context, id string, body ApplicationV1ProjectNewDomainParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/domains", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create a new environment variable for a project
func (r *ApplicationV1ProjectService) NewEnv(ctx context.Context, id string, body ApplicationV1ProjectNewEnvParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/env", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Remove a domain from a project
func (r *ApplicationV1ProjectService) DeleteDomain(ctx context.Context, id string, domain string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/domains/%s", id, domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete an environment variable from a project
func (r *ApplicationV1ProjectService) DeleteEnv(ctx context.Context, id string, envID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if envID == "" {
		err = errors.New("missing required envId parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/env/%s", id, envID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get runtime/function logs for a project
func (r *ApplicationV1ProjectService) GetRuntimeLogs(ctx context.Context, id string, query ApplicationV1ProjectGetRuntimeLogsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/runtime-logs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// List deployments for a specific project
func (r *ApplicationV1ProjectService) ListDeployments(ctx context.Context, id string, query ApplicationV1ProjectListDeploymentsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/deployments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// List all domains configured for a project
func (r *ApplicationV1ProjectService) ListDomains(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/domains", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// List all environment variables for a project (values are hidden unless
// decrypt=true)
func (r *ApplicationV1ProjectService) ListEnv(ctx context.Context, id string, query ApplicationV1ProjectListEnvParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("applications/v1/projects/%s/env", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ApplicationV1ProjectListResponse struct {
	Projects []ApplicationV1ProjectListResponseProject `json:"projects"`
	JSON     applicationV1ProjectListResponseJSON      `json:"-"`
}

// applicationV1ProjectListResponseJSON contains the JSON metadata for the struct
// [ApplicationV1ProjectListResponse]
type applicationV1ProjectListResponseJSON struct {
	Projects    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationV1ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationV1ProjectListResponseJSON) RawJSON() string {
	return r.raw
}

type ApplicationV1ProjectListResponseProject struct {
	ID              string                                           `json:"id"`
	CreatedAt       string                                           `json:"createdAt"`
	Domains         []ApplicationV1ProjectListResponseProjectsDomain `json:"domains"`
	Framework       string                                           `json:"framework"`
	GitBranch       string                                           `json:"gitBranch"`
	GitRepo         string                                           `json:"gitRepo"`
	Name            string                                           `json:"name"`
	Status          string                                           `json:"status"`
	UpdatedAt       string                                           `json:"updatedAt"`
	VercelProjectID string                                           `json:"vercelProjectId"`
	JSON            applicationV1ProjectListResponseProjectJSON      `json:"-"`
}

// applicationV1ProjectListResponseProjectJSON contains the JSON metadata for the
// struct [ApplicationV1ProjectListResponseProject]
type applicationV1ProjectListResponseProjectJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Domains         apijson.Field
	Framework       apijson.Field
	GitBranch       apijson.Field
	GitRepo         apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	UpdatedAt       apijson.Field
	VercelProjectID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ApplicationV1ProjectListResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationV1ProjectListResponseProjectJSON) RawJSON() string {
	return r.raw
}

type ApplicationV1ProjectListResponseProjectsDomain struct {
	ID         string                                             `json:"id"`
	Domain     string                                             `json:"domain"`
	IsPrimary  bool                                               `json:"isPrimary"`
	IsVerified bool                                               `json:"isVerified"`
	JSON       applicationV1ProjectListResponseProjectsDomainJSON `json:"-"`
}

// applicationV1ProjectListResponseProjectsDomainJSON contains the JSON metadata
// for the struct [ApplicationV1ProjectListResponseProjectsDomain]
type applicationV1ProjectListResponseProjectsDomainJSON struct {
	ID          apijson.Field
	Domain      apijson.Field
	IsPrimary   apijson.Field
	IsVerified  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationV1ProjectListResponseProjectsDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationV1ProjectListResponseProjectsDomainJSON) RawJSON() string {
	return r.raw
}

type ApplicationV1ProjectNewParams struct {
	// GitHub repository URL or "owner/repo"
	GitRepo param.Field[string] `json:"gitRepo" api:"required"`
	// Project name
	Name param.Field[string] `json:"name" api:"required"`
	// Custom build command
	BuildCommand param.Field[string] `json:"buildCommand"`
	// Environment variables to set on the project
	EnvironmentVariables param.Field[[]ApplicationV1ProjectNewParamsEnvironmentVariable] `json:"environmentVariables"`
	// Framework (e.g., "nextjs", "remix", "astro")
	Framework param.Field[string] `json:"framework"`
	// Git branch to deploy
	GitBranch param.Field[string] `json:"gitBranch"`
	// Custom install command
	InstallCommand param.Field[string] `json:"installCommand"`
	// Build output directory
	OutputDirectory param.Field[string] `json:"outputDirectory"`
	// Root directory of the project
	RootDirectory param.Field[string] `json:"rootDirectory"`
}

func (r ApplicationV1ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1ProjectNewParamsEnvironmentVariable struct {
	// Environment variable name
	Key param.Field[string] `json:"key" api:"required"`
	// Deployment targets for this variable
	Target param.Field[[]ApplicationV1ProjectNewParamsEnvironmentVariablesTarget] `json:"target" api:"required"`
	// Environment variable value
	Value param.Field[string] `json:"value" api:"required"`
	// Variable type
	Type param.Field[ApplicationV1ProjectNewParamsEnvironmentVariablesType] `json:"type"`
}

func (r ApplicationV1ProjectNewParamsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1ProjectNewParamsEnvironmentVariablesTarget string

const (
	ApplicationV1ProjectNewParamsEnvironmentVariablesTargetProduction  ApplicationV1ProjectNewParamsEnvironmentVariablesTarget = "production"
	ApplicationV1ProjectNewParamsEnvironmentVariablesTargetPreview     ApplicationV1ProjectNewParamsEnvironmentVariablesTarget = "preview"
	ApplicationV1ProjectNewParamsEnvironmentVariablesTargetDevelopment ApplicationV1ProjectNewParamsEnvironmentVariablesTarget = "development"
)

func (r ApplicationV1ProjectNewParamsEnvironmentVariablesTarget) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectNewParamsEnvironmentVariablesTargetProduction, ApplicationV1ProjectNewParamsEnvironmentVariablesTargetPreview, ApplicationV1ProjectNewParamsEnvironmentVariablesTargetDevelopment:
		return true
	}
	return false
}

// Variable type
type ApplicationV1ProjectNewParamsEnvironmentVariablesType string

const (
	ApplicationV1ProjectNewParamsEnvironmentVariablesTypePlain     ApplicationV1ProjectNewParamsEnvironmentVariablesType = "plain"
	ApplicationV1ProjectNewParamsEnvironmentVariablesTypeEncrypted ApplicationV1ProjectNewParamsEnvironmentVariablesType = "encrypted"
	ApplicationV1ProjectNewParamsEnvironmentVariablesTypeSecret    ApplicationV1ProjectNewParamsEnvironmentVariablesType = "secret"
)

func (r ApplicationV1ProjectNewParamsEnvironmentVariablesType) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectNewParamsEnvironmentVariablesTypePlain, ApplicationV1ProjectNewParamsEnvironmentVariablesTypeEncrypted, ApplicationV1ProjectNewParamsEnvironmentVariablesTypeSecret:
		return true
	}
	return false
}

type ApplicationV1ProjectDeleteParams struct {
	// Also delete the project from hosting (default: true)
	DeleteFromHosting param.Field[bool] `query:"deleteFromHosting"`
}

// URLQuery serializes [ApplicationV1ProjectDeleteParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1ProjectDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ApplicationV1ProjectNewDeploymentParams struct {
	// Additional environment variables to set or update before deployment
	EnvironmentVariables param.Field[[]ApplicationV1ProjectNewDeploymentParamsEnvironmentVariable] `json:"environmentVariables"`
}

func (r ApplicationV1ProjectNewDeploymentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1ProjectNewDeploymentParamsEnvironmentVariable struct {
	// Environment variable name
	Key param.Field[string] `json:"key" api:"required"`
	// Deployment targets for this variable
	Target param.Field[[]ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget] `json:"target" api:"required"`
	// Environment variable value
	Value param.Field[string] `json:"value" api:"required"`
	// Variable type
	Type param.Field[ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesType] `json:"type"`
}

func (r ApplicationV1ProjectNewDeploymentParamsEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget string

const (
	ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetProduction  ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget = "production"
	ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetPreview     ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget = "preview"
	ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetDevelopment ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget = "development"
)

func (r ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTarget) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetProduction, ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetPreview, ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTargetDevelopment:
		return true
	}
	return false
}

// Variable type
type ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesType string

const (
	ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypePlain     ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesType = "plain"
	ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypeEncrypted ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesType = "encrypted"
	ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypeSecret    ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesType = "secret"
)

func (r ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesType) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypePlain, ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypeEncrypted, ApplicationV1ProjectNewDeploymentParamsEnvironmentVariablesTypeSecret:
		return true
	}
	return false
}

type ApplicationV1ProjectNewDomainParams struct {
	// Domain name to add
	Domain param.Field[string] `json:"domain" api:"required"`
	// Git branch to associate with this domain
	GitBranch param.Field[string] `json:"gitBranch"`
}

func (r ApplicationV1ProjectNewDomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1ProjectNewEnvParams struct {
	// Environment variable name
	Key param.Field[string] `json:"key" api:"required"`
	// Deployment targets for this variable
	Target param.Field[[]ApplicationV1ProjectNewEnvParamsTarget] `json:"target" api:"required"`
	// Environment variable value
	Value param.Field[string] `json:"value" api:"required"`
	// Specific git branch (for preview deployments)
	GitBranch param.Field[string] `json:"gitBranch"`
	// Variable type
	Type param.Field[ApplicationV1ProjectNewEnvParamsType] `json:"type"`
}

func (r ApplicationV1ProjectNewEnvParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationV1ProjectNewEnvParamsTarget string

const (
	ApplicationV1ProjectNewEnvParamsTargetProduction  ApplicationV1ProjectNewEnvParamsTarget = "production"
	ApplicationV1ProjectNewEnvParamsTargetPreview     ApplicationV1ProjectNewEnvParamsTarget = "preview"
	ApplicationV1ProjectNewEnvParamsTargetDevelopment ApplicationV1ProjectNewEnvParamsTarget = "development"
)

func (r ApplicationV1ProjectNewEnvParamsTarget) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectNewEnvParamsTargetProduction, ApplicationV1ProjectNewEnvParamsTargetPreview, ApplicationV1ProjectNewEnvParamsTargetDevelopment:
		return true
	}
	return false
}

// Variable type
type ApplicationV1ProjectNewEnvParamsType string

const (
	ApplicationV1ProjectNewEnvParamsTypePlain     ApplicationV1ProjectNewEnvParamsType = "plain"
	ApplicationV1ProjectNewEnvParamsTypeEncrypted ApplicationV1ProjectNewEnvParamsType = "encrypted"
	ApplicationV1ProjectNewEnvParamsTypeSecret    ApplicationV1ProjectNewEnvParamsType = "secret"
)

func (r ApplicationV1ProjectNewEnvParamsType) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectNewEnvParamsTypePlain, ApplicationV1ProjectNewEnvParamsTypeEncrypted, ApplicationV1ProjectNewEnvParamsTypeSecret:
		return true
	}
	return false
}

type ApplicationV1ProjectGetRuntimeLogsParams struct {
	// Maximum number of logs to return
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [ApplicationV1ProjectGetRuntimeLogsParams]'s query
// parameters as `url.Values`.
func (r ApplicationV1ProjectGetRuntimeLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ApplicationV1ProjectListDeploymentsParams struct {
	// Maximum number of deployments to return
	Limit param.Field[float64] `query:"limit"`
	// Filter by deployment state
	State param.Field[string] `query:"state"`
	// Filter by deployment target
	Target param.Field[ApplicationV1ProjectListDeploymentsParamsTarget] `query:"target"`
}

// URLQuery serializes [ApplicationV1ProjectListDeploymentsParams]'s query
// parameters as `url.Values`.
func (r ApplicationV1ProjectListDeploymentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by deployment target
type ApplicationV1ProjectListDeploymentsParamsTarget string

const (
	ApplicationV1ProjectListDeploymentsParamsTargetProduction ApplicationV1ProjectListDeploymentsParamsTarget = "production"
	ApplicationV1ProjectListDeploymentsParamsTargetStaging    ApplicationV1ProjectListDeploymentsParamsTarget = "staging"
)

func (r ApplicationV1ProjectListDeploymentsParamsTarget) IsKnown() bool {
	switch r {
	case ApplicationV1ProjectListDeploymentsParamsTargetProduction, ApplicationV1ProjectListDeploymentsParamsTargetStaging:
		return true
	}
	return false
}

type ApplicationV1ProjectListEnvParams struct {
	// Whether to decrypt and return values (requires appropriate permissions)
	Decrypt param.Field[bool] `query:"decrypt"`
}

// URLQuery serializes [ApplicationV1ProjectListEnvParams]'s query parameters as
// `url.Values`.
func (r ApplicationV1ProjectListEnvParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
