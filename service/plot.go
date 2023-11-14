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

func Plot_Limit(conn string, userId string, memberType string) error {
	resultSQL := sqlstring.Plot_Count(userId)
	count := db.Query(conn, resultSQL)
	if len(count) > 0 {
		result, ok := count[0]["COUNT(id)"].(int64)
		if !ok {
			panic("Invalid count result")
		}

		switch memberType {
		case "guest", "standard":
			if result >= 20 {
				panic("You've reached the maximum limit of plot.")
			}
		case "gold", "premium":
			if result >= 100 {
				panic("You've reached the maximum limit of plot.")
			}
		case "testmember":
			if result >= 5 {
				panic("You've reached the maximum limit of plot.")
			}
		case "enterprise":
			// ไม่จำกัด
		default:
			panic("Please register as a member before using Application")
		}
	} else {
		panic("Count result not found")
	}
	return nil

}

func Plot_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Plot_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"name", "start_date", "end_date", "plot_type", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("plot_type", lib.T(filters, "plot_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))
	filter += lib.AddSqlFilter("name", lib.T(filters, "name"))

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

	payload := lib.GetMask(r.Payload, []string{"user_id", "name", "num", "plot_type", "code", "area", "geo_field", "lat", "lng", "address", "area_type", "detail", "land_ownership", "pics"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	// check COUTN(id) items
	checkID := lib.T(r.User, "id")
	checkMember := lib.T(r.User, "member")
	err := Plot_Limit(r.Conn, checkID, checkMember)
	if err != nil {
		panic(err.Error())
	}

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
