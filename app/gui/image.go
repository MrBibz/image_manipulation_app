package gui

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
)

type Image struct {
	Src      paint.ImageOp
	Fit      Fit
	Position layout.Direction
	Scale    float32
}

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
