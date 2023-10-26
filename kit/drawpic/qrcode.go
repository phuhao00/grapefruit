package drawpic

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	imgtype "github.com/shamsher31/goimgtype"
	"image"
	"image/draw"
	"image/png"
	"os"
)

var err error

//二维码中间贴logo
func CreateAvatar() (image.Image, error) {
	var (
		bgImg      image.Image
		offset     image.Point
		avatarFile *os.File
		avatarImg  image.Image

		datatype string
	)

	bgImg, err = createQrCode("https://www.baidu.com") //调用生成二维码的方法

	if err != nil {
		fmt.Println("创建二维码失败:", err)
		return nil, errors.New("创建二维码失败")
	}
	avatarFile, err = os.Open("./20231025-165747.png")

	datatype, err = imgtype.Get("./20231025-165747.png")
	if err != nil {
		fmt.Printf("这不是图片文件")
	} else {
		if datatype == `image/png` {
			fmt.Printf("这是png图片")
		} else {
			fmt.Printf("这不是png图片")
		}
	}
	avatarImg, err = png.Decode(avatarFile)
	if err != nil {
		fmt.Println("创建二维码失败啊:", err.Error())
		return nil, errors.New("创建二维码失败啊")
	}

	avatarImg = ImageResize(avatarImg, 40, 40)
	b := bgImg.Bounds()

	// 设置为居中
	offset = image.Pt((b.Max.X-avatarImg.Bounds().Max.X)/2, (b.Max.Y-avatarImg.Bounds().Max.Y)/2)

	m := image.NewRGBA(b)

	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)

	draw.Draw(m, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)

	return m, err
}

//创建二维码的方法
func createQrCode(context string) (img image.Image, err error) {
	var qrCode *qrcode.QRCode

	qrCode, err = qrcode.New(context, qrcode.Highest) //创建二维码
	if err != nil {
		return nil, errors.New("二维码创建失败")
	}
	qrCode.DisableBorder = true //禁用二维码的边框
	img = qrCode.Image(150)     //二维码的大小设置(正方形大小 150x150)

	return img, nil
}

//设置二维码的大小
func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
