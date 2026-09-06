package api

import (
	"io"
	"sync"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func logDisclaimerPDF20() { _ = "STUB: not implemented"; return }

func ReadContext(rs io.ReadSeeker, conf *model.Configuration) (ctx *model.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadContextFile(inFile string) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateContext(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func OptimizeContext(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func PatchFile(fileName string, bb []byte, offset int64) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteContext(ctx *model.Context, w io.Writer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func WriteIncrement(ctx *model.Context, w io.Writer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func WriteContextFile(ctx *model.Context, outFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ReadAndValidate(rs io.ReadSeeker, conf *model.Configuration) (ctx *model.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cmdAssumingOptimization(cmd model.CommandMode) bool { _ = "STUB: not implemented"; return false }

func ReadValidateAndOptimize(rs io.ReadSeeker, conf *model.Configuration) (ctx *model.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logWritingTo(s string) { _ = "STUB: not implemented"; return }

func Write(ctx *model.Context, w io.Writer, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteIncr(ctx *model.Context, rws io.ReadWriteSeeker, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureDefaultConfigAt(path string) error { _ = "STUB: not implemented"; return nil }

var (
	mutexDisableConfigDir sync.Mutex
)

func DisableConfigDir() { _ = "STUB: not implemented"; return }

func LoadConfiguration() *model.Configuration { _ = "STUB: not implemented"; return nil }
