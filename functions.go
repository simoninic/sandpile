package main

import (
	"fmt"
	"math/rand"
	"time"
	//"math"
)

func SimulateSandpile(size, pile int, placement string) (GameBoard, GameBoard) {	

	//fmt.Println("starting simulation with ", placement)
	board1 := InitializeBoard(size)
	AddStartingCoins(board1, pile, placement)
	board2 := CopyBoard(board1)

	fmt.Println("Initial Board")
	PrintBoard(board1)

	start := time.Now()
	numProcs := 2
	for !IsStable(board1, size, size) {
		SandpileMultiprocs(board1, numProcs)
	}
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Println("Final Board (Parallel)")
	PrintBoard(board1)
	fmt.Println()


	// board2 := InitializeBoard(size)
	// AddStartingCoins(board2, pile, placement)
	fmt.Println("Initial Board")
	PrintBoard(board2)

	start2 := time.Now()
	ToppleSubboardSerial(board2)
	elapsed2 := time.Since(start2)

	fmt.Println()
	fmt.Println("Final Board (Serial)")
	PrintBoard(board2)

	fmt.Println("Time Elapsed (Parallel): ", elapsed)
	fmt.Println("Time Elapsed (Serial): ", elapsed2)

	fmt.Println("Do boards match?: ", BoardsMatch(board1, board2))

	return board1, board2
}

func CopyBoard(board GameBoard) GameBoard {
	size := len(board)
	var newBoard GameBoard
	newBoard = make([]([]int), size)

	for r := range board {
		newBoard[r] = make([]int, size)
	}
  
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			newBoard[i][j] = board[i][j]
		}
	}

	return newBoard
}

func BoardsMatch(board1, board2 GameBoard) bool {
	if (len(board1) != len(board2)) || (len(board1[0]) != len(board2[0])) {
		return false
	}
	for i := 0; i < len(board1); i++ {
		for j := 0; j < len(board1[0]); j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}

func ToppleSubboardSerial(board GameBoard) GameBoard {
	size := len(board)
	for !IsStable(board, size, size) {
		for row := 0; row < size; row++ {
			for col := 0; col < size; col++ {
				if board[row][col] >= 4 {
					ToppleCellSerial(board, size, size, row, col)	
				}		
			}
		}
	}
	return board
}

func ToppleSubboard(board GameBoard, startRow int, endRow int, c chan SubBoard) {
	var miniBoard SubBoard

	numRows := len(board)
	numCols := len(board[0])

	topFalloff := make([]*int, numCols)
	for i := 0; i < numCols; i++ {
		topFalloff[i] = new(int)
		*topFalloff[i] = 0
	}
	bottomFalloff := make([]*int, numCols)
	for i := 0; i < numCols; i++ {
		bottomFalloff[i] = new(int)
		*bottomFalloff[i] = 0
	}
	//fmt.Println("brownie: ", bottomFalloff)



	//fmt.Println("subboard initial")
	//PrintBoard(board)


	//fmt.Println(numRows)
	//fmt.Println(numCols)

	//fmt.Println("Macaron: ", startRow, "      ", endRow)

	for !IsStable(board, numRows, numCols) {
		//fmt.Println("Iroh")
		for row := 0; row < numRows; row++ {
			for col := 0; col < numCols; col++ {
				if board[row][col] >= 4 {
					//fmt.Println("Thomas")
					ToppleCell(board, numRows, numCols, row, col, topFalloff, bottomFalloff)	
				}		
			}
		}
	}
	//fmt.Println("subboard final result:")
	//PrintBoard(board)
	//NEED TO KEEP TRACK OF COINS THAT FALL OUT!!!!!

	miniBoard.startRow = startRow
	miniBoard.endRow = endRow
	miniBoard.topFalloff = topFalloff
	miniBoard.bottomFalloff = bottomFalloff

	//fmt.Println("fall off slices")
	//fmt.Println(topFalloff)
	//fmt.Println(bottomFalloff)

	// divide the boards correctly (use only subboards)
	// include another thing in channel: the overflow of values going into the top and bottom sides of the board...

	c <- miniBoard
}

func AddStartingCoins(board GameBoard, pile int, placement string) {
	width := len(board)

	if placement == "central" {
		board[width/2][width/2] = pile
	} else { // placement == "random"
		volume := 100
		// add the coins randomly into 100 different cells (can repeat into same cell)
		remainder := pile % volume
		
		for i := 0; i < volume; i++ {
			randRow := rand.Intn(width)
			randCol := rand.Intn(width)
			if i < remainder {
				board[randRow][randCol]++
			}
			board[randRow][randCol] += pile/volume
		}
	}
}

func InitializeBoard(size int) GameBoard {
	var board GameBoard
	board = make([]([]int), size)
  
	for r := range board {
	  board[r] = make([]int, size)
	}
	return board
}

func IsStable(board GameBoard, numRows, numCols int) bool {
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if board[row][col] >= 4 {
				return false
			}	
		}
	}
	return true
}

func ToppleCell(board GameBoard, numRows, numCols, row, col int, topFalloff []*int, bottomFalloff []*int) {
	// topFalloff := make([]int, numCols)
	// bottomFalloff := make([]int, numCols)

	for board[row][col] >= 4 {
		board[row][col] -= 4
		if OnBoard(len(board), numRows, numCols, row-1, col) {
			board[row-1][col] += 1
		}
		if row == 0 && col >= 0 && col < numCols {
			//fmt.Println("length of top falloff: ", len(topFalloff), "   col: ", col)
			(*topFalloff[col])++
		}
		if row == (numRows - 1) && col >= 0 && col < numCols {
			//fmt.Println("length of bot falloff: ", len(bottomFalloff), "   col: ", col)
			//fmt.Println("bottom falloff: ", bottomFalloff[col])
			(*bottomFalloff[col])++
			//fmt.Println("meow")
		}
		if OnBoard(len(board), numRows, numCols, row+1, col) {
			board[row+1][col] += 1
		}
		if OnBoard(len(board), numRows, numCols, row, col-1) {
			board[row][col-1] += 1
		}
		if OnBoard(len(board), numRows, numCols, row, col+1) {
			board[row][col+1] += 1
		}
	}
}
func ToppleCellSerial(board GameBoard, numRows, numCols, row, col int) {
	for board[row][col] >= 4 {
		board[row][col] -= 4
		if OnBoard(len(board), numRows, numCols, row-1, col) {
			board[row-1][col] += 1
		}
		if OnBoard(len(board), numRows, numCols, row+1, col) {
			board[row+1][col] += 1
		}
		if OnBoard(len(board), numRows, numCols, row, col-1) {
			board[row][col-1] += 1
		}
		if OnBoard(len(board), numRows, numCols, row, col+1) {
			board[row][col+1] += 1
		}
	}
}

func OnBoard(width, numRows, numCols, row, col int) bool {
	if (row >= 0 && row < numRows && col >= 0 && col < numCols) {
		return true
	}
	return false
}

func PrintBoard(board GameBoard) {
	for row := 0; row < len(board); row++ {
		fmt.Println(board[row])
	}
}


///// Parallelization
func SandpileMultiprocs(board GameBoard, numProcs int) {
	//size := len(board)

	// set up final empty board
	//var finalBoard GameBoard

	//fmt.Println("starting final board")
	//fmt.Println(finalBoard)

	n := len(board)
	c := make(chan SubBoard, numProcs)

	for i := 0; i < numProcs; i++ {
		startIndex := i * (n / numProcs)
		endIndex := (i + 1) * (n / numProcs)
		//fmt.Println("s: ", startIndex, "     e: ", endIndex)
		if i < numProcs - 1 {
			go ToppleSubboard(board[startIndex:endIndex], startIndex, endIndex, c)
		} else {
			go ToppleSubboard(board[startIndex:], startIndex, endIndex, c)
		}
	}

	var mergedBoard []SubBoard

	for i := 0; i < numProcs; i++ {
		// startIndex := i * (n / numProcs)
		// endIndex := (i + 1) * (n / numProcs)
		miniBoard := <- c
		mergedBoard = append(mergedBoard, miniBoard)

		//fmt.Println("merging boards")

		//fmt.Println(miniBoard)
	}

	HandleLostCoins(board, mergedBoard)
}

func HandleLostCoins(board GameBoard, mergedBoard []SubBoard) {
	for i := 0; i < len(mergedBoard); i++ {
		for j := 0; j < len(mergedBoard); j++ { // loop through twice to compare two different miniBoards together
			// define both mini boards
			miniBoardA := mergedBoard[i]
			miniBoardB := mergedBoard[j]
			
			// if miniBoards are adjacent, then add the fallen coins to each others first/last row
			if miniBoardA.endRow == miniBoardB.startRow {
				numCols := len(board)

				// add the fallen off coins
				for colIndex := 0; colIndex < numCols; colIndex++ {
					AddFallenCoins(board, mergedBoard, miniBoardA, miniBoardB, colIndex)
				}
			}
		}
	}
}

// AddFallenCoins: adds coins that have fallen off the subBoards back to the main board while merging
func AddFallenCoins(board GameBoard, mergedBoard []SubBoard, miniBoardA SubBoard, miniBoardB SubBoard, colIndex int) {
	if len(miniBoardB.topFalloff) > 0 {
		board[miniBoardA.endRow - 1][colIndex] += *(miniBoardB.topFalloff)[colIndex]
	}
	if len(miniBoardA.bottomFalloff) > 0 {
		board[miniBoardB.startRow][colIndex] += *(miniBoardA.bottomFalloff)[colIndex]
	}
}
