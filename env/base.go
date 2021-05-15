package env

import (
    "os"
    "strconv"
)

func GetEnvStr(key string, defaultValue string) string {
    v := os.Getenv(key)
    if v == "" {
        return defaultValue
    } else {
        return v
    }
}

func GetEnvInt(key string, defaultValue int) int {
    v := os.Getenv(key)
    if n, err := strconv.Atoi(v); err != nil {
        return defaultValue
    } else {
        return n
    }
}
