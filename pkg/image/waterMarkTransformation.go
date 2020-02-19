package image

import (
	"github.com/pkg/errors"
	"image"
	"image/draw"
)

type Point image.Point
type Resize struct{
	width float32
	height float32
}
//WaterMarkProcessing is a class that can be used for applying watermark
type WaterMarkProcessing struct {
	//WaterMark is a watermark image
	WaterMark *Image

	//RelativePosition where the watermark image being positioned
	RelativePosition RelativePosition

	//Point is a margin information used in accordance with the RelativePosition
	Point Point

	//ResizePercentage is resize percentage being applied to watermark image, relative to the source/base image.
	//Example:
	//-`Resize{0.4, 0.3}` will resize the watermark's width to 40% of the source image width and watermark's height to 30% of source image height.
	ResizePercentage Resize
}

// RelativePosition
type RelativePosition int

const (
	// TopLeft is relative from top-left
	TopLeft RelativePosition = iota
	// TopRight is relative from top-right
	TopRight
	// BottomLeft is relative from bottom-left
	BottomLeft
	// BottomRight is relative from bottom-right
	BottomRight
)

//NewWaterMarkProcessing return an instance of WaterMarkProcessing
func NewWaterMarkProcessing(watermark *Image) WaterMarkProcessing {
	return WaterMarkProcessing{
		WaterMark: watermark,
		RelativePosition:BottomRight,
		Point:
	}
}

//Apply is a waterMarkProcessing apply method
func (p WaterMarkProcessing) Apply(img *Image) Result {
	if img == nil || img.InternalImage == nil{
		return Result{Error:errors.New("input `image` is nil")}
	}

	targetBound := img.InternalImage.Bounds()
	targetImg := image.NewRGBA(targetBound)


	draw.Draw(targetImg, targetBound, img.InternalImage, image.Pt(0,0), draw.Src)
	draw.Draw(targetImg, targetBound.Add(image.Pt(300, 200)), p.WaterMark.InternalImage, image.Pt(0, 0), draw.Over)


	result := Result{
		Data: &Image{InternalImage:targetImg},
		Error:nil,
	}
	return result
}
