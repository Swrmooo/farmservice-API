package bu

import (
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
	"farmservice/sqlstring"
)

func Driver_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Driver_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Driver_Create(trans db.Transaction, userId string) string {
	code := lib.GenerateRandomString(10)
	trans.Execute(sqlstring.Driver_Create(code, userId))
	if list := trans.Query(sqlstring.Driver_GetFromCode(code)); len(list) == 1 {
		return lib.T(list[0], "id")
	}else{
		panic("error.ContactAdmin")
	}
}

func Driver_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Driver_GetFromId(id)); len(list) == 1 {
		return list[0]
	}else{
		panic("error.ContactAdmin")
	}
}
