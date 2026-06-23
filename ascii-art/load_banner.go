package ascii_art

import (
	"os"
	"strings"
)

func LoadBanner(bannerFile string) (map[rune][]string, error) {
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}
	stringFile := strings.ReplaceAll(string(file), "\r\n", "\n")
	bannerLines := strings.Split(stringFile, "\n")
	bannerLines = bannerLines[1:]
	glyphMap := make(map[rune][]string)

	for char := ' '; char <= '~'; char++ {
		start := (char - 32) * 9
		glyphMap[char] = bannerLines[start : start+8]
	}

	return glyphMap, nil
}
