package sqlstring

import (
	"fmt"

	lib "github.com/ttoonn112/ktgolib"
)

func Driver_get() string {
	//sql := " SELECT id, code, driver_type, doc_date name, detail FROM drivers "
	sql := " SELECT id, user_id, firstname, lastname, tel, mood, pics FROM drivers "
	sql += " WHERE "
	return sql
}

func Driver_GetFromId(id string) string {
	sql := Driver_get()
	sql += " id = '" + id + "' "
	return sql
}

func Driver_GetFromCode(code string) string {
	sql := Driver_get()
	sql += " code = '" + code + "' "
	return sql
}

func Driver_GetFromFilter(filter string) string {
	sql := Driver_get()
	sql += filter
	fmt.Println("/-/-/-/-/-/-/-/-/-/-/test", filter)
	fmt.Println("/-/-/-/-/-/-/-/-/-/-/test", sql)
	return sql
}

func Driver_Create(code string, userId string) string {
	sql := " insert into drivers (code, doc_date, user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "'); "
	return sql
}

func Driver_update(data map[string]interface{}) string {
	sql := " UPDATE drivers set "
	for k, _ := range data {
		sql += " " + k + " = '" + lib.T(data, k) + "', "
	}
	sql += " last_updated_time = NOW() "
	sql += " WHERE "
	return sql
}

func Driver_UpdateFromId(id string, data map[string]interface{}) string {
	sql := Driver_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func Driver_DeleteFromId(id string) string {
	sql := "DELETE FROM drivers"
	sql += " WHERE id IN (" + id + ")"
	return sql
}

func Driver_Count(userId string) string {
	sql := "SELECT COUNT(id) FROM drivers WHERE user_id = '" + userId + "'"
	return sql
}
