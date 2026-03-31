// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1Service] method instead.
type MatterV1Service struct {
	Options []option.RequestOption
	// Matter-native legal workspaces and orchestration primitives
	AgentTypes *MatterV1AgentTypeService
	// Matter-native legal workspaces and orchestration primitives
	Parties *MatterV1PartyService
	// Matter-native legal workspaces and orchestration primitives
	Types  *MatterV1TypeService
	Events *MatterV1EventService
	// Matter-native legal workspaces and orchestration primitives
	Log *MatterV1LogService
	// Matter-native legal workspaces and orchestration primitives
	MatterParties *MatterV1MatterPartyService
	// Matter-native legal workspaces and orchestration primitives
	Shares *MatterV1ShareService
	// Matter-native legal workspaces and orchestration primitives
	WorkItems *MatterV1WorkItemService
}

// NewMatterV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1Service(opts ...option.RequestOption) (r *MatterV1Service) {
	r = &MatterV1Service{}
	r.Options = opts
	r.AgentTypes = NewMatterV1AgentTypeService(opts...)
	r.Parties = NewMatterV1PartyService(opts...)
	r.Types = NewMatterV1TypeService(opts...)
	r.Events = NewMatterV1EventService(opts...)
	r.Log = NewMatterV1LogService(opts...)
	r.MatterParties = NewMatterV1MatterPartyService(opts...)
	r.Shares = NewMatterV1ShareService(opts...)
	r.WorkItems = NewMatterV1WorkItemService(opts...)
	return
}

// Create a new legal matter and optionally link an existing primary vault.
func (r *MatterV1Service) New(ctx context.Context, body MatterV1NewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get a single matter by ID.
func (r *MatterV1Service) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Update mutable matter fields.
func (r *MatterV1Service) Update(ctx context.Context, id string, body MatterV1UpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// List matters for the authenticated organization.
func (r *MatterV1Service) List(ctx context.Context, query MatterV1ListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type MatterV1NewParams struct {
	Title                 param.Field[string]                  `json:"title" api:"required"`
	Billing               param.Field[map[string]interface{}]  `json:"billing"`
	ClientName            param.Field[string]                  `json:"client_name"`
	ClientPartyID         param.Field[string]                  `json:"client_party_id"`
	CustomFields          param.Field[map[string]interface{}]  `json:"custom_fields"`
	Description           param.Field[string]                  `json:"description"`
	DisplayID             param.Field[string]                  `json:"display_id"`
	ImportantDates        param.Field[map[string]interface{}]  `json:"important_dates"`
	Jurisdiction          param.Field[map[string]interface{}]  `json:"jurisdiction"`
	MatterType            param.Field[string]                  `json:"matter_type"`
	Metadata              param.Field[map[string]interface{}]  `json:"metadata"`
	PracticeArea          param.Field[string]                  `json:"practice_area"`
	ResponsibleAttorneyID param.Field[string]                  `json:"responsible_attorney_id"`
	Status                param.Field[MatterV1NewParamsStatus] `json:"status"`
	Subtype               param.Field[string]                  `json:"subtype"`
	Vault                 param.Field[MatterV1NewParamsVault]  `json:"vault"`
	VaultID               param.Field[string]                  `json:"vault_id"`
}

func (r MatterV1NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1NewParamsStatus string

const (
	MatterV1NewParamsStatusIntake   MatterV1NewParamsStatus = "intake"
	MatterV1NewParamsStatusOpen     MatterV1NewParamsStatus = "open"
	MatterV1NewParamsStatusPending  MatterV1NewParamsStatus = "pending"
	MatterV1NewParamsStatusClosed   MatterV1NewParamsStatus = "closed"
	MatterV1NewParamsStatusArchived MatterV1NewParamsStatus = "archived"
)

func (r MatterV1NewParamsStatus) IsKnown() bool {
	switch r {
	case MatterV1NewParamsStatusIntake, MatterV1NewParamsStatusOpen, MatterV1NewParamsStatusPending, MatterV1NewParamsStatusClosed, MatterV1NewParamsStatusArchived:
		return true
	}
	return false
}

type MatterV1NewParamsVault struct {
	Description    param.Field[string]                 `json:"description"`
	EnableGraph    param.Field[bool]                   `json:"enableGraph"`
	EnableIndexing param.Field[bool]                   `json:"enableIndexing"`
	Metadata       param.Field[map[string]interface{}] `json:"metadata"`
}

func (r MatterV1NewParamsVault) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1UpdateParams struct {
	ArchivedAt            param.Field[time.Time]                  `json:"archived_at" format:"date-time"`
	Billing               param.Field[map[string]interface{}]     `json:"billing"`
	ClientName            param.Field[string]                     `json:"client_name"`
	ClientPartyID         param.Field[string]                     `json:"client_party_id"`
	CustomFields          param.Field[map[string]interface{}]     `json:"custom_fields"`
	Description           param.Field[string]                     `json:"description"`
	DisplayID             param.Field[string]                     `json:"display_id"`
	ImportantDates        param.Field[map[string]interface{}]     `json:"important_dates"`
	Jurisdiction          param.Field[map[string]interface{}]     `json:"jurisdiction"`
	MatterType            param.Field[string]                     `json:"matter_type"`
	Metadata              param.Field[map[string]interface{}]     `json:"metadata"`
	PracticeArea          param.Field[string]                     `json:"practice_area"`
	ResponsibleAttorneyID param.Field[string]                     `json:"responsible_attorney_id"`
	Status                param.Field[MatterV1UpdateParamsStatus] `json:"status"`
	Subtype               param.Field[string]                     `json:"subtype"`
	Title                 param.Field[string]                     `json:"title"`
}

func (r MatterV1UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1UpdateParamsStatus string

const (
	MatterV1UpdateParamsStatusIntake   MatterV1UpdateParamsStatus = "intake"
	MatterV1UpdateParamsStatusOpen     MatterV1UpdateParamsStatus = "open"
	MatterV1UpdateParamsStatusPending  MatterV1UpdateParamsStatus = "pending"
	MatterV1UpdateParamsStatusClosed   MatterV1UpdateParamsStatus = "closed"
	MatterV1UpdateParamsStatusArchived MatterV1UpdateParamsStatus = "archived"
)

func (r MatterV1UpdateParamsStatus) IsKnown() bool {
	switch r {
	case MatterV1UpdateParamsStatusIntake, MatterV1UpdateParamsStatusOpen, MatterV1UpdateParamsStatusPending, MatterV1UpdateParamsStatusClosed, MatterV1UpdateParamsStatusArchived:
		return true
	}
	return false
}

type MatterV1ListParams struct {
	MatterType   param.Field[string] `query:"matter_type"`
	PracticeArea param.Field[string] `query:"practice_area"`
	Query        param.Field[string] `query:"query"`
	Status       param.Field[string] `query:"status"`
}

// URLQuery serializes [MatterV1ListParams]'s query parameters as `url.Values`.
func (r MatterV1ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
