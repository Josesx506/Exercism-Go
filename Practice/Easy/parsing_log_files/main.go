package main

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	// panic("Please implement the IsValidLine function")
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	ok := re.MatchString(text)
	return ok
}

func SplitLogLine(text string) []string {
	// panic("Please implement the SplitLogLine function")
	re, _ := regexp.Compile(`(\<[~*=-]+\>|<>)`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	// panic("Please implement the CountQuotedPasswords function")
	total := 0

	re, _ := regexp.Compile(`".*[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd].*"`)

	for _, text := range lines {
		if re.MatchString(string(text)) {
			total++
		}
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	// panic("Please implement the RemoveEndOfLineText function")
	re, _ := regexp.Compile(`end-of-line\S+`)
	clean := re.ReplaceAllString(text, "")
	return clean
}

func TagWithUserName(lines []string) []string {
	// panic("Please implement the TagWithUserName function")
	re, _ := regexp.Compile(`User\s+(\S+)`)
	for i, text := range lines {
		sl := re.FindStringSubmatch(text)
		if len(sl) > 1 {
			lines[i] = "[USR] " + sl[1] + " " + text
			// fmt.Println("[USR] " + sl[1] + " " + text + "\n")
		}
	}
	return lines
}

func main() {
	fmt.Println(IsValidLine("[ERR] A good error here"))
	fmt.Println(IsValidLine("Any old [ERR] text"))
	fmt.Println(SplitLogLine("section 1<*>section 2<~~~>section 3"))

	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	fmt.Println(CountQuotedPasswords(lines))
	fmt.Println(RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27"))
	fmt.Println(TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	}))
}
