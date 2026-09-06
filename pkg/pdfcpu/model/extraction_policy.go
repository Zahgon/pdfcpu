package model

type UnsupportedResourcePolicy int

const (
	UnsupportedResourceSkip UnsupportedResourcePolicy = iota

	UnsupportedResourceFail
)
