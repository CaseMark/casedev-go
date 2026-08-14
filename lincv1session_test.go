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

func TestLincV1SessionNewWithOptionalParams(t *testing.T) {
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
	err := client.Linc.V1.Sessions.New(context.TODO(), githubcomcasemarkcasedevgo.LincV1SessionNewParams{
		DocumentTemplateSlugs:    githubcomcasemarkcasedevgo.F([]string{"string"}),
		IdleTimeoutMs:            githubcomcasemarkcasedevgo.F(int64(0)),
		IncludeDocumentTemplates: githubcomcasemarkcasedevgo.F(true),
		Instructions:             githubcomcasemarkcasedevgo.F("instructions"),
		Model:                    githubcomcasemarkcasedevgo.F("model"),
		ScopedAPIKey:             githubcomcasemarkcasedevgo.F("scopedApiKey"),
		ServiceTier:              githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LincV1SessionNewParamsServiceTierDefault),
		SkillSlugs:               githubcomcasemarkcasedevgo.F([]string{"string"}),
		Title:                    githubcomcasemarkcasedevgo.F("title"),
		VaultIDs:                 githubcomcasemarkcasedevgo.F([]string{"string"}),
		AIReportingTags:          githubcomcasemarkcasedevgo.F("ai-reporting-tags"),
		AIReportingUser:          githubcomcasemarkcasedevgo.F("ai-reporting-user"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLincV1SessionDelete(t *testing.T) {
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
	err := client.Linc.V1.Sessions.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLincV1SessionCancelWithOptionalParams(t *testing.T) {
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
	err := client.Linc.V1.Sessions.Cancel(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.LincV1SessionCancelParams{
			ClearQueue: githubcomcasemarkcasedevgo.F(true),
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

func TestLincV1SessionIngestEvents(t *testing.T) {
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
	err := client.Linc.V1.Sessions.IngestEvents(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.LincV1SessionIngestEventsParams{
			Frames: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.LincV1SessionIngestEventsParamsFrame{{
				Event: githubcomcasemarkcasedevgo.F(map[string]interface{}{
					"foo": "bar",
				}),
				Seq:  githubcomcasemarkcasedevgo.F(int64(1)),
				Type: githubcomcasemarkcasedevgo.F("type"),
			}}),
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

func TestLincV1SessionGetEventsWithOptionalParams(t *testing.T) {
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
	err := client.Linc.V1.Sessions.GetEvents(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.LincV1SessionGetEventsParams{
			AfterSeq:          githubcomcasemarkcasedevgo.F(int64(0)),
			Cursor:            githubcomcasemarkcasedevgo.F(int64(0)),
			ExcludeEventTypes: githubcomcasemarkcasedevgo.F([]string{"string"}),
			Limit:             githubcomcasemarkcasedevgo.F(int64(1)),
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

func TestLincV1SessionGetMessagesWithOptionalParams(t *testing.T) {
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
	err := client.Linc.V1.Sessions.GetMessages(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.LincV1SessionGetMessagesParams{
			AfterSeq: githubcomcasemarkcasedevgo.F(int64(0)),
			Cursor:   githubcomcasemarkcasedevgo.F(int64(0)),
			Limit:    githubcomcasemarkcasedevgo.F(int64(1)),
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

func TestLincV1SessionGetState(t *testing.T) {
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
	err := client.Linc.V1.Sessions.GetState(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLincV1SessionSendRpcWithOptionalParams(t *testing.T) {
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
	err := client.Linc.V1.Sessions.SendRpc(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.LincV1SessionSendRpcParams{
			Type: githubcomcasemarkcasedevgo.F("type"),
			ID:   githubcomcasemarkcasedevgo.F("id"),
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
