package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Service_get() string {
	//sql := " SELECT id, code, service_type, doc_date name, detail FROM services "
	sql := " SELECT id, user_id, service_type, service, fee, detail FROM services "
	sql += " WHERE "
	return sql
}

func Service_GetFromUserId(id string) string {
	sql := Service_get()
	sql += " id = '" + id + "' "
	return sql
}

func Service_GetFromId(id string) string {
	sql := Service_get()
	sql += " id = '" + id + "' "
	return sql
}

func Service_GetFromCode(code string) string {
	sql := Service_get()
	sql += " code = '" + code + "' "
	return sql
}

func Service_GetFromFilter(filter string) string {
	sql := Service_get()
	sql += filter
	return sql
}

func Service_Create(code string, userId string) string {
	sql := " insert into services (code, doc_date, user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "'); "
	return sql
}

func Service_update(data map[string]interface{}) string {
	sql := " UPDATE services set "
	for k, _ := range data {
		sql += " " + k + " = '" + lib.T(data, k) + "', "
	}
	sql += " last_updated_time = NOW() "
	sql += " WHERE "
	return sql
}

func Service_UpdateFromId(id string, data map[string]interface{}) string {
	sql := Service_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func Service_DeleteFromId(id string) string {
	sql := "DELETE FROM services"
	sql += " WHERE id IN (" + id + ")"
	return sql
}
