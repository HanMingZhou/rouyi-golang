package ruoyiConv

func SubStr(str string, startIndex, endIndex int) string {
	rs := []rune(str)
	return string(rs[startIndex:endIndex])
}
