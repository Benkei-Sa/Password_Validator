package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	var result bool
	var a int
	for _, n := range s {
		if unicode.IsDigit(n) == true {
			a++
			break
		}
	}
	if a == 1 {
		result = true
	} else {
		result = false
	}
	return result
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	var result bool
	var a int
	for _, n := range s {
		if unicode.IsLetter(n) == true {
			a++
			break
		}
	}
	if a == 1 {
		result = true
	} else {
		result = false
	}
	return result
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	var validator = func(s string) bool {
		var result bool
		if utf8.RuneCountInString(s) == length {
			result = true
		} else {
			result = false
		}
		return result
	}
	return validator
}

// and возвращает валидатор, который проверяет, что все
// переданные ему валидаторы вернули true
func and(funcs ...validator) validator {
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) == false {
				return false
				break
			}
		}
		return true
	}
}

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
func or(funcs ...validator) validator {
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) == true {
				return true
				break
			}
		}
		return false
	}
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	result := p.validator(p.value)
	return result
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	var s string
	fmt.Scan(&s)
	// валидатор, который проверяет, что пароль содержит буквы и цифры,
	// либо его длина не менее 10 символов
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
