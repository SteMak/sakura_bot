package string_control

import (
	"regexp"
	"strings"
)

func ClearStrange(data string) string {
	
	codePattern := regexp.MustCompile(`\w\w\w\w/\w\w\w\w`)
	match := codePattern.FindStringSubmatch(data)
	res := "error"
	if len(match) > 0 {
		res = match[0]
	}
	return res
}

func ReplaceBadSymbols(str string) string {
	var (
		mes [300][2]string = [300][2]string {
			[2]string {"S", "5"},
			[2]string {"é", "e"},
			[2]string {"A", "4"},
			[2]string {"I", "1"},
			[2]string {"¢", "с"},
			[2]string {"O", "0"},
			[2]string {"Q", "0"},
			[2]string {"J", "7"},
			[2]string {"C", "c"},
			[2]string {"|", "1"},
			[2]string {"$", "8"},
		}
	)

	for i := 0; i < len(mes); i++ {
		str = strings.Replace(str, mes[i][0], mes[i][1], -1)
	}

	return str
}
