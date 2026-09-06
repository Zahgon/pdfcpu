package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type PageMode int

const (
	PageModeUseNone PageMode = iota
	PageModeUseOutlines
	PageModeUseThumbs
	PageModeFullScreen
	PageModeUseOC
	PageModeUseAttachments
)

func PageModeFor(s string) *PageMode { _ = "STUB: not implemented"; return nil }

func (pm *PageMode) String() string { _ = "STUB: not implemented"; return "" }

type PageLayout int

const (
	PageLayoutSinglePage PageLayout = iota
	PageLayoutTwoColumnLeft
	PageLayoutTwoColumnRight
	PageLayoutTwoPageLeft
	PageLayoutTwoPageRight
	PageLayoutOneColumn
)

func PageLayoutFor(s string) *PageLayout { _ = "STUB: not implemented"; return nil }

func (pl *PageLayout) String() string { _ = "STUB: not implemented"; return "" }

type NonFullScreenPageMode PageMode

const (
	NFSPageModeUseNone NonFullScreenPageMode = iota
	NFSPageModeUseOutlines
	NFSPageModeUseThumb
	NFSPageModeUseOC
)

type PageBoundary int

const (
	MediaBox PageBoundary = iota
	CropBox
	TrimBox
	BleedBox
	ArtBox
)

func PageBoundaryFor(s string) *PageBoundary { _ = "STUB: not implemented"; return nil }

func (pb *PageBoundary) String() string { _ = "STUB: not implemented"; return "" }

type PrintScaling int

const (
	PrintScalingNone PrintScaling = iota
	PrintScalingAppDefault
)

func PrintScalingFor(s string) *PrintScaling { _ = "STUB: not implemented"; return nil }

func (ps *PrintScaling) String() string { _ = "STUB: not implemented"; return "" }

type Direction int

const (
	L2R Direction = iota
	R2L
)

func DirectionFor(s string) *Direction { _ = "STUB: not implemented"; return nil }

func (d *Direction) String() string { _ = "STUB: not implemented"; return "" }

type PaperHandling int

const (
	Simplex PaperHandling = iota
	DuplexFlipShortEdge
	DuplexFlipLongEdge
)

func PaperHandlingFor(s string) *PaperHandling { _ = "STUB: not implemented"; return nil }

func (ph *PaperHandling) String() string { _ = "STUB: not implemented"; return "" }

type ViewerPreferences struct {
	HideToolbar           *bool
	HideMenubar           *bool
	HideWindowUI          *bool
	FitWindow             *bool
	CenterWindow          *bool
	DisplayDocTitle       *bool
	NonFullScreenPageMode *NonFullScreenPageMode
	Direction             *Direction
	ViewArea              *PageBoundary
	ViewClip              *PageBoundary
	PrintArea             *PageBoundary
	PrintClip             *PageBoundary
	PrintScaling          *PrintScaling
	Duplex                *PaperHandling
	PickTrayByPDFSize     *bool
	PrintPageRange        types.Array
	NumCopies             *types.Integer
	Enforce               types.Array
}

func (vp *ViewerPreferences) validatePrinterPreferences(version Version) error {
	_ = "STUB: not implemented"
	return nil
}

func (vp *ViewerPreferences) Validate(version Version) error { _ = "STUB: not implemented"; return nil }

func (vp *ViewerPreferences) SetHideToolBar(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetHideMenuBar(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetHideWindowUI(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetFitWindow(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetCenterWindow(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetDisplayDocTitle(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetPickTrayByPDFSize(val bool) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) SetNumCopies(i int) { _ = "STUB: not implemented"; return }

func (vp *ViewerPreferences) populatePrinterPreferences(vp1 *ViewerPreferences) {
	_ = "STUB: not implemented"
	return
}

func (vp *ViewerPreferences) Populate(vp1 *ViewerPreferences) { _ = "STUB: not implemented"; return }

func DefaultViewerPreferences(version Version) *ViewerPreferences {
	_ = "STUB: not implemented"
	return nil
}

func ViewerPreferencesWithDefaults(vp *ViewerPreferences, version Version) (*ViewerPreferences, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ViewerPrefJSON struct {
	HideToolbar           *bool    `json:"hideToolbar,omitempty"`
	HideMenubar           *bool    `json:"hideMenubar,omitempty"`
	HideWindowUI          *bool    `json:"hideWindowUI,omitempty"`
	FitWindow             *bool    `json:"fitWindow,omitempty"`
	CenterWindow          *bool    `json:"centerWindow,omitempty"`
	DisplayDocTitle       *bool    `json:"displayDocTitle,omitempty"`
	NonFullScreenPageMode string   `json:"nonFullScreenPageMode,omitempty"`
	Direction             string   `json:"direction,omitempty"`
	ViewArea              string   `json:"viewArea,omitempty"`
	ViewClip              string   `json:"viewClip,omitempty"`
	PrintArea             string   `json:"printArea,omitempty"`
	PrintClip             string   `json:"printClip,omitempty"`
	PrintScaling          string   `json:"printScaling,omitempty"`
	Duplex                string   `json:"duplex,omitempty"`
	PickTrayByPDFSize     *bool    `json:"pickTrayByPDFSize,omitempty"`
	PrintPageRange        []int    `json:"printPageRange,omitempty"`
	NumCopies             *int     `json:"numCopies,omitempty"`
	Enforce               []string `json:"enforce,omitempty"`
}

func (vp *ViewerPreferences) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vp *ViewerPreferences) unmarshalPrintPageRange(vpJSON ViewerPrefJSON) error {
	_ = "STUB: not implemented"
	return nil
}

func (vp *ViewerPreferences) unmarshalPrinterPreferences(vpJSON ViewerPrefJSON) error {
	_ = "STUB: not implemented"
	return nil
}

func (vp *ViewerPreferences) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func renderViewerFlags(vp ViewerPreferences, ss *[]string) { _ = "STUB: not implemented"; return }

func listViewerFlags(vp ViewerPreferences, ss *[]string) { _ = "STUB: not implemented"; return }

func (vp ViewerPreferences) listPrinterPreferences() []string {
	_ = "STUB: not implemented"
	return nil
}

func (vp ViewerPreferences) List() []string { _ = "STUB: not implemented"; return nil }

func (vp ViewerPreferences) String() string { _ = "STUB: not implemented"; return "" }
