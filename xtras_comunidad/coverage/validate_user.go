package validate

import (
	"fmt"
	"regexp"
)

// Username valida un nombre de usuario. Solo permite
// nombres de usuario contener letras, números,
// y caracteres especiales "_", "-" y "."
func Username(u string) (bool, error) {
	if len(u) == 0 {
		return false, fmt.Errorf("username must be > 0 chars")
	}

	if len(u) > 30 {
		return false, fmt.Errorf("username too long (must be < 30 chars)")
	}

	validChars := regexp.MustCompile(`^[a-zA-Z1-9-_.]+$`)
	if !validChars.MatchString(u) {
		return false, fmt.Errorf("username contains invalid character")
	}

	return true, nil

}
