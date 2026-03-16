// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// OperatorV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperatorV1Service] method instead.
type OperatorV1Service struct {
	Options []option.RequestOption
}

// NewOperatorV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperatorV1Service(opts ...option.RequestOption) (r *OperatorV1Service) {
	r = &OperatorV1Service{}
	r.Options = opts
	return
}

// Provision a new operator instance for the organization.
func (r *OperatorV1Service) New(ctx context.Context, body OperatorV1NewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "operator/v1/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Proxy a chat completion request to the organization's operator instance.
func (r *OperatorV1Service) NewChatCompletion(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "operator/v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Proxy a response request to the organization's operator instance.
func (r *OperatorV1Service) NewResponse(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "operator/v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Get the status of the organization's operator instance.
func (r *OperatorV1Service) GetStatus(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "operator/v1/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type OperatorV1NewParams struct {
	// Operator name
	Name param.Field[string] `json:"name" api:"required"`
	// Model to use
	Model param.Field[string] `json:"model"`
	// Instance size
	Size param.Field[OperatorV1NewParamsSize] `json:"size"`
}

func (r OperatorV1NewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Instance size
type OperatorV1NewParamsSize string

const (
	OperatorV1NewParamsSizeSmall  OperatorV1NewParamsSize = "small"
	OperatorV1NewParamsSizeMedium OperatorV1NewParamsSize = "medium"
	OperatorV1NewParamsSizeLarge  OperatorV1NewParamsSize = "large"
)

func (r OperatorV1NewParamsSize) IsKnown() bool {
	switch r {
	case OperatorV1NewParamsSizeSmall, OperatorV1NewParamsSizeMedium, OperatorV1NewParamsSizeLarge:
		return true
	}
	return false
}
