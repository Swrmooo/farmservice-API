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

func Driver_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Driver_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "tel", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("tel", lib.T(filters, "tel"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))

	list := bu.Driver_List(filter)

	return r.Success(list)
}

func Driver_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Driver_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.Driver_Detail(id)

	return r.Success(detail)
}

func Driver_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Driver_Update")

	id := lib.T(r.Payload, "id")

	payload := lib.GetMask(r.Payload, []string{"num", "firstname", "lastname", "tel", "mood", "pics"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	if id == "" {
		id = bu.Driver_Create(trans, lib.T(r.User, "id"))
	}

	trans.Execute(sqlstring.Driver_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	detail := bu.Driver_Detail(id)

	return r.Success(detail)
}

func Driver_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Driver_Delete")

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

	sql := sqlstring.Driver_DeleteFromId(idString)
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
