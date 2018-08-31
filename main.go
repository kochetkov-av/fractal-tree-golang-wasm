// +build js,wasm

package main

import (
	"fmt"
	"fractal_trees/bindings"
	"fractal_trees/fractals"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

const (
	TreesMax     = 3
	CanvasWidth  = 1920.0
	CanvasHeight = 1080.0
	ShowFps      = false
)

var (
	window          = js.Global().Get("window")
	document        = js.Global().Get("document")
	canvas          = document.Call("getElementById", "canvas")
	mainCtx         = bindings.CanvasContext{Value: canvas.Call("getContext", "2d")}
	offscreenCanvas = document.Call("createElement", "canvas")

	trees         = fractals.FractalTrees{CanvasHeight: CanvasHeight}
	needRedraw    = false
	lastFrameTime = time.Now().UnixNano()
)

func main() {
	c := make(chan int)

	rand.Seed(time.Now().UnixNano())

	window.Call("requestAnimationFrame", js.NewCallback(render))

	go grow()

	<-c
}

type WeightedRandomSelectItem struct {
	itemType       string
	itemId, weight int
}

type WeightedRandomSelect struct {
	List []WeightedRandomSelectItem
}

func (wrs *WeightedRandomSelect) Add(item WeightedRandomSelectItem) {
	wrs.List = append(wrs.List, item)
}

func (wrs *WeightedRandomSelect) Random() WeightedRandomSelectItem {
	sumWeight := 0
	for _, item := range wrs.List {
		sumWeight += item.weight
	}

	if sumWeight == 0 {
		return WeightedRandomSelectItem{}
	}

	rnd := rand.Intn(sumWeight)
	sumWeight = 0
	for _, item := range wrs.List {
		sumWeight += item.weight
		if sumWeight >= rnd {
			return item
		}
	}
	return WeightedRandomSelectItem{}
}

func grow() {
	for {
		wrs := WeightedRandomSelect{}

		for i, tree := range trees.List {
			if !tree.StopGrow {
				wrs.Add(WeightedRandomSelectItem{itemType: "branch", itemId: i, weight: 40})
			}
		}

		if len(trees.List) < TreesMax {
			wrs.Add(WeightedRandomSelectItem{itemType: "tree", weight: 10})
		}

		random := wrs.Random()

		if random.itemType == "branch" {
			trees.List[random.itemId].AddBrunches()
		}

		if random.itemType == "tree" {
			x := rand.Intn(1720) + 100
			length := 100 + rand.Intn(150)
			width := int(15.0 * (float64(length) / 250.0))

			trees.Plant(x, length, width, 5+rand.Intn(50), 5+rand.Intn(50), randomColorString())
		}

		if random.itemType == "" {
			trees.List = trees.List[1:]
		}

		needRedraw = true
		time.Sleep(300 * time.Millisecond)
	}
}

func randomColorString() string {
	r := rand.Intn(175)
	g := rand.Intn(175)
	b := rand.Intn(175)
	colorString := fmt.Sprintf("#%02X%02X%02X", r, g, b)
	return colorString
}

func render(_ []js.Value) {
	if needRedraw {
		draw()
		needRedraw = false
	}

	frameTime := time.Now().UnixNano()
	fps := int64(time.Second) / (frameTime - lastFrameTime)
	lastFrameTime = frameTime

	if ShowFps {
		mainCtx.ClearRect(0, 0, 80, 20)
		mainCtx.FillStyle("Black")
		mainCtx.Font("normal 16pt Arial")
		mainCtx.FillText("fps "+strconv.Itoa(int(fps)), 5, 15)
	}

	window.Call("requestAnimationFrame", js.NewCallback(render))
}

func draw() {
	width := window.Get("innerWidth").Int()
	height := window.Get("innerHeight").Int()

	canvasW := canvas.Get("width").Int()
	canvasH := canvas.Get("height").Int()

	if (canvasW != width) || (canvasH != height) {
		// resize canvas to fullscreen
		canvas.Set("width", width)
		canvas.Set("height", height)
	}

	offscreenCanvas.Set("width", width)
	offscreenCanvas.Set("height", height)

	ctx := bindings.CanvasContext{Value: offscreenCanvas.Call("getContext", "2d")}

	scaleX := float64(width) / CanvasWidth
	scaleY := float64(height) / CanvasHeight

	ctx.Scale(scaleX, scaleY)

	for _, tree := range trees.List {
		ctx.StrokeStyle(tree.Color)
		for _, line := range tree.Lines {
			ctx.BeginPath()
			ctx.MoveTo(line.X0, line.Y0)
			ctx.LineTo(line.X1, line.Y1)
			ctx.LineWidth(line.Width)
			ctx.Stroke()
		}
	}

	mainCtx.ClearRect(0, 0, width, height)
	mainCtx.DrawImage(offscreenCanvas, 0, 0)
}
