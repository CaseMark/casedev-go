// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Matter-native legal workspaces and orchestration primitives
//
// MatterV1PartyService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatterV1PartyService] method instead.
type MatterV1PartyService struct {
	Options []option.RequestOption
}

// NewMatterV1PartyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatterV1PartyService(opts ...option.RequestOption) (r *MatterV1PartyService) {
	r = &MatterV1PartyService{}
	r.Options = opts
	return
}

// Create a reusable legal party for the authenticated organization.
func (r *MatterV1PartyService) New(ctx context.Context, body MatterV1PartyNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "matters/v1/parties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get a reusable legal party by ID.
func (r *MatterV1PartyService) Get(ctx context.Context, partyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if partyID == "" {
		err = errors.New("missing required partyId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/parties/%s", partyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Update a reusable legal party.
func (r *MatterV1PartyService) Update(ctx context.Context, partyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if partyID == "" {
		err = errors.New("missing required partyId parameter")
		return err
	}
	path := fmt.Sprintf("matters/v1/parties/%s", partyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return err
}

// List reusable legal parties for the authenticated organization, newest update
// first. Pagination is opt-in: pass `limit` (1-200) to receive a bounded page,
// then replay `pagination.next_cursor` as `?cursor=` while `pagination.has_more`
// is true. A request with neither `limit` nor `cursor` still returns every party,
// and `pagination.limit` is null. That default will become a bounded page in a
// future release — paginate now to avoid the change.
func (r *MatterV1PartyService) List(ctx context.Context, query MatterV1PartyListParams, opts ...option.RequestOption) (res *MatterV1PartyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "matters/v1/parties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MatterV1PartyListResponse struct {
	Data       []interface{}                       `json:"data"`
	Pagination MatterV1PartyListResponsePagination `json:"pagination"`
	JSON       matterV1PartyListResponseJSON       `json:"-"`
}

// matterV1PartyListResponseJSON contains the JSON metadata for the struct
// [MatterV1PartyListResponse]
type matterV1PartyListResponseJSON struct {
	Data        apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatterV1PartyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1PartyListResponseJSON) RawJSON() string {
	return r.raw
}

type MatterV1PartyListResponsePagination struct {
	HasMore    bool                                    `json:"has_more"`
	Limit      int64                                   `json:"limit" api:"nullable"`
	NextCursor string                                  `json:"next_cursor" api:"nullable"`
	JSON       matterV1PartyListResponsePaginationJSON `json:"-"`
}

// matterV1PartyListResponsePaginationJSON contains the JSON metadata for the
// struct [MatterV1PartyListResponsePagination]
type matterV1PartyListResponsePaginationJSON struct {
	HasMore     apijson.Field
	Limit       apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatterV1PartyListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matterV1PartyListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type MatterV1PartyNewParams struct {
	Name         param.Field[string]                     `json:"name" api:"required"`
	Addresses    param.Field[[]map[string]interface{}]   `json:"addresses"`
	CustomFields param.Field[map[string]interface{}]     `json:"custom_fields"`
	Email        param.Field[string]                     `json:"email"`
	Metadata     param.Field[map[string]interface{}]     `json:"metadata"`
	Notes        param.Field[string]                     `json:"notes"`
	Phone        param.Field[string]                     `json:"phone"`
	Type         param.Field[MatterV1PartyNewParamsType] `json:"type"`
}

func (r MatterV1PartyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MatterV1PartyNewParamsType string

const (
	MatterV1PartyNewParamsTypePerson       MatterV1PartyNewParamsType = "person"
	MatterV1PartyNewParamsTypeOrganization MatterV1PartyNewParamsType = "organization"
)

func (r MatterV1PartyNewParamsType) IsKnown() bool {
	switch r {
	case MatterV1PartyNewParamsTypePerson, MatterV1PartyNewParamsTypeOrganization:
		return true
	}
	return false
}

type MatterV1PartyListParams struct {
	// Opaque continuation cursor from `pagination.next_cursor` of the previous page.
	// Must be replayed with the same filters that produced it.
	Cursor param.Field[string] `query:"cursor"`
	Email  param.Field[string] `query:"email"`
	// Parties per page (1-200). Omit to receive every party. Supplying a cursor
	// without a limit uses 50.
	Limit param.Field[int64]                       `query:"limit"`
	Query param.Field[string]                      `query:"query"`
	Type  param.Field[MatterV1PartyListParamsType] `query:"type"`
}

// URLQuery serializes [MatterV1PartyListParams]'s query parameters as
// `url.Values`.
func (r MatterV1PartyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MatterV1PartyListParamsType string

const (
	MatterV1PartyListParamsTypePerson       MatterV1PartyListParamsType = "person"
	MatterV1PartyListParamsTypeOrganization MatterV1PartyListParamsType = "organization"
)

func (r MatterV1PartyListParamsType) IsKnown() bool {
	switch r {
	case MatterV1PartyListParamsTypePerson, MatterV1PartyListParamsTypeOrganization:
		return true
	}
	return false
}
