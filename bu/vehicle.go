package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Vehicle_GenCode() string {
	code := "BR000001"
	filter := " left(code,2) = 'BR' and length(code) = 8 "
	sql := " select right(max(code),6) as last_code, count(code) as num from ( "
	sql += sqlstring.Vehicle_GetFromFilter(filter)
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

func Vehicle_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Vehicle_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Vehicle_Create(trans db.Transaction, userId string) string {
	code := Vehicle_GenCode()
	trans.Execute(sqlstring.Vehicle_Create(code, userId))
	if list := trans.Query(sqlstring.Vehicle_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func Vehicle_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Vehicle_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
