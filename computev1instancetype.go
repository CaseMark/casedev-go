// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// ComputeV1InstanceTypeService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeV1InstanceTypeService] method instead.
type ComputeV1InstanceTypeService struct {
	Options []option.RequestOption
}

// NewComputeV1InstanceTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeV1InstanceTypeService(opts ...option.RequestOption) (r *ComputeV1InstanceTypeService) {
	r = &ComputeV1InstanceTypeService{}
	r.Options = opts
	return
}

// Retrieves all available GPU instance types with pricing, specifications, and
// regional availability. Includes T4, A10, A100, H100, and H200 GPUs powered by
// Lambda Labs. Perfect for AI model training, inference workloads, and legal
// document OCR processing at scale.
func (r *ComputeV1InstanceTypeService) List(ctx context.Context, opts ...option.RequestOption) (res *ComputeV1InstanceTypeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/instance-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ComputeV1InstanceTypeListResponse struct {
	// Total number of instance types
	Count         int64                                           `json:"count,required"`
	InstanceTypes []ComputeV1InstanceTypeListResponseInstanceType `json:"instanceTypes,required"`
	JSON          computeV1InstanceTypeListResponseJSON           `json:"-"`
}

// computeV1InstanceTypeListResponseJSON contains the JSON metadata for the struct
// [ComputeV1InstanceTypeListResponse]
type computeV1InstanceTypeListResponseJSON struct {
	Count         apijson.Field
	InstanceTypes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ComputeV1InstanceTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceTypeListResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceTypeListResponseInstanceType struct {
	// Instance description
	Description string `json:"description"`
	// GPU model and count
	GPU string `json:"gpu"`
	// Instance type identifier
	Name string `json:"name"`
	// Price per hour (e.g. '$1.20')
	PricePerHour string `json:"pricePerHour"`
	// Available regions
	RegionsAvailable []string                                            `json:"regionsAvailable"`
	Specs            ComputeV1InstanceTypeListResponseInstanceTypesSpecs `json:"specs"`
	JSON             computeV1InstanceTypeListResponseInstanceTypeJSON   `json:"-"`
}

// computeV1InstanceTypeListResponseInstanceTypeJSON contains the JSON metadata for
// the struct [ComputeV1InstanceTypeListResponseInstanceType]
type computeV1InstanceTypeListResponseInstanceTypeJSON struct {
	Description      apijson.Field
	GPU              apijson.Field
	Name             apijson.Field
	PricePerHour     apijson.Field
	RegionsAvailable apijson.Field
	Specs            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ComputeV1InstanceTypeListResponseInstanceType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceTypeListResponseInstanceTypeJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceTypeListResponseInstanceTypesSpecs struct {
	// RAM in GiB
	MemoryGib int64 `json:"memoryGib"`
	// Storage in GiB
	StorageGib int64 `json:"storageGib"`
	// Number of vCPUs
	Vcpus int64                                                   `json:"vcpus"`
	JSON  computeV1InstanceTypeListResponseInstanceTypesSpecsJSON `json:"-"`
}

// computeV1InstanceTypeListResponseInstanceTypesSpecsJSON contains the JSON
// metadata for the struct [ComputeV1InstanceTypeListResponseInstanceTypesSpecs]
type computeV1InstanceTypeListResponseInstanceTypesSpecsJSON struct {
	MemoryGib   apijson.Field
	StorageGib  apijson.Field
	Vcpus       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1InstanceTypeListResponseInstanceTypesSpecs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceTypeListResponseInstanceTypesSpecsJSON) RawJSON() string {
	return r.raw
}
