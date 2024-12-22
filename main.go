package main

import (
	database "tri_darma/database"
	types "tri_darma/types"
	view "tri_darma/view"
)

/*
	TODO:
	1. ui kurang gacor
	2. belum ada handle buat cancle
	3. bagian view codenya belum terlalu rapi (buru")
	4. (ini baru inget) belum ditambahin buat update data di tampilan, kalau di service udah
	5. kurang penerapan sorting dan binary search
*/

func main() {
	database.InitDb()
	var setDataTriDarma types.TriDarma
	view.MainMenu(&setDataTriDarma)
	if setDataTriDarma.Id != 0 {
		view.PenelitianMenu(&setDataTriDarma)
		main()
	}
}
