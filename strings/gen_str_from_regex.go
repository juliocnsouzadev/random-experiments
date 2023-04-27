package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// GenerateRandomStringFromRegex generates a random string from a regular expression
func GenerateRandomStringFromRegex(regex string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}

	var result []byte
	expr := r.String()

	for len(expr) > 0 {
		switch expr[0] {
		case '[':
			// handle character class
			i := 1
			for expr[i] != ']' {
				i++
			}
			class := expr[2:i]
			char := class[rand.Intn(len(class))]
			result = append(result, byte(char))
			expr = expr[i+1:]

		case '\\':
			// handle escaped characters
			result = append(result, byte(expr[1]))
			expr = expr[2:]

		case '(':
			// handle capturing groups
			i := strings.IndexByte(expr, ')')
			if i == -1 {
				return "", regexp.ErrMissingParen
			}
			group := expr[1:i]
			subR, err := regexp.Compile(group)
			if err != nil {
				return "", err
			}
			subStr, err := GenerateRandomStringFromRegex(subR)
			if err != nil {
				return "", err
			}
			result = append(result, []byte(subStr)...)
			expr = expr[i+1:]

		case '|':
			// handle alternation
			return "", regexp.ErrInternalError

		default:
			// handle normal characters
			result = append(result, byte(expr[0]))
			expr = expr[1:]
		}
	}

	return string(result), nil
}
