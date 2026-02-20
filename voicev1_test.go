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

func TestVoiceV1ListVoicesWithOptionalParams(t *testing.T) {
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
	_, err := client.Voice.V1.ListVoices(context.TODO(), githubcomcasemarkcasedevgo.VoiceV1ListVoicesParams{
		Category:          githubcomcasemarkcasedevgo.F("category"),
		CollectionID:      githubcomcasemarkcasedevgo.F("collection_id"),
		IncludeTotalCount: githubcomcasemarkcasedevgo.F(true),
		NextPageToken:     githubcomcasemarkcasedevgo.F("next_page_token"),
		PageSize:          githubcomcasemarkcasedevgo.F(int64(1)),
		Search:            githubcomcasemarkcasedevgo.F("search"),
		Sort:              githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceV1ListVoicesParamsSortName),
		SortDirection:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceV1ListVoicesParamsSortDirectionAsc),
		VoiceType:         githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceV1ListVoicesParamsVoiceTypePremade),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
