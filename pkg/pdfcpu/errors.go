package pdfcpu

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

var ErrUnsupportedVersion = errors.New("PDF 2.0 unsupported for this operation")

var ErrUnsupportedResource = errors.New("unsupported resource")

var ErrNoSignatures = errors.New("no signatures present")

var ErrMissingPDFContext = model.ErrMissingPDFContext

var ErrMissingPDFInfo = errors.New("missing PDF info")

var ErrMissingReadContext = errors.New("missing PDF read context")

var ErrMissingWriteContext = errors.New("missing PDF write context")

var ErrMissingXRefTable = model.ErrMissingXRefTable

var ErrMissingStreamDict = errors.New("missing PDF stream dictionary")

var ErrMissingOptimizationContext = errors.New("missing optimization context")

var ErrMissingPageNumbers = errors.New("missing page numbers")

var ErrMissingAnnotation = errors.New("missing annotation")

var ErrMissingReader = errors.New("missing reader")

var ErrMissingImageReader = model.ErrMissingImageReader

var ErrMissingWatermarkConfiguration = errors.New("missing watermark configuration")

var ErrMissingWatermarks = errors.New("missing watermarks")

var ErrInvalidPageNumber = errors.New("invalid page number")
