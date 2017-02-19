package sigi

type Point complex128
type Renderer func(mandi *Mandi, ch chan int)

type Mandi struct {
	View  View
	Coord Coord

	MandiPos       Point
	mandiW, mandiH float64

	Map []uint16
}

type View struct {
	Width, Height    int
	AspectW, AspectH float64
}

type Coord struct {
	Position Point
	Depth    float64
	Time     uint64
}

func NewView(w, h int) (v View) {
	v.Width = w
	v.Height = h
	v.AspectW = 3
	v.AspectH = (float64(h) * 3) / float64(w)
	return
}

func NewCoord() (v Coord) {
	v.Position = complex(0, 0)
	v.Depth = 0.5
	v.Time = 2000
	return
}

func NewMandi(width, height int) (mandi Mandi) {

	mandi.SetCoord(NewCoord())
	mandi.SetView(NewView(width, height))

	return
}
