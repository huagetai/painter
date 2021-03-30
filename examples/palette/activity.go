package palette

import (
	"github.com/huagetai/painter"
	"io"
)

type Activity struct {
	AvatarUrl string
	Title     string
	Images    []string
	BtnName   string
}

func (a *Activity) Gen(writer io.Writer) {
	var views []painter.View
	iv := &painter.ImageView{
		X:            Padding,
		Y:            5,
		Width:        45,
		Height:       45,
		URI:          a.AvatarUrl,
		Mode:         "smart",
		BorderRadius: 22.5,
	}
	views = append(views, iv)
	title := &painter.TextView{
		X:          Padding + 8 + 45,
		Y:          5,
		Width:      float64(CanvasWidth - Padding - 8 - 45 - Padding),
		LineHeight: 45,
		LineClamp:  "1",
		Text:       a.Title,
		FontWeight: "bolder",
		FontSize:   20,
		Color:      "#000000",
		Align:      "left",
	}
	views = append(views, title)
	if len(a.Images) >= 4 {
		x := Padding
		y := 5 + 45 + 16
		w := (CanvasWidth - (Padding * 2) - (8 * 2)) / 3
		h := (CanvasHeight - Padding - 90 - y - 16 - 8) / 2
		for j := 0; j < 2; j++ {
			for i := 0; i < 3; i++ {
				k := j*3 + i
				if k >= len(a.Images) {
					break
				}
				iv := &painter.ImageView{
					X:            x + i*w + i*8,
					Y:            y + j*h + j*8,
					Width:        w,
					Height:       h,
					URI:          a.Images[k],
					Mode:         "smart",
					BorderRadius: 4,
				}
				views = append(views, iv)
			}
		}
	} else if len(a.Images) > 0 {
		iv := &painter.ImageView{
			X:            Padding,
			Y:            5 + 45 + 16,
			Width:        CanvasWidth - (Padding * 2),
			Height:       CanvasHeight - Padding - 90 - 66 - 16,
			URI:          a.Images[0],
			Mode:         "smart",
			BorderRadius: 4,
		}
		views = append(views, iv)
	}

	rv := &painter.RectangleView{
		X:              Padding,
		Y:              CanvasHeight - Padding - 90,
		Width:          float64(CanvasWidth - Padding*2),
		Height:         90,
		BorderRadius:   48,
		BackgroudColor: "#06ae56",
	}
	views = append(views, rv)
	tv := &painter.TextView{
		X:          Padding,
		Y:          CanvasHeight - Padding - 90,
		Width:      float64(CanvasWidth - Padding*2),
		LineHeight: 90,
		LineClamp:  "none",
		Text:       a.BtnName,
		FontSize:   40,
		Color:      "#FFFFFF",
		Align:      "center",
	}
	views = append(views, tv)
	p := painter.Palette{
		BackgroundColor: "#FFFFFF",
		Width:           CanvasWidth,
		Height:          CanvasHeight,
		Views:           views,
	}
	p.Paint(writer)
}
