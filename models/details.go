package models

import (
	"regexp"
	"strings"
)

type SuraDeetsCol [114]SuraDeets

type SuraDeets struct {
	SurahNum            int    `json:"surah_num"`
	NumberVerses        int    `json:"num_verses"`
	NoldekeChronoOrder  int    `json:"noldeke_chrono_order"`
	EgyptianChronoOrder int    `json:"egyptian_chrono_order"`
	Location            string `json:"location"`
	Year                string `json:"year"`
	ArabicTitles        string `json:"arabic_titles"`
	EnglishTitles       string `json:"english_titles"`
	Topics              string `json:"topics_themes"`
}

func splitDeets(list, seperator string) []string {
	var col []string
	for _, val := range splitOnSeperator(list, seperator) {
		col = append(col, processString(val)...)
	}
	return col
}

func splitOnSeperator(list, seperator string) []string {
	var stringCol []string
	if seperator != "" {
		stringCol = strings.Split(list, seperator)
	} else {
		stringCol = []string{list}
	}
	return stringCol
}

func processString(val string) []string {
	re := regexp.MustCompile("\\[.*\\]")
	st := strings.TrimPrefix(val, " ")
	st = strings.Title(st)
	st = re.ReplaceAllString(st, "")
	splitSt := strings.Split(st, ",")

	if len(splitSt) < 2 {
		return []string{st}
	}

	return splitSt
}
