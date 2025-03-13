package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"topmusicstreaming/models"
	"regexp"
)

func TrimStringArtist(s string) string {
	if idx := strings.Index(s, " - "); idx != -1 {
		s = s[:idx]
	}

	separators := []string{" & ", " x ", " feat.", " ft."}

	for _, sep := range separators {
		if idx := strings.Index(s, sep); idx != -1 {
			s = s[:idx]
			break
		}
	}

	s = strings.TrimSpace(s)

	return s
}

func TrimStringTrack(s string) string {
	idx1 := strings.Index(s, " - ")
	if idx1 == -1 {
		return s
	}

	track := s[idx1+3:]

	re := regexp.MustCompile(`\s*\([^)]*\)`)
	track = re.ReplaceAllString(track, "")

	track = strings.TrimSpace(track)

	return track
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

func ensurePath(dir, file string) (string, error) {
	cwd, _ := os.Getwd()
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return "", err
		}
	}

	path := filepath.Join(cwd, dir, file)
	newFilePath := filepath.FromSlash(path)

	if _, err := os.Stat(newFilePath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(newFilePath)
		if err != nil {
			return "", err
		}
		file.Close()

		Logger.Infof("%s created successfully at %s \n", file.Name(), newFilePath)
	}

	return newFilePath, nil
}

func BuildFilePath(dir, name, ext string) (string, error) {
	file := name + "." + ext
	path, err := ensurePath(dir, file)
	if err != nil {
		return "", err
	}
	return path, nil
}

func WriteJSON(data models.Final, file string) error {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(file, dataFile, 0666); err != nil {
		return err
	}

	return nil
}

func BuildCollectorUrl(platform, country string) string {
	switch platform {
	case Spotify:
		return SpotifyBaseCollectorUri + country + "_daily" + HTMLEXT
	case AppleMusic:
		return AppleMusicCollectorBaseUri + country + HTMLEXT
	case Deezer:
		return DeezerCollectorBaseUri + country + HTMLEXT
	}
	return PlatformNotSupported
}
