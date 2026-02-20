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

func TestMemoryV1NewWithOptionalParams(t *testing.T) {
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
	_, err := client.Memory.V1.New(context.TODO(), githubcomcasemarkcasedevgo.MemoryV1NewParams{
		Messages: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.MemoryV1NewParamsMessage{{
			Content: githubcomcasemarkcasedevgo.F("content"),
			Role:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MemoryV1NewParamsMessagesRoleUser),
		}}),
		Category:         githubcomcasemarkcasedevgo.F("category"),
		ExtractionPrompt: githubcomcasemarkcasedevgo.F("extraction_prompt"),
		Infer:            githubcomcasemarkcasedevgo.F(true),
		Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Tag1:  githubcomcasemarkcasedevgo.F("tag_1"),
		Tag10: githubcomcasemarkcasedevgo.F("tag_10"),
		Tag11: githubcomcasemarkcasedevgo.F("tag_11"),
		Tag12: githubcomcasemarkcasedevgo.F("tag_12"),
		Tag2:  githubcomcasemarkcasedevgo.F("tag_2"),
		Tag3:  githubcomcasemarkcasedevgo.F("tag_3"),
		Tag4:  githubcomcasemarkcasedevgo.F("tag_4"),
		Tag5:  githubcomcasemarkcasedevgo.F("tag_5"),
		Tag6:  githubcomcasemarkcasedevgo.F("tag_6"),
		Tag7:  githubcomcasemarkcasedevgo.F("tag_7"),
		Tag8:  githubcomcasemarkcasedevgo.F("tag_8"),
		Tag9:  githubcomcasemarkcasedevgo.F("tag_9"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryV1Get(t *testing.T) {
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
	_, err := client.Memory.V1.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryV1ListWithOptionalParams(t *testing.T) {
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
	_, err := client.Memory.V1.List(context.TODO(), githubcomcasemarkcasedevgo.MemoryV1ListParams{
		Category: githubcomcasemarkcasedevgo.F("category"),
		Limit:    githubcomcasemarkcasedevgo.F(int64(0)),
		Offset:   githubcomcasemarkcasedevgo.F(int64(0)),
		Tag1:     githubcomcasemarkcasedevgo.F("tag_1"),
		Tag10:    githubcomcasemarkcasedevgo.F("tag_10"),
		Tag11:    githubcomcasemarkcasedevgo.F("tag_11"),
		Tag12:    githubcomcasemarkcasedevgo.F("tag_12"),
		Tag2:     githubcomcasemarkcasedevgo.F("tag_2"),
		Tag3:     githubcomcasemarkcasedevgo.F("tag_3"),
		Tag4:     githubcomcasemarkcasedevgo.F("tag_4"),
		Tag5:     githubcomcasemarkcasedevgo.F("tag_5"),
		Tag6:     githubcomcasemarkcasedevgo.F("tag_6"),
		Tag7:     githubcomcasemarkcasedevgo.F("tag_7"),
		Tag8:     githubcomcasemarkcasedevgo.F("tag_8"),
		Tag9:     githubcomcasemarkcasedevgo.F("tag_9"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryV1Delete(t *testing.T) {
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
	_, err := client.Memory.V1.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryV1DeleteAllWithOptionalParams(t *testing.T) {
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
	_, err := client.Memory.V1.DeleteAll(context.TODO(), githubcomcasemarkcasedevgo.MemoryV1DeleteAllParams{
		Tag1:  githubcomcasemarkcasedevgo.F("tag_1"),
		Tag10: githubcomcasemarkcasedevgo.F("tag_10"),
		Tag11: githubcomcasemarkcasedevgo.F("tag_11"),
		Tag12: githubcomcasemarkcasedevgo.F("tag_12"),
		Tag2:  githubcomcasemarkcasedevgo.F("tag_2"),
		Tag3:  githubcomcasemarkcasedevgo.F("tag_3"),
		Tag4:  githubcomcasemarkcasedevgo.F("tag_4"),
		Tag5:  githubcomcasemarkcasedevgo.F("tag_5"),
		Tag6:  githubcomcasemarkcasedevgo.F("tag_6"),
		Tag7:  githubcomcasemarkcasedevgo.F("tag_7"),
		Tag8:  githubcomcasemarkcasedevgo.F("tag_8"),
		Tag9:  githubcomcasemarkcasedevgo.F("tag_9"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryV1SearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Memory.V1.Search(context.TODO(), githubcomcasemarkcasedevgo.MemoryV1SearchParams{
		Query:    githubcomcasemarkcasedevgo.F("query"),
		Category: githubcomcasemarkcasedevgo.F("category"),
		Tag1:     githubcomcasemarkcasedevgo.F("tag_1"),
		Tag10:    githubcomcasemarkcasedevgo.F("tag_10"),
		Tag11:    githubcomcasemarkcasedevgo.F("tag_11"),
		Tag12:    githubcomcasemarkcasedevgo.F("tag_12"),
		Tag2:     githubcomcasemarkcasedevgo.F("tag_2"),
		Tag3:     githubcomcasemarkcasedevgo.F("tag_3"),
		Tag4:     githubcomcasemarkcasedevgo.F("tag_4"),
		Tag5:     githubcomcasemarkcasedevgo.F("tag_5"),
		Tag6:     githubcomcasemarkcasedevgo.F("tag_6"),
		Tag7:     githubcomcasemarkcasedevgo.F("tag_7"),
		Tag8:     githubcomcasemarkcasedevgo.F("tag_8"),
		Tag9:     githubcomcasemarkcasedevgo.F("tag_9"),
		TopK:     githubcomcasemarkcasedevgo.F(int64(1)),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
