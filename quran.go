package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	m "github.com/jsteenb2/boltDBquran/models"
	"github.com/jsteenb2/boltDBquran/parser"

	"github.com/boltdb/bolt"
)

var quranBucket = []byte("quran")

func main() {
	dir, osErr := os.Getwd()
	if osErr != nil {
		log.Fatal(osErr)
	}

	db, err := bolt.Open(dir+"/quran.db", 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sura, err := m.GetSurah(quranBucket, []byte{byte(1)}, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", sura)
}

func gobDecode(data []byte) (*m.Sura, error) {
	var s *m.Sura
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func BuildQuran(db *bolt.DB, bucket []byte) error {
	f, err := ioutil.ReadFile("quran.json")
	checkErr(err)

	var quran parser.QuranCollection
	checkErr(json.Unmarshal(f, &quran))

	f, err = ioutil.ReadFile("quranDeets.txt")
	checkErr(err)

	var deets m.SuraDeetsCol
	checkErr(json.Unmarshal(f, &deets))
	for i, sura := range quran {
		sura.AddDeets(&deets[i])
		checkErr(sura.Save(db, quranBucket))
		if err != nil {
			break
		}
	}
	return err
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
