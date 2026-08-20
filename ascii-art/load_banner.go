package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(bannerFile string) (map[rune][]string, int, error) {
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, 0, err
	}
	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	if len(lines) < 2 || lines[0] != "" {
		return nil, 0, fmt.Errorf("invalid banner header")
	}
	lines = lines[1:]
	if len(lines) == 0 || len(lines)%95 != 0 {
		return nil, 0, fmt.Errorf("invalid banner dimensions")
	}
	recordHeight := len(lines) / 95
	if recordHeight < 2 {
		return nil, 0, fmt.Errorf("invalid banner height")
	}
	height := recordHeight - 1
	glyphMap := make(map[rune][]string, 95)
	for i, char := 0, rune(' '); i < 95; i, char = i+1, char+1 {
		start := i * recordHeight
		glyphMap[char] = lines[start : start+height]
	}
	return glyphMap, height, nil
}
