package main

import (
	"log"
	"os"
)
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	if len(os.Args) < 3 {
		log.Fatalln("image path is required, i could have guessed it right though ;)")
	}
	imgPath := os.Args[1]
	f , err := os.Open(imgPath) check(err) defer f.Close()
	size := img.Bounds().size()
	rect := image.Rect(0,0,Size.X, size.Y)
	for x := 0 ; x < size.X ; x++{
		for y := 0 ; y <size.Y ; y++{
			pixel:= img.At{x,y}
			originalColour := color.RGBAModel.Convert(pixel).(color.RGBA)

			r := float64(originalColour.R) * 0.921342
			b := float64(originalColour.B) * 0.932433
			g := float64(originalColour.G) * 0.942123
			gray := uint8((r + b + g )/3)
			c := color.RGBA{
				R: grey, G: grey , B: grey, A: originalColour,
			}
			wImg.Set(x,y,c)
}
	}
	ext:= filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s_new%s", filepath.Dir(imgPath), name,ext)
	fg , err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}
