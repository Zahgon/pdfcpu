package primitives

type Padding struct {
	Name                     string
	Width                    float64
	Top, Right, Bottom, Left float64
}

func (p *Padding) validate() error { _ = "STUB: not implemented"; return nil }

func (p *Padding) mergeIn(p0 *Padding) { _ = "STUB: not implemented"; return }
