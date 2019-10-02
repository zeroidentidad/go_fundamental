package extras

import (
	"os"
	"strings"
)

// ParseDirsEnv dt
func ParseDirsEnv(dir string) string {
	return strings.NewReplacer(
		"$HOME", os.Getenv("HOME"),
	).Replace(dir)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
