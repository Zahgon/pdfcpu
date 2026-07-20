package filter

import (
	"bytes"
	"io"
)

const (
	PredictorNo      = 1
	PredictorTIFF    = 2
	PredictorNone    = 10
	PredictorSub     = 11
	PredictorUp      = 12
	PredictorAverage = 13
	PredictorPaeth   = 14
	PredictorOptimum = 15
)

const (
	PNGNone    = 0x00
	PNGSub     = 0x01
	PNGUp      = 0x02
	PNGAverage = 0x03
	PNGPaeth   = 0x04
)

type flate struct {
	baseFilter
}

func (f flate) Encode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f flate) Decode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f flate) DecodeLength(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f flate) passThru(rin io.Reader, maxLen int64) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logUnexpectedEOFFlateDecode() { _ = "STUB: not implemented"; return }

func intMemberOf(i int, list []int) bool { _ = "STUB: not implemented"; return false }

func validatePredictor(predictor int) error { _ = "STUB: not implemented"; return nil }

func predictorRowParams(predictor, colors, bpc, columns int) (rowSize, rowLen, bytesPerPixel int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func applyHorDiff(row []byte, colors int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processRow(pr, cr []byte, p, colors, bytesPerPixel int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f flate) parameters() (colors, bpc, columns int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func checkBufLen(b bytes.Buffer, maxLen int64) bool { _ = "STUB: not implemented"; return false }

func process(w io.Writer, pr, cr []byte, predictor, colors, bytesPerPixel int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f flate) decodePostProcessRows(r io.Reader, maxLen int64, m, predictor, colors, bytesPerPixel int) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f flate) decodePostProcess(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
