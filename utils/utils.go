package utils

import (
	"strings"
)

func TrimStringTrack(s string) string {
	if idx := strings.Index(s, " ("); idx != -1 {
		return s[:idx]
	}
	return s
}

func TrimStringArtist(s string) string {
	if idx := strings.Index(s, ","); idx != -1 {
		return s[:idx]
	}
	return s
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
