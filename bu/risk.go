package bu

import (
	"farmservice/sqlstring"
	"fmt"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Risk_List(filter string) []map[string]interface{} {
	list := db.Query("fs", sqlstring.Risk_GetFromFilter(filter))
	//for k,v := range list {
	//	do something
	//}
	return list
}

func Risk_Create(trans db.Transaction, userId string) string {
	code := lib.GenerateRandomString(10)
	fmt.Println("bu1")
	trans.Execute(sqlstring.Risk_Create(code, userId))
	fmt.Println("bu2")
	if list := trans.Query(sqlstring.Risk_GetFromCode(code)); len(list) == 1 {
		fmt.Println("bu3")
		return lib.T(list[0], "id")
	} else {
		fmt.Println("ererererere")
		panic("error.ContactAdmin")
	}
}

// func Risk_Create(trans db.Transaction, userId string) string {
// 	trans.Execute(sqlstring.Risk_Create(userId))

// 	return "moo success"
// }

func Risk_Detail(id string) map[string]interface{} {
	if list := db.Query("fs", sqlstring.Risk_GetFromId(id)); len(list) == 1 {
		return list[0]
	} else {
		panic("error.ContactAdmin")
	}
}
