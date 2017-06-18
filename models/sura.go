package models

import (
	"bytes"
	"encoding/gob"

	"github.com/boltdb/bolt"
)

type Sura struct {
	Name                string   `json:"sura_name"`
	Number              int      `json:"sura_number"`
	EgyptianChronoOrder int      `json:"-"`
	NoldekeChronoOrder  int      `json:"-"`
	NumberVerses        int      `json:"-"`
	Location            string   `json:"-"`
	Year                string   `json:"-"`
	ArabicTitles        []string `json:"-"`
	EnglishTitles       []string `json:"-"`
	Topics              []string `json:"-"`
	Ayaat               []Ayah   `json:"-"`
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
