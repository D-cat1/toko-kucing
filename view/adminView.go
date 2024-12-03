package view

import (
	"fmt"
	"time"
	"toko_kucing/services"
	"toko_kucing/types"
)

func handleCategoryCrt() {
	var catName string
	Clrscr()
	fmt.Println(border("-", "Add Product", 50))
	fmt.Print("Category : ")
	HandleLongInput(&catName)
	services.AddCategory(types.Category{Nama: catName})
	fmt.Println(border("-", "", 50))
	fmt.Println("Category telah disimpan!")
	time.Sleep(1 * time.Second)
}

func handleProductCrt() {
	var prodSpec types.Product
	var choice int
	Clrscr()
	fmt.Println(border("-", "Create Product", 50))
	fmt.Print("Name : ")
	HandleLongInput(&prodSpec.Nama)
	fmt.Print("Stok : ")
	fmt.Scan(&prodSpec.Stok)
	fmt.Print("Price : ")
	fmt.Scan(&prodSpec.Price)
	Clrscr()
	fmt.Println(border("-", "Select Category Product", 50))
	showCategory()
	var success bool = false
	for !success {
		fmt.Println()
		fmt.Print("Pilih id : ")
		fmt.Scan(&choice)
		done, data := services.GetCategoryById(choice)
		if done {
			fmt.Println("category", data, "telah dipilih!")
			prodSpec.CategoryId = choice
		} else {
			fmt.Println("id tersebut tidak ada!")
		}
		success = done
	}
	services.AddProduct(prodSpec)
	fmt.Println(border("-", "", 50))
	fmt.Println("Product telah disimpan!")
	time.Sleep(1 * time.Second)
}

func showCategory() {
	var dataCat types.DataCategory = services.ListCategory()
	printTable(dataCat.Data)

}

func CategoryMenu() {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("-", "Category Menu", 50))
		fmt.Println("1. Add Category")
		fmt.Println("2. Delete Category")
		fmt.Println("3. Show Category")
		fmt.Println("4. Back to previous")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			handleCategoryCrt()
		case 2:
			Clrscr()
			fmt.Println(border("-", "Delete Category", 50))
			showCategory()
			var success bool = false
			for !success {
				fmt.Println()
				fmt.Print("Pilih id yang akan dihapus | (-1) exit: ")
				fmt.Scan(&choice)
				if choice == -1 {
					return
				}
				done, err := services.RemoveCategoryById(choice)
				if !done {
					fmt.Println("Erorr : ", err)
				} else {
					fmt.Println("")
				}
				success = done
			}
		case 3:
			Clrscr()
			fmt.Println(border("-", "Show Category", 50))
			showCategory()
			fmt.Scanln()
		case 4:
			return
		}
	}
}

func ProductMenu() {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("-", "Product Menu", 50))
		fmt.Println("1. Add Product")
		fmt.Println("2. Delete Product")
		fmt.Println("3. Show Product")
		fmt.Println("4. Back to previous")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			handleProductCrt()
		case 2:
			Clrscr()
			fmt.Println(border("-", "Delete Category", 50))
			showCategory()
			var success bool = false
			for !success {
				fmt.Println()
				fmt.Print("Pilih id yang akan dihapus | (-1) exit: ")
				fmt.Scan(&choice)
				if choice == -1 {
					return
				}
				done, err := services.RemoveCategoryById(choice)
				if !done {
					fmt.Println("Erorr : ", err)
				} else {
					fmt.Println("")
				}
				success = done
			}
		case 3:
			Clrscr()
			fmt.Println(border("-", "Show Product", 50))
			showCategory()
			fmt.Scanln()
		case 4:
			return
		}
	}
}

func AdminMenu(userLog types.User) {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("-", "Halo "+userLog.Nama+" Silahkan Pilih menu dibawah", 50))
		fmt.Println("1. Category")
		fmt.Println("2. Products")
		fmt.Println("3. Log Out")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			CategoryMenu()
		case 2:
			ProductMenu()
		case 3:
			return
		}
	}
}
