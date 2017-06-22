package models

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/boltdb/bolt"
)

type Sura struct {
	Name                string   `json:"sura_name"`
	Number              int      `json:"sura_number"`
	EgyptianChronoOrder int      `json:"egyptian_chrono_order"`
	NoldekeChronoOrder  int      `json:"noldeke_chrono_order"`
	NumberVerses        int      `json:"number_verses"`
	Location            string   `json:"location"`
	Year                string   `json:"year"`
	ArabicTitles        []string `json:"arabic_titles"`
	EnglishTitles       []string `json:"english_titles"`
	Topics              []string `json:"topics"`
	Ayaat               []Ayah   `json:"ayaat"`
}

func (s *Sura) AddDeets(deets *SuraDeets) {
	s.Location = deets.Location
	s.Year = deets.Year
	s.EgyptianChronoOrder = deets.EgyptianChronoOrder
	s.NoldekeChronoOrder = deets.NoldekeChronoOrder
	s.NumberVerses = deets.NumberVerses
	s.ArabicTitles = splitDeets(deets.ArabicTitles, "aka:")
	s.EnglishTitles = splitDeets(deets.EnglishTitles, "aka:")
	s.Topics = splitDeets(deets.Topics, "")
}

func (s Sura) gobEncode() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *Sura) Save(db *bolt.DB, bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		data, err := s.gobEncode()

		var surahNum = []byte{byte(s.Number)}
		return b.Put(surahNum, data)
	})
}

func GetSurah(bucket, surahNum []byte, db *bolt.DB) (*Sura, error) {
	var sura *Sura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get(surahNum)

		var decodeErr error
		sura, decodeErr = gobDecode(v)
		if decodeErr != nil {
			log.Println(decodeErr)
			return decodeErr
		}
		return nil
	})
	return sura, err
}

func gobDecode(data []byte) (*Sura, error) {
	var s *Sura
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
