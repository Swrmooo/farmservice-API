package bu

import (
	"crypto/md5"
	"encoding/hex"
	"farmservice/sqlstring"
	"fmt"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func User_Login(tel string, pass string) map[string]interface{} {
	encodedPass := md5.Sum([]byte(pass))
	passMd5 := hex.EncodeToString(encodedPass[:])
	fmt.Println("Encoded Password: ", passMd5)

	if result := db.Query("fs", sqlstring.User_CheckLogin(tel, passMd5)); len(result) == 1 {
		id := lib.T(result[0], "id")
		token := lib.T(result[0], "username") + "-" + lib.GenerateRandomString(60) //Generate token
		db.Execute("fs", sqlstring.User_UpdateTokenFromId(id, token))
		detail := User_Detail(id)
		detail["token"] = token
		return detail
	}

	return nil
}

func User_Detail(id string) map[string]interface{} {
	if users := db.Query("fs", sqlstring.User_GetFromId(id)); len(users) == 1 {
		return users[0]
	} else {
		panic("error.ContactAdmin")
	}
	return nil
}

func User_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.User_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func User_Create(trans db.Transaction, tel string, pass string) string {
	fmt.Println("passwordBu=====", pass)
	if users := trans.Query(sqlstring.User_GetFromPhone(tel)); len(users) > 0 {
		// panic("error.user.PhoneExists")
		trans.Execute(sqlstring.User_CreateWithPhone(tel, pass))
	}
	trans.Execute(sqlstring.User_CreateWithPhone(tel, pass))
	if users := trans.Query(sqlstring.User_GetFromPhone(tel)); len(users) == 1 {
		return lib.T(users[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func User_Register(trans db.Transaction, tel string, pass string) string {
	if users := trans.Query(sqlstring.User_GetFromPhone(tel)); len(users) == 0 {
		// panic("error.user.PhoneExists")
		trans.Execute(sqlstring.User_CreateWithPhone(tel, pass))
	}
	//trans.Execute(sqlstring.User_CreateWithPhone(tel, pass))
	if users := trans.Query(sqlstring.User_GetFromPhone(tel)); len(users) == 1 {
		return lib.T(users[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}
