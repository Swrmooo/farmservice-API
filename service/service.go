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

func Service_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Service_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "service_type", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("service_type", lib.T(filters, "service_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))

	list := bu.Service_List(filter)

	return r.Success(list)
}

func Service_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Service_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Service_Detail(id)

	return r.Success(detail)
}

func Service_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Service_Update")

	id := lib.T(r.Payload, "id")
	if lib.T(r.Payload, "service_type") == "" {
		panic("require.Service.ServiceType")
	}

	payload := lib.GetMask(r.Payload, []string{"service_type", "service", "fee", "detail"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	if id == "" {
		id = bu.Service_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Service_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Service_Detail(id)

	return r.Success(detail)
}

func Service_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Service_Delete")

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

	sql := sqlstring.Service_DeleteFromId(idString)
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
