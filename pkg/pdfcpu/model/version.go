package model

var VersionStr = "v0.13.0 dev"

type Version int

const (
	V10 Version = iota
	V11
	V12
	V13
	V14
	V15
	V16
	V17
	V20
)

func PDFVersion(versionStr string) (Version, error) {
	_ = "STUB: not implemented"
	return *new(Version), nil
}

func PDFVersionRelaxed(versionStr string) (Version, error) {
	_ = "STUB: not implemented"
	return *new(Version), nil
}

func (v Version) String() string { _ = "STUB: not implemented"; return "" }

func identicalMajorAndMinorVersions(v1, v2 string) bool { _ = "STUB: not implemented"; return false }

func CheckConfigVersion(v string) { _ = "STUB: not implemented"; return }
