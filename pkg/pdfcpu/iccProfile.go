package pdfcpu

type iccProfile struct {
	b          []byte
	rX, rY, rZ float32
	gX, gY, gZ float32
	bX, bY, bZ float32
}

func (p iccProfile) tag(sig string) (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (p *iccProfile) matrixCol(sig string) (float32, float32, float32, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func (p *iccProfile) init() error {

	var err error

	p.rX, p.rY, p.rZ, err = p.matrixCol("rXYZ")
	if err != nil {
		return err
	}

	p.gX, p.gY, p.gZ, err = p.matrixCol("gXYZ")
	if err != nil {
		return err
	}

	p.bX, p.bY, p.bZ, err = p.matrixCol("bXYZ")

	return err
}

func (p iccProfile) size() uint32 { _ = "STUB: not implemented"; return 0 }

func (p iccProfile) preferredCMM() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) version() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) class() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) dataColorSpace() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) pcs() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) creationTS() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) fileSig() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) primaryPlatform() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) deviceManufacturer() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) deviceModel() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) renderingIntent() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) xyz(i int) (x, y, z float32) { _ = "STUB: not implemented"; return 0, 0, 0 }

func (p iccProfile) PCSIlluminant() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) creator() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) id() string { _ = "STUB: not implemented"; return "" }

func (p iccProfile) tagCount() int { _ = "STUB: not implemented"; return 0 }

func (p iccProfile) String() string { _ = "STUB: not implemented"; return "" }
