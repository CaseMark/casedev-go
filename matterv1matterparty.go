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

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1MatterPartyService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1MatterPartyService] method instead.
type MatterV1MatterPartyService struct {
	Options []option.RequestOption
}

// NewMatterV1MatterPartyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatterV1MatterPartyService(opts ...option.RequestOption) (r *MatterV1MatterPartyService) {
	r = &MatterV1MatterPartyService{}
	r.Options = opts
	return
}

// Attach a reusable party to a matter with a matter-specific role.
func (r *MatterV1MatterPartyService) New(ctx context.Context, id string, body MatterV1MatterPartyNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/parties", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// List parties attached to a matter.
func (r *MatterV1MatterPartyService) List(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s/parties", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type MatterV1MatterPartyNewParams struct {
	PartyID      param.Field[string]                           `json:"party_id" api:"required"`
	Role         param.Field[MatterV1MatterPartyNewParamsRole] `json:"role" api:"required"`
	CustomFields param.Field[map[string]interface{}]           `json:"custom_fields"`
	IsPrimary    param.Field[bool]                             `json:"is_primary"`
	Metadata     param.Field[map[string]interface{}]           `json:"metadata"`
	Notes        param.Field[string]                           `json:"notes"`
	SetAsClient  param.Field[bool]                             `json:"set_as_client"`
}

func (r MatterV1MatterPartyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1MatterPartyNewParamsRole string

const (
	MatterV1MatterPartyNewParamsRoleClient          MatterV1MatterPartyNewParamsRole = "client"
	MatterV1MatterPartyNewParamsRoleProspect        MatterV1MatterPartyNewParamsRole = "prospect"
	MatterV1MatterPartyNewParamsRoleOpposingParty   MatterV1MatterPartyNewParamsRole = "opposing_party"
	MatterV1MatterPartyNewParamsRoleOpposingCounsel MatterV1MatterPartyNewParamsRole = "opposing_counsel"
	MatterV1MatterPartyNewParamsRoleCoCounsel       MatterV1MatterPartyNewParamsRole = "co_counsel"
	MatterV1MatterPartyNewParamsRoleJudge           MatterV1MatterPartyNewParamsRole = "judge"
	MatterV1MatterPartyNewParamsRoleExpert          MatterV1MatterPartyNewParamsRole = "expert"
	MatterV1MatterPartyNewParamsRoleWitness         MatterV1MatterPartyNewParamsRole = "witness"
	MatterV1MatterPartyNewParamsRoleVendor          MatterV1MatterPartyNewParamsRole = "vendor"
	MatterV1MatterPartyNewParamsRoleInsurer         MatterV1MatterPartyNewParamsRole = "insurer"
	MatterV1MatterPartyNewParamsRoleOther           MatterV1MatterPartyNewParamsRole = "other"
)

func (r MatterV1MatterPartyNewParamsRole) IsKnown() bool {
	switch r {
	case MatterV1MatterPartyNewParamsRoleClient, MatterV1MatterPartyNewParamsRoleProspect, MatterV1MatterPartyNewParamsRoleOpposingParty, MatterV1MatterPartyNewParamsRoleOpposingCounsel, MatterV1MatterPartyNewParamsRoleCoCounsel, MatterV1MatterPartyNewParamsRoleJudge, MatterV1MatterPartyNewParamsRoleExpert, MatterV1MatterPartyNewParamsRoleWitness, MatterV1MatterPartyNewParamsRoleVendor, MatterV1MatterPartyNewParamsRoleInsurer, MatterV1MatterPartyNewParamsRoleOther:
		return true
	}
	return false
}
