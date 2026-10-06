// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Import and export between provider folders and vaults
//
// ConnectorV1InstallationTokenService contains methods and other services that
// help with interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorV1InstallationTokenService] method instead.
type ConnectorV1InstallationTokenService struct {
	Options []option.RequestOption
}

// NewConnectorV1InstallationTokenService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectorV1InstallationTokenService(opts ...option.RequestOption) (r *ConnectorV1InstallationTokenService) {
	r = &ConnectorV1InstallationTokenService{}
	r.Options = opts
	return
}

// Bound application management API key only. Returns a 256-bit opaque bearer
// secret once; valid for 15 minutes, connector data-plane only. No refresh token.
// At most 4 overlapping credentials and 20 mints/minute per installation. The
// application server must derive tenant identity; this is not browser
// authentication.
func (r *ConnectorV1InstallationTokenService) New(ctx context.Context, id string, body ConnectorV1InstallationTokenNewParams, opts ...option.RequestOption) (res *ConnectorV1InstallationTokenNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connectors/v1/installations/%s/tokens", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Issuing API key only; installation credentials cannot revoke or mint
// credentials. Immediate revocation on the next request.
func (r *ConnectorV1InstallationTokenService) Revoke(ctx context.Context, id string, tokenID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if tokenID == "" {
		err = errors.New("missing required tokenId parameter")
		return err
	}
	path := fmt.Sprintf("connectors/v1/installations/%s/tokens/%s", id, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ConnectorV1InstallationTokenNewResponse struct {
	ID             string                                      `json:"id" api:"required"`
	Token          string                                      `json:"token" api:"required"`
	ExpiresAt      time.Time                                   `json:"expires_at" api:"required" format:"date-time"`
	InstallationID string                                      `json:"installation_id" api:"required"`
	JSON           connectorV1InstallationTokenNewResponseJSON `json:"-"`
}

// connectorV1InstallationTokenNewResponseJSON contains the JSON metadata for the
// struct [ConnectorV1InstallationTokenNewResponse]
type connectorV1InstallationTokenNewResponseJSON struct {
	ID             apijson.Field
	Token          apijson.Field
	ExpiresAt      apijson.Field
	InstallationID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectorV1InstallationTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorV1InstallationTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorV1InstallationTokenNewParams struct {
	Scopes param.Field[[]ConnectorV1InstallationTokenNewParamsScope] `json:"scopes" api:"required"`
}

func (r ConnectorV1InstallationTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorV1InstallationTokenNewParamsScope string

const (
	ConnectorV1InstallationTokenNewParamsScopeRead  ConnectorV1InstallationTokenNewParamsScope = "read"
	ConnectorV1InstallationTokenNewParamsScopeWrite ConnectorV1InstallationTokenNewParamsScope = "write"
)

func (r ConnectorV1InstallationTokenNewParamsScope) IsKnown() bool {
	switch r {
	case ConnectorV1InstallationTokenNewParamsScopeRead, ConnectorV1InstallationTokenNewParamsScopeWrite:
		return true
	}
	return false
}
