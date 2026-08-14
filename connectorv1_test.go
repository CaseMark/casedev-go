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

func TestConnectorV1SyncLinkWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectors.V1.SyncLink(context.TODO(), githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParams{
		ConnectionID: githubcomcasemarkcasedevgo.F("connection_id"),
		Direction:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsDirectionImport),
		Remote: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsRemote{
			FolderID:    githubcomcasemarkcasedevgo.F("folder_id"),
			ContainerID: githubcomcasemarkcasedevgo.F("container_id"),
			Path:        githubcomcasemarkcasedevgo.F("path"),
			SiteID:      githubcomcasemarkcasedevgo.F("site_id"),
		}),
		VaultID: githubcomcasemarkcasedevgo.F("vault_id"),
		ExportDestination: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsExportDestination{
			FolderID:    githubcomcasemarkcasedevgo.F("folder_id"),
			ContainerID: githubcomcasemarkcasedevgo.F("container_id"),
			Path:        githubcomcasemarkcasedevgo.F("path"),
			SiteID:      githubcomcasemarkcasedevgo.F("site_id"),
		}),
		MatterID: githubcomcasemarkcasedevgo.F("matter_id"),
		Policy: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsPolicy{
			Collisions: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsPolicyCollisionsVersion),
			Deletes:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsPolicyDeletesMirror),
			Filters: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1SyncLinkParamsPolicyFilters{
				ExcludeMime:  githubcomcasemarkcasedevgo.F([]string{"string"}),
				MaxSizeBytes: githubcomcasemarkcasedevgo.F(int64(0)),
			}),
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

func TestConnectorV1TransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Connectors.V1.Transfer(context.TODO(), githubcomcasemarkcasedevgo.ConnectorV1TransferParams{
		ConnectionID: githubcomcasemarkcasedevgo.F("connection_id"),
		Direction:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsDirectionImport),
		Remote: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsRemote{
			FolderID:    githubcomcasemarkcasedevgo.F("folder_id"),
			ContainerID: githubcomcasemarkcasedevgo.F("container_id"),
			Path:        githubcomcasemarkcasedevgo.F("path"),
			SiteID:      githubcomcasemarkcasedevgo.F("site_id"),
		}),
		VaultID: githubcomcasemarkcasedevgo.F("vault_id"),
		ExportDestination: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsExportDestination{
			FolderID:    githubcomcasemarkcasedevgo.F("folder_id"),
			ContainerID: githubcomcasemarkcasedevgo.F("container_id"),
			Path:        githubcomcasemarkcasedevgo.F("path"),
			SiteID:      githubcomcasemarkcasedevgo.F("site_id"),
		}),
		MatterID: githubcomcasemarkcasedevgo.F("matter_id"),
		Policy: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsPolicy{
			Collisions: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsPolicyCollisionsVersion),
			Deletes:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsPolicyDeletesMirror),
			Filters: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsPolicyFilters{
				ExcludeMime:  githubcomcasemarkcasedevgo.F([]string{"string"}),
				MaxSizeBytes: githubcomcasemarkcasedevgo.F(int64(0)),
			}),
		}),
		RunMode: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.ConnectorV1TransferParamsRunModeAuto),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
