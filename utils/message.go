package utils

import "fmt"

func Alpha(field string) string {
	return fmt.Sprintf("O campo %s deve conter apenas letras", field)
}

func AlphaNum(field string) string {
	return fmt.Sprintf("O campo %s deve conter apenas letras e números", field)
}

func Required(field string) string {
	return fmt.Sprintf("O campo %s é obrigatário", field)
}

func Invalid(field string) string {
	return fmt.Sprintf("O campo %s é inválido", field)
}

func MaxLength(field string, length string) string {
	return fmt.Sprintf("O campo %s deve ter no máximo %s caracteres", field, length)
}

func MinLength(field string, length string) string {
	return fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres", field, length)
}

func Email() string {
	return Invalid("email")
}
