package model

import (
	"github.com/lib/pq"
	"time"
)

type Cate struct {
	CateId    string      `json:"cateId,omitempty" db:"cate_id,omitempty"`
	CateName  string      `json:"cateName,omitempty" db:"cate_name,omitempty" valid:"required"`
	CreatedAt time.Time   `json:"createdAt,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty" db:"updated_at,omitempty"`
	DeletedAt pq.NullTime `json:"-"  db:"deleted_at,omitempty"`
}
