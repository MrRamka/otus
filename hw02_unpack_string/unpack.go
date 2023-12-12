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
		currentElement := inputStingRuneArray[i]
		nextElement := inputStingRuneArray[i+1]
		if currentElement == BACKSLASH &&
			(unicode.IsDigit(nextElement) || nextElement == BACKSLASH) {
			i++
		}
		// Обновляем значение, так как index мог поменятся
		currentElement = inputStingRuneArray[i]
		nextElement = inputStingRuneArray[i+1]

		if i < arrayLen-1 {
			if unicode.IsDigit(nextElement) {
				repeatCount, _ := strconv.Atoi(string(nextElement))
				stringBuilder.WriteString(strings.Repeat(string(currentElement), repeatCount))
				i += 2
				continue
			}
		}

		// Обновляем значение, так как index мог поменятся
		currentElement = inputStingRuneArray[i]
		stringBuilder.WriteRune(currentElement)
		i++
	}
	return stringBuilder.String(), nil
}
