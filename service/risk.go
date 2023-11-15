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

func Risk_Limit(conn string, userId string, memberType string) error {
	resultSQL := sqlstring.Risk_Count(userId)
	count := db.Query(conn, resultSQL)
	if len(count) > 0 {
		result, ok := count[0]["COUNT(id)"].(int64)
		if !ok {
			panic("Invalid count result")
		}

		switch memberType {
		case "guest", "standard":
			if result >= 80 {
				panic("You've reached the maximum limit of risk area.")
			}
		case "gold":
			if result >= 400 {
				panic("You've reached the maximum limit of risk area.")
			}
		case "testmember":
			if result >= 3 {
				panic("You've reached the maximum limit of risk area.")
			}
		case "premium", "enterprise":
			// ไม่จำกัด
		default:
			panic("Please register as a member before using Application")
		}
	} else {
		panic("Count result not found")
	}
	return nil

}

func Risk_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Risk_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"risk_type", "user_id", "lat", "lng", "radius"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("risk_type", lib.T(filters, "risk_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))
	filter += lib.AddSqlFilter("lat", lib.T(filters, "lat"))
	filter += lib.AddSqlFilter("lng", lib.T(filters, "lng"))
	filter += lib.AddSqlFilter("radius", lib.T(filters, "radius"))

	list := bu.Risk_List(filter)

	return r.Success(list)
}

func Risk_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Risk_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Risk_Detail(id)

	return r.Success(detail)
}

func Risk_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Risk_Update")

	id := lib.T(r.Payload, "id")
	if lib.T(r.Payload, "risk_type") == "" {
		panic("require risk_type")
	}
	if lib.T(r.Payload, "plot_id") == "" {
		panic("require plot_id")
	}

	payload := lib.GetMask(r.Payload, []string{"plot_id", "risk_type", "lat", "lng", "radius", "geo_field"})

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
	err := Risk_Limit(r.Conn, checkID, checkMember)
	if err != nil {
		panic(err.Error())
	}

	if id == "" {
		id = bu.Risk_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Risk_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Risk_Detail(id)

	return r.Success(detail)
}

func Risk_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Risk_Delete")

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

	sql := sqlstring.Risk_DeleteFromId(idString)
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
