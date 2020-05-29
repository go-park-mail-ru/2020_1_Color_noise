package utils

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"bytes"
	"io/ioutil"
	"os"
)

func SaveImage(buffer *bytes.Buffer) (string, error) {
	if buffer == nil {
		return "", New("Empty buffer")
	}

	name := RandStringRunes(30) + ".jpg"
	path := "../storage/" + name
	if err := ioutil.WriteFile(path, buffer.Bytes(), 0666); err != nil {
		return "", Wrap(err, "Saving image error")
	}

	return name, nil
}

func DeleteImage(name string) (error) {
	path := "../storage/" + name

	err := os.Remove(path)
	if err != nil {
		return Wrap(err, "Deleting image error")
	}

	return nil
}
