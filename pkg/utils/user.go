package utils

import "regexp"

func IsValidPassword(password string) bool {
	lengthRegex := regexp.MustCompile(`^.{6,}$`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*()_+{}":;'?/.,<>]`)

	return lengthRegex.MatchString(password) &&
		uppercaseRegex.MatchString(password) &&
		specialCharRegex.MatchString(password)
}
