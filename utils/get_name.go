package utils

import "strings"

func GetNameAndExt(filename string) (string, string) {
	splt := strings.Split(filename, ".")
	n := len(splt)
	if n == 0 {
		return "", ""
	}
	if n == 1 {
		return filename, ""
	}
	return strings.Join(splt[:n-1], ""), splt[n-1]
}
