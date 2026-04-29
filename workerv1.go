// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// WorkerV1Service contains methods and other services that help with interacting
// with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerV1Service] method instead.
type WorkerV1Service struct {
	Options []option.RequestOption
}

// NewWorkerV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerV1Service(opts ...option.RequestOption) (r *WorkerV1Service) {
	r = &WorkerV1Service{}
	r.Options = opts
	return
}

// Creates a Daytona-backed worker runtime. The worker exposes its native runtime
// API through /worker/v1/:id/\* without reshaping payloads or events.
func (r *WorkerV1Service) New(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "worker/v1"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Get worker
func (r *WorkerV1Service) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// End worker
func (r *WorkerV1Service) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Starts or resumes the worker sandbox and OpenCode server. Native
// /worker/v1/:id/\* proxy routes require this lifecycle primitive to have
// completed first.
func (r *WorkerV1Service) Boot(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s/boot", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Forwards a DELETE request to the worker runtime without translating response
// shapes.
func (r *WorkerV1Service) ProxyDelete(ctx context.Context, id string, workerPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workerPath == "" {
		err = errors.New("missing required workerPath parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s/%s", id, workerPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Forwards a GET request to the worker runtime without translating response or SSE
// event shapes.
func (r *WorkerV1Service) ProxyGet(ctx context.Context, id string, workerPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workerPath == "" {
		err = errors.New("missing required workerPath parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s/%s", id, workerPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Forwards a PATCH request to the worker runtime without translating request or
// response shapes.
func (r *WorkerV1Service) ProxyPatch(ctx context.Context, id string, workerPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workerPath == "" {
		err = errors.New("missing required workerPath parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s/%s", id, workerPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return err
}

// Forwards a POST request to the worker runtime without translating request,
// response, or SSE event shapes.
func (r *WorkerV1Service) ProxyPost(ctx context.Context, id string, workerPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workerPath == "" {
		err = errors.New("missing required workerPath parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s/%s", id, workerPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Forwards a PUT request to the worker runtime without translating request or
// response shapes.
func (r *WorkerV1Service) ProxyPut(ctx context.Context, id string, workerPath string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workerPath == "" {
		err = errors.New("missing required workerPath parameter")
		return err
	}
	path := fmt.Sprintf("worker/v1/%s/%s", id, workerPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}
