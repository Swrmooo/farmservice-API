package service

import (
	"farmservice/bu"
	"farmservice/middleware"
	"farmservice/sqlstring"

	"github.com/gofiber/fiber/v2"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Plan_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plan_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "Plan_type", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("Plan_type", lib.T(filters, "Plan_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))

	list := bu.Plan_List(filter)

	return r.Success(list)
}

func Plan_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plan_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Plan_Detail(id)

	return r.Success(detail)
}

func Plan_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plan_Update")

	id := lib.T(r.Payload, "id")
	if lib.T(r.Payload, "Plan_type") == "" {
		panic("require.Plan.PlanType")
	}

	payload := lib.GetMask(r.Payload, []string{"Plan_type", "num", "license_plate", "brand", "model", "driver"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	if id == "" {
		id = bu.Plan_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Plan_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Plan_Detail(id)

	return r.Success(detail)
}

func Plan_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plan_Delete")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	db.Execute(r.Conn, sqlstring.Plan_DeleteFromId(id))

	return r.Success(nil)
}
