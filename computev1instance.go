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

// ComputeV1InstanceService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeV1InstanceService] method instead.
type ComputeV1InstanceService struct {
	Options []option.RequestOption
}

// NewComputeV1InstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeV1InstanceService(opts ...option.RequestOption) (r *ComputeV1InstanceService) {
	r = &ComputeV1InstanceService{}
	r.Options = opts
	return
}

// Launches a new GPU compute instance with automatic SSH key generation. Supports
// mounting Case.dev Vaults as filesystems and configurable auto-shutdown. Instance
// boots in ~2-5 minutes. Perfect for batch OCR processing, AI model training, and
// intensive document analysis workloads.
func (r *ComputeV1InstanceService) New(ctx context.Context, body ComputeV1InstanceNewParams, opts ...option.RequestOption) (res *ComputeV1InstanceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information about a GPU instance including SSH connection
// details, vault mount scripts, real-time cost tracking, and current status. SSH
// private key included for secure access.
func (r *ComputeV1InstanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputeV1InstanceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all GPU compute instances for your organization with real-time status
// updates from Lambda Labs. Includes pricing, runtime metrics, and auto-shutdown
// configuration. Perfect for monitoring AI workloads, document processing jobs,
// and cost tracking.
func (r *ComputeV1InstanceService) List(ctx context.Context, opts ...option.RequestOption) (res *ComputeV1InstanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Terminates a running GPU instance, calculates final cost, and cleans up SSH
// keys. This action is permanent and cannot be undone. All data on the instance
// will be lost.
func (r *ComputeV1InstanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputeV1InstanceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("compute/v1/instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ComputeV1InstanceNewResponse struct {
	ID                  string                           `json:"id"`
	AutoShutdownMinutes int64                            `json:"autoShutdownMinutes" api:"nullable"`
	CreatedAt           string                           `json:"createdAt"`
	GPU                 string                           `json:"gpu"`
	InstanceType        string                           `json:"instanceType"`
	Message             string                           `json:"message"`
	Name                string                           `json:"name"`
	PricePerHour        string                           `json:"pricePerHour"`
	Region              string                           `json:"region"`
	Specs               interface{}                      `json:"specs"`
	Status              string                           `json:"status"`
	Vaults              []interface{}                    `json:"vaults"`
	JSON                computeV1InstanceNewResponseJSON `json:"-"`
}

// computeV1InstanceNewResponseJSON contains the JSON metadata for the struct
// [ComputeV1InstanceNewResponse]
type computeV1InstanceNewResponseJSON struct {
	ID                  apijson.Field
	AutoShutdownMinutes apijson.Field
	CreatedAt           apijson.Field
	GPU                 apijson.Field
	InstanceType        apijson.Field
	Message             apijson.Field
	Name                apijson.Field
	PricePerHour        apijson.Field
	Region              apijson.Field
	Specs               apijson.Field
	Status              apijson.Field
	Vaults              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ComputeV1InstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceNewResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceGetResponse struct {
	ID                    string                           `json:"id"`
	AutoShutdownMinutes   int64                            `json:"autoShutdownMinutes" api:"nullable"`
	CreatedAt             string                           `json:"createdAt"`
	CurrentCost           string                           `json:"currentCost"`
	CurrentRuntimeSeconds int64                            `json:"currentRuntimeSeconds"`
	GPU                   string                           `json:"gpu"`
	InstanceType          string                           `json:"instanceType"`
	IP                    string                           `json:"ip" api:"nullable"`
	Name                  string                           `json:"name"`
	PricePerHour          string                           `json:"pricePerHour"`
	Region                string                           `json:"region"`
	Specs                 interface{}                      `json:"specs"`
	SSH                   ComputeV1InstanceGetResponseSSH  `json:"ssh" api:"nullable"`
	StartedAt             string                           `json:"startedAt" api:"nullable"`
	Status                string                           `json:"status"`
	VaultMounts           interface{}                      `json:"vaultMounts" api:"nullable"`
	JSON                  computeV1InstanceGetResponseJSON `json:"-"`
}

// computeV1InstanceGetResponseJSON contains the JSON metadata for the struct
// [ComputeV1InstanceGetResponse]
type computeV1InstanceGetResponseJSON struct {
	ID                    apijson.Field
	AutoShutdownMinutes   apijson.Field
	CreatedAt             apijson.Field
	CurrentCost           apijson.Field
	CurrentRuntimeSeconds apijson.Field
	GPU                   apijson.Field
	InstanceType          apijson.Field
	IP                    apijson.Field
	Name                  apijson.Field
	PricePerHour          apijson.Field
	Region                apijson.Field
	Specs                 apijson.Field
	SSH                   apijson.Field
	StartedAt             apijson.Field
	Status                apijson.Field
	VaultMounts           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ComputeV1InstanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceGetResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceGetResponseSSH struct {
	Command      string                              `json:"command"`
	Host         string                              `json:"host"`
	Instructions []interface{}                       `json:"instructions"`
	PrivateKey   string                              `json:"privateKey"`
	User         string                              `json:"user"`
	JSON         computeV1InstanceGetResponseSSHJSON `json:"-"`
}

// computeV1InstanceGetResponseSSHJSON contains the JSON metadata for the struct
// [ComputeV1InstanceGetResponseSSH]
type computeV1InstanceGetResponseSSHJSON struct {
	Command      apijson.Field
	Host         apijson.Field
	Instructions apijson.Field
	PrivateKey   apijson.Field
	User         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ComputeV1InstanceGetResponseSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceGetResponseSSHJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceListResponse struct {
	Count     int64                                   `json:"count"`
	Instances []ComputeV1InstanceListResponseInstance `json:"instances"`
	JSON      computeV1InstanceListResponseJSON       `json:"-"`
}

// computeV1InstanceListResponseJSON contains the JSON metadata for the struct
// [ComputeV1InstanceListResponse]
type computeV1InstanceListResponseJSON struct {
	Count       apijson.Field
	Instances   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComputeV1InstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceListResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceListResponseInstance struct {
	ID                  string                                       `json:"id"`
	AutoShutdownMinutes int64                                        `json:"autoShutdownMinutes" api:"nullable"`
	CreatedAt           time.Time                                    `json:"createdAt" format:"date-time"`
	GPU                 string                                       `json:"gpu"`
	InstanceType        string                                       `json:"instanceType"`
	IP                  string                                       `json:"ip" api:"nullable"`
	Name                string                                       `json:"name"`
	PricePerHour        string                                       `json:"pricePerHour"`
	Region              string                                       `json:"region"`
	StartedAt           time.Time                                    `json:"startedAt" api:"nullable" format:"date-time"`
	Status              ComputeV1InstanceListResponseInstancesStatus `json:"status"`
	TotalCost           string                                       `json:"totalCost"`
	TotalRuntimeSeconds int64                                        `json:"totalRuntimeSeconds"`
	JSON                computeV1InstanceListResponseInstanceJSON    `json:"-"`
}

// computeV1InstanceListResponseInstanceJSON contains the JSON metadata for the
// struct [ComputeV1InstanceListResponseInstance]
type computeV1InstanceListResponseInstanceJSON struct {
	ID                  apijson.Field
	AutoShutdownMinutes apijson.Field
	CreatedAt           apijson.Field
	GPU                 apijson.Field
	InstanceType        apijson.Field
	IP                  apijson.Field
	Name                apijson.Field
	PricePerHour        apijson.Field
	Region              apijson.Field
	StartedAt           apijson.Field
	Status              apijson.Field
	TotalCost           apijson.Field
	TotalRuntimeSeconds apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ComputeV1InstanceListResponseInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceListResponseInstanceJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceListResponseInstancesStatus string

const (
	ComputeV1InstanceListResponseInstancesStatusBooting    ComputeV1InstanceListResponseInstancesStatus = "booting"
	ComputeV1InstanceListResponseInstancesStatusRunning    ComputeV1InstanceListResponseInstancesStatus = "running"
	ComputeV1InstanceListResponseInstancesStatusStopping   ComputeV1InstanceListResponseInstancesStatus = "stopping"
	ComputeV1InstanceListResponseInstancesStatusStopped    ComputeV1InstanceListResponseInstancesStatus = "stopped"
	ComputeV1InstanceListResponseInstancesStatusTerminated ComputeV1InstanceListResponseInstancesStatus = "terminated"
	ComputeV1InstanceListResponseInstancesStatusFailed     ComputeV1InstanceListResponseInstancesStatus = "failed"
)

func (r ComputeV1InstanceListResponseInstancesStatus) IsKnown() bool {
	switch r {
	case ComputeV1InstanceListResponseInstancesStatusBooting, ComputeV1InstanceListResponseInstancesStatusRunning, ComputeV1InstanceListResponseInstancesStatusStopping, ComputeV1InstanceListResponseInstancesStatusStopped, ComputeV1InstanceListResponseInstancesStatusTerminated, ComputeV1InstanceListResponseInstancesStatusFailed:
		return true
	}
	return false
}

type ComputeV1InstanceDeleteResponse struct {
	ID                  string                              `json:"id"`
	Message             string                              `json:"message"`
	Name                string                              `json:"name"`
	Status              string                              `json:"status"`
	TotalCost           string                              `json:"totalCost"`
	TotalRuntimeSeconds int64                               `json:"totalRuntimeSeconds"`
	JSON                computeV1InstanceDeleteResponseJSON `json:"-"`
}

// computeV1InstanceDeleteResponseJSON contains the JSON metadata for the struct
// [ComputeV1InstanceDeleteResponse]
type computeV1InstanceDeleteResponseJSON struct {
	ID                  apijson.Field
	Message             apijson.Field
	Name                apijson.Field
	Status              apijson.Field
	TotalCost           apijson.Field
	TotalRuntimeSeconds apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ComputeV1InstanceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r computeV1InstanceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ComputeV1InstanceNewParams struct {
	// GPU type (e.g., 'gpu_1x_h100_sxm5')
	InstanceType param.Field[string] `json:"instanceType" api:"required"`
	// Instance name
	Name param.Field[string] `json:"name" api:"required"`
	// Region (e.g., 'us-west-1')
	Region param.Field[string] `json:"region" api:"required"`
	// Auto-shutdown timer (null = never)
	AutoShutdownMinutes param.Field[int64] `json:"autoShutdownMinutes"`
	// Vault IDs to mount
	VaultIDs param.Field[[]string] `json:"vaultIds"`
}

func (r ComputeV1InstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
