package vec

import (
	"testing"
	"math"
)

func TestCosineSimilarity(t *testing.T) {
	v := Vector{1.0, 2.0, 3.0}
	u := Vector{2.0, 3.0, 4.0}

	got := v.CosineSimilarity(u)

	want := 20 / (math.Sqrt(14) * math.Sqrt(29))

	if got != want {
		t.Errorf("CosineSimilarity() = %v, want %v", got, want);
	}
}

