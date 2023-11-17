package lang

var TH map[string]string = map[string]string{
  "error.JsonArrayParsingFailed" : "Json array parsing failed",
  "error.InvalidJSONFormat" : "Invalid Json format",
  "error.TokenNotFound" : "User token not found",
  "error.IncorrectLogin" : "Username หรือ Password ไม่ถูกต้อง",
  "error.OperationFailed" : "Operation Failed !!!",
  "error.DBOperationFailed" : "Database operation failed !!!",
  "error.JsonFailed" : "Json Failed !!!",
  "error.ContactAdmin" : "ไม่สามารถดำเนินการได้ กรุณาติดต่อผู้ดูแลระบบ",

  "error.user.PhoneExists": "เบอร์โทรนี้มีผู้ใช้ลงทะเบียนแล้ว ไม่สามารถลงทะเบียนซ้ำได้",
  "error.user.OTPSendFailed": "ไม่สามารถส่ง OTP ได้ในขณะนี้ กรุณาลองอีกครั้งหรือติดต่อฝ่าย Support ลูกค้า",
  "error.user.InvalidOTP": "รหัส OTP ที่ระบุไม่ถูกต้อง",

  "require.Id" : "ไม่พบข้อมูล ID",
  "require.OTP" : "กรุณาระบุรหัส OTP",
  "require.Name" : "กรุณาระบุชื่อ-นามสกุล",
  "require.Username" : "กรุณาระบุชื่อผู้ใช้งานหรือเบอร์โทร",
  "require.Phone" : "กรุณาระบุเบอร์โทร",
  "require.PhoneNotValid" : "กรุณาระบุเบอร์โทร 10 หลักให้ถูกต้อง",
  "require.Password" : "กรุณาระบุรหัสผ่าน",

  "require.Plot.PlotType": "กรุณาระบุประเภทแปลง",
}
