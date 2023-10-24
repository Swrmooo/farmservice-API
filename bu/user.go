package bu

import (
	//"farmservice/lib"
	"farmservice/lib/db"
	"farmservice/sqlstring"
)

// func User_Login(username string, pass string) map[string]interface{} {
//   if users := db.Query("fs", sqlstring.User_Login(username, pass)); len(users) == 1 {
//     return users[0]
//   }else {
//     return nil
//   }
// }

func User_Login(tel string, pass string) map[string]interface{} {
	users := db.Query("fs", sqlstring.User_Login(tel, pass))
	if len(users) == 1 {
		return users[0]
	} else {
		return nil
	}
}
