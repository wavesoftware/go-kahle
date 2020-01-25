package base

import (
	"os"
	"strconv"
)

func Env(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	} else {
		return defaultValue
	}
}

func EnvI(key string, defaultValue int) int {
	val := Env(key, strconv.Itoa(defaultValue))
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return ret
}
