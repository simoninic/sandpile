package main

import (
	"canvas"
	"image"
	//"fmt"
)

func (board GameBoard) DrawToCanvas(canvasWidth int) image.Image {
	if board == nil {
		panic("Cannot draw a nil board.")
	}

	width := len(board)


	// set a new square canvas
	c := canvas.CreateNewCanvas(1000, 1000)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0)) //black
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	rwidth := int(float64(canvasWidth) / float64(width))

	// // range over all the bodies and draw them.
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 3 {
				c.SetFillColor(canvas.MakeColor(255, 255, 255))
			} else if board[i][j] == 2 {
				//c.SetFillColor(canvas.MakeColor(0, 255, 0))
				c.SetFillColor(canvas.MakeColor(170, 170, 170))
			} else if board[i][j] == 1 {
				//c.SetFillColor(canvas.MakeColor(0, 0, 255))
				c.SetFillColor(canvas.MakeColor(85, 85, 85))
			} else {
				//c.SetFillColor(canvas.MakeColor(255, 0, 0))
				c.SetFillColor(canvas.MakeColor(0, 0, 0))
			}

			c.ClearRect(i * rwidth, j * rwidth, i * rwidth + rwidth, j * rwidth + rwidth)
			//fmt.Println(i, "    ", j,  "     ", rwidth)
			//c.Close()

			//c.Rect(rx, ry, rwidth, rwidth)
			//c.Circle(rx, ry, rwidth)

			c.Fill()
		}
	}

	// c.SetFillColor(canvas.MakeColor(0, 220, 180))
	// c.Circle(1, 1, 0.5)
	// c.Fill()

	c.SetFillColor(canvas.MakeColor(0, 220, 12))
	c.ClearRect(1, 1, 2, 2)
	c.Fill()

	// c.SetFillColor(canvas.MakeColor(255, 255, 255))
	// c.ClearRect(10, 10, 200, 200)
	// c.Fill()

	// we want draw the image image!
	return c.GetImage()
}

// Draws an empty rectangle
// Fill the given circle with the fill color
// Stroke() each time to avoid connected circles
// func (c *Canvas) Rect(x, y, w, h float64) {
// 	c.gc.LineTo(x, x + w)
// 	c.gc.LineTo(x + w, y + h)
// 	c.gc.LineTo(x, y + h)
// 	c.gc.LineTo(x, y)
// 	c.gc.Close()
// }