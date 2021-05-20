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
}
