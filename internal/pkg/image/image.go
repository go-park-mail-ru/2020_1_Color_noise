package image

import (
	"os"
	. "pinterest/internal/pkg/error"
	"pinterest/internal/pkg/utils"
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
