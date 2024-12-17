package types

type TriDarma struct {
	Id           int
	Nama         string
	Tipe         string
	CountLuaran  int
	CountAnggota int
	SumDana      int
	Tahun        int
}

type DataTriDarma struct {
	LastId int
	Length int
	Data   [100]TriDarma
}
