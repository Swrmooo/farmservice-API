package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Plan_Join() string {
	sql := "SELECT p.id, p.start_date, p.end_date, p.user_id, p.vehicle_id, p.driver_id, p.plan_type, p.plan, p.code, p.status, d.firstname, d.lastname, v.license_plate "
	sql += "FROM plan p "
	sql += "JOIN drivers d ON p.driver_id = d.id "
	sql += "JOIN vehicles v ON p.vehicle_id = v.id "
	sql += "WHERE "
	return sql
}

func Plan_get() string {
	sql := " SELECT id, start_date, end_date, user_id, vehicle_id, driver_id, job, plan, plan_type, status, code, doc_date, last_updated_time "
	sql += " FROM plan "
	sql += " WHERE "
	return sql
}

func Plan_GetFromUserId(id string) string {
	sql := Plan_get()
	sql += " id = '" + id + "' "
	return sql
}

func Plan_GetFromId(id string) string {
	sql := Plan_get()
	sql += " id = '" + id + "' "
	return sql
}

func Plan_GetFromCode(code string) string {
	sql := Plan_get()
	// sql := Plan_Join(code)
	sql += " code = '" + code + "' "
	return sql
}

func Plan_GetFromFilterToJoin(filter string) string {
	sql := Plan_Join()
	sql += filter
	//sql += lib.AddSqlFilter()
	return sql
}

func Plan_GetFromFilter(filter string) string {
	sql := Plan_get()
	// sql := Plan_Join(filter)
	sql += filter
	return sql
}

func Plan_Create(code string, userId string) string {
	sql := " insert into plan (code, doc_date, user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "'); "
	return sql
}

func Plan_update(data map[string]interface{}) string {
	sql := " UPDATE plan set "
	for k, _ := range data {
		sql += " " + k + " = '" + lib.T(data, k) + "', "
	}
	sql += " last_updated_time = NOW() "
	sql += " WHERE "
	return sql
}

func Plan_UpdateFromId(id string, data map[string]interface{}) string {
	sql := Plan_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func Plan_DeleteFromId(id string) string {
	sql := "DELETE FROM plan"
	sql += " WHERE id IN (" + id + ")"
	return sql
}

func Plan_Count(userId string) string {
	sql := "SELECT COUNT(id) FROM plan WHERE user_id = '" + userId + "'"
	return sql
}
