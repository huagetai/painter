package painter

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/huagetai/painter/utils/checksum"
	"github.com/huagetai/painter/utils/log"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
	"gopkg.in/fogleman/gg.v1"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
)

type ImageView struct {
	X            int
	Y            int
	Width        int
	Height       int
	Mode         string
	URI          string
	BorderRadius float64
}

func (v *ImageView) Paint(ctx *PosterContext) {
	srcReader, err := imageResourceReader(v.URI)
	if err != nil {
		log.Error("%v URIResourceReader err：%v \n", v.URI, err)
		return
	}
	srcImage, format, err := image.Decode(srcReader)
	if err != nil {
		log.Error("%v image.Decode err：%v \n", v.URI, err)
		return
	}
	log.Info("%v %v \n", v.URI, format)

	dc := gg.NewContext(v.Width, v.Height)
	if v.BorderRadius > 0 {
		dc.DrawRoundedRectangle(0, 0, float64(v.Width), float64(v.Height), v.BorderRadius)
		dc.Clip()
	}
	switch v.Mode {
	case "smart":
		v.smart(srcImage, dc)
	case "scaleToFill":
		srcImage = imaging.Resize(srcImage, v.Width, v.Height, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "aspectFit":
		srcImage = imaging.Fit(srcImage, v.Width, v.Height, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "aspectFill":
		v.aspectFill(srcImage, dc)
	case "top":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.Top, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "bottom":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.Bottom, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "center":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.Center, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "left":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.Left, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "right":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.Right, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "top left":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "top right":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.TopRight, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "bottom left":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.BottomLeft, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "bottom right":
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.BottomRight, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	default:
		v.aspectFill(srcImage, dc)
	}
	rdc := ctx.Ctx
	rdc.DrawImage(dc.Image(), v.X, v.Y)
}

func (v *ImageView) smart(srcImage image.Image, dc *gg.Context) {
	analyzer := smartcrop.NewAnalyzer(nfnt.NewResizer(resize.Lanczos2))
	topCrop, err := analyzer.FindBestCrop(srcImage, v.Width, v.Height)
	if err != nil {
		v.aspectFill(srcImage, dc)
	}
	sub, ok := srcImage.(SubImager)
	if ok {
		cropImage := sub.SubImage(topCrop)
		i := imaging.Fill(cropImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
		dc.DrawImage(i, 0, 0)
	} else {
		v.aspectFill(srcImage, dc)
	}
}

func (v *ImageView) aspectFill(srcImage image.Image, dc *gg.Context) {
	srcBounds := srcImage.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	if srcW > srcH {
		srcImage = imaging.Resize(srcImage, 0, v.Height, imaging.Lanczos)
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	} else if srcW < srcH {
		srcImage = imaging.Resize(srcImage, v.Width, 0, imaging.Lanczos)
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	} else {
		srcImage = imaging.Resize(srcImage, v.Width, 0, imaging.Lanczos)
		srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.Center, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	}
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func imageResourceReader(uri string) (r *bytes.Reader, err error) {
	var fileBytes []byte
	if uri[0:4] == "http" {
		filename := "image-" + checksum.MD5Str(uri)
		fileBytes, err = os.ReadFile("./image/" + filename + ".jpeg")
		if err == nil {
			log.Info("image cache hit:%v \n", uri)
		} else {
			log.Info("image cache miss:%v \n", uri)
			resp, err := http.Get(uri)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			fileBytes, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			os.WriteFile("./image/"+filename+".jpeg", fileBytes, 0777)
		}
	} else {
		log.Info("local image:%v \n", uri)
		fileBytes, err = os.ReadFile(uri)
		if err != nil {
			return nil, err
		}
	}
	r = bytes.NewReader(fileBytes)
	return r, nil
}
