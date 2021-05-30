package gohelp

type If bool

func (c If) Int64(a, b int64) int64 {
	if c {
		return a
	}
	return b
}

func (c If) String(a, b string) string {
	if c {
		return a
	}
	return b
}

func (c If) Int(a, b int) int {
	if c {
		return a
	}
	return b
}
func (c If) Float64(a, b float64) float64 {
	if c {
		return a
	}
	return b
}

