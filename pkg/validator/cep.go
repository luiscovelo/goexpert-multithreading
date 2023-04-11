package validator

import (
	"regexp"
)

// Validate a raw cep, must to contains 8 digits, e.g: 37540000.
func Cep(cep string) bool {
	r := regexp.MustCompile(`^\d{8}$`)
	return r.MatchString(cep)
}
