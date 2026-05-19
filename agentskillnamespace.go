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

// Create, manage, and execute AI agents with tool access, sandbox environments,
// and async run workflows
//
// AgentSkillNamespaceService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentSkillNamespaceService] method instead.
type AgentSkillNamespaceService struct {
	Options []option.RequestOption
}

// NewAgentSkillNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentSkillNamespaceService(opts ...option.RequestOption) (r *AgentSkillNamespaceService) {
	r = &AgentSkillNamespaceService{}
	r.Options = opts
	return
}

// Create a private skill namespace owned by the authenticated org and receive a
// one-time bearer token used by the case-skills publisher.
func (r *AgentSkillNamespaceService) New(ctx context.Context, body AgentSkillNamespaceNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "agent/skills/namespaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Read skill namespace
func (r *AgentSkillNamespaceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/skills/namespaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// List all active skill namespaces owned by the authenticated organization.
func (r *AgentSkillNamespaceService) List(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "agent/skills/namespaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Delete skill namespace
func (r *AgentSkillNamespaceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/skills/namespaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Upload a tree of skill files for the namespace. Authenticated by the namespace
// bearer token. Atomic at the version-bump level: a partial upload leaves the
// namespace pinned to the previous version.
func (r *AgentSkillNamespaceService) Publish(ctx context.Context, id string, body AgentSkillNamespacePublishParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/skills/namespaces/%s/publish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Returns the active version's file manifest with short-lived presigned S3 URLs.
// Sandboxes use this to materialize the tree at /workspace/.agents/skills/ before
// opencode boots.
func (r *AgentSkillNamespaceService) Pull(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/skills/namespaces/%s/pull", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Rotate skill namespace token
func (r *AgentSkillNamespaceService) RotateToken(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("agent/skills/namespaces/%s/rotate-token", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type AgentSkillNamespaceNewParams struct {
	// URL-safe slug, e.g. "curi" or "client-firm-abc". Lowercase alphanumeric with
	// single hyphens, 2-64 chars.
	NamespaceID param.Field[string]      `json:"namespaceId" api:"required"`
	Description param.Field[string]      `json:"description"`
	Label       param.Field[string]      `json:"label"`
	Metadata    param.Field[interface{}] `json:"metadata"`
}

func (r AgentSkillNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentSkillNamespacePublishParams struct {
	Files param.Field[[]AgentSkillNamespacePublishParamsFile] `json:"files" api:"required"`
}

func (r AgentSkillNamespacePublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentSkillNamespacePublishParamsFile struct {
	Content     param.Field[string]                                        `json:"content" api:"required"`
	Encoding    param.Field[AgentSkillNamespacePublishParamsFilesEncoding] `json:"encoding" api:"required"`
	Path        param.Field[string]                                        `json:"path" api:"required"`
	ContentType param.Field[string]                                        `json:"contentType"`
}

func (r AgentSkillNamespacePublishParamsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentSkillNamespacePublishParamsFilesEncoding string

const (
	AgentSkillNamespacePublishParamsFilesEncodingUtf8   AgentSkillNamespacePublishParamsFilesEncoding = "utf8"
	AgentSkillNamespacePublishParamsFilesEncodingBase64 AgentSkillNamespacePublishParamsFilesEncoding = "base64"
)

func (r AgentSkillNamespacePublishParamsFilesEncoding) IsKnown() bool {
	switch r {
	case AgentSkillNamespacePublishParamsFilesEncodingUtf8, AgentSkillNamespacePublishParamsFilesEncodingBase64:
		return true
	}
	return false
}
