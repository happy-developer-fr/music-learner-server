package utils

import (
	"errors"
	"os"
	"strconv"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string, ifEmptyValue string) (string, error) {
	v := os.Getenv(k)
	var error error = nil
	if v == "" {
		if ifEmptyValue == "" {
			error = errors.New("ENV missing, key: " + k)
		} else {
			return ifEmptyValue, error
		}
	}
	return v, error
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string, ifEmptyValue bool) (bool, error) {
	v := os.Getenv(k)
	var error error = nil
	var b bool = true
	if v == "" {
		return ifEmptyValue, nil
	} else {
		b, error = strconv.ParseBool(v)
	}
	return b, error
}
