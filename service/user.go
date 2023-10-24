package service

import (
	"farmservice/lib"
	"farmservice/middleware"

	//"farmservice/lib/db"
	"farmservice/bu"

	"github.com/gofiber/fiber/v2"
)

func User_Login(c *fiber.Ctx) error {
	// userId := c.Params("userid")
	// var data map[string]string

	// err := c.BodyParser(&data)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "Invalid Post Request",
	// 	})
	// }

	// if data["passcode"] == ""{
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"success":false,
	// 		"message": "Password is required",
	// 		"error": map[string]interface{}{},
	// 	})
	// }

	// var user models.User
	// db.DB.Where("id=?", userId).First(&user)

	r := middleware.GetAnonymousRequestToken(c, "fs", "User_Login")

	if lib.T(r.Payload, "tel") == "" {
		panic("Please input tel")
	} else if lib.T(r.Payload, "password") == "" {
		panic("Please input password")
	}

	if user := bu.User_Login(lib.T(r.Payload, "tel"), lib.T(r.Payload, "password")); user != nil {
		return r.Success(user)
	} else {
		panic("error.IncorrectLogin")
	}

}
