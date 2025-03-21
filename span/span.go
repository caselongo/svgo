package main

import (
	svg "github.com/caselongo/svgo"
	"os"
)

func main() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Gstyle("text-anchor:middle;font-family:sans;fill:white")
	canvas.Textspan(width/2, height/2, "Hello ", "font-size:30px")
	canvas.Span(0, 0, "SVG", "font-family:serif;font-size:50px;fill:yellow")
	canvas.TextEnd()
	canvas.Gend()
	canvas.End()
}
