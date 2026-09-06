package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func WatermarkContext(ctx *model.Context, selectedPages types.IntSet, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func watermarkOperation(wm *model.Watermark) string { _ = "STUB: not implemented"; return "" }

func AddWatermarksMap(rs io.ReadSeeker, w io.Writer, m map[int]*model.Watermark, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func sortedWatermarkMapPages[T any](m map[int]T) []int { _ = "STUB: not implemented"; return nil }

func validateWatermarkMap(m map[int]*model.Watermark) error { _ = "STUB: not implemented"; return nil }

func validateWatermarkSliceMap(m map[int][]*model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarksMapFile(inFile, outFile string, m map[int]*model.Watermark, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarksSliceMap(rs io.ReadSeeker, w io.Writer, m map[int][]*model.Watermark, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarksSliceMapFile(inFile, outFile string, m map[int][]*model.Watermark, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarks(rs io.ReadSeeker, w io.Writer, selectedPages []string, wm *model.Watermark, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarksFile(inFile, outFile string, selectedPages []string, wm *model.Watermark, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveWatermarks(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveWatermarksFile(inFile, outFile string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func HasWatermarks(rs io.ReadSeeker, conf *model.Configuration) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func HasWatermarksFile(inFile string, conf *model.Configuration) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ImageWatermarkForReader(rd io.Reader, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFWatermarkForReadSeeker(rs io.ReadSeeker, pageNrSrc int, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFMultiWatermarkForReadSeeker(rs io.ReadSeeker, startPageNrSrc, startPageNrDest int, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseWatermark(mode int, modeParm, desc string, onTop bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func watermark(mode int, modeParm, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TextWatermark(text, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageWatermark(fileName, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFWatermark(fileName, desc string, onTop, update bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddTextWatermarksFile(inFile, outFile string, selectedPages []string, onTop bool, text, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func AddImageWatermarksFile(inFile, outFile string, selectedPages []string, onTop bool, fileName, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func AddImageWatermarksForReaderFile(inFile, outFile string, selectedPages []string, onTop bool, r io.Reader, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPDFWatermarksFile(inFile, outFile string, selectedPages []string, onTop bool, fileName, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPDFWatermarksForReadSeekerFile(inFile, outFile string, selectedPages []string, onTop bool, rs io.ReadSeeker, pageNrSrc int, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateTextWatermarksFile(inFile, outFile string, selectedPages []string, onTop bool, text, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateImageWatermarksFile(inFile, outFile string, selectedPages []string, onTop bool, fileName, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdatePDFWatermarksFile(inFile, outFile string, selectedPages []string, onTop bool, fileName, desc string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
