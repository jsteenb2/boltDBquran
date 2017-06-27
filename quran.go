package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	m "github.com/jsteenb2/boltDBquran/models"
	"github.com/jsteenb2/boltDBquran/parser"

	"os"

	"github.com/boltdb/bolt"
)

var quranBucket = []byte("quran")

func main() {
	dir, osErr := os.Getwd()
	if osErr != nil {
		log.Fatal(osErr)
	}
	fmt.Println(dir)
	db, err := bolt.Open(dir+"/quran.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// checkErr(BuildQuran(db, quranBucket))
	sura, _ := m.GetSurah(quranBucket, []byte{byte(1)}, db)
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
