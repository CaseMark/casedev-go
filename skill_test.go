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

func TestSkillNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Skills.New(context.TODO(), githubcomcasemarkcasedevgo.SkillNewParams{
		Content: githubcomcasemarkcasedevgo.F("x"),
		Name:    githubcomcasemarkcasedevgo.F("x"),
		Files: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.SkillNewParamsFile{{
			Content:     githubcomcasemarkcasedevgo.F("content"),
			Path:        githubcomcasemarkcasedevgo.F("path"),
			ContentType: githubcomcasemarkcasedevgo.F("contentType"),
			Metadata:    githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
			Name:        githubcomcasemarkcasedevgo.F("name"),
			Summary:     githubcomcasemarkcasedevgo.F("summary"),
			Tags:        githubcomcasemarkcasedevgo.F([]string{"string"}),
		}}),
		Metadata: githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
		Slug:     githubcomcasemarkcasedevgo.F("slug"),
		Summary:  githubcomcasemarkcasedevgo.F("summary"),
		Tags:     githubcomcasemarkcasedevgo.F([]string{"string"}),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkillUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Skills.Update(
		context.TODO(),
		"slug",
		githubcomcasemarkcasedevgo.SkillUpdateParams{
			Content: githubcomcasemarkcasedevgo.F("content"),
			Files: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.SkillUpdateParamsFile{{
				Content:     githubcomcasemarkcasedevgo.F("content"),
				Path:        githubcomcasemarkcasedevgo.F("path"),
				ContentType: githubcomcasemarkcasedevgo.F("contentType"),
				Metadata:    githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
				Name:        githubcomcasemarkcasedevgo.F("name"),
				Summary:     githubcomcasemarkcasedevgo.F("summary"),
				Tags:        githubcomcasemarkcasedevgo.F([]string{"string"}),
			}}),
			Metadata: githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
			Name:     githubcomcasemarkcasedevgo.F("name"),
			Slug:     githubcomcasemarkcasedevgo.F("slug"),
			Summary:  githubcomcasemarkcasedevgo.F("summary"),
			Tags:     githubcomcasemarkcasedevgo.F([]string{"string"}),
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

func TestSkillDelete(t *testing.T) {
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
	_, err := client.Skills.Delete(context.TODO(), "slug")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkillExportWithOptionalParams(t *testing.T) {
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
	_, err := client.Skills.Export(
		context.TODO(),
		"slug",
		githubcomcasemarkcasedevgo.SkillExportParams{
			Target: githubcomcasemarkcasedevgo.F("target"),
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

func TestSkillRead(t *testing.T) {
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
	_, err := client.Skills.Read(context.TODO(), "slug")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkillResolveWithOptionalParams(t *testing.T) {
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
	_, err := client.Skills.Resolve(context.TODO(), githubcomcasemarkcasedevgo.SkillResolveParams{
		Q:     githubcomcasemarkcasedevgo.F("q"),
		Limit: githubcomcasemarkcasedevgo.F(int64(1)),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
