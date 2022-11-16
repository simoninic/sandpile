package main

import (
	"canvas"
	"image"
	//"fmt"
)

// DrawToCanvas: uses canvas to draw an image showing the resulting pattern from the coin toppling
// input: the width of the canvas (int)
// output: the image representing the resulting pattern from coin toppling
func (board GameBoard) DrawToCanvas(canvasWidth int) image.Image {
	if board == nil {
		panic("Cannot draw a nil board.")
	}

	width := len(board)

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0)) //black
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// the width of each cell of the board
	rwidth := int(float64(canvasWidth) / float64(width))

	// range over each cell in the board and color it accordingly
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 3 { // color cell white if it has 3 coins
				c.SetFillColor(canvas.MakeColor(255, 255, 255))
			} else if board[i][j] == 2 { // color cell light gray if it has 2 coins
				c.SetFillColor(canvas.MakeColor(170, 170, 170))
			} else if board[i][j] == 1 { // color cell dark gray if it has 1 coin
				c.SetFillColor(canvas.MakeColor(85, 85, 85))
			} else { // color cell black if it has no coins
				c.SetFillColor(canvas.MakeColor(0, 0, 0))
			}

			// draw a rectangle to represent the cell of the board
			c.ClearRect(i * rwidth, j * rwidth, i * rwidth + rwidth, j * rwidth + rwidth)

			// color the rectangle according to the conditions above
			c.Fill()
		}
	}

	// draw the image!
	return c.GetImage()
}