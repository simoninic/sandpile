package main

import (
	"fmt"
	"math/rand"
	"time"
	//"runtime"
	//"math"
)

// SimulateSandpile: runs the sandpile process on a board of coins both serially and parallel-ly
// inputs: size - width of board (int), pile - number of coins on starting board (int), placement - central or random (string)
// outputs: two identical, final boards (GameBoard) - one derived serially, one derived parallel-ly
func SimulateSandpile(size, pile int, placement string) (GameBoard, GameBoard) {	

	//fmt.Println("starting simulation with ", placement)
	board1 := InitializeBoard(size)
	AddStartingCoins(board1, pile, placement)
	board2 := CopyBoard(board1)

	// fmt.Println("Initial Board")
	// PrintBoard(board1)

	numProcs := 8
	//println("Num CPU: ", runtime.NumCPU())
	println("Num Procs: ", numProcs)

	start := time.Now()
	//for !IsStable(board1, size, size) {

	SandpileMultiprocs(board1, numProcs)

	//}
	elapsed := time.Since(start)
	//fmt.Println()
	// fmt.Println("Final Board (Parallel)")
	// PrintBoard(board1)
	// fmt.Println()

	//fmt.Println("Initial Board")
	//PrintBoard(board2)

	start2 := time.Now()
	ToppleSubboardSerial(board2)
	elapsed2 := time.Since(start2)

	// fmt.Println()
	// fmt.Println("Final Board (Serial)")
	// PrintBoard(board2)

	fmt.Println("Time Elapsed (Parallel): ", elapsed)
	fmt.Println("Time Elapsed (Serial): ", elapsed2)
	// fmt.Println()
	// fmt.Println("Total # Runs of Sandpile Multiprocs: ", len(elapsedPar))

	// for i := 0; i < len(elapsedPar); i+=25 {
	// 	fmt.Println("Time Elapsed - Sandpile Multiproc Run ", i, ": ", elapsedPar[i])
	// }
	fmt.Println("Do boards match?: ", BoardsMatch(board1, board2))

	return board1, board2
}

// InitializeBoard: sets up an empty board of given size
// input: size - width of board (int)
// output: the empty board of given size (GameBoard)
func InitializeBoard(size int) GameBoard {
	var board GameBoard
	board = make([]([]int), size)
  
	for r := range board {
	  board[r] = make([]int, size)
	}
	return board
}

// AddStartingCoins: distribute 'pile' number of coins across the board
// inputs: board (GameBoard), pile - number of coins to add ot board (int),
//         placement - central [all coins in center] or random (string)
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

// CopyBoard: copies the current board's values onto a new board
// inputs: given, current board (GameBoard)
// output: a new board (GameBoard) with the same values
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

// BoardsMatch: determines if two boards are identical (assumes the boards are rectangles)
// inputs: two boards (GameBoards)
// output: whether or not the boards share identical values across all cells (boolean)
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

// ToppleSubboardSerial: topples every topple-able cell in the board until the board is stable
// inputs: the given board (GameBoard)
func ToppleSubboardSerial(board GameBoard) {
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
}

// ToppleCellSerial: topples a single cell until it has less than 4 coins
// inputs: board (Gameboard), the number of rows and columns, as well as the row and column index (all ints)
func ToppleCellSerial(board GameBoard, numRows, numCols, row, col int) {
	for board[row][col] >= 4 {
		UpdateOnBoardCells(board, numRows, numCols, row, col)
	}
}

// SandpileMultiprocs: runs the topple process throughout the entire board as it gets broken down into sub-boards (parallelization)
// inputs: board (GameBoard), number of processors (int)
func SandpileMultiprocs(board GameBoard, numProcs int) {
	//var startPar []time.Time
	var elapsedPar []time.Duration

	size := len(board)

	for !IsStable(board, size, size) {
		startPar := time.Now()
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
		}
		HandleLostCoins(board, mergedBoard)
		elapsedPar = append(elapsedPar, time.Since(startPar))
	}

	fmt.Println("Total # Runs of Sandpile Multiprocs: ", len(elapsedPar))

	for i := 0; i < len(elapsedPar); i+=25 {
		fmt.Println("Time Elapsed - Sandpile Multiproc Run ", i, ": ", elapsedPar[i])
	}
}

// HandleLostCoins: look for adjacent subboards where their "fallen off" coins can be added accordingly
// inputs: board (GameBoard), the complete board as a slice of SubBoards
func HandleLostCoins(board GameBoard, mergedBoard []SubBoard) {
	for i := 0; i < len(mergedBoard); i++ {
		for j := i + 1; j < len(mergedBoard); j++ { // loop through twice to compare two different miniBoards together
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

			// looking at miniBoards the other way
			if miniBoardB.endRow == miniBoardA.startRow {
				numCols := len(board)

				// add the fallen off coins
				for colIndex := 0; colIndex < numCols; colIndex++ {
					AddFallenCoins(board, mergedBoard, miniBoardB, miniBoardA, colIndex)
				}
			}
		}
	}
}

// AddFallenCoins: adds coins that have fallen off the subBoards back to the main board while merging
// inputs: the board (GameBoard), the completeBoard as a slice of SubBoards, two adjacent subBoards, the column index (int)
func AddFallenCoins(board GameBoard, mergedBoard []SubBoard, miniBoardA SubBoard, miniBoardB SubBoard, colIndex int) {
	if len(miniBoardB.topFalloff) > 0 {
		board[miniBoardA.endRow - 1][colIndex] += *(miniBoardB.topFalloff)[colIndex]
	}
	if len(miniBoardA.bottomFalloff) > 0 {
		board[miniBoardB.startRow][colIndex] += *(miniBoardA.bottomFalloff)[colIndex]
	}
}

// ToppleSubboards: topples all the cells in each subboard
// inputs: board (GameBoard), the start and end rows of the subboard (ints), the channel that returns SubBoard info
func ToppleSubboard(board GameBoard, startRow int, endRow int, c chan SubBoard) {
	var miniBoard SubBoard

	numRows := len(board)
	numCols := len(board[0])

	// set up and initialize top and bottom fallof slices
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

	// topples cells in subboard until subboard is stable
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

	// track miniBoard info to pass on from the channel
	miniBoard.startRow = startRow
	miniBoard.endRow = endRow
	miniBoard.topFalloff = topFalloff
	miniBoard.bottomFalloff = bottomFalloff

	//fmt.Println("fall off slices")
	//fmt.Println(topFalloff)
	//fmt.Println(bottomFalloff)

	c <- miniBoard
}

// ToppleCell: topples the cell until it has less than 4 coins
// inputs: board (GameBoard), numRows & numCols (ints) - borders of sub-board, row and column indices (int),
//         top and bottom falloff numbers just above & below the borders of the sub-board
func ToppleCell(board GameBoard, numRows, numCols, row, col int, topFalloff []*int, bottomFalloff []*int) {
	for board[row][col] >= 4 { // while the cell is topple-able
		UpdateOnBoardCells(board, numRows, numCols, row, col)
		if row == 0 && col >= 0 && col < numCols {
			(*topFalloff[col])++
		}
		if row == (numRows - 1) && col >= 0 && col < numCols {
			(*bottomFalloff[col])++
		}
	}
}

// UpdateOnBoardCells: removes 4 coins from target cell and adds 1 coin to the 4 adjacent cells
// inputs: board (GameBoard), numRows & numCols - bottom & right borders of sub-board, row & col - indices of target cell
func UpdateOnBoardCells(board GameBoard, numRows, numCols, row, col int) {
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

// IsStable: checks if a board is stable (no cells have 4 coins or more)
// inputs: the given board (GameBoard), numRows & numCols - the bottom & right borders of sub-board
// output: if board is stable (bool)
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

// OnBoard: checks if target cell is on the board
// inputs: width of board, numRows & numCols - bottom & right borders of board, row & col indices of target cell (all ints)
// output: if target cell is on the board (bool)
func OnBoard(width, numRows, numCols, row, col int) bool {
	if (row >= 0 && row < numRows && col >= 0 && col < numCols) {
		return true
	}
	return false
}

// PrintBoard: prints out the board's values
// input: the board (GameBoard)
func PrintBoard(board GameBoard) {
	for row := 0; row < len(board); row++ {
		fmt.Println(board[row])
	}
}
