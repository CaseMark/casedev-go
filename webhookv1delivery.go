// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Webhook endpoint management
//
// WebhookV1DeliveryService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookV1DeliveryService] method instead.
type WebhookV1DeliveryService struct {
	Options []option.RequestOption
}

// NewWebhookV1DeliveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookV1DeliveryService(opts ...option.RequestOption) (r *WebhookV1DeliveryService) {
	r = &WebhookV1DeliveryService{}
	r.Options = opts
	return
}

// Get webhook delivery
func (r *WebhookV1DeliveryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/deliveries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Returns delivery attempts for the organization, newest first. Filter by
// endpoint_id or status to narrow results.
func (r *WebhookV1DeliveryService) List(ctx context.Context, query WebhookV1DeliveryListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks/v1/deliveries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Re-sends the original event to its endpoint. The payload is reconstructed from
// the delivery record (same eventId, eventType, and occurred_at). The signature
// header includes `svix-delivery-attempt: replay` so receivers can distinguish
// replays from first-time deliveries. Uses the endpoint's current signing secret —
// not the one in force at the original delivery time.
func (r *WebhookV1DeliveryService) Replay(ctx context.Context, id string, body WebhookV1DeliveryReplayParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/deliveries/%s/replay", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type WebhookV1DeliveryListParams struct {
	EndpointID param.Field[string]                            `query:"endpoint_id"`
	Limit      param.Field[int64]                             `query:"limit"`
	Status     param.Field[WebhookV1DeliveryListParamsStatus] `query:"status"`
}

// URLQuery serializes [WebhookV1DeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookV1DeliveryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookV1DeliveryListParamsStatus string

const (
	WebhookV1DeliveryListParamsStatusPending   WebhookV1DeliveryListParamsStatus = "pending"
	WebhookV1DeliveryListParamsStatusDelivered WebhookV1DeliveryListParamsStatus = "delivered"
	WebhookV1DeliveryListParamsStatusFailed    WebhookV1DeliveryListParamsStatus = "failed"
)

func (r WebhookV1DeliveryListParamsStatus) IsKnown() bool {
	switch r {
	case WebhookV1DeliveryListParamsStatusPending, WebhookV1DeliveryListParamsStatusDelivered, WebhookV1DeliveryListParamsStatusFailed:
		return true
	}
	return false
}

type WebhookV1DeliveryReplayParams struct {
	// Override payload to deliver. Must only be supplied when the delivery record
	// lacks enough context to reconstruct the original event (rare). Defaults to an
	// empty data envelope.
	Payload param.Field[interface{}] `json:"payload"`
}

func (r WebhookV1DeliveryReplayParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
