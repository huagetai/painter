package view

import (
	"math"
	"strconv"
	"strings"
)

type TextView struct {
	X          float64
	Y          float64
	Width      float64
	LineHeight float64
	LineClamp  string
	Text       string
	FontSize   float64
	FontWeight string
	Color      string
	Align      string
}

func (v *TextView) Paint(ctx *PosterContext) {
	dc := ctx.Ctx
	dc.SetHexColor(v.Color)
	v.loadFontFace(ctx)
	x := v.X
	y := v.Y
	width := v.Width
	lineHeight := math.Max(v.LineHeight, dc.FontHeight()*1.25)
	lines := v.lineWrap(ctx)
	lc, err := strconv.Atoi(v.LineClamp)
	if err != nil {
		ax := 0.0
		ay := 1.0
		switch v.Align {
		case "left":
			ax = 0
		case "center":
			ax = 0.5
			x += width / 2
		case "right":
			ax = 1
			x += width
		}
		y = y + (lineHeight-dc.FontHeight())/2
		dc.DrawStringAnchored(lines[0], x, y, ax, ay)
	} else {
		for index, line := range lines {
			if index >= lc {
				break
			}
			ax := 0.0
			ay := 1.0
			switch v.Align {
			case "left":
				ax = 0
			case "center":
				ax = 0.5
				x += width / 2
			case "right":
				ax = 1
				x += width
			}
			y += lineHeight*float64(index) + (lineHeight-dc.FontHeight())/2
			dc.DrawStringAnchored(line, x, y, ax, ay)
		}
	}
}

func (v *TextView) loadFontFace(ctx *PosterContext) {
	dc := ctx.Ctx
	switch v.FontWeight {
	case "heavy":
		if err := dc.LoadFontFace("assets/fonts/Alibaba-PuHuiTi-Heavy.ttf", v.FontSize); err != nil {
			panic(err)
		}
	case "bold":
		if err := dc.LoadFontFace("assets/fonts/Alibaba-PuHuiTi-Bold.ttf", v.FontSize); err != nil {
			panic(err)
		}
	case "bolder":
		if err := dc.LoadFontFace("assets/fonts/Alibaba-PuHuiTi-Medium.ttf", v.FontSize); err != nil {
			panic(err)
		}
	case "normal":
		if err := dc.LoadFontFace("assets/fonts/Alibaba-PuHuiTi-Regular.ttf", v.FontSize); err != nil {
			panic(err)
		}
	case "lighter":
		if err := dc.LoadFontFace("assets/fonts/Alibaba-PuHuiTi-Light.ttf", v.FontSize); err != nil {
			panic(err)
		}
	default:
		if err := dc.LoadFontFace("assets/fonts/Alibaba-PuHuiTi-Regular.ttf", v.FontSize); err != nil {
			panic(err)
		}
	}
}

func (v *TextView) lineWrap(ctx *PosterContext) []string {
	var result []string
	dc := ctx.Ctx
	width := v.Width
	s := v.Text
	s = strings.ReplaceAll(s, "\n", "")
	words := split(s)
	if len(words)%2 == 1 {
		words = append(words, "")
	}
	x := ""
	for i := 0; i < len(words); i += 2 {
		w, _ := dc.MeasureString(x + words[i])
		if w > width {
			if x == "" {
				result = append(result, words[i])
				x = ""
				continue
			} else {
				result = append(result, x)
				x = ""
			}
		}
		x += words[i] + words[i+1]
	}
	if x != "" {
		result = append(result, x)
	}
	return result
}

func split(x string) []string {
	var result []string
	pi := 0
	for i, _ := range x {
		if i > 0 {
			result = append(result, x[pi:i])
			pi = i
		}
	}
	result = append(result, x[pi:])
	return result
}
