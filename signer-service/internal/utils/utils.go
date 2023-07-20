package utils

import (
	"strconv"
)

func MatchProviderID(fromPath string, fromToken int) bool {
	if fromPath == "" || fromPath == "0" {
		return true
	}

	if fromToken == 0 {
		return true
	}

	fromPathInt, err := strconv.Atoi(fromPath)
	if err != nil {
		return false
	}

	return fromPathInt == fromToken
}
