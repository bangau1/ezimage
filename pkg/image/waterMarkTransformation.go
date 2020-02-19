package image

//WaterMarkProcessing is a class that can be used for applying watermark
type WaterMarkProcessing struct {
	//WaterMark is a watermark image
	WaterMark Image
}

//NewWaterMarkProcessing return an instance of WaterMarkProcessing
func NewWaterMarkProcessing(watermark Image) WaterMarkProcessing {
	return WaterMarkProcessing{
		WaterMark: watermark,
	}
}

//Apply is a waterMarkProcessing apply method
func (p WaterMarkProcessing) Apply(image Image) Result {
	return Result{}
}
