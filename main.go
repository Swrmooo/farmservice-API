package main

import (
  "farmservice/middleware"
  "farmservice/service"
  "farmservice/lib/db"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

    db.DB_Connections = map[string]map[string]string{
      "fs" : map[string]string{
        "server" : "clouddb01.bestgeosystem.com:10012",
        "user" : "farmserviceapp",
        "pass" : "BT@farmservice893",
        "dbname" : "btfarmservice_db",
      },
    }

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        //AllowOrigins: "https://gofiber.io, https://gofiber.net",
        AllowHeaders:  "Origin, Content-Type, Accept",
        AllowMethods: "GET, POST, PUT, DELETE",
    }))

    app.Use(middleware.JSONOnly)
    app.Use(middleware.HandleErrors)

    app.Post("/user/login", service.User_Login)
    app.Post("/example/query", service.Example_Query)
    app.Post("/example/update", service.Example_Update)
    app.Post("/example/transaction", service.Example_Transaction)

    app.Listen(":8090")
}
