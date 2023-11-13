package util

import (
	lib "github.com/ttoonn112/ktgolib"
)

// func GenCode() string {
// 	code := "BR000001"
// 	filter := " left(code,2) = 'BR' and length(code) = 8 "
// 	sql := " select right(max(code),6) as last_code, count(code) as num from ( "
// 	sql += sqlstring.Vehicle_GetFromFilter(filter)
// 	sql += " ) A "
// 	if list := db.Query("fs", sql); len(list) == 1 {
// 		if lib.SI64(list[0], "num") > 0 {
// 			code = "BR" + ZeroString(lib.SI64(list[0], "last_code")+1, 6)
// 		}
// 	} else {
// 		panic("error.ContactAdmin")
// 	}
// 	return code
// }

func ZeroString(value int64, numberOfZero int) string {
	padStr := ""
	for k := 0; k < numberOfZero; k++ {
		padStr += "0"
	}
	thestr := padStr + lib.I64_S(value)
	return thestr[len(thestr)-numberOfZero:]
}
