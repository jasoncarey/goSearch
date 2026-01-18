package vec

import (
	"testing"
	"math"
)

func TestMagnitude(t *testing.T) {
	// Arrange
	v := Vector{3.0, 4.0, 5.0}

	// Act
	got := v.Magnitude()

	// Assert
	want := math.Sqrt(50)
	if got != want {
		t.Errorf("Magnitude() = %v, want %v", got, want)
	}
}

func TestDotProduct(t * testing.T) {
	v := Vector{1.0, 3.0, 5.0}
	u := Vector{2.0, 4.0, 6.0}

	got := v.DotProduct(u)
	
	want := 44.0

	if got != want {
		t.Errorf("DotProduct() = %v, want %v", got, want)
	}
}

