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

func TestLegalV1FindWithOptionalParams(t *testing.T) {
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
	_, err := client.Legal.V1.Find(context.TODO(), githubcomcasemarkcasedevgo.LegalV1FindParams{
		Query:        githubcomcasemarkcasedevgo.F("xxx"),
		Jurisdiction: githubcomcasemarkcasedevgo.F("jurisdiction"),
		NumResults:   githubcomcasemarkcasedevgo.F(int64(1)),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1GetCitations(t *testing.T) {
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
	_, err := client.Legal.V1.GetCitations(context.TODO(), githubcomcasemarkcasedevgo.LegalV1GetCitationsParams{
		Text: githubcomcasemarkcasedevgo.F("text"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1GetCitationsFromURL(t *testing.T) {
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
	_, err := client.Legal.V1.GetCitationsFromURL(context.TODO(), githubcomcasemarkcasedevgo.LegalV1GetCitationsFromURLParams{
		URL: githubcomcasemarkcasedevgo.F("https://example.com"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1GetFullTextWithOptionalParams(t *testing.T) {
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
	_, err := client.Legal.V1.GetFullText(context.TODO(), githubcomcasemarkcasedevgo.LegalV1GetFullTextParams{
		URL:            githubcomcasemarkcasedevgo.F("https://example.com"),
		HighlightQuery: githubcomcasemarkcasedevgo.F("highlightQuery"),
		MaxCharacters:  githubcomcasemarkcasedevgo.F(int64(1000)),
		SummaryQuery:   githubcomcasemarkcasedevgo.F("summaryQuery"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1ListJurisdictions(t *testing.T) {
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
	_, err := client.Legal.V1.ListJurisdictions(context.TODO(), githubcomcasemarkcasedevgo.LegalV1ListJurisdictionsParams{
		Name: githubcomcasemarkcasedevgo.F("xx"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1PatentSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Legal.V1.PatentSearch(context.TODO(), githubcomcasemarkcasedevgo.LegalV1PatentSearchParams{
		Query:             githubcomcasemarkcasedevgo.F("x"),
		ApplicationStatus: githubcomcasemarkcasedevgo.F("applicationStatus"),
		ApplicationType:   githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LegalV1PatentSearchParamsApplicationTypeUtility),
		Assignee:          githubcomcasemarkcasedevgo.F("assignee"),
		FilingDateFrom:    githubcomcasemarkcasedevgo.F(time.Now()),
		FilingDateTo:      githubcomcasemarkcasedevgo.F(time.Now()),
		GrantDateFrom:     githubcomcasemarkcasedevgo.F(time.Now()),
		GrantDateTo:       githubcomcasemarkcasedevgo.F(time.Now()),
		Inventor:          githubcomcasemarkcasedevgo.F("inventor"),
		Limit:             githubcomcasemarkcasedevgo.F(int64(1)),
		Offset:            githubcomcasemarkcasedevgo.F(int64(0)),
		SortBy:            githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LegalV1PatentSearchParamsSortByFilingDate),
		SortOrder:         githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LegalV1PatentSearchParamsSortOrderAsc),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1ResearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Legal.V1.Research(context.TODO(), githubcomcasemarkcasedevgo.LegalV1ResearchParams{
		Query:             githubcomcasemarkcasedevgo.F("xxx"),
		AdditionalQueries: githubcomcasemarkcasedevgo.F([]string{"string"}),
		Jurisdiction:      githubcomcasemarkcasedevgo.F("jurisdiction"),
		NumResults:        githubcomcasemarkcasedevgo.F(int64(1)),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalV1SimilarWithOptionalParams(t *testing.T) {
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
	_, err := client.Legal.V1.Similar(context.TODO(), githubcomcasemarkcasedevgo.LegalV1SimilarParams{
		URL:                githubcomcasemarkcasedevgo.F("https://example.com"),
		Jurisdiction:       githubcomcasemarkcasedevgo.F("jurisdiction"),
		NumResults:         githubcomcasemarkcasedevgo.F(int64(1)),
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

func TestLegalV1Verify(t *testing.T) {
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
	_, err := client.Legal.V1.Verify(context.TODO(), githubcomcasemarkcasedevgo.LegalV1VerifyParams{
		Text: githubcomcasemarkcasedevgo.F("text"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
