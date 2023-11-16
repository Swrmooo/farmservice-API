package util

import (
	"github.com/ttoonn112/ktgolib/sms"
  "github.com/ttoonn112/ktgolib/log"
)

var sms_provider_api_key string = "0d75690cfee73f7b3523c9b25d07e859"
var sms_provider_secret_key string = "XjLNNY0hLBLICd4d"
var project_key string = "FarmService"

func SendOTP(phone string) {
  msg, ok := sms.SendOTP_MKT(sms_provider_api_key, sms_provider_secret_key, project_key, phone);
  log.LogHidden("SendOTP", "", phone, msg, "OTP")
  if !ok {
    log.LogHidden("SendOTP", "", phone, msg, "OTPFailed")
  }
}

func ValidateOTP(phone string, token string, otp string) {
  msg, ok := sms.ValidateOTP_MKT(sms_provider_api_key, sms_provider_secret_key, token, otp);
  log.LogHidden("ValidateOTP", "", phone, "("+token+","+otp+") : "+msg, "OTP")
  if !ok {
    log.LogHidden("ValidateOTP", "", phone, "("+token+","+otp+") : "+msg, "OTPFailed")
  }
}
