// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/apiquery"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// TranslateV1Service contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTranslateV1Service] method instead.
type TranslateV1Service struct {
	Options []option.RequestOption
}

// NewTranslateV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTranslateV1Service(opts ...option.RequestOption) (r *TranslateV1Service) {
	r = &TranslateV1Service{}
	r.Options = opts
	return
}

// Detect the language of text. Returns the most likely language code and
// confidence score. Supports batch detection for multiple texts.
func (r *TranslateV1Service) Detect(ctx context.Context, body TranslateV1DetectParams, opts ...option.RequestOption) (res *TranslateV1DetectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "translate/v1/detect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the list of languages supported for translation. Optionally specify a target
// language to get translated language names.
func (r *TranslateV1Service) ListLanguages(ctx context.Context, query TranslateV1ListLanguagesParams, opts ...option.RequestOption) (res *TranslateV1ListLanguagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "translate/v1/languages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Translate text between languages using Google Cloud Translation API. Supports
// 100+ languages, automatic language detection, HTML preservation, and batch
// translation.
func (r *TranslateV1Service) Translate(ctx context.Context, body TranslateV1TranslateParams, opts ...option.RequestOption) (res *TranslateV1TranslateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "translate/v1/translate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TranslateV1DetectResponse struct {
	Data TranslateV1DetectResponseData `json:"data"`
	JSON translateV1DetectResponseJSON `json:"-"`
}

// translateV1DetectResponseJSON contains the JSON metadata for the struct
// [TranslateV1DetectResponse]
type translateV1DetectResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1DetectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1DetectResponseJSON) RawJSON() string {
	return r.raw
}

type TranslateV1DetectResponseData struct {
	Detections [][]TranslateV1DetectResponseDataDetection `json:"detections"`
	JSON       translateV1DetectResponseDataJSON          `json:"-"`
}

// translateV1DetectResponseDataJSON contains the JSON metadata for the struct
// [TranslateV1DetectResponseData]
type translateV1DetectResponseDataJSON struct {
	Detections  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1DetectResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1DetectResponseDataJSON) RawJSON() string {
	return r.raw
}

type TranslateV1DetectResponseDataDetection struct {
	// Confidence score (0-1)
	Confidence float64 `json:"confidence"`
	// Whether the detection is reliable
	IsReliable bool `json:"isReliable"`
	// Detected language code (ISO 639-1)
	Language string                                     `json:"language"`
	JSON     translateV1DetectResponseDataDetectionJSON `json:"-"`
}

// translateV1DetectResponseDataDetectionJSON contains the JSON metadata for the
// struct [TranslateV1DetectResponseDataDetection]
type translateV1DetectResponseDataDetectionJSON struct {
	Confidence  apijson.Field
	IsReliable  apijson.Field
	Language    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1DetectResponseDataDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1DetectResponseDataDetectionJSON) RawJSON() string {
	return r.raw
}

type TranslateV1ListLanguagesResponse struct {
	Data TranslateV1ListLanguagesResponseData `json:"data"`
	JSON translateV1ListLanguagesResponseJSON `json:"-"`
}

// translateV1ListLanguagesResponseJSON contains the JSON metadata for the struct
// [TranslateV1ListLanguagesResponse]
type translateV1ListLanguagesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1ListLanguagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1ListLanguagesResponseJSON) RawJSON() string {
	return r.raw
}

type TranslateV1ListLanguagesResponseData struct {
	Languages []TranslateV1ListLanguagesResponseDataLanguage `json:"languages"`
	JSON      translateV1ListLanguagesResponseDataJSON       `json:"-"`
}

// translateV1ListLanguagesResponseDataJSON contains the JSON metadata for the
// struct [TranslateV1ListLanguagesResponseData]
type translateV1ListLanguagesResponseDataJSON struct {
	Languages   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1ListLanguagesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1ListLanguagesResponseDataJSON) RawJSON() string {
	return r.raw
}

type TranslateV1ListLanguagesResponseDataLanguage struct {
	// Language code (ISO 639-1)
	Language string `json:"language"`
	// Language name (if target specified)
	Name string                                           `json:"name"`
	JSON translateV1ListLanguagesResponseDataLanguageJSON `json:"-"`
}

// translateV1ListLanguagesResponseDataLanguageJSON contains the JSON metadata for
// the struct [TranslateV1ListLanguagesResponseDataLanguage]
type translateV1ListLanguagesResponseDataLanguageJSON struct {
	Language    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1ListLanguagesResponseDataLanguage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1ListLanguagesResponseDataLanguageJSON) RawJSON() string {
	return r.raw
}

type TranslateV1TranslateResponse struct {
	Data TranslateV1TranslateResponseData `json:"data"`
	JSON translateV1TranslateResponseJSON `json:"-"`
}

// translateV1TranslateResponseJSON contains the JSON metadata for the struct
// [TranslateV1TranslateResponse]
type translateV1TranslateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslateV1TranslateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1TranslateResponseJSON) RawJSON() string {
	return r.raw
}

type TranslateV1TranslateResponseData struct {
	Translations []TranslateV1TranslateResponseDataTranslation `json:"translations"`
	JSON         translateV1TranslateResponseDataJSON          `json:"-"`
}

// translateV1TranslateResponseDataJSON contains the JSON metadata for the struct
// [TranslateV1TranslateResponseData]
type translateV1TranslateResponseDataJSON struct {
	Translations apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TranslateV1TranslateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1TranslateResponseDataJSON) RawJSON() string {
	return r.raw
}

type TranslateV1TranslateResponseDataTranslation struct {
	// Detected source language (if source not specified)
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
	// Model used for translation
	Model string `json:"model"`
	// Translated text
	TranslatedText string                                          `json:"translatedText"`
	JSON           translateV1TranslateResponseDataTranslationJSON `json:"-"`
}

// translateV1TranslateResponseDataTranslationJSON contains the JSON metadata for
// the struct [TranslateV1TranslateResponseDataTranslation]
type translateV1TranslateResponseDataTranslationJSON struct {
	DetectedSourceLanguage apijson.Field
	Model                  apijson.Field
	TranslatedText         apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TranslateV1TranslateResponseDataTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translateV1TranslateResponseDataTranslationJSON) RawJSON() string {
	return r.raw
}

type TranslateV1DetectParams struct {
	// Text to detect language for. Can be a single string or an array for batch
	// detection.
	Q param.Field[TranslateV1DetectParamsQUnion] `json:"q,required"`
}

func (r TranslateV1DetectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Text to detect language for. Can be a single string or an array for batch
// detection.
//
// Satisfied by [shared.UnionString], [TranslateV1DetectParamsQArray].
type TranslateV1DetectParamsQUnion interface {
	ImplementsTranslateV1DetectParamsQUnion()
}

type TranslateV1DetectParamsQArray []string

func (r TranslateV1DetectParamsQArray) ImplementsTranslateV1DetectParamsQUnion() {}

type TranslateV1ListLanguagesParams struct {
	// Translation model to check language support for
	Model param.Field[TranslateV1ListLanguagesParamsModel] `query:"model"`
	// Target language code for translating language names (e.g., 'es' for Spanish
	// names)
	Target param.Field[string] `query:"target"`
}

// URLQuery serializes [TranslateV1ListLanguagesParams]'s query parameters as
// `url.Values`.
func (r TranslateV1ListLanguagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Translation model to check language support for
type TranslateV1ListLanguagesParamsModel string

const (
	TranslateV1ListLanguagesParamsModelNmt  TranslateV1ListLanguagesParamsModel = "nmt"
	TranslateV1ListLanguagesParamsModelBase TranslateV1ListLanguagesParamsModel = "base"
)

func (r TranslateV1ListLanguagesParamsModel) IsKnown() bool {
	switch r {
	case TranslateV1ListLanguagesParamsModelNmt, TranslateV1ListLanguagesParamsModelBase:
		return true
	}
	return false
}

type TranslateV1TranslateParams struct {
	// Text to translate. Can be a single string or an array for batch translation.
	Q param.Field[TranslateV1TranslateParamsQUnion] `json:"q,required"`
	// Target language code (ISO 639-1)
	Target param.Field[string] `json:"target,required"`
	// Format of the source text. Use 'html' to preserve HTML tags.
	Format param.Field[TranslateV1TranslateParamsFormat] `json:"format"`
	// Translation model. 'nmt' (Neural Machine Translation) is recommended for
	// quality.
	Model param.Field[TranslateV1TranslateParamsModel] `json:"model"`
	// Source language code (ISO 639-1). If not specified, language is auto-detected.
	Source param.Field[string] `json:"source"`
}

func (r TranslateV1TranslateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Text to translate. Can be a single string or an array for batch translation.
//
// Satisfied by [shared.UnionString], [TranslateV1TranslateParamsQArray].
type TranslateV1TranslateParamsQUnion interface {
	ImplementsTranslateV1TranslateParamsQUnion()
}

type TranslateV1TranslateParamsQArray []string

func (r TranslateV1TranslateParamsQArray) ImplementsTranslateV1TranslateParamsQUnion() {}

// Format of the source text. Use 'html' to preserve HTML tags.
type TranslateV1TranslateParamsFormat string

const (
	TranslateV1TranslateParamsFormatText TranslateV1TranslateParamsFormat = "text"
	TranslateV1TranslateParamsFormatHTML TranslateV1TranslateParamsFormat = "html"
)

func (r TranslateV1TranslateParamsFormat) IsKnown() bool {
	switch r {
	case TranslateV1TranslateParamsFormatText, TranslateV1TranslateParamsFormatHTML:
		return true
	}
	return false
}

// Translation model. 'nmt' (Neural Machine Translation) is recommended for
// quality.
type TranslateV1TranslateParamsModel string

const (
	TranslateV1TranslateParamsModelNmt  TranslateV1TranslateParamsModel = "nmt"
	TranslateV1TranslateParamsModelBase TranslateV1TranslateParamsModel = "base"
)

func (r TranslateV1TranslateParamsModel) IsKnown() bool {
	switch r {
	case TranslateV1TranslateParamsModelNmt, TranslateV1TranslateParamsModelBase:
		return true
	}
	return false
}
