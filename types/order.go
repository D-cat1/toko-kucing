package types

type Item struct {
	Nama     string
	qty      int
	Price    int
	Category string
}

type Order struct {
	Id         int
	InvNum     string
	UserData   User
	TotalPrice int
	Item       []Item
}

type DataOrder struct {
	LastId int
	Length int
	Data   []Order
}
