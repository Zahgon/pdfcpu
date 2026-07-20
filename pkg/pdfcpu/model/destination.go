package model

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"

type DestinationType int

const (
	DestXYZ DestinationType = iota
	DestFit
	DestFitH
	DestFitV
	DestFitR
	DestFitB
	DestFitBH
	DestFitBV
)

var DestinationTypeStrings = map[DestinationType]string{
	DestXYZ:   "XYZ",
	DestFit:   "Fit",
	DestFitH:  "FitH",
	DestFitV:  "FitV",
	DestFitR:  "FitR",
	DestFitB:  "FitB",
	DestFitBH: "FitBH",
	DestFitBV: "FitBV",
}

type Destination struct {
	Typ                      DestinationType
	PageNr                   int
	Left, Bottom, Right, Top int
	Zoom                     float32
}

func (dest Destination) String() string { _ = "STUB: not implemented"; return "" }

func (dest Destination) Name() types.Name { _ = "STUB: not implemented"; return *new(types.Name) }

func (dest Destination) Array(indRef types.IndirectRef) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}
