package internals

import "math"

type Circle struct {
	Radius float32
}

func (c Circle) GetWide() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) GetAround() float32 {
	return 2 * math.Pi * c.Radius
}
