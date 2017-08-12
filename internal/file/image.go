package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

// IsAnImage return true the the path seem like an image
func IsAnImage(path string) bool {

	fileExt := filepath.Ext(path)
	imageExt := [3]string{".jpg", ".jpeg", ".png"}
	for _, ext := range imageExt {

		if ext == strings.ToLower(fileExt) {
			return true
		}
	}
	return false
}

// DateStringFromImage Return the date path for an image path
func DateStringFromImage(fname string) (string, error) {
	if !IsAnImage(fname) {
		return "", errors.New("NOT_AN_IMAGE")
	}

	f, err := os.Open(fname)
	if err != nil {
		return "", errors.New("READ_FAILED")
	}

	x, err := exif.Decode(f)
	if err != nil {
		return "", errors.New("DECODE ERROR")
	}

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, dateError := x.DateTime()
	if dateError != nil {
		return "", errors.New("DATE ERROR")
	}

	return tm.Format("2006/01/02"), nil
}
