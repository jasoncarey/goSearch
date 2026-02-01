package engine

import (
	"testing"

	"search/vec"
)

func TestIVFIndexSearch(t *testing.T) {
	ivf := &IVFIndex{}

	ivf.Add(vec.Vector{1, 0, 0})
	ivf.Add(vec.Vector{0, 1, 0})
	ivf.Add(vec.Vector{0, 0, 1})
	ivf.Add(vec.Vector{1, 1, 0})

	query := vec.Vector{1, 0, 0}

	ivf.Train(2)

	results := ivf.Search(query, 2)

	if len(results) != 2 {
		t.Fatalf("got %d results, want 2", len(results))
	}

	if results[0].index != 0 {
		t.Errorf("first result index = %d, want 0", results[0].index)
	}

	if results[0].score != 1.0 {
		t.Errorf("first result score = %f, want 1.0", results[0].score)
	}

	if results[1].index != 3 {
		t.Errorf("second result index = %d, want 3", results[1].index)
	}
}
