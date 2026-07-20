package cli

import (
	"io"
)

func readSeekerFromStdin() (io.ReadSeeker, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil
}

func streamInOut(inFile, outFile string) (io.ReadSeeker, io.Writer, func(), error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), *new(io.Writer), nil, nil
}
