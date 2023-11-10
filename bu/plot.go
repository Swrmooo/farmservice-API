package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Plot_GenCode() string {
	code := "PT000001"
	//filter := " left(A.code,2) = 'PT' and length(A.code) = 8 "
	filter := " left(code,2) = 'PT' and length(code) = 8 "
	sql := " select right(max(code),6) as last_code, count(code) as num from ( "
	sql += sqlstring.Plot_GetFromFilter(filter)
	sql += " ) A "
	if list := db.Query("fs", sql); len(list) == 1 {
		if lib.SI64(list[0], "num") > 0 {
			code = "PT" + util.ZeroString(lib.SI64(list[0], "last_code")+1, 6)
		}
	} else {
		panic("error.ContactAdmin")
	}
	return code
}

func Plot_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Plot_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Plot_Create(trans db.Transaction, userId string) string {
	code := Plot_GenCode()
	// code := ZeroString("PT000001", 6)
	trans.Execute(sqlstring.Plot_Create(code, userId))
	if list := trans.Query(sqlstring.Plot_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func Plot_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Plot_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
