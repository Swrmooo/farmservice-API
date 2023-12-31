package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Vehicle_get() string {
	sql := " SELECT id, user_id, install, license_plate, brand, model, vehicle_type, driver, vehicle, catagory, code FROM vehicles "
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
	sql := " insert into vehicles (code, doc_date, user_id, install) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "', '" + "0" + "'); "
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

func Vehicle_DeleteFromId(id string) string {
	sql := "DELETE FROM vehicles"
	sql += " WHERE id IN (" + id + ")"
	return sql
}

func Vehicle_Count(userId string) string {
	sql := "SELECT COUNT(id) FROM vehicles WHERE user_id = '" + userId + "'"
	return sql
}
