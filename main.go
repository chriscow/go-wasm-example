package main

import (
	"fmt"
	"syscall/js"

	"github.com/chriscow/go-wasm-example/js/canvas"
	"github.com/chriscow/go-wasm-example/vector2d"
)

func renderPolygon(ctx *canvas.Context2D, s *vector2d.Polygon) {
	ctx.Save()

	ctx.BeginPath()
	first := true
	var firstPoint vector2d.Vector

	for _, v := range s.Matrix {
		if first {
			ctx.MoveTo(v[0], v[1])
			first = false
			firstPoint = v
		} else {
			ctx.LineTo(v[0], v[1])
		}
	}

	ctx.LineTo(firstPoint[0], firstPoint[1])
	ctx.Stroke()
	ctx.Restore()
}

func initScreen(window js.Value, w, h int) *canvas.Context2D {
	document := window.Get("document")
	canvasE := document.Call("createElement", "canvas")

	canvasE.Set("id", "canvas")
	canvasE.Set("width", w)
	canvasE.Set("height", h)
	document.Get("body").Call("appendChild", canvasE)
	document.Get("body").Set("style", "margin: 0px; overflow: hidden")

	return canvas.NewContext2D(canvasE)
}

func main() {
	window := js.Global()
	ctx := initScreen(window, 1280, 720)

	// sets the type of compositing operation to apply when drawing new shapes
	// destination-over:
	// 	New shapes are drawn behind the existing canvas content.
	ctx.SetGlobalCompositeOperation("destination-over")

	poly := vector2d.NewPolygon(
		0, 2,
		-1.5, -2,
		-1, -1,
		1, -1,
		1.5, -2,
	)

	proj := poly.Clone().Translate(1280/2, 720/2).Scale(10, 10)
	renderPolygon(ctx, proj)

	ctx.Save()
	ctx.SetTextAlign("center")
	ctx.SetFont("bold 20px Courier New")
	ctx.SetStrokeStyle("rgba(0, 0, 0, 1)")
	ctx.FillText(fmt.Sprintf("Score:"), 640, 40, 1280)
	ctx.Restore()

	// Never return
	done := make(chan struct{}, 0)
	<-done
}
