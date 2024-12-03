package constants

import "path/filepath"

const (
	PATH_DATA        = "data"
	PATH_DB          = "db"
	DB_USERS_FILE    = "user.json"
	DB_ORDERS_FILE   = "order.json"
	DB_PRODUCTS_FILE = "products.json"
	DB_CATEGORY_FILE = "category.json"
)

func FullPathDB(file string) string {
	return filepath.Join(PATH_DATA, PATH_DB, file)
}
