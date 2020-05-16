package main

import (
	. "2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/utils"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var table [][]string

func init() {
	data, err := ioutil.ReadFile("./internal/pkg/image/data.csv")
	if err != nil {
		log.Println("data.csv is not found ", err.Error())
		return
	}
	r := csv.NewReader(strings.NewReader(string(data)))
	r.Comma = ';'
	for {
		line, err := r.Read()
		if err != nil {
			break
		}
		table = append(table, line)
	}

}

func Analyze(image string) {
	cmd := exec.Command("python3", "./internal/pkg/image/analyze.py", "../storage/" + image)
	out, _ := cmd.CombinedOutput()
	i, err := strconv.Atoi(string(out))
	if err != nil {
		log.Println("bad return from script")
	}
	if len(table) < i {
		log.Println("bad read from csv file")
		return
	}
	fmt.Println(table[i])
}

func SaveImage(b *[]byte) (string, error) {
	name := utils.RandStringRunes(30) + ".jpg"
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
