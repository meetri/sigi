package sigi

import (
	"math/cmplx"
	"runtime"
)

func SlowRenderer(mandi *Mandi, ch chan int) {

	numcores := runtime.GOMAXPROCS(0)
	if numcores > runtime.NumCPU() {
		numcores = runtime.NumCPU()
	}
	cores := make(chan int, numcores)
	for y := 0; y < mandi.View.Height; y = y + 1 {
		cores <- 1
		go slowRowRenderer(mandi, 0, y, cores)
	}

	close(ch)

}

func slowRowRenderer(mandi *Mandi, x, y int, ch chan int) {

	pos := y*mandi.View.Width + x
	buf := mandi.Map[pos : pos+mandi.View.Width-x]

	sx := mandi.mandiW / float64(mandi.View.Width)
	sy := mandi.mandiH / float64(mandi.View.Height)

	cx := real(mandi.MandiPos)
	cy := -imag(mandi.MandiPos) + (sy * float64(y))

	c := complex(cx, cy)
	for p := uint64(0); p < uint64(mandi.View.Width-x); p++ {
		z := complex(sx, sy)
		escape := uint16(0)
		for i := uint16(0); i < uint16(mandi.Coord.Time); i++ {
			z = z*z + c
			if cmplx.Abs(z) > 2 {
				escape = i
				break
			}
		}
		c = complex(real(c)+sx, imag(c))
		buf[p] = escape
	}
	<-ch

}
