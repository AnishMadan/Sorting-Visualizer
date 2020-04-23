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
		time.Sleep(time.Second / time.Duration(config.SortSpeed/2))
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

			time.Sleep(time.Second / time.Duration(config.SortSpeed/2))
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

			fmt.Println(A)
			config.RunningTotal++
			time.Sleep(time.Second / time.Duration(config.SortSpeed/2))
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

func choosePivot(A []int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(len(A))
}

// func parition(A []int, p int) int {
// 	n := len(A)

// 	A[n-1], A[p] = A[p], A[n-1]

// 	i := -1
// 	j := n - 1
// 	v := A[n-1]

// 	for {

// 		for {
// 			i++
// 			if i <= n && A[i] <= v {
// 				break
// 			}
// 		}

// 		for {
// 			j--
// 			if j >= 0 && A[i] >= v {
// 				break
// 			}
// 		}

// 		if i >= j {
// 			break
// 		} else {
// 			A[i], A[j] = A[j], A[i]
// 		}
// 	}

// 	A[n-1], A[i] = A[i], A[n-1]

// 	return i
// }

func partition(lo int, piv int, arr []int) int {
	is := lo

	for i := lo; i < piv; i++ {
		if arr[i] < arr[piv] {
			if i != is {
				arr[i], arr[is] = arr[is], arr[i]
			}

			is++
		}
	}

	arr[is], arr[piv] = arr[piv], arr[is]

	if is-1 > lo {
		partition(lo, is-1, arr)
	}
	if is+1 < piv {
		partition(is+1, piv, arr)
	}

	return is
}

func QuickSort(A []int) {
	if len(A) <= 1 {
		return
	}
	p := len(A) - 1
	i := partition(0, p, A)
	fmt.Println(A)
	QuickSort(A[:i])
	QuickSort(A[i+1:])
}

// func QuickSort(a []int) []int {
// 	if len(a) < 2 {
// 		return a
// 	}

// 	left, right := 0, len(a)-1

// 	pivot := rand.Int() % len(a)

// 	a[pivot], a[right] = a[right], a[pivot]

// 	for i, _ := range a {
// 		if a[i] < a[right] {
// 			a[left], a[i] = a[i], a[left]
// 			left++
// 		}
// 	}

// 	a[left], a[right] = a[right], a[left]

// 	QuickSort(a[:left])
// 	QuickSort(a[left+1:])

// 	return a
// }
