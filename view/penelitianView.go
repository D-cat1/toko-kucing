package view

import (
	"fmt"
	"os"
	"strconv"
	services "tri_darma/services"
	types "tri_darma/types"

	tablewriter "github.com/olekukonko/tablewriter"
)

func choosePeran() string {
	var choice int
	Clrscr()
	fmt.Println(border("-", "Pilih Jabatan", 50))
	fmt.Println("1. Anggota (Mahasiswa)")
	fmt.Println("2. Anggota (Dosen)")
	fmt.Println(border("-", "", 50))
	for {
		fmt.Print("Pilih : ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			return "Anggota (Mahasiswa)"
		case 2:
			return "Anggota (Dosen)"
		default:
			fmt.Println("inputan salah! Ulangi")
			delay(2)
		}
	}
}

func PenelitianMenu(penelitianManage *types.TriDarma) {
	var choice int
	var anggota types.Anggota
	for {
		Clrscr()
		if penelitianManage.CountAnggota == 0 {
			fmt.Println(border("-", "Data Ketua", 50))
			fmt.Print("Nama Ketua : ")
			HandleLongInput(&anggota.Nama)
			anggota.IdTridarma = penelitianManage.Id
			anggota.Role = "Ketua"
			services.AddAnggota(anggota)
			_, user := services.GetTridarById(penelitianManage.Id)
			*penelitianManage = user
		} else {
			Clrscr()
			_, user := services.GetTridarById(penelitianManage.Id)
			*penelitianManage = user
			fmt.Println(border("-", "Detail "+penelitianManage.Tipe, 50))
			formatPrint("Nama", penelitianManage.Nama)
			formatPrint("Tahun", penelitianManage.Tahun)
			formatPrint("Banyak Anggota", penelitianManage.CountAnggota)
			formatPrint("Banyak Luaran", penelitianManage.CountLuaran)
			formatPrint("Total Pendanaan", penelitianManage.SumDana)
			fmt.Println(border("-", "", 50))
			fmt.Println("1. Tambah Anggota")
			fmt.Println("2. Hapus Anggota")
			fmt.Println("3. Lihat Anggota")
			fmt.Println("4. Tambah Pendanaan")
			fmt.Println("5. Lihat Pendanaan")
			fmt.Println("6. Hapus Pendanaan")
			fmt.Println("7. Tambah Luaran")
			fmt.Println("8. Lihat Luaran")
			fmt.Println("9. Hapus Luaran")
			fmt.Println("0. Exit")
			fmt.Println(border("-", "", 50))
			fmt.Print("Pilih : ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				if penelitianManage.CountAnggota < 4 {
					Clrscr()
					fmt.Println(border("-", "Data Anggota", 50))
					var anggotaTemp types.Anggota
					fmt.Print("Nama : ")
					HandleLongInput(&anggotaTemp.Nama)
					anggotaTemp.Role = choosePeran()
					anggotaTemp.IdTridarma = penelitianManage.Id
					services.AddAnggota(anggotaTemp)
					fmt.Println("Data berhasil disimpan, akan dialihkan dalam 2 dtk")
				} else {
					fmt.Println("Error!, Max. 4 Orang")
				}
				delay(2)
			case 2:
				var sementara [99]int
				var countTmp int
				var dataAnggota = services.ListAnggota()
				for i := 0; i < dataAnggota.Length; i++ {
					if dataAnggota.Data[i].IdTridarma == penelitianManage.Id {
						sementara[countTmp] = i
						countTmp++
					}
				}
				if countTmp == 0 {
					fmt.Println("Tidak Ada Anggota disini!")
					delay(2)
				} else {
					Clrscr()
					fmt.Println(border("-", "Hapus Anggota", 50))
					for i := 0; i < countTmp; i++ {
						fmt.Println(i+1, ". ", dataAnggota.Data[sementara[i]].Nama)
					}
					fmt.Println(border("-", "", 50))
					fmt.Print("Pilih : ")
					fmt.Scanln(&choice)
					for choice > countTmp || choice < 1 {
						fmt.Println("Pilihan tidak valid!, ulang!")
						fmt.Print("Pilih : ")
						fmt.Scanln(&choice)
					}

					services.RemoveAnggotaById(dataAnggota.Data[sementara[choice-1]].Id)

					fmt.Println("Data Berhasil dihapus!, dialihkan dalam 2 detik")
					delay(2)
				}
			case 3:
				Clrscr()
				fmt.Println(border("-", "Data Anggota", 50))
				var dataAnggota = services.ListAnggota()
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"No", "Nama", "Role"})
				for i := 0; i < dataAnggota.Length; i++ {
					if dataAnggota.Data[i].IdTridarma == penelitianManage.Id {
						table.Append([]string{strconv.Itoa(i + 1), dataAnggota.Data[i].Nama, dataAnggota.Data[i].Role})
					}
				}
				table.Render()

				fmt.Println("Klik [enter] untuk lanjut")
				fmt.Scanln()
			case 4:
				Clrscr()
				var dataPay types.Dana
				fmt.Println(border("-", "Tambah Pendanaan", 50))
				fmt.Println("1. Internal")
				fmt.Println("2. External")
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scanln(&choice)
				for choice > 2 || choice < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scanln(&choice)
				}
				if choice == 1 {
					dataPay.Sumber = "Internal"
				} else {
					dataPay.Sumber = "External"
				}
				fmt.Print("Keterangan : ")
				HandleLongInput(&dataPay.Keterangan)
				fmt.Print("Nominal : ")
				fmt.Scanln(&dataPay.Nominal)
				dataPay.IdTridarma = penelitianManage.Id
				services.AddDana(dataPay)
				fmt.Println("Data berhasil disimpan, akan dialihkan dalam 2 dtk")
				delay(2)
			case 5:
				Clrscr()
				fmt.Println(border("-", "Data Pendanaan", 50))
				var dataDana = services.ListDana()
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"No", "Sumber Dana", "Asal Dana", "Nomimal"})
				for i := 0; i < dataDana.Length; i++ {
					if dataDana.Data[i].IdTridarma == penelitianManage.Id {
						table.Append([]string{strconv.Itoa(i + 1), dataDana.Data[i].Sumber, dataDana.Data[i].Keterangan, strconv.Itoa(dataDana.Data[i].Nominal)})
					}
				}
				table.Render()

				fmt.Println("Klik [enter] untuk lanjut")
				fmt.Scanln()
			case 6:
				var sementara [99]int
				var countTmp int
				var dataDana = services.ListDana()
				for i := 0; i < dataDana.Length; i++ {
					if dataDana.Data[i].IdTridarma == penelitianManage.Id {
						sementara[countTmp] = i
						countTmp++
					}
				}
				if countTmp == 0 {
					fmt.Println("Tidak Ada Dana, Kembali ke menu!")
					delay(2)
				} else {
					Clrscr()
					fmt.Println(border("-", "Hapus Pendanaan", 50))
					for i := 0; i < countTmp; i++ {
						fmt.Println(i+1, ". ", dataDana.Data[sementara[i]].Sumber, "(", dataDana.Data[sementara[i]].Keterangan, ")")
					}
					fmt.Println(border("-", "", 50))
					fmt.Print("Pilih : ")
					fmt.Scanln(&choice)
					for choice > countTmp || choice < 1 {
						fmt.Println("Pilihan tidak valid!, ulang!")
						fmt.Print("Pilih : ")
						fmt.Scanln(&choice)
					}
					services.RemoveDanaById(dataDana.Data[sementara[choice-1]].Id)
					fmt.Println("Data Berhasil dihapus!, dialihkan dalam 2 detik")
					delay(2)
				}
			case 7:
				Clrscr()
				var tempLuaran types.Luaran
				fmt.Println(border("-", "Tambah Luaran", 50))
				fmt.Println("1. Publikasi")
				fmt.Println("2. Produk")
				if penelitianManage.Tipe == "Abdimas" {
					fmt.Println("3. Seminar")
					fmt.Println("4. Pelatihan")
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scanln(&choice)
				for (choice > 2 && penelitianManage.Tipe == "Penelitian") || (choice > 4 && penelitianManage.Tipe == "Abdimas") {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scanln(&choice)
				}
				switch choice {
				case 1:
					tempLuaran.BentukLuaran = "Publikasi"
				case 2:
					tempLuaran.BentukLuaran = "Produk"
				case 3:
					tempLuaran.BentukLuaran = "Seminar"
				case 4:
					tempLuaran.BentukLuaran = "Pelatihan"
				}
				fmt.Print("Tanggal Pelaksanaan (dd/mm/yyyy) : ")
				HandleLongInput(&tempLuaran.Pelaksanaan)
				tempLuaran.IdTridarma = penelitianManage.Id
				services.AddLuaran(tempLuaran)
				fmt.Println("Data berhasil disimpan, akan dialihkan dalam 2 dtk")
				delay(2)
			case 8:
				Clrscr()
				fmt.Println(border("-", "Data Pendanaan", 50))
				var dataLuaran = services.ListLuaran()
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"No", "Bentuk", "Pelaksanaan (dd/mm/yyyy)"})
				for i := 0; i < dataLuaran.Length; i++ {
					if dataLuaran.Data[i].IdTridarma == penelitianManage.Id {
						table.Append([]string{strconv.Itoa(i + 1), dataLuaran.Data[i].BentukLuaran, dataLuaran.Data[i].Pelaksanaan})
					}
				}
				fmt.Println("Klik [enter] untuk lanjut")
				fmt.Scanln()
			case 9:
				var sementara [99]int
				var countTmp int
				var dataLuaran = services.ListLuaran()
				for i := 0; i < dataLuaran.Length; i++ {
					if dataLuaran.Data[i].IdTridarma == penelitianManage.Id {
						sementara[countTmp] = i
						countTmp++
					}
				}
				if countTmp == 0 {
					fmt.Println("Tidak Ada Luaran, Kembali ke menu!")
					delay(2)
				} else {
					Clrscr()
					fmt.Println(border("-", "Hapus Pendanaan", 50))
					for i := 0; i < countTmp; i++ {
						fmt.Println(i+1, ". ", dataLuaran.Data[sementara[i]].BentukLuaran, "(", dataLuaran.Data[sementara[i]].Pelaksanaan, ")")
					}
					fmt.Println(border("-", "", 50))
					fmt.Print("Pilih : ")
					fmt.Scanln(&choice)
					for choice > countTmp || choice < 1 {
						fmt.Println("Pilihan tidak valid!, ulang!")
						fmt.Print("Pilih : ")
						fmt.Scanln(&choice)
					}
					services.RemoveProductById(dataLuaran.Data[sementara[choice-1]].Id)
					fmt.Println("Data Berhasil dihapus!, dialihkan dalam 2 detik")
					delay(2)
				}
			case 0:
				return
			}
		}

	}
}
