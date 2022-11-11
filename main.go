package main

import(
	//"fmt"
	"strconv"
	"os"
	//"image"
	"image/png"
)

func main() {

	// reads in size, pile and placement input arguments
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
	
	// run sandpile simulations serially and parallel-ly to create the two respective final boards
	var parallelBoard GameBoard
	var serialBoard GameBoard
	parallelBoard, serialBoard = SimulateSandpile(size, pile, placement) // 3 4 central


	// create the 2 PNG images resulting from serial and parallel strategies
	canvasWidth := 1000
	img := parallelBoard.DrawToCanvas(canvasWidth)
	out, err := os.Create("rename_parallel.png")
	if err != nil {
		panic(err)
	}
	png.Encode(out, img)
	out.Close()

	img2 := serialBoard.DrawToCanvas(canvasWidth)
	out2, err2 := os.Create("rename_serial.png")
	if err2 != nil {
		panic(err2)
	}
	png.Encode(out2, img2)
	out2.Close()
}
