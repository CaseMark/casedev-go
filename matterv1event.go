// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// MatterV1EventService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1EventService] method instead.
type MatterV1EventService struct {
	Options []option.RequestOption
	// Matter-native legal workspaces and orchestration primitives
	Subscriptions *MatterV1EventSubscriptionService
}

// NewMatterV1EventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1EventService(opts ...option.RequestOption) (r *MatterV1EventService) {
	r = &MatterV1EventService{}
	r.Options = opts
	r.Subscriptions = NewMatterV1EventSubscriptionService(opts...)
	return
}
