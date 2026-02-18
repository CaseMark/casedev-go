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

// ComputeV1SecretService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeV1SecretService] method instead.
type ComputeV1SecretService struct {
	Options []option.RequestOption
}

// NewComputeV1SecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeV1SecretService(opts ...option.RequestOption) (r *ComputeV1SecretService) {
	r = &ComputeV1SecretService{}
	r.Options = opts
	return
}

// Creates a new secret group in a compute environment. Secret groups organize
// related secrets for use in serverless functions and workflows. If no environment
// is specified, the group is created in the default environment.
//
// **Features:**
//
//   - Organize secrets by logical groups (e.g., database, APIs, third-party
//     services)
//   - Environment-based isolation
//   - Validation of group names
//   - Conflict detection for existing groups
func (r *ComputeV1SecretService) New(ctx context.Context, body ComputeV1SecretNewParams, opts ...option.RequestOption) (res *ComputeV1SecretNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all secret groups for a compute environment. Secret groups organize
// related secrets (API keys, credentials, etc.) that can be securely accessed by
// compute jobs during execution.
func (r *ComputeV1SecretService) List(ctx context.Context, query ComputeV1SecretListParams, opts ...option.RequestOption) (res *ComputeV1SecretListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an entire secret group or a specific key within a secret group. When
// deleting a specific key, the remaining secrets in the group are preserved. When
// deleting the entire group, all secrets and the group itself are removed.
func (r *ComputeV1SecretService) DeleteGroup(ctx context.Context, group string, body ComputeV1SecretDeleteGroupParams, opts ...option.RequestOption) (res *ComputeV1SecretDeleteGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/secrets/%s", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retrieve the keys (names) of secrets in a specified group within a compute
// environment. For security reasons, actual secret values are not returned - only
// the keys and metadata.
func (r *ComputeV1SecretService) GetGroup(ctx context.Context, group string, query ComputeV1SecretGetGroupParams, opts ...option.RequestOption) (res *ComputeV1SecretGetGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/secrets/%s", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Set or update secrets in a compute secret group. Secrets are encrypted with
// AES-256-GCM. Use this to manage environment variables and API keys for your
// compute workloads.
func (r *ComputeV1SecretService) UpdateGroup(ctx context.Context, group string, body ComputeV1SecretUpdateGroupParams, opts ...option.RequestOption) (res *ComputeV1SecretUpdateGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if group == "" {
		err = errors.New("missing required group parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/secrets/%s", group)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ComputeV1SecretNewResponse struct {
	// Unique identifier for the secret group
	ID string `json:"id"`
	// Creation timestamp
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Description of the secret group
	Description string `json:"description"`
	// Name of the secret group
	Name string                         `json:"name"`
	JSON computeV1SecretNewResponseJSON `json:"-"`
}

// computeV1SecretNewResponseJSON contains the JSON metadata for the struct
// [ComputeV1SecretNewResponse]
type computeV1SecretNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretNewResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretListResponse struct {
	Groups []ComputeV1SecretListResponseGroup `json:"groups"`
	JSON   computeV1SecretListResponseJSON    `json:"-"`
}

// computeV1SecretListResponseJSON contains the JSON metadata for the struct
// [ComputeV1SecretListResponse]
type computeV1SecretListResponseJSON struct {
	Groups      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretListResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretListResponseGroup struct {
	// Unique identifier for the secret group
	ID string `json:"id"`
	// When the secret group was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Description of the secret group
	Description string `json:"description"`
	// Name of the secret group
	Name string `json:"name"`
	// When the secret group was last updated
	UpdatedAt time.Time                            `json:"updatedAt" format:"date-time"`
	JSON      computeV1SecretListResponseGroupJSON `json:"-"`
}

// computeV1SecretListResponseGroupJSON contains the JSON metadata for the struct
// [ComputeV1SecretListResponseGroup]
type computeV1SecretListResponseGroupJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretListResponseGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretListResponseGroupJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretDeleteGroupResponse struct {
	Message string                                 `json:"message"`
	Success bool                                   `json:"success"`
	JSON    computeV1SecretDeleteGroupResponseJSON `json:"-"`
}

// computeV1SecretDeleteGroupResponseJSON contains the JSON metadata for the struct
// [ComputeV1SecretDeleteGroupResponse]
type computeV1SecretDeleteGroupResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretDeleteGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretDeleteGroupResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretGetGroupResponse struct {
	Group ComputeV1SecretGetGroupResponseGroup `json:"group"`
	Keys  []ComputeV1SecretGetGroupResponseKey `json:"keys"`
	JSON  computeV1SecretGetGroupResponseJSON  `json:"-"`
}

// computeV1SecretGetGroupResponseJSON contains the JSON metadata for the struct
// [ComputeV1SecretGetGroupResponse]
type computeV1SecretGetGroupResponseJSON struct {
	Group       apijson.Field
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretGetGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretGetGroupResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretGetGroupResponseGroup struct {
	// Unique identifier of the secret group
	ID string `json:"id"`
	// Description of the secret group
	Description string `json:"description"`
	// Name of the secret group
	Name string                                   `json:"name"`
	JSON computeV1SecretGetGroupResponseGroupJSON `json:"-"`
}

// computeV1SecretGetGroupResponseGroupJSON contains the JSON metadata for the
// struct [ComputeV1SecretGetGroupResponseGroup]
type computeV1SecretGetGroupResponseGroupJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretGetGroupResponseGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretGetGroupResponseGroupJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretGetGroupResponseKey struct {
	// When the secret was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Name of the secret key
	Key string `json:"key"`
	// When the secret was last updated
	UpdatedAt time.Time                              `json:"updatedAt" format:"date-time"`
	JSON      computeV1SecretGetGroupResponseKeyJSON `json:"-"`
}

// computeV1SecretGetGroupResponseKeyJSON contains the JSON metadata for the struct
// [ComputeV1SecretGetGroupResponseKey]
type computeV1SecretGetGroupResponseKeyJSON struct {
	CreatedAt   apijson.Field
	Key         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretGetGroupResponseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretGetGroupResponseKeyJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretUpdateGroupResponse struct {
	// Number of new secrets created
	Created float64 `json:"created"`
	// Name of the secret group
	Group   string `json:"group"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	// Number of existing secrets updated
	Updated float64                                `json:"updated"`
	JSON    computeV1SecretUpdateGroupResponseJSON `json:"-"`
}

// computeV1SecretUpdateGroupResponseJSON contains the JSON metadata for the struct
// [ComputeV1SecretUpdateGroupResponse]
type computeV1SecretUpdateGroupResponseJSON struct {
	Created     apijson.Field
	Group       apijson.Field
	Message     apijson.Field
	Success     apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1SecretUpdateGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1SecretUpdateGroupResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1SecretNewParams struct {
	// Unique name for the secret group. Must contain only letters, numbers, hyphens,
	// and underscores.
	Name param.Field[string] `json:"name,required"`
	// Optional description of the secret group's purpose
	Description param.Field[string] `json:"description"`
	// Environment name where the secret group will be created. Uses default
	// environment if not specified.
	Env param.Field[string] `json:"env"`
}

func (r ComputeV1SecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputeV1SecretListParams struct {
	// Environment name to list secret groups for. If not specified, uses the default
	// environment.
	Env param.Field[string] `query:"env"`
}

// URLQuery serializes [ComputeV1SecretListParams]'s query parameters as
// `url.Values`.
func (r ComputeV1SecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeV1SecretDeleteGroupParams struct {
	// Environment name. If not provided, uses the default environment
	Env param.Field[string] `query:"env"`
	// Specific key to delete within the group. If not provided, the entire group is
	// deleted
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [ComputeV1SecretDeleteGroupParams]'s query parameters as
// `url.Values`.
func (r ComputeV1SecretDeleteGroupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeV1SecretGetGroupParams struct {
	// Environment name. If not specified, uses the default environment
	Env param.Field[string] `query:"env"`
}

// URLQuery serializes [ComputeV1SecretGetGroupParams]'s query parameters as
// `url.Values`.
func (r ComputeV1SecretGetGroupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeV1SecretUpdateGroupParams struct {
	// Key-value pairs of secrets to set
	Secrets param.Field[map[string]string] `json:"secrets,required"`
	// Environment name (optional, uses default if not specified)
	Env param.Field[string] `json:"env"`
}

func (r ComputeV1SecretUpdateGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
