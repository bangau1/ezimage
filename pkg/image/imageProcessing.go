package image

import (
	"fmt"
	"github.com/pkg/errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

//Image is a struct
type Image struct {
	InternalImage image.Image
}

//Transformation is a transformation interface
type Transformation interface {
	Apply(*Image) Result
}

//Result is a result
type Result struct {
	Data  *Image
	Error error
}

const (
	MIME_TYPE_JPEG = "image/jpeg"
	MIME_TYPE_PNG  = "image/png"
)

//NewImageFromLocation return an image from location
func NewImageFromLocation(location string) (*Image, error) {
	imageFile, err := os.Open(location)

	if err != nil {
		return nil, errors.Wrap(err, "Couldn't open the image")
	}

	defer imageFile.Close()
	return NewImageFromReader(imageFile)
}

//NewImageFromReader return an image from Reader
func NewImageFromReader(reader io.Reader) (*Image, error) {
	if reader == nil {
		return nil, errors.New("Reader can't be nil")
	}

	var result *Image
	var pErr error
	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Can't decode the image"))
	}
	fmt.Printf("Format is %s\n", format)

	result = &Image{InternalImage: img}
	return result, pErr
}

func (img Image) Save(location, mimeType string, jpegQuality int) error {

	imageFile, err := os.Create(location)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Can't write to file location %s", location))
	}
	defer imageFile.Close()

	if MIME_TYPE_JPEG == mimeType {
		return jpeg.Encode(imageFile, img.InternalImage, &jpeg.Options{Quality:jpegQuality})
	} else if MIME_TYPE_PNG == mimeType {
		return png.Encode(imageFile, img.InternalImage)
	} else {
		return errors.New(fmt.Sprintf("Unsupported mimeType: %s", mimeType))
	}
}
