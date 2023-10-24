package service

import (
	//lib "github.com/ttoonn112/ktgolib"
	//"github.com/ttoonn112/ktgolib/db"
	"farmservice/middleware"
	_ "farmservice/middleware"

	//"farmservice/lib/db"
	_ "farmservice/bu"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
)

func Friend_Update(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "Friend_Update")

	var inputData struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Tel       string `json:"tel"`
	}

	if err := c.BodyParser(&inputData); err != nil {
		panic("Failed to parse JSON: " + err.Error())
	}

	id := lib.T(r.Payload, "id")

	if id == "" {
		panic("Id is not found")
	}

	sql := "update users set "

	if inputData.FirstName != "" {
		sql += " first_name = '" + inputData.FirstName + "', "
	}

	if inputData.LastName != "" {
		sql += " last_name = '" + inputData.LastName + "', "
	}

	sql += " token_expire_time = NOW() "
	sql += " where id = '" + id + "' "

	db.Execute(r.Conn, sql)

	return r.Success(nil)
}
