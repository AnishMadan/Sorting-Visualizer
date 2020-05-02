package sorting

import (
	"fmt"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type InsertionSort struct {
	A              []int
	area           *ui.Area
	iterationLabel *ui.Label
}

func (algo InsertionSort) Sort() {
	A := algo.A
	i := 1
	for i < len(A) {
		j := i

		for j > 0 && A[j-1] > A[j] {
			A[j], A[j-1] = A[j-1], A[j]
			j--

			config.RunningTotal++

			if config.Stop || config.SortType != 0 {
				i = len(A) // outer loop
				break
			}

			redraw(algo.area, algo.iterationLabel)
		}

		i++
	}

	if config.SortType != 0 {
		SortDecider()
	}

	fmt.Printf("%d iterations for Insertion Sort \n", config.RunningTotal)
}
