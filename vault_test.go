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
	"github.com/CaseMark/casedev-go/shared"
)

func TestVaultNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.New(context.TODO(), githubcomcasemarkcasedevgo.VaultNewParams{
		Name:           githubcomcasemarkcasedevgo.F("Contract Review Archive"),
		Description:    githubcomcasemarkcasedevgo.F("Repository for all client contract reviews and analysis"),
		EmbeddingModel: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultNewParamsEmbeddingModelCasemarkEmbedV1),
		EnableIndexing: githubcomcasemarkcasedevgo.F(true),
		GroupID:        githubcomcasemarkcasedevgo.F("grp_abc123"),
		Metadata: githubcomcasemarkcasedevgo.F[any](map[string]interface{}{
			"containsPHI":    true,
			"hipaaCompliant": true,
		}),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVaultGet(t *testing.T) {
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
	_, err := client.Vault.Get(context.TODO(), "vault_abc123")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVaultUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.Update(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultUpdateParams{
			Description: githubcomcasemarkcasedevgo.F("description"),
			GroupID:     githubcomcasemarkcasedevgo.F("groupId"),
			Name:        githubcomcasemarkcasedevgo.F("Updated Vault Name"),
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

func TestVaultListWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.List(context.TODO(), githubcomcasemarkcasedevgo.VaultListParams{
		Cursor:        githubcomcasemarkcasedevgo.F("cursor"),
		IncludeTotals: githubcomcasemarkcasedevgo.F(true),
		Limit:         githubcomcasemarkcasedevgo.F(int64(1)),
		Query:         githubcomcasemarkcasedevgo.F("query"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVaultDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.Delete(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultDeleteParams{
			Async: githubcomcasemarkcasedevgo.F(true),
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

func TestVaultConfirmUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.ConfirmUpload(
		context.TODO(),
		"id",
		"objectId",
		githubcomcasemarkcasedevgo.VaultConfirmUploadParams{
			Success:      githubcomcasemarkcasedevgo.F(true),
			AutoIngest:   githubcomcasemarkcasedevgo.F(true),
			ErrorCode:    githubcomcasemarkcasedevgo.F("errorCode"),
			ErrorMessage: githubcomcasemarkcasedevgo.F("errorMessage"),
			Etag:         githubcomcasemarkcasedevgo.F("etag"),
			SizeBytes:    githubcomcasemarkcasedevgo.F(int64(0)),
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

func TestVaultIngestWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.Ingest(
		context.TODO(),
		"id",
		"objectId",
		githubcomcasemarkcasedevgo.VaultIngestParams{
			CallbackURL:    githubcomcasemarkcasedevgo.F("https://example.com"),
			PageBoundaries: githubcomcasemarkcasedevgo.F([]int64{int64(2)}),
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

func TestVaultSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.Search(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultSearchParams{
			Query: githubcomcasemarkcasedevgo.F("query"),
			Filters: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultSearchParamsFilters{
				ObjectID: githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.VaultSearchParamsFiltersObjectIDUnion](shared.UnionString("string")),
				PageRange: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultSearchParamsFiltersPageRange{
					Start: githubcomcasemarkcasedevgo.F(int64(1)),
					End:   githubcomcasemarkcasedevgo.F(int64(1)),
				}),
			}),
			Method: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultSearchParamsMethodHybrid),
			TopK:   githubcomcasemarkcasedevgo.F(int64(1)),
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

func TestVaultUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.Upload(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultUploadParams{
			ContentType: githubcomcasemarkcasedevgo.F("contentType"),
			Filename:    githubcomcasemarkcasedevgo.F("filename"),
			AutoIndex:   githubcomcasemarkcasedevgo.F(true),
			FileOrigin: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			IsAIGenerated:  githubcomcasemarkcasedevgo.F(true),
			Metadata:       githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
			Path:           githubcomcasemarkcasedevgo.F("path"),
			SizeBytes:      githubcomcasemarkcasedevgo.F(int64(0)),
			IdempotencyKey: githubcomcasemarkcasedevgo.F("Idempotency-Key"),
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
