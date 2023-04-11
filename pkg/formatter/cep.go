package formatter

import "strings"

// Formatter cep input to standard XXXXX-XXX, e.g: 37540-000
func Cep(cep string) string {
	return strings.Join([]string{cep[:5], cep[5:]}, "-")
}

func SanitalizeCep(cep string) string {
	return strings.ReplaceAll(cep, "-", "")
}
