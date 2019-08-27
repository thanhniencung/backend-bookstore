package model

type Role int

const (
	ADMIN Role = iota
	MEMBER
)

func (r Role) String() string {
	return [...]string{"ADMIN", "MEMBER"}[r]
}
