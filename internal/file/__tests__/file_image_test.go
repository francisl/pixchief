package file_test

import (
	"testing"

	"github.com/francisl/pixchief/internal/file"
)

func TestIsAnImage(t *testing.T) {
	isAnImage := file.IsAnImage("test.txt")
	if isAnImage == true {
		t.Errorf(".txt should not be an image file")
	}
}

func TestIsAnImage_shouldReturnTrueForJpg(t *testing.T) {
	isAnImage := file.IsAnImage("test.jpg")
	if isAnImage != true {
		t.Errorf(".jpg should be an image file")
	}

	isAnImage = file.IsAnImage("test.jpeg")
	if isAnImage != true {
		t.Errorf(".jpeg should be an image file")
	}

	isAnImage = file.IsAnImage("test.png")
	if isAnImage != true {
		t.Errorf(".jpg should be an image file")
	}
}

func TestDateStringFromImage_shouldGetEnglishDate(t *testing.T) {
	date, err := file.DateStringFromImage("./photo1.jpg")
	expectedDate := "2017/08/01"
	if err != nil {
		t.Errorf("Date Was not correctly read from image metadata")
	} else if date != expectedDate {
		t.Errorf("Invalid date format return %s is not %s", date, expectedDate)
	}

	_, err = file.DateStringFromImage("./photo_nodate.jpg")
	if err == nil {
		t.Errorf("Date should not be read from image metadata")
	}
}

func TestDateStringFromImage_shouldReturnError_WhenNotAnImage(t *testing.T) {
	_, err := file.DateStringFromImage("./notanimage.txt")

	if err == nil {
		t.Errorf("Should Return NOT_AN_IMAGE")
	} else if err.Error() != "NOT_AN_IMAGE" {
		t.Errorf("Return Invalid Error")
	}
}

func TestDateStringFromImage_shouldReturnError_WhenCannotRead(t *testing.T) {
	_, err := file.DateStringFromImage("./notafile.png")

	if err == nil {
		t.Errorf("Should Return NOT_AN_IMAGE")
	} else if err.Error() != "READ_FAILED" {
		t.Errorf("Return Invalid Error")
	}
}
