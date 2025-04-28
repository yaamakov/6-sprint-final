// Пакет service реализует функционал автоматического определения кода Морзе или обычного текста из переданной строки.

package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var (
	ErrInvalidCharacter = errors.New("invalid character in input text")
	ErrEmptyInput       = errors.New("empty input")
)

func DetectAndConvert(input string) (string, error) {
	if input == "" {
		return "", ErrEmptyInput
	}

	if IsMorse(input) {
		return MorseToText(input)
	}
	return TextToMorse(input)
}

func TextToMorse(text string) (string, error) {
	if text == "" {
		return "", ErrEmptyInput
	}

	converter := morse.NewConverter(
		morse.DefaultMorse,
		morse.WithLowercaseHandling(true),
		// morse.WithHandler(func(err error) string {
		// 	return ""
		// }),
	)

	for _, char := range strings.ToUpper(text) {
		if _, ok := morse.DefaultMorse[char]; !ok && char != ' ' && !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", char) {
			return "", errors.New("invalid character: " + string(char))
		}
	}

	result := converter.ToMorse(text)
	return result, nil
}

func MorseToText(morseString string) (string, error) {
	if morseString == "" {
		return "", ErrEmptyInput
	}

	for _, char := range morseString {
		if char != '.' && char != '-' && char != '/' && char != ' ' {
			return "", errors.New("invalid morse code character: " + string(char))
		}
	}

	converter := morse.DefaultConverter

	result := converter.ToText(morseString)

	return result, nil
}

func IsMorse(input string) bool {
	for _, char := range input {
		if char != '.' && char != '-' && char != '/' && char != ' ' {
			return false
		}
	}

	return strings.ContainsAny(input, ".-")
}
