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

// VaultEventSubscriptionService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultEventSubscriptionService] method instead.
type VaultEventSubscriptionService struct {
	Options []option.RequestOption
}

// NewVaultEventSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVaultEventSubscriptionService(opts ...option.RequestOption) (r *VaultEventSubscriptionService) {
	r = &VaultEventSubscriptionService{}
	r.Options = opts
	return
}

// Creates a webhook subscription for vault lifecycle events. Optional object
// filters can limit notifications to specific vault objects.
func (r *VaultEventSubscriptionService) New(ctx context.Context, id string, body VaultEventSubscriptionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/events/subscriptions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Updates callback URL, filters, active state, or signing secret for a vault
// webhook subscription.
func (r *VaultEventSubscriptionService) Update(ctx context.Context, id string, subscriptionID string, body VaultEventSubscriptionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/events/subscriptions/%s", id, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Lists webhook subscriptions configured for a vault.
func (r *VaultEventSubscriptionService) List(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/events/subscriptions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Deactivates a vault webhook subscription.
func (r *VaultEventSubscriptionService) Delete(ctx context.Context, id string, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/events/subscriptions/%s", id, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delivers a test event to a single vault webhook subscription. Uses the same
// payload shape, signature, and retry behavior as production event delivery.
func (r *VaultEventSubscriptionService) Test(ctx context.Context, id string, subscriptionID string, body VaultEventSubscriptionTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("vault/%s/events/subscriptions/%s/test", id, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type VaultEventSubscriptionNewParams struct {
	CallbackURL   param.Field[string]   `json:"callbackUrl,required" format:"uri"`
	EventTypes    param.Field[[]string] `json:"eventTypes"`
	ObjectIDs     param.Field[[]string] `json:"objectIds"`
	SigningSecret param.Field[string]   `json:"signingSecret"`
}

func (r VaultEventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultEventSubscriptionUpdateParams struct {
	CallbackURL        param.Field[string]   `json:"callbackUrl" format:"uri"`
	ClearSigningSecret param.Field[bool]     `json:"clearSigningSecret"`
	EventTypes         param.Field[[]string] `json:"eventTypes"`
	IsActive           param.Field[bool]     `json:"isActive"`
	ObjectIDs          param.Field[[]string] `json:"objectIds"`
	SigningSecret      param.Field[string]   `json:"signingSecret"`
}

func (r VaultEventSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultEventSubscriptionTestParams struct {
	// Optional event type override for this test
	EventType param.Field[string] `json:"eventType"`
	// Optional object ID for object-scoped payload testing
	ObjectID param.Field[string] `json:"objectId"`
	// Optional additional fields merged into payload.data
	Payload param.Field[map[string]interface{}] `json:"payload"`
}

func (r VaultEventSubscriptionTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
