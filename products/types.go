package products

import (
	"time"
)

type Products struct {
	// gorm.Model
	ProductID          uint              `json:"product_id" gorm:"primary_key"`
	ProductName        string            `json:"product_name"`
	ProductQuantity    string            `json:"products_quantity"`
	ProductDescription string            `json:"product_description"`
	ProductPrice       uint              `json:"product_price"`
	SKU                string            `json:"sku"`
	CreatedAt          time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time         `json:"updated_at" gorm:"autoUpdateTime:milli"`
	ProductCategory    []CategoryProduct `json:"category" gorm:"foreignkey:ProductID"`
}

type CategoryProduct struct {
	// gorm.Model
	ID         uint      `json:"id" gorm:"primary_key"`
	CategoryID uint      `json:"categoryId"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
	ProductID  uint      `json:"-"`
}
