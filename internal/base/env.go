package base

import (
	"os"
	"strconv"
)

// Env gets a environment variable, or returns a default value if it hasn't been installed
func Env(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultValue
}

// EnvI acts similar to Env, but also maps result to integer
func EnvI(key string, defaultValue int) int {
	val := Env(key, strconv.Itoa(defaultValue))
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return ret
}
