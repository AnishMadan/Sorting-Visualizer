package sorting

import (
	"strconv"
	"time"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
)

type Sort interface {
	Sort()
}

var (
	algorithms []Sort
)

func SortSetup(A []int, area *ui.Area, iterationLabel *ui.Label) {
	algorithms = make([]Sort, 6)
	algorithms[0] = InsertionSort{A, area, iterationLabel}
	algorithms[1] = BubbleSort{A, area, iterationLabel}
	algorithms[2] = BogoSort{A, area, iterationLabel}
	algorithms[3] = QuickSort{A, area, iterationLabel}
	algorithms[4] = MergeSort{A, area, iterationLabel}
	algorithms[5] = HeapSort{A, area, iterationLabel}
}

func SortDecider() {
	go algorithms[config.SortType].Sort()
}

func redraw(area *ui.Area, iterationLabel *ui.Label) {
	time.Sleep(sleepTime())
	ui.QueueMain(func() {
		area.QueueRedrawAll()
		iterationLabel.SetText(strconv.Itoa(config.RunningTotal))
	})
}

func sleepTime() time.Duration {
	return (time.Second / time.Duration(config.SortSpeed/5))
}
