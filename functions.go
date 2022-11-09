package main

import (
	"fmt"
	"math/rand"
	"math"
)

func SimulateSandpile(size, pile int, placement string) GameBoard {
	fmt.Println("starting simulation with ", placement)
	board := InitializeBoard(size)

	AddStartingCoins(board, pile, placement)

	fmt.Println("Initial Board")
	PrintBoard(board)

	board = ToppleSubboard(board, size)

	return board
}

func ToppleSubboard(board GameBoard, size int) GameBoard {
	for !IsStable(board) {
		for row := 0; row < size; row++ {
			for col := 0; col < size; col++ {
				if board[row][col] >= 4 {
					board = ToppleCell(board, row, col)	
				}		
			}
		}
	}
	return board
}

func AddStartingCoins(board GameBoard, pile int, placement string) {
	width := len(board)

	if placement == "central" {
		board[width/2][width/2] = pile
	} else { // placement == "random"

		volume := 100

		// optimized random number generation for relatively small piles
		if pile < volume {
			fmt.Println("small pile rng")
			randomNumbers := make([]int, volume)

			for i := 0; i < pile; i++ {
				index := rand.Intn(100)
				randomNumbers[index]++
			}

			for i := 0; i < volume; i++ {
				randRow := rand.Intn(width)
				randCol := rand.Intn(width)
				board[randRow][randCol] += randomNumbers[i]
			}
		} else { // optimized random number generation for relatively large piles
			fmt.Println("large pile rng")
			randomNumbers := make([]float64, volume)
			randomSum := 0.0

			// generate random coins 100 times
			for i := 0; i < volume; i++ {
				randCoins := rand.Float64()
				randomNumbers[i] = randCoins
				randomSum += randCoins
			}
			// convert to integers
			pileTracker := 0.0
			for i := 0; i < volume; i++ {
				randomNumbers[i] = math.Round(randomNumbers[i] * float64(pile) / randomSum)
				pileTracker += randomNumbers[i]
			}
			remainder := int(math.Round(float64(pile) - pileTracker))

			// add the random coins to 100 random cells of the board
			for i := 0; i < volume; i++ {
				randRow := rand.Intn(width)
				randCol := rand.Intn(width)
				if i < remainder {
					board[randRow][randCol] += 1
				}
				board[randRow][randCol] += int(randomNumbers[i])
			}
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

func IsStable(board GameBoard) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] >= 4 {
				return false
			}	
		}
	}
	return true
}

func ToppleCell(board GameBoard, row, col int) GameBoard {
	for board[row][col] >= 4 {
		board[row][col] -= 4
		if OnBoard(len(board), row-1, col) {
			board[row-1][col] += 1
		}
		if OnBoard(len(board), row+1, col) {
			board[row+1][col] += 1
		}
		if OnBoard(len(board), row, col-1) {
			board[row][col-1] += 1
		}
		if OnBoard(len(board), row, col+1) {
			board[row][col+1] += 1
		}
	}
	return board
}

func OnBoard(width, row, col int) bool {
	if (row >= 0 && row < width && col >= 0 && col < width) {
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

	n := len(board)
	c := make(chan [][]int, numProcs)

	for i := 0; i < numProcs; i++ {
		startIndex := i * (n / numProcs)
		endIndex := (i + 1) * (n / numProcs)
		if i < numProcs - 1 {
			go ToppleSubboard(board[startIndex:endIndex], size, c)
		} else {
			go ToppleSubboard(board[startIndex:], c)
		}
	}

	for i := 0; i < numProcs; i++ {
		miniBoard := <- c
		for j := 0; j < len(miniBoard); j++ {
			finalBoard = append(finalBoard, miniBoard) //should this be miniBoard[j]?

		}
	}
	/*
	range through numProcs
		divide the board into subboards that maintain  (so each col is col of entire board, each row is ~row of entire board / numProcs)
		each time call goroutine - SandpileSingleproc on each subboard
		
		get results back from asynchronous channel
		combine subboards together
			add together the coins that fall off if the subboards are adjacent (one directly on top of the other)
				this follows the condition last row index of subboard-A + 1 == first row index of subboard-B
	*/
}

// func SandpileSingleprocs(board Gameboard, c chan GameBoard) {

// }
