package main

import (
	"farmservice/middleware"
	"farmservice/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ttoonn112/ktgolib/db"
	//_ "log"
	//_ "gorm.io/driver/mysql"
	//_ "gorm.io/gorm"
)

func main() {

	db.DB_Connections = map[string]map[string]string{
		"fs": map[string]string{
			"server": "clouddb01.bestgeosystem.com:10012",
			"user":   "farmserviceapp",
			"pass":   "BT@farmservice893",
			"dbname": "btfarmservice_db",
		},
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		//AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(middleware.JSONOnly)
	app.Use(middleware.HandleErrors)

	app.Post("/user/login", service.User_Login)
	app.Post("/user/register", service.User_Profile)
	app.Post("/user/profile", service.User_Profile)             // Get ข้อมูล user ที่ login ปัจจุบัน
	app.Post("/user/updateprofile", service.User_UpdateProfile) // Update ข้อมูล user ที่ login ปัจจุบัน เช่น ชื่อ, นามสกุล
	app.Post("/user/list", service.User_List)                   // รายการ User ในระบบ ใช้สำหรับ Admin (บน Web)
	app.Post("/user/detail", service.User_Detail)               // Get ข้อมูล user จาก Id (ใช้ร่วมกับ /user/list)
	app.Post("/user/update", service.User_Update)               // Update ข้อมูล user จาก Id (ใช้ร่วมกับ /user/list)
	app.Post("/user/delete", service.User_Delete)               // Delete ข้อมูล user จาก Id (ใช้ร่วมกับ /user/list)

	app.Post("/plot/list", service.Plot_List)
	app.Post("/plot/detail", service.Plot_Detail)
	app.Post("/plot/update", service.Plot_Update)
	app.Post("/plot/delete", service.Plot_Delete)

	app.Post("/vehicle/list", service.Vehicle_List)
	app.Post("/vehicle/detail", service.Vehicle_Detail)
	app.Post("/vehicle/update", service.Vehicle_Update)

	app.Post("/driver/list", service.Driver_List)
	app.Post("/driver/detail", service.Driver_Detail)
	app.Post("/driver/update", service.Driver_Update)
	app.Post("/driver/delete", service.Driver_Delete)

	app.Post("/friend/list", service.Friend_List)
	app.Post("/friend/detail", service.Friend_Detail)
	app.Post("/friend/update", service.Friend_Update)
	app.Post("/friend/delete", service.Friend_Delete)

	app.Post("/plan/list", service.Plan_List)
	app.Post("/plan/detail", service.Plan_Detail)
	app.Post("/plan/update", service.Plan_Update)
	app.Post("/plan/delete", service.Plan_Delete)

	app.Post("/service/list", service.Service_List)
	app.Post("/service/detail", service.Service_Detail)
	app.Post("/service/update", service.Service_Update)
	app.Post("/service/delete", service.Service_Delete)
	/*
		app.Post("/friend/list", service.Friend_List)
		app.Post("/friend/detail", service.Friend_Detail)
		app.Post("/friend/update", service.Friend_Update)
		app.Post("/friend/delete", service.Friend_Delete)

		app.Post("/plan/list", service.Plan_List)
		app.Post("/plan/detail", service.Plan_Detail)
		app.Post("/plan/update", service.Plan_Update)
		app.Post("/plan/delete", service.Plan_Delete)
	*/

	app.Post("/user/friend", service.Friend_Update)
	app.Post("/example/query", service.Example_Query)
	app.Post("/example/update", service.Example_Update)
	app.Post("/example/transaction", service.Example_Transaction)

	app.Listen(":8090")
}
