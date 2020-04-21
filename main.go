package main

import (
	"math/rand"
	"time"

	"github.com/anishmadan/Sorting-Visualizer/sorting"
)

func setup(a []int) {
	for i := range a {
		a[i] = i + 1
	}
}

func shuffle(a []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(a); n > 0; n-- {
		randIndex := r.Intn(n)
		a[n-1], a[randIndex] = a[randIndex], a[n-1]
	}
}

func main() {
	A := make([]int, 50)

	setup(A)

	shuffle(A)

	sorting.InsertionSort(A)
}
