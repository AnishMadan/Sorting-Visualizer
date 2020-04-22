package main

import (
	"math/rand"
	"time"

	"github.com/andlabs/ui"
	"github.com/anishmadan/Sorting-Visualizer/sorting"
)

func setup(a []int) {
	for i := range a {
		a[i] = i + 1
	}
}

func shuffle(a []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(a); n > 0; n-- {
		randIndex := r.Intn(n)
		a[n-1], a[randIndex] = a[randIndex], a[n-1]
	}
}

var (
	sortButton    *ui.Button
	shuffleButton *ui.Button
	typeOfSort    *ui.Combobox
	attrstr       *ui.AttributedString
	A             []int
	sortSelected  int
)

type areaHandler struct{}

func (areaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {

	for i, x := range A {
		p := ui.DrawNewPath(ui.DrawFillModeWinding)
		p.NewFigure(0, 0)
		p.AddRectangle((float64)(i*10+15), 0, 5, (float64)(5*x))
		p.End()
		dp.Context.Fill(p, &ui.DrawBrush{Type: ui.DrawBrushTypeSolid, R: .75, G: .25, B: 0, A: 1})
		p.Free()
	}

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
		if sortSelected == 0 {
			sorting.InsertionSort(A)
		}
		area.QueueRedrawAll()
	})
	vbox.Append(sortButton, false)

	shuffleButton = ui.NewButton("Shuffle")
	shuffleButton.OnClicked(func(*ui.Button) {
		shuffle(A)
		area.QueueRedrawAll()
	})
	vbox.Append(shuffleButton, false)

	form := ui.NewForm()
	form.SetPadded(true)
	// TODO on OS X if this is set to 1 then the window can't resize; does the form not have the concept of stretchy trailing space?
	vbox.Append(form, false)

	typeOfSort = ui.NewCombobox()
	// note that the items match with the values of the uiDrawTextAlign values
	typeOfSort.Append("Insertion")
	typeOfSort.SetSelected(0) // start with insertion sort
	typeOfSort.OnSelected(func(*ui.Combobox) {
		sortSelected = typeOfSort.Selected()
	})

	form.Append("Type of Sort", typeOfSort, false)

	hbox.Append(area, true)

	mainwin.Show()
}

func main() {
	A = make([]int, 100)
	setup(A)
	shuffle(A)

	ui.Main(setupUI)
}
