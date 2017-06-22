package parser

import (
	"encoding/json"

	m "github.com/jsteenb2/boltDBquran/models"
)

type QuranCollection [114]m.Sura

func (q *QuranCollection) UnmarshalJSON(data []byte) error {
	var ayaat []m.Ayah
	if err := json.Unmarshal(data, &ayaat); err != nil {
		return err
	}

	for _, ayah := range ayaat {
		suraNum := ayah.ChapterNumber - 1
		if q[suraNum].Name == "" {
			q[suraNum].Name = ayah.SuraName
		}
		if q[suraNum].Number == 0 {
			q[suraNum].Number = ayah.ChapterNumber
		}
		q[suraNum].Ayaat = append(q[suraNum].Ayaat, m.NewFormattedAyah(ayah))
	}
	return nil
}
