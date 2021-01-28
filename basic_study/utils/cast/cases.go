package cast

import (
	"strings"
)

const (
	snakeStateWaitUppercase  = -1
	snakeStateWaitUnderscore = -2
)

const expectedMaxWords = 10

var acronyms = map[string]struct{}{
	"JSON": {},
	"HTTP": {},
	"API":  {},
	"URL":  {},
	"ID":   {},
}

func isUppercase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func toLower(c byte) byte {
	return c + 'a' - 'A'
}

// ToSnake converts camel case string identifier to snake case.
func ToSnake(camel string) string {
	var b strings.Builder
	b.Grow(len(camel) + expectedMaxWords)
	s := snakeStateWaitUppercase
	for i := 0; i < len(camel); i++ {
		c := camel[i]
		if isUppercase(c) {
			if s >= 0 {
				if _, ok := acronyms[camel[s:i+1]]; ok {
					for j := s; j <= i; j++ {
						b.WriteByte(toLower(camel[j]))
					}
					s = snakeStateWaitUnderscore
					continue
				}
			}
			if s < 0 {
				if i > 0 {
					b.WriteByte('_')
				}
				s = i
			}
		} else {
			if i > 0 && s >= 0 {
				if i-1 > s {
					for j := s; j < i-1; j++ {
						b.WriteByte(toLower(camel[j]))
					}
					b.WriteByte('_')
				}
				b.WriteByte(toLower(camel[i-1]))
			}
			if s == snakeStateWaitUnderscore {
				b.WriteByte('_')
			}
			b.WriteByte(c)
			s = snakeStateWaitUppercase
		}
	}
	if s >= 0 {
		for j := s; j < len(camel); j++ {
			b.WriteByte(toLower(camel[j]))
		}
	}
	return b.String()
}
