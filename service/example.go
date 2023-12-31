package service

import (
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
	"farmservice/middleware"
	"github.com/gofiber/fiber/v2"
)

// ตัวอย่างการใช้งาน Transaction
func Example_Transaction(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "Example_Transaction")

	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	list := trans.Query("select * from users")
	for _, rdata := range list {
		trans.Execute("update users set token_expire_time = NOW() where id = '" + lib.T(rdata, "id") + "' ")
	}

	trans.Commit()
	trans.Close()

	return r.Success(list)
}

// ตัวอย่าง Query operation
func Example_Query(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "Example_Query")

	list := db.Query("fs", "select * from users")

	return r.Success(list)
}

// ตัวอย่าง Update operation
func Example_Update(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "Example_Update")

	id := lib.T(r.Payload, "id")
	tel := lib.T(r.Payload, "tel")
	if id == "" {
		panic("Id is not found")
	}

	sql := "update users set "
	if tel != "" {
		sql += " tel = '" + tel + "', "
	}
	sql += " token_expire_time = NOW() "
	sql += " where id = '" + id + "' "
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
