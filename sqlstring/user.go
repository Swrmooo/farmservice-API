package sqlstring

import (

)

func User_GetToken(token string) string {

  sql := " select username, firstname, phone from users "
	sql += " where token = '"+token+"' "

  return sql
}

func User_Login(username string, pass string) string {

  sql := " select username, firstname, phone from users "
	sql += " where username = '"+username+"' and password = md5('"+pass+"') "

  return sql
}
