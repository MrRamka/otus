package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const BACKSLASH = '\\'

func Unpack(str string) (string, error) {
	inputStingRuneArray := []rune(str)

	arrayLen := len(inputStingRuneArray)

	if arrayLen == 0 {
		return "", nil
	}

	// Проверка строки на правльное экранирование и расположение цифр относительно букв
	/*
		[^\\\d]+(\d?) - проверка что перед числом идет не слэш и не число
		(\\+\d+) - проверка что после слэша идет число
	*/
	matchedSingleDigitAndBackslash, _ := regexp.MatchString(`^(([^\\\d]+(\d?))+|(\\+\d+)+)+$`, str)
	if !matchedSingleDigitAndBackslash {
		return "", ErrInvalidString
	}

	var stringBuilder strings.Builder

	i := 0
	for i < arrayLen {
		if inputStingRuneArray[i] == BACKSLASH &&
			(unicode.IsDigit(inputStingRuneArray[i+1]) || inputStingRuneArray[i+1] == BACKSLASH) {
			i++
		}

		if i < arrayLen-1 {
			if unicode.IsDigit(inputStingRuneArray[i+1]) {
				repeatCount, _ := strconv.Atoi(string(inputStingRuneArray[i+1]))
				stringBuilder.WriteString(strings.Repeat(string(inputStingRuneArray[i]), repeatCount))
				i += 2
				continue
			}
		}

		stringBuilder.WriteRune(inputStingRuneArray[i])
		i++
	}
	return stringBuilder.String(), nil
}
