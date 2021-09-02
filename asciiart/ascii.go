package AsciiArt

import (
	"errors"
	"os"
	"strings"
)

type sign struct {
	color   string
	picture [8][]rune
	style   string
}

// GetASCII returns word in ascii-art
func GetASCII(word, FontStyle string) (string, error) {
	template, err := PreparingTemplate(FontStyle)
	if err != nil {
		return "", err
	}
	words := strings.Split(word, "\n")
	res := ""
	//building sign
	for _, word := range words {
		signs := []*sign{}
		for _, runes := range word {
			newSign := sign{picture: template[runes-32]}
			signs = append(signs, &newSign)
		}
		if len(signs) != 0 {
			res += getOneLine(signs)
		}
	}
	return res, nil
}

func getOneLine(signs []*sign) string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, Sign := range signs {
			res += string(Sign.picture[i])
		}
		res += "\n"
	}
	return res
}

//PreparingTemplate is preparing template for translator
func PreparingTemplate(filename string) ([95][8][]rune, error) {
	newBuf, err := os.ReadFile("./asciiart/templates/" + filename + ".txt")
	if err != nil {
		return [95][8][]rune{}, errors.New("There is no such FontStyle")
	}
	newBuf = []byte(strings.ReplaceAll(string(newBuf), string(rune(13)), ""))
	GeneralLen := len(newBuf)
	res := [95][8][]rune{}
	start, index, queue := 1, 1, 0
	tempLen := 0
	for ; index < GeneralLen; index++ {
		tempLen = 0
		row := 0
		//reading symbol
		for queue < 95 && index < GeneralLen && row < 8 {
			//checking symbol before appending
			if newBuf[index] == 10 {
				if row > 0 {
					if index-start != tempLen {
						return [95][8][]rune{}, errors.New("FontStyle is damaged")
					}
				} else {
					tempLen = index - start
				}
				res[queue][row] = []rune(string(newBuf[start:index]))
				row++
				start = index + 1
			}
			index++
		}
		queue++
		start++
	}

	if index-start != 0 {
		if index-start != tempLen {
			return [95][8][]rune{}, errors.New("FontStyle is damaged")
		}
		res[94][7] = []rune(string(newBuf[start-1 : index]))
	}

	if queue != 95 {
		return [95][8][]rune{}, errors.New("FontStyle is damaged")
	}

	return res, nil
}
