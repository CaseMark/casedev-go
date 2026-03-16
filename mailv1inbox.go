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
// MailV1InboxService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailV1InboxService] method instead.
type MailV1InboxService struct {
	Options []option.RequestOption
}

// NewMailV1InboxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMailV1InboxService(opts ...option.RequestOption) (r *MailV1InboxService) {
	r = &MailV1InboxService{}
	r.Options = opts
	return
}

// Create an inbox owned by the authenticated organization.
func (r *MailV1InboxService) New(ctx context.Context, body MailV1InboxNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "mail/v1/inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Get an inbox owned by the authenticated organization.
func (r *MailV1InboxService) Get(ctx context.Context, inboxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// List inboxes owned by the authenticated organization.
func (r *MailV1InboxService) List(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "mail/v1/inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Delete an inbox owned by the authenticated organization.
func (r *MailV1InboxService) Delete(ctx context.Context, inboxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get attachment metadata for a message in an inbox owned by the authenticated
// organization.
func (r *MailV1InboxService) GetAttachment(ctx context.Context, inboxID string, messageID string, attachmentID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return err
	}
	if attachmentID == "" {
		err = errors.New("missing required attachmentId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s/messages/%s/attachments/%s", inboxID, messageID, attachmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Get a message for an inbox owned by the authenticated organization.
func (r *MailV1InboxService) GetMessage(ctx context.Context, inboxID string, messageID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s/messages/%s", inboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// List messages for an inbox owned by the authenticated organization.
func (r *MailV1InboxService) ListMessages(ctx context.Context, inboxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s/messages", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Reply to a message in an inbox owned by the authenticated organization.
func (r *MailV1InboxService) Reply(ctx context.Context, inboxID string, messageID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s/messages/%s/reply", inboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Send a message from an inbox owned by the authenticated organization.
func (r *MailV1InboxService) Send(ctx context.Context, inboxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inboxId parameter")
		return err
	}
	path := fmt.Sprintf("mail/v1/inboxes/%s/messages/send", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type MailV1InboxNewParams struct {
	Address     param.Field[string] `json:"address"`
	DisplayName param.Field[string] `json:"displayName"`
}

func (r MailV1InboxNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
