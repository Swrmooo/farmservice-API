package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Risk_GenCode() string {
	code := util.GenCode("BR", "risk")
	return code
}

func Risk_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Risk_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Risk_Create(trans db.Transaction, userId string) string {
	code := Risk_GenCode()
	// code := code.GenCode("SW000001")
	trans.Execute(sqlstring.Risk_Create(code, userId))
	if list := trans.Query(sqlstring.Risk_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	} else {
		panic("error.ContactAdmin")
	}
}

func Risk_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Risk_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
