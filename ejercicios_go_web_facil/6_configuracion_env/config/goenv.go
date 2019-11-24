package config

import "os"

func init() {
}

func GetEnvOrCreate(env, defVal string) string {
	if val := os.Getenv(env); val == "" {
		os.Setenv(env, defVal)
		return defVal
	} else {
		return val
	}
}

func GetStringEnv(env, defEnv string) string {
	val := GetEnv(env, defEnv)
	return val.(string)
}

func GetIntEnv(env string, defEnv int) int {
	val := GetEnv(env, defEnv)
	return val.(int)
}

func GetBoolEnv(env string, defEnv bool) bool {
	val := GetEnv(env, defEnv)
	return val.(bool)
}

func GetEnv(env string, defEnv interface{}) interface{} {
	if val := os.Getenv(env); val == "" {
		return defEnv
	} else {
		return val
	}
}
