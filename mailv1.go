// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// MailV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailV1Service] method instead.
type MailV1Service struct {
	Options []option.RequestOption
	// Create, manage, and execute AI agents with tool access, sandbox environments,
	// and async run workflows
	Inboxes *MailV1InboxService
}

// NewMailV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMailV1Service(opts ...option.RequestOption) (r *MailV1Service) {
	r = &MailV1Service{}
	r.Options = opts
	r.Inboxes = NewMailV1InboxService(opts...)
	return
}
