package service

import (
  "farmservice/middleware"
  "farmservice/lib"
  //"farmservice/lib/db"
  "farmservice/bu"
  "github.com/gofiber/fiber/v2"
)

func User_Login(c *fiber.Ctx) error {
  r :=  middleware.GetAnonymousRequestToken(c, "fs", "User_Login")

  if lib.T(r.Payload, "username") == "" {
    panic("Please input username")
  }else if lib.T(r.Payload, "password") == "" {
    panic("Please input password")
  }

  if user := bu.User_Login( lib.T(r.Payload, "username"), lib.T(r.Payload, "password") ); user != nil {
    return r.Success(user)
  }else{
    panic("error.IncorrectLogin")
  }

}
