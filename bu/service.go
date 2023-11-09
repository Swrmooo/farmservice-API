package bu

import (
	"farmservice/sqlstring"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Service_GenCode() string {
	code := "SV000001"
	filter := " left(code,2) = 'SV' and length(code) = 8 "
	sql := " select right(max(code),6) as last_code, count(code) as num from ( "
	sql += sqlstring.Service_GetFromFilter(filter)
	sql += " ) A "
	if list := db.Query("fs", sql); len(list) == 1 {
		if lib.SI64(list[0], "num") > 0 {
			code = "SV" + lib.ZeroString(lib.SI64(list[0], "last_code")+1, 6)
		}
	} else {
		panic("error.ContactAdmin")
	}
	return code
}

func Service_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Service_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Service_Create(trans db.Transaction, userId string) string {
	//code := lib.GenerateRandomString(10)
	code := Service_GenCode()
	trans.Execute(sqlstring.Service_Create(code, userId))
	if list := trans.Query(sqlstring.Service_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func Service_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Service_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
