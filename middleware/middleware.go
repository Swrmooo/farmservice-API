package middleware

import (
  "btfarmservice/lib"
  "github.com/gofiber/fiber/v2"
)

type Request struct {
	Operation string
  User map[string]interface{}
  Payload map[string]interface{}
  Ctx *fiber.Ctx
}

func JSONOnly(c *fiber.Ctx) error {
    if c.Method() == "POST" || c.Method() == "PUT" {
        contentType := c.Get("Content-Type")
        if contentType != "application/json" {
            return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
                "what": "ok",
                "error_code": "InvalidJSONFormat",
                "msg": "Only JSON content type is supported",
            })
        }
    }
    return c.Next()
}

func Payload(c *fiber.Ctx, operation string) (Request, error) {
  var request map[string]interface{}
  if err := c.BodyParser(&request); err != nil {
      lib.Log(operation, "", "InvalidRequestFormat", "Invalid request format", "OperationUnexpected")
      return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "what": "error",
        "error_code": "InvalidRequestFormat",
        "msg": "Invalid request format",
      })
  }
  // Getting user token
  user := map[string]interface{}{
    "username": lib.T(request, "username"),
  }
  lib.Log(operation, lib.T(user, "username"), "Called", "Success", "Operation")
  return Request{Ctx:c, User:user, Operation:operation, Payload:request}, nil
}

func (r *Request) Unauthorized(msg string) error {
  lib.Log(r.Operation, lib.T(r.User,"username"), "Unauthorized", msg, "UserUnauthorized")
  return r.Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
    "what": "error",
    "error_code": "Unauthorized",
    "msg": msg,
  })
}

func (r *Request) Error(error string, msg string) error {
  lib.Log(r.Operation, lib.T(r.User,"username"), error, msg, "OperationError")
  return r.Ctx.JSON(fiber.Map{
    "what": "error",
    "error_code": error,
    "msg": msg,
  })
}

func (r *Request) Success(info map[string]interface{}) error {
  return r.Ctx.JSON(fiber.Map{
    "what": "ok",
    "info": info,
  })
}
