package utils

import (
	"strings"
)

func TrimStringArtist(s string) string {
	if idx := strings.Index(s, " - "); idx != -1 {
		return s[:idx]
	}
	return s
}

func TrimStringTrack(s string) string {
	idx1 := strings.Index(s, " - ")
	idx2 := strings.Index(s, " (")
	if idx1 != -1 && idx2 != -1 {
		return s[(idx1 + 3):idx2]
	} else if idx1 != -1 {
		return s[(idx1 + 3):]
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

func TrimTweet(s string) string {
	val := s
	if idx1 := strings.Index(val, "&"); idx1 != -1 {
		val = strings.ReplaceAll(val, "&", "")
	}
	if idx2 := strings.Index(val, " "); idx2 != -1 {
		val = strings.ReplaceAll(val, " ", "")
	}
	return val
}
