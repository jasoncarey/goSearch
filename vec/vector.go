package vec

import (
	"math"
)

type Vector []float64

func (v Vector) Dimension() int {
	return len(v)
}

func (v Vector) DotProduct(u Vector) float64 {
	dot := 0.0
	for i := range v {
		cur := v[i] * u[i]
		dot += cur
	}  
	return dot
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.DotProduct(v))
}
