// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Import and export between provider folders and vaults
//
// ConnectorV1ApplicationKeyService contains methods and other services that help
// with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1ApplicationKeyService] method instead.
type ConnectorV1ApplicationKeyService struct {
	Options []option.RequestOption
}

// NewConnectorV1ApplicationKeyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectorV1ApplicationKeyService(opts ...option.RequestOption) (r *ConnectorV1ApplicationKeyService) {
	r = &ConnectorV1ApplicationKeyService{}
	r.Options = opts
	return
}

// Requires an owner/admin Clerk session. Explicitly binds a non-system connector
// API key to one application. Allows multiple keys for controlled rotation. Does
// not grant Vault access or mint credentials.
func (r *ConnectorV1ApplicationKeyService) Bind(ctx context.Context, id string, keyID string, opts ...option.RequestOption) (res *ConnectorV1ApplicationKeyBindResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required keyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/applications/%s/keys/%s", id, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Requires an owner/admin Clerk session. Immediately invalidates tokens minted
// through this binding. Rebinding later does not revive those tokens.
func (r *ConnectorV1ApplicationKeyService) Revoke(ctx context.Context, id string, keyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if keyID == "" {
		err = errors.New("missing required keyId parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/applications/%s/keys/%s", id, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ConnectorV1ApplicationKeyBindResponse struct {
	ID            string                                    `json:"id" api:"required"`
	APIKeyID      string                                    `json:"api_key_id" api:"required"`
	ApplicationID string                                    `json:"application_id" api:"required"`
	JSON          connectorV1ApplicationKeyBindResponseJSON `json:"-"`
}

// connectorV1ApplicationKeyBindResponseJSON contains the JSON metadata for the
// struct [ConnectorV1ApplicationKeyBindResponse]
type connectorV1ApplicationKeyBindResponseJSON struct {
	ID            apijson.Field
	APIKeyID      apijson.Field
	ApplicationID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectorV1ApplicationKeyBindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1ApplicationKeyBindResponseJSON) RawJSON() string {
	return r.raw
}
