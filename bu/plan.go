package bu

import (
	"farmservice/sqlstring"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Plan_GenCode() string {
	code := "JOB000001"
	filter := " left(code,3) = 'JOB' and length(code) = 9 "
	sql := " select right(max(code),6) as last_code, count(code) as num from ( "
	sql += sqlstring.Plan_GetFromFilter(filter)
	sql += " ) A "
	if list := db.Query("fs", sql); len(list) == 1 {
		if lib.SI64(list[0], "num") > 0 {
			code = "JOB" + lib.ZeroString(lib.SI64(list[0], "last_code")+1, 6)
		}
	} else {
		panic("error.ContactAdmin")
	}
	return code
}

func Plan_Join(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Plan_GetFromFilterToJoin(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Plan_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Plan_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Plan_Create(trans db.Transaction, userId string) string {
	//code := lib.GenerateRandomString(10)
	code := Plan_GenCode()
	trans.Execute(sqlstring.Plan_Create(code, userId))
	if list := trans.Query(sqlstring.Plan_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func Plan_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Plan_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
