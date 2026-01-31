package engine

import (
	"sort"

	"search/vec"
)

type Data struct {
	index  int
	vector vec.Vector
}

type Result struct {
	index int
	score float64
}

type Engine struct {
	data []Data
}

func (e *Engine) Add(v vec.Vector) {
	e.data = append(e.data, Data{index: len(e.data) + 1, vector: v})
}

func (e *Engine) BruteForceSearch(query vec.Vector, k int) []Result {
	// check query against everything in the data store
	results := []Result{}
	for i, d := range e.data {
		score := query.CosineSimilarity(d.vector)
		results = append(results, Result{index: i, score: score})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})
	return results[:min(k, len(results))]
}
