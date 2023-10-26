package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Plan_get() string {
	//sql := " SELECT id, code, plan_type, doc_date name, detail FROM Plans "
	sql := " SELECT id, user_id, plan, num, driver, model, Plan_type, brand FROM plan "
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
	sql += " code = '" + code + "' "
	return sql
}

func Plan_GetFromFilter(filter string) string {
	sql := Plan_get()
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
	sql := " DELETE from plan "
	sql += " WHERE id = '" + id + "' "
	return sql
}
