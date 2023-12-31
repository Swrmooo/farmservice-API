package sqlstring

import (
	lib "github.com/ttoonn112/ktgolib"
)

func Friend_get() string {
	sql := " SELECT id, user_id, friend_user_id, firstname, lastname, tel, mood, pics, code FROM friends "
	sql += " WHERE "
	return sql
}

func Friend_GetFromId(id string) string {
	sql := Friend_get()
	sql += " id = '" + id + "' "
	return sql
}

func Friend_GetFromCode(code string) string {
	sql := Friend_get()
	sql += " code = '" + code + "' "
	return sql
}

func Friend_GetFromFilter(filter string) string {
	sql := Friend_get()
	sql += filter
	return sql
}

func Friend_Create(code string, userId string, friendId string) string {
	sql := " insert into friends (code, doc_date, user_id, friend_user_id) values ('" + code + "', '" + lib.NowDate() + "', '" + userId + "', '" + friendId + "'); "
	return sql
}

func Friend_update(data map[string]interface{}) string {
	sql := " UPDATE friends set "
	for k, _ := range data {
		sql += " " + k + " = '" + lib.T(data, k) + "', "
	}
	sql += " last_updated_time = NOW() "
	sql += " WHERE "
	return sql
}

func Friend_UpdateFromId(id string, data map[string]interface{}) string {
	sql := Friend_update(data)
	sql += " id = '" + id + "' "
	return sql
}

func Friend_DeleteFromId(id string) string {
	sql := "DELETE FROM friends"
	sql += " WHERE id IN (" + id + ")"
	return sql
}

func Friend_Count(userId string) string {
	sql := "SELECT COUNT(id) FROM friends WHERE user_id = '" + userId + "'"
	return sql
}
