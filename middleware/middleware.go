package middleware

import (
  "time"
  "farmservice/lib/db"
  "farmservice/sqlstring"
  "farmservice/lib"
  "farmservice/lang"
  "github.com/gofiber/fiber/v2"
)

type RequestToken struct {
  Ctx *fiber.Ctx                    // fiber context
  Conn string                       // connection name (refer to server and database)
  Time time.Time                    // Request timestamp
	Operation string                  // Operation name
  User map[string]interface{}       // User object (if available)
  Payload map[string]interface{}    // Payload object
}

func JSONOnly(c *fiber.Ctx) error {
    if c.Method() == "POST" || c.Method() == "PUT" {
        contentType := c.Get("Content-Type")
        if contentType != "application/json" {
            return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
                "what": "error",
                "error_code": "InvalidContentType",
                "msg": "Only JSON content type is supported",
            })
        }
    }
    return c.Next()
}

func HandleErrors(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			switch err := r.(type) {
  			case error:
          c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "what": "error",
            "error_code": "Unexpected",
            "msg": "Internal Server Error ("+err.Error()+")",
          })
          break
        case string:
          error_code := err
          error_text := lang.Msg("th", err)
          if error_text == error_code {
            error_code = ""
          }
          c.JSON(fiber.Map{
            "what": "error",
            "error_code": error_code,
            "msg": error_text,
          })
          break
  			default:
          c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "what": "error",
            "error_code": "Unexpected",
            "msg": "Internal Server Error",
          })
			}

		}
	}()
	return c.Next()
}

func TryCatch(callback func(errStr string) ){
	 if r := recover(); r != nil {
		 errStr := ""
		 if err,ok := r.(error); ok {
			 errStr = err.Error()
		 }else if errS,ok := r.(string); ok {
			 errStr = errS
		 }
		 if(callback != nil){
			 callback(errStr)
		 }
	 }
}

func GetAnonymousRequestToken(c *fiber.Ctx, conn string, operation string) (*RequestToken) {
  return GetRequestToken(c, conn, operation, true)
}

func GetUserRequestToken(c *fiber.Ctx, conn string, operation string) (*RequestToken) {
  return GetRequestToken(c, conn, operation, false)
}

func GetRequestToken(c *fiber.Ctx, conn string, operation string, allowAnonymous bool) (*RequestToken) {

  t := time.Now()

  // Get user from token
  var user map[string]interface{}
  authorizationHeader := c.Get("Authorization")
  if authorizationHeader != "" {
    if users := db.Query(conn, sqlstring.User_GetToken(authorizationHeader)); len(users) == 1 {
      user = users[0]
    }
  }
  if user == nil && allowAnonymous == false {
    panic("error.TokenNotFound")
  }

  // Get payload
  var payload map[string]interface{}
  if err := c.BodyParser(&payload); err != nil {
    panic("error.InvalidJSONFormat")
  }

  if user != nil {
    lib.Log(operation, lib.T(user, "username"), "", lib.MapToString(payload), "Operation")
  }else{
    lib.Log(operation, "", "", lib.MapToString(payload), "Operation")
  }

  return &RequestToken{Ctx:c, Conn:conn, Time:t, Operation:operation, User:user, Payload:payload}
}

// info = ข้อมูลที่ส่งกลับ Client
func (r *RequestToken) Success(info interface{}) error {
  username := ""
  if r.User != nil {
    username = lib.T(r.User,"username")
  }
  text := "Unable parse data"
  if iinfo, ok := info.(map[string]interface{}); ok {
    text = lib.MapToString(iinfo)
  }else if iinfo, ok := info.([]map[string]interface{}); ok {
    text = lib.ArrayOfMapToString(iinfo)
  }
  lib.LogHiddenWithDuration(r.Operation, username, "Success", text, lib.I64_S(lib.DateTimeValueDiffSec(r.Time, time.Now())), "OperationSuccess")
  return r.Ctx.JSON(fiber.Map{
    "what": "ok",
    "info": info,
  })
}
