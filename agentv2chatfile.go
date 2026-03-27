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

// Create, manage, and execute AI agents with tool access, sandbox environments,
// and async run workflows
//
// AgentV2ChatFileService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV2ChatFileService] method instead.
type AgentV2ChatFileService struct {
	Options []option.RequestOption
}

// NewAgentV2ChatFileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV2ChatFileService(opts ...option.RequestOption) (r *AgentV2ChatFileService) {
	r = &AgentV2ChatFileService{}
	r.Options = opts
	return
}

// Lists files created by the agent in the Daytona runtime workspace. Stopped or
// archived runtimes are transparently resumed or recovered.
func (r *AgentV2ChatFileService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV2ChatFileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/chat/%s/files", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Downloads a file from the Daytona runtime workspace by path. Stopped or archived
// runtimes are transparently resumed or recovered.
func (r *AgentV2ChatFileService) Download(ctx context.Context, id string, filePath string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v2/chat/%s/files/%s", id, filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentV2ChatFileListResponse struct {
	ChatID string                            `json:"chatId"`
	Files  []AgentV2ChatFileListResponseFile `json:"files"`
	JSON   agentV2ChatFileListResponseJSON   `json:"-"`
}

// agentV2ChatFileListResponseJSON contains the JSON metadata for the struct
// [AgentV2ChatFileListResponse]
type agentV2ChatFileListResponseJSON struct {
	ChatID      apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV2ChatFileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ChatFileListResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV2ChatFileListResponseFile struct {
	Name string `json:"name"`
	// Relative path from /workspace
	Path      string                              `json:"path"`
	SizeBytes int64                               `json:"sizeBytes"`
	JSON      agentV2ChatFileListResponseFileJSON `json:"-"`
}

// agentV2ChatFileListResponseFileJSON contains the JSON metadata for the struct
// [AgentV2ChatFileListResponseFile]
type agentV2ChatFileListResponseFileJSON struct {
	Name        apijson.Field
	Path        apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV2ChatFileListResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV2ChatFileListResponseFileJSON) RawJSON() string {
	return r.raw
}
