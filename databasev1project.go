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

// DatabaseV1ProjectService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabaseV1ProjectService] method instead.
type DatabaseV1ProjectService struct {
	Options []option.RequestOption
}

// NewDatabaseV1ProjectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDatabaseV1ProjectService(opts ...option.RequestOption) (r *DatabaseV1ProjectService) {
	r = &DatabaseV1ProjectService{}
	r.Options = opts
	return
}

// Creates a new serverless Postgres database project powered by Neon. Includes
// automatic scaling, connection pooling, and a default 'main' branch with 'neondb'
// database. Supports branching for isolated dev/staging environments. Perfect for
// case management applications, document workflows, and litigation support
// systems.
func (r *DatabaseV1ProjectService) New(ctx context.Context, body DatabaseV1ProjectNewParams, opts ...option.RequestOption) (res *DatabaseV1ProjectNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "database/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information about a specific database project including
// branches, databases, storage/compute metrics, connection host, and linked
// deployments. Fetches live usage statistics from Neon API.
func (r *DatabaseV1ProjectService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DatabaseV1ProjectGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("database/v1/projects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all serverless Postgres database projects for the authenticated
// organization. Includes storage and compute metrics, plus linked deployments from
// Thurgood apps and Compute instances.
func (r *DatabaseV1ProjectService) List(ctx context.Context, opts ...option.RequestOption) (res *DatabaseV1ProjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "database/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a database project from Neon and marks it as deleted in
// Case.dev. This action cannot be undone and will destroy all data including
// branches and databases. Use with caution.
func (r *DatabaseV1ProjectService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DatabaseV1ProjectDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("database/v1/projects/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new branch from the specified parent branch (or default 'main'
// branch). Branches provide instant point-in-time clones of your database for
// isolated development, staging, testing, or feature work. Perfect for testing
// schema changes, running migrations safely, or creating ephemeral preview
// environments.
func (r *DatabaseV1ProjectService) NewBranch(ctx context.Context, id string, body DatabaseV1ProjectNewBranchParams, opts ...option.RequestOption) (res *DatabaseV1ProjectNewBranchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("database/v1/projects/%s/branches", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the PostgreSQL connection URI for a database project. Supports
// selecting specific branches and pooled vs direct connections. Connection strings
// include credentials and should be stored securely. Use for configuring
// applications and deployment environments.
func (r *DatabaseV1ProjectService) GetConnection(ctx context.Context, id string, query DatabaseV1ProjectGetConnectionParams, opts ...option.RequestOption) (res *DatabaseV1ProjectGetConnectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("database/v1/projects/%s/connection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all branches for a database project. Branches enable isolated
// development and testing environments with instant point-in-time cloning. Each
// branch includes the default branch and any custom branches created for staging,
// testing, or feature development.
func (r *DatabaseV1ProjectService) ListBranches(ctx context.Context, id string, opts ...option.RequestOption) (res *DatabaseV1ProjectListBranchesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("database/v1/projects/%s/branches", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DatabaseV1ProjectNewResponse struct {
	// Project ID
	ID string `json:"id" api:"required"`
	// Project creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Default 'main' branch details
	DefaultBranch DatabaseV1ProjectNewResponseDefaultBranch `json:"defaultBranch" api:"required"`
	// Project name
	Name string `json:"name" api:"required"`
	// PostgreSQL major version
	PgVersion int64 `json:"pgVersion" api:"required"`
	// AWS region
	Region string `json:"region" api:"required"`
	// Project status
	Status DatabaseV1ProjectNewResponseStatus `json:"status" api:"required"`
	// Project description
	Description string                           `json:"description" api:"nullable"`
	JSON        databaseV1ProjectNewResponseJSON `json:"-"`
}

// databaseV1ProjectNewResponseJSON contains the JSON metadata for the struct
// [DatabaseV1ProjectNewResponse]
type databaseV1ProjectNewResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	DefaultBranch apijson.Field
	Name          apijson.Field
	PgVersion     apijson.Field
	Region        apijson.Field
	Status        apijson.Field
	Description   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DatabaseV1ProjectNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectNewResponseJSON) RawJSON() string {
	return r.raw
}

// Default 'main' branch details
type DatabaseV1ProjectNewResponseDefaultBranch struct {
	// Branch ID
	ID string `json:"id"`
	// Branch name
	Name string                                        `json:"name"`
	JSON databaseV1ProjectNewResponseDefaultBranchJSON `json:"-"`
}

// databaseV1ProjectNewResponseDefaultBranchJSON contains the JSON metadata for the
// struct [DatabaseV1ProjectNewResponseDefaultBranch]
type databaseV1ProjectNewResponseDefaultBranchJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectNewResponseDefaultBranch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectNewResponseDefaultBranchJSON) RawJSON() string {
	return r.raw
}

// Project status
type DatabaseV1ProjectNewResponseStatus string

const (
	DatabaseV1ProjectNewResponseStatusActive    DatabaseV1ProjectNewResponseStatus = "active"
	DatabaseV1ProjectNewResponseStatusSuspended DatabaseV1ProjectNewResponseStatus = "suspended"
	DatabaseV1ProjectNewResponseStatusDeleted   DatabaseV1ProjectNewResponseStatus = "deleted"
)

func (r DatabaseV1ProjectNewResponseStatus) IsKnown() bool {
	switch r {
	case DatabaseV1ProjectNewResponseStatusActive, DatabaseV1ProjectNewResponseStatusSuspended, DatabaseV1ProjectNewResponseStatusDeleted:
		return true
	}
	return false
}

type DatabaseV1ProjectGetResponse struct {
	// Project ID
	ID string `json:"id" api:"required"`
	// All branches in this project
	Branches []DatabaseV1ProjectGetResponseBranch `json:"branches" api:"required"`
	// Total compute time consumed in seconds
	ComputeTimeSeconds float64 `json:"computeTimeSeconds" api:"required"`
	// Database connection hostname (masked for security)
	ConnectionHost string `json:"connectionHost" api:"required"`
	// Project creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Databases in the default branch
	Databases []DatabaseV1ProjectGetResponseDatabase `json:"databases" api:"required"`
	// Linked deployments using this database
	LinkedDeployments []DatabaseV1ProjectGetResponseLinkedDeployment `json:"linkedDeployments" api:"required"`
	// Project name
	Name string `json:"name" api:"required"`
	// PostgreSQL major version
	PgVersion int64 `json:"pgVersion" api:"required"`
	// AWS region
	Region string `json:"region" api:"required"`
	// Project status
	Status DatabaseV1ProjectGetResponseStatus `json:"status" api:"required"`
	// Current storage usage in bytes
	StorageSizeBytes float64 `json:"storageSizeBytes" api:"required"`
	// Project last update timestamp
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Project description
	Description string                           `json:"description" api:"nullable"`
	JSON        databaseV1ProjectGetResponseJSON `json:"-"`
}

// databaseV1ProjectGetResponseJSON contains the JSON metadata for the struct
// [DatabaseV1ProjectGetResponse]
type databaseV1ProjectGetResponseJSON struct {
	ID                 apijson.Field
	Branches           apijson.Field
	ComputeTimeSeconds apijson.Field
	ConnectionHost     apijson.Field
	CreatedAt          apijson.Field
	Databases          apijson.Field
	LinkedDeployments  apijson.Field
	Name               apijson.Field
	PgVersion          apijson.Field
	Region             apijson.Field
	Status             apijson.Field
	StorageSizeBytes   apijson.Field
	UpdatedAt          apijson.Field
	Description        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DatabaseV1ProjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectGetResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectGetResponseBranch struct {
	// Branch ID
	ID string `json:"id"`
	// Branch creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Whether this is the default branch
	IsDefault bool `json:"isDefault"`
	// Branch name
	Name string `json:"name"`
	// Branch status
	Status string                                 `json:"status"`
	JSON   databaseV1ProjectGetResponseBranchJSON `json:"-"`
}

// databaseV1ProjectGetResponseBranchJSON contains the JSON metadata for the struct
// [DatabaseV1ProjectGetResponseBranch]
type databaseV1ProjectGetResponseBranchJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectGetResponseBranch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectGetResponseBranchJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectGetResponseDatabase struct {
	// Database ID
	ID string `json:"id"`
	// Database name
	Name string `json:"name"`
	// Database owner role name
	OwnerName string                                   `json:"ownerName"`
	JSON      databaseV1ProjectGetResponseDatabaseJSON `json:"-"`
}

// databaseV1ProjectGetResponseDatabaseJSON contains the JSON metadata for the
// struct [DatabaseV1ProjectGetResponseDatabase]
type databaseV1ProjectGetResponseDatabaseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	OwnerName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectGetResponseDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectGetResponseDatabaseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectGetResponseLinkedDeployment struct {
	// Deployment ID
	ID string `json:"id"`
	// Environment variable name for connection string
	EnvVarName string `json:"envVarName"`
	// Deployment name
	Name string `json:"name"`
	// Deployment type
	Type DatabaseV1ProjectGetResponseLinkedDeploymentsType `json:"type"`
	// Deployment URL
	URL  string                                           `json:"url"`
	JSON databaseV1ProjectGetResponseLinkedDeploymentJSON `json:"-"`
}

// databaseV1ProjectGetResponseLinkedDeploymentJSON contains the JSON metadata for
// the struct [DatabaseV1ProjectGetResponseLinkedDeployment]
type databaseV1ProjectGetResponseLinkedDeploymentJSON struct {
	ID          apijson.Field
	EnvVarName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectGetResponseLinkedDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectGetResponseLinkedDeploymentJSON) RawJSON() string {
	return r.raw
}

// Deployment type
type DatabaseV1ProjectGetResponseLinkedDeploymentsType string

const (
	DatabaseV1ProjectGetResponseLinkedDeploymentsTypeThurgood DatabaseV1ProjectGetResponseLinkedDeploymentsType = "thurgood"
	DatabaseV1ProjectGetResponseLinkedDeploymentsTypeCompute  DatabaseV1ProjectGetResponseLinkedDeploymentsType = "compute"
)

func (r DatabaseV1ProjectGetResponseLinkedDeploymentsType) IsKnown() bool {
	switch r {
	case DatabaseV1ProjectGetResponseLinkedDeploymentsTypeThurgood, DatabaseV1ProjectGetResponseLinkedDeploymentsTypeCompute:
		return true
	}
	return false
}

// Project status
type DatabaseV1ProjectGetResponseStatus string

const (
	DatabaseV1ProjectGetResponseStatusActive    DatabaseV1ProjectGetResponseStatus = "active"
	DatabaseV1ProjectGetResponseStatusSuspended DatabaseV1ProjectGetResponseStatus = "suspended"
	DatabaseV1ProjectGetResponseStatusDeleted   DatabaseV1ProjectGetResponseStatus = "deleted"
)

func (r DatabaseV1ProjectGetResponseStatus) IsKnown() bool {
	switch r {
	case DatabaseV1ProjectGetResponseStatusActive, DatabaseV1ProjectGetResponseStatusSuspended, DatabaseV1ProjectGetResponseStatusDeleted:
		return true
	}
	return false
}

type DatabaseV1ProjectListResponse struct {
	Projects []DatabaseV1ProjectListResponseProject `json:"projects" api:"required"`
	JSON     databaseV1ProjectListResponseJSON      `json:"-"`
}

// databaseV1ProjectListResponseJSON contains the JSON metadata for the struct
// [DatabaseV1ProjectListResponse]
type databaseV1ProjectListResponseJSON struct {
	Projects    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectListResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectListResponseProject struct {
	// Project ID
	ID string `json:"id"`
	// Total compute time consumed in seconds
	ComputeTimeSeconds float64 `json:"computeTimeSeconds"`
	// Project creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Project description
	Description string `json:"description" api:"nullable"`
	// Linked application deployments using this database
	LinkedDeployments []DatabaseV1ProjectListResponseProjectsLinkedDeployment `json:"linkedDeployments"`
	// Project name
	Name string `json:"name"`
	// PostgreSQL major version
	PgVersion int64 `json:"pgVersion"`
	// AWS region where database is deployed
	Region string `json:"region"`
	// Current project status
	Status DatabaseV1ProjectListResponseProjectsStatus `json:"status"`
	// Current storage usage in bytes
	StorageSizeBytes float64 `json:"storageSizeBytes"`
	// Project last update timestamp
	UpdatedAt time.Time                                `json:"updatedAt" format:"date-time"`
	JSON      databaseV1ProjectListResponseProjectJSON `json:"-"`
}

// databaseV1ProjectListResponseProjectJSON contains the JSON metadata for the
// struct [DatabaseV1ProjectListResponseProject]
type databaseV1ProjectListResponseProjectJSON struct {
	ID                 apijson.Field
	ComputeTimeSeconds apijson.Field
	CreatedAt          apijson.Field
	Description        apijson.Field
	LinkedDeployments  apijson.Field
	Name               apijson.Field
	PgVersion          apijson.Field
	Region             apijson.Field
	Status             apijson.Field
	StorageSizeBytes   apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DatabaseV1ProjectListResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectListResponseProjectJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectListResponseProjectsLinkedDeployment struct {
	// Deployment ID
	ID string `json:"id"`
	// Deployment name
	Name string `json:"name"`
	// Type of deployment
	Type DatabaseV1ProjectListResponseProjectsLinkedDeploymentsType `json:"type"`
	// Deployment URL (for Thurgood apps)
	URL  string                                                    `json:"url"`
	JSON databaseV1ProjectListResponseProjectsLinkedDeploymentJSON `json:"-"`
}

// databaseV1ProjectListResponseProjectsLinkedDeploymentJSON contains the JSON
// metadata for the struct [DatabaseV1ProjectListResponseProjectsLinkedDeployment]
type databaseV1ProjectListResponseProjectsLinkedDeploymentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectListResponseProjectsLinkedDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectListResponseProjectsLinkedDeploymentJSON) RawJSON() string {
	return r.raw
}

// Type of deployment
type DatabaseV1ProjectListResponseProjectsLinkedDeploymentsType string

const (
	DatabaseV1ProjectListResponseProjectsLinkedDeploymentsTypeThurgood DatabaseV1ProjectListResponseProjectsLinkedDeploymentsType = "thurgood"
	DatabaseV1ProjectListResponseProjectsLinkedDeploymentsTypeCompute  DatabaseV1ProjectListResponseProjectsLinkedDeploymentsType = "compute"
)

func (r DatabaseV1ProjectListResponseProjectsLinkedDeploymentsType) IsKnown() bool {
	switch r {
	case DatabaseV1ProjectListResponseProjectsLinkedDeploymentsTypeThurgood, DatabaseV1ProjectListResponseProjectsLinkedDeploymentsTypeCompute:
		return true
	}
	return false
}

// Current project status
type DatabaseV1ProjectListResponseProjectsStatus string

const (
	DatabaseV1ProjectListResponseProjectsStatusActive    DatabaseV1ProjectListResponseProjectsStatus = "active"
	DatabaseV1ProjectListResponseProjectsStatusSuspended DatabaseV1ProjectListResponseProjectsStatus = "suspended"
	DatabaseV1ProjectListResponseProjectsStatusDeleted   DatabaseV1ProjectListResponseProjectsStatus = "deleted"
)

func (r DatabaseV1ProjectListResponseProjectsStatus) IsKnown() bool {
	switch r {
	case DatabaseV1ProjectListResponseProjectsStatusActive, DatabaseV1ProjectListResponseProjectsStatusSuspended, DatabaseV1ProjectListResponseProjectsStatusDeleted:
		return true
	}
	return false
}

type DatabaseV1ProjectDeleteResponse struct {
	// Confirmation message
	Message string `json:"message" api:"required"`
	// Deletion success indicator
	Success bool                                `json:"success" api:"required"`
	JSON    databaseV1ProjectDeleteResponseJSON `json:"-"`
}

// databaseV1ProjectDeleteResponseJSON contains the JSON metadata for the struct
// [DatabaseV1ProjectDeleteResponse]
type databaseV1ProjectDeleteResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectNewBranchResponse struct {
	// Branch ID
	ID string `json:"id" api:"required"`
	// Branch creation timestamp
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether this is the default branch (always false for new branches)
	IsDefault bool `json:"isDefault" api:"required"`
	// Branch name
	Name string `json:"name" api:"required"`
	// Parent branch ID
	ParentBranchID string `json:"parentBranchId" api:"required,nullable"`
	// Branch status
	Status string                                 `json:"status" api:"required"`
	JSON   databaseV1ProjectNewBranchResponseJSON `json:"-"`
}

// databaseV1ProjectNewBranchResponseJSON contains the JSON metadata for the struct
// [DatabaseV1ProjectNewBranchResponse]
type databaseV1ProjectNewBranchResponseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	IsDefault      apijson.Field
	Name           apijson.Field
	ParentBranchID apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DatabaseV1ProjectNewBranchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectNewBranchResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectGetConnectionResponse struct {
	// Branch name for this connection
	Branch string `json:"branch" api:"required"`
	// PostgreSQL connection string (includes credentials)
	ConnectionUri string `json:"connectionUri" api:"required" format:"uri"`
	// Whether this is a pooled connection
	Pooled bool                                       `json:"pooled" api:"required"`
	JSON   databaseV1ProjectGetConnectionResponseJSON `json:"-"`
}

// databaseV1ProjectGetConnectionResponseJSON contains the JSON metadata for the
// struct [DatabaseV1ProjectGetConnectionResponse]
type databaseV1ProjectGetConnectionResponseJSON struct {
	Branch        apijson.Field
	ConnectionUri apijson.Field
	Pooled        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DatabaseV1ProjectGetConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectGetConnectionResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectListBranchesResponse struct {
	Branches []DatabaseV1ProjectListBranchesResponseBranch `json:"branches" api:"required"`
	JSON     databaseV1ProjectListBranchesResponseJSON     `json:"-"`
}

// databaseV1ProjectListBranchesResponseJSON contains the JSON metadata for the
// struct [DatabaseV1ProjectListBranchesResponse]
type databaseV1ProjectListBranchesResponseJSON struct {
	Branches    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1ProjectListBranchesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectListBranchesResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectListBranchesResponseBranch struct {
	// Branch ID
	ID string `json:"id"`
	// Branch creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Whether this is the default branch
	IsDefault bool `json:"isDefault"`
	// Branch name
	Name string `json:"name"`
	// Parent branch ID (null for default branch)
	ParentBranchID string `json:"parentBranchId" api:"nullable"`
	// Branch status
	Status string `json:"status"`
	// Branch last update timestamp
	UpdatedAt time.Time                                       `json:"updatedAt" format:"date-time"`
	JSON      databaseV1ProjectListBranchesResponseBranchJSON `json:"-"`
}

// databaseV1ProjectListBranchesResponseBranchJSON contains the JSON metadata for
// the struct [DatabaseV1ProjectListBranchesResponseBranch]
type databaseV1ProjectListBranchesResponseBranchJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	IsDefault      apijson.Field
	Name           apijson.Field
	ParentBranchID apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DatabaseV1ProjectListBranchesResponseBranch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1ProjectListBranchesResponseBranchJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1ProjectNewParams struct {
	// Project name (letters, numbers, hyphens, underscores only)
	Name param.Field[string] `json:"name" api:"required"`
	// Optional project description
	Description param.Field[string] `json:"description"`
	// AWS region for database deployment
	Region param.Field[DatabaseV1ProjectNewParamsRegion] `json:"region"`
}

func (r DatabaseV1ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AWS region for database deployment
type DatabaseV1ProjectNewParamsRegion string

const (
	DatabaseV1ProjectNewParamsRegionAwsUsEast1      DatabaseV1ProjectNewParamsRegion = "aws-us-east-1"
	DatabaseV1ProjectNewParamsRegionAwsUsEast2      DatabaseV1ProjectNewParamsRegion = "aws-us-east-2"
	DatabaseV1ProjectNewParamsRegionAwsUsWest2      DatabaseV1ProjectNewParamsRegion = "aws-us-west-2"
	DatabaseV1ProjectNewParamsRegionAwsEuCentral1   DatabaseV1ProjectNewParamsRegion = "aws-eu-central-1"
	DatabaseV1ProjectNewParamsRegionAwsEuWest1      DatabaseV1ProjectNewParamsRegion = "aws-eu-west-1"
	DatabaseV1ProjectNewParamsRegionAwsEuWest2      DatabaseV1ProjectNewParamsRegion = "aws-eu-west-2"
	DatabaseV1ProjectNewParamsRegionAwsApSoutheast1 DatabaseV1ProjectNewParamsRegion = "aws-ap-southeast-1"
	DatabaseV1ProjectNewParamsRegionAwsApSoutheast2 DatabaseV1ProjectNewParamsRegion = "aws-ap-southeast-2"
)

func (r DatabaseV1ProjectNewParamsRegion) IsKnown() bool {
	switch r {
	case DatabaseV1ProjectNewParamsRegionAwsUsEast1, DatabaseV1ProjectNewParamsRegionAwsUsEast2, DatabaseV1ProjectNewParamsRegionAwsUsWest2, DatabaseV1ProjectNewParamsRegionAwsEuCentral1, DatabaseV1ProjectNewParamsRegionAwsEuWest1, DatabaseV1ProjectNewParamsRegionAwsEuWest2, DatabaseV1ProjectNewParamsRegionAwsApSoutheast1, DatabaseV1ProjectNewParamsRegionAwsApSoutheast2:
		return true
	}
	return false
}

type DatabaseV1ProjectNewBranchParams struct {
	// Branch name (letters, numbers, hyphens, underscores only)
	Name param.Field[string] `json:"name" api:"required"`
	// Parent branch ID to clone from (defaults to main branch)
	ParentBranchID param.Field[string] `json:"parentBranchId"`
}

func (r DatabaseV1ProjectNewBranchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseV1ProjectGetConnectionParams struct {
	// Branch name (defaults to 'main')
	Branch param.Field[string] `query:"branch"`
	// Use pooled connection (PgBouncer)
	Pooled param.Field[bool] `query:"pooled"`
}

// URLQuery serializes [DatabaseV1ProjectGetConnectionParams]'s query parameters as
// `url.Values`.
func (r DatabaseV1ProjectGetConnectionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
