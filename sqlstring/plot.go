package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func plot_get() string {
	sql := " SELECT id, user_id, area, geo_field, lat, lng, address, area_type, detail, land_ownership, pics, plot_type, doc_date, code FROM plot "
	sql += " WHERE "
	return sql
}

func Plot_GetFromId(id string) string {
	sql := plot_get()
	sql += " id = '" + id + "' "
	return sql
}

func Plot_GetFromCode(code string) string {
	sql := plot_get()
	sql += " code = '" + code + "' "
	return sql
}

func Plot_GetFromFilter(filter string) string {
	sql := plot_get()
	sql += filter
	return sql
}

func Plot_Create(code string, userId string) string {
	sql := " insert into plot (code, doc_date, user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "'); "
	return sql
}

// ('POLYGON(" + v.(string) + ")')
func plot_update(data map[string]interface{}) string {
	sql := "UPDATE plot SET "
	for k, v := range data {
		if k == "geo_field" {
			sql += k + " = ST_GeomFromText('POLYGON(" + v.(string) + ")'), "
		} else {
			sql += k + " = '" + lib.T(data, k) + "', "
		}
	}
	sql += "last_updated_time = NOW() "
	sql += "WHERE "
	return sql
}

func Plot_UpdateFromId(id string, data map[string]interface{}) string {
	sql := plot_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func Plot_DeleteFromId(id string) string {
	sql := "DELETE FROM plot"
	sql += " WHERE id IN (" + id + ")"
	return sql
}

func Plot_Count(userId string) string {
	sql := "SELECT COUNT(id) FROM plot WHERE user_id = '" + userId + "'"
	return sql
}
