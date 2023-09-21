package service

import (
  "btfarmservice/middleware"
  "btfarmservice/lib"
  "github.com/gofiber/fiber/v2"
)

func User_Login(c *fiber.Ctx) error {
  r, err :=  middleware.Payload(c, "User_Login")
  if err != nil { return err }

  if lib.T(r.Payload, "username") != "admin" {
    return r.Error("", "Invalid username or password")
  }

  info := map[string]interface{}{}
  info["username"] = lib.T(r.Payload, "username")

  return r.Success(info)
}


func User_Register(c *fiber.Ctx) error {
  operation := "User_Register"
  payload, user, err :=  middleware.Payload(c, operation)
  if err != nil { return err }

  if lib.T(payload, "name") == "" {
    return middleware.Error(c, operation, user, "", "Name is required")
  }

  info := map[string]interface{}{}
  info["name"] = lib.T(payload, "name")

  return middleware.Success(c, operation, user, info)
}
