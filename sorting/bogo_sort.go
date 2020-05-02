package sorting

import (
	"math/rand"
	"time"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type BogoSort struct {
	A              []int
	area           *ui.Area
	iterationLabel *ui.Label
}

func shuffle(A []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(A); n > 0; n-- {
		randIndex := r.Intn(n)
		A[n-1], A[randIndex] = A[randIndex], A[n-1]
	}
}

func isSorted(A []int) bool {
	for i := range A {

		if i == len(A)-1 {
			break
		}

		if A[i] > A[i+1] {
			return false
		}
	}

	return true
}

func (algo BogoSort) Sort() {
	A := algo.A

	for !isSorted(A) {
		shuffle(A)
		if config.Stop || config.SortType != 2 {
			break
		}
		config.RunningTotal++
		redraw(algo.area, algo.iterationLabel)
	}

	if config.SortType != 2 {
		SortDecider()
	}

}
