package usecase

import (
	"encoding/json"
	"github.com/mb-14/gomarkov"
	"io/ioutil"
	"log"
	"strings"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

var chain *gomarkov.Chain

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

	for _, i := range *tags {
			tokens := []string{i}

			next, _ := chain.Generate(tokens[(len(tokens) - 1):])

			println(next)

			tokens = append(tokens, next)

			str := strings.Join(tokens[0:len(tokens)-1], " ")
			generated = append(generated, str)
	}

	return &generated, nil
}

