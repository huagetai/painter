package view

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/huagetai/painter/log"
	"github.com/huagetai/painter/utils/checksum"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
	"gopkg.in/fogleman/gg.v1"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
)

type ImageView struct {
	X           int
	Y           int
	Width       int
	Height      int
	Mode        string
	URI         string
	BorderRadis float64
}

func (v *ImageView) Paint(ctx *PosterContext) {
	srcReader, format, err := URIResourceReader(v.URI)
	if err != nil {
		fmt.Errorf("URIResourceReader err：%v", err)
	}
	srcImage, err := decoodeImage(srcReader, format)
	if err != nil {
		fmt.Errorf("image.Decode err：%v", err)
	}
	srcBounds := srcImage.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	dc := gg.NewContext(v.Width, v.Height)
	if v.BorderRadis > 0 {
		dc.DrawRoundedRectangle(0, 0, float64(v.Width), float64(v.Height), v.BorderRadis)
		dc.Clip()
	}
	switch v.Mode {
	case "smart":
		topCrop, err := smartCrop(srcImage, v.Width, v.Height)
		if err != nil {
			fmt.Errorf("smartCrop err：%v", err)
		}
		sub, ok := srcImage.(SubImager)
		if ok {
			cropImage := sub.SubImage(topCrop)
			i := imaging.Fill(cropImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
			dc.DrawImage(i, 0, 0)
		} else {
			fmt.Errorf("No SubImage support：%v", err)
		}
	case "scaleToFill":
		srcImage = imaging.Resize(srcImage, v.Width, v.Height, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "aspectFit":
		srcImage = imaging.Fit(srcImage, v.Width, v.Height, imaging.Lanczos)
		dc.DrawImage(srcImage, 0, 0)
	case "aspectFill":
		if srcW > srcH {
			srcImage = imaging.Resize(srcImage, 0, v.Height, imaging.Lanczos)
			srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
			dc.DrawImage(srcImage, 0, 0)
		} else {
			srcImage = imaging.Resize(srcImage, v.Width, 0, imaging.Lanczos)
			srcImage = imaging.Fill(srcImage, v.Width, v.Height, imaging.TopLeft, imaging.Lanczos)
			dc.DrawImage(srcImage, 0, 0)
		}
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
	rdc := ctx.Ctx
	rdc.DrawImage(dc.Image(), v.X, v.Y)
}

func (v *ImageView) calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(float64(v.Width)/float64(srcWidth), float64(v.Height)/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// Moved here and unexported to decouple the resizer implementation.
func smartCrop(img image.Image, width, height int) (image.Rectangle, error) {
	analyzer := smartcrop.NewAnalyzer(nfnt.NewResizer(resize.Lanczos3))
	return analyzer.FindBestCrop(img, width, height)
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func decoodeImage(r io.Reader, format string) (srcImage image.Image, err error) {
	switch format {
	case "png":
		return png.Decode(r)
	case "jpeg":
		return jpeg.Decode(r)
	case "webp":
		return webp.Decode(r)
	case "gif":
		return gif.Decode(r)
	default:
		return jpeg.Decode(r)
	}
}

func URIResourceReader(uri string) (r *bytes.Reader, format string, err error) {
	if uri[0:4] == "http" {
		filename := "image-" + checksum.MD5Str(uri)
		fileBytes, err := os.ReadFile("./assets/" + filename + ".jpeg")
		if err == nil {
			log.Info("remote image hit:%v \n", uri)
			format = "jpeg"
			r = bytes.NewReader(fileBytes)
		} else {
			log.Info("remote image miss:%v \n", uri)
			resp, err := http.Get(uri)
			if err != nil {
				return nil, "", err
			}
			defer resp.Body.Close()
			fileBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, "", err
			}
			_, format, _ = image.DecodeConfig(bytes.NewReader(fileBytes))
			if format == "" {
				format = resFormat(resp)
			}
			r = bytes.NewReader(fileBytes)
			os.WriteFile("./assets/"+filename+"."+format, fileBytes, 777)
		}
	} else {
		log.Info("local image hit:%v \n", uri)
		fileBytes, err := os.ReadFile(uri)
		if err != nil {
			return nil, "", err
		}
		_, format, _ = image.DecodeConfig(bytes.NewReader(fileBytes))
		if format == "" {
			format = fileFormat(uri)
		}
		r = bytes.NewReader(fileBytes)
	}
	return r, format, nil
}

func fileFormat(path string) (format string) {
	ext := filepath.Ext(path)
	switch ext {
	case ".png":
		format = "png"
	case ".jpeg":
		format = "jpeg"
	case ".jpg":
		format = "jpeg"
	case ".webp":
		format = "webp"
	case ".gif":
		format = "gif"
	default:
		format = "jpeg"
	}
	return
}

func resFormat(resp *http.Response) (format string) {
	contentType := resp.Header.Get("Content-Type")
	switch contentType {
	case "image/png":
		format = "png"
	case "image/jpeg":
		format = "jpeg"
	case "image/jpg":
		format = "jpeg"
	case "image/webp":
		format = "webp"
	case "image/gif":
		format = "gif"
	default:
		format = "jpeg"
	}
	return
}
