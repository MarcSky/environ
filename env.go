package environ

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	failedToGetVar  = "Failed to get environment variable "
	failedToConvVar = "Failed to convert environment variable "
)

func IntGetEnv(name string, def int) int {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return i
}

func Int64GetEnv(name string, def int64) int64 {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	return i
}

func UInt64GetEnv(name string, def uint64) uint64 {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return def
	}
	return i
}

func FloatGetEnv(name string, def float64) float64 {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	i, err := strconv.ParseFloat(value, 10)
	if err != nil {
		return def
	}
	return i
}

func StrGetEnv(name, def string) string {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	return value
}

func BoolGetEnv(name string, def bool) bool {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	i, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return i
}

func TimeDurationEnv(name string, defaultVal string) time.Duration {
	result, ok := os.LookupEnv(name)
	if !ok {
		result = defaultVal
	}

	d, err := time.ParseDuration(result)
	if err != nil {
		d, _ = time.ParseDuration(defaultVal)
	}
	return d
}

func MustGetString(name string) string {
	value := StrGetEnv(name, "")
	if value == "" {
		log.Panic(failedToGetVar, name)
	}
	return value
}

func MustGetInt(name string) int {
	value := os.Getenv(name)
	if value == "" {
		log.Panic(failedToGetVar, name)
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Panic(failedToConvVar, name)
	}
	return i
}

func MustGetInt64(name string) int64 {
	value := os.Getenv(name)
	if value == "" {
		log.Panic(failedToGetVar, name)
	}
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Panic(failedToConvVar, name)
	}
	return i
}

func MustGetUInt64(name string) uint64 {
	value := os.Getenv(name)
	if value == "" {
		log.Panic(failedToGetVar, name)
	}
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Panic(failedToConvVar, name)
	}
	return i
}

func MustGetFloat(name string) float64 {
	value := os.Getenv(name)
	if value == "" {
		log.Panic(failedToGetVar, name)
	}
	i, err := strconv.ParseFloat(value, 10)
	if err != nil {
		log.Panic(failedToConvVar, name)
	}
	return i
}

func MustGetBool(name string) bool {
	value := os.Getenv(name)
	if value == "" {
		log.Panic(failedToGetVar, name)
	}
	i, err := strconv.ParseBool(value)
	if err != nil {
		log.Panic(failedToConvVar, name)
	}
	return i
}

func MustGetTimeDuration(name string) time.Duration {
	result, ok := os.LookupEnv(name)
	if !ok {
		log.Panic(failedToGetVar, name)
	}
	d, err := time.ParseDuration(result)
	if err != nil {
		log.Panic(failedToConvVar, name)
	}
	return d
}

