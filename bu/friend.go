package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Friend_GenCode() string {
	code := "BR000001"
	filter := " left(code,2) = 'BR' and length(code) = 8 "
	sql := " select right(max(code),6) as last_code, count(code) as num from ( "
	sql += sqlstring.Friend_GetFromFilter(filter)
	sql += " ) A "
	if list := db.Query("fs", sql); len(list) == 1 {
		if lib.SI64(list[0], "num") > 0 {
			code = "BR" + util.ZeroString(lib.SI64(list[0], "last_code")+1, 6)
		}
	} else {
		panic("error.ContactAdmin")
	}
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
