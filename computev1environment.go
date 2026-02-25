// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// ComputeV1EnvironmentService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeV1EnvironmentService] method instead.
type ComputeV1EnvironmentService struct {
	Options []option.RequestOption
}

// NewComputeV1EnvironmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeV1EnvironmentService(opts ...option.RequestOption) (r *ComputeV1EnvironmentService) {
	r = &ComputeV1EnvironmentService{}
	r.Options = opts
	return
}

// Creates a new compute environment for running serverless workloads. Each
// environment gets its own isolated namespace with a unique domain for hosting
// applications and APIs. The first environment created becomes the default
// environment for the organization.
func (r *ComputeV1EnvironmentService) New(ctx context.Context, body ComputeV1EnvironmentNewParams, opts ...option.RequestOption) (res *ComputeV1EnvironmentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/environments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific compute environment by name. Returns environment
// configuration including status, domain, and metadata for your serverless compute
// infrastructure.
func (r *ComputeV1EnvironmentService) Get(ctx context.Context, name string, opts ...option.RequestOption) (res *ComputeV1EnvironmentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/environments/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all compute environments for your organization. Environments provide
// isolated execution contexts for running code and workflows.
func (r *ComputeV1EnvironmentService) List(ctx context.Context, opts ...option.RequestOption) (res *ComputeV1EnvironmentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/environments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently delete a compute environment and all its associated resources. This
// will stop all running deployments and clean up related configurations. The
// default environment cannot be deleted if other environments exist.
func (r *ComputeV1EnvironmentService) Delete(ctx context.Context, name string, opts ...option.RequestOption) (res *ComputeV1EnvironmentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/environments/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Sets a compute environment as the default for the organization. Only one
// environment can be default at a time - setting a new default will automatically
// unset the previous one.
func (r *ComputeV1EnvironmentService) SetDefault(ctx context.Context, name string, opts ...option.RequestOption) (res *ComputeV1EnvironmentSetDefaultResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/environments/%s/default", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ComputeV1EnvironmentNewResponse struct {
	// Unique environment identifier
	ID string `json:"id"`
	// Environment creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Unique domain for this environment
	Domain string `json:"domain"`
	// Whether this is the default environment
	IsDefault bool `json:"isDefault"`
	// Environment name
	Name string `json:"name"`
	// URL-friendly slug derived from name
	Slug string `json:"slug"`
	// Environment status
	Status ComputeV1EnvironmentNewResponseStatus `json:"status"`
	JSON   computeV1EnvironmentNewResponseJSON   `json:"-"`
}

// computeV1EnvironmentNewResponseJSON contains the JSON metadata for the struct
// [ComputeV1EnvironmentNewResponse]
type computeV1EnvironmentNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Domain      apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1EnvironmentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1EnvironmentNewResponseJSON) RawJSON() string {
	return r.raw
}

// Environment status
type ComputeV1EnvironmentNewResponseStatus string

const (
	ComputeV1EnvironmentNewResponseStatusActive   ComputeV1EnvironmentNewResponseStatus = "active"
	ComputeV1EnvironmentNewResponseStatusInactive ComputeV1EnvironmentNewResponseStatus = "inactive"
)

func (r ComputeV1EnvironmentNewResponseStatus) IsKnown() bool {
	switch r {
	case ComputeV1EnvironmentNewResponseStatusActive, ComputeV1EnvironmentNewResponseStatusInactive:
		return true
	}
	return false
}

type ComputeV1EnvironmentGetResponse struct {
	// Unique environment identifier
	ID string `json:"id"`
	// Environment creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Environment domain URL
	Domain string `json:"domain"`
	// Whether this is the default environment
	IsDefault bool `json:"isDefault"`
	// Environment name
	Name string `json:"name"`
	// URL-safe environment slug
	Slug string `json:"slug"`
	// Environment status (active, inactive, etc.)
	Status string `json:"status"`
	// Environment last update timestamp
	UpdatedAt time.Time                           `json:"updatedAt" format:"date-time"`
	JSON      computeV1EnvironmentGetResponseJSON `json:"-"`
}

// computeV1EnvironmentGetResponseJSON contains the JSON metadata for the struct
// [ComputeV1EnvironmentGetResponse]
type computeV1EnvironmentGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Domain      apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1EnvironmentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1EnvironmentGetResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1EnvironmentListResponse struct {
	Environments []ComputeV1EnvironmentListResponseEnvironment `json:"environments"`
	JSON         computeV1EnvironmentListResponseJSON          `json:"-"`
}

// computeV1EnvironmentListResponseJSON contains the JSON metadata for the struct
// [ComputeV1EnvironmentListResponse]
type computeV1EnvironmentListResponseJSON struct {
	Environments apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ComputeV1EnvironmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1EnvironmentListResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1EnvironmentListResponseEnvironment struct {
	// Unique environment identifier
	ID string `json:"id"`
	// Environment creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Environment domain
	Domain string `json:"domain"`
	// Whether this is the default environment
	IsDefault bool `json:"isDefault"`
	// Human-readable environment name
	Name string `json:"name"`
	// URL-safe environment identifier
	Slug string `json:"slug"`
	// Environment status
	Status string `json:"status"`
	// Last update timestamp
	UpdatedAt time.Time                                       `json:"updatedAt" format:"date-time"`
	JSON      computeV1EnvironmentListResponseEnvironmentJSON `json:"-"`
}

// computeV1EnvironmentListResponseEnvironmentJSON contains the JSON metadata for
// the struct [ComputeV1EnvironmentListResponseEnvironment]
type computeV1EnvironmentListResponseEnvironmentJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Domain      apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1EnvironmentListResponseEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1EnvironmentListResponseEnvironmentJSON) RawJSON() string {
	return r.raw
}

type ComputeV1EnvironmentDeleteResponse struct {
	Message string                                 `json:"message" api:"required"`
	Success bool                                   `json:"success" api:"required"`
	JSON    computeV1EnvironmentDeleteResponseJSON `json:"-"`
}

// computeV1EnvironmentDeleteResponseJSON contains the JSON metadata for the struct
// [ComputeV1EnvironmentDeleteResponse]
type computeV1EnvironmentDeleteResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1EnvironmentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1EnvironmentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1EnvironmentSetDefaultResponse struct {
	// Unique environment identifier
	ID string `json:"id"`
	// Environment creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Environment domain
	Domain string `json:"domain"`
	// Whether this is the default environment
	IsDefault bool `json:"isDefault"`
	// Environment name
	Name string `json:"name"`
	// URL-friendly environment identifier
	Slug string `json:"slug"`
	// Current environment status
	Status string `json:"status"`
	// Last update timestamp
	UpdatedAt time.Time                                  `json:"updatedAt" format:"date-time"`
	JSON      computeV1EnvironmentSetDefaultResponseJSON `json:"-"`
}

// computeV1EnvironmentSetDefaultResponseJSON contains the JSON metadata for the
// struct [ComputeV1EnvironmentSetDefaultResponse]
type computeV1EnvironmentSetDefaultResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Domain      apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1EnvironmentSetDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1EnvironmentSetDefaultResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1EnvironmentNewParams struct {
	// Environment name (alphanumeric, hyphens, and underscores only)
	Name param.Field[string] `json:"name" api:"required"`
}

func (r ComputeV1EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
