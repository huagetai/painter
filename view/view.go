package view

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

type BaseView struct {
}
