package fileupload

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type FileUpload interface {
	Upload(file multipart.FileHeader) (*string, error)
	Delete(filepath string) error
}

type FileUploadImpl struct {
}

// Delete implements FileUpload.
func (*FileUploadImpl) Delete(filepath string) error {
	panic("unimplemented")
}

// Upload implements FileUpload.
func (*FileUploadImpl) Upload(file multipart.FileHeader) (*string, error) {

	url := os.Getenv("CLOUDINARY_URL")
	cloud, err := cloudinary.NewFromURL(url)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	binary, _ := file.Open()

	uploadResult, err := cloud.Upload.Upload(ctx, binary, uploader.UploadParams{
		Folder:   "gobook",
		PublicID: uuid.New().String(),
	})

	if err != nil {
		return nil, err
	}

	return &uploadResult.SecureURL, nil
}

func NewFileUpload() FileUpload {
	return &FileUploadImpl{}
}
