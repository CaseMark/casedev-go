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

func TestAgentV1ExecuteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent.V1.Execute.New(context.TODO(), githubcomcasemarkcasedevgo.AgentV1ExecuteNewParams{
		Prompt:        githubcomcasemarkcasedevgo.F("prompt"),
		DisabledTools: githubcomcasemarkcasedevgo.F([]string{"string"}),
		EnabledTools:  githubcomcasemarkcasedevgo.F([]string{"string"}),
		Guidance:      githubcomcasemarkcasedevgo.F("guidance"),
		Instructions:  githubcomcasemarkcasedevgo.F("instructions"),
		Model:         githubcomcasemarkcasedevgo.F("model"),
		ObjectIDs:     githubcomcasemarkcasedevgo.F([]string{"string"}),
		Sandbox: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.AgentV1ExecuteNewParamsSandbox{
			CPU:       githubcomcasemarkcasedevgo.F(int64(0)),
			MemoryMiB: githubcomcasemarkcasedevgo.F(int64(0)),
		}),
		VaultIDs: githubcomcasemarkcasedevgo.F([]string{"string"}),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
