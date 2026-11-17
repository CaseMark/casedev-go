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

func TestVaultAssemblyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Vault.Assemblies.New(
		context.TODO(),
		"id",
		githubcomcasemarkcasedevgo.VaultAssemblyNewParams{
			Sources: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.VaultAssemblyNewParamsSource{{
				ObjectID:  githubcomcasemarkcasedevgo.F("objectId"),
				Rendition: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsSourcesRenditionSearchablePdf),
			}}),
			Volumes: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolume{{
				Filename: githubcomcasemarkcasedevgo.F("filename"),
				Pages: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolumesPage{{
					End:    githubcomcasemarkcasedevgo.F(int64(1)),
					Source: githubcomcasemarkcasedevgo.F(int64(0)),
					Start:  githubcomcasemarkcasedevgo.F(int64(1)),
					Bates:  githubcomcasemarkcasedevgo.F(true),
				}}),
				Bates: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolumesBates{
					FontSize:     githubcomcasemarkcasedevgo.F(6.000000),
					Legend:       githubcomcasemarkcasedevgo.F("legend"),
					MarginBottom: githubcomcasemarkcasedevgo.F(0.000000),
					MarginSide:   githubcomcasemarkcasedevgo.F(0.000000),
					PadTo:        githubcomcasemarkcasedevgo.F(int64(0)),
					Position:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolumesBatesPositionBottomRight),
					Prefix:       githubcomcasemarkcasedevgo.F("prefix"),
					Shrink:       githubcomcasemarkcasedevgo.F(true),
					Start:        githubcomcasemarkcasedevgo.F(int64(1)),
					Suffix:       githubcomcasemarkcasedevgo.F("suffix"),
				}),
				Contents: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolumesContents{
					Depth: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolumesContentsDepth1),
					Title: githubcomcasemarkcasedevgo.F("title"),
				}),
				Outline: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.VaultAssemblyNewParamsVolumesOutline{{
					Page:     githubcomcasemarkcasedevgo.F(int64(1)),
					Title:    githubcomcasemarkcasedevgo.F("title"),
					Children: githubcomcasemarkcasedevgo.F([]interface{}{map[string]interface{}{}}),
				}}),
				Title: githubcomcasemarkcasedevgo.F("title"),
			}}),
			IdempotencyKey: githubcomcasemarkcasedevgo.F("x"),
			Bates: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsBates{
				FontSize:     githubcomcasemarkcasedevgo.F(6.000000),
				Legend:       githubcomcasemarkcasedevgo.F("legend"),
				MarginBottom: githubcomcasemarkcasedevgo.F(0.000000),
				MarginSide:   githubcomcasemarkcasedevgo.F(0.000000),
				PadTo:        githubcomcasemarkcasedevgo.F(int64(0)),
				Position:     githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsBatesPositionBottomRight),
				Prefix:       githubcomcasemarkcasedevgo.F("prefix"),
				Shrink:       githubcomcasemarkcasedevgo.F(true),
				Start:        githubcomcasemarkcasedevgo.F(int64(1)),
				Suffix:       githubcomcasemarkcasedevgo.F("suffix"),
			}),
			ClientReference: githubcomcasemarkcasedevgo.F("clientReference"),
			Compression: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsCompression{
				Bitonal:    githubcomcasemarkcasedevgo.F(true),
				CleanScans: githubcomcasemarkcasedevgo.F(true),
				Encoding:   githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsCompressionEncodingJpeg2000),
				MaxDpi:     githubcomcasemarkcasedevgo.F(int64(72)),
				Quality:    githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsCompressionQualityHigh),
			}),
			Linearize: githubcomcasemarkcasedevgo.F(true),
			PageSize:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VaultAssemblyNewParamsPageSizeLetter),
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

func TestVaultAssemblyGet(t *testing.T) {
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
	_, err := client.Vault.Assemblies.Get(
		context.TODO(),
		"id",
		"assemblyId",
	)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
