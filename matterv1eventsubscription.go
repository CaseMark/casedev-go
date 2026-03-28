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

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1EventSubscriptionService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1EventSubscriptionService] method instead.
type MatterV1EventSubscriptionService struct {
	Options []option.RequestOption
}

// NewMatterV1EventSubscriptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMatterV1EventSubscriptionService(opts ...option.RequestOption) (r *MatterV1EventSubscriptionService) {
	r = &MatterV1EventSubscriptionService{}
	r.Options = opts
	return
}

// Creates a webhook subscription for matter and work-item events.
func (r *MatterV1EventSubscriptionService) New(ctx context.Context, id string, body MatterV1EventSubscriptionNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/events/subscriptions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Lists webhook subscriptions configured for a matter.
func (r *MatterV1EventSubscriptionService) List(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/events/subscriptions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Deactivates a matter webhook subscription.
func (r *MatterV1EventSubscriptionService) Delete(ctx context.Context, id string, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/events/subscriptions/%s", id, subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type MatterV1EventSubscriptionNewParams struct {
	CallbackURL   param.Field[string]   `json:"callbackUrl" api:"required" format:"uri"`
	EventTypes    param.Field[[]string] `json:"eventTypes"`
	SigningSecret param.Field[string]   `json:"signingSecret"`
}

func (r MatterV1EventSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
