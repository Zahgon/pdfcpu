package cli

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type dispatchFunc func(*Command) ([]string, error)

var dispatchTable = map[model.CommandMode]dispatchFunc{

	model.VALIDATE:       Validate,
	model.OPTIMIZE:       Optimize,
	model.LISTINFO:       ListInfo,
	model.DUMP:           Dump,
	model.CREATE:         Create,
	model.MERGECREATE:    MergeCreate,
	model.MERGECREATEZIP: MergeCreateZip,
	model.MERGEAPPEND:    MergeAppend,
	model.SPLIT:          Split,
	model.SPLITBYPAGENR:  SplitByPageNr,
	model.TRIM:           Trim,
	model.COLLECT:        Collect,

	model.INSERTPAGESBEFORE: dispatchPages,
	model.INSERTPAGESAFTER:  dispatchPages,
	model.REMOVEPAGES:       dispatchPages,
	model.ROTATE:            Rotate,
	model.NUP:               NUp,
	model.BOOKLET:           Booklet,
	model.RESIZE:            Resize,
	model.POSTER:            Poster,
	model.NDOWN:             NDown,
	model.CUT:               Cut,
	model.CROP:              dispatchPageBoundaries,
	model.ZOOM:              Zoom,

	model.ADDWATERMARKS:          AddWatermarks,
	model.REMOVEWATERMARKS:       RemoveWatermarks,
	model.LISTANNOTATIONS:        dispatchPageAnnotations,
	model.REMOVEANNOTATIONS:      dispatchPageAnnotations,
	model.LISTBOOKMARKS:          dispatchBookmarks,
	model.EXPORTBOOKMARKS:        dispatchBookmarks,
	model.IMPORTBOOKMARKS:        dispatchBookmarks,
	model.REMOVEBOOKMARKS:        dispatchBookmarks,
	model.LISTPAGEMODE:           dispatchPageMode,
	model.SETPAGEMODE:            dispatchPageMode,
	model.RESETPAGEMODE:          dispatchPageMode,
	model.LISTPAGELAYOUT:         dispatchPageLayout,
	model.SETPAGELAYOUT:          dispatchPageLayout,
	model.RESETPAGELAYOUT:        dispatchPageLayout,
	model.LISTVIEWERPREFERENCES:  dispatchViewerPreferences,
	model.SETVIEWERPREFERENCES:   dispatchViewerPreferences,
	model.RESETVIEWERPREFERENCES: dispatchViewerPreferences,

	model.IMPORTIMAGES:            ImportImages,
	model.CHEATSHEETSFONTS:        CreateCheatSheetsFonts,
	model.INSTALLFONTS:            InstallFonts,
	model.LISTFONTS:               ListFonts,
	model.LISTIMAGES:              dispatchImages,
	model.UPDATEIMAGES:            dispatchImages,
	model.LISTATTACHMENTS:         dispatchAttachments,
	model.ADDATTACHMENTS:          dispatchAttachments,
	model.ADDATTACHMENTSPORTFOLIO: dispatchAttachments,
	model.REMOVEATTACHMENTS:       dispatchAttachments,
	model.EXTRACTATTACHMENTS:      dispatchAttachments,
	model.LISTKEYWORDS:            dispatchKeywords,
	model.ADDKEYWORDS:             dispatchKeywords,
	model.REMOVEKEYWORDS:          dispatchKeywords,
	model.LISTPROPERTIES:          dispatchProperties,
	model.ADDPROPERTIES:           dispatchProperties,
	model.REMOVEPROPERTIES:        dispatchProperties,
	model.LISTBOXES:               dispatchPageBoundaries,
	model.ADDBOXES:                dispatchPageBoundaries,
	model.REMOVEBOXES:             dispatchPageBoundaries,

	model.EXTRACTIMAGES:   ExtractImages,
	model.EXTRACTFONTS:    ExtractFonts,
	model.EXTRACTPAGES:    ExtractPages,
	model.EXTRACTCONTENT:  ExtractContent,
	model.EXTRACTMETADATA: ExtractMetadata,

	model.LISTFORMFIELDS:      dispatchForm,
	model.REMOVEFORMFIELDS:    dispatchForm,
	model.LOCKFORMFIELDS:      dispatchForm,
	model.UNLOCKFORMFIELDS:    dispatchForm,
	model.RESETFORMFIELDS:     dispatchForm,
	model.EXPORTFORMFIELDS:    dispatchForm,
	model.FILLFORMFIELDS:      dispatchForm,
	model.MULTIFILLFORMFIELDS: dispatchForm,

	model.ENCRYPT:         dispatchEncryption,
	model.DECRYPT:         dispatchEncryption,
	model.CHANGEUPW:       dispatchEncryption,
	model.CHANGEOPW:       dispatchEncryption,
	model.LISTPERMISSIONS: dispatchPermissions,
	model.SETPERMISSIONS:  dispatchPermissions,

	model.LISTCERTIFICATES:    dispatchCertificates,
	model.INSPECTCERTIFICATES: dispatchCertificates,
	model.IMPORTCERTIFICATES:  dispatchCertificates,
	model.VALIDATESIGNATURES:  dispatchSignatures,
	model.REMOVESIGNATURES:    dispatchSignatures,
	model.ADDSIGNATURE:        dispatchSignatures,
}

func Dispatch(cmd *Command) (out []string, err error) { _ = "STUB: not implemented"; return nil, nil }

func dispatchAttachments(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchBookmarks(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchEncryption(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchForm(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchImages(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchKeywords(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchPageAnnotations(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchPageBoundaries(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchPageLayout(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchPageMode(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchPages(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchPermissions(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchProperties(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchViewerPreferences(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchCertificates(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchSignatures(cmd *Command) (out []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
