package model

type Card struct {
	OrderId      string      `json:"orderId,omitempty" db:"order_id,omitempty"`
	//CateId       string      `json:"cateId,omitempty" db:"cate_id,omitempty"`
	ProductId    string      `json:"productId,omitempty" db:"product_id,omitempty"`
	ProductName  string      `json:"productName,omitempty" db:"product_name,omitempty" valid:"required"`
	ProductImage string      `json:"productImage,omitempty" db:"product_image,omitempty" valid:"required,url"`
	Quantity     int         `json:"quantity,omitempty" db:"quantity,omitempty"`
	Price        float64     `json:"price,omitempty" db:"price,omitempty" valid:"required,numeric"`
}
