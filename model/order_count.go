package model

type OrderCount struct {
	OrderId    string   `json:"orderId,omitempty" db:"order_id,omitempty"`
	Total      int      `json:"total,omitempty" db:"total,omitempty"`
}
