// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// DocumentTemplateFromVaultObjectService contains methods and other services that
// help with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentTemplateFromVaultObjectService] method instead.
type DocumentTemplateFromVaultObjectService struct {
	Options []option.RequestOption
}

// NewDocumentTemplateFromVaultObjectService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDocumentTemplateFromVaultObjectService(opts ...option.RequestOption) (r *DocumentTemplateFromVaultObjectService) {
	r = &DocumentTemplateFromVaultObjectService{}
	r.Options = opts
	return
}

// Promote vault object to document template
func (r *DocumentTemplateFromVaultObjectService) New(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "document-templates/from-vault-object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}
