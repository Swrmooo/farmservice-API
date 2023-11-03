package sqlstring

import (
	"crypto/md5"
	"encoding/hex"

	lib "github.com/ttoonn112/ktgolib"
)

func User_CheckLogin(username string, pass string) string {
	sql := " SELECT id, username FROM users "
	sql += " WHERE (username = '" + username + "' OR tel = '" + username + "') AND password = '" + pass + "' "
	return sql
}

func User_UpdateTokenFromId(id string, token string) string {
	sql := " UPDATE users set token = '" + token + "', token_expire_time = NOW() + INTERVAL 1 DAY " // กำหนดให้ Access Token มีอายุ 1 วันนับจากเวลาที่ Access ล่าสุด
	sql += " WHERE id = '" + id + "' "
	return sql
}

func User_UpdateTokenTime(token string) string {
	sql := " UPDATE users set token_expire_time = NOW() + INTERVAL 1 DAY " // กำหนดให้ Access Token มีอายุ 1 วันนับจากเวลาที่ Access ล่าสุด
	sql += " WHERE token = '" + token + "' "
	return sql
}

func user_get() string {
	sql := " SELECT id, username, tel, firstname, lastname, password, member FROM users "
	sql += " WHERE "
	return sql
}

func User_GetFromToken(token string) string {
	sql := user_get()
	sql += " token = '" + token + "' and token_expire_time >= NOW() "
	return sql
}

func User_GetFromId(id string) string {
	sql := user_get()
	sql += " id = '" + id + "' "
	return sql
}

func User_GetFromPhone(tel string) string {
	sql := user_get()
	sql += " tel = '" + tel + "' "
	return sql
}

func User_GetFromFilter(filter string) string {
	sql := user_get()
	sql += filter
	return sql
}

func User_CreateWithPhone(tel string) string {
	member := "standard"
	encodedPass := md5.Sum([]byte(tel))
	passMd5 := hex.EncodeToString(encodedPass[:])
	// fmt.Println("-------role : ", member)
	// fmt.Println("-------passMd5 : ", passMd5)

	// sql := " insert into users (tel, username, password) values ('"+tel+"', '"+tel+"', '"+lib.GenerateRandomString(10)+"'); "
	sql := " insert into users (tel, username, password, member) values ('" + tel + "', '" + tel + "', '" + passMd5 + "', '" + member + "'); "
	return sql
}

func user_update(data map[string]interface{}) string {
	sql := " UPDATE users set "
	for k, _ := range data {
		sql += " " + k + " = '" + lib.T(data, k) + "', "
	}
	sql += " last_updated_time = NOW() "
	sql += " WHERE "
	return sql
}

func User_UpdateFromId(id string, data map[string]interface{}) string {
	sql := user_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func User_DeleteFromId(id string) string {
	sql := "DELETE FROM users"
	sql += " WHERE id IN (" + id + ")"
	return sql
}
