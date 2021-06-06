package categories

import "time"

type Categories struct {
	CategoryID          uint      `json:"category_id" gorm:"primary_key"`
	CategoryName        string    `json:"category_name"`
	CategoryDescription string    `json:"category_description"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
