package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Service_GenCode() string {
	code := util.GenCode("SV", "")
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
