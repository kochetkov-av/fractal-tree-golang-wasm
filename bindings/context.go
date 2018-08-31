package bindings

import (
	"syscall/js"
)

type CanvasContext struct {
	Value js.Value
}

func (ctx CanvasContext) BeginPath() {
	ctx.Value.Call("beginPath")
}

func (ctx CanvasContext) MoveTo(x, y int) {
	ctx.Value.Call("moveTo", x, y)
}

func (ctx CanvasContext) LineTo(x, y int) {
	ctx.Value.Call("lineTo", x, y)
}

func (ctx CanvasContext) Stroke() {
	ctx.Value.Call("stroke")
}

func (ctx CanvasContext) Scale(x, y float64) {
	ctx.Value.Call("scale", x, y)
}

func (ctx CanvasContext) StrokeStyle(color string) {
	ctx.Value.Set("strokeStyle", color)
}

func (ctx CanvasContext) LineWidth(width int) {
	ctx.Value.Set("lineWidth", width)
}

func (ctx CanvasContext) DrawImage(image js.Value, x, y int) {
	ctx.Value.Call("drawImage", image, x, y)
}

func (ctx CanvasContext) ClearRect(x, y, width, height int) {
	ctx.Value.Call("clearRect", x, y, width, height)
}

func (ctx CanvasContext) FillStyle(color string) {
	ctx.Value.Set("fillStyle ", color)
}

func (ctx CanvasContext) Font(font string) {
	ctx.Value.Set("font", font)
}

func (ctx CanvasContext) FillText(text string, x, y int) {
	ctx.Value.Call("fillText", text, x, y)
}
