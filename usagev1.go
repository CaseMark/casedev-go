// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Usage reporting and webhook subscriptions
//
// UsageV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageV1Service] method instead.
type UsageV1Service struct {
	Options []option.RequestOption
	// Usage reporting and webhook subscriptions
	Subscriptions *UsageV1SubscriptionService
}

// NewUsageV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageV1Service(opts ...option.RequestOption) (r *UsageV1Service) {
	r = &UsageV1Service{}
	r.Options = opts
	r.Subscriptions = NewUsageV1SubscriptionService(opts...)
	return
}

// Returns customer-facing usage metrics and costs for the requested period.
// Supports summary totals and daily buckets for timestamped usage sources. Vault
// storage is intentionally omitted from totals because it is not yet periodized
// for arbitrary windows.
func (r *UsageV1Service) Get(ctx context.Context, query UsageV1GetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "usage/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type UsageV1GetParams struct {
	// Whether to return period totals only or include daily buckets.
	Granularity param.Field[UsageV1GetParamsGranularity] `query:"granularity"`
	// Period end date. Defaults to now.
	PeriodEnd param.Field[time.Time] `query:"periodEnd" format:"date-time"`
	// Period start date. Defaults to the start of the current calendar month.
	PeriodStart param.Field[time.Time] `query:"periodStart" format:"date-time"`
}

// URLQuery serializes [UsageV1GetParams]'s query parameters as `url.Values`.
func (r UsageV1GetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to return period totals only or include daily buckets.
type UsageV1GetParamsGranularity string

const (
	UsageV1GetParamsGranularitySummary UsageV1GetParamsGranularity = "summary"
	UsageV1GetParamsGranularityDaily   UsageV1GetParamsGranularity = "daily"
)

func (r UsageV1GetParamsGranularity) IsKnown() bool {
	switch r {
	case UsageV1GetParamsGranularitySummary, UsageV1GetParamsGranularityDaily:
		return true
	}
	return false
}
