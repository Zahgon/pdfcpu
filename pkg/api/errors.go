package api

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/font"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

var (
	ErrAttachmentOutputCollision = errors.New("attachment output collision")

	ErrBookletImageOutputConflict = errors.New("booklet image output aliases input")

	ErrCircularBookmarks = pdfcpu.ErrCircularBookmarks

	ErrDuplicateCertificateDestination = errors.New("duplicate certificate destination")

	ErrDuplicatePostScriptName = font.ErrDuplicatePostScriptName

	ErrExistingBookmarks = pdfcpu.ErrExistingBookmarks

	ErrGridImageOutputConflict = errors.New("grid image output aliases input")

	ErrImportImagesOutputConflict = errors.New("import images output aliases image input")

	ErrInvalidBookmark = pdfcpu.ErrInvalidBookmark

	ErrInvalidBookmarkJSON = pdfcpu.ErrInvalidBookmarkJSON

	ErrInvalidCSV = errors.New("invalid csv input file")

	ErrInvalidCutConfiguration = errors.New("invalid cut configuration")

	ErrInvalidFormData = errors.New("invalid form data")

	ErrInvalidImageSelection = errors.New("invalid image selection")

	ErrInvalidImportConfiguration = errors.New("invalid import configuration")

	ErrInvalidJSON = errors.New("invalid JSON encoding")

	ErrInvalidPageBoundaries = errors.New("invalid page boundaries")

	ErrInvalidPageConfiguration = errors.New("invalid page configuration")

	ErrInvalidPageLayout = errors.New("invalid page layout")

	ErrInvalidPageMode = errors.New("invalid page mode")

	ErrInvalidRotation = errors.New("invalid rotation")

	ErrInvalidResizeConfiguration = errors.New("invalid resize configuration")

	ErrInvalidSplitPageNumberSequence = errors.New("invalid split page number sequence")

	ErrInvalidSplitSpan = errors.New("invalid split span")

	ErrInvalidUnicodePlane = errors.New("invalid Unicode plane")

	ErrInvalidZoomConfiguration = errors.New("invalid zoom configuration")

	ErrMissingAnnotation = pdfcpu.ErrMissingAnnotation

	ErrMissingBookmarks = errors.New("missing bookmarks")

	ErrMissingBookletConfiguration = errors.New("missing booklet configuration")

	ErrMissingBoxConfiguration = errors.New("missing box configuration")

	ErrMissingCertificate = pdfcpu.ErrMissingCertificate

	ErrMissingCertificateInput = errors.New("missing certificate input")

	ErrMissingConfiguration = errors.New("missing configuration")

	ErrMissingCutConfiguration = errors.New("missing cut configuration")

	ErrMissingDigestFunction = errors.New("missing digest function")

	ErrMissingFontInput = errors.New("missing font input")

	ErrMissingFormInput = errors.New("missing form input")

	ErrMissingGridConfiguration = errors.New("missing grid configuration")

	ErrMissingImageInput = errors.New("missing image input")

	ErrMissingImageReader = pdfcpu.ErrMissingImageReader

	ErrMissingJSONInput = errors.New("missing JSON input")

	ErrMissingJSONOutput = errors.New("missing JSON output")

	ErrMissingJSONReader = errors.New("missing JSON reader")

	ErrMissingJSONWriter = errors.New("missing JSON writer")

	ErrMissingNUpConfiguration = errors.New("missing n-up configuration")

	ErrMissingPageBoundaries = errors.New("missing page boundaries")

	ErrMissingPDFContext = pdfcpu.ErrMissingPDFContext

	ErrMissingPDFInput = errors.New("missing PDF input")

	ErrMissingPDFOutput = errors.New("missing PDF output")

	ErrMissingPDFReadSeeker = errors.New("missing PDF read seeker")

	ErrMissingPDFReadWriteSeeker = errors.New("missing PDF read write seeker")

	ErrMissingPDFWriter = errors.New("missing PDF writer")

	ErrMissingReader = pdfcpu.ErrMissingReader

	ErrMissingResizeConfiguration = errors.New("missing resize configuration")

	ErrMissingSplitPageNumbers = errors.New("missing split page numbers")

	ErrMissingWatermarkConfiguration = pdfcpu.ErrMissingWatermarkConfiguration

	ErrMissingWatermarks = pdfcpu.ErrMissingWatermarks

	ErrMissingXRefTable = pdfcpu.ErrMissingXRefTable

	ErrMissingZoomConfiguration = errors.New("missing zoom configuration")

	ErrNoAttachmentAdded = errors.New("no attachment added")

	ErrNoAttachmentRemoved = errors.New("no attachment removed")

	ErrNoBookmarks = pdfcpu.ErrNoBookmarks

	ErrNoCertificates = pdfcpu.ErrNoCertificates

	ErrNoFormData = errors.New("missing form data")

	ErrNoFormFieldsAffected = errors.New("no form fields affected")

	ErrNoKeywordRemoved = errors.New("no keyword removed")

	ErrNoOp = errors.New("no operation")

	ErrNoPropertyRemoved = errors.New("no property removed")

	ErrNoSignatures = pdfcpu.ErrNoSignatures

	ErrNoOutlines = ErrNoBookmarks

	ErrNUpImageOutputConflict = errors.New("n-up image output aliases input")

	ErrOutlines = ErrExistingBookmarks

	ErrUnknownFont = font.ErrUnknownFont

	ErrUnsupportedCertificateFile = pdfcpu.ErrUnsupportedCertificateFile

	ErrUnsupportedFontFile = errors.New("unsupported font file")

	ErrUnsupportedFormDataFormat = errors.New("unsupported data format")

	ErrUpdateImagesOutputConflict = errors.New("update images output aliases image input")

	ErrUserFontNotFound = font.ErrUnknownFont
)
