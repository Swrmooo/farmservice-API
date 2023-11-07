package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Risk_get() string {
	sql := " SELECT id, user_id, plot_id, icon_type, risk_type, lat, lng, radius, field, code, doc_date, last_updated_time FROM plot_risk "
	sql += " WHERE "
	return sql
}

func Risk_GetFromId(id string) string {
	sql := Risk_get()
	sql += " id = '" + id + "' "
	return sql
}

func Risk_GetFromCode(code string) string {
	sql := Risk_get()
	sql += " code = '" + code + "' "
	return sql
}

func Risk_GetFromFilter(filter string) string {
	sql := Risk_get()
	sql += filter
	return sql
}

func Risk_Create(code string, userId string) string {
	sql := " insert into plot_risk (code, doc_date, user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "'); "
	return sql
}

func Risk_update(data map[string]interface{}) string {
	sql := "UPDATE plot_risk SET "
	for k, v := range data {
		if k == "field" {
			sql += k + " = ST_GeomFromText('" + v.(string) + "'), "
		} else {
			sql += k + " = '" + lib.T(data, k) + "', "
		}
	}
	sql += "last_updated_time = NOW() "
	sql += "WHERE "
	return sql
}

func Risk_UpdateFromId(id string, data map[string]interface{}) string {
	sql := Risk_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func Risk_DeleteFromId(id string) string {
	sql := "DELETE FROM plot_risk"
	sql += " WHERE id IN (" + id + ")"
	return sql
}

func Risk_Count(userId string) string {
	sql := "SELECT COUNT(id) FROM plot_risk WHERE user_id = '" + userId + "'"
	return sql
}
