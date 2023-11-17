package util

import (
	"github.com/ttoonn112/ktgolib/sms"
  "github.com/ttoonn112/ktgolib/log"
)

var sms_provider_api_key string = "0d75690cfee73f7b3523c9b25d07e859"
var sms_provider_secret_key string = "XjLNNY0hLBLICd4d"
var project_key string = "4d9028cd56"		// FarmService

func SendOTP(phone string) (otp_token string) {
  msg, ok := sms.SendOTP_MKT(sms_provider_api_key, sms_provider_secret_key, project_key, phone);
  log.Log("SendOTP", "", phone, msg, "OTP")
  if !ok {
    log.Log("SendOTP", "", phone, msg, "OTPFailed")
		return ""
  }else{
		return msg
	}
}

func ValidateOTP(phone string, token string, otp string) bool {
  msg, ok := sms.ValidateOTP_MKT(sms_provider_api_key, sms_provider_secret_key, token, otp);
  log.Log("ValidateOTP", "", phone, "("+token+","+otp+") : "+msg, "OTP")
  if !ok {
    log.Log("ValidateOTP", "", phone, "("+token+","+otp+") : "+msg, "OTPFailed")
		return false
  }else{
		return true
	}
}
