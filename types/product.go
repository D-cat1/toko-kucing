package types

type Category struct {
	Id   int
	Nama string
}

type DataCategory struct {
	LastId int
	Length int
	Data   []Category
}

type Product struct {
	Id         int
	Nama       string
	Stok       int
	Price      int
	CategoryId int
}

type DataProduct struct {
	LastId int
	Length int
	Data   []Product
}
