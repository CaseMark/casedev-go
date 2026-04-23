// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// WebhookV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookV1Service] method instead.
type WebhookV1Service struct {
	Options []option.RequestOption
	// Webhook endpoint management
	Endpoints *WebhookV1EndpointService
	// Webhook endpoint management
	Deliveries *WebhookV1DeliveryService
	// Webhook endpoint management
	EventTypes *WebhookV1EventTypeService
}

// NewWebhookV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookV1Service(opts ...option.RequestOption) (r *WebhookV1Service) {
	r = &WebhookV1Service{}
	r.Options = opts
	r.Endpoints = NewWebhookV1EndpointService(opts...)
	r.Deliveries = NewWebhookV1DeliveryService(opts...)
	r.EventTypes = NewWebhookV1EventTypeService(opts...)
	return
}
