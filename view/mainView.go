package view

import (
	"fmt"
	"time"
	"toko_kucing/services"
	"toko_kucing/types"
)

func handleRegister() {
	Clrscr()
	var dataUser types.User
	var input string
	var validate bool = true
	fmt.Println(border("-", "Register", 50))
	fmt.Print("Nama : ")
	HandleLongInput(&dataUser.Nama)
	for validate {
		fmt.Print("Username : ")
		fmt.Scan(&input)
		if services.GetUserByUsername(input) != nil {
			fmt.Println("Username sudah terpakai!")
		} else {
			validate = false
			dataUser.Username = input
		}
	}
	validate = true
	fmt.Print("Password : ")
	HandleLongInput(&dataUser.Password)
	dataUser.Role = "user"
	services.AddUser(dataUser)
	fmt.Println("User berhasil ditambahkan!")
	time.Sleep(1 * time.Second)
	MainMenu(&dataUser)
}

func HandleLogin(dataUser *types.User) {
	Clrscr()
	var username, password string
	fmt.Println(border("-", "Log In", 50))
	fmt.Print("Username : ")
	fmt.Scanln(&username)
	fmt.Print("Password : ")
	fmt.Scanln(&password)
	loginSuccess := services.LoginAction(username, password)
	if loginSuccess != nil {
		fmt.Println("log success hallo", loginSuccess.Nama)
		*dataUser = *loginSuccess
	} else {
		fmt.Println("log fail back to menu in 2 seconds")
		time.Sleep(2 * time.Second)
		MainMenu(dataUser)
	}
}

func aserMenu() { // if we use recursive
	var choice int
	var dataUser types.User
	Clrscr()
	fmt.Println(border("-", "Silahkan Pilih menu dibawah", 50))
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println(border("-", "", 50))
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		handleRegister()
	case 2:
		HandleLogin(&dataUser)
		if dataUser.Nama != "" {
			fmt.Println(dataUser.Nama)
		}
		return
	case 3:
		fmt.Println("babay!")
		return
	default:
		fmt.Println("Opsi tidak tersedia")
		time.Sleep(1 * time.Second)
		MainMenu(&dataUser)
	}
}

func MainMenu(userLogin *types.User) {
	var choice int
	var dataUser types.User
	for {
		Clrscr()
		fmt.Println(border("-", "Silahkan Pilih menu dibawah", 50))
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			handleRegister()
		case 2:
			HandleLogin(&dataUser)
			if dataUser.Nama != "" {
				*userLogin = dataUser
				return
			}
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
