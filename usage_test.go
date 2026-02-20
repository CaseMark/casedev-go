// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/internal/testutil"
	"github.com/CaseMark/casedev-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Mock server tests are disabled")
	response, err := client.Llm.V1.Chat.NewCompletion(context.TODO(), githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParams{
		Messages: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParamsMessage{{
			Role:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParamsMessagesRoleUser),
			Content: githubcomcasemarkcasedevgo.F("Hello!"),
		}}),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.ID)
}
