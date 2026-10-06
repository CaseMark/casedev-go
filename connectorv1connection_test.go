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

func TestConnectorV1ConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectors.V1.Connections.New(context.TODO(), githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewParams{
		Provider:              githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewParamsProviderBox),
		ReturnURL:             githubcomcasemarkcasedevgo.F("return_url"),
		ScopeTier:             githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1ConnectionNewParamsScopeTierBoxReadwrite),
		XCaseConnectorSubject: githubcomcasemarkcasedevgo.F("x-case-connector-subject"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectorV1ConnectionGetWithOptionalParams(t *testing.T) {
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
	err := client.Connectors.V1.Connections.Get(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ConnectorV1ConnectionGetParams{
			XCaseConnectorSubject: githubcomcasemarkcasedevgo.F("x-case-connector-subject"),
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

func TestConnectorV1ConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectors.V1.Connections.List(context.TODO(), githubcomcasemarkcasedevgo.ConnectorV1ConnectionListParams{
		Cursor:                githubcomcasemarkcasedevgo.F("cursor"),
		Limit:                 githubcomcasemarkcasedevgo.F(int64(1)),
		Provider:              githubcomcasemarkcasedevgo.F("provider"),
		Status:                githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1ConnectionListParamsStatusPending),
		XCaseConnectorSubject: githubcomcasemarkcasedevgo.F("x-case-connector-subject"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectorV1ConnectionDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Connectors.V1.Connections.Delete(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ConnectorV1ConnectionDeleteParams{
			Purge:                 githubcomcasemarkcasedevgo.F(true),
			XCaseConnectorSubject: githubcomcasemarkcasedevgo.F("x-case-connector-subject"),
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

func TestConnectorV1ConnectionBrowseWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectors.V1.Connections.Browse(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.ConnectorV1ConnectionBrowseParams{
			Container:             githubcomcasemarkcasedevgo.F("container"),
			Cursor:                githubcomcasemarkcasedevgo.F("cursor"),
			PageSize:              githubcomcasemarkcasedevgo.F(int64(1000)),
			Parent:                githubcomcasemarkcasedevgo.F("parent"),
			Query:                 githubcomcasemarkcasedevgo.F("query"),
			Site:                  githubcomcasemarkcasedevgo.F("site"),
			XCaseConnectorSubject: githubcomcasemarkcasedevgo.F("x-case-connector-subject"),
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

func TestConnectorV1ConnectionUpdateAll(t *testing.T) {
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
	err := client.Connectors.V1.Connections.UpdateAll(context.TODO(), githubcomcasemarkcasedevgo.ConnectorV1ConnectionUpdateAllParams{
		ConfirmOrganizationWide: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1ConnectionUpdateAllParamsConfirmOrganizationWideTrue),
		Enabled:                 githubcomcasemarkcasedevgo.F(true),
		Provider:                githubcomcasemarkcasedevgo.F("provider"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
