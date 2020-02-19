package image

import "errors"

//NoOpTransformation is a no-op transformation
type NoOpTransformation struct {
}

//NoOpTransformation return an instance of WaterMarkProcessing
func NewNoOpTransformation() NoOpTransformation {
	return NoOpTransformation{}
}

//Apply is a NoOpTransformation apply method
func (p NoOpTransformation) Apply(image *Image) Result {
	if image == nil || image.InternalImage == nil{
		return Result{Error:errors.New("input `image` is nil")}
	}
	result := Result{
		Data: image,
		Error:nil,
	}
	return result
}
