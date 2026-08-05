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

func TestMatterV1WorkItemNewWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.WorkItems.New(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1WorkItemNewParams{
			Title:        githubcomcasemarkcasedevgo.F("title"),
			AssigneeID:   githubcomcasemarkcasedevgo.F("assignee_id"),
			Description:  githubcomcasemarkcasedevgo.F("description"),
			DueAt:        githubcomcasemarkcasedevgo.F(time.Now()),
			ExitCriteria: githubcomcasemarkcasedevgo.F([]string{"string"}),
			Instructions: githubcomcasemarkcasedevgo.F("instructions"),
			Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Priority: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1WorkItemNewParamsPriorityLow),
			Type:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1WorkItemNewParamsTypeTask),
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

func TestMatterV1WorkItemGet(t *testing.T) {
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
	err := client.Matters.V1.WorkItems.Get(
		context.TODO(),
		"id",
		"workItemId",
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatterV1WorkItemUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.WorkItems.Update(
		context.TODO(),
		"id",
		"workItemId",
		githubcomcasemarkcasedevgo.MatterV1WorkItemUpdateParams{
			AssigneeID:   githubcomcasemarkcasedevgo.F("assignee_id"),
			CompletedAt:  githubcomcasemarkcasedevgo.F(time.Now()),
			Description:  githubcomcasemarkcasedevgo.F("description"),
			DueAt:        githubcomcasemarkcasedevgo.F(time.Now()),
			ExitCriteria: githubcomcasemarkcasedevgo.F([]string{"string"}),
			Instructions: githubcomcasemarkcasedevgo.F("instructions"),
			Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Priority:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1WorkItemUpdateParamsPriorityLow),
			StartedAt: githubcomcasemarkcasedevgo.F(time.Now()),
			Status:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1WorkItemUpdateParamsStatusDraft),
			Title:     githubcomcasemarkcasedevgo.F("title"),
			Type:      githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1WorkItemUpdateParamsTypeTask),
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

func TestMatterV1WorkItemListWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.WorkItems.List(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1WorkItemListParams{
			AssigneeID: githubcomcasemarkcasedevgo.F("assignee_id"),
			Status:     githubcomcasemarkcasedevgo.F("status"),
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

func TestMatterV1WorkItemDecideWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.WorkItems.Decide(
		context.TODO(),
		"id",
		"workItemId",
		githubcomcasemarkcasedevgo.MatterV1WorkItemDecideParams{
			Decision: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1WorkItemDecideParamsDecisionApprove),
			Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Reason: githubcomcasemarkcasedevgo.F("reason"),
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
