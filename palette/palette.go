package palette

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/huagetai/painter/log"
	"github.com/huagetai/painter/view"
	"gopkg.in/fogleman/gg.v1"
	"io"
	"os"
)

type Palette struct {
	Background   string
	Width        int
	Height       int
	BorderRadius string
	Views        *[]view.View
}

func (p *Palette) Paint(w *io.Writer) {
	j, _ := json.Marshal(*p)
	filename := "poster-" + fmt.Sprintf("%x", md5.Sum(j))
	fileBytes, err := os.ReadFile("./assets/" + filename + ".png")
	if err == nil {
		log.Info("poster hit:%v \n", filename)
		io.Copy(*w, bytes.NewReader(fileBytes))
		return
	} else {
		log.Info("poster miss:%v \n", filename)
	}
	dc := gg.NewContext(p.Width, p.Height)
	dc.SetHexColor(p.Background)
	dc.Clear()
	ctx := &view.PosterContext{
		p.Width,
		p.Height,
		dc,
	}
	for _, v := range *p.Views {
		v.Paint(ctx)
	}
	dc.EncodePNG(*w)
	dc.SavePNG("./assets/" + filename + ".png")
}
