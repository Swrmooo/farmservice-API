package service

import (
	"farmservice/bu"
	"farmservice/middleware"
	"farmservice/sqlstring"

	"github.com/gofiber/fiber/v2"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Friend_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Friend_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "tel", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("tel", lib.T(filters, "tel"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))

	list := bu.Friend_List(filter)

	return r.Success(list)
}

func Friend_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Friend_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Friend_Detail(id)

	return r.Success(detail)
}

func Friend_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Friend_Update")

	id := lib.T(r.Payload, "id")
	// if lib.T(r.Payload, "tel") == "" {
	// 	panic("require.PhoneNumber")
	// } else if lib.T(r.Payload, "firstname") == "" {
	// 	panic("require.Firstname")
	// } else if lib.T(r.Payload, "lastname") == "" {
	// 	panic("require.Lastname")
	// }

	payload := lib.GetMask(r.Payload, []string{"num", "firstname", "lastname", "tel", "mood", "pics"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	if id == "" {
		id = bu.Friend_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Friend_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Friend_Detail(id)

	return r.Success(detail)
}

func Friend_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Friend_Delete")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	db.Execute(r.Conn, sqlstring.Friend_DeleteFromId(id))

	return r.Success(nil)
}
