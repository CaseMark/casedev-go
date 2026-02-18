// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"github.com/CaseMark/casedev-go/option"
)

// ApplicationV1Service contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationV1Service] method instead.
type ApplicationV1Service struct {
	Options     []option.RequestOption
	Deployments *ApplicationV1DeploymentService
	Projects    *ApplicationV1ProjectService
	Workflows   *ApplicationV1WorkflowService
}

// NewApplicationV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApplicationV1Service(opts ...option.RequestOption) (r *ApplicationV1Service) {
	r = &ApplicationV1Service{}
	r.Options = opts
	r.Deployments = NewApplicationV1DeploymentService(opts...)
	r.Projects = NewApplicationV1ProjectService(opts...)
	r.Workflows = NewApplicationV1WorkflowService(opts...)
	return
}
