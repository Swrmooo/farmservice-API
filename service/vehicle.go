package service

import (
	"farmservice/bu"
	"farmservice/middleware"
	"farmservice/sqlstring"

	"github.com/gofiber/fiber/v2"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Vehicle_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Vehicle_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "vehicle_type", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("vehicle_type", lib.T(filters, "vehicle_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))

	list := bu.Vehicle_List(filter)

	return r.Success(list)
}

func Vehicle_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Vehicle_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Vehicle_Detail(id)

	return r.Success(detail)
}

func Vehicle_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Vehicle_Update")

	id := lib.T(r.Payload, "id")
	if lib.T(r.Payload, "vehicle_type") == "" {
		panic("require.Vehicle.VehicleType")
	}

	payload := lib.GetMask(r.Payload, []string{"vehicle_type", "num", "license_plate", "brand", "model", "driver"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	if id == "" {
		id = bu.Vehicle_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Vehicle_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Vehicle_Detail(id)

	return r.Success(detail)
}

func Vehicle_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Vehicle_Delete")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	db.Execute(r.Conn, sqlstring.Vehicle_DeleteFromId(id))

	return r.Success(nil)
}
