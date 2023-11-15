package bu

import (
	"farmservice/sqlstring"
	"farmservice/util"
	"fmt"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Plot_GenCode() string {
	code := util.GenCode("PT", "")
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
	fmt.Println("code Plot=====", code)
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
