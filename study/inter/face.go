package inter

type Integer int
func (a Integer) Less (b Integer) bool {
	return a < b
}

