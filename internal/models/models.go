package models

type Info struct {
	ID  uint
	Smt string
}
type User struct {
	ID             uint
	Username       string
	Email          string
	HashedPassword string
	IsSuperuser    bool
}
