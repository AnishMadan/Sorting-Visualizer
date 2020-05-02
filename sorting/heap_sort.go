package sorting

import (
	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type HeapSort struct {
	A              []int
	area           *ui.Area
	iterationLabel *ui.Label
}

// BubbleSort implementation
func (algo HeapSort) Sort() {
	A := algo.A
	var t, index, child int
	n := len(A)
	parent := len(A) / 2
	/* loop until array is sorted */
	for {
		if parent > 0 {
			/* first stage - Sorting the heap */
			parent--
			t = A[parent] /* save old value to t */
		} else {
			/* second stage - Extracting elements in-place */
			n-- /* make the heap smaller */
			if n == 0 {
				return /* When the heap is empty, we are done */
			}
			t = A[n]    /* save lost heap entry to temporary */
			A[n] = A[0] /* save root entry beyond heap */
		}
		/* insert operation - pushing t down the heap to replace the parent */
		index = parent      /* start at the parent index */
		child = index*2 + 1 /* get its left child index */
		for child < n {
			/* choose the largest child */
			if child+1 < n && A[child+1] > A[child] {
				child++ /* right child exists and is bigger */
			}
			/* is the largest child larger than the entry? */
			if A[child] > t {
				A[index] = A[child] /* overwrite entry with child */
				index = child       /* move index to the child */
				child = index*2 + 1 /* get the left child and go around again */
			} else {
				break /* t's place is found */
			}

			config.RunningTotal++
			redraw(algo.area, algo.iterationLabel)
		}
		/* store the temporary value at its new location */
		A[index] = t

		config.RunningTotal++
		redraw(algo.area, algo.iterationLabel)

		if config.Stop || config.SortType != 5 {
			break
		}

	}

	if config.Stop || config.SortType != 5 {
		SortDecider()
	}
}
