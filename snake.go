package anycase

import (
	"strings"
)

// ToSnake converts a string to snake_case
func ToSnake(s string) string {

	return ToDelimited(s, '_')
}

// ToSnakeAndIgnore converts to snake_case and ignores the following characters arg
func ToSnakeAndIgnore(s string, ignore uint8) string {

	return ToDelimitedUppercase(s, '_', ignore, false)
}

// ToSnakeUppercase converts a string to UPPERCASE_SNAKE_CASE
func ToSnakeUppercase(s string) string {
	return ToDelimitedUppercase(s, '_', 0, true)
}

// ToKebab converts a string to kebab-case
func ToKebab(s string) string {
	return ToDelimited(s, '-')
}

// ToKebabUppercase converts a string to UPPERCASE-KEBAB-CASE
func ToKebabUppercase(s string) string {
	return ToDelimitedUppercase(s, '-', 0, true)
}

// ToDelimited converts a string to delimited.snake.case
// (in this case `delimiter = '.'`)
func ToDelimited(s string, delimiter uint8) string {
	return ToDelimitedUppercase(s, delimiter, 0, false)
}

// ToDelimitedUppercase converts a string to UPPERCASE.DELIMITED.SNAKE.CASE
// (in this case `delimiter = '.'; uppercase = true`)
// or delimited.snake.case
// (in this case `delimiter = '.'; uppercase = false`)
func ToDelimitedUppercase(s string, delimiter uint8, ignore uint8, uppercase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			vIsCap := v >= 'A' && v <= 'Z'
			vIsLow := v >= 'a' && v <= 'z'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			if (vIsCap && nextIsLow) || (vIsLow && nextIsCap) {
				nextCaseIsChanged = true
			}
			if ignore > 0 && i-1 >= 0 && s[i-1] == ignore && nextCaseIsChanged {
				nextCaseIsChanged = false
			}
		}

		if i > 0 && n[len(n)-1] != delimiter && nextCaseIsChanged {
			// add underscore if next letter case type is changed
			if v >= 'A' && v <= 'Z' {
				n += string(delimiter) + string(v)
			} else if v >= 'a' && v <= 'z' {
				n += string(v) + string(delimiter)
			}
		} else if v == ' ' || v == '_' || v == '-' {
			// replace spaces/underscores with delimiters
			if uint8(v) == ignore {
				n += string(v)
			} else {
				n += string(delimiter)
			}
		} else {
			n = n + string(v)
		}
	}

	if uppercase {
		n = strings.ToUpper(n)
	} else {
		n = strings.ToLower(n)
	}
	return n
}
