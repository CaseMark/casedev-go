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

func TestVoiceTranscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Voice.Transcription.New(context.TODO(), githubcomcasemarkcasedevgo.VoiceTranscriptionNewParams{
		AudioURL:            githubcomcasemarkcasedevgo.F("audio_url"),
		AutoHighlights:      githubcomcasemarkcasedevgo.F(true),
		BoostParam:          githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceTranscriptionNewParamsBoostParamLow),
		ContentSafetyLabels: githubcomcasemarkcasedevgo.F(true),
		Format:              githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceTranscriptionNewParamsFormatJson),
		FormatText:          githubcomcasemarkcasedevgo.F(true),
		LanguageCode:        githubcomcasemarkcasedevgo.F("language_code"),
		LanguageDetection:   githubcomcasemarkcasedevgo.F(true),
		ObjectID:            githubcomcasemarkcasedevgo.F("object_id"),
		Punctuate:           githubcomcasemarkcasedevgo.F(true),
		SpeakerLabels:       githubcomcasemarkcasedevgo.F(true),
		SpeakersExpected:    githubcomcasemarkcasedevgo.F(int64(0)),
		SpeechModels:        githubcomcasemarkcasedevgo.F([]string{"string"}),
		VaultID:             githubcomcasemarkcasedevgo.F("vault_id"),
		WordBoost:           githubcomcasemarkcasedevgo.F([]string{"string"}),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceTranscriptionGet(t *testing.T) {
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
	_, err := client.Voice.Transcription.Get(context.TODO(), "tr_abc123def456")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceTranscriptionDelete(t *testing.T) {
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
	err := client.Voice.Transcription.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
