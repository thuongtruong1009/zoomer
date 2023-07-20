package usecase

import (
	"io"
	"strings"
	"time"
	"fmt"
	"context"
	"mime/multipart"
	"path/filepath"
	"os"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

type localResourceUsecase struct {}

func NewLocalResourceUseCase() ILocalResourceUseCase {
	return &localResourceUsecase{}
}

func (lu *localResourceUsecase) UploadSingleFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	fileExt := filepath.Ext(file.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(constants.UploadPath + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}
	return filename, nil
}

func (lu *localResourceUsecase) UploadMultipleFile(ctx context.Context, files []*multipart.FileHeader) ([]string, error) {
	filePaths := []string{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

		filePaths = append(filePaths, filename)
		out, err := os.Create(constants.UploadPath + filename)
		if err != nil {
			return nil, err
		}
		defer out.Close()

		readerFile, _ := file.Open()
		if _, err := io.Copy(out, readerFile); err != nil {
			return nil, err
		}
	}
	return filePaths, nil
}

func (lu *localResourceUsecase) DeleteSingleFile(ctx context.Context, fileName string) error {
	err := os.Remove(constants.UploadPath + fileName)
	if err != nil {
		return err
	}
	return nil
}

func (lu *localResourceUsecase) DeleteMultipleFile(ctx context.Context, fileNames []string) error {
	for _, fileName := range fileNames {
		err := os.Remove(constants.UploadPath + fileName)
		if err != nil {
			return err
		}
	}
	return nil
}
