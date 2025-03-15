package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	reg := regexp.MustCompile(`(?i)\[TRC]|\[ERR]|\[INF]|\[DBG]|\[WRN]|\[FTL]`)
	result := reg.FindIndex([]byte(text[:]))
	if result != nil {
		if result[0] == 0 {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	reg := regexp.MustCompile(`<[*=~-]*>`)
	result := reg.Split(text, -1)

	return result
}

func CountQuotedPasswords(lines []string) int {
	reg := regexp.MustCompile(`(?i)".*password"`)
	var count int

	for _, line := range lines {
		if reg.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	reg := regexp.MustCompile(`end-of-line[0-9]*`)
	result := reg.Split(text, -1)

	var log string
	for _, line := range result {
		log += line
	}
	return log
}

func TagWithUserName(lines []string) []string {
	reg := regexp.MustCompile(`User\s+`)

	for i, line := range lines {
		//prefix := ""
		if reg.MatchString(line) {
			split := reg.Split(line, -1)
			userName := strings.Split(split[1], " ")[0]
			fmt.Println(userName)

			lines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		}
	}
	return lines
}

//
//func TagWithUserName(lines []string) []string {
//	reg := regexp.MustCompile(`User\s+`)
//
//	for in, line := range lines {
//		if reg.MatchString(line) {
//			var response = "[USR] "
//			split := strings.Split(line, " ")
//			for i, tag := range split {
//				fmt.Println(tag)
//				if reg.MatchString(tag) {
//					response += split[i+1]
//				}
//			}
//			lines[in] = response + line
//		} else {
//			continue
//		}
//	}
//	return lines
//}
