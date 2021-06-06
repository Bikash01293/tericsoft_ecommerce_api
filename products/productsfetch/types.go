package productsfetch

import (
	"time"
)

type Products struct {
	// gorm.Model
	ProductID          uint            `json:"product_id" gorm:"primary_key"`
	ProductName        string          `json:"product_name"`
	ProductQuantity    string          `json:"products_quantity"`
	ProductDescription string          `json:"product_description"`
	ProductCategory    CategoryProduct `json:"product_category" gorm:"foreignkey:ProductID"`
	ProductPrice       uint            `json:"product_price"`
	SKU                string          `json:"sku"`
	CreatedAt          time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time       `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
type CategoryProduct struct {
	// gorm.Model
	ID         uint         `json:"-"`
	CategoryID uint         `json:"-" gorm:"primary_key"`
	CreatedAt  time.Time    `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `json:"-" gorm:"autoUpdateTime:milli"`
	ProductID  uint         `json:"-"`
	Category   []Categories `json:"category" gorm:"foreignkey:CategoryID"`
}
type Categories struct {
	CategoryId          uint      `json:"category_id" gorm:"primary_key"`
	CategoryName        string    `json:"category_name"`
	CategoryDescription string    `json:"category_description"`
	CreatedAt           time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"-" gorm:"autoUpdateTime:milli"`
	CategoryID          uint      `json:"-"`
}
