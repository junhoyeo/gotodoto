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
	parsed := make([]string, 0, len(lines))

	var pattern string = "TODO: [^*/\n]*"
	re := regexp.MustCompile(pattern)

	var index uint64 = 0
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		matched := strings.TrimSpace(matches[0])
		parsed = append(parsed, matched)
		index++
	}
	return parsed
}
