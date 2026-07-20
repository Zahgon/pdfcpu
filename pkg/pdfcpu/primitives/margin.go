package primitives

type Margin struct {
	Name                     string
	Width                    float64
	Top, Right, Bottom, Left float64
}

func (m *Margin) validate() error { _ = "STUB: not implemented"; return nil }

func (m *Margin) mergeIn(m0 *Margin) { _ = "STUB: not implemented"; return }
