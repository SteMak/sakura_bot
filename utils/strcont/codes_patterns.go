package strcont

import (
	"regexp"
	"strings"
)

// ClearStrange find codes in text
func ClearStrange(data string) string {

	codePattern := regexp.MustCompile(`\w\w\w\w/\w\w\w\w`)
	match := codePattern.FindStringSubmatch(data)
	res := "error"
	if len(match) > 0 {
		res = match[0]
	}
	return res
}

// ReplaceBadSymbols replaces unusual symbols
func ReplaceBadSymbols(str string) string {
	var (
		mes [300][2]string = [300][2]string{
			{
				"S",
				"5",
			},
			{
				"é",
				"e",
			},
			{
				"A",
				"4",
			},
			{
				"I",
				"1",
			},
			{
				"¢",
				"с",
			},
			{
				"O",
				"0",
			},
			{
				"Q",
				"0",
			},
			{
				"J",
				"7",
			},
			{
				"C",
				"c",
			},
			{
				"|",
				"1",
			},
			{
				"$",
				"8",
			},
			{
				".",
				"",
			},
			{
				",",
				"",
			},
		}
	)

	for i := 0; i < len(mes); i++ {
		str = strings.Replace(str, mes[i][0], mes[i][1], -1)
	}

	return str
}
