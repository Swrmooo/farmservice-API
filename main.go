package main

import (
  "btfarmservice/middleware"
  "btfarmservice/service"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

  /*
    DB_server := "clouddb01.bestgeosystem.com"
    DB_port := "10012"
    DB_user := "farmserviceapp"
    DB_pass := "@farmserviceappDB997"
  */

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        //AllowOrigins: "https://gofiber.io, https://gofiber.net",
        AllowHeaders:  "Origin, Content-Type, Accept",
        AllowMethods: "GET, POST, PUT, DELETE",
    }))

    app.Use(middleware.JSONOnly)

    app.Post("/login", service.User_Login)
    app.Post("/register", service.User_Register)

    app.Listen(":8090")
}
