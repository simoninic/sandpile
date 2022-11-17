package main

import(
	"strconv"
	"os"
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
	img := parallelBoard.DrawToCanvas()
	MakeImage(img, "rename_parallel.png")

	img2 := serialBoard.DrawToCanvas()
	MakeImage(img2, "rename_serial.png")
}
