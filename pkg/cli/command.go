package cli

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type Command struct {
	Mode              model.CommandMode
	InFile            *string
	InFileCert        *string
	InFilePrivateKey  *string
	InFileJSON        *string
	InFiles           []string
	InDir             *string
	OutFile           *string
	OutFileJSON       *string
	OutDir            *string
	PageSelection     []string
	PWOld             *string
	PWNew             *string
	StringVal         string
	IntVal            int
	BoolVal1          bool
	BoolVal2          bool
	BoolVal3          bool
	IntVals           []int
	StringVals        []string
	StringMap         map[string]string
	Input             io.ReadSeeker
	Inputs            []io.ReadSeeker
	Output            io.Writer
	ErrorOutput       io.Writer
	Box               *model.Box
	Import            *pdfcpu.Import
	NUp               *model.NUp
	Cut               *model.Cut
	PageBoundaries    *model.PageBoundaries
	Resize            *model.Resize
	Zoom              *model.Zoom
	Watermark         *model.Watermark
	ViewerPreferences *model.ViewerPreferences
	PageConf          *pdfcpu.PageConfiguration
	Conf              *model.Configuration
}
