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

func TestMatterV1MatterPartyNewWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.MatterParties.New(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1MatterPartyNewParams{
			PartyID: githubcomcasemarkcasedevgo.F("party_id"),
			Role:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1MatterPartyNewParamsRoleClient),
			CustomFields: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			IsPrimary: githubcomcasemarkcasedevgo.F(true),
			Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Notes:       githubcomcasemarkcasedevgo.F("notes"),
			SetAsClient: githubcomcasemarkcasedevgo.F(true),
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

func TestMatterV1MatterPartyList(t *testing.T) {
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
	err := client.Matters.V1.MatterParties.List(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
