package main

import (
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

var language Language

func translate(label string) *string {
	for _, phrase := range language.Phrases {
		if phrase.Label == label {
			return &phrase.Translated
		}
	}
	return &label
}

func readTranslationFile(langCode string) (string) {
	content, err := ioutil.ReadFile("translations/" + langCode + ".json")
	if err != nil {
		return ""; // TODO: Error
	}

	return string(content);
}

func setLocale(lc string) {
	// Attempt to load translation
	content := readTranslationFile(lc);

	if content == "" && lc != "en" {
		// Fallback to English
		content = readTranslationFile("en")
	}

	err := json.Unmarshal([]byte(content), &language)

	if err != nil {
		return; // TODO: Error
	}
}

func detectLocale() {
	lc, err := jibber_jabber.DetectLanguage()

	if err == nil {
		setLocale(lc)
	} else {
		setLocale("en")
	}
}
