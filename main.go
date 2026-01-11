package main

import "fmt"

type Vector []float64

type Data struct {
	index int
  vector Vector
}

type DataStore struct {
	tail int
	data []Data
}

func (v Vector) Dimension() int {
	return len(v)
}

func (d *DataStore) Add(vector Vector) {
	d.tail++
	newData := Data{d.tail, vector}
	d.data = append(d.data, newData)
}

func main() {
	fmt.Println("Hi")

// map id -> data
//vectorsfn := map[int]Vector{}

// map id -> neighbours
//neighbours := map[int][]int{}

var store DataStore

// init vectors
vectors := []Vector{
	{0.1, 0.2, 0.3, 0.4},
	{0.2, 0.3, 0.4, 0.5},
	{0.3, 0.4, 0.5, 0.6},
}

for i, value := range(vectors) {
	fmt.Println(i, value)
	store.Add(value)
}

fmt.Println(store)

}
