package image

import (
	"github.com/pkg/errors"
	"image"
	"image/draw"
	"github.com/nfnt/resize"
)

type Margin struct {
	width float32
	height float32
}

type Resize struct{
	width float32
	height float32
}

var noResize = Resize{
	width:  1,
	height: 1,
}
//WaterMarkProcessing is a class that can be used for applying watermark
type WaterMarkProcessing struct {
	//WaterMark is a watermark image
	WaterMark *Image

	//RelativePosition where the watermark image being positioned
	RelativePosition RelativePosition

	//Margin is a margin information used in accordance with the RelativePosition
	Margin Margin

	//ResizePercentage is resize percentage being applied to watermark image
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
		RelativePosition: BottomRight,
		Margin: Margin{
			width:  0.5,
			height: 0.5,
		},
		ResizePercentage: Resize{
			width:  1,
			height: 1,
		},
	}
}

//Apply is a waterMarkProcessing apply method
func (p WaterMarkProcessing) Apply(img *Image) Result {
	if img == nil || img.InternalImage == nil{
		return Result{Error:errors.New("input `image` is nil")}
	}

	targetBound := img.InternalImage.Bounds()
	targetImg := image.NewRGBA(targetBound)

	// draw base image
	draw.Draw(targetImg, targetBound, img.InternalImage, image.Pt(0,0), draw.Src)

	// resize the watermark image
	var resizedWatermarkImg = Image{InternalImage:p.WaterMark.InternalImage}
	if p.ResizePercentage != noResize{
		rect := resizedWatermarkImg.InternalImage.Bounds()

		targetW := uint(float32(rect.Size().X) * p.ResizePercentage.width)
		targetH := uint(float32(rect.Size().Y) * p.ResizePercentage.height)

		resizedWatermarkImg.InternalImage = resize.Resize(targetW, targetH,resizedWatermarkImg.InternalImage, resize.MitchellNetravali)
	}

	// draw the watermark with margin
	watermarkBound := resizedWatermarkImg.InternalImage.Bounds()
	var drawPoint image.Point
	marginW := int(p.Margin.width * float32(targetBound.Size().X))
	marginH := int(p.Margin.height * float32(targetBound.Size().Y))

	switch p.RelativePosition {
	case TopLeft:
		drawPoint = image.Point{
			X: marginW,
			Y: marginH,
		}
		break
	case TopRight:
		drawPoint = image.Point{
			X: targetBound.Size().X - (marginW + watermarkBound.Size().X),
			Y: marginH,
		}
		break
	case BottomLeft:
		drawPoint = image.Point{
			X: marginW,
			Y: targetBound.Size().Y - (marginH + watermarkBound.Size().Y),
		}
		break
	case BottomRight:
		drawPoint = image.Point{
			X: targetBound.Size().X - (marginW + watermarkBound.Size().X),
			Y: targetBound.Size().Y - (marginH + watermarkBound.Size().Y),
		}
		break
	}

	draw.Draw(targetImg, targetBound.Add(drawPoint), resizedWatermarkImg.InternalImage, image.Pt(0, 0), draw.Over)


	result := Result{
		Data: &Image{InternalImage:targetImg},
		Error:nil,
	}
	return result
}
