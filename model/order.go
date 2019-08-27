package model

import "time"

type Order struct {
	OrderId      string      `json:"orderId,omitempty" db:"order_id,omitempty"`
	UserId       string      `json:"UserId,omitempty" db:"user_id,omitempty"`
	Status  	 string      `json:"status,omitempty" db:"status,omitempty"`
	CreatedAt    time.Time   `json:"createdAt,omitempty" db:"created_at,omitempty"`
	UpdatedAt    time.Time   `json:"updatedAt,omitempty" db:"updated_at,omitempty"`
	Total    	 float64     `json:"total,omitempty"`
	Items 		 []Card		 `json:"items,omitempty"`
}
