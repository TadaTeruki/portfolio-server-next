package repository

import (
	"github.com/TadaTeruki/portfolio-server-next/api/domain/entity"
)

type IActivityRepository interface {
	PostActivity(activity *entity.Activity) (string, error)
	GetActivity(activityID string) (*entity.Activity, error)
	DeleteActivity(activityID string) error
	UpdateActivity(activityID string, updatedActivity *entity.Activity) error
	GetActivities() ([]entity.ListActivity, error)
}
