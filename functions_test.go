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

// // roundFloat
// // inputs: the value (float64) to be rounded, and precision (uint)
// // output: the rounded number as a float64
// func roundFloat(val float64, precision uint) float64 {
// 	ratio := math.Pow(10, float64(precision))
// 	return math.Round(val*ratio) / ratio
// }

func TestUpdateOnBoardCells(t *testing.T) {

	type test struct {
		board                        GameBoard
		numRows, numCols, row, col   int
		answer                       int
	}
	
	inputDirectory := "tests/UpdateOnBoardCells/input/"
	outputDirectory := "tests/UpdateOnBoardCells/output/"

	inputFiles := ReadFilesFromDirectory(inputDirectory)
	outputFiles := ReadFilesFromDirectory(outputDirectory)

	fmt.Println("A")

	//assert that files are non-empty and have the same length
	AssertEqualAndNonzero(len(inputFiles), len(outputFiles))

	//we now will need to create our array of tests
	tests := make([]test, len(inputFiles))

	var board GameBoard
	var numRows int
	var numCols int
	var row int
	var col int


	//first, range through the input and output files and set the test values
	for i := range inputFiles {
		
		board, numRows, numCols, row, col = ReadBoardFourIntsFromFile(inputDirectory, inputFiles[i])
		tests[i].board = board
		tests[i].numRows = numRows
		tests[i].numCols = numCols
		tests[i].row = row
		tests[i].col = col

		tests[i].answer = ReadIntegerFromFile(outputDirectory, outputFiles[i])
	}

	//are the tests correct?
	for _, test := range tests {
		fmt.Println("Thomas")
		fmt.Println(test)
		fmt.Println("start test board")
		PrintBoard(test.board)
		fmt.Println("end test board")
		outcome := UpdateOnBoardCells(test.board, test.numRows, test.numCols, test.row, test.col)
		fmt.Println("Outcome: ", outcome)
		fmt.Println("test answer: ", test.answer)
			
		if outcome == test.answer {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}
}

func ReadBoardFourIntsFromFile(directory string, inputFile os.FileInfo) (GameBoard, int, int, int, int) {
	fileName := inputFile.Name() //grab file name

	//now, read in the input file
	fileContents, err := ioutil.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}

	inputLines := strings.Split(strings.TrimSpace(strings.Replace(string(fileContents), "\r\n", "\n", -1)), "\n")

	index := 0

	numRows, err2 := strconv.Atoi(string(inputLines[0]))
	if err2 != nil {
		panic(err2)
	}
	index++
	fmt.Println("numRows: ", numRows)
	numCols, err3 := strconv.Atoi(string(inputLines[index][0]))
	if err3 != nil {
		panic(err3)
	}
	index++
	fmt.Println("numCols: ", numRows)
	row, err4 := strconv.Atoi(string(inputLines[index][0]))
	if err4 != nil {
		panic(err4)
	}
	index++
	fmt.Println("row: ", row)
	col, err5 := strconv.Atoi(string(inputLines[index][0]))
	if err5 != nil {
		panic(err5)
	}
	fmt.Println("col: ", col)
	index++


	// rawLines := string(fileContents)
	// index := 0
	// for i, line := range rawLines {
	// 	fmt.Println("line: ", string(line))
	// 	if string(line) == "/" {
	// 		index = i
	// 	}
	// }
	// rawLines = rawLines[0:index]

	//first, read lines and split along blank space


	// index := 0
	// for i, inputLine := range inputLines {
	// 	//currentRow := strings.Split(inputLine, " ")
	// 	if inputLine == "/" {
	// 		index = i
	// 	}
	// }

	//make the map that will store our frequency map
	var frequencyMap [][]int

	frequencyMap = make([]([]int), numRows)
	for r := range frequencyMap {
		frequencyMap[r] = make([]int, numCols)
	}

	//each line of the file corresponds to a single line of the frequency map
	for i := index + 1; i < len(inputLines); i++ {

		//read out the current line
		currentRow := strings.Split(inputLines[i], " ")
		fmt.Println("currentRow: ", currentRow)
		fmt.Println("len current row: ", len(currentRow))

		//currentLine has two strings corresponding to the key and value

		//frequencyMap = append(frequencyMap, make([]int, numCols))
		//maxIndex := len(frequencyMap) - 1

		//frequencyMap[maxIndex] = make([]int, 0)

		for j, inputCell := range currentRow {

			currentValue, err := strconv.Atoi(inputCell)
			fmt.Println(currentValue)
			if err != nil {
				panic(err)
			}
			frequencyMap[i - index - 1][j] = currentValue
			//frequencyMap[maxIndex] = append(frequencyMap[maxIndex], currentValue)
			fmt.Println("row in progress: ", frequencyMap[i - index - 1])
		}

		//if we make it here, everything is OK, so append to the input map
		
	}
	fmt.Println("index: ", index)

	return frequencyMap, numRows, numCols, row, col
	//return frequencyMap, 0, 0, 0, 0
}

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