// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/internal/testutil"
	"github.com/CaseMark/casedev-go/option"
)

func TestVaultMultipartAbort(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Vault.Multipart.Abort(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultMultipartAbortParams{
			ObjectID: githubcomcasemarkcasedevgo.F("objectId"),
			UploadID: githubcomcasemarkcasedevgo.F("uploadId"),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVaultMultipartGetPartURLs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Vault.Multipart.GetPartURLs(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultMultipartGetPartURLsParams{
			ObjectID: githubcomcasemarkcasedevgo.F("objectId"),
			Parts: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.VaultMultipartGetPartURLsParamsPart{{
				PartNumber: githubcomcasemarkcasedevgo.F(int64(1)),
				SizeBytes:  githubcomcasemarkcasedevgo.F(int64(1)),
			}}),
			UploadID: githubcomcasemarkcasedevgo.F("uploadId"),
		},
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
