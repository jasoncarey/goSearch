package engine

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"search/vec"
)

type Centroid struct {
	vector  vec.Vector
	members []Data
}

type IVFIndex struct {
	pending   []Data
	centroids []Centroid
	nprobe    int
}

func Average(data []Data) vec.Vector {
	dim := len(data[0].vector)
	sum := make([]float64, dim)

	for _, x := range data {
		for d := range dim {
			sum[d] += x.vector[d]
		}
	}

	count := float64(len(data))
	for d := range dim {
		sum[d] /= count
	}

	return sum
}

func (idx *IVFIndex) Add(v vec.Vector) {
	idx.pending = append(idx.pending, Data{len(idx.pending) + 1, v})
}

func (idx *IVFIndex) Train(k int) {
	for range k {
		centroidIdx := rand.Intn(len(idx.pending))
		idx.centroids = append(idx.centroids, Centroid{idx.pending[centroidIdx].vector, []Data{}})
	}

	for range 10 {
		for i := range k {
			idx.centroids[i].members = []Data{}
		}

		for i := range len(idx.pending) {
			maxSimilarity := math.Inf(-1)
			maxSimilarityIdx := 0
			for j := range len(idx.centroids) {
				similarity := idx.centroids[j].vector.CosineSimilarity(idx.pending[i].vector)
				if similarity >= maxSimilarity {
					maxSimilarity = similarity
					maxSimilarityIdx = j
				}
			}
			idx.centroids[maxSimilarityIdx].members = append(idx.centroids[maxSimilarityIdx].members, idx.pending[i])
		}

		for i := range len(idx.centroids) {
			idx.centroids[i].vector = Average(idx.centroids[i].members)
		}
	}

	for i := range k {
		fmt.Print(idx.centroids[i])
	}

	idx.pending = nil
}

func (idx *IVFIndex) Search(query vec.Vector, k int) []Result {
	results := []Result{}
	centroidSimilarity := math.Inf(-1)
	centroidIdx := 0

	for n := range len(idx.centroids) {
		similarity := query.CosineSimilarity(idx.centroids[n].vector)
		if similarity >= centroidSimilarity {
			centroidSimilarity = similarity
			centroidIdx = n
		}
	}

	cluster := idx.centroids[centroidIdx].members
	for _, d := range cluster {
		score := query.CosineSimilarity(d.vector)
		results = append(results, Result{d.index, score})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].score > results[j].score
	})
	return results[:min(k, len(results))]
}
