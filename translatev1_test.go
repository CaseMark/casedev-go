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

func TestTranslateV1Detect(t *testing.T) {
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
	_, err := client.Translate.V1.Detect(context.TODO(), githubcomcasemarkcasedevgo.TranslateV1DetectParams{
		Q: githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.TranslateV1DetectParamsQUnion](shared.UnionString("string")),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTranslateV1ListLanguagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Translate.V1.ListLanguages(context.TODO(), githubcomcasemarkcasedevgo.TranslateV1ListLanguagesParams{
		Model:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.TranslateV1ListLanguagesParamsModelNmt),
		Target: githubcomcasemarkcasedevgo.F("target"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTranslateV1TranslateWithOptionalParams(t *testing.T) {
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
	_, err := client.Translate.V1.Translate(context.TODO(), githubcomcasemarkcasedevgo.TranslateV1TranslateParams{
		Q:      githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.TranslateV1TranslateParamsQUnion](shared.UnionString("string")),
		Target: githubcomcasemarkcasedevgo.F("es"),
		Format: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.TranslateV1TranslateParamsFormatText),
		Model:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.TranslateV1TranslateParamsModelNmt),
		Source: githubcomcasemarkcasedevgo.F("en"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
