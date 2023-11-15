package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Friend_GenCode() string {
	code := util.GenCode("BR", "friend")
	return code
}

func Friend_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Friend_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Friend_Create(trans db.Transaction, userId string, friendId string) string {
	code := Friend_GenCode()
	trans.Execute(sqlstring.Friend_Create(code, userId, friendId))
	if list := trans.Query(sqlstring.Friend_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func Friend_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Friend_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
