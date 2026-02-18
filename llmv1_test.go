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

func TestLlmV1NewEmbeddingWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Llm.V1.NewEmbedding(context.TODO(), githubcomcasemarkcasedevgo.LlmV1NewEmbeddingParams{
		Input:          githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.LlmV1NewEmbeddingParamsInputUnion](shared.UnionString("string")),
		Model:          githubcomcasemarkcasedevgo.F("model"),
		Dimensions:     githubcomcasemarkcasedevgo.F(int64(0)),
		EncodingFormat: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LlmV1NewEmbeddingParamsEncodingFormatFloat),
		User:           githubcomcasemarkcasedevgo.F("user"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmV1ListModels(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Llm.V1.ListModels(context.TODO())
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
