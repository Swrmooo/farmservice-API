package bu

import (
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
	"farmservice/sqlstring"
)

func Plan_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Plan_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Plan_Create(trans db.Transaction, userId string) string {
	code := lib.GenerateRandomString(10)
	trans.Execute(sqlstring.Plan_Create(code, userId))
	if list := trans.Query(sqlstring.Plan_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	}else{
		panic("error.ContactAdmin")
	}
}

func Plan_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Plan_GetFromId(id)); len(list) == 1 {
		return list[0]
	}else{
		panic("error.ContactAdmin")
	}
}
