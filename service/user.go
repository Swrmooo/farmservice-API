package service

import (
	"farmservice/middleware"
	"farmservice/bu"
	lib "github.com/ttoonn112/ktgolib"
	"github.com/ttoonn112/ktgolib/db"
	"github.com/gofiber/fiber/v2"
)

func User_Login(c *fiber.Ctx) error {
	r := middleware.GetAnonymousRequestToken(c, "fs", "User_Login")

	username := lib.T(r.Payload, "username")
	pass := lib.T(r.Payload, "password")

	if username == "" {
		panic("require.Username")
	} else if pass == "" {
		panic("require.Password")
	}

	if user := bu.User_Login(username, pass); user != nil {
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

	return r.Success(profile)										// ตอบกลับ Success พร้อมค่า profile
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

	return r.Success(profile)										// ตอบกลับ Success พร้อมค่า profile
}

func User_List(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_List")

	// ค้นหา User จาก member, tel
	filters := lib.GetMask(r.Payload, []string{"member", "tel"})
	filter := " id <> 0 "
	filter += lib.AddSqlFilter("member", lib.T(filters, "member"))
	filter += lib.AddSqlFilter("tel", lib.T(filters, "tel"))

	list := bu.User_List(filter)

	return r.Success(list)										// ตอบกลับ Success พร้อมค่า profile
}

func User_Detail(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Detail")

	id := lib.T(r.Payload, "id")
	if id == "" { panic("require.Id") }

	profile := bu.User_Detail(id)

	return r.Success(profile)										// ตอบกลับ Success พร้อมค่า profile
}

func User_Update(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Update")

	id := lib.T(r.Payload, "id")
	tel := lib.T(r.Payload, "tel")

	if tel == "" { panic("require.Phone") } else
	if lib.T(r.Payload, "firstname") == "" || lib.T(r.Payload, "lastname") == "" { panic("require.Name") }

	// ดึงค่า field ที่ต้องการมาจาก r.Payload เช่น tel, firstname, lastname
	payload := lib.GetMask(r.Payload, []string{"tel", "firstname", "lastname"})

	// Start transaction
	trans := db.OpenTrans(r.Conn)
	defer middleware.TryCatch(func(errStr string) {
		trans.Rollback()
		trans.Close()
		panic(errStr)
	})

	// กรณีสร้าง User ใหม่ (ถ้าไม่ส่งค่า ID มา)
	if id == "" {
		id = bu.User_Create(trans, tel)
	}

	// อัพเดทข้อมูล
	trans.Execute(sqlstring.User_UpdateFromId(id, payload))

	// End transaction
	trans.Commit()
	trans.Close()

	// ดึงข้อมูล User Profile จาก ID
	profile = bu.User_Detail(id)

	return r.Success(profile)										// ตอบกลับ Success พร้อมค่า profile
}

func User_Delete(c *fiber.Ctx) error {
	r := middleware.GetUserRequestToken(c, "fs", "User_Delete")

	id := lib.T(r.Payload, "id")
	if id == "" { panic("require.Id") }

	db.Execute(r.Conn, sqlstring.User_DeleteFromId(id))

	return r.Success(nil)
}
