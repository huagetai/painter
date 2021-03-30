package painter

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/huagetai/painter/utils/log"
	"gopkg.in/fogleman/gg.v1"
	"io"
	"os"
)

type Palette struct {
	BackgroundImage string
	BackgroundColor string
	Width           int
	Height          int
	BorderRadius    float64
	Views           []View
}

func (p *Palette) Paint(w io.Writer) {
	if p.BackgroundColor == "" {
		p.BackgroundColor = "#ffffff"
	}
	j, _ := json.Marshal(*p)
	filename := "poster-" + fmt.Sprintf("%x", md5.Sum(j))
	fileBytes, err := os.ReadFile("./image/" + filename + ".png")
	if err == nil {
		log.Info("poster hit:%v \n", filename)
		io.Copy(w, bytes.NewReader(fileBytes))
		return
	} else {
		log.Info("poster miss:%v \n", filename)
	}
	dc := gg.NewContext(p.Width, p.Height)
	if p.BorderRadius > 0 {
		dc.DrawRoundedRectangle(0, 0, float64(p.Width), float64(p.Height), p.BorderRadius)
		dc.Clip()
		dc.DrawRectangle(0, 0, float64(p.Width), float64(p.Height))
		dc.SetHexColor(p.BackgroundColor)
		dc.Fill()
	} else {
		dc.SetHexColor(p.BackgroundColor)
		dc.Clear()
	}
	ctx := &PosterContext{
		p.Width,
		p.Height,
		dc,
	}
	if p.BackgroundImage != "" {
		bgView := ImageView{
			X:            0,
			Y:            0,
			Width:        p.Width,
			Height:       p.Height,
			URI:          p.BackgroundImage,
			Mode:         "scaleToFill",
			BorderRadius: p.BorderRadius,
		}
		bgView.Paint(ctx)
	}
	for _, v := range p.Views {
		v.Paint(ctx)
	}
	dc.EncodePNG(w)
	dc.SavePNG("./image/" + filename + ".png")
}
