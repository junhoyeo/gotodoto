package parser

import (
	"regexp"
	"strings"
)

func Map(arr []string, callback func(string) string) []string {
	new := make([]string, len(arr))
	for i, v := range arr {
		new[i] = callback(v)
	}
	return new
}

func ToLines(content string) []string {
	trimmed := strings.TrimSpace(content)
	if len(trimmed) == 0 {
		return []string{}
	}
	lines := strings.Split(trimmed, "\n")
	lines = Map(lines, strings.TrimSpace)
	return lines
}

func Parse(content string) []string {
	lines := ToLines(content)
	var pattern string = "TODO: [^\n]*"

	return Map(lines, func(line string) string {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(line, -1)
		if len(matches) == 0 {
			return ""
		}
		return matches[0]
	})
}
