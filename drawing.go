package main

import (
	"canvas"
	"image"
	"image/png"
	"os"
)

// DrawToCanvas: uses canvas to draw an image showing the resulting pattern from the coin toppling
// input: the width of the canvas (int)
// output: the image representing the resulting pattern from coin toppling
func (board GameBoard) DrawToCanvas() image.Image {
	if board == nil {
		panic("Cannot draw a nil board.")
	}

	width := len(board)
	canvasWidth := width * 3

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
				c.SetFillColor(canvas.MakeColor(90, 130, 90))
			} else if board[i][j] == 2 { // color cell light gray if it has 2 coins
				c.SetFillColor(canvas.MakeColor(130, 180, 130))
			} else if board[i][j] == 1 { // color cell dark gray if it has 1 coin
				c.SetFillColor(canvas.MakeColor(170, 225, 170))
			} else { // color cell black if it has no coins
				c.SetFillColor(canvas.MakeColor(204, 255, 255))
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

// MakeSerialImage: draws the PNG image of the final board state after toppling
// input: the image to draw, the file name of the new PNG file containing the image
func MakeImage(img image.Image, fileName string) {
	if _, err := os.Stat(fileName); err == nil { // if file exists, delete it
		os.Remove(fileName)
	}
	// create new image file with new png image of final board state
	out, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	png.Encode(out, img)
	out.Close()	
}