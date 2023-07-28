package usecase

import (
	"context"
	"fmt"
	"github.com/thuongtruong1009/zoomer/internal/resources/local/presenter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type localResourceUsecase struct{}

func NewLocalResourceUseCase() ILocalResourceUseCase {
	return &localResourceUsecase{}
}

func (lu *localResourceUsecase) UploadSingleFile(ctx context.Context, file *multipart.FileHeader) (*presenter.SingleUploadResponse, error) {
	fileExt := filepath.Ext(file.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

	src, err := file.Open()
	if err != nil {
		log.Println("Failed to open file", err)
		return nil, err
	}
	defer src.Close()

	dst, err := helpers.LockFuncOneInTwoOut[string, *os.File, error](os.Create)(constants.UploadPath + filename)
	if err != nil {
		log.Println("Failed to create file", err)
		return nil, err
	}
	defer dst.Close()

	fileSize, err := io.Copy(dst, src)
	if err != nil {
		log.Println("Failed to calculate file size", err)
		return nil, err
	}

	res := &presenter.SingleUploadResponse{
		Image: constants.UploadPath + filename,
		Size:  fileSize,
	}

	return res, nil
}

func (lu *localResourceUsecase) UploadMultipleFile(ctx context.Context, files []*multipart.FileHeader) (*presenter.MultipleUploadResponse, error) {
	res := &presenter.MultipleUploadResponse{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

		out, err := helpers.LockFuncOneInTwoOut[string, *os.File, error](os.Create)(constants.UploadPath + filename)
		if err != nil {
			return nil, err
		}
		defer out.Close()

		readerFile, err2 := file.Open()
		if err2 != nil {
			log.Println("Failed to open file", err2)
			return nil, err2
		}
		defer readerFile.Close()

		fileSize, err3 := io.Copy(out, readerFile)
		if err3 != nil {
			return nil, err3
		}

		once := &presenter.SingleUploadResponse{
			Image: constants.UploadPath + filename,
			Size:  fileSize,
		}

		res.Images = append(res.Images, *once)
	}
	return res, nil
}

func (lu *localResourceUsecase) DeleteSingleFile(ctx context.Context, fileName string) error {
	err := helpers.LockFuncOneInOneOut[string, error](os.Remove)(constants.UploadPath + fileName)
	if err != nil {
		return err
	}
	return nil
}

func (lu *localResourceUsecase) DeleteMultipleFile(ctx context.Context, fileNames []string) error {
	for _, fileName := range fileNames {
		err := helpers.LockFuncOneInOneOut[string, error](os.Remove)(constants.UploadPath + fileName)
		if err != nil {
			return err
		}
	}
	return nil
}
