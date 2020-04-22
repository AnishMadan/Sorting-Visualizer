package main

import (
	"math/rand"
	"time"

	"github.com/anishmadan/Sorting-Visualizer/sorting"

	"github.com/andlabs/ui"
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
	sortButton *ui.Button
	typeOfSort *ui.Combobox
	attrstr    *ui.AttributedString
	A          []int
)

type areaHandler struct{}

func (areaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {

	for i, x := range A {
		p := ui.DrawNewPath(ui.DrawFillModeWinding)
		p.NewFigure(0, 0)
		p.AddRectangle((float64)(i*10+15), 0, 5, (float64)(10*x))
		p.End()
		dp.Context.Fill(p, &ui.DrawBrush{Type: ui.DrawBrushTypeSolid, R: .75, G: .25, B: 0, A: 1})
		p.Free()
	}

	// fmt.Println(len(A))

	// p.NewFigure(0, 0)
	// p.AddRectangle(0, 0, 10, 10)
	// p.End()

	// p.NewFigure(10, 10)
	// p.LineTo(dp.ClipWidth-10, 10)
	// p.LineTo(dp.ClipWidth-10, dp.ClipHeight-10)
	// p.LineTo(10, dp.ClipHeight-10)
	// p.CloseFigure()
	// p.End()
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
	mainwin := ui.NewWindow("Sorting Example", 640, 480, true)
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
		sorting.InsertionSort(A, area)
		//area.QueueRedrawAll()
	})
	vbox.Append(sortButton, false)

	form := ui.NewForm()
	form.SetPadded(true)
	// TODO on OS X if this is set to 1 then the window can't resize; does the form not have the concept of stretchy trailing space?
	vbox.Append(form, false)

	typeOfSort = ui.NewCombobox()
	// note that the items match with the values of the uiDrawTextAlign values
	typeOfSort.Append("Insertion")
	typeOfSort.SetSelected(0) // start with insertion sort
	typeOfSort.OnSelected(func(*ui.Combobox) {
		area.QueueRedrawAll()
	})
	form.Append("Type of Sort", typeOfSort, false)

	hbox.Append(area, true)

	mainwin.Show()
}

func main() {
	A = make([]int, 50)
	setup(A)
	shuffle(A)

	ui.Main(setupUI)
}
