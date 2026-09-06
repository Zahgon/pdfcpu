package primitives

type DateFormat struct {
	Int string
	Ext string
}

var dateFormats = []DateFormat{

	{"2006-1-2", "yyyy-m-d"},
	{"2006-2-1", "yyyy-d-m"},
	{"2006-01-02", "yyyy-mm-dd"},
	{"2006-02-01", "yyyy-dd-mm"},
	{"02-01-2006", "dd-mm-yyyy"},
	{"01-02-2006", "mm-dd-yyyy"},
	{"2-1-2006", "d-m-yyyy"},
	{"1-2-2006", "m-d-yyyy"},

	{"2006/1/2", "yyyy/m/d"},
	{"2006/2/1", "yyyy/d/m"},
	{"2006/01/02", "yyyy/mm/dd"},
	{"2006/02/01", "yyyy/dd/mm"},
	{"02/01/2006", "dd/mm/yyyy"},
	{"01/02/2006", "mm/dd/yyyy"},
	{"2/1/2006", "d/m/yyyy"},
	{"1/2/2006", "m/d/yyyy"},

	{"2006.1.2", "yyyy.m.d"},
	{"2006.2.1", "yyyy.d.m"},
	{"2006.01.02", "yyyy.mm.dd"},
	{"2006.02.01", "yyyy.dd.mm"},
	{"02.01.2006", "dd.mm.yyyy"},
	{"01.02.2006", "mm.dd.yyyy"},
	{"2.1.2006", "d.m.yyyy"},
	{"1.2.2006", "m.d.yyyy"},
}

func DateFormatForFmtInt(fmtInt string) (*DateFormat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DateFormatForFmtExt(fmtExt string) (*DateFormat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DateFormatForDate(date string) (*DateFormat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (df DateFormat) validate(date string) error { _ = "STUB: not implemented"; return nil }
