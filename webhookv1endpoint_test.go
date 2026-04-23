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

func TestWebhookV1EndpointNewWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.New(context.TODO(), githubcomcasemarkcasedevgo.WebhookV1EndpointNewParams{
		EventTypeFilters: githubcomcasemarkcasedevgo.F([]string{"string"}),
		URL:              githubcomcasemarkcasedevgo.F("https://example.com"),
		Description:      githubcomcasemarkcasedevgo.F("description"),
		ResourceScopes: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.WebhookV1EndpointNewParamsResourceScopes{
			MatterIDs: githubcomcasemarkcasedevgo.F([]string{"string"}),
			VaultIDs:  githubcomcasemarkcasedevgo.F([]string{"string"}),
		}),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookV1EndpointGet(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookV1EndpointUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.Update(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.WebhookV1EndpointUpdateParams{
			Description:      githubcomcasemarkcasedevgo.F("description"),
			EventTypeFilters: githubcomcasemarkcasedevgo.F([]string{"string"}),
			ResourceScopes: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.WebhookV1EndpointUpdateParamsResourceScopes{
				MatterIDs: githubcomcasemarkcasedevgo.F([]string{"string"}),
				VaultIDs:  githubcomcasemarkcasedevgo.F([]string{"string"}),
			}),
			Status: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.WebhookV1EndpointUpdateParamsStatusActive),
			URL:    githubcomcasemarkcasedevgo.F("https://example.com"),
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

func TestWebhookV1EndpointListWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.List(context.TODO(), githubcomcasemarkcasedevgo.WebhookV1EndpointListParams{
		Limit:  githubcomcasemarkcasedevgo.F(int64(1)),
		Status: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.WebhookV1EndpointListParamsStatusActive),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookV1EndpointDelete(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookV1EndpointRotateSecretWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.RotateSecret(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.WebhookV1EndpointRotateSecretParams{
			PreviousSecretExpiresInSec: githubcomcasemarkcasedevgo.F(int64(0)),
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

func TestWebhookV1EndpointTestWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.V1.Endpoints.Test(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.WebhookV1EndpointTestParams{
			EventType: githubcomcasemarkcasedevgo.F("eventType"),
			Payload:   githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
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
