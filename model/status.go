package model

type OrderStatus int

const (
	ORDERING OrderStatus = iota
	CONFIRM
)

func (o OrderStatus) String() string {
	return [...]string{"ORDERING", "CONFIRM"}[o]
}
