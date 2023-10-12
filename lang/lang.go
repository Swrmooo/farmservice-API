package lang

func Msg(lang string, code string) string {
	switch lang {
		case "th":
			if str, ok := TH[code]; ok {
				return str
			}
			break
		default:
			if str, ok := TH[code]; ok {
				return str
			}
	}
	return code
}
