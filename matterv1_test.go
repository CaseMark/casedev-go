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

func TestMatterV1NewWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.New(context.TODO(), githubcomcasemarkcasedevgo.MatterV1NewParams{
		Title: githubcomcasemarkcasedevgo.F("title"),
		Billing: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		ClientName:    githubcomcasemarkcasedevgo.F("client_name"),
		ClientPartyID: githubcomcasemarkcasedevgo.F("client_party_id"),
		CustomFields: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Description: githubcomcasemarkcasedevgo.F("description"),
		DisplayID:   githubcomcasemarkcasedevgo.F("display_id"),
		ImportantDates: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		Jurisdiction: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		MatterType: githubcomcasemarkcasedevgo.F("matter_type"),
		Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
			"foo": "bar",
		}),
		PracticeArea:          githubcomcasemarkcasedevgo.F("practice_area"),
		ResponsibleAttorneyID: githubcomcasemarkcasedevgo.F("responsible_attorney_id"),
		Status:                githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1NewParamsStatusIntake),
		Subtype:               githubcomcasemarkcasedevgo.F("subtype"),
		Vault: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1NewParamsVault{
			Description:    githubcomcasemarkcasedevgo.F("description"),
			EnableGraph:    githubcomcasemarkcasedevgo.F(true),
			EnableIndexing: githubcomcasemarkcasedevgo.F(true),
			Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
		}),
		VaultID: githubcomcasemarkcasedevgo.F("vault_id"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatterV1Get(t *testing.T) {
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
	err := client.Matters.V1.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatterV1UpdateWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.Update(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.MatterV1UpdateParams{
			ArchivedAt: githubcomcasemarkcasedevgo.F(time.Now()),
			Billing: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			ClientName:    githubcomcasemarkcasedevgo.F("client_name"),
			ClientPartyID: githubcomcasemarkcasedevgo.F("client_party_id"),
			CustomFields: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Description: githubcomcasemarkcasedevgo.F("description"),
			DisplayID:   githubcomcasemarkcasedevgo.F("display_id"),
			ImportantDates: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			Jurisdiction: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			MatterType: githubcomcasemarkcasedevgo.F("matter_type"),
			Metadata: githubcomcasemarkcasedevgo.F(map[string]interface{}{
				"foo": "bar",
			}),
			PracticeArea:          githubcomcasemarkcasedevgo.F("practice_area"),
			ResponsibleAttorneyID: githubcomcasemarkcasedevgo.F("responsible_attorney_id"),
			Status:                githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.MatterV1UpdateParamsStatusIntake),
			Subtype:               githubcomcasemarkcasedevgo.F("subtype"),
			Title:                 githubcomcasemarkcasedevgo.F("title"),
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

func TestMatterV1ListWithOptionalParams(t *testing.T) {
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
	err := client.Matters.V1.List(context.TODO(), githubcomcasemarkcasedevgo.MatterV1ListParams{
		MatterType:   githubcomcasemarkcasedevgo.F("matter_type"),
		PracticeArea: githubcomcasemarkcasedevgo.F("practice_area"),
		Query:        githubcomcasemarkcasedevgo.F("query"),
		Status:       githubcomcasemarkcasedevgo.F("status"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
