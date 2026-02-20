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

func TestLlmV1ChatNewCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Llm.V1.Chat.NewCompletion(context.TODO(), githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParams{
		Messages: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParamsMessage{{
			Content: githubcomcasemarkcasedevgo.F("content"),
			Role:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.LlmV1ChatNewCompletionParamsMessagesRoleSystem),
		}}),
		CasemarkShowReasoning: githubcomcasemarkcasedevgo.F(false),
		FrequencyPenalty:      githubcomcasemarkcasedevgo.F(0.000000),
		MaxTokens:             githubcomcasemarkcasedevgo.F(int64(1000)),
		Model:                 githubcomcasemarkcasedevgo.F("casemark/casemark-core-3"),
		PresencePenalty:       githubcomcasemarkcasedevgo.F(0.000000),
		Stream:                githubcomcasemarkcasedevgo.F(false),
		Temperature:           githubcomcasemarkcasedevgo.F(0.700000),
		TopP:                  githubcomcasemarkcasedevgo.F(0.000000),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
