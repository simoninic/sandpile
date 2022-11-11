package main

// GameBoard is a nested slice of ints
type GameBoard [][]int

// SubBoard contains info that is important to acquire from the channels
// startRow: the first row index of the sub-board
// endRow: the last row index + 1 of the sub-board
// topFalloff: the slice of ints that represents coins falling off the top of the sub-board
// bottomFalloff: the slice of ints that represent coins falling off the bottom of the sub-board
type SubBoard struct {
	startRow                         int
	endRow                           int
	topFalloff                       []*int
	bottomFalloff                    []*int
}
