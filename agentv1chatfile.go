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
// AgentV1ChatFileService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentV1ChatFileService] method instead.
type AgentV1ChatFileService struct {
	Options []option.RequestOption
}

// NewAgentV1ChatFileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentV1ChatFileService(opts ...option.RequestOption) (r *AgentV1ChatFileService) {
	r = &AgentV1ChatFileService{}
	r.Options = opts
	return
}

// Lists files created by the agent in the sandbox workspace. Only available while
// the sandbox is running.
func (r *AgentV1ChatFileService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *AgentV1ChatFileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v1/chat/%s/files", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Downloads a file from the sandbox workspace by path. Only available while the
// sandbox is running.
func (r *AgentV1ChatFileService) Download(ctx context.Context, id string, path string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if path == "" {
		err = errors.New("missing required path parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent/v1/chat/%s/files/%s", id, path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentV1ChatFileListResponse struct {
	ChatID string                            `json:"chatId"`
	Files  []AgentV1ChatFileListResponseFile `json:"files"`
	JSON   agentV1ChatFileListResponseJSON   `json:"-"`
}

// agentV1ChatFileListResponseJSON contains the JSON metadata for the struct
// [AgentV1ChatFileListResponse]
type agentV1ChatFileListResponseJSON struct {
	ChatID      apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1ChatFileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1ChatFileListResponseJSON) RawJSON() string {
	return r.raw
}

type AgentV1ChatFileListResponseFile struct {
	Name string `json:"name"`
	// Relative path from /workspace
	Path      string                              `json:"path"`
	SizeBytes int64                               `json:"sizeBytes"`
	JSON      agentV1ChatFileListResponseFileJSON `json:"-"`
}

// agentV1ChatFileListResponseFileJSON contains the JSON metadata for the struct
// [AgentV1ChatFileListResponseFile]
type agentV1ChatFileListResponseFileJSON struct {
	Name        apijson.Field
	Path        apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentV1ChatFileListResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentV1ChatFileListResponseFileJSON) RawJSON() string {
	return r.raw
}
