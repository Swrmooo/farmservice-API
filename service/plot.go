package service

import (
	"farmservice/sqlstring"
	"farmservice/middleware"
	"farmservice/bu"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
	"github.com/gofiber/fiber/v2"
)

func Plot_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "plot_type"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("plot_type", lib.T(filters, "plot_type"))

	list := bu.Plot_List(filter)

	return r.Success(list)
}

func Plot_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" { panic("require.Id") }

	detail := bu.Plot_Detail(id)

	return r.Success(detail)
}

func Plot_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_Update")

	id := lib.T(r.Payload, "id")
	if lib.T(r.Payload, "plot_type") == "" { panic("require.Plot.PlotType") }

	payload := lib.GetMask(r.Payload, []string{"plot_type", "name", "detail"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	if id == "" {
		id = bu.Plot_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Plot_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Plot_Detail(id)

	return r.Success(detail)
}

func Plot_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_Delete")

	id := lib.T(r.Payload, "id")
	if id == "" { panic("require.Id") }

	db.Execute(r.Conn, sqlstring.Plot_DeleteFromId(id))

	return r.Success(nil)
}
