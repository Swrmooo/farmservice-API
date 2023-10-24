package sqlstring

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func User_GetToken(token string) string {
	sql := " SELECT tel FROM btfarmservice_db.users "
	sql += " WHERE token = '" + token + "' "

	return sql
}

func User_Login(tel string, pass string) string {
	// encodepass := " ";
	// decodepass := " ";
	encodedPass := md5.Sum([]byte(pass))
	encodedPassStr := hex.EncodeToString(encodedPass[:])

	fmt.Println("Encoded Password: ", encodedPassStr)
	sql := " SELECT tel FROM btfarmservice_db.users "
	sql += " WHERE tel = '" + tel + "' AND password = ('" + pass + "') "

	return sql
}
