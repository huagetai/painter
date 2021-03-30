package main

import (
	"gopkg.in/fogleman/gg.v1"
)

func main() {
	dc := gg.NewContext(500, 500)
	dc.DrawCircle(50, 50, 50)
	dc.SetHexColor("#00aaaa")
	dc.Stroke()
	dc.SavePNG("out.png")
}

func DrawRoundedRectangleWithColor(dc *gg.Context, x, y, w, h, r float64, c string) {
	x0, x1, x2, x3 := x, x+r, x+w-r, x+w
	y0, y1, y2, y3 := y, y+r, y+h-r, y+h
	dc.NewSubPath()
	dc.MoveTo(x1, y0)
	dc.LineTo(x2, y0)
	dc.DrawArc(x2, y1, r, gg.Radians(270), gg.Radians(360))
	dc.LineTo(x3, y2)
	dc.DrawArc(x2, y2, r, gg.Radians(0), gg.Radians(90))
	dc.LineTo(x1, y3)
	dc.DrawArc(x1, y2, r, gg.Radians(90), gg.Radians(180))
	dc.LineTo(x0, y1)
	dc.DrawArc(x1, y1, r, gg.Radians(180), gg.Radians(270))
	dc.SetHexColor(c)
	dc.Stroke()
	dc.ClosePath()
}
