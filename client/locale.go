package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/pivotal-cf-experimental/jibber_jabber"
)

type Phrase struct {
	Label      string `json:"label"`
	Translated string `json:"string"`
}

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`

	Phrases []Phrase `json:"strings"`
}

func (l *Language)getPhrase(label string) string {
	for _, phrase := range l.Phrases {
		if phrase.Label == label {
			return phrase.Translated
		}
	}
	return label
}

func NewLanguage(lc string)(*Language, error) {
	// Attempt to load translation
	content, err := ioutil.ReadFile("translations/" + lc + ".json")
	if err != nil {
		fmt.Printf("Failed to set %q as locale.\n", lc)
		if lc != "en" {
			return NewLanguage("en")
		}
		return nil, err
	}

	var language Language
	err = json.Unmarshal([]byte(content), &language)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Setting locale %q.\n", lc)
	return &language, nil
}

func DetectLocale()(string) {
	lc, err := jibber_jabber.DetectLanguage()

	if err == nil {
		return lc
	}
	return "en"
}
