package pdfcpu

type parameterMap[T any] map[string]func(string, *T) error

func handleParameter[T any](m map[string]func(string, *T) error, paramPrefix, paramValueStr string, v *T) error {
	_ = "STUB: not implemented"
	return nil
}
