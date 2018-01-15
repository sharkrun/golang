package util

type Rect struct {
	X, Y float64
	Width, Height float64
}


func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

func NewRect(x, y, width, height float64) *Rect{
	return &Rect{x, y, width, height}
}