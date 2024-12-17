package view

import (
	"fmt"
	"time"
	"tri_darma/services"
	"tri_darma/types"
)

func handleRegister(state *types.TriDarma) {
	var choice int
	var dataTriDarma types.TriDarma
	Clrscr()
	fmt.Println(border("-", "Register", 50))
	fmt.Println("Pilih Tipe")
	fmt.Println("1. Penelitian")
	fmt.Println("2. Abdimas")
	fmt.Println("3. Exit")
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Clrscr()
		fmt.Println(border("-", "Detail Penelitian", 50))
		fmt.Print("Judul Penelitian : ")
		HandleLongInput(&dataTriDarma.Nama)
		dataTriDarma.Tipe = "Penelitian"
		fmt.Print("Tahun Pelaksanaan: ")
		fmt.Scan(&dataTriDarma.Tahun)
	case 2:
		Clrscr()
		fmt.Println(border("-", "Detail Abdimas", 50))
		fmt.Print("Judul Abdimas : ")
		HandleLongInput(&dataTriDarma.Nama)
		dataTriDarma.Tipe = "Abdimas"
		fmt.Print("Tahun Pelaksanaan: ")
		fmt.Scan(&dataTriDarma.Tahun)
	case 3:
		return
	default:
		fmt.Println("Opsi tidak tersedia")
		time.Sleep(1 * time.Second)
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

		fmt.Println(i+1, ". ", lastidx.Data[i].Nama)
	}
	fmt.Println(border("-", "", 50))

	for choice > lastidx.Length || choice == 0 {
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
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
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleRegister(managedState)
			return
		case 2:
			HandleManagemen(managedState)
			return
		case 3:
			fmt.Println("babay!")
			return
		default:
			fmt.Println("Opsi tidak tersedia")
			time.Sleep(1 * time.Second)
		}
	}
}
