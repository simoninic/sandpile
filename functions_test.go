package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Tests for the function: UpdateOnBoardCells()
func TestUpdateOnBoardCells(t *testing.T) {

	// test object (inputs and answer)
	type test struct {
		board                        GameBoard
		numRows, numCols, row, col   int
		answer                       int
	}
	
	// obtain directories and input & output files
	inputDirectory := "tests/UpdateOnBoardCells/input/"
	outputDirectory := "tests/UpdateOnBoardCells/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	// set up input variables
	var board GameBoard
	var numRows int
	var numCols int
	var row int
	var col int

	// go through the input and output files and set the test values
	for i := range inputFiles {
		
		board, numRows, numCols, row, col = ReadBoardFourIntsFromFile(inputDirectory, inputFiles[i])
		tests[i].board = board
		tests[i].numRows = numRows
		tests[i].numCols = numCols
		tests[i].row = row
		tests[i].col = col

		tests[i].answer = ReadIntegerFromFile(outputDirectory, outputFiles[i])
	}

	// test the function and prints out whether the test passes or fails
	fmt.Println("Testing UpdateOnBoardCells()")
	for _, test := range tests {
		output := UpdateOnBoardCells(test.board, test.numRows, test.numCols, test.row, test.col)
		
		if output == test.answer {
			fmt.Println("Correct! Answer: ", output)
		} else {
			fmt.Println("The function's output is ", output, ". The correct answer is ", test.answer, ".")
		}
		fmt.Println()
	}
}

// Tests for the function: IsStable()
func TestIsStable(t *testing.T) {

	// test object (inputs and answer)
	type test struct {
		board                        GameBoard
		numRows, numCols, row, col   int
		answer                       bool
	}
	
	// obtain directories and input & output files
	inputDirectory := "tests/IsStable/input/"
	outputDirectory := "tests/IsStable/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	// set up input variables
	var board GameBoard
	var numRows int
	var numCols int

	// go through the input and output files and set the test values
	for i := range inputFiles {
		
		board, numRows, numCols = ReadBoardTwoIntsFromFile(inputDirectory, inputFiles[i])
		tests[i].board = board
		tests[i].numRows = numRows
		tests[i].numCols = numCols

		tests[i].answer = ReadBooleanFromFile(outputDirectory, outputFiles[i])
	}

	// test the function and prints out whether the test passes or fails
	fmt.Println("Testing IsStable()")
	for _, test := range tests {
		output := IsStable(test.board, test.numRows, test.numCols)

		
		if output == test.answer {
			fmt.Println("Correct! Answer: ", output)
		} else {
			fmt.Println("The function's output is ", output, ". The correct answer is ", test.answer, ".")
		}
		fmt.Println()
	}
}

// Tests for the function: OnBoard()
func TestOnBoard(t *testing.T) {

	// test object (inputs and answer)
	type test struct {
		board                        GameBoard
		numRows, numCols, row, col   int
		answer                       bool
	}

	//we now will need to create our array of tests
	numTests := 3
	tests := make([]test, numTests)

	// normal, true case
	index := 0
	tests[index].numRows = 5
	tests[index].numCols = 5
	tests[index].row = 2
	tests[index].col = 2
	tests[index].answer = true

	// false case
	index++
	tests[index].numRows = 3
	tests[index].numCols = 3
	tests[index].row = 3
	tests[index].col = 4
	tests[index].answer = false

	// border case
	index++
	tests[index].numRows = 4
	tests[index].numCols = 4
	tests[index].row = 3
	tests[index].col = 3
	tests[index].answer = true

	fmt.Println("Testing OnBoard()")
	// test the function and prints out whether the test passes or fails
	for _, test := range tests {
		output := OnBoard(test.numRows, test.numCols, test.row, test.col)

		if output == test.answer {
			fmt.Println("Correct! Answer: ", output)
		} else {
			fmt.Println("The function's output is ", output, ". The correct answer is ", test.answer, ".")
		}
	}
	fmt.Println()
}

// Tests for the function: ToppleCell()
func TestToppleCell(t *testing.T) {

	// test object (inputs and answer)
	type test struct {
		board                        GameBoard
		numRows, numCols, row, col   int
		topFalloff                   []*int
		bottomFalloff                []*int
		answer                       GameBoard
	}

	//we now will need to create our array of tests
	numTests := 1
	tests := make([]test, numTests)

	// normal, true case
	index := 0
	tests[index].numRows = 3
	tests[index].numCols = 3
	var board GameBoard
	board = make([]([]int), tests[index].numRows)
	for r := range board {
	  board[r] = make([]int, tests[index].numCols)
	}
	board[0][0] = 0
	board[0][1] = 0
	board[0][2] = 0
	board[1][0] = 0
	board[1][1] = 4
	board[1][2] = 0
	board[2][0] = 0
	board[2][1] = 0
	board[2][2] = 0
	tests[index].board = board
	tests[index].row = 1
	tests[index].col = 1
	tests[index].topFalloff = make([]*int, tests[index].numRows)
	tests[index].bottomFalloff = make([]*int, tests[index].numCols)
	var answerBoard GameBoard
	answerBoard = make([]([]int), tests[index].numRows)
	for r := range answerBoard {
	  answerBoard[r] = make([]int, tests[index].numCols)
	}
	board[0][0] = 0
	board[0][1] = 1
	board[0][2] = 0
	board[1][0] = 1
	board[1][1] = 0
	board[1][2] = 1
	board[2][0] = 0
	board[2][1] = 1
	board[2][2] = 0
	tests[index].answer = answerBoard

	fmt.Println("Testing ToppleCell()")
	// test the function and prints out whether the test passes or fails
	for _, test := range tests {
		ToppleCell(test.board, test.numRows, test.numCols, test.row, test.col, test.topFalloff, test.bottomFalloff)

		output := test.board

		if BoardsMatch(output, test.board) {
			fmt.Println("Correct! Answer: ")
			PrintBoard(output)
		} else {
			fmt.Println("The function's output is...")
			PrintBoard(output)
			fmt.Println("The correct answer is...")
			PrintBoard(test.answer)
		}
	}
	fmt.Println()
}

// ReadBoardFourIntsFromFile: reads the given board and accompanying untegers
func ReadBoardFourIntsFromFile(directory string, inputFile os.FileInfo) (GameBoard, int, int, int, int) {
	fileName := inputFile.Name() //grab file name

	// read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	// format input line by line
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	// read in the first lines which are just integers
	numInts := 4
	intParams := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		intParams[i], err = strconv.Atoi((strings.Split(inputLines[i], " "))[0])
		if err != nil {
			panic(err)
		}
	}
	// initialize parameters' values
	numRows := intParams[0]
	numCols := intParams[1]
	row := intParams[2]
	col := intParams[3]

	index := numInts + 1

	//make the board
	var board [][]int
	board = make([]([]int), numRows)
	for r := range board {
		board[r] = make([]int, numCols)
	}

	// read the remaining part of the file, which is the given board
	for i := index + 1; i < len(inputLines); i++ {

		//read out the current line
		currentRow := strings.Split(inputLines[i], " ")

		// go integer by integer in the current line
		for j, cell := range currentRow {
			currentValue, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			// add the integer to the board
			board[i - index - 1][j] = currentValue
		}
	}
	return board, numRows, numCols, row, col
}

// ReadBoardFourIntsFromFile: reads the given board and accompanying 2 integers
func ReadBoardTwoIntsFromFile(directory string, inputFile os.FileInfo) (GameBoard, int, int) {
	fileName := inputFile.Name() //grab file name

	// read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	// format input line by line
	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	// read in the first lines which are just integers
	numInts := 2
	intParams := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		intParams[i], err = strconv.Atoi((strings.Split(inputLines[i], " "))[0])
		if err != nil {
			panic(err)
		}
	}
	// initialize parameters' values
	numRows := intParams[0]
	numCols := intParams[1]

	index := numInts + 1

	//make the board
	var board [][]int
	board = make([]([]int), numRows)
	for r := range board {
		board[r] = make([]int, numCols)
	}

	// read the remaining part of the file, which is the given board
	for i := index + 1; i < len(inputLines); i++ {

		//read out the current line
		currentRow := strings.Split(inputLines[i], " ")

		// go integer by integer in the current line
		for j, cell := range currentRow {
			currentValue, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			// add the integer to the board
			board[i - index - 1][j] = currentValue
		}
	}
	return board, numRows, numCols
}

// func ReadFourIntsFromFile(directory string, inputFile os.FileInfo) (int, int, int, int) {
// 	fileName := inputFile.Name() //grab file name

// 	// read in the input file
// 	fileContents, err := ioutil.ReadFile(directory + fileName)
// 	if err != nil {
// 		panic(err)
// 	}


// 	// format input line by line
// 	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

// 	numInts := 4
// 	intParams := make([]int, numInts)

// 	// read the remaining part of the file, which is the given board

// 	for i := 0; i < numInts; i++ {

// 		//read out the current line
// 		currentRow := strings.Split(inputLines[i], " ")
// 		fmt.Println(currentRow)

// 		// go integer by integer in the current line
// 		intParams[i], err = strconv.Atoi(currentRow[0])
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	// initialize parameters
// 	numRows := intParams[0]
// 	numCols := intParams[1]
// 	row := intParams[2]
// 	col := intParams[3]
	
// 	return numRows, numCols, row, col
// }




// func ReadBoardFromFile(directory string, inputFile os.FileInfo) GameBoard {
// 	fileName := inputFile.Name() //grab file name

// 	//now, read in the input file
// 	fileContents, err := ioutil.ReadFile(directory + fileName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//first, read lines and split along blank space
// 	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

// 	//make the map that will store our frequency map
// 	frequencyMap := make([][]int)

// 	//each line of the file corresponds to a single line of the frequency map
// 	for _, inputLine := range inputLines {

// 		//read out the current line
// 		currentRow := strings.Split(inputLine, " ")
// 		//currentLine has two strings corresponding to the key and value

// 		frequencyMap = append(frequencyMap, currentRow)

// 		for j, inputCell := range currentRow {
// 			currentValue, err := strconv.Atoi(inputCell)
// 			if err != nil {
// 				panic(err)
// 			}
// 			frequencyMap[currentRow] = append(frequencyMap[currentRow], currentValue)
// 		}

// 		//if we make it here, everything is OK, so append to the input map
		
// 	}

// 	return frequencyMap
// }

func ReadBooleanFromFile(directory string, file os.FileInfo) bool {
	//now, consult the associated output file.
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.ParseBool(outputLines[0])

	if err != nil {
		panic(err)
	}

	return answer
}

func ReadIntegerFromFile(directory string, file os.FileInfo) int {
	//now, consult the associated output file.
	fileName := file.Name() //grab file name

	//now, read out the file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	//trim out extra space and store as a slice of strings, each containing one line.
	outputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	//parse the float
	answer, err := strconv.Atoi(outputLines[0])

	if err != nil {
		panic(err)
	}

	return answer
}


func ReadFilesFromDirectory(directory string) []os.FileInfo {
	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory: " + directory)
	}

	return dirContents
}


func AssertEqualAndNonzero(length0, length1 int) {
	if length0 == 0 {
		panic("No files present in input directory.")
	}
	if length1 == 0 {
		panic("No files present in output directory.")
	}
	if length0 != length1 {
		panic("Number of files in directories doesn't match.")
	}
}


// func TestRichness(t *testing.T) {
// 	//first, declare our test type
// 	type test struct {
// 		frequencyMap map[string]int
// 		answer       int
// 	}

// 	inputDirectory := "tests/Richness/input/"
// 	outputDirectory := "tests/Richness/output/"

// 	inputFiles := ReadFilesFromDirectory(inputDirectory)
// 	outputFiles := ReadFilesFromDirectory(outputDirectory)

// 	//assert that files are non-empty and have the same length
// 	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

// 	//we now will need to create our array of tests
// 	tests := make([]test, len(inputFiles))

// 	//first, range through the input and output files and set the test values
// 	for i := range inputFiles {
// 		tests[i].frequencyMap = ReadFrequencyMapFromFile(inputDirectory, inputFiles[i])
// 		tests[i].answer = ReadIntegerFromFile(outputDirectory, outputFiles[i])
// 	}

// 	for i, test := range tests {
// 		outcome := Richness(test.frequencyMap)

// 		if outcome != test.answer {
// 			t.Errorf("Error! For input test dataset %d, your code gives %d, and the correct richness is %d", i, outcome, test.answer)
// 		} else {
// 			fmt.Println("Correct! When the frequency map is", test.frequencyMap, "the richness is", test.answer)
// 		}
// 	}

// }

// func TestCalcDistance(t *testing.T) {
// 	//first, declare our test type
// 	type testObj struct {
// 		p1 OrderedPair
// 		p2 OrderedPair
// 		answer float64
// 	}

// 	num_tests := 2
// 	tests := make([]testObj, num_tests)

// 	// normal case test for distance
// 	index := 0
// 	tests[index].p1.x = 0.0
// 	tests[index].p1.y = 0.0
// 	tests[index].p2.x = 3.0
// 	tests[index].p2.y = 4.0
// 	tests[index].answer = 5.0

// 	// test case for distance 0
// 	index++
// 	tests[index].p1.x = 3.0
// 	tests[index].p1.y = 4.0
// 	tests[index].p2.x = 3.0
// 	tests[index].p2.y = 4.0
// 	tests[index].answer = 0

// 	// range over the tests and check if the function produces the correct output
// 	for i := range tests {
// 		output := CalcDistance(tests[i].p1, tests[i].p2)
// 		output = roundFloat(output, 1)

// 		fmt.Println("ComputeDistance")
// 		fmt.Println("Test ", i, ":")
// 		// if function output and correct answer mismatch
// 		if output != tests[i].answer {
// 			fmt.Println("Incorrect! The function's output is ", output, ". The correct answer is ", tests[i].answer, ".")
// 		} else { // if function output and correct answer match
// 			fmt.Println("Correct! Answer: ", output)
// 		}
// 	}
// 	fmt.Println("")
// }

// func TestComputeSingleForce(t *testing.T) {
// 	//first, declare our test type
// 	type testObj struct {
// 		s1 *Star
// 		s2 *Star
// 		answer OrderedPair
// 	}

// 	num_tests := 1
// 	tests := make([]testObj, num_tests)

// 	// normal test case
// 	index := 0
// 	// tests[index].s1.position.x = 50.0
// 	// tests[index].s1.position.y = 60.0
// 	// tests[index].s2.position.x = 90.0
// 	// tests[index].s2.position.y = 90.0
// 	// tests[index].s1.mass = 20.0
// 	// tests[index].s2.mass = 30.0
// 	(*tests[index].s1).position.x = 50.0
// 	(*tests[index].s1).position.y = 60.0
// 	(*tests[index].s2).position.x = 90.0
// 	(*tests[index].s2).position.y = 90.0
// 	(*tests[index].s1).mass = 20.0
// 	(*tests[index].s2).mass = 30.0
// 	tests[index].answer.x = 0.0000000000128
// 	tests[index].answer.y = 0.00000000000961


// 	// range over the tests and check if the function produces the correct output
// 	for i := range tests {
// 		output := ComputeSingleForce(tests[i].s1, tests[i].s2)
// 		output.x = roundFloat(output.x, 13)
// 		output.y = roundFloat(output.y, 14)

// 		fmt.Println("ComputeSeparationForce")
// 		fmt.Println("Test ", i, ":")
// 		// if function output and correct answer mismatch
// 		if output.x != tests[i].answer.x || output.y != tests[i].answer.y {
// 			fmt.Println("Incorrect! The function's output is ", output.x, " and ", output.y, ". The correct answer is ", tests[i].answer.x, " and ", tests[i].answer.y, ".")
// 		} else { // if function output and correct answer match
// 			fmt.Println("Correct! Answer: (", output.x, ", ", output.y, ")")
// 		}
// 	}
// 	fmt.Println("")
// }