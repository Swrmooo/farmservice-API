package main

import (
	"farmservice/lib/db"
	"farmservice/middleware"
	"farmservice/service"
	_ "log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

// var Db *gorm.DB
// var err error

// func InitDB() {
// 	dsn := "farmserviceapp:BT@farmservice893@tcp(clouddb01.bestgeosystem.com:10012)/btfarmservice_db"
// 	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v ", err)
// 	}

// 	// Db.AutoMigrate(&User{})
// }

func main() {

	db.DB_Connections = map[string]map[string]string{
		"fs": map[string]string{
			"server": "clouddb01.bestgeosystem.com:10012",
			"user":   "farmserviceapp",
			"pass":   "BT@farmservice893",
			"dbname": "btfarmservice_db",
		},
	}
	// InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		//AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(middleware.JSONOnly)
	app.Use(middleware.HandleErrors)

	app.Post("/user/login", service.User_Login)
	app.Post("/user/friend", service.Friend_Update)

	app.Post("/example/query", service.Example_Query)
	app.Post("/example/update", service.Example_Update)
	app.Post("/example/transaction", service.Example_Transaction)

	app.Listen(":8090")
}
