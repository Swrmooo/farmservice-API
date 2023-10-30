package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Vehicle_get() string {
	//sql := " SELECT id, code, vehicle_type, doc_date name, detail FROM vehicles "
	sql := " SELECT id, user_id, license_plate, vehicle, num, driver, model, vehicle_type, brand FROM vehicles "
	sql += " WHERE "
	return sql
}

func Vehicle_GetFromUserId(id string) string {
	sql := Vehicle_get()
	sql += " id = '" + id + "' "
	return sql
}

func Vehicle_GetFromId(id string) string {
	sql := Vehicle_get()
	sql += " id = '" + id + "' "
	return sql
}

func Vehicle_GetFromCode(code string) string {
	sql := Vehicle_get()
	sql += " code = '" + code + "' "
	return sql
}

func Vehicle_GetFromFilter(filter string) string {
	sql := Vehicle_get()
	sql += filter
	return sql
}

func Vehicle_Create(code string, userId string) string {
	sql := " insert into vehicles (code, doc_date, user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "'); "
	return sql
}

func Vehicle_update(data map[string]interface{}) string {
	sql := " UPDATE vehicles set "
	for k, _ := range data {
		sql += " " + k + " = '" + lib.T(data, k) + "', "
	}
	sql += " last_updated_time = NOW() "
	sql += " WHERE "
	return sql
}

func Vehicle_UpdateFromId(id string, data map[string]interface{}) string {
	sql := Vehicle_update(data)
	sql += " id = '" + id + "' "
	return sql
}

// func Vehicle_DeleteFromId(id string) string {
// 	sql := " DELETE from vehicles "
// 	sql += " WHERE id = '" + id + "' "
// 	return sql
// }

func Vehicle_DeleteFromId(id string) string {
	sql := "DELETE FROM vehicles"
	sql += " WHERE id IN (" + id + ")"
	return sql
}
