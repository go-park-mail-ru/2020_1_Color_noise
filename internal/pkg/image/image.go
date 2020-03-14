package image

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/utils"
	"os"
)

func SaveImage(b *[]byte) (string, error) {
	name := utils.RandStringRunes(30) + ".jpg"
	path := "static/" + name
	file, err := os.Create(path)
	if err != nil{
		return "", Wrap(err, "Creating image error")
	}

	defer file.Close()

	_, err = file.Write(*b)
	if err != nil {
		return "", Wrap(err, "Saving image error")
	}

	return name, nil
}