package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	RootVersion = iota
	RootExtensions
	RootPageLabels
	RootNames
	RootDests
	RootViewerPrefs
	RootPageLayout
	RootPageMode
	RootOutlines
	RootThreads
	RootOpenAction
	RootAA
	RootURI
	RootAcroForm
	RootMetadata
	RootStructTreeRoot
	RootMarkInfo
	RootLang
	RootSpiderInfo
	RootOutputIntents
	RootPieceInfo
	RootOCProperties
	RootPerms
	RootLegal
	RootRequirements
	RootCollection
	RootNeedsRendering
)

const (
	PageLastModified = iota
	PageResources
	PageMediaBox
	PageCropBox
	PageBleedBox
	PageTrimBox
	PageArtBox
	PageBoxColorInfo
	PageContents
	PageRotate
	PageGroup
	PageThumb
	PageB
	PageDur
	PageTrans
	PageAnnots
	PageAA
	PageMetadata
	PagePieceInfo
	PageStructParents
	PageID
	PagePZ
	PageSeparationInfo
	PageTabs
	PageTemplateInstantiated
	PagePresSteps
	PageUserUnit
	PageVP
)

type PDFStats struct {
	rootAttrs types.IntSet

	pageAttrs types.IntSet
}

func NewPDFStats() PDFStats { _ = "STUB: not implemented"; return *new(PDFStats) }

func (stats PDFStats) AddRootAttr(name int) { _ = "STUB: not implemented"; return }

func (stats PDFStats) UsesRootAttr(name int) bool { _ = "STUB: not implemented"; return false }

func (stats PDFStats) AddPageAttr(name int) { _ = "STUB: not implemented"; return }

func (stats PDFStats) UsesPageAttr(name int) bool { _ = "STUB: not implemented"; return false }

func ValidationTimingStats(dur1, dur2, dur float64) { _ = "STUB: not implemented"; return }

func TimingStats(op string, durRead, durVal, durOpt, durWrite, durTotal float64) {
	_ = "STUB: not implemented"
	return
}
