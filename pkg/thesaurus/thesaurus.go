package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

// BigHuge : API handler, initialize with API key
type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

// Synonyms : get the synonyms of any word
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var (
		syns []string
		data synonyms
	)
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")

	if err != nil {
		return syns, errors.New(`bighuge: Failed when looking for synonyms for ` + term + `` + err.Error())
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}

	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}

	return syns, nil
}

// Thesaurus : interface for any thesaurus API to extend
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
