package services

import (
	"errors"
	"toko_kucing/constants"
	"toko_kucing/database"
	"toko_kucing/types"
)

func AddCategory(catData types.Category) (bool, error) {
	var dataCat = database.ReadJSON[types.DataCategory](constants.FullPathDB(constants.DB_CATEGORY_FILE))
	for _, data := range dataCat.Data {
		if data.Nama == catData.Nama {
			return false, errors.New("Duplicate Name")
		}
	}
	catData.Id = dataCat.LastId + 1
	dataCat.LastId = catData.Id
	dataCat.Data = append(dataCat.Data, catData)
	dataCat.Length = len(dataCat.Data)
	database.WriteJSON(
		constants.FullPathDB(constants.DB_CATEGORY_FILE),
		dataCat,
	)
	return true, nil
}

func RemoveCategoryById(id int) (bool, error) {
	var dataCat = database.ReadJSON[types.DataCategory](constants.FullPathDB(constants.DB_CATEGORY_FILE))
	for inx, data := range dataCat.Data {
		if data.Id == id {
			dataCat.Data = append(dataCat.Data[:inx], dataCat.Data[inx+1:]...)
			dataCat.Length--
			database.WriteJSON(
				constants.FullPathDB(constants.DB_CATEGORY_FILE),
				dataCat,
			)
			return true, nil
		}
	}
	return false, errors.New("CategoryId Not Found")
}

func GetCategoryById(id int) (bool, string) {
	var dataCat = database.ReadJSON[types.DataCategory](constants.FullPathDB(constants.DB_CATEGORY_FILE))
	for _, data := range dataCat.Data {
		if data.Id == id {
			return true, data.Nama
		}
	}
	return false, ""
}

func ListCategory() types.DataCategory {
	return database.ReadJSON[types.DataCategory](constants.FullPathDB(constants.DB_CATEGORY_FILE))
}
