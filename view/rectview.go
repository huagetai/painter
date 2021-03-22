package view

type RectangleView struct {
	X              float64
	Y              float64
	Width          float64
	Height         float64
	BorderRadis    float64
	BackgroudColor string
}

func (v *RectangleView) Paint(ctx *PosterContext) {
	dc := ctx.Ctx
	dc.DrawRoundedRectangle(v.X, v.Y, v.Width, v.Height, v.BorderRadis)
	dc.SetHexColor(v.BackgroudColor)
	dc.Fill()
}
