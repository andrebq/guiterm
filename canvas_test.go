package guiterm

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteString(t *testing.T) {
	fontData, err := ioutil.ReadFile("./Anonymous-Pro/Anonymous_Pro.ttf")
	if err != nil {
		t.Fatalf("Unable to read font: %v", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	canvas, err := NewCanvas(img, 18, fontData)

	if err != nil {
		t.Fatalf("Unable to create canvas: %v", err)
	}
	canvas.Clear()

	err = canvas.SetCell(5, 1, 'm', color.Black, color.White)
	if err != nil {
		t.Fatalf("Unable to render cell. %v", err)
	}

	file, err := os.Create("out.png")
	if err != nil {
		t.Fatalf("Unable to open out.png for writing")
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Errorf("Unable to save image to out.png")
	}

	_ = canvas
}
