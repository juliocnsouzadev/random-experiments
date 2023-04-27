package main

import (
    "math/rand"
    "regexp"
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
    for _, c := range r.String() {
        switch c {
        case '[':
            // handle character class
            i := 1
            for r.String()[i] != ']' {
                i++
            }
            class := r.String()[2:i]\n\t\t	char := class[rand.Intn(len(class))]
            result = append(result, byte(char))
            r = regexp.MustCompile(r.String()[i+1:])
        case '\\':
            // handle escaped characters
            result = append(result, byte(r.String()[1]))
            r = regexp.MustCompile(r.String()[2:])
        default:
            // handle normal characters
            result = append(result, byte(c))
            r = regexp.MustCompile(r.String()[1:])
        }
    }

    return string(result), nil
}
