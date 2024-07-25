package utils

import "regexp"

func IsValidPassword(password string) bool {
	lengthRegex := regexp.MustCompile(`^.{6,}$`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*()_+{}":;'?/.,<>]`)

	return lengthRegex.MatchString(password) &&
		uppercaseRegex.MatchString(password) &&
		specialCharRegex.MatchString(password) &&
		lowercaseRegex.MatchString(password) &&
		digitRegex.MatchString(password)
}
