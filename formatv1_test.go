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

func TestFormatV1NewDocumentWithOptionalParams(t *testing.T) {
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
	resp, err := client.Format.V1.NewDocument(context.TODO(), githubcomcasemarkcasedevgo.FormatV1NewDocumentParams{
		Content:      githubcomcasemarkcasedevgo.F("content"),
		OutputFormat: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.FormatV1NewDocumentParamsOutputFormatPdf),
		InputFormat:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.FormatV1NewDocumentParamsInputFormatMd),
		Options: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.FormatV1NewDocumentParamsOptions{
			Components: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.FormatV1NewDocumentParamsOptionsComponent{{
				Content:    githubcomcasemarkcasedevgo.F("content"),
				Styles:     githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
				TemplateID: githubcomcasemarkcasedevgo.F("templateId"),
				Variables:  githubcomcasemarkcasedevgo.F[any](map[string]interface{}{}),
			}}),
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
