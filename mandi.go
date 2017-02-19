package sigi

/*
Mandelbrot renderer
*/

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func (self *Mandi) Init() {

	self.mandiW = self.View.AspectW / self.Coord.Depth
	self.mandiH = self.View.AspectH / self.Coord.Depth

	self.Map = make([]uint16, self.View.Width*self.View.Height)
}

func (self *Mandi) SetCoord(coord Coord) {
	self.Coord = coord
}

func (self *Mandi) SetView(view View) {
	self.View = view
}

func (self *Mandi) Center() {

	mx := real(self.Coord.Position) - self.mandiW/2
	my := imag(self.Coord.Position) + self.mandiH/2
	self.MandiPos = Point(complex(mx, my))
}

func (mandi *Mandi) Render(renderer Renderer) chan int {

	mandi.Init()
	mandi.Center()

	ch := make(chan int, 10)
	go renderer(mandi, ch)

	return ch
}

func (self *Mandi) ToImage() *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, self.View.Width, self.View.Height))
	for y := 0; y < self.View.Height; y++ {
		for x := 0; x < self.View.Width; x++ {
			n := self.Map[y*self.View.Width+x] % 255
			c := color.Gray{uint8(n)}
			img.Set(x, y, c)
		}
	}

	return img
}

func (self *Mandi) ToImage2() *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, self.View.Width, self.View.Height))
	pos := 0

	l := self.View.Height * self.View.Width
	for i := 0; i < l; i++ {
		n := uint8(self.Map[i] % 255)
		img.Pix[pos] = n
		img.Pix[pos+1] = n
		img.Pix[pos+2] = n
		//img.Pix[pos+3] = n
		pos += 4
	}

	return img
}

func (self *Mandi) ToPng(fn string) error {

	fmt.Println("Converting...")
	img := self.ToImage()

	//TODO: Build image from Map data
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("Enocding...")
	png.Encode(f, img)
	return nil
}
