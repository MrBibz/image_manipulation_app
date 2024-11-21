package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// Fit scales a widget to fit and clip to the constraints.
type Fit uint8

const (
	// Unscaled does not alter the scale of a widget.
	Unscaled Fit = iota
	// Contain scales widget as large as possible without cropping
	// and it preserves aspect-ratio.
	Contain
	// Cover scales the widget to cover the constraint area and
	// preserves aspect-ratio.
	Cover
	// ScaleDown scales the widget smaller without cropping,
	// when it exceeds the constraint area.
	// It preserves aspect-ratio.
	ScaleDown
	// Fill stretches the widget to the constraints and does not
	// preserve aspect-ratio.
	Fill
)

// scale computes the new dimensions and transformation required to fit dims to cs, given the position.
func (fit Fit) scale(cs layout.Constraints, pos layout.Direction, dims layout.Dimensions) (layout.Dimensions, f32.Affine2D) {
	widgetSize := dims.Size

	// If no scaling is needed or dimensions are zero, return the constrained size and offset.
	if fit == Unscaled || dims.Size.X == 0 || dims.Size.Y == 0 {
		dims.Size = cs.Constrain(dims.Size)
		offset := pos.Position(widgetSize, dims.Size)
		dims.Baseline += offset.Y
		return dims, f32.Affine2D{}.Offset(layout.FPt(offset))
	}

	// Calculate the scaling factors for both dimensions.
	scale := f32.Point{
		X: float32(cs.Max.X) / float32(dims.Size.X),
		Y: float32(cs.Max.Y) / float32(dims.Size.Y),
	}

	// Adjust the scaling factors based on the fit type.
	switch fit {
	case Contain:
		if scale.Y < scale.X {
			scale.X = scale.Y
		} else {
			scale.Y = scale.X
		}
	case Cover:
		if scale.Y > scale.X {
			scale.X = scale.Y
		} else {
			scale.Y = scale.X
		}
	case ScaleDown:
		if scale.Y < scale.X {
			scale.X = scale.Y
		} else {
			scale.Y = scale.X
		}
		// If scaling up is not needed, return the constrained size and offset.
		if scale.X >= 1 {
			dims.Size = cs.Constrain(dims.Size)
			offset := pos.Position(widgetSize, dims.Size)
			dims.Baseline += offset.Y
			return dims, f32.Affine2D{}.Offset(layout.FPt(offset))
		}
	case Fill:
		// No additional adjustments needed for Fill.
	}

	// Calculate the scaled size and transformation.
	var scaledSize image.Point
	scaledSize.X = int(float32(widgetSize.X) * scale.X)
	scaledSize.Y = int(float32(widgetSize.Y) * scale.Y)
	dims.Size = cs.Constrain(scaledSize)
	dims.Baseline = int(float32(dims.Baseline) * scale.Y)

	offset := pos.Position(scaledSize, dims.Size)
	trans := f32.Affine2D{}.
		Scale(f32.Point{}, scale).
		Offset(layout.FPt(offset))

	dims.Baseline += offset.Y

	return dims, trans
}

// Image represents an image widget with scaling and positioning options.
type Image struct {
	Src      paint.ImageOp
	Fit      Fit
	Position layout.Direction
	Scale    float32
}

// Layout lays out the image widget within the given context.
func (im Image) Layout(gtx layout.Context) layout.Dimensions {
	scale := im.Scale
	if scale == 0 {
		scale = 1
	}

	size := im.Src.Size()
	wf, hf := float32(size.X), float32(size.Y)
	w, h := gtx.Dp(unit.Dp(wf*scale)), gtx.Dp(unit.Dp(hf*scale))

	dims, trans := im.Fit.scale(gtx.Constraints, im.Position, layout.Dimensions{Size: image.Pt(w, h)})
	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()

	pixelScale := scale * gtx.Metric.PxPerDp
	trans = trans.Mul(f32.Affine2D{}.Scale(f32.Point{}, f32.Pt(pixelScale, pixelScale)))
	defer op.Affine(trans).Push(gtx.Ops).Pop()

	im.Src.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return dims
}

func main() {
	go func() {
		window := new(app.Window)
		if err := imageDisplayLoop(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

// imageDisplayLoop handles the main event loop for displaying the image.
func imageDisplayLoop(window *app.Window) error {
	file, err := os.Open("../images/bread.jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	imgOp := paint.NewImageOp(img)

	for {
		event := window.Event()
		switch typ := event.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&op.Ops{}, typ)

			imageWidget := Image{
				Src:      imgOp,
				Fit:      Contain,
				Position: layout.Center,
				Scale:    1,
			}

			imageWidget.Layout(gtx)
			typ.Frame(gtx.Ops)
		case app.DestroyEvent:
			return typ.Err
		}
	}
}
