// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/internal/testutil"
	"github.com/CaseMark/casedev-go/option"
)

func TestSearchV1AnswerWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.V1.Answer(context.TODO(), githubcomcasemarkcasedevgo.SearchV1AnswerParams{
		Query:          githubcomcasemarkcasedevgo.F("query"),
		ExcludeDomains: githubcomcasemarkcasedevgo.F([]string{"string"}),
		IncludeDomains: githubcomcasemarkcasedevgo.F([]string{"string"}),
		MaxTokens:      githubcomcasemarkcasedevgo.F(int64(0)),
		Model:          githubcomcasemarkcasedevgo.F("model"),
		NumResults:     githubcomcasemarkcasedevgo.F(int64(1)),
		SearchType:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SearchV1AnswerParamsSearchTypeAuto),
		Stream:         githubcomcasemarkcasedevgo.F(true),
		Temperature:    githubcomcasemarkcasedevgo.F(0.000000),
		Text:           githubcomcasemarkcasedevgo.F(true),
		UseCustomLlm:   githubcomcasemarkcasedevgo.F(true),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchV1ContentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.V1.Contents(context.TODO(), githubcomcasemarkcasedevgo.SearchV1ContentsParams{
		URLs:             githubcomcasemarkcasedevgo.F([]string{"https://example.com"}),
		Context:          githubcomcasemarkcasedevgo.F("context"),
		Extras:           githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
		Highlights:       githubcomcasemarkcasedevgo.F(true),
		Livecrawl:        githubcomcasemarkcasedevgo.F(true),
		LivecrawlTimeout: githubcomcasemarkcasedevgo.F(int64(0)),
		Subpages:         githubcomcasemarkcasedevgo.F(true),
		SubpageTarget:    githubcomcasemarkcasedevgo.F(int64(0)),
		Summary:          githubcomcasemarkcasedevgo.F(true),
		Text:             githubcomcasemarkcasedevgo.F(true),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchV1ResearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.V1.Research(context.TODO(), githubcomcasemarkcasedevgo.SearchV1ResearchParams{
		Instructions: githubcomcasemarkcasedevgo.F("instructions"),
		Model:        githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SearchV1ResearchParamsModelFast),
		OutputSchema: githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
		Query:        githubcomcasemarkcasedevgo.F("query"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchV1GetResearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.V1.GetResearch(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.SearchV1GetResearchParams{
			Events: githubcomcasemarkcasedevgo.F("events"),
			Stream: githubcomcasemarkcasedevgo.F(true),
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

func TestSearchV1SearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.V1.Search(context.TODO(), githubcomcasemarkcasedevgo.SearchV1SearchParams{
		Query:              githubcomcasemarkcasedevgo.F("query"),
		AdditionalQueries:  githubcomcasemarkcasedevgo.F([]string{"string"}),
		Category:           githubcomcasemarkcasedevgo.F("category"),
		Contents:           githubcomcasemarkcasedevgo.F("contents"),
		EndCrawlDate:       githubcomcasemarkcasedevgo.F(time.Now()),
		EndPublishedDate:   githubcomcasemarkcasedevgo.F(time.Now()),
		ExcludeDomains:     githubcomcasemarkcasedevgo.F([]string{"string"}),
		IncludeDomains:     githubcomcasemarkcasedevgo.F([]string{"string"}),
		IncludeText:        githubcomcasemarkcasedevgo.F(true),
		NumResults:         githubcomcasemarkcasedevgo.F(int64(1)),
		StartCrawlDate:     githubcomcasemarkcasedevgo.F(time.Now()),
		StartPublishedDate: githubcomcasemarkcasedevgo.F(time.Now()),
		Type:               githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SearchV1SearchParamsTypeAuto),
		UserLocation:       githubcomcasemarkcasedevgo.F("userLocation"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchV1SimilarWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.V1.Similar(context.TODO(), githubcomcasemarkcasedevgo.SearchV1SimilarParams{
		URL:                githubcomcasemarkcasedevgo.F("https://example.com"),
		Contents:           githubcomcasemarkcasedevgo.F("contents"),
		EndCrawlDate:       githubcomcasemarkcasedevgo.F(time.Now()),
		EndPublishedDate:   githubcomcasemarkcasedevgo.F(time.Now()),
		ExcludeDomains:     githubcomcasemarkcasedevgo.F([]string{"string"}),
		IncludeDomains:     githubcomcasemarkcasedevgo.F([]string{"string"}),
		IncludeText:        githubcomcasemarkcasedevgo.F(true),
		NumResults:         githubcomcasemarkcasedevgo.F(int64(1)),
		StartCrawlDate:     githubcomcasemarkcasedevgo.F(time.Now()),
		StartPublishedDate: githubcomcasemarkcasedevgo.F(time.Now()),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
