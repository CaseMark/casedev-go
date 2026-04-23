// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Webhook endpoint management
//
// WebhookV1EventTypeService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookV1EventTypeService] method instead.
type WebhookV1EventTypeService struct {
	Options []option.RequestOption
}

// NewWebhookV1EventTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookV1EventTypeService(opts ...option.RequestOption) (r *WebhookV1EventTypeService) {
	r = &WebhookV1EventTypeService{}
	r.Options = opts
	return
}

// Returns the catalog of event types that can be subscribed to via webhook
// endpoints. Each entry lists the required service scope the API key must carry to
// subscribe, plus the stability level.
func (r *WebhookV1EventTypeService) List(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "webhooks/v1/event_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}
