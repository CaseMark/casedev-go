// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CaseMark/casedev-go"
	"github.com/CaseMark/casedev-go/option"
)

func TestVoiceV1SpeakNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := githubcomcasemarkcasedevgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Voice.V1.Speak.New(context.TODO(), githubcomcasemarkcasedevgo.VoiceV1SpeakNewParams{
		Text:                     githubcomcasemarkcasedevgo.F("text"),
		ApplyTextNormalization:   githubcomcasemarkcasedevgo.F(true),
		EnableLogging:            githubcomcasemarkcasedevgo.F(true),
		LanguageCode:             githubcomcasemarkcasedevgo.F("en"),
		ModelID:                  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceV1SpeakNewParamsModelIDElevenMultilingualV2),
		NextText:                 githubcomcasemarkcasedevgo.F("next_text"),
		OptimizeStreamingLatency: githubcomcasemarkcasedevgo.F(int64(0)),
		OutputFormat:             githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceV1SpeakNewParamsOutputFormatMP3_44100_128),
		PreviousText:             githubcomcasemarkcasedevgo.F("previous_text"),
		Seed:                     githubcomcasemarkcasedevgo.F(int64(0)),
		VoiceID:                  githubcomcasemarkcasedevgo.F("voice_id"),
		VoiceSettings: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.VoiceV1SpeakNewParamsVoiceSettings{
			SimilarityBoost: githubcomcasemarkcasedevgo.F(0.000000),
			Stability:       githubcomcasemarkcasedevgo.F(0.000000),
			Style:           githubcomcasemarkcasedevgo.F(0.000000),
			UseSpeakerBoost: githubcomcasemarkcasedevgo.F(true),
		}),
	})
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *githubcomcasemarkcasedevgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
