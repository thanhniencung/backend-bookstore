package model

import (
	"github.com/lib/pq"
	"time"
)

type Product struct {
	UserId  	 string      `json:"UserId,omitempty" db:"user_id,omitempty""`
	ProductId  	 string    	 `json:"productId,omitempty" db:"product_id,omitempty""`
	ProductName  string    	 `json:"productName,omitempty" db:"product_name,omitempty" valid:"required"`
	ProductImage string    	 `json:"productImage,omitempty" db:"product_image,omitempty" valid:"required,url"`
	Quantity 	 int 	   	 `json:"quantity,omitempty" db:"quantity,omitempty" valid:"required,int"`
	SoldItems 	 int 	   	 `json:"soldItems" db:"sold_items,omitempty"`
	Price 		 float64   	 `json:"price,omitempty" db:"price,omitempty" valid:"required,numeric"`
	CateId 		 string    	 `json:"cateId,omitempty" db:"cate_id,omitempty" valid:"required"`
	CreatedAt 	 time.Time 	 `json:"createdAt,omitempty" db:"created_at,omitempty"`
	UpdatedAt	 time.Time 	 `json:"updatedAt,omitempty" db:"updated_at,omitempty"`
	DeletedAt	 pq.NullTime `json:"-"  db:"deleted_at,omitempty"`
}