package main

import (
	"fmt"
	"proj/images"
)

func main() {
	fmt.Println("I am Monalisa")
	fmt.Println("I am Rupam")
	images.UTI("Boom Boom Bumer")

	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	for j := 340; j > 0; j-- {
		fmt.Println(j)
	}
	images.NewFT()
}
