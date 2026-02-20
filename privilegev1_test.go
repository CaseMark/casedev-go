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

func TestPrivilegeV1DetectWithOptionalParams(t *testing.T) {
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
	_, err := client.Privilege.V1.Detect(context.TODO(), githubcomcasemarkcasedevgo.PrivilegeV1DetectParams{
		Categories:       githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.PrivilegeV1DetectParamsCategory{githubcomcasemarkcasedevgo.PrivilegeV1DetectParamsCategoryAttorneyClient}),
		Content:          githubcomcasemarkcasedevgo.F("content"),
		DocumentID:       githubcomcasemarkcasedevgo.F("document_id"),
		IncludeRationale: githubcomcasemarkcasedevgo.F(true),
		Jurisdiction:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.PrivilegeV1DetectParamsJurisdictionUsFederal),
		Model:            githubcomcasemarkcasedevgo.F("model"),
		VaultID:          githubcomcasemarkcasedevgo.F("vault_id"),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
