package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conv(s string) (string, error) {

	s = strings.TrimSpace(s)

	if s == "" {
		return "", fmt.Errorf("пустой ввод")
	}
	var out string
	if isMorse(s) {

		out = morse.ToText(s)

	} else {
		out = morse.ToMorse(s)

	}
	return out, nil
}

func isMorse(s string) bool {

	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return false
		}
	}
	return true

}
