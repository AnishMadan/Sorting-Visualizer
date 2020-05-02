package sorting

import (
	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type QuickSort struct {
	A              []int
	area           *ui.Area
	iterationLabel *ui.Label
}

func partition(A []int, l int, h int, area *ui.Area, iterationLabel *ui.Label) int {
	x := A[h]
	i := l - 1

	for j := l; j <= h-1; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}

		if config.Stop || config.SortType != 3 {
			break
		}

		config.RunningTotal++
		redraw(area, iterationLabel)
	}
	A[i+1], A[h] = A[h], A[i+1]

	return (i + 1)
}

func (algo QuickSort) Sort() {
	A := algo.A
	l := 0
	h := len(A) - 1

	stack := make([]int, h-l+1)

	top := -1

	top++
	stack[top] = l

	top++
	stack[top] = h

	for top >= 0 {
		h = stack[top]
		top--
		l = stack[top]
		top--

		if config.Stop || config.SortType != 3 {
			break
		}

		p := partition(A, l, h, algo.area, algo.iterationLabel)

		if p-1 > l {
			top++
			stack[top] = l
			top++
			stack[top] = p - 1
		}

		if p+1 < h {
			top++
			stack[top] = p + 1
			top++
			stack[top] = h
		}
	}

	if config.Stop || config.SortType != 3 {
		SortDecider()
	}

}
