package view

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"tri_darma/services"
	"tri_darma/types"

	"github.com/olekukonko/tablewriter"
)

func inputRegTriDarma(data *types.TriDarma, data_type string) {
	Clrscr()
	fmt.Printf(border("-", "Detail %v", 50)+"\n", data_type)
	fmt.Printf("Judul %v : ", data_type)
	HandleLongInput(&data.Nama)
	data.Tipe = data_type
	fmt.Printf("Prodi %v : ", data_type)
	HandleLongInput(&data.Prodi)
	fmt.Printf("Tahun %v : ", data_type)
	fmt.Scan(&data.Tahun)
}

func handleRegister(state *types.TriDarma) int {
	var choice int
	var dataTriDarma types.TriDarma
	for {
		Clrscr()
		fmt.Println(border("-", "Register", 50))
		fmt.Println("1. Penelitian")
		fmt.Println("2. Abdimas")
		fmt.Println("0. Exit")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			inputRegTriDarma(&dataTriDarma, "Penelitian")
		case 2:
			inputRegTriDarma(&dataTriDarma, "Abdimas")
		case 0:
			return 0
		default:
			fmt.Println("Opsi tidak tersedia")
			delay(1)
		}

		services.Add3Darma(dataTriDarma)

		var lastidx = services.ListTridar().Length
		*state = services.ListTridar().Data[lastidx-1]
		return 1
	}
}

func ShowTriDarma(dataTriSaved *types.TriDarma) int {
	var choice int
	var lastidx = services.ListTridar()
	Clrscr()
	fmt.Println(border("-", "Pilih Judul Penelitian", 50))
	for i := 0; i < lastidx.Length; i++ {
		fmt.Printf("%v. %v\n", i+1, lastidx.Data[i].Nama)
	}
	fmt.Println("0. Exit")
	fmt.Println(border("-", "", 50))

	for choice > lastidx.Length || choice == 0 {
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
		if choice == 0 {
			return 0
		}
		fmt.Println("Inputan Salah! ulangi")
	}
	*dataTriSaved = lastidx.Data[choice-1]
	return 1
}

func PrintTriDarmaTable(data *[]types.TriDarma) {
	dataLength := len(*data)
	Clrscr()
	fmt.Println(border("-", "List Tri Darma", 50))
	if dataLength < 1 {
		fmt.Println("Tidak ada data.")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Judul", "Prodi", "Tahun"})
		for i := 0; i < dataLength; i++ {
			table.Append([]string{strconv.Itoa(i + 1), (*data)[i].Nama, (*data)[i].Prodi, strconv.Itoa((*data)[i].Tahun)})
		}
		table.Render()
	}

	fmt.Println("Klik [enter] untuk lanjut")
	fmt.Scanln()
}

func SearchTahun() {
	var tahun int
	triDarma := services.ListTridar()
	Clrscr()
	fmt.Println(border("-", "Cari Tri Darma", 50))
	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)

	temp := []types.TriDarma{}

	found := false
	left := 0
	right := triDarma.Length - 1
	println(left, right, tahun)
	for left <= right && !found {
		mid := (left + right) / 2
		if triDarma.Data[mid].Tahun == tahun {
			temp = append(temp, triDarma.Data[mid])
			// search after
			for i := mid + 1; i < triDarma.Length && triDarma.Data[i].Tahun == tahun; i++ {
				temp = append(temp, triDarma.Data[i])
			}
			// search before
			for i := mid - 1; i >= 0 && triDarma.Data[i].Tahun == tahun; i-- {
				temp = append(temp, triDarma.Data[i])
			}
			found = true
		} else if triDarma.Data[mid].Tahun > tahun {
			left = mid + 1
		} else {
			right = mid - 1
		}
		println(left, right, mid, triDarma.Data[mid].Tahun, tahun)
	}
	fmt.Scanln()
	PrintTriDarmaTable(&temp)
}

// func BinarySearchTahun(triDarma *types.DataTriDarma, tahun int, getFirstIndex bool) int {
// 	left := 0
// 	right := triDarma.Length - 1
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if triDarma.Data[mid].Tahun < tahun {
// 			left = mid + 1
// 		} else if triDarma.Data[mid].Tahun > tahun {
// 			right = mid - 1
// 		} else {
// 			if getFirstIndex {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 	}
// 	return -1
// }

// func SearchTahun() {
// 	var tahun int
// 	triDarma := services.ListTridar()
// 	Clrscr()
// 	fmt.Println(border("-", "Cari Tri Darma", 50))
// 	fmt.Print("Tahun : ")
// 	fmt.Scan(&tahun)

// 	left, right := BinarySearchTahun(&triDarma, tahun, true), BinarySearchTahun(&triDarma, tahun, false)
// 	if left == -1 || right == -1 {
// 		fmt.Println(border("-", "List Tri Darma", 50))
// 		fmt.Println("Tidak ada data.")
// 		fmt.Scanln()
// 		fmt.Println("\nKlik [enter] untuk lanjut")
// 		fmt.Scanln()
// 	} else {
// 		fmt.Scanln()
// 		temp := triDarma.Data[left : right+1]
// 		PrintTriDarmaTable(&temp)
// 	}

// }

func SearchProdi() {
	var prodi string
	triDarma := services.ListTridar()
	Clrscr()
	fmt.Println(border("-", "Cari Tri Darma", 50))
	fmt.Print("Prodi : ")
	HandleLongInput(&prodi)
	prodi = strings.ToLower(prodi)

	var temp []types.TriDarma
	for i := 0; i < triDarma.Length; i++ {
		if strings.ToLower(triDarma.Data[i].Prodi) == prodi {
			temp = append(temp, triDarma.Data[i])
		}
	}
	PrintTriDarmaTable(&temp)
}

func InsertionSort() {
	data := services.ListTridar()
	for i := 1; i < data.Length; i++ {
		temp := data.Data[i]
		j := i - 1
		// desc <, asc >
		for j >= 0 && data.Data[j].Tahun < temp.Tahun {
			data.Data[j+1] = data.Data[j]
			j--
		}
		data.Data[j+1] = temp
	}
	services.UpdateAll(data)
}

type YearRank struct {
	Year  int
	Count int
}

func SelectionSort() {
	data := services.ListTridar()
	tahun := [100]YearRank{{data.Data[0].Tahun, 1}}
	n := 1

	for i := 1; i < data.Length; i++ {
		exist := false
		for j := 0; j < n && !exist; j++ {
			if data.Data[i].Tahun == tahun[j].Year {
				tahun[j].Count++
				exist = true
			}
		}
		if !exist {
			tahun[n] = YearRank{data.Data[i].Tahun, 1}
			n++
		}
	}

	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if tahun[j].Count > tahun[maxIdx].Count {
				maxIdx = j
			}
		}
		tahun[i], tahun[maxIdx] = tahun[maxIdx], tahun[i]
	}

	var temp [100]types.TriDarma
	tempCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < data.Length; j++ {
			if data.Data[j].Tahun == tahun[i].Year {
				temp[tempCount] = data.Data[j]
				tempCount++
			}
		}
	}
	data.Data = temp
	services.UpdateAll(data)
}

func HandleManagement(dataTriSaved *types.TriDarma) int {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("-", "Manage Tri Darma", 50))
		fmt.Println("1. List & Detail Tri Darma")
		fmt.Println("2. Cari Tri Darma (Tahun)") // maybe just one is fine
		fmt.Println("3. Cari Tri Darma (Prodi)")
		fmt.Println("4. Sort by Tahun")
		fmt.Println("5. Sort by Jumlah Kegiatan Per Tahun")
		fmt.Println("0. Exit")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if ShowTriDarma(dataTriSaved) == 1 {
				return 1
			}
		case 2:
			InsertionSort()
			SearchTahun()
		case 3:
			SearchProdi()
		case 4:
			InsertionSort()
			fmt.Println("Data tersortir. Kembali ke menu...")
			delay(1)
		case 5:
			SelectionSort()
		case 0:
			return 0
		default:
			fmt.Println("Opsi tidak tersedia")
			delay(1)
		}
	}
}

func MainMenu(managedState *types.TriDarma) {
	var choice int
	for {
		Clrscr()
		fmt.Println(border("-", "Silahkan Pilih menu dibawah", 50))
		fmt.Println("1. Registrasi Tri Darma")
		fmt.Println("2. Manage Tri Darma")
		fmt.Println("0. Exit")
		fmt.Println(border("-", "", 50))
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if handleRegister(managedState) == 1 {
				return
			}
		case 2:
			if services.ListTridar().Length < 1 {
				fmt.Println("Data Tri Darma masih kosong!")
				delay(1)
				fmt.Println("Kembali ke menu awal...")
				delay(2)
			} else {
				if HandleManagement(managedState) == 1 {
					return
				}
			}
		case 0:
			fmt.Println("babay!")
			return
		default:
			fmt.Println("Opsi tidak tersedia")
			delay(1)
		}
	}
}
