package repository

import (
	"github.com/TadaTeruki/portfolio-server-next/api/domain/entity"
)

type IPhotoRepository interface {
	PostPhoto(photo *entity.Photo) (string, error)
	GetPhoto(photoID string) (*entity.Photo, error)
	DeletePhoto(photoID string) error
	UpdatePhoto(photoID string, updatedPhoto *entity.Photo) error
	GetPhotos() ([]entity.ListPhoto, error)
}
