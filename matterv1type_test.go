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

func TestMatterV1TypeNewWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.Types.New(context.TODO(), githubcomcasemarkcasedevgo.MatterV1TypeNewParams{
		Name:               githubcomcasemarkcasedevgo.F("name"),
		DefaultAgentTypeID: githubcomcasemarkcasedevgo.F("default_agent_type_id"),
		DefaultMetadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		DefaultWorkItems: githubcomcasemarkcasedevgo.F([]map[string]interface{}{{
			"foo": "bar",
		}}),
		Description:        githubcomcasemarkcasedevgo.F("description"),
		ExitCriteria:       githubcomcasemarkcasedevgo.F([]string{"string"}),
		Instructions:       githubcomcasemarkcasedevgo.F("instructions"),
		IntakeRequirements: githubcomcasemarkcasedevgo.F([]string{"string"}),
		IsActive:           githubcomcasemarkcasedevgo.F(true),
		OrchestrationMode:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1TypeNewParamsOrchestrationModeAuto),
		ReviewAgentTypeID:  githubcomcasemarkcasedevgo.F("review_agent_type_id"),
		ReviewCriteria:     githubcomcasemarkcasedevgo.F([]string{"string"}),
		SkillRefs:          githubcomcasemarkcasedevgo.F([]string{"string"}),
		Slug:               githubcomcasemarkcasedevgo.F("slug"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatterV1TypeGet(t *testing.T) {
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
	err := client.Matters.V1.Types.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatterV1TypeUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.Types.Update(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1TypeUpdateParams{
			DefaultAgentTypeID: githubcomcasemarkcasedevgo.F("default_agent_type_id"),
			DefaultMetadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			DefaultWorkItems: githubcomcasemarkcasedevgo.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			Description:        githubcomcasemarkcasedevgo.F("description"),
			ExitCriteria:       githubcomcasemarkcasedevgo.F([]string{"string"}),
			Instructions:       githubcomcasemarkcasedevgo.F("instructions"),
			IntakeRequirements: githubcomcasemarkcasedevgo.F([]string{"string"}),
			IsActive:           githubcomcasemarkcasedevgo.F(true),
			Name:               githubcomcasemarkcasedevgo.F("name"),
			OrchestrationMode:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1TypeUpdateParamsOrchestrationModeAuto),
			ReviewAgentTypeID:  githubcomcasemarkcasedevgo.F("review_agent_type_id"),
			ReviewCriteria:     githubcomcasemarkcasedevgo.F([]string{"string"}),
			SkillRefs:          githubcomcasemarkcasedevgo.F([]string{"string"}),
			Slug:               githubcomcasemarkcasedevgo.F("slug"),
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

func TestMatterV1TypeListWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.Types.List(context.TODO(), githubcomcasemarkcasedevgo.MatterV1TypeListParams{
		Active: githubcomcasemarkcasedevgo.F(true),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
