package usecase

import (
	"encoding/json"
	"github.com/mb-14/gomarkov"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

var chain *gomarkov.Chain

var Smile = []string{"❤", "💥","😁","🤩", "🎈", "🌸"}

func init(){
	chain = gomarkov.NewChain(1)
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Fatal("no data")
	}

	err = json.Unmarshal(data, &chain)
	if err != nil {
		log.Fatal("incorrect data")
	}



}

func (us* Usecase) Predict(tags *[]string) (*[]string, error){

	chain := gomarkov.Chain{Order:1}
	data, _ := ioutil.ReadFile("model.json")
	 json.Unmarshal(data, &chain)

	var generated []string

	m := make(map[string]bool)

	if len(*tags) < 3 {
		return nil, nil
	}

	for i := 0; i < 4; i++ {
		tokens := []string{(*tags)[3]}
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])

		if next == gomarkov.EndToken {
			next = Smile[rand.Intn(len(Smile))]
		}
		tokens = append(tokens, next)

		str := strings.Join(tokens, " ")

		m[str] = true
	}

	for key, _ := range m {
		generated = append(generated, key)
	}

	return &generated, nil
}

