package main

// type board
type GameBoard [][]int

// Star is analogous to the "Body" object from the jupiter simulations.
type SubBoard struct {
	startRow                         int
	endRow                           int
	topFalloff                       []*int
	bottomFalloff                    []*int
}
