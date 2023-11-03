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

// func CheckVehicleLimit(conn string, userId string, memberType string) error {
// 	resultSQL := sqlstring.Vehicle_Count(userId)
// 	count := db.Query(conn, resultSQL)
// 	if len(count) > 0 {
// 		result, ok := count[0]["COUNT(id)"].(int64)
// 		if !ok {
// 			panic("Invalid count result")
// 		}

// 		switch memberType {
// 		case "guest", "standard":
// 			if result >= 20 {
// 				panic("You've reached the maximum limit of vehicles.")
// 			}
// 		case "gold":
// 			if result >= 100 {
// 				panic("You've reached the maximum limit of vehicles.")
// 			}
// 		case "premium":
// 			if result >= 200 {
// 				panic("You've reached the maximum limit of vehicles.")
// 			}
// 		case "testmember":
// 			if result >= 5 {
// 				panic("You've reached the maximum limit of vehicles.")
// 			}
// 		case "enterprise":
// 			// ไม่จำกัด
// 		default:
// 			panic("Invalid member type.")
// 		}
// 	} else {
// 		panic("Count result not found")
// 	}
// 	return nil

// }

func Vehicle_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "Vehicle_List")
	// userRole := db.GetUserRoleFromDatabase(r.User.ID)

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"start_date", "end_date", "vehicle_type", "vehicle", "user_id"})
	filter := " id <> 0 "
	filter += lib.AddSqlDateRangeFilter("doc_date", lib.T(filters, "start_date"), lib.T(filters, "end_date"))
	filter += lib.AddSqlFilter("vehicle_type", lib.T(filters, "vehicle_type"))
	filter += lib.AddSqlFilter("user_id", lib.T(filters, "user_id"))
	filter += lib.AddSqlFilter("vehicle", lib.T(filters, "vehicle"))

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

	payload := lib.GetMask(r.Payload, []string{"vehicle_type", "vehicle", "category", "num", "license_plate", "brand", "model", "driver"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	// // check COUTN(id) items
	// checkID := lib.T(r.User, "id")
	// checkMember := lib.T(r.User, "member")
	// err := CheckVehicleLimit(r.Conn, checkID, checkMember)
	// if err != nil {
	// 	panic(err.Error())
	// }

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

	sql := sqlstring.Vehicle_DeleteFromId(idString)
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
