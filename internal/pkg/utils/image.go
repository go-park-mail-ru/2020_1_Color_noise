package utils

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"bytes"
	"image"
	"io/ioutil"
	"log"
	"os"
	"github.com/chai2010/webp"
)

func SaveImage(buffer *bytes.Buffer) (string, error) {
	if buffer == nil {
		return "", New("Empty buffer")
	}

	img, _, err := image.Decode(buffer)
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	if err = webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Println(err)
	}

	name := RandStringRunes(30) + ".webp"
	path := "../storage/" + name
	if err = ioutil.WriteFile(path, buf.Bytes(), 0666); err != nil {
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
