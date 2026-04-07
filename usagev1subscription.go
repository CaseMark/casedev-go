// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Usage reporting and webhook subscriptions
//
// UsageV1SubscriptionService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageV1SubscriptionService] method instead.
type UsageV1SubscriptionService struct {
	Options []option.RequestOption
}

// NewUsageV1SubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUsageV1SubscriptionService(opts ...option.RequestOption) (r *UsageV1SubscriptionService) {
	r = &UsageV1SubscriptionService{}
	r.Options = opts
	return
}

// Creates a webhook subscription for usage, balance, and billing events.
func (r *UsageV1SubscriptionService) New(ctx context.Context, body UsageV1SubscriptionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "usage/v1/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Updates callback URL, event filters, active state, or signing secret.
func (r *UsageV1SubscriptionService) Update(ctx context.Context, subscriptionID string, body UsageV1SubscriptionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return err
	}
	path := fmt.Sprintf("usage/v1/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Lists webhook subscriptions configured for usage and billing events.
func (r *UsageV1SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "usage/v1/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Deactivates a usage webhook subscription.
func (r *UsageV1SubscriptionService) Delete(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return err
	}
	path := fmt.Sprintf("usage/v1/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Delivers a test event to a single usage webhook subscription using the same
// payload shape and signing behavior as production delivery.
func (r *UsageV1SubscriptionService) Test(ctx context.Context, subscriptionID string, body UsageV1SubscriptionTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return err
	}
	path := fmt.Sprintf("usage/v1/subscriptions/%s/test", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type UsageV1SubscriptionNewParams struct {
	CallbackURL   param.Field[string]   `json:"callbackUrl" api:"required" format:"uri"`
	EventTypes    param.Field[[]string] `json:"eventTypes"`
	SigningSecret param.Field[string]   `json:"signingSecret"`
}

func (r UsageV1SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageV1SubscriptionUpdateParams struct {
	CallbackURL        param.Field[string]   `json:"callbackUrl" format:"uri"`
	ClearSigningSecret param.Field[bool]     `json:"clearSigningSecret"`
	EventTypes         param.Field[[]string] `json:"eventTypes"`
	IsActive           param.Field[bool]     `json:"isActive"`
	SigningSecret      param.Field[string]   `json:"signingSecret"`
}

func (r UsageV1SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageV1SubscriptionTestParams struct {
	EventType param.Field[string]                 `json:"eventType"`
	Payload   param.Field[map[string]interface{}] `json:"payload"`
}

func (r UsageV1SubscriptionTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
