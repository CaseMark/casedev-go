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
// WebhookV1EndpointService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookV1EndpointService] method instead.
type WebhookV1EndpointService struct {
	Options []option.RequestOption
}

// NewWebhookV1EndpointService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookV1EndpointService(opts ...option.RequestOption) (r *WebhookV1EndpointService) {
	r = &WebhookV1EndpointService{}
	r.Options = opts
	return
}

// Creates a webhook endpoint that receives platform events matching the supplied
// event-type filters. Returns the generated signing secret ONCE — the response is
// the only time it is shown in plaintext.
func (r *WebhookV1EndpointService) New(ctx context.Context, body WebhookV1EndpointNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks/v1/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get webhook endpoint
func (r *WebhookV1EndpointService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Partially updates a webhook endpoint. Any omitted field is left unchanged.
// Signing secrets are rotated via the separate /rotate_secret endpoint.
func (r *WebhookV1EndpointService) Update(ctx context.Context, id string, body WebhookV1EndpointUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Returns the organization's webhook endpoints, newest first. Signing secrets are
// never included.
func (r *WebhookV1EndpointService) List(ctx context.Context, query WebhookV1EndpointListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks/v1/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Soft-deletes a webhook endpoint. Delivery stops immediately and the endpoint no
// longer appears in list results. Delivery history is preserved (and can be
// fetched via GET /deliveries with the endpoint_id filter) so audit trails and
// post-mortem debugging remain possible.
func (r *WebhookV1EndpointService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/endpoints/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Generates a new signing secret for the endpoint. The previous secret remains
// valid until `previousSecretExpiresInSec` elapses (default 24h, max 30 days).
// During the grace window deliveries are signed with both secrets so receivers can
// migrate without downtime. Returns the new secret — this is the only time it is
// shown in plaintext.
func (r *WebhookV1EndpointService) RotateSecret(ctx context.Context, id string, body WebhookV1EndpointRotateSecretParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/endpoints/%s/rotate_secret", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Synchronously delivers a synthetic `webhook.test` event to the endpoint and
// returns the HTTP result. No retries. Useful for validating that a new endpoint
// is reachable and its signature verifier works. The delivery is not persisted in
// the delivery history.
func (r *WebhookV1EndpointService) Test(ctx context.Context, id string, body WebhookV1EndpointTestParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("webhooks/v1/endpoints/%s/test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type WebhookV1EndpointNewParams struct {
	// Glob patterns of event types to deliver (e.g. "vault._", "ocr.job.completed",
	// "_")
	EventTypeFilters param.Field[[]string] `json:"eventTypeFilters" api:"required"`
	// HTTPS callback URL that will receive event deliveries
	URL param.Field[string] `json:"url" api:"required" format:"uri"`
	// Human-readable label for this endpoint
	Description param.Field[string] `json:"description"`
	// Optional per-resource allowlists. If vaultIds is set, only events for those
	// vaults are delivered. Same for matterIds.
	ResourceScopes param.Field[WebhookV1EndpointNewParamsResourceScopes] `json:"resourceScopes"`
}

func (r WebhookV1EndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional per-resource allowlists. If vaultIds is set, only events for those
// vaults are delivered. Same for matterIds.
type WebhookV1EndpointNewParamsResourceScopes struct {
	MatterIDs param.Field[[]string] `json:"matterIds"`
	VaultIDs  param.Field[[]string] `json:"vaultIds"`
}

func (r WebhookV1EndpointNewParamsResourceScopes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookV1EndpointUpdateParams struct {
	Description      param.Field[string]                                      `json:"description"`
	EventTypeFilters param.Field[[]string]                                    `json:"eventTypeFilters"`
	ResourceScopes   param.Field[WebhookV1EndpointUpdateParamsResourceScopes] `json:"resourceScopes"`
	Status           param.Field[WebhookV1EndpointUpdateParamsStatus]         `json:"status"`
	URL              param.Field[string]                                      `json:"url" format:"uri"`
}

func (r WebhookV1EndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookV1EndpointUpdateParamsResourceScopes struct {
	MatterIDs param.Field[[]string] `json:"matterIds"`
	VaultIDs  param.Field[[]string] `json:"vaultIds"`
}

func (r WebhookV1EndpointUpdateParamsResourceScopes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookV1EndpointUpdateParamsStatus string

const (
	WebhookV1EndpointUpdateParamsStatusActive   WebhookV1EndpointUpdateParamsStatus = "active"
	WebhookV1EndpointUpdateParamsStatusDisabled WebhookV1EndpointUpdateParamsStatus = "disabled"
)

func (r WebhookV1EndpointUpdateParamsStatus) IsKnown() bool {
	switch r {
	case WebhookV1EndpointUpdateParamsStatusActive, WebhookV1EndpointUpdateParamsStatusDisabled:
		return true
	}
	return false
}

type WebhookV1EndpointListParams struct {
	Limit param.Field[int64] `query:"limit"`
	// Filter by endpoint status
	Status param.Field[WebhookV1EndpointListParamsStatus] `query:"status"`
}

// URLQuery serializes [WebhookV1EndpointListParams]'s query parameters as
// `url.Values`.
func (r WebhookV1EndpointListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by endpoint status
type WebhookV1EndpointListParamsStatus string

const (
	WebhookV1EndpointListParamsStatusActive       WebhookV1EndpointListParamsStatus = "active"
	WebhookV1EndpointListParamsStatusDisabled     WebhookV1EndpointListParamsStatus = "disabled"
	WebhookV1EndpointListParamsStatusAutoDisabled WebhookV1EndpointListParamsStatus = "auto_disabled"
)

func (r WebhookV1EndpointListParamsStatus) IsKnown() bool {
	switch r {
	case WebhookV1EndpointListParamsStatusActive, WebhookV1EndpointListParamsStatusDisabled, WebhookV1EndpointListParamsStatusAutoDisabled:
		return true
	}
	return false
}

type WebhookV1EndpointRotateSecretParams struct {
	// How long (seconds) the old secret continues to be accepted. 0 invalidates
	// immediately. Default: 86400 (24h).
	PreviousSecretExpiresInSec param.Field[int64] `json:"previousSecretExpiresInSec"`
}

func (r WebhookV1EndpointRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookV1EndpointTestParams struct {
	// Event type to simulate. Defaults to "webhook.test".
	EventType param.Field[string] `json:"eventType"`
	// Custom `data` payload. Defaults to a small placeholder.
	Payload param.Field[interface{}] `json:"payload"`
}

func (r WebhookV1EndpointTestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
