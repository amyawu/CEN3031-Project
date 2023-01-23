package services

import (
	"github.com/go-playground/validator/v10"
	"web-service-gin/helper"
	"web-service-gin/models"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.File) (string, error)
	RemoteUpload(url models.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file models.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url models.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}

//--------------------------------------------------------

type mediaViewer interface {
	FileView(url string) ([]byte, error)
	RemoteView(url string) ([]byte, error)
}

func NewMediaViewer() mediaViewer {
	return &media{}
}

func (*media) FileView(url string) ([]byte, error) {
	// Retrieve image data from Cloudinary
	//data, err := helper.ImageRetrieveHelper(url)
	//if err != nil {
	//	return nil, err
	//}
	//return data, nil
	return nil, nil
}

func (*media) RemoteView(url string) ([]byte, error) {
	// Retrieve image data from Cloudinary
	//data, err := helper.ImageRetrieveHelper(url)
	//if err != nil {
	//	return nil, err
	//}
	//return data, nil
	return nil, nil
}
