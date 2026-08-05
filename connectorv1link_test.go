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

func TestConnectorV1LinkGet(t *testing.T) {
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
	err := client.Connectors.V1.Links.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectorV1LinkUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Connectors.V1.Links.Update(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ConnectorV1LinkUpdateParams{
			Mode:   githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkUpdateParamsModeOnce),
			Policy: githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
			State:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkUpdateParamsStatePaused),
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

func TestConnectorV1LinkListWithOptionalParams(t *testing.T) {
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
	err := client.Connectors.V1.Links.List(context.TODO(), githubcomcasemarkcasedevgo.ConnectorV1LinkListParams{
		ConnectionID: githubcomcasemarkcasedevgo.F("connection_id"),
		Direction:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkListParamsDirectionImport),
		Mode:         githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkListParamsModeOnce),
		PairID:       githubcomcasemarkcasedevgo.F("pair_id"),
		State:        githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkListParamsStateReady),
		VaultID:      githubcomcasemarkcasedevgo.F("vault_id"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectorV1LinkDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Connectors.V1.Links.Delete(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ConnectorV1LinkDeleteParams{
			VaultDocs: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkDeleteParamsVaultDocsKeep),
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

func TestConnectorV1LinkListObjectsWithOptionalParams(t *testing.T) {
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
	err := client.Connectors.V1.Links.ListObjects(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ConnectorV1LinkListObjectsParams{
			Cursor: githubcomcasemarkcasedevgo.F("cursor"),
			State:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1LinkListObjectsParamsStatePending),
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
