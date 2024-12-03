package services

import (
	"errors"
	"toko_kucing/constants"
	"toko_kucing/database"
	"toko_kucing/types"
)

func AddProduct(productData types.Product) (bool, error) {
	var dataProd = database.ReadJSON[types.DataProduct](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	productData.Id = dataProd.LastId + 1
	dataProd.LastId = productData.Id
	dataProd.Data = append(dataProd.Data, productData)
	dataProd.Length = len(dataProd.Data)
	database.WriteJSON(
		constants.FullPathDB(constants.DB_USERS_FILE),
		dataProd,
	)
	return true, nil
}

func RemoveProductById(id int) (bool, error) {
	var dataProd = database.ReadJSON[types.DataProduct](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	for inx, data := range dataProd.Data {
		if data.Id == id {
			dataProd.Data = append(dataProd.Data[:inx], dataProd.Data[inx+1:]...)
			dataProd.Length--
			database.WriteJSON(
				constants.FullPathDB(constants.DB_CATEGORY_FILE),
				dataProd,
			)
			return true, nil
		}
	}
	return false, errors.New("ProductId Not Found")
}

func ModifyProductById(id int, productData types.Product) (bool, error) {
	var dataProd = database.ReadJSON[types.DataProduct](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	for inx, data := range dataProd.Data {
		if data.Id == id {
			dataProd.Data[inx] = productData
			database.WriteJSON(
				constants.FullPathDB(constants.DB_CATEGORY_FILE),
				dataProd,
			)
			return true, nil
		}
	}
	return false, errors.New("ProductId Not Found")
}
