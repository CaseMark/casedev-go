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
	"github.com/CaseMark/casedev-go/shared"
)

func TestSuperdocV1AnnotateWithOptionalParams(t *testing.T) {
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
	resp, err := client.Superdoc.V1.Annotate(context.TODO(), githubcomcasemarkcasedevgo.SuperdocV1AnnotateParams{
		Document: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SuperdocV1AnnotateParamsDocument{
			Base64: githubcomcasemarkcasedevgo.F("base64"),
			URL:    githubcomcasemarkcasedevgo.F("url"),
		}),
		Fields: githubcomcasemarkcasedevgo.F([]githubcomcasemarkcasedevgo.SuperdocV1AnnotateParamsField{{
			Type:  githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SuperdocV1AnnotateParamsFieldsTypeText),
			Value: githubcomcasemarkcasedevgo.F[githubcomcasemarkcasedevgo.SuperdocV1AnnotateParamsFieldsValueUnion](shared.UnionString("string")),
			ID:    githubcomcasemarkcasedevgo.F("id"),
			Group: githubcomcasemarkcasedevgo.F("group"),
			Options: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SuperdocV1AnnotateParamsFieldsOptions{
				Height: githubcomcasemarkcasedevgo.F(0.000000),
				Width:  githubcomcasemarkcasedevgo.F(0.000000),
			}),
		}}),
		OutputFormat: githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SuperdocV1AnnotateParamsOutputFormatDocx),
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

func TestSuperdocV1ConvertWithOptionalParams(t *testing.T) {
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
	resp, err := client.Superdoc.V1.Convert(context.TODO(), githubcomcasemarkcasedevgo.SuperdocV1ConvertParams{
		From:           githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SuperdocV1ConvertParamsFromDocx),
		DocumentBase64: githubcomcasemarkcasedevgo.F("document_base64"),
		DocumentURL:    githubcomcasemarkcasedevgo.F("document_url"),
		To:             githubcomcasemarkcasedevgo.F(githubcomcasemarkcasedevgo.SuperdocV1ConvertParamsToPdf),
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
