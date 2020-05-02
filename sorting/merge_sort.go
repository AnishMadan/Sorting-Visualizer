package sorting

import (
	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type MergeSort struct {
	A              []int
	area           *ui.Area
	iterationLabel *ui.Label
}

func merge(A []int, l int, m int, r int, area *ui.Area, iterationLabel *ui.Label) {
	n1 := m - l + 1
	n2 := r - m

	/* create temp arrays */
	L := make([]int, n1)
	R := make([]int, n2)

	/* Copy data to temp arrays L[] and R[] */
	for i := 0; i < n1; i++ {
		L[i] = A[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[m+1+j]
	}

	/* Merge the temp arrays back into arr[l..r]*/
	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}

		config.RunningTotal++
		redraw(area, iterationLabel)

		k++
	}

	/* Copy the remaining elements of L[], if there are any */
	for i < n1 {
		A[k] = L[i]
		i++
		k++
	}

	/* Copy the remaining elements of R[], if there are any */
	for j < n2 {
		A[k] = R[j]
		j++
		k++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (algo MergeSort) Sort() {
	A := algo.A
	n := len(A)

	for currSize := 1; currSize <= n-1; currSize = 2 * currSize {
		for leftStart := 0; leftStart < n-1; leftStart += 2 * currSize {

			mid := min(leftStart+currSize-1, n-1)

			rightEnd := min(leftStart+2*currSize-1, n-1)

			if config.Stop || config.SortType != 4 {
				currSize = n
				break
			}

			// Merge Subarrays arr[left_start...mid] & arr[mid+1...right_end]
			merge(A, leftStart, mid, rightEnd, algo.area, algo.iterationLabel)

		}
	}

	if config.Stop || config.SortType != 4 {
		SortDecider()
	}

}
