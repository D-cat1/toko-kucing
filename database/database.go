package database

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"toko_kucing/constants"
	"toko_kucing/types"
)

type ConstarintData interface {
	types.DataCategory | types.DataDBInit | types.DataProduct | types.DataUser
}

func initFolder() {
	paths := []string{constants.PATH_DATA, filepath.Join(constants.PATH_DATA, constants.PATH_DB)}
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}
	}
}

func ReadJSON[V ConstarintData](path string) V {
	data, _ := os.ReadFile(path)
	var dataUser V
	err := json.Unmarshal(data, &dataUser)
	if err != nil {
		fmt.Println("error JSON not in valid format")
	}
	return dataUser
}

func WriteJSON[V ConstarintData](path string, dataUser V) {
	dataJson, _ := json.Marshal(dataUser)
	os.WriteFile(path, dataJson, os.ModePerm)
}

func initFileDBIfNotExist() {
	var rootDBPath string = filepath.Join(constants.PATH_DATA, constants.PATH_DB)
	paths := []string{
		constants.DB_USERS_FILE,
		constants.DB_PRODUCTS_FILE,
		constants.DB_ORDERS_FILE,
		constants.DB_CATEGORY_FILE,
	}
	var dataDummy = types.DataDBInit{LastId: 0, Length: 0, Data: make([]int, 0)}
	for _, path := range paths {
		if _, err := os.Stat(filepath.Join(rootDBPath, path)); os.IsNotExist(err) {
			WriteJSON(filepath.Join(rootDBPath, path), dataDummy)
		}
	}
}

func InitDb() {
	initFolder()
	initFileDBIfNotExist()
}
