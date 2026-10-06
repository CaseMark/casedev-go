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
func (r *WebhookV1EndpointService) New(ctx context.Context, body WebhookV1EndpointNewParams, opts ...option.RequestOption) (res *WebhookV1EndpointNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks/v1/endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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

type WebhookV1EndpointNewResponse struct {
	Endpoint WebhookV1EndpointNewResponseEndpoint `json:"endpoint" api:"required"`
	// One-time webhook signing secret.
	SigningSecret string                           `json:"signingSecret" api:"required"`
	JSON          webhookV1EndpointNewResponseJSON `json:"-"`
}

// webhookV1EndpointNewResponseJSON contains the JSON metadata for the struct
// [WebhookV1EndpointNewResponse]
type webhookV1EndpointNewResponseJSON struct {
	Endpoint      apijson.Field
	SigningSecret apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WebhookV1EndpointNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookV1EndpointNewResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookV1EndpointNewResponseEndpoint struct {
	ID                             string                                             `json:"id" api:"required"`
	ConsecutiveFailureCount        int64                                              `json:"consecutiveFailureCount" api:"required"`
	CreatedAt                      time.Time                                          `json:"createdAt" api:"required" format:"date-time"`
	Description                    string                                             `json:"description" api:"required,nullable"`
	DisabledReason                 string                                             `json:"disabledReason" api:"required,nullable"`
	EventTypeFilters               []string                                           `json:"eventTypeFilters" api:"required"`
	HasPreviousSigningSecret       bool                                               `json:"hasPreviousSigningSecret" api:"required"`
	LastFailureAt                  time.Time                                          `json:"lastFailureAt" api:"required,nullable" format:"date-time"`
	LastSuccessAt                  time.Time                                          `json:"lastSuccessAt" api:"required,nullable" format:"date-time"`
	PreviousSigningSecretExpiresAt time.Time                                          `json:"previousSigningSecretExpiresAt" api:"required,nullable" format:"date-time"`
	ResourceScopes                 WebhookV1EndpointNewResponseEndpointResourceScopes `json:"resourceScopes" api:"required,nullable"`
	Status                         WebhookV1EndpointNewResponseEndpointStatus         `json:"status" api:"required"`
	UpdatedAt                      time.Time                                          `json:"updatedAt" api:"required" format:"date-time"`
	URL                            string                                             `json:"url" api:"required" format:"uri"`
	JSON                           webhookV1EndpointNewResponseEndpointJSON           `json:"-"`
}

// webhookV1EndpointNewResponseEndpointJSON contains the JSON metadata for the
// struct [WebhookV1EndpointNewResponseEndpoint]
type webhookV1EndpointNewResponseEndpointJSON struct {
	ID                             apijson.Field
	ConsecutiveFailureCount        apijson.Field
	CreatedAt                      apijson.Field
	Description                    apijson.Field
	DisabledReason                 apijson.Field
	EventTypeFilters               apijson.Field
	HasPreviousSigningSecret       apijson.Field
	LastFailureAt                  apijson.Field
	LastSuccessAt                  apijson.Field
	PreviousSigningSecretExpiresAt apijson.Field
	ResourceScopes                 apijson.Field
	Status                         apijson.Field
	UpdatedAt                      apijson.Field
	URL                            apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *WebhookV1EndpointNewResponseEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookV1EndpointNewResponseEndpointJSON) RawJSON() string {
	return r.raw
}

type WebhookV1EndpointNewResponseEndpointResourceScopes struct {
	MatterIDs []string                                               `json:"matterIds"`
	VaultIDs  []string                                               `json:"vaultIds"`
	JSON      webhookV1EndpointNewResponseEndpointResourceScopesJSON `json:"-"`
}

// webhookV1EndpointNewResponseEndpointResourceScopesJSON contains the JSON
// metadata for the struct [WebhookV1EndpointNewResponseEndpointResourceScopes]
type webhookV1EndpointNewResponseEndpointResourceScopesJSON struct {
	MatterIDs   apijson.Field
	VaultIDs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookV1EndpointNewResponseEndpointResourceScopes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookV1EndpointNewResponseEndpointResourceScopesJSON) RawJSON() string {
	return r.raw
}

type WebhookV1EndpointNewResponseEndpointStatus string

const (
	WebhookV1EndpointNewResponseEndpointStatusActive       WebhookV1EndpointNewResponseEndpointStatus = "active"
	WebhookV1EndpointNewResponseEndpointStatusDisabled     WebhookV1EndpointNewResponseEndpointStatus = "disabled"
	WebhookV1EndpointNewResponseEndpointStatusAutoDisabled WebhookV1EndpointNewResponseEndpointStatus = "auto_disabled"
)

func (r WebhookV1EndpointNewResponseEndpointStatus) IsKnown() bool {
	switch r {
	case WebhookV1EndpointNewResponseEndpointStatusActive, WebhookV1EndpointNewResponseEndpointStatusDisabled, WebhookV1EndpointNewResponseEndpointStatusAutoDisabled:
		return true
	}
	return false
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
