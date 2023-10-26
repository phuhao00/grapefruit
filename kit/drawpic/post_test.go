package drawpic

import (
	"fmt"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	textBrush, _ := NewTextBrush("./font.ttc", 80, image.Black, 50)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(textBrush.FontType)
	c.SetHinting(font.HintingFull)
	c.SetFontSize(textBrush.FontSize)

	textBrush.FontColor = image.NewUniform(color.RGBA{
		R: 0xFF,
		G: 255,
		B: 255,
		A: 255,
	})
	c.SetSrc(textBrush.FontColor)
	now := time.Now()
	if _, err := Poster(c); err != nil {
	} else {
		//f, _ := os.Create("./image1.png")
		//png.Encode(f, img)
	}
	fmt.Println(time.Since(now).Seconds())
}
