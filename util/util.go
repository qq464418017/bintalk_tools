package util

func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}

	return string(strArry)
}

func StrFirstToUpperBySql(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	s := string(strArry)
	if s == "Id" {
		return "ID"
	}

	return s
}

func StrToUpper(str string) string {
	strArry := []rune(str)
	for i := 0; i < len(strArry); i++ {
		if strArry[i] >= 97 && strArry[i] <= 122 {
			strArry[i] -= 32
		}
	}
	return string(strArry)
}