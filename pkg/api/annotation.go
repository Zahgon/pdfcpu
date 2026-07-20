package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Annotations(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (m map[int]model.PgAnnots, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddAnnotations(rs io.ReadSeeker, w io.Writer, selectedPages []string, ann model.AnnotationRenderer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAnnotationsAsIncrement(rws io.ReadWriteSeeker, selectedPages []string, ar model.AnnotationRenderer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAnnotationsFile(inFile, outFile string, selectedPages []string, ar model.AnnotationRenderer, conf *model.Configuration, incr bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAnnotationsMap(rs io.ReadSeeker, w io.Writer, m map[int][]model.AnnotationRenderer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAnnotationsMapAsIncrement(rws io.ReadWriteSeeker, m map[int][]model.AnnotationRenderer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAnnotationsMapFile(inFile, outFile string, m map[int][]model.AnnotationRenderer, conf *model.Configuration, incr bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAnnotations(rs io.ReadSeeker, w io.Writer, selectedPages, idsAndTypes []string, objNrs []int, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAnnotationsAsIncrement(rws io.ReadWriteSeeker, selectedPages, idsAndTypes []string, objNrs []int, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAnnotationsFile(inFile, outFile string, selectedPages, idsAndTypes []string, objNrs []int, conf *model.Configuration, incr bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}
