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
}
