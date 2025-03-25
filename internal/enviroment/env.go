package enviroment

// some helper functions i've used in the past
// for simple environment variables lookups

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func String(name string, defaultValue ...string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}
	return val
}

func StringSlice(name, delimiter string, defaultValue ...string) []string {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue
	}

	return strings.FieldsFunc(val, func(r rune) bool {
		return r == ','
	})
}

func Int(name string, defaultValue ...int) int {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(name + " is not a valid int: " + err.Error())
	}
	return intVal
}

func Int32(name string, defaultValue ...int32) int32 {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	intVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		panic(name + " is not a valid int64: " + err.Error())
	}
	return int32(intVal)
}

func Int64(name string, defaultValue ...int64) int64 {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(name + " is not a valid int64: " + err.Error())
	}
	return intVal
}

func Float(name string, defaultValue ...float64) float64 {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(name + " is not a valid float64: " + err.Error())
	}
	return floatVal
}

func Float32(name string, defaultValue ...float32) float32 {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	floatVal, err := strconv.ParseFloat(val, 32)
	if err != nil {
		panic(name + " is not a valid float32: " + err.Error())
	}
	return float32(floatVal)
}

func Duration(name string, defaultValue ...time.Duration) time.Duration {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	duration, err := time.ParseDuration(val)
	if err != nil {
		panic(name + " is not a valid duration: " + err.Error())
	}
	return duration
}

func Bool(name string, defaultValue ...bool) bool {
	val, ok := os.LookupEnv(name)
	if !ok {
		if !hasDefault(defaultValue) {
			panic(name + " is required")
		}
		return defaultValue[0]
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		panic(name + " is not a valid duration: " + err.Error())
	}
	return boolVal
}

func hasDefault(value any) bool {
	if value == nil {
		return false
	}
	return reflect.ValueOf(value).Len() > 0
}
