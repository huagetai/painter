package painter

import (
	"gopkg.in/fogleman/gg.v1"
)

type PosterContext struct {
	Width  int
	Height int
	Ctx    *gg.Context
}

type View interface {
	Paint(ctx *PosterContext)
}

type LineView struct {
	X         float64
	Y         float64
	X2        float64
	Y2        float64
	LineWidth float64
	LineColor string
}

func (v *LineView) Paint(pc *PosterContext) {
	dc := pc.Ctx
	dc.DrawLine(v.X, v.Y, v.X2, v.Y2)
	dc.SetLineWidth(v.LineWidth)
	dc.SetHexColor(v.LineColor)
	dc.Stroke()
}

type RectangleView struct {
	X              float64
	Y              float64
	Width          float64
	Height         float64
	BackgroudColor string
	BorderStyle    string
	BorderRadius   float64
	BorderWidth    float64
	BorderColor    string
}

func (v *RectangleView) Paint(ctx *PosterContext) {
	dc := ctx.Ctx
	dc.NewSubPath()
	if v.BorderStyle != "" && v.BorderStyle != "none" {
		dc.SetLineWidth(v.BorderWidth)
		dc.DrawRoundedRectangle(v.X, v.Y, v.Width, v.Height, v.BorderRadius)
		dc.SetHexColor(v.BorderColor)
		dc.Stroke()
	}
	if v.BackgroudColor != "" {
		dc.DrawRoundedRectangle(v.X, v.Y, v.Width, v.Height, v.BorderRadius)
		dc.SetHexColor(v.BackgroudColor)
		dc.Fill()
	}
	dc.ClosePath()
}

type CircleView struct {
	X              float64
	Y              float64
	Radius         float64
	BackgroudColor string
	BorderStyle    string
	BorderWidth    float64
	BorderColor    string
}

func (v *CircleView) Paint(pc *PosterContext) {
	dc := pc.Ctx
	dc.NewSubPath()
	if v.BorderStyle != "" && v.BorderStyle != "none" {
		dc.SetLineWidth(v.BorderWidth)
		dc.DrawCircle(v.X, v.Y, v.Radius)
		dc.SetHexColor(v.BorderColor)
		dc.Stroke()
	}
	if v.BackgroudColor != "" {
		dc.DrawCircle(v.X, v.Y, v.Radius)
		dc.SetHexColor(v.BackgroudColor)
		dc.Fill()
	}
	dc.ClosePath()
}
