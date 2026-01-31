package engine

import (
	"testing"

	"search/vec"
)

func TestBruteForceSearch(t *testing.T) {
	// Setup: create engine and add vectors
	e := &Engine{}
	e.Add(vec.Vector{1, 0, 0})
	e.Add(vec.Vector{0, 1, 0})
	e.Add(vec.Vector{0, 0, 1})
	e.Add(vec.Vector{1, 1, 0})

	// Query for vector most similar to {1, 0, 0}
	// Expected: index 0 (exact match), then index 3 (shares x component)
	query := vec.Vector{1, 0, 0}
	results := e.BruteForceSearch(query, 2)

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

func TestBruteForceSearchKGreaterThanData(t *testing.T) {
	e := &Engine{}
	e.Add(vec.Vector{1, 0, 0})
	e.Add(vec.Vector{0, 1, 0})

	query := vec.Vector{1, 0, 0}
	results := e.BruteForceSearch(query, 10)

	if len(results) != 2 {
		t.Errorf("go %d results, want 2", len(results))
	}
}
