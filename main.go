package main

import(
	"fmt"
	"os"
	"strconv"
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

	fmt.Println("")
	fmt.Println("Final Board")
	PrintBoard(coinBoard)
}
