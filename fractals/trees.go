package fractals

import "math"

type Line struct {
	X0, Y0, X1, Y1, Width int
	angleRad, length      float64
}

type FractalTrees struct {
	List         []FractalTree
	CanvasHeight float64
}

func (ts *FractalTrees) Plant(x, length, width, angleFirst, angleSecond int, color string) {
	t := FractalTree{}
	t.x = x
	t.y = int(ts.CanvasHeight)
	t.length = length
	t.angleFirst = angleFirst
	t.angleSecond = angleSecond
	t.Color = color

	t.Lines = []Line{{X0: t.x, Y0: t.y, X1: t.x, Y1: t.y - length, Width: width, length: float64(t.length), angleRad: math.Pi / 2}}
	ts.List = append(ts.List, t)
}

type FractalTree struct {
	Lines                                        []Line
	x, y, length, angleFirst, angleSecond, level int
	Color                                        string
	StopGrow                                     bool
}

func (t *FractalTree) AddBrunches() {
	if t.StopGrow {
		return
	}

	angleRadOffsetFirst := float64(t.angleFirst) * math.Pi / float64(180)
	angleRadOffsetSecond := float64(t.angleSecond) * math.Pi / float64(180)

	initialLinesCount := len(t.Lines)

	branchLength := float64(t.Lines[initialLinesCount-1].length) * 0.8

	if branchLength < 10 {
		t.StopGrow = true
		return
	}

	for n := (len(t.Lines)+1)/2 - 1; n < initialLinesCount; n++ {

		for i := 0; i < 2; i++ {
			line := t.Lines[n]

			mirror := 1.0
			if t.level%2 == 1 {
				mirror = -1.0
			}

			angleRad := 0.0
			if i == 0 {
				angleRad = line.angleRad + angleRadOffsetFirst*mirror
			} else {
				angleRad = line.angleRad - angleRadOffsetSecond*mirror
			}

			x := line.X1
			y := line.Y1

			xOffset := int(math.Cos(angleRad) * float64(branchLength))
			yOffset := int(math.Sin(angleRad) * float64(branchLength))

			nextWidth := line.Width - 1
			if nextWidth < 1 {
				nextWidth = 1
			}
			t.Lines = append(t.Lines, Line{X0: x, Y0: y, X1: x + xOffset, Y1: y - yOffset, Width: nextWidth, length: branchLength, angleRad: angleRad})
		}
	}

	t.level++
}
