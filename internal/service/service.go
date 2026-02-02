package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DataDetect(data string) (string, error) {
	trim := strings.TrimSpace(data)
	if trim == "" {
		return "", errors.New("empty input")
	}

	isMorse := true
	for _, r := range trim {
		switch r {
		case '.', '-', ' ', '\n', '\r', '\t':
		default:
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(trim), nil
	}

	for _, r := range trim {
		if unicode.IsLetter(r) && unicode.Is(unicode.Latin, r) {
			return "", errors.New("latin letters are not supported")
		}
	}

	return morse.ToMorse(trim), nil
}

func Reader(rawData multipart.File) (string, error) {
	defer func() {
		if err := rawData.Close(); err != nil {
			log.Printf("close uploaded file: %v", err)
		}
	}()
	b, err := io.ReadAll(rawData)
	if err != nil {
		return "", fmt.Errorf("read uploaded file: %w", err)
	}
	s := string(b)
	encodedData, err := DataDetect(s)
	if err != nil {
		return "", err
	}
	return encodedData, nil
}
