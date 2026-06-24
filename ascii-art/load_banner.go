package ascii_art

import (
	"os"
	"strings"
)

func LoadBanner(bannerFile string) (map[rune][]string, int, error) {
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, 0, err
	}
	stringFile := strings.ReplaceAll(string(file), "\r\n", "\n")
	bannerLines := strings.Split(stringFile, "\n")
	bannerLines = bannerLines[1:]
	glyphMap := make(map[rune][]string)

	charLenght := 9

	name := strings.TrimPrefix(bannerFile, "banners/")
	name = strings.TrimSuffix(name, ".txt")
	switch name {
	case "acrobat":
		charLenght = 13
	case "graceful", "miniwi":
		charLenght = 7
	case "temper":
		charLenght = 6
	}
	var actualHeight = charLenght - 1

	for char := ' '; char <= '~'; char++ {
		start := int((char - 32)) * charLenght
		glyphMap[char] = bannerLines[start : start+actualHeight]
	}

	return glyphMap, actualHeight, nil
}
