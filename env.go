package env

import (
	"os"
	"strconv"
	"strings"
)

// Set env
func Set(key string, vals ...string) error {
	var val string
	for _, v := range vals {
		val = v
		break
	}
	return os.Setenv(key, val)
}

// Raw to read environment key
func Raw(name string, defs ...string) string {
	var def string
	for _, d := range defs {
		def = d
		break
	}
	if str := os.Getenv(name); str != "" {
		return str
	}
	return def
}

// Multi to read environment key and return value in array format
func Multi(name, separator string, defs ...string) []string {
	s := os.Getenv(name)
	if s != "" {
		return strings.Split(s, separator)
	}
	return defs
}

// Int to read environment key and return value in int format
func Int(name string, defs ...int) int {
	var def int
	for _, d := range defs {
		def = d
		break
	}
	v := Raw(name)
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}

// I32 to read environment key and return value in int32 format
func I32(name string, defs ...int32) int32 {
	var def int32
	for _, d := range defs {
		def = d
		break
	}
	v := Raw(name)
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return def
	}
	return int32(i)
}

// I64 to read environment key and return value in int64 format
func I64(name string, defs ...int64) int64 {
	var def int64
	for _, d := range defs {
		def = d
		break
	}
	v := Raw(name)
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return def
	}
	return i
}

// Bool to read environment key and return value in bool format
func Bool(name string, defs ...bool) bool {
	var def bool
	for _, d := range defs {
		def = d
		break
	}
	switch v := strings.ToLower(Raw(name)); {
	case v == "true":
		return true
	case v == "false":
		return false
	default:
		return def
	}
}
