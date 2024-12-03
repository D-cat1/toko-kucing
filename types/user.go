package types

type User struct {
	Id       int
	Username string
	Password string
	Nama     string
	Role     string
}

type DataUser struct {
	LastId int
	Length int
	Data   []User
}
