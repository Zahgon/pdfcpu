package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type AnnotationFlags int

const (
	AnnInvisible AnnotationFlags = 1 << iota
	AnnHidden
	AnnPrint
	AnnNoZoom
	AnnNoRotate
	AnnNoView
	AnnReadOnly
	AnnLocked
	AnnToggleNoView
	AnnLockedContents
)

type AnnotationType int

const (
	AnnText AnnotationType = iota
	AnnLink
	AnnFreeText
	AnnLine
	AnnSquare
	AnnCircle
	AnnPolygon
	AnnPolyLine
	AnnHighLight
	AnnUnderline
	AnnSquiggly
	AnnStrikeOut
	AnnStamp
	AnnCaret
	AnnInk
	AnnPopup
	AnnFileAttachment
	AnnSound
	AnnMovie
	AnnWidget
	AnnScreen
	AnnPrinterMark
	AnnTrapNet
	AnnWatermark
	Ann3D
	AnnRedact
	AnnCustom
)

var AnnotTypes = map[string]AnnotationType{
	"Text":           AnnText,
	"Link":           AnnLink,
	"FreeText":       AnnFreeText,
	"Line":           AnnLine,
	"Square":         AnnSquare,
	"Circle":         AnnCircle,
	"Polygon":        AnnPolygon,
	"PolyLine":       AnnPolyLine,
	"Highlight":      AnnHighLight,
	"Underline":      AnnUnderline,
	"Squiggly":       AnnSquiggly,
	"StrikeOut":      AnnStrikeOut,
	"Stamp":          AnnStamp,
	"Caret":          AnnCaret,
	"Ink":            AnnInk,
	"Popup":          AnnPopup,
	"FileAttachment": AnnFileAttachment,
	"Sound":          AnnSound,
	"Movie":          AnnMovie,
	"Widget":         AnnWidget,
	"Screen":         AnnScreen,
	"PrinterMark":    AnnPrinterMark,
	"TrapNet":        AnnTrapNet,
	"Watermark":      AnnWatermark,
	"3D":             Ann3D,
	"Redact":         AnnRedact,
	"Custom":         AnnCustom,
}

var AnnotTypeStrings = map[AnnotationType]string{
	AnnText:           "Text",
	AnnLink:           "Link",
	AnnFreeText:       "FreeText",
	AnnLine:           "Line",
	AnnSquare:         "Square",
	AnnCircle:         "Circle",
	AnnPolygon:        "Polygon",
	AnnPolyLine:       "PolyLine",
	AnnHighLight:      "Highlight",
	AnnUnderline:      "Underline",
	AnnSquiggly:       "Squiggly",
	AnnStrikeOut:      "StrikeOut",
	AnnStamp:          "Stamp",
	AnnCaret:          "Caret",
	AnnInk:            "Ink",
	AnnPopup:          "Popup",
	AnnFileAttachment: "FileAttachment",
	AnnSound:          "Sound",
	AnnMovie:          "Movie",
	AnnWidget:         "Widget",
	AnnScreen:         "Screen",
	AnnPrinterMark:    "PrinterMark",
	AnnTrapNet:        "TrapNet",
	AnnWatermark:      "Watermark",
	Ann3D:             "3D",
	AnnRedact:         "Redact",
	AnnCustom:         "Custom",
}

type BorderStyle int

const (
	BSSolid BorderStyle = iota
	BSDashed
	BSBeveled
	BSInset
	BSUnderline
)

func borderStyleDict(width float64, style BorderStyle) types.Dict {
	_ = "STUB: not implemented"
	return *new(types.Dict)
}

func borderEffectDict(cloudyBorder bool, intensity int) types.Dict {
	_ = "STUB: not implemented"
	return *new(types.Dict)
}

func borderArray(rx, ry, width float64) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}

type LineEndingStyle int

const (
	LESquare LineEndingStyle = iota
	LECircle
	LEDiamond
	LEOpenArrow
	LEClosedArrow
	LENone
	LEButt
	LEROpenArrow
	LERClosedArrow
	LESlash
)

func LineEndingStyleName(les LineEndingStyle) string { _ = "STUB: not implemented"; return "" }

type AnnotationRenderer interface {
	RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error)
	Type() AnnotationType
	Rectangle() types.Rectangle
	RectString() string
	APObjNrInt() int
	ID() string
	Content() string
	ContentString() string
	CustomTypeString() string
}

type Annotation struct {
	SubType          AnnotationType
	CustomSubType    string
	Rect             types.Rectangle
	APObjNr          int
	Contents         string
	NM               string
	ModificationDate string
	P                *types.IndirectRef
	F                AnnotationFlags
	C                *color.SimpleColor
	BorderRadX       float64
	BorderRadY       float64
	BorderWidth      float64
	Hash             uint32
}

func NewAnnotation(
	typ AnnotationType,
	customTyp string,
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64) Annotation {
	_ = "STUB: not implemented"
	return *new(Annotation)
}

func NewAnnotationForRawType(
	typ string,
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,

	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64) Annotation {
	_ = "STUB: not implemented"
	return *new(Annotation)
}

func (ann Annotation) ID() string { _ = "STUB: not implemented"; return "" }

func (ann Annotation) ContentString() string { _ = "STUB: not implemented"; return "" }

func (ann Annotation) Content() string { _ = "STUB: not implemented"; return "" }

func (ann Annotation) CustomTypeString() string { _ = "STUB: not implemented"; return "" }

func (ann Annotation) RectString() string { _ = "STUB: not implemented"; return "" }

func (ann Annotation) Rectangle() types.Rectangle {
	_ = "STUB: not implemented"
	return *new(types.Rectangle)
}

func (ann Annotation) APObjNrInt() int { _ = "STUB: not implemented"; return 0 }

func (ann Annotation) Type() AnnotationType { _ = "STUB: not implemented"; return *new(AnnotationType) }

func (ann Annotation) TypeString() string { _ = "STUB: not implemented"; return "" }

func (ann Annotation) HashString() uint32 { _ = "STUB: not implemented"; return 0 }

func (ann Annotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type PopupAnnotation struct {
	Annotation
	ParentIndRef *types.IndirectRef
	Open         bool
}

func NewPopupAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,

	parentIndRef *types.IndirectRef,
	displayOpen bool) PopupAnnotation {
	_ = "STUB: not implemented"
	return *new(PopupAnnotation)
}

func (ann PopupAnnotation) ContentString() string { _ = "STUB: not implemented"; return "" }

func (ann PopupAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type LinkAnnotation struct {
	Annotation
	Dest        *Destination
	URI         string
	Quad        types.QuadPoints
	Border      bool
	BorderWidth float64
	BorderStyle BorderStyle
}

func NewLinkAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	borderCol *color.SimpleColor,

	dest *Destination,
	uri string,
	quad types.QuadPoints,
	border bool,
	borderWidth float64,
	borderStyle BorderStyle) LinkAnnotation {
	_ = "STUB: not implemented"
	return *new(LinkAnnotation)
}

func (ann LinkAnnotation) ContentString() string { _ = "STUB: not implemented"; return "" }

func (ann LinkAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type MarkupAnnotation struct {
	Annotation
	T            string
	PopupIndRef  *types.IndirectRef
	CA           *float64
	RC           string
	CreationDate string
	Subj         string
}

func NewMarkupAnnotation(
	subType AnnotationType,
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,

	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string) MarkupAnnotation {
	_ = "STUB: not implemented"
	return *new(MarkupAnnotation)
}

func (ann MarkupAnnotation) ContentString() string { _ = "STUB: not implemented"; return "" }

func (ann MarkupAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type TextAnnotation struct {
	MarkupAnnotation
	Open bool
	Name string
}

func NewTextAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,

	displayOpen bool,
	name string) TextAnnotation {
	_ = "STUB: not implemented"
	return *new(TextAnnotation)
}

func (ann TextAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type FreeTextIntent int

const (
	IntentFreeText FreeTextIntent = 1 << iota
	IntentFreeTextCallout
	IntentFreeTextTypeWriter
)

func FreeTextIntentName(fti FreeTextIntent) string { _ = "STUB: not implemented"; return "" }

type FreeTextAnnotation struct {
	MarkupAnnotation
	Text                   string
	HAlign                 types.HAlignment
	FontName               string
	FontSize               int
	FontCol                *color.SimpleColor
	DS                     string
	Intent                 string
	CallOutLine            types.Array
	CallOutLineEndingStyle string
	Margins                types.Array
	BorderWidth            float64
	BorderStyle            BorderStyle
	CloudyBorder           bool
	CloudyBorderIntensity  int
}

func NewFreeTextAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	text string,
	hAlign types.HAlignment,
	fontName string,
	fontSize int,
	fontCol *color.SimpleColor,
	ds string,
	intent *FreeTextIntent,
	callOutLine types.Array,
	callOutLineEndingStyle *LineEndingStyle,
	MLeft, MTop, MRight, MBot float64,
	borderWidth float64,
	borderStyle BorderStyle,
	cloudyBorder bool,
	cloudyBorderIntensity int) FreeTextAnnotation {
	_ = "STUB: not implemented"
	return *new(FreeTextAnnotation)
}

func (ann FreeTextAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type LineIntent int

const (
	IntentLineArrow LineIntent = 1 << iota
	IntentLineDimension
)

func LineIntentName(li LineIntent) string { _ = "STUB: not implemented"; return "" }

type LineAnnotation struct {
	MarkupAnnotation
	P1, P2                    types.Point
	LineEndings               types.Array
	LeaderLineLength          float64
	LeaderLineOffset          float64
	LeaderLineExtensionLength float64
	Intent                    string
	Measure                   types.Dict
	Caption                   bool
	CaptionPositionTop        bool
	CaptionOffsetX            float64
	CaptionOffsetY            float64
	FillCol                   *color.SimpleColor
	BorderWidth               float64
	BorderStyle               BorderStyle
}

func NewLineAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	p1, p2 types.Point,
	beginLineEndingStyle *LineEndingStyle,
	endLineEndingStyle *LineEndingStyle,
	leaderLineLength float64,
	leaderLineOffset float64,
	leaderLineExtensionLength float64,
	intent *LineIntent,
	measure types.Dict,
	caption bool,
	captionPosTop bool,
	captionOffsetX float64,
	captionOffsetY float64,
	fillCol *color.SimpleColor,
	borderWidth float64,
	borderStyle BorderStyle) LineAnnotation {
	_ = "STUB: not implemented"
	return *new(LineAnnotation)
}

func (ann LineAnnotation) validateLeaderLineAttrs() error { _ = "STUB: not implemented"; return nil }

func (ann LineAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type SquareAnnotation struct {
	MarkupAnnotation
	FillCol               *color.SimpleColor
	Margins               types.Array
	BorderWidth           float64
	BorderStyle           BorderStyle
	CloudyBorder          bool
	CloudyBorderIntensity int
}

func NewSquareAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	fillCol *color.SimpleColor,
	MLeft, MTop, MRight, MBot float64,
	borderWidth float64,
	borderStyle BorderStyle,
	cloudyBorder bool,
	cloudyBorderIntensity int) SquareAnnotation {
	_ = "STUB: not implemented"
	return *new(SquareAnnotation)
}

func (ann SquareAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type CircleAnnotation struct {
	MarkupAnnotation
	FillCol               *color.SimpleColor
	Margins               types.Array
	BorderWidth           float64
	BorderStyle           BorderStyle
	CloudyBorder          bool
	CloudyBorderIntensity int
}

func NewCircleAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	fillCol *color.SimpleColor,
	MLeft, MTop, MRight, MBot float64,
	borderWidth float64,
	borderStyle BorderStyle,
	cloudyBorder bool,
	cloudyBorderIntensity int) CircleAnnotation {
	_ = "STUB: not implemented"
	return *new(CircleAnnotation)
}

func (ann CircleAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type PolygonIntent int

const (
	IntentPolygonCloud PolygonIntent = 1 << iota
	IntentPolygonDimension
)

func PolygonIntentName(pi PolygonIntent) string { _ = "STUB: not implemented"; return "" }

type PolygonAnnotation struct {
	MarkupAnnotation
	Vertices              types.Array
	Path                  types.Array
	Intent                string
	Measure               types.Dict
	FillCol               *color.SimpleColor
	BorderWidth           float64
	BorderStyle           BorderStyle
	CloudyBorder          bool
	CloudyBorderIntensity int
}

func NewPolygonAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	vertices types.Array,
	path types.Array,
	intent *PolygonIntent,
	measure types.Dict,
	fillCol *color.SimpleColor,
	borderWidth float64,
	borderStyle BorderStyle,
	cloudyBorder bool,
	cloudyBorderIntensity int) PolygonAnnotation {
	_ = "STUB: not implemented"
	return *new(PolygonAnnotation)
}

func (ann PolygonAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type PolyLineIntent int

const (
	IntentPolyLinePolygonCloud PolyLineIntent = 1 << iota
	IntentPolyLineDimension
)

func PolyLineIntentName(pi PolyLineIntent) string { _ = "STUB: not implemented"; return "" }

type PolyLineAnnotation struct {
	MarkupAnnotation
	Vertices    types.Array
	Path        types.Array
	Intent      string
	Measure     types.Dict
	FillCol     *color.SimpleColor
	BorderWidth float64
	BorderStyle BorderStyle
	LineEndings types.Array
}

func NewPolyLineAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	vertices types.Array,
	path types.Array,
	intent *PolyLineIntent,
	measure types.Dict,
	fillCol *color.SimpleColor,
	borderWidth float64,
	borderStyle BorderStyle,
	beginLineEndingStyle *LineEndingStyle,
	endLineEndingStyle *LineEndingStyle) PolyLineAnnotation {
	_ = "STUB: not implemented"
	return *new(PolyLineAnnotation)
}

func (ann PolyLineAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type TextMarkupAnnotation struct {
	MarkupAnnotation
	Quad types.QuadPoints
}

func NewTextMarkupAnnotation(
	subType AnnotationType,
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	quad types.QuadPoints) TextMarkupAnnotation {
	_ = "STUB: not implemented"
	return *new(TextMarkupAnnotation)
}

func (ann TextMarkupAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type HighlightAnnotation struct {
	TextMarkupAnnotation
}

func NewHighlightAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	quad types.QuadPoints) HighlightAnnotation {
	_ = "STUB: not implemented"
	return *new(HighlightAnnotation)
}

type UnderlineAnnotation struct {
	TextMarkupAnnotation
}

func NewUnderlineAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	quad types.QuadPoints) UnderlineAnnotation {
	_ = "STUB: not implemented"
	return *new(UnderlineAnnotation)
}

type SquigglyAnnotation struct {
	TextMarkupAnnotation
}

func NewSquigglyAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	quad types.QuadPoints) SquigglyAnnotation {
	_ = "STUB: not implemented"
	return *new(SquigglyAnnotation)
}

type StrikeOutAnnotation struct {
	TextMarkupAnnotation
}

func NewStrikeOutAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	quad types.QuadPoints) StrikeOutAnnotation {
	_ = "STUB: not implemented"
	return *new(StrikeOutAnnotation)
}

type CaretAnnotation struct {
	MarkupAnnotation
	RD        *types.Rectangle
	Paragraph bool
}

func NewCaretAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	borderRadX float64,
	borderRadY float64,
	borderWidth float64,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	rd *types.Rectangle,
	paragraph bool) CaretAnnotation {
	_ = "STUB: not implemented"
	return *new(CaretAnnotation)
}

func (ann CaretAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

type InkPath []float64

type InkAnnotation struct {
	MarkupAnnotation
	InkList     []InkPath
	BorderWidth float64
	BorderStyle BorderStyle
}

func NewInkAnnotation(
	rect types.Rectangle,
	apObjNr int,
	contents, id string,
	modDate string,
	f AnnotationFlags,
	col *color.SimpleColor,
	title string,
	popupIndRef *types.IndirectRef,
	ca *float64,
	rc, subject string,

	ink []InkPath,
	borderWidth float64,
	borderStyle BorderStyle) InkAnnotation {
	_ = "STUB: not implemented"
	return *new(InkAnnotation)
}

func (ann InkAnnotation) RenderDict(xRefTable *XRefTable, pageIndRef *types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}
