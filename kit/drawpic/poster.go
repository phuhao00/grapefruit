package drawpic

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"os"
)

func Poster(c *freetype.Context) (image.Image, error) {
	var (
		bgFile *os.File
		bgImg  image.Image
		dingZi image.Image
		//qrCodeImg    image.Image
		pic image.Image
		//offset       image.Point
		dingZiOffset image.Point
		picOffset    image.Point
	)

	// 01: 打开背景图片
	bgFile, err = os.Open("./20231025-165747.png")
	if err != nil {
		fmt.Println("打开背景图片失败", err)
		return nil, err
	}

	defer bgFile.Close()

	// 02: 编码为图片格式
	bgImg, err = png.Decode(bgFile)
	if err != nil {
		fmt.Println("背景图片编码失败:", err)
		return nil, err
	}
	//dingZi
	dingziFile, err := os.Open("./20231025-171032.png")
	if err != nil {
		fmt.Println("打开背景图片失败", err)
		return nil, err
	}

	defer dingziFile.Close()

	// 02: 编码为图片格式
	dingZi, err = png.Decode(dingziFile)
	if err != nil {
		fmt.Println("背景图片编码失败:", err)
		return nil, err
	}
	//pic
	url2 := "https://www.ourchinastory.com/images/cover/snapshot/2021/01/square/%E7%95%B6%E4%BB%A3%E4%B8%AD%E5%9C%8B-%E4%B8%AD%E5%9C%8B%E6%97%85%E9%81%8A-%E4%B8%AD%E5%9C%8B%E6%96%87%E5%8C%96-%E4%B8%8A%E6%B5%B7-%E5%A4%96%E7%81%98_x1.jpg"
	pic, err = ImageZoom(url2, 1200)
	if err != nil {
		fmt.Println("Error fetching image 2:", err)
		return nil, err
	}
	//picFile, err := os.Open("./3hamburger-01.png")
	//if err != nil {
	//	fmt.Println("打开背景图片失败", err)
	//	return nil, err
	//}
	//
	//defer picFile.Close()

	// 02: 编码为图片格式
	//pic, err = png.Decode(picFile)
	//if err != nil {
	//	fmt.Println("背景图片编码失败:", err)
	//	return nil, err
	//}
	// 03: 生成二维码
	//qrCodeImg, err = CreateAvatar()
	//if err != nil {
	//	fmt.Println("生成二维码失败:", err)
	//	return nil, err
	//}

	//offset = image.Pt(1200, 1600)      //用于调整二维码在背景图片上的位置
	dingZiOffset = image.Pt(1200, 200) //用于调整二维码在背景图片上的位置
	picOffset = image.Pt(150, 200)     //用于调整二维码在背景图片上的位置
	// Set the corner radius
	//radius := 100.0
	//
	//// Create a new context with the same size as the image
	//dc := gg.NewContextForImage(bgImg)
	//// Draw the image on the context
	//dc.DrawImage(bgImg, 0, 0)
	//
	//// Create a rounded corner path for the top-right corner
	//x := float64(bgImg.Bounds().Dx())
	//y := 0.0
	//dc.NewSubPath()
	//dc.MoveTo(x-radius, y)
	//dc.LineTo(x, y)
	//dc.LineTo(x, y+radius)
	//controlX := x
	//controlY := y
	//endX := x - radius
	//endY := y + radius
	//dc.CubicTo(controlX, controlY, controlX, controlY, endX, endY)
	//
	//// Clip the context to the rounded corner path
	//dc.Clip()
	//
	//// Clear the context and redraw the image
	//dc.Clear()
	//dc.DrawImage(bgImg, 0, 0)

	b := bgImg.Bounds()

	m := image.NewRGBA(b)

	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	//draw.Draw(m, qrCodeImg.Bounds().Add(offset), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Over)
	draw.Draw(m, pic.Bounds().Add(picOffset), pic, image.Point{X: 0, Y: 0}, draw.Over)
	draw.Draw(m, dingZi.Bounds().Add(dingZiOffset), dingZi, image.Point{X: 0, Y: 0}, draw.Over)
	des := m

	//新建笔刷

	//Px Py 绘图开始坐标 text要绘制的文字
	//调整颜色
	c.SetClip(des.Bounds())
	c.SetDst(des)
	c.DrawString("8月24日 city walk 路线  ", freetype.Pt(150, 1300))

	fSave, err := os.Create("./image1.png")
	if err != nil {
		return nil, err
	}
	defer fSave.Close()

	err = png.Encode(fSave, des)
	return nil, err

}

func fetchImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ImageZoom(url string, width uint) (image.Image, error) {
	m, err := fetchImage(url)
	if err != nil {
		return nil, err
	}
	if width == 0 {
		return m, nil
	}
	thImg := resize.Resize(width, 0, m, resize.Lanczos3)
	return thImg, nil
}
