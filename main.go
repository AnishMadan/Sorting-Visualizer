package main

import (
	"math/rand"
	"time"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/config"
	"github.com/anishmadan/Sorting-Visualizer/sorting"
)

var (
	iterationLabel     *ui.Label
	sortButton         *ui.Button
	shuffleButton      *ui.Button
	stopButton         *ui.Button
	continueButton     *ui.Button
	typeOfSortComboBox *ui.Combobox
	sortSpeedSlider    *ui.Slider

	// A is an array for sorting
	A            []int
	sortSelected int
)

func setup(A []int) {
	for i := range A {
		A[i] = i + 1
	}
}

func shuffle(A []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(A); n > 0; n-- {
		randIndex := r.Intn(n)
		A[n-1], A[randIndex] = A[randIndex], A[n-1]
	}
}

type areaHandler struct{}

func (areaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	p := ui.DrawNewPath(ui.DrawFillModeWinding)

	for i, x := range A {
		p.NewFigure(0, 0)
		p.AddRectangle((float64)(i*10+15), 0, 5, (float64)(7*x))
	}
	p.End()

	dp.Context.Fill(p, &ui.DrawBrush{Type: ui.DrawBrushTypeSolid, R: .75, G: .25, B: 0, A: 1})
	p.Free()
}

func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}

func setupUI() {
	mainwin := ui.NewWindow("Sorting Examples", 1640, 1480, true)
	mainwin.SetMargined(true)
	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	mainwin.SetChild(hbox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	area := ui.NewArea(areaHandler{})

	sortButton = ui.NewButton("Sort")
	sortButton.OnClicked(func(*ui.Button) {
		config.Stop = false
		sorting.Sort(A, area, iterationLabel)
	})
	vbox.Append(sortButton, false)

	shuffleButton = ui.NewButton("Shuffle")
	shuffleButton.OnClicked(func(*ui.Button) {
		config.Stop = false
		config.RunningTotal = 0
		shuffle(A)
		iterationLabel.SetText("")
		area.QueueRedrawAll()
	})
	vbox.Append(shuffleButton, false)

	stopButton = ui.NewButton("Stop")
	stopButton.OnClicked(func(*ui.Button) {
		config.Stop = true
	})
	vbox.Append(stopButton, false)

	continueButton = ui.NewButton("Continue")
	continueButton.OnClicked(func(*ui.Button) {
		config.Stop = false
		sorting.Sort(A, area, iterationLabel)
	})
	vbox.Append(continueButton, false)

	form := ui.NewForm()
	form.SetPadded(true)
	vbox.Append(form, false)

	typeOfSortComboBox = ui.NewCombobox()
	typeOfSortComboBox.Append("Insertion")
	typeOfSortComboBox.Append("Bubble")
	typeOfSortComboBox.Append("BogoSort")
	typeOfSortComboBox.SetSelected(0) // start with insertion sort
	typeOfSortComboBox.OnSelected(func(*ui.Combobox) {
		sortSelected = typeOfSortComboBox.Selected()
		config.SortType = typeOfSortComboBox.Selected()
	})

	form.Append("Type of Sort ", typeOfSortComboBox, false) //TODO align left

	sortSpeedSlider = ui.NewSlider(2, 1000)
	sortSpeedSlider.SetValue(500)
	config.SortSpeed = 500
	sortSpeedSlider.OnChanged(func(*ui.Slider) {
		config.SortSpeed = sortSpeedSlider.Value()
	})

	form.Append("Sort Speed", sortSpeedSlider, false)

	iterationLabel = ui.NewLabel("")
	form.Append("Number of Iterations: ", iterationLabel, false)

	hbox.Append(area, true)

	mainwin.Show()
}

func main() {
	A = make([]int, 100)
	setup(A)
	shuffle(A)

	ui.Main(setupUI)
}
