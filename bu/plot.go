package bu

import (
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
	"farmservice/sqlstring"
)

func Plot_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Plot_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Plot_Create(trans *db.Transaction, userId string) string {
	code := lib.GenerateRandomString(10)
	trans.Execute(sqlstring.Plot_Create(code, userId))
	if list := trans.Query(sqlstring.Plot_GetFromCode(code)); len(list) == 1 {
		return lib.T(users[0], "id")
	}else{
		panic("error.ContactAdmin")
	}
	return nil
}

func Plot_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Plot_GetFromId(id)); len(list) == 1 {
		return list[0]
	}else{
		panic("error.ContactAdmin")
	}
	return nil
}
