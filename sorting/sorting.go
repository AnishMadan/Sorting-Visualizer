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

	if config.SortType != 3 {
		Sort(A, area, iterationLabel)
	}

}
