package util

import (
	"farmservice/sqlstring"

	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func GenCode(prefix string, from string) string {
	code := ""
	sql := ""
	switch prefix {
	case "BR":
		code = "BR000001"
		filter := " left(code,2) = 'BR' and length(code) = 8 "
		sql = " select right(max(code),6) as last_code, count(code) as num from ( "
		switch from {
		case "driver":
			sql += sqlstring.Driver_GetFromFilter(filter)
			sql += " ) A "
		case "friend":
			sql += sqlstring.Friend_GetFromFilter(filter)
			sql += " ) A "
		case "risk":
			sql += sqlstring.Risk_GetFromFilter(filter)
			sql += " ) A "
		case "vehicle":
			sql += sqlstring.Vehicle_GetFromFilter(filter)
			sql += " ) A "
		}

	case "PT":
		code = "PT000001"
		filter := " left(code,2) = 'PT' and length(code) = 8 "
		sql = " select right(max(code),6) as last_code, count(code) as num from ( "
		sql += sqlstring.Plot_GetFromFilter(filter)
		sql += " ) A "
	case "SV":
		code = "SV000001"
		filter := " left(code,2) = 'SV' and length(code) = 8 "
		sql = " select right(max(code),6) as last_code, count(code) as num from ( "
		sql += sqlstring.Service_GetFromFilter(filter)
		sql += " ) A "
	case "JOB":
		code = "JOB000001"
		filter := " left(code,3) = 'JOB' and length(code) = 9 "
		sql := " select right(max(code),6) as last_code, count(code) as num from ( "
		sql += sqlstring.Plan_GetFromFilter(filter)
		sql += " ) A "
	default:
		panic("103")
	}

	if list := db.Query("fs", sql); len(list) == 1 {
		if lib.SI64(list[0], "num") > 0 {
			code = prefix + ZeroString(lib.SI64(list[0], "last_code")+1, 6)
		}
	} else {
		panic("error.ContactAdmin")
	}
	return code
}

func ZeroString(value int64, numberOfZero int) string {
	padStr := ""
	for k := 0; k < numberOfZero; k++ {
		padStr += "0"
	}
	thestr := padStr + lib.I64_S(value)
	return thestr[len(thestr)-numberOfZero:]
}
