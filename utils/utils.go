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

func TrimStringCoverAppleMusic(s string) string {
	idx1 := strings.Index(s, "40w, ")
	idx2 := strings.Index(s, " 80w")
	if idx1 != -1 && idx2 != -1 {
		return s[(idx1 + 5):idx2]
	}
	return s
}

func TrimStringCoverDeezer(s string) string {
	idx1 := strings.Index(s, "(")
	idx2 := strings.Index(s, ")")
	if idx1 != -1 && idx2 != -1 {
		return s[(idx1 + 1):idx2]
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
