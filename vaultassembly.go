// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomcasemarkcasedevgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CaseMark/casedev-go/internal/apijson"
	"github.com/CaseMark/casedev-go/internal/param"
	"github.com/CaseMark/casedev-go/internal/requestconfig"
	"github.com/CaseMark/casedev-go/option"
)

// Vault object management, content access, and document operations
//
// VaultAssemblyService contains methods and other services that help with
// interacting with the casedev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultAssemblyService] method instead.
type VaultAssemblyService struct {
	Options []option.RequestOption
}

// NewVaultAssemblyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVaultAssemblyService(opts ...option.RequestOption) (r *VaultAssemblyService) {
	r = &VaultAssemblyService{}
	r.Options = opts
	return
}

// Starts an asynchronous, deterministic PDF assembly. The manifest selects page
// ranges from source PDFs in output order, adds an outline, an optional linked
// contents page and optional Bates numbering, and splits the result into volumes.
// Optional compression downsamples oversized scans. Each volume becomes a new
// vault object. Source bookmarks, form fields and page labels are not carried
// over. Outputs are not indexed automatically. Completion is reported through
// vault.assembly webhooks.
func (r *VaultAssemblyService) New(ctx context.Context, id string, params VaultAssemblyNewParams, opts ...option.RequestOption) (res *VaultAssemblyNewResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/assemblies", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the status of a PDF assembly. Completed volumes include their page
// count, size, SHA-256 and Bates range.
func (r *VaultAssemblyService) Get(ctx context.Context, id string, assemblyID string, opts ...option.RequestOption) (res *VaultAssemblyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if assemblyID == "" {
		err = errors.New("missing required assemblyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("vault/%s/assemblies/%s", id, assemblyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type VaultAssemblyNewResponse struct {
	ID              string                           `json:"id"`
	ClientReference string                           `json:"clientReference"`
	Status          VaultAssemblyNewResponseStatus   `json:"status"`
	Volumes         []VaultAssemblyNewResponseVolume `json:"volumes"`
	JSON            vaultAssemblyNewResponseJSON     `json:"-"`
}

// vaultAssemblyNewResponseJSON contains the JSON metadata for the struct
// [VaultAssemblyNewResponse]
type vaultAssemblyNewResponseJSON struct {
	ID              apijson.Field
	ClientReference apijson.Field
	Status          apijson.Field
	Volumes         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaultAssemblyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyNewResponseJSON) RawJSON() string {
	return r.raw
}

type VaultAssemblyNewResponseStatus string

const (
	VaultAssemblyNewResponseStatusQueued     VaultAssemblyNewResponseStatus = "queued"
	VaultAssemblyNewResponseStatusProcessing VaultAssemblyNewResponseStatus = "processing"
	VaultAssemblyNewResponseStatusCompleted  VaultAssemblyNewResponseStatus = "completed"
	VaultAssemblyNewResponseStatusFailed     VaultAssemblyNewResponseStatus = "failed"
)

func (r VaultAssemblyNewResponseStatus) IsKnown() bool {
	switch r {
	case VaultAssemblyNewResponseStatusQueued, VaultAssemblyNewResponseStatusProcessing, VaultAssemblyNewResponseStatusCompleted, VaultAssemblyNewResponseStatusFailed:
		return true
	}
	return false
}

type VaultAssemblyNewResponseVolume struct {
	Filename string                             `json:"filename"`
	ObjectID string                             `json:"objectId"`
	JSON     vaultAssemblyNewResponseVolumeJSON `json:"-"`
}

// vaultAssemblyNewResponseVolumeJSON contains the JSON metadata for the struct
// [VaultAssemblyNewResponseVolume]
type vaultAssemblyNewResponseVolumeJSON struct {
	Filename    apijson.Field
	ObjectID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultAssemblyNewResponseVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyNewResponseVolumeJSON) RawJSON() string {
	return r.raw
}

type VaultAssemblyGetResponse struct {
	ID              string `json:"id"`
	ClientReference string `json:"clientReference"`
	// What the assembly cost Case.dev, once it is completed or failed. Internal COGS,
	// not a price.
	Cogs    VaultAssemblyGetResponseCogs     `json:"cogs"`
	Error   VaultAssemblyGetResponseError    `json:"error"`
	Status  VaultAssemblyGetResponseStatus   `json:"status"`
	Timings VaultAssemblyGetResponseTimings  `json:"timings"`
	VaultID string                           `json:"vaultId"`
	Volumes []VaultAssemblyGetResponseVolume `json:"volumes"`
	JSON    vaultAssemblyGetResponseJSON     `json:"-"`
}

// vaultAssemblyGetResponseJSON contains the JSON metadata for the struct
// [VaultAssemblyGetResponse]
type vaultAssemblyGetResponseJSON struct {
	ID              apijson.Field
	ClientReference apijson.Field
	Cogs            apijson.Field
	Error           apijson.Field
	Status          apijson.Field
	Timings         apijson.Field
	VaultID         apijson.Field
	Volumes         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaultAssemblyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyGetResponseJSON) RawJSON() string {
	return r.raw
}

// What the assembly cost Case.dev, once it is completed or failed. Internal COGS,
// not a price.
type VaultAssemblyGetResponseCogs struct {
	ComputeMicros  int64                            `json:"computeMicros"`
	CPU            float64                          `json:"cpu"`
	MemoryMib      int64                            `json:"memoryMib"`
	OutputBytes    int64                            `json:"outputBytes"`
	RateVersion    string                           `json:"rateVersion"`
	SourceBytes    int64                            `json:"sourceBytes"`
	TransferMicros int64                            `json:"transferMicros"`
	WorkerSeconds  float64                          `json:"workerSeconds"`
	JSON           vaultAssemblyGetResponseCogsJSON `json:"-"`
}

// vaultAssemblyGetResponseCogsJSON contains the JSON metadata for the struct
// [VaultAssemblyGetResponseCogs]
type vaultAssemblyGetResponseCogsJSON struct {
	ComputeMicros  apijson.Field
	CPU            apijson.Field
	MemoryMib      apijson.Field
	OutputBytes    apijson.Field
	RateVersion    apijson.Field
	SourceBytes    apijson.Field
	TransferMicros apijson.Field
	WorkerSeconds  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VaultAssemblyGetResponseCogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyGetResponseCogsJSON) RawJSON() string {
	return r.raw
}

type VaultAssemblyGetResponseError struct {
	Code        string                            `json:"code"`
	Message     string                            `json:"message"`
	Retryable   bool                              `json:"retryable"`
	SourceIndex int64                             `json:"sourceIndex"`
	Stage       string                            `json:"stage"`
	VolumeIndex int64                             `json:"volumeIndex"`
	JSON        vaultAssemblyGetResponseErrorJSON `json:"-"`
}

// vaultAssemblyGetResponseErrorJSON contains the JSON metadata for the struct
// [VaultAssemblyGetResponseError]
type vaultAssemblyGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Retryable   apijson.Field
	SourceIndex apijson.Field
	Stage       apijson.Field
	VolumeIndex apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultAssemblyGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type VaultAssemblyGetResponseStatus string

const (
	VaultAssemblyGetResponseStatusQueued     VaultAssemblyGetResponseStatus = "queued"
	VaultAssemblyGetResponseStatusProcessing VaultAssemblyGetResponseStatus = "processing"
	VaultAssemblyGetResponseStatusCompleted  VaultAssemblyGetResponseStatus = "completed"
	VaultAssemblyGetResponseStatusFailed     VaultAssemblyGetResponseStatus = "failed"
	VaultAssemblyGetResponseStatusCancelled  VaultAssemblyGetResponseStatus = "cancelled"
)

func (r VaultAssemblyGetResponseStatus) IsKnown() bool {
	switch r {
	case VaultAssemblyGetResponseStatusQueued, VaultAssemblyGetResponseStatusProcessing, VaultAssemblyGetResponseStatusCompleted, VaultAssemblyGetResponseStatusFailed, VaultAssemblyGetResponseStatusCancelled:
		return true
	}
	return false
}

type VaultAssemblyGetResponseTimings struct {
	AcceptedAt      time.Time                           `json:"acceptedAt" format:"date-time"`
	CompletedAt     time.Time                           `json:"completedAt" format:"date-time"`
	WorkerElapsedMs int64                               `json:"workerElapsedMs"`
	WorkerStartedAt time.Time                           `json:"workerStartedAt" format:"date-time"`
	JSON            vaultAssemblyGetResponseTimingsJSON `json:"-"`
}

// vaultAssemblyGetResponseTimingsJSON contains the JSON metadata for the struct
// [VaultAssemblyGetResponseTimings]
type vaultAssemblyGetResponseTimingsJSON struct {
	AcceptedAt      apijson.Field
	CompletedAt     apijson.Field
	WorkerElapsedMs apijson.Field
	WorkerStartedAt apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *VaultAssemblyGetResponseTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyGetResponseTimingsJSON) RawJSON() string {
	return r.raw
}

type VaultAssemblyGetResponseVolume struct {
	BatesEnd   string                             `json:"batesEnd" api:"nullable"`
	BatesStart string                             `json:"batesStart" api:"nullable"`
	Filename   string                             `json:"filename"`
	ObjectID   string                             `json:"objectId"`
	PageCount  int64                              `json:"pageCount"`
	Sha256     string                             `json:"sha256"`
	SizeBytes  int64                              `json:"sizeBytes"`
	JSON       vaultAssemblyGetResponseVolumeJSON `json:"-"`
}

// vaultAssemblyGetResponseVolumeJSON contains the JSON metadata for the struct
// [VaultAssemblyGetResponseVolume]
type vaultAssemblyGetResponseVolumeJSON struct {
	BatesEnd    apijson.Field
	BatesStart  apijson.Field
	Filename    apijson.Field
	ObjectID    apijson.Field
	PageCount   apijson.Field
	Sha256      apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VaultAssemblyGetResponseVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vaultAssemblyGetResponseVolumeJSON) RawJSON() string {
	return r.raw
}

type VaultAssemblyNewParams struct {
	// Source PDFs, or ingested office files (spreadsheets, CSV/TSV, Word), which are
	// assembled from the PDF version ingestion saved for them. A source can also be a
	// PDF you made, such as a cover or notice letter, sent inline. Page ranges refer
	// to a source by its index in this list.
	Sources        param.Field[[]VaultAssemblyNewParamsSourceUnion] `json:"sources" api:"required"`
	Volumes        param.Field[[]VaultAssemblyNewParamsVolume]      `json:"volumes" api:"required"`
	IdempotencyKey param.Field[string]                              `header:"Idempotency-Key" api:"required"`
	// Bates numbering that continues across every volume without its own series. Text
	// must be printable ASCII.
	Bates           param.Field[VaultAssemblyNewParamsBates] `json:"bates"`
	ClientReference param.Field[string]                      `json:"clientReference"`
	// Downsamples colour and grayscale scans above maxDpi and re-encodes them. Images
	// already stored as 1-bit, JPEG 2000, indexed or CMYK keep their bytes.
	Compression param.Field[VaultAssemblyNewParamsCompression] `json:"compression"`
	// Optimize volumes for web viewing
	Linearize param.Field[bool] `json:"linearize"`
	// Put every page upright on a US Letter sheet: portrait pages on a portrait sheet,
	// landscape pages on a landscape sheet, other sizes scaled to fit
	PageSize param.Field[VaultAssemblyNewParamsPageSize] `json:"pageSize"`
}

func (r VaultAssemblyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultAssemblyNewParamsSource struct {
	ObjectID param.Field[string] `json:"objectId"`
	// A base64 PDF of up to 10 pages and 1 MB, assembled as given. Up to 10 per
	// request.
	Pdf       param.Field[string]                                 `json:"pdf" format:"byte"`
	Rendition param.Field[VaultAssemblyNewParamsSourcesRendition] `json:"rendition"`
}

func (r VaultAssemblyNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VaultAssemblyNewParamsSource) implementsVaultAssemblyNewParamsSourceUnion() {}

// Satisfied by [VaultAssemblyNewParamsSourcesObject],
// [VaultAssemblyNewParamsSourcesPdf], [VaultAssemblyNewParamsSource].
type VaultAssemblyNewParamsSourceUnion interface {
	implementsVaultAssemblyNewParamsSourceUnion()
}

type VaultAssemblyNewParamsSourcesObject struct {
	ObjectID  param.Field[string]                                       `json:"objectId" api:"required"`
	Rendition param.Field[VaultAssemblyNewParamsSourcesObjectRendition] `json:"rendition"`
}

func (r VaultAssemblyNewParamsSourcesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VaultAssemblyNewParamsSourcesObject) implementsVaultAssemblyNewParamsSourceUnion() {}

type VaultAssemblyNewParamsSourcesObjectRendition string

const (
	VaultAssemblyNewParamsSourcesObjectRenditionSearchablePdf VaultAssemblyNewParamsSourcesObjectRendition = "searchable_pdf"
	VaultAssemblyNewParamsSourcesObjectRenditionOriginal      VaultAssemblyNewParamsSourcesObjectRendition = "original"
)

func (r VaultAssemblyNewParamsSourcesObjectRendition) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsSourcesObjectRenditionSearchablePdf, VaultAssemblyNewParamsSourcesObjectRenditionOriginal:
		return true
	}
	return false
}

type VaultAssemblyNewParamsSourcesPdf struct {
	// A base64 PDF of up to 10 pages and 1 MB, assembled as given. Up to 10 per
	// request.
	Pdf param.Field[string] `json:"pdf" api:"required" format:"byte"`
}

func (r VaultAssemblyNewParamsSourcesPdf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VaultAssemblyNewParamsSourcesPdf) implementsVaultAssemblyNewParamsSourceUnion() {}

type VaultAssemblyNewParamsSourcesRendition string

const (
	VaultAssemblyNewParamsSourcesRenditionSearchablePdf VaultAssemblyNewParamsSourcesRendition = "searchable_pdf"
	VaultAssemblyNewParamsSourcesRenditionOriginal      VaultAssemblyNewParamsSourcesRendition = "original"
)

func (r VaultAssemblyNewParamsSourcesRendition) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsSourcesRenditionSearchablePdf, VaultAssemblyNewParamsSourcesRenditionOriginal:
		return true
	}
	return false
}

type VaultAssemblyNewParamsVolume struct {
	// Output filename, ending in .pdf
	Filename param.Field[string] `json:"filename" api:"required"`
	// Page ranges in output order. Pages are 1-based and inclusive; repeats are
	// allowed.
	Pages param.Field[[]VaultAssemblyNewParamsVolumesPage] `json:"pages" api:"required"`
	// This volume’s own Bates series, numbered from its own start. Same fields as the
	// top-level bates.
	Bates param.Field[VaultAssemblyNewParamsVolumesBates] `json:"bates"`
	// Adds linked contents pages at the front of the volume, built from its outline.
	// Outline pages stay numbered within the manifest pages.
	Contents param.Field[VaultAssemblyNewParamsVolumesContents] `json:"contents"`
	// Bookmarks, up to 8 levels deep
	Outline param.Field[[]VaultAssemblyNewParamsVolumesOutline] `json:"outline"`
	// PDF document title
	Title param.Field[string] `json:"title"`
}

func (r VaultAssemblyNewParamsVolume) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultAssemblyNewParamsVolumesPage struct {
	End    param.Field[int64] `json:"end" api:"required"`
	Source param.Field[int64] `json:"source" api:"required"`
	Start  param.Field[int64] `json:"start" api:"required"`
	// false leaves these pages unstamped; they use up no Bates number
	Bates param.Field[bool] `json:"bates"`
}

func (r VaultAssemblyNewParamsVolumesPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This volume’s own Bates series, numbered from its own start. Same fields as the
// top-level bates.
type VaultAssemblyNewParamsVolumesBates struct {
	FontSize param.Field[float64] `json:"fontSize"`
	Legend   param.Field[string]  `json:"legend"`
	// Points from the bottom edge to the label
	MarginBottom param.Field[float64] `json:"marginBottom"`
	// Points from the side edge to the label
	MarginSide param.Field[float64]                                    `json:"marginSide"`
	PadTo      param.Field[int64]                                      `json:"padTo"`
	Position   param.Field[VaultAssemblyNewParamsVolumesBatesPosition] `json:"position"`
	Prefix     param.Field[string]                                     `json:"prefix"`
	// Shrinks each stamped page slightly so the label sits in a clear strip below its
	// content
	Shrink param.Field[bool]   `json:"shrink"`
	Start  param.Field[int64]  `json:"start"`
	Suffix param.Field[string] `json:"suffix"`
}

func (r VaultAssemblyNewParamsVolumesBates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultAssemblyNewParamsVolumesBatesPosition string

const (
	VaultAssemblyNewParamsVolumesBatesPositionBottomRight  VaultAssemblyNewParamsVolumesBatesPosition = "bottom_right"
	VaultAssemblyNewParamsVolumesBatesPositionBottomCenter VaultAssemblyNewParamsVolumesBatesPosition = "bottom_center"
	VaultAssemblyNewParamsVolumesBatesPositionBottomLeft   VaultAssemblyNewParamsVolumesBatesPosition = "bottom_left"
)

func (r VaultAssemblyNewParamsVolumesBatesPosition) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsVolumesBatesPositionBottomRight, VaultAssemblyNewParamsVolumesBatesPositionBottomCenter, VaultAssemblyNewParamsVolumesBatesPositionBottomLeft:
		return true
	}
	return false
}

// Adds linked contents pages at the front of the volume, built from its outline.
// Outline pages stay numbered within the manifest pages.
type VaultAssemblyNewParamsVolumesContents struct {
	// Outline levels listed
	Depth param.Field[VaultAssemblyNewParamsVolumesContentsDepth] `json:"depth"`
	Title param.Field[string]                                     `json:"title"`
}

func (r VaultAssemblyNewParamsVolumesContents) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outline levels listed
type VaultAssemblyNewParamsVolumesContentsDepth int64

const (
	VaultAssemblyNewParamsVolumesContentsDepth1 VaultAssemblyNewParamsVolumesContentsDepth = 1
	VaultAssemblyNewParamsVolumesContentsDepth2 VaultAssemblyNewParamsVolumesContentsDepth = 2
)

func (r VaultAssemblyNewParamsVolumesContentsDepth) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsVolumesContentsDepth1, VaultAssemblyNewParamsVolumesContentsDepth2:
		return true
	}
	return false
}

type VaultAssemblyNewParamsVolumesOutline struct {
	// 1-based page within this volume
	Page  param.Field[int64]  `json:"page" api:"required"`
	Title param.Field[string] `json:"title" api:"required"`
	// Nested entries, same shape
	Children param.Field[[]interface{}] `json:"children"`
}

func (r VaultAssemblyNewParamsVolumesOutline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Bates numbering that continues across every volume without its own series. Text
// must be printable ASCII.
type VaultAssemblyNewParamsBates struct {
	FontSize param.Field[float64] `json:"fontSize"`
	// Line printed above each label
	Legend param.Field[string] `json:"legend"`
	// Points from the bottom edge to the label
	MarginBottom param.Field[float64] `json:"marginBottom"`
	// Points from the side edge to the label
	MarginSide param.Field[float64]                             `json:"marginSide"`
	PadTo      param.Field[int64]                               `json:"padTo"`
	Position   param.Field[VaultAssemblyNewParamsBatesPosition] `json:"position"`
	Prefix     param.Field[string]                              `json:"prefix"`
	// Shrinks each stamped page slightly so the label sits in a clear strip below its
	// content instead of over existing footer text
	Shrink param.Field[bool]   `json:"shrink"`
	Start  param.Field[int64]  `json:"start"`
	Suffix param.Field[string] `json:"suffix"`
}

func (r VaultAssemblyNewParamsBates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VaultAssemblyNewParamsBatesPosition string

const (
	VaultAssemblyNewParamsBatesPositionBottomRight  VaultAssemblyNewParamsBatesPosition = "bottom_right"
	VaultAssemblyNewParamsBatesPositionBottomCenter VaultAssemblyNewParamsBatesPosition = "bottom_center"
	VaultAssemblyNewParamsBatesPositionBottomLeft   VaultAssemblyNewParamsBatesPosition = "bottom_left"
)

func (r VaultAssemblyNewParamsBatesPosition) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsBatesPositionBottomRight, VaultAssemblyNewParamsBatesPositionBottomCenter, VaultAssemblyNewParamsBatesPositionBottomLeft:
		return true
	}
	return false
}

// Downsamples colour and grayscale scans above maxDpi and re-encodes them. Images
// already stored as 1-bit, JPEG 2000, indexed or CMYK keep their bytes.
type VaultAssemblyNewParamsCompression struct {
	// Store pages that are only ink on paper as 1-bit images
	Bitonal param.Field[bool] `json:"bitonal"`
	// Clean up scanned document pages: 300 dpi, white paper, dark ink. Photos and
	// imaging keep their tones; gray stays gray and colour stays colour
	CleanScans param.Field[bool] `json:"cleanScans"`
	// JPEG 2000 is smaller; JPEG is faster
	Encoding param.Field[VaultAssemblyNewParamsCompressionEncoding] `json:"encoding"`
	MaxDpi   param.Field[int64]                                     `json:"maxDpi"`
	Quality  param.Field[VaultAssemblyNewParamsCompressionQuality]  `json:"quality"`
}

func (r VaultAssemblyNewParamsCompression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JPEG 2000 is smaller; JPEG is faster
type VaultAssemblyNewParamsCompressionEncoding string

const (
	VaultAssemblyNewParamsCompressionEncodingJpeg2000 VaultAssemblyNewParamsCompressionEncoding = "jpeg2000"
	VaultAssemblyNewParamsCompressionEncodingJpeg     VaultAssemblyNewParamsCompressionEncoding = "jpeg"
)

func (r VaultAssemblyNewParamsCompressionEncoding) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsCompressionEncodingJpeg2000, VaultAssemblyNewParamsCompressionEncodingJpeg:
		return true
	}
	return false
}

type VaultAssemblyNewParamsCompressionQuality string

const (
	VaultAssemblyNewParamsCompressionQualityHigh     VaultAssemblyNewParamsCompressionQuality = "high"
	VaultAssemblyNewParamsCompressionQualityStandard VaultAssemblyNewParamsCompressionQuality = "standard"
	VaultAssemblyNewParamsCompressionQualitySmall    VaultAssemblyNewParamsCompressionQuality = "small"
)

func (r VaultAssemblyNewParamsCompressionQuality) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsCompressionQualityHigh, VaultAssemblyNewParamsCompressionQualityStandard, VaultAssemblyNewParamsCompressionQualitySmall:
		return true
	}
	return false
}

// Put every page upright on a US Letter sheet: portrait pages on a portrait sheet,
// landscape pages on a landscape sheet, other sizes scaled to fit
type VaultAssemblyNewParamsPageSize string

const (
	VaultAssemblyNewParamsPageSizeLetter VaultAssemblyNewParamsPageSize = "letter"
)

func (r VaultAssemblyNewParamsPageSize) IsKnown() bool {
	switch r {
	case VaultAssemblyNewParamsPageSizeLetter:
		return true
	}
	return false
}
