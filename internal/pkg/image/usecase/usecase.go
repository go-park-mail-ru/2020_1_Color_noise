package usecase

import (
	"2020_1_Color_noise/internal/pkg/board"
	"2020_1_Color_noise/internal/pkg/pin"
	userService "2020_1_Color_noise/internal/pkg/proto/user"
	"context"
	"encoding/csv"
	"io/ioutil"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

var table [][]string

func init() {
	data, err := ioutil.ReadFile("data.csv")
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

type ImageUsecase struct {
	repoPin   pin.IRepository
	repoBoard board.IRepository
	us        userService.UserServiceClient
}

func NewImageUsecase(repoPin pin.IRepository, repoBoard board.IRepository, us userService.UserServiceClient) *ImageUsecase {
	return &ImageUsecase{
		repoPin,
		repoBoard,
		us,
	}
}

func (im *ImageUsecase) Analyze(pinId uint, userId uint, image string) {
	cmd := exec.Command("python3", "analyze.py", "../storage/"+image)
	out, _ := cmd.CombinedOutput()

	i, err := strconv.Atoi(string(out))
	if err != nil {
		log.Println("Analazing error: bad return from script ", string(out))
	}

	if len(table) <= i {
		log.Println("Analazing error: bad read from csv file")
		return
	}

	err = im.repoPin.AddTags(pinId, table[i])
	if err != nil {
		log.Println("Analazing error: ", err)
		return
	}

	preferences := make(map[string]int)

	pins, err := im.repoPin.GetByUserID(userId, 0, 30)
	if err != nil {
		log.Println("Analazing error: ", err)
		return
	}

	for _, pin := range pins {
		for _, tag := range pin.Tags {
			preferences[tag] = preferences[tag] + 1
		}
	}

	boards, err := im.repoBoard.GetByUserID(userId, 0, 2)
	if err != nil {
		log.Println("Analazing error: ", err)
		return
	}

	if len(boards) != 0 {
		for _, p := range boards[0].Pins {
			if p.UserId != userId {
				for _, tag := range p.Tags {
					preferences[tag] = preferences[tag] + 2

				}
			}
		}
	} else {
		log.Println("Analazing error: boards is not found, userID: ", userId)
	}

	type kv struct {
		Key   string
		Value int
	}

	var kvs []kv
	for k, v := range preferences {
		kvs = append(kvs, kv{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})

	maxPref := 4
	if len(kvs) < maxPref {
		maxPref = len(kvs)
	}

	pref := make([]string, 0, maxPref)

	for i := 0; i < maxPref; i++ {
		pref = append(pref, kvs[i].Key)

	}

	_, err = im.us.UpdatePreferences(context.Background(), &userService.Pref{UserId: int32(userId), Preferences: pref})
	if err != nil {
		log.Println(err)
	}

}
