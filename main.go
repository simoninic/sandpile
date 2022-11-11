package main

import(
	//"fmt"
	"strconv"
	"os"
	//"image"
	"image/png"
)

func main() {

	size, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	pile, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	placement := os.Args[3]
	if (placement != "central" && placement != "random") {
		panic("Incorrect argument value for placement")
	}
	
	var coinBoard GameBoard

	coinBoard = SimulateSandpile(size, pile, placement) // 3 4 central


	canvasWidth := 1000
	img := coinBoard.DrawToCanvas(canvasWidth)
	out, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	png.Encode(out, img)
	out.Close()
}
