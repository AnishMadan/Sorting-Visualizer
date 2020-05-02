package sorting

import (
	"fmt"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type BubbleSort struct {
	A              []int
	area           *ui.Area
	iterationLabel *ui.Label
}

// BubbleSort implementation
func (algo BubbleSort) Sort() {
	A := algo.A
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}

			if config.Stop || config.SortType != 1 {
				i = len(A) // outer loop
				break
			}

			config.RunningTotal++
			redraw(algo.area, algo.iterationLabel)
		}
	}

	if config.SortType != 1 {
		SortDecider()
	}

	fmt.Printf("%d iterations for Bubble Sort \n", config.RunningTotal)
}
