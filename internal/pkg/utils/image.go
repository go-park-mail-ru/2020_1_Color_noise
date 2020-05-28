package utils

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"os"
)

func SaveImage(b *[]byte) (string, error) {
	name := RandStringRunes(30) + ".jpg"
	path := "../storage/" + name
	file, err := os.Create(path)
	if err != nil {
		return "", Wrap(err, "Creating image error")
	}

	defer file.Close()

	_, err = file.Write(*b)
	if err != nil {
		return "", Wrap(err, "Saving image error")
	}

	return name, nil
}
