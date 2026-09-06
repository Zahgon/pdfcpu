package fault

type Panic struct {
	Err   error
	Stack []byte
}

func (p Panic) Error() string { _ = "STUB: not implemented"; return "" }

func (p Panic) Unwrap() error { _ = "STUB: not implemented"; return nil }

func Fail(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func Catch(err *error) { _ = "STUB: not implemented"; return }
