package sign

import (
	"errors"
)

var (
	errCertificateParse = errors.New("certificate parse error")

	errMalformedByteRange = errors.New("malformed signature ByteRange")

	errUnsupportedPublicKey = errors.New("unsupported public key algorithm")

	errMalformedPublicKey = errors.New("malformed public key")

	errESSCertificateMismatch = errors.New("ESS: signing certificate mismatch")

	errUnsupportedESSCertificateProfile = errors.New("ESS: unsupported signing certificate profile")

	errMalformedAdobePKCS7Profile = errors.New("Adobe PKCS#7: malformed profile")

	errMalformedCAdESBaselineBProfile = errors.New("CAdES baseline B: malformed profile")

	errUnsupportedCAdESBaselineBProfile = errors.New("CAdES baseline B: unsupported profile")

	errCAdESCertificateBindingMismatch = errors.New("CAdES baseline B: certificate binding mismatch")
)

type certificateParseError struct {
	cause error
}

func (e *certificateParseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *certificateParseError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *certificateParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type malformedByteRangeError struct {
	cause error
}

func (e *malformedByteRangeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *malformedByteRangeError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *malformedByteRangeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type byteRangeReadError struct {
	cause error
}

func (e *byteRangeReadError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *byteRangeReadError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func malformedByteRange(cause error) error { _ = "STUB: not implemented"; return nil }

func fatalByteRangeRead(cause error) error { _ = "STUB: not implemented"; return nil }
