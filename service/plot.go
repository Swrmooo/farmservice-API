package service

import (
	"farmservice/bu"
	"farmservice/middleware"
	"farmservice/sqlstring"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func Plot_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "plot_type"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("plot_type", lib.T(filters, "plot_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))

	list := bu.Plot_List(filter)

	return r.Success(list)
}

func Plot_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Plot_Detail(id)

	return r.Success(detail)
}

func Plot_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_Update")

	id := lib.T(r.Payload, "id")
	if lib.T(r.Payload, "plot_type") == "" {
		panic("require.Plot.PlotType")
	}

	payload := lib.GetMask(r.Payload, []string{"plot_type", "code", "area", "geo_field", "lat", "lng", "address", "detail", "pics"})

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

// func Plot_Delete(c *fiber.Ctx) error {
// 	r := middleware.GetUserRequestToken(c, "fs", "Plot_Delete")

// 	id := lib.T(r.Payload, "id")
// 	if id == "" { panic("require.Id") }

// 	db.Execute(r.Conn, sqlstring.Plot_DeleteFromId(id))

// 	return r.Success(nil)
// }

func Plot_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_Delete")

	var request struct {
		ID []int `json:"id"`
	}

	if err := c.BodyParser(&request); err != nil {
		panic("error.InvalidJSONFormat")
	}

	if len(request.ID) == 0 {
		panic("require.Id")
	}

	ids := make([]string, len(request.ID))
	for i, id := range request.ID {
		ids[i] = strconv.Itoa(id)
	}
	idString := strings.Join(ids, ",")

	sql := sqlstring.Plot_DeleteFromId(idString)
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
