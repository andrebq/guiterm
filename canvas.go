package wdeterm

import (
	"github.com/Agon/freetype-go"
	"image"
	"image/color"
	"image/draw"
)

type Canvas struct {
	img      draw.Image
	ctx      *freetype.Context
	fontSize int32
}

func NewCanvas(img draw.Image, fontSize int32, font []byte) (*Canvas, error) {
	ftFont, err := freetype.ParseFont(font)
	if err != nil {
		return nil, err
	}

	ctx := freetype.NewContext()
	ctx.SetFont(ftFont)
	ctx.SetFontSize(float64(fontSize))
	spaceIdx := ftFont.Index(rune(20)) // space
	width := ftFont.HMetric(fontSize, spaceIdx)
	cellWidth := width.AdvanceWidth + width.LeftSideBearing

	bounds := img.Bounds()
	lines := make([][]rune, bounds.Dy()/int(fontSize))
	for i, _ := range lines {
		lines[i] = make([]rune, bounds.Dx()/int(cellWidth))
	}
	ctx.SetDst(img)
	ctx.SetSrc(image.NewUniform(color.Black))
	return &Canvas{img: img, ctx: ctx, fontSize: fontSize}, nil
}
