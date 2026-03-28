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

func TestMatterV1AgentTypeNewWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.AgentTypes.New(context.TODO(), githubcomcasemarkcasedevgo.MatterV1AgentTypeNewParams{
		Instructions:  githubcomcasemarkcasedevgo.F("instructions"),
		Name:          githubcomcasemarkcasedevgo.F("name"),
		Description:   githubcomcasemarkcasedevgo.F("description"),
		DisabledTools: githubcomcasemarkcasedevgo.F([]string{"string"}),
		EnabledTools:  githubcomcasemarkcasedevgo.F([]string{"string"}),
		IsActive:      githubcomcasemarkcasedevgo.F(true),
		IsDefault:     githubcomcasemarkcasedevgo.F(true),
		Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Model:     githubcomcasemarkcasedevgo.F("model"),
		SkillRefs: githubcomcasemarkcasedevgo.F([]string{"string"}),
		Slug:      githubcomcasemarkcasedevgo.F("slug"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatterV1AgentTypeListWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.AgentTypes.List(context.TODO(), githubcomcasemarkcasedevgo.MatterV1AgentTypeListParams{
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
