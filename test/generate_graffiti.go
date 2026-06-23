package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("mod.txt")
	if err != nil {
		panic(err)
	}
	bannerLines := strings.Split(string(data), "\n")

	m := map[rune][]string{}
	for char := ' '; char <= 'a'; char++ {
		start := (char - 32) * 6
		m[char] = bannerLines[start : start+6]
	}

	for char := ' '; char <= '~'; char++ {
		fmt.Print(string(char))
		pad := strings.Repeat(" ", len(m[char][1]))
		slice := []string{}
		slice = append(slice, "\n")
		slice = append(slice, pad)
		slice = append(slice, m[char]...)
		slice = append(slice, pad)
		//slice = append(slice, pad)

		m[char] = slice

	}

	var s strings.Builder
	for char := ' '; char <= '~'; char++ {
		for _, eachIndex := range m[char] {
			s.WriteString(eachIndex)
			s.WriteRune('\n')
		}

	}

	fmt.Print(strings.ReplaceAll(s.String(), "\n\n", "\n"))

}
