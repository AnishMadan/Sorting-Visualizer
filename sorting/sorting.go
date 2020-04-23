package sorting

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

func Sort(A []int, area *ui.Area, iterationLabel *ui.Label) {
	if config.SortType == 0 {
		go insertionSort(A, area, iterationLabel)
	} else if config.SortType == 1 {
		go bubbleSort(A, area, iterationLabel)
	} else if config.SortType == 2 {
		go bogoSort(A, area, iterationLabel)
	} else if config.SortType == 3 {
		go quickSort(A, area, iterationLabel)
	} else if config.SortType == 4 {
		go mergeSort(A, area, iterationLabel)
	}
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

func bogoSort(A []int, area *ui.Area, iterationLabel *ui.Label) {
	for !isSorted(A) {
		shuffle(A)
		config.RunningTotal++
		if config.Stop || config.SortType != 2 {
			break
		}
		time.Sleep(time.Second / time.Duration(config.SortSpeed/5))
		ui.QueueMain(func() {
			area.QueueRedrawAll()
			iterationLabel.SetText(strconv.Itoa(config.RunningTotal))
		})
		fmt.Println(A)
	}

	if config.SortType != 2 {
		Sort(A, area, iterationLabel)
	}

}

// InsertionSort implementation
func insertionSort(A []int, area *ui.Area, iterationLabel *ui.Label) {
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

			time.Sleep(time.Second / time.Duration(config.SortSpeed/5))
			ui.QueueMain(func() {
				area.QueueRedrawAll()
				iterationLabel.SetText(strconv.Itoa(config.RunningTotal))
			})
			fmt.Println(A)
		}

		i++
	}

	if config.SortType != 0 {
		Sort(A, area, iterationLabel)
	}

	fmt.Printf("%d iterations for Insertion Sort \n", config.RunningTotal)
}

// BubbleSort implementation
func bubbleSort(A []int, area *ui.Area, iterationLabel *ui.Label) {
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
			time.Sleep(time.Second / time.Duration(config.SortSpeed/5))
			ui.QueueMain(func() {
				area.QueueRedrawAll()
				iterationLabel.SetText(strconv.Itoa(config.RunningTotal))
			})
		}
	}

	if config.SortType != 1 {
		Sort(A, area, iterationLabel)
	}

	fmt.Printf("%d iterations for Bubble Sort \n", config.RunningTotal)
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
		time.Sleep(time.Second / time.Duration(config.SortSpeed/5))
		ui.QueueMain(func() {
			area.QueueRedrawAll()
			iterationLabel.SetText(strconv.Itoa(config.RunningTotal))
		})
	}
	A[i+1], A[h] = A[h], A[i+1]

	return (i + 1)
}

func quickSort(A []int, area *ui.Area, iterationLabel *ui.Label) {
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

		p := partition(A, l, h, area, iterationLabel)

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
		Sort(A, area, iterationLabel)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mergeSort(A []int, area *ui.Area, iterationLabel *ui.Label) {

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
			merge(A, leftStart, mid, rightEnd, area, iterationLabel)

		}
	}

	if config.Stop || config.SortType != 4 {
		Sort(A, area, iterationLabel)
	}

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
		time.Sleep(time.Second / time.Duration(config.SortSpeed/5))
		ui.QueueMain(func() {
			area.QueueRedrawAll()
			iterationLabel.SetText(strconv.Itoa(config.RunningTotal))
		})

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
