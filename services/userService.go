package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"toko_kucing/constants"
	"toko_kucing/database"
	"toko_kucing/types"
)

func encryptPass(pwd string) string {
	var sha = sha1.New()
	sha.Write([]byte(pwd))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	return encryptedString
}

func AddUser(userData types.User) (bool, error) {
	var dataUser = database.ReadJSON[types.DataUser](constants.FullPathDB(constants.DB_USERS_FILE))
	for _, data := range dataUser.Data {
		if data.Username == userData.Username {
			return false, errors.New("Duplicate Username")
		}
	}
	userData.Id = dataUser.LastId + 1
	userData.Password = encryptPass(userData.Password)
	dataUser.LastId = userData.Id
	dataUser.Data = append(dataUser.Data, userData)
	dataUser.Length = len(dataUser.Data)
	database.WriteJSON(
		constants.FullPathDB(constants.DB_USERS_FILE),
		dataUser,
	)
	return true, nil
}

func RemoveUserByUsername(username string) (bool, error) {
	var dataUser = database.ReadJSON[types.DataUser](constants.FullPathDB(constants.DB_USERS_FILE))
	for inx, data := range dataUser.Data {
		if data.Username == username {
			dataUser.Data = append(dataUser.Data[:inx], dataUser.Data[inx+1:]...)
			dataUser.Length--
			database.WriteJSON(
				constants.FullPathDB(constants.DB_USERS_FILE),
				dataUser,
			)
			return true, nil
		}
	}
	return false, errors.New("User Not Found")
}

func ChangePwdByUsername(username, newPassword string) (bool, error) {
	var dataUser = database.ReadJSON[types.DataUser](constants.FullPathDB(constants.DB_USERS_FILE))
	for inx, data := range dataUser.Data {
		if data.Username == username {
			dataUser.Data[inx].Password = newPassword
			database.WriteJSON(
				constants.FullPathDB(constants.DB_USERS_FILE),
				dataUser,
			)
			return true, nil
		}
	}
	return false, errors.New("User Not Found")
}

func ListUser() types.DataUser {
	return database.ReadJSON[types.DataUser](constants.FullPathDB(constants.DB_USERS_FILE))
}

func GetUserByUsername(username string) *types.User {
	var dataUser = database.ReadJSON[types.DataUser](constants.FullPathDB(constants.DB_USERS_FILE))
	for _, data := range dataUser.Data {
		if data.Username == username {
			return &data
		}
	}
	return nil
}

func LoginAction(username string, password string) *types.User {
	var dataUser = GetUserByUsername(username)
	if dataUser != nil && dataUser.Password == encryptPass(password) {
		return dataUser
	}
	return nil
}
