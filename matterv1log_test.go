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
	"github.com/CaseMark/casedev-go/shared"
)

func TestMatterV1LogNewWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.Log.New(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1LogNewParams{
			Summary: githubcomcasemarkcasedevgo.F("summary"),
			Details: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			EventType:  githubcomcasemarkcasedevgo.F("event_type"),
			WorkItemID: githubcomcasemarkcasedevgo.F("work_item_id"),
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

func TestMatterV1LogListWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.Log.List(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1LogListParams{
			ActorID:    githubcomcasemarkcasedevgo.F("actor_id"),
			ActorType:  githubcomcasemarkcasedevgo.F("actor_type"),
			EndTime:    githubcomcasemarkcasedevgo.F(time.Now()),
			EventType:  githubcomcasemarkcasedevgo.F("event_type"),
			Limit:      githubcomcasemarkcasedevgo.F(int64(200)),
			Offset:     githubcomcasemarkcasedevgo.F(int64(0)),
			Scope:      githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.MatterV1LogListParamsScopeUnion](shared.UnionString("string")),
			StartTime:  githubcomcasemarkcasedevgo.F(time.Now()),
			WorkItemID: githubcomcasemarkcasedevgo.F("work_item_id"),
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

func TestMatterV1LogExportWithOptionalParams(t *testing.T) {
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
	_, err := client.Matters.V1.Log.Export(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1LogExportParams{
			ActorID:    githubcomcasemarkcasedevgo.F("actor_id"),
			ActorType:  githubcomcasemarkcasedevgo.F("actor_type"),
			EndTime:    githubcomcasemarkcasedevgo.F(time.Now()),
			EventType:  githubcomcasemarkcasedevgo.F("event_type"),
			Format:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1LogExportParamsFormatJson),
			Scope:      githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.MatterV1LogExportParamsScopeUnion](shared.UnionString("string")),
			StartTime:  githubcomcasemarkcasedevgo.F(time.Now()),
			WorkItemID: githubcomcasemarkcasedevgo.F("work_item_id"),
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
