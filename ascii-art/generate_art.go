package ascii_art

import (
	"strings"
)

func Generate(input string, bannerFile string) (string, error) {
	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, "\r\n", "\n")

	glyphMap, charHeight, err := LoadBanner(bannerFile)
	if err != nil {
		return "", err
	}

	result := RenderBanner(glyphMap, input, charHeight)
	return result, nil
}
