package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"xiang/backend/internal/model"
	"xiang/backend/internal/repository"

	"github.com/google/uuid"
)

type UploadService struct {
	uploadDir string
	mediaRepo *repository.MediaRepository
}

func NewUploadService(uploadDir string, mediaRepo *repository.MediaRepository) *UploadService {
	return &UploadService{uploadDir: uploadDir, mediaRepo: mediaRepo}
}

func (s *UploadService) SaveImage(fileHeader *multipart.FileHeader) (*model.Media, error) {
	year, month, day := time.Now().UTC().Date()
	relDir := filepath.Join(fmt.Sprintf("%04d", year), fmt.Sprintf("%02d", month), fmt.Sprintf("%02d", day))
	absDir := filepath.Join(s.uploadDir, relDir)

	if err := os.MkdirAll(absDir, 0o755); err != nil {
		return nil, fmt.Errorf("create upload dir: %w", err)
	}

	ext := filepath.Ext(fileHeader.Filename)
	fileName := uuid.New().String() + ext
	absPath := filepath.Join(absDir, fileName)

	src, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("open upload file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(absPath)
	if err != nil {
		return nil, fmt.Errorf("create target file: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return nil, fmt.Errorf("save upload file: %w", err)
	}

	fileURL := "/uploads/" + filepath.ToSlash(filepath.Join(relDir, fileName))
	media := &model.Media{
		FileName: fileHeader.Filename,
		FileURL:  fileURL,
		MimeType: fileHeader.Header.Get("Content-Type"),
		FileSize: fileHeader.Size,
	}

	if err := s.mediaRepo.Create(media); err != nil {
		return nil, err
	}

	return media, nil
}

func (s *UploadService) SaveImages(fileHeaders []*multipart.FileHeader) ([]*model.Media, error) {
	items := make([]*model.Media, 0, len(fileHeaders))
	for _, file := range fileHeaders {
		media, err := s.SaveImage(file)
		if err != nil {
			return nil, err
		}
		items = append(items, media)
	}
	return items, nil
}
