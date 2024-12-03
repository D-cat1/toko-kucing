package types

type ProductTemp struct {
	Id         int
	UserId     int
	Nama       string
	Stok       int
	Price      int
	CategoryId int
}

type DataProductTemp struct {
	LastId int
	Length int
	Data   []ProductTemp
}
