package ascii_art

import (
	"strings"
)

func RenderBanner(glyphMap map[rune][]string, input string) string {
	if input == "\n" {
		return input
	}
	if input == "" {
		return ""
	}

	var s strings.Builder

	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			s.WriteRune('\n')
			continue
		}
		for i := range 8 {
			for _, ch := range line {
				if ch < 32 || ch > 126 {
					ch = ' '
				}
				s.WriteString(glyphMap[ch][i])
			}
			s.WriteRune('\n')
		}
	}
	return s.String()
}
