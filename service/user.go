package service

import (
	"crypto/md5"
	"encoding/hex"
	"farmservice/bu"
	"farmservice/middleware"
	"farmservice/sqlstring"
	"farmservice/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
)

func User_Login(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "User_Login")

	// username := lib.T(r.Payload, "username")
	tel := lib.T(r.Payload, "tel")
	pass := lib.T(r.Payload, "password")

	// if username == "" {
	// 	panic("require.Username")
	// }

	if tel == "" {
		panic("require.Telephone")
	}

	if pass == "" {
		panic("require.Password")
	}

	if user := bu.User_Login(tel, pass); user != nil {
		return r.Success(user)
	} else {
		panic("error.IncorrectLogin")
	}

}

func User_Profile(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Profile")

	id := lib.T(r.User, "id")

	// ดึงข้อมูล User Profile จาก ID
	profile := bu.User_Detail(id)

	return r.Success(profile) // ตอบกลับ Success พร้อมค่า profile
}

func User_UpdateProfile(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_UpdateProfile")

	id := lib.T(r.User, "id")

	// ดึงค่า field ที่ต้องการมาจาก r.Payload เช่น firstname, lastname
	payload := lib.GetMask(r.Payload, []string{"firstname", "lastname"})

	// อัพเดทข้อมูลลง Database
	db.Execute(r.Conn, sqlstring.User_UpdateFromId(id, payload))

	// ดึงข้อมูล User Profile จาก ID
	profile := bu.User_Detail(id)

	return r.Success(profile) // ตอบกลับ Success พร้อมค่า profile
}

func User_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"member", "tel"})
	filter := " id <> 0 "
	filter += lib.AddSqlFilter("member", lib.T(filters, "member"))
	filter += lib.AddSqlFilter("tel", lib.T(filters, "tel"))

	list := bu.User_List(filter)

	return r.Success(list) // ตอบกลับ Success พร้อมค่า profile
}

func User_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" {
		panic("require.Id")
	}

	detail := bu.User_Detail(id)

	return r.Success(detail) // ตอบกลับ Success พร้อมค่า profile
}

func User_Register(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "User_Register")

	id := lib.T(r.Payload, "id")
	// firstname := lib.T(r.Payload, "firstname")
	// lastname := lib.T(r.Payload, "lastname")
	tel := lib.T(r.Payload, "tel")
	//email := lib.T(r.Payload, "email")
	password := lib.T(r.Payload, "password")

	if tel == "" {
		panic("require.Phone")
	} else if len(tel) != 10 {
		panic("require.PhoneNotValid")
	}
	if lib.T(r.Payload, "firstname") == "" || lib.T(r.Payload, "lastname") == "" {
		panic("require.Name")
	}
	//if email == "" {
	//	panic("require.Email")
	//}
	if password == "" {
		panic("require.Password")
	}

	encodedPass := md5.Sum([]byte(password))
	passMd5 := hex.EncodeToString(encodedPass[:])
	fmt.Println("function passMd5: ", passMd5)
	password = passMd5
	fmt.Println("password : ", password)

	payload := lib.GetMask(r.Payload, []string{"tel", "firstname", "lastname", "username", "member"})

	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	// fmt.Println("service tel : ", tel)
	// fmt.Println("service trans : ", trans)

	// กรณีสร้าง User ใหม่ (ถ้าไม่ส่งค่า ID มา)
	if id == "" {
		id = bu.User_Register(trans, tel, password)
		if otp_token := util.SendOTP(tel); otp_token != "" {
			trans.Execute(sqlstring.User_UpdateFromId(id, map[string]interface{}{
				"otp_token": otp_token,
			}))
		} else {
			panic("error.user.OTPSendFailed")
		}
	}

	// อัพเดทข้อมูล
	trans.Execute(sqlstring.User_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	// ดึงข้อมูล User Profile จาก ID
	detail := bu.User_Detail(id)

	return r.Success(detail) // ตอบกลับ Success พร้อมค่า profile
}

func User_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Update")
	// r := middleware.GetAnonymousRequestToken(c, "fs", "User_Update")

	id := lib.T(r.Payload, "id")
	tel := lib.T(r.Payload, "tel")
	email := lib.T(r.Payload, "email")
	password := lib.T(r.Payload, "password")

	if tel == "" {
		panic("require.Phone")
	}
	if lib.T(r.Payload, "firstname") == "" || lib.T(r.Payload, "lastname") == "" {
		panic("require.Name")
	}
	if email == "" {
		panic("require.Email")
	}
	if password == "" {
		panic("require.Password")
	}

	// ดึงค่า field ที่ต้องการมาจาก r.Payload เช่น tel, firstname, lastname
	payload := lib.GetMask(r.Payload, []string{"tel", "firstname", "lastname", "password"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	// กรณีสร้าง User ใหม่ (ถ้าไม่ส่งค่า ID มา)
	if id == "" {
		id = bu.User_Create(trans, tel, password)
	}

	// อัพเดทข้อมูล
	trans.Execute(sqlstring.User_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	// ดึงข้อมูล User Profile จาก ID
	// ดึงข้อมูล User Profile จาก ID
	detail := bu.User_Detail(id)

	return r.Success(detail) // ตอบกลับ Success พร้อมค่า profile
}

func User_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Delete")

	var request struct {
		ID []int `json:"id"`
	}

	if err := c.BodyParser(&request); err != nil {
		panic("error.InvalidJSONFormat")
	}

	if len(request.ID) == 0 {
		panic("require.Id")
	}

	ids := make([]string, len(request.ID))
	for i, id := range request.ID {
		ids[i] = strconv.Itoa(id)
	}
	idString := strings.Join(ids, ",")

	sql := sqlstring.User_DeleteFromId(idString)
	db.Execute(r.Conn, sql)

	return r.Success(nil)
}

func User_OTPCheck(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "User_OTPCheck")

	tel := lib.T(r.Payload, "tel")
	otp := lib.T(r.Payload, "otp")

	if tel == "" {
		panic("require.Phone")
	} else if otp == "" {
		panic("require.OTP")
	}

	details := db.Query(r.Conn, sqlstring.User_GetAccessTokenFromPhone(tel))
	uac := details[0]

	if isValid := util.ValidateOTP(tel, lib.T(uac, "otp_token"), otp); !isValid {
		panic("error.user.InvalidOTP")
	}

	db.Execute(r.Conn, sqlstring.User_UpdateFromId(lib.T(uac, "id"), map[string]interface{}{
		"otp_token": "Validated on " + lib.Now(),
	}))

	return r.Success(nil)
}
