package main

import (
	"fmt"
	"toko_kucing/database"
	"toko_kucing/types"
	"toko_kucing/view"
)

func main() {
	var saveUserLogin types.User
	database.InitDb()
	view.MainMenu(&saveUserLogin)
	switch saveUserLogin.Role {
	case "admin":
		view.AdminMenu(saveUserLogin)
		main()
	case "user":
		fmt.Print("role user approved")
	}
}
