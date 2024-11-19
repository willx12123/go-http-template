package util

func SafeStringSlice(str string, from, to int) string {
	strRunes := []rune(str)
	if from > len(strRunes) || from < 0 {
		return ""
	}
	if to > len(strRunes) {
		to = len(strRunes)
	}
	if to < from {
		return ""
	}
	return string(strRunes[from:to])
}
