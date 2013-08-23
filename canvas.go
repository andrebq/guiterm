package guiterm

import (
	"github.com/Agon/freetype-go"
	"github.com/Agon/freetype-go/raster"
	"image"
	"image/color"
	"image/draw"
)

type Canvas struct {
	img   draw.Image
	ctx   *freetype.Context
	fontH int32
	fontW int32
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
	println("width: ", width.AdvanceWidth, "left side bearing: ", width.LeftSideBearing)
	cellWidth := width.AdvanceWidth + width.LeftSideBearing

	bounds := img.Bounds()
	lines := make([][]rune, bounds.Dy()/int(fontSize))
	for i, _ := range lines {
		lines[i] = make([]rune, bounds.Dx()/int(cellWidth))
	}
	ctx.SetDst(img)
	ctx.SetSrc(image.White)
	ctx.SetClip(img.Bounds())
	return &Canvas{img: img, ctx: ctx, fontH: fontSize, fontW: cellWidth}, nil
}

func (c *Canvas) Clear() {
	draw.Draw(c.img, c.img.Bounds(), image.Black, image.ZP, draw.Src)
}

func (c *Canvas) SetCell(line, col int32, r rune, fg, bg color.Color) error {
	println("line: ", line, " col: ", col)
	pt := c.cellCoordToPt(line, col)
	println("pt.x: ", pt.X, " pt.y: ", pt.Y)
	pt = freetype.Pt(10, 10+int(c.ctx.PointToFix32(12)>>8))
	_, err := c.ctx.DrawString(string(r), pt)
	return err
}

func (c *Canvas) cellCoordToPt(line, col int32) raster.Point {
	return freetype.Pt(int(line*c.fontH), int(col*c.fontW))
}
