package view

import (
	"fmt"
	"tri_darma/services"
	"tri_darma/types"
)

func inputRegTriDarma(data *types.TriDarma, data_type string) {
	Clrscr()
	fmt.Printf(border("-", "Detail %v", 50)+"\n", data_type)
	fmt.Printf("Judul %v : ", data_type)
	HandleLongInput(&data.Nama)
	data.Tipe = data_type
	fmt.Printf("Tahun %v : ", data_type)
	fmt.Scan(&data.Tahun)
}

func handleRegister(state *types.TriDarma) {
	var choice int
	var dataTriDarma types.TriDarma
	Clrscr()
	fmt.Println(border("-", "Register", 50))
	fmt.Println("1. Penelitian")
	fmt.Println("2. Abdimas")
	fmt.Println("3. Exit")
	fmt.Println(border("-", "", 50))
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		inputRegTriDarma(&dataTriDarma, "Penelitian")
	case 2:
		inputRegTriDarma(&dataTriDarma, "Abdimas")
	case 3:
		return
	default:
		fmt.Println("Opsi tidak tersedia")
		delay(1)
		return
	}

	services.Add3Darma(dataTriDarma)

	var lastidx = services.ListTridar().Length
	*state = services.ListTridar().Data[lastidx-1]
}

func HandleManagemen(dataTriSaved *types.TriDarma) {
	var choice int
	var lastidx = services.ListTridar()
	Clrscr()
	fmt.Println(border("-", "Pilih Judul Penelitian", 50))
	for i := 0; i < lastidx.Length; i++ {

		fmt.Printf("%v. %v\n", i+1, lastidx.Data[i].Nama)
	}
	fmt.Println(border("-", "", 50))

	for choice > lastidx.Length || choice == 0 {
		fmt.Print("Pilih : ")
		fmt.Scanln(&choice)
		fmt.Println("Inputan Salah! ulangi")
	}
	*dataTriSaved = lastidx.Data[choice-1]
	return
}

func MainMenu(managedState *types.TriDarma) {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("-", "Silahkan Pilih menu dibawah", 50))
		fmt.Println("1. Registrasi Tri Darma")
		fmt.Println("2. Manage Tri Darma")
		fmt.Println("3. Exit")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handleRegister(managedState)
			return
		case 2:
			if services.ListTridar().Length < 1 {
				fmt.Println("Data Tri Darma masih kosong!")
				delay(1)
				fmt.Println("Kembali ke menu awal...")
				delay(2)
			} else {
				HandleManagemen(managedState)
				return
			}
		case 3:
			fmt.Println("babay!")
			return
		default:
			fmt.Println("Opsi tidak tersedia")
			delay(1)
		}
	}
}
