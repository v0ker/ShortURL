package utils

var mapping = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Int2String(code int64, minLength int32) string {
	//map int to a-zA-Z0-9
	var result string
	for code > 0 {
		result = string(mapping[code%62]) + result
		code /= 62
	}
	if int32(len(result)) < minLength {
		for i := int32(len(result)); i < minLength; i++ {
			result = string(mapping[0]) + result
		}
	}
	return result
}

func String2Int(code string) int64 {
	//map a-zA-Z0-9 to int
	var result int64
	for _, c := range code {
		result = result*62 + mapping2Int(c)
	}
	return result
}

func mapping2Int(c rune) int64 {
	if c >= 'a' && c <= 'z' {
		return int64(c - 'a')
	} else if c >= 'A' && c <= 'Z' {
		return int64(c - 'A' + 26)
	} else if c >= '0' && c <= '9' {
		return int64(c - '0' + 52)
	}
	return 0
}
