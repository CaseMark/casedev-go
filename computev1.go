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

// ComputeV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeV1Service] method instead.
type ComputeV1Service struct {
	Options       []option.RequestOption
	Environments  *ComputeV1EnvironmentService
	InstanceTypes *ComputeV1InstanceTypeService
	Instances     *ComputeV1InstanceService
	Secrets       *ComputeV1SecretService
}

// NewComputeV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeV1Service(opts ...option.RequestOption) (r *ComputeV1Service) {
	r = &ComputeV1Service{}
	r.Options = opts
	r.Environments = NewComputeV1EnvironmentService(opts...)
	r.InstanceTypes = NewComputeV1InstanceTypeService(opts...)
	r.Instances = NewComputeV1InstanceService(opts...)
	r.Secrets = NewComputeV1SecretService(opts...)
	return
}

// Returns current pricing for GPU instances. Prices are fetched in real-time and
// include a 20% platform fee. For detailed instance types and availability, use
// GET /compute/v1/instance-types.
func (r *ComputeV1Service) GetPricing(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "compute/v1/pricing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Returns detailed compute usage statistics and billing information for your
// organization. Includes GPU and CPU hours, total runs, costs, and breakdowns by
// environment. Use optional query parameters to filter by specific year and month.
func (r *ComputeV1Service) GetUsage(ctx context.Context, query ComputeV1GetUsageParams, opts ...option.RequestOption) (res *ComputeV1GetUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ComputeV1GetUsageResponse struct {
	ByEnvironment []ComputeV1GetUsageResponseByEnvironment `json:"byEnvironment"`
	Period        ComputeV1GetUsageResponsePeriod          `json:"period"`
	Summary       ComputeV1GetUsageResponseSummary         `json:"summary"`
	JSON          computeV1GetUsageResponseJSON            `json:"-"`
}

// computeV1GetUsageResponseJSON contains the JSON metadata for the struct
// [ComputeV1GetUsageResponse]
type computeV1GetUsageResponseJSON struct {
	ByEnvironment apijson.Field
	Period        apijson.Field
	Summary       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ComputeV1GetUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1GetUsageResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1GetUsageResponseByEnvironment struct {
	Environment        string                                     `json:"environment"`
	TotalCostCents     int64                                      `json:"totalCostCents"`
	TotalCostFormatted string                                     `json:"totalCostFormatted"`
	TotalCPUSeconds    int64                                      `json:"totalCpuSeconds"`
	TotalGPUSeconds    int64                                      `json:"totalGpuSeconds"`
	TotalRuns          int64                                      `json:"totalRuns"`
	JSON               computeV1GetUsageResponseByEnvironmentJSON `json:"-"`
}

// computeV1GetUsageResponseByEnvironmentJSON contains the JSON metadata for the
// struct [ComputeV1GetUsageResponseByEnvironment]
type computeV1GetUsageResponseByEnvironmentJSON struct {
	Environment        apijson.Field
	TotalCostCents     apijson.Field
	TotalCostFormatted apijson.Field
	TotalCPUSeconds    apijson.Field
	TotalGPUSeconds    apijson.Field
	TotalRuns          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ComputeV1GetUsageResponseByEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1GetUsageResponseByEnvironmentJSON) RawJSON() string {
	return r.raw
}

type ComputeV1GetUsageResponsePeriod struct {
	Month     int64                               `json:"month"`
	MonthName string                              `json:"monthName"`
	Year      int64                               `json:"year"`
	JSON      computeV1GetUsageResponsePeriodJSON `json:"-"`
}

// computeV1GetUsageResponsePeriodJSON contains the JSON metadata for the struct
// [ComputeV1GetUsageResponsePeriod]
type computeV1GetUsageResponsePeriodJSON struct {
	Month       apijson.Field
	MonthName   apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1GetUsageResponsePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1GetUsageResponsePeriodJSON) RawJSON() string {
	return r.raw
}

type ComputeV1GetUsageResponseSummary struct {
	TotalCostCents     int64                                `json:"totalCostCents"`
	TotalCostFormatted string                               `json:"totalCostFormatted"`
	TotalCPUHours      float64                              `json:"totalCpuHours"`
	TotalGPUHours      float64                              `json:"totalGpuHours"`
	TotalRuns          int64                                `json:"totalRuns"`
	JSON               computeV1GetUsageResponseSummaryJSON `json:"-"`
}

// computeV1GetUsageResponseSummaryJSON contains the JSON metadata for the struct
// [ComputeV1GetUsageResponseSummary]
type computeV1GetUsageResponseSummaryJSON struct {
	TotalCostCents     apijson.Field
	TotalCostFormatted apijson.Field
	TotalCPUHours      apijson.Field
	TotalGPUHours      apijson.Field
	TotalRuns          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ComputeV1GetUsageResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1GetUsageResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type ComputeV1GetUsageParams struct {
	// Month to filter usage data (1-12, defaults to current month)
	Month param.Field[int64] `query:"month"`
	// Year to filter usage data (defaults to current year)
	Year param.Field[int64] `query:"year"`
}

// URLQuery serializes [ComputeV1GetUsageParams]'s query parameters as
// `url.Values`.
func (r ComputeV1GetUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
