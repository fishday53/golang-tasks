//go:build !solution

package mycheck

import (
	"errors"
	"fmt"
	"unicode"
)

type Errs []error

func Error(e Errs) error {
	var result error
	if len(e) > 0 {
		result = e[0]
		for i := 1; i < len(e); i++ {
			result = fmt.Errorf("%v;%v", result, e[i])
		}
	}
	return result
}

func MyCheck(input string) error {
	var e Errs
	var digits bool
	var spaces int
	for _, char := range input {
		if unicode.IsDigit(char) {
			digits = true
		}
		if unicode.IsSpace(char) {
			spaces++
		}
	}
	if digits {
		e = append(e, errors.New("found numbers"))
	}
	if len(input) >= 20 {
		e = append(e, errors.New("line is too long"))
	}
	if spaces != 2 {
		e = append(e, errors.New("no two spaces"))
	}
	return Error(e)
}
