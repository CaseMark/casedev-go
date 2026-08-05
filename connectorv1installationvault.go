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

// Import and export between provider folders (Google Drive) and vaults
//
// ConnectorV1InstallationVaultService contains methods and other services that
// help with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1InstallationVaultService] method instead.
type ConnectorV1InstallationVaultService struct {
	Options []option.RequestOption
}

// NewConnectorV1InstallationVaultService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectorV1InstallationVaultService(opts ...option.RequestOption) (r *ConnectorV1InstallationVaultService) {
	r = &ConnectorV1InstallationVaultService{}
	r.Options = opts
	return
}

// List the vaults an installation may use, with capabilities and revocation state.
func (r *ConnectorV1InstallationVaultService) List(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/installations/%s/vaults", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Grant (or update) an installation's access to a vault. Re-granting a revoked
// vault reactivates it. Import links need can_write; export links need can_read;
// mirror deletion and purge need can_manage.
func (r *ConnectorV1InstallationVaultService) Grant(ctx context.Context, id string, vaultID string, body ConnectorV1InstallationVaultGrantParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if vaultID == "" {
		err = errors.New("missing required vaultId parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/installations/%s/vaults/%s", id, vaultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// Revoke an installation's access to a vault. Links using the vault pause at their
// next run; nothing is deleted.
func (r *ConnectorV1InstallationVaultService) Revoke(ctx context.Context, id string, vaultID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if vaultID == "" {
		err = errors.New("missing required vaultId parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/installations/%s/vaults/%s", id, vaultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ConnectorV1InstallationVaultGrantParams struct {
	CanManage    param.Field[bool]                                                `json:"can_manage"`
	CanRead      param.Field[bool]                                                `json:"can_read"`
	CanWrite     param.Field[bool]                                                `json:"can_write"`
	Relationship param.Field[ConnectorV1InstallationVaultGrantParamsRelationship] `json:"relationship"`
	Source       param.Field[ConnectorV1InstallationVaultGrantParamsSource]       `json:"source"`
}

func (r ConnectorV1InstallationVaultGrantParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1InstallationVaultGrantParamsRelationship string

const (
	ConnectorV1InstallationVaultGrantParamsRelationshipOwned  ConnectorV1InstallationVaultGrantParamsRelationship = "owned"
	ConnectorV1InstallationVaultGrantParamsRelationshipShared ConnectorV1InstallationVaultGrantParamsRelationship = "shared"
)

func (r ConnectorV1InstallationVaultGrantParamsRelationship) IsKnown() bool {
	switch r {
	case ConnectorV1InstallationVaultGrantParamsRelationshipOwned, ConnectorV1InstallationVaultGrantParamsRelationshipShared:
		return true
	}
	return false
}

type ConnectorV1InstallationVaultGrantParamsSource string

const (
	ConnectorV1InstallationVaultGrantParamsSourceProvisioning  ConnectorV1InstallationVaultGrantParamsSource = "provisioning"
	ConnectorV1InstallationVaultGrantParamsSourceLazyReconcile ConnectorV1InstallationVaultGrantParamsSource = "lazy_reconcile"
	ConnectorV1InstallationVaultGrantParamsSourceExplicitShare ConnectorV1InstallationVaultGrantParamsSource = "explicit_share"
)

func (r ConnectorV1InstallationVaultGrantParamsSource) IsKnown() bool {
	switch r {
	case ConnectorV1InstallationVaultGrantParamsSourceProvisioning, ConnectorV1InstallationVaultGrantParamsSourceLazyReconcile, ConnectorV1InstallationVaultGrantParamsSourceExplicitShare:
		return true
	}
	return false
}
