package controller

type Record struct {
	Id        uint64 `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt int64  `json:"deleted_at" gorm:"index;column:deleted_at;default:0"`

	Name string `json:"name" gorm:"name"`
}
