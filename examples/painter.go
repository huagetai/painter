package main

import (
	"bufio"
	"github.com/huagetai/painter"
	"github.com/huagetai/painter/examples/palette"
	"os"
)

func main() {
	drawBase()
	activity()
}

func drawBase() {
	views := []painter.View{
		&painter.LineView{
			X:         0,
			Y:         0,
			X2:        300,
			Y2:        300,
			LineWidth: 1,
			LineColor: "#000000",
		},
		&painter.RectangleView{
			X:              200,
			Y:              100,
			Width:          200,
			Height:         50,
			BackgroudColor: "#888888",
			BorderStyle:    "solid",
			BorderRadius:   8,
			BorderColor:    "#000000",
		},
		&painter.CircleView{
			X:              300,
			Y:              300,
			Radius:         100,
			BackgroudColor: "#888888",
			BorderStyle:    "solid",
			BorderColor:    "#000000",
		},
	}
	p := painter.Palette{
		BackgroundColor: "#ffffff",
		Width:           600,
		Height:          600,
		Views:           views,
	}
	f, _ := os.Create("./image/example-rect.png")
	w := bufio.NewWriter(f)
	defer w.Flush()
	p.Paint(w)
}

func activity() {
	a := palette.Activity{
		Title: "【预售3月24日排单发】1号莲雾，连续十年有机认证认证🏆与那个让你惊艳的牛奶莲雾平分秋色 🏆基地自主研发 🔥  独有品种🔥 不一样的口感和味觉体验💥 必须要尝试一下！",
		Images: []string{
			"./image/test-1.jpeg",
			"./image/test-2.jpeg",
			"./image/test-3.jpeg",
			"./image/test-4.jpeg",
			"./image/test-5.jpeg",
			"./image/test-6.jpeg",
		},
		BtnName:   "我要团购",
		AvatarUrl: "./image/logo.png",
	}
	f, _ := os.Create("./image/activity-out.png")
	w := bufio.NewWriter(f)
	a.Gen(w)
	w.Flush()
}
