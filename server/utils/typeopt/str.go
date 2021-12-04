package typeopt

import "strings"

func ReplaceStr(str string, re []string) string {

	for _, v := range re {
		str = strings.Replace(str, v, "", -1)
	}
	return str
}
