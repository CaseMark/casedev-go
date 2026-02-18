// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// DatabaseV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabaseV1Service] method instead.
type DatabaseV1Service struct {
	Options  []option.RequestOption
	Projects *DatabaseV1ProjectService
}

// NewDatabaseV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatabaseV1Service(opts ...option.RequestOption) (r *DatabaseV1Service) {
	r = &DatabaseV1Service{}
	r.Options = opts
	r.Projects = NewDatabaseV1ProjectService(opts...)
	return
}

// Returns detailed database usage statistics and billing information for the
// current billing period. Includes compute hours, storage, data transfer, and
// branch counts with associated costs broken down by project.
func (r *DatabaseV1Service) GetUsage(ctx context.Context, opts ...option.RequestOption) (res *DatabaseV1GetUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "database/v1/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DatabaseV1GetUsageResponse struct {
	Period DatabaseV1GetUsageResponsePeriod `json:"period"`
	// Current pricing rates
	Pricing DatabaseV1GetUsageResponsePricing `json:"pricing"`
	// Total number of projects with usage
	ProjectCount int64 `json:"projectCount"`
	// Usage breakdown by project
	Projects []DatabaseV1GetUsageResponseProject `json:"projects"`
	// Aggregated totals across all projects
	Totals DatabaseV1GetUsageResponseTotals `json:"totals"`
	JSON   databaseV1GetUsageResponseJSON   `json:"-"`
}

// databaseV1GetUsageResponseJSON contains the JSON metadata for the struct
// [DatabaseV1GetUsageResponse]
type databaseV1GetUsageResponseJSON struct {
	Period       apijson.Field
	Pricing      apijson.Field
	ProjectCount apijson.Field
	Projects     apijson.Field
	Totals       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DatabaseV1GetUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1GetUsageResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1GetUsageResponsePeriod struct {
	// End of the billing period
	End time.Time `json:"end" format:"date-time"`
	// Start of the billing period
	Start time.Time                            `json:"start" format:"date-time"`
	JSON  databaseV1GetUsageResponsePeriodJSON `json:"-"`
}

// databaseV1GetUsageResponsePeriodJSON contains the JSON metadata for the struct
// [DatabaseV1GetUsageResponsePeriod]
type databaseV1GetUsageResponsePeriodJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1GetUsageResponsePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1GetUsageResponsePeriodJSON) RawJSON() string {
	return r.raw
}

// Current pricing rates
type DatabaseV1GetUsageResponsePricing struct {
	// Cost per branch per month in dollars
	BranchPerMonth float64 `json:"branchPerMonth"`
	// Cost per compute unit hour in dollars
	ComputePerCuHour float64 `json:"computePerCuHour"`
	// Number of free branches included
	FreeBranches int64 `json:"freeBranches"`
	// Cost per GB of storage per month in dollars
	StoragePerGBMonth float64 `json:"storagePerGbMonth"`
	// Cost per GB of data transfer in dollars
	TransferPerGB float64                               `json:"transferPerGb"`
	JSON          databaseV1GetUsageResponsePricingJSON `json:"-"`
}

// databaseV1GetUsageResponsePricingJSON contains the JSON metadata for the struct
// [DatabaseV1GetUsageResponsePricing]
type databaseV1GetUsageResponsePricingJSON struct {
	BranchPerMonth    apijson.Field
	ComputePerCuHour  apijson.Field
	FreeBranches      apijson.Field
	StoragePerGBMonth apijson.Field
	TransferPerGB     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DatabaseV1GetUsageResponsePricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1GetUsageResponsePricingJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1GetUsageResponseProject struct {
	ID              string                                  `json:"id"`
	BranchCount     int64                                   `json:"branchCount"`
	ComputeCuHours  float64                                 `json:"computeCuHours"`
	Costs           DatabaseV1GetUsageResponseProjectsCosts `json:"costs"`
	LastUpdated     time.Time                               `json:"lastUpdated" format:"date-time"`
	ProjectID       string                                  `json:"projectId"`
	ProjectName     string                                  `json:"projectName,nullable"`
	StorageGBMonths float64                                 `json:"storageGbMonths"`
	TransferGB      float64                                 `json:"transferGb"`
	JSON            databaseV1GetUsageResponseProjectJSON   `json:"-"`
}

// databaseV1GetUsageResponseProjectJSON contains the JSON metadata for the struct
// [DatabaseV1GetUsageResponseProject]
type databaseV1GetUsageResponseProjectJSON struct {
	ID              apijson.Field
	BranchCount     apijson.Field
	ComputeCuHours  apijson.Field
	Costs           apijson.Field
	LastUpdated     apijson.Field
	ProjectID       apijson.Field
	ProjectName     apijson.Field
	StorageGBMonths apijson.Field
	TransferGB      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DatabaseV1GetUsageResponseProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1GetUsageResponseProjectJSON) RawJSON() string {
	return r.raw
}

type DatabaseV1GetUsageResponseProjectsCosts struct {
	Branches string                                      `json:"branches"`
	Compute  string                                      `json:"compute"`
	Storage  string                                      `json:"storage"`
	Total    string                                      `json:"total"`
	Transfer string                                      `json:"transfer"`
	JSON     databaseV1GetUsageResponseProjectsCostsJSON `json:"-"`
}

// databaseV1GetUsageResponseProjectsCostsJSON contains the JSON metadata for the
// struct [DatabaseV1GetUsageResponseProjectsCosts]
type databaseV1GetUsageResponseProjectsCostsJSON struct {
	Branches    apijson.Field
	Compute     apijson.Field
	Storage     apijson.Field
	Total       apijson.Field
	Transfer    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseV1GetUsageResponseProjectsCosts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1GetUsageResponseProjectsCostsJSON) RawJSON() string {
	return r.raw
}

// Aggregated totals across all projects
type DatabaseV1GetUsageResponseTotals struct {
	// Total branch cost formatted as dollars
	BranchCostDollars string `json:"branchCostDollars"`
	// Total compute cost formatted as dollars
	ComputeCostDollars string `json:"computeCostDollars"`
	// Total compute unit hours
	ComputeCuHours float64 `json:"computeCuHours"`
	// Total storage cost formatted as dollars
	StorageCostDollars string `json:"storageCostDollars"`
	// Total storage in GB-months
	StorageGBMonths float64 `json:"storageGbMonths"`
	// Total number of branches
	TotalBranches int64 `json:"totalBranches"`
	// Total cost formatted as dollars
	TotalCostDollars string `json:"totalCostDollars"`
	// Total transfer cost formatted as dollars
	TransferCostDollars string `json:"transferCostDollars"`
	// Total data transfer in GB
	TransferGB float64                              `json:"transferGb"`
	JSON       databaseV1GetUsageResponseTotalsJSON `json:"-"`
}

// databaseV1GetUsageResponseTotalsJSON contains the JSON metadata for the struct
// [DatabaseV1GetUsageResponseTotals]
type databaseV1GetUsageResponseTotalsJSON struct {
	BranchCostDollars   apijson.Field
	ComputeCostDollars  apijson.Field
	ComputeCuHours      apijson.Field
	StorageCostDollars  apijson.Field
	StorageGBMonths     apijson.Field
	TotalBranches       apijson.Field
	TotalCostDollars    apijson.Field
	TransferCostDollars apijson.Field
	TransferGB          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DatabaseV1GetUsageResponseTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseV1GetUsageResponseTotalsJSON) RawJSON() string {
	return r.raw
}
