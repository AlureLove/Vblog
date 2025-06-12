package utils

import "time"

type ResourceMeta struct {
	Id        uint       `json:"id" gorm:"primary_key;column:id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy string     `json:"created_by" gorm:"created_by;type:varchar(100)"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func NewResourceMeta() *ResourceMeta {
	return &ResourceMeta{
		CreatedAt: time.Now(),
	}
}
