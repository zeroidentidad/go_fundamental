package config

import (
	"fmt"
	"os"
)

func GetHostame() string {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("no Hostame: %v\n", err)
		return ""
	}
	return name
}
