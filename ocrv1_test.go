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

func TestOcrV1Get(t *testing.T) {
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
	_, err := client.Ocr.V1.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOcrV1Download(t *testing.T) {
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
	_, err := client.Ocr.V1.Download(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.OcrV1DownloadParamsTypeText,
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOcrV1ProcessWithOptionalParams(t *testing.T) {
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
	_, err := client.Ocr.V1.Process(context.TODO(), githubcomcasemarkcasedevgo.OcrV1ProcessParams{
		DocumentURL: githubcomcasemarkcasedevgo.F("https://example.com/contract.pdf"),
		CallbackURL: githubcomcasemarkcasedevgo.F("https://your-app.com/webhooks/ocr-complete"),
		DocumentID:  githubcomcasemarkcasedevgo.F("contract-2024-001"),
		Engine:      githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.OcrV1ProcessParamsEngineDoctr),
		Features: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.OcrV1ProcessParamsFeatures{
			Embed: githubcomcasemarkcasedevgo.F(map[string]interface{}{}),
			Forms: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Tables: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.OcrV1ProcessParamsFeaturesTables{
				Format: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.OcrV1ProcessParamsFeaturesTablesFormatCsv),
			}),
		}),
		ResultBucket: githubcomcasemarkcasedevgo.F("my-ocr-results"),
		ResultPrefix: githubcomcasemarkcasedevgo.F("ocr/2024/"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
