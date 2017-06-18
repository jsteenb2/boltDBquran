package models

import (
	"strings"
)

type Ayah struct {
	AyahTitle           string `json:"ayah-title"`
	AyahID              int    `json:"ayahid"`
	AyahNumber          int    `json:"Ayah_number"`
	PageTag             string `json:"page-tag"`
	PageID              int    `json:"pageid"`
	RubuTag             string `json:"rubu-tag"`
	RubuhizbID          int    `json:"rubuhizbid"`
	HizbTag             string `json:"hizb-tag"`
	HizBID              int    `json:"hizbid"`
	JuzTag              string `json:"juz-tag"`
	JuzID               int    `json:"juzid"`
	SuraName            string `json:"sura name"`
	SuraTag             string `json:"Sura-tag"`
	ChapterNumber       int    `json:"chapter_number"`
	Meaning             string `json:"meaning"`
	ContentArabic       string `json:"content_ar"`
	ContentEnglish      string `json:"content_en"`
	ContentEnglishSi    string `json:"content_en_si"`
	ContentEnglishTrans string `json:"content_en_tr"`
	IbrahimWalk         string `json:"ibrahim_walk"`
	AbdulbasitMurrattal string `json:"abdulbasit_murattal"`
	Basfar              string `json:"basfar"`
	Assudais            string `json:"assudais"`
	Alafasy             string `json:"alafasy"`
	Algamdi             string `json:"algamdi"`
	Alshatri            string `json:"alshatri"`
	Alhudhaify          string `json:"alhudhaify"`
	Alhusary            string `json:"alhusary"`
	Arrefai             string `json:"arrefai"`
	Alakhdar            string `json:"alakhdar"`
	Almuaiqly           string `json:"almuaiqly"`
	Jebril              string `json:"jebril"`
	Minshawi            string `json:"minshawi"`
	Attablawi           string `json:"attablawi"`
	Alshuraym           string `json:"alshuraym"`
	Abdulbasit          string `json:"abdulbasit"`
	Ayyub               string `json:"ayyub"`
	Bukhatir            string `json:"bukhatir"`
	KhanUrdu            string `json:"khan_urdu"`
	Ibrahimdosary       string `json:"ibrahimdosary"`
	Jazaery             string `json:"jazaery"`
	Kabiri              string `json:"kabiri"`
	Juhayne             string `json:"juhayne"`
	Matroud             string `json:"matroud"`
	Neena               string `json:"neena"`
	Jaber               string `json:"jaber"`
	Alaqimy             string `json:"alaqimy"`
	Hajjaj              string `json:"hajjaj"`
	Baliyev             string `json:"baliyev"`
	Bosnian             string `json:"bosnian"`
	Abbad               string `json:"abbad"`
	Hussarym            string `json:"hussarym"`
	HussaryMujawad      string `json:"hussary_mujawad"`
	Tunaji              string `json:"tunaji"`
	Qahdhani            string `json:"qahdhani"`
	Albana              string `json:"albana"`
	Alqasim             string `json:"alqasim"`
	Alqatami            string `json:"alqatami"`
	Alajamy             string `json:"alajamy"`
}

func NewFormattedAyah(ayah Ayah) Ayah {
	ayah.IbrahimWalk = removeHost(ayah.IbrahimWalk)
	ayah.AbdulbasitMurrattal = removeHost(ayah.AbdulbasitMurrattal)
	ayah.Basfar = removeHost(ayah.Basfar)
	ayah.Assudais = removeHost(ayah.Assudais)
	ayah.Alafasy = removeHost(ayah.Alafasy)
	ayah.Algamdi = removeHost(ayah.Algamdi)
	ayah.Alshatri = removeHost(ayah.Alshatri)
	ayah.Alhudhaify = removeHost(ayah.Alhudhaify)
	ayah.Alhusary = removeHost(ayah.Alhusary)
	ayah.Arrefai = removeHost(ayah.Arrefai)
	ayah.Alakhdar = removeHost(ayah.Alakhdar)
	ayah.Almuaiqly = removeHost(ayah.Almuaiqly)
	ayah.Jebril = removeHost(ayah.Jebril)
	ayah.Minshawi = removeHost(ayah.Minshawi)
	ayah.Attablawi = removeHost(ayah.Attablawi)
	ayah.Alshuraym = removeHost(ayah.Alshuraym)
	ayah.Abdulbasit = removeHost(ayah.Abdulbasit)
	ayah.Ayyub = removeHost(ayah.Ayyub)
	ayah.Bukhatir = removeHost(ayah.Bukhatir)
	ayah.KhanUrdu = removeHost(ayah.KhanUrdu)
	ayah.Ibrahimdosary = removeHost(ayah.Ibrahimdosary)
	ayah.Jazaery = removeHost(ayah.Jazaery)
	ayah.Kabiri = removeHost(ayah.Kabiri)
	ayah.Juhayne = removeHost(ayah.Juhayne)
	ayah.Matroud = removeHost(ayah.Matroud)
	ayah.Neena = removeHost(ayah.Neena)
	ayah.Jaber = removeHost(ayah.Jaber)
	ayah.Alaqimy = removeHost(ayah.Alaqimy)
	ayah.Hajjaj = removeHost(ayah.Hajjaj)
	ayah.Baliyev = removeHost(ayah.Baliyev)
	ayah.Bosnian = removeHost(ayah.Bosnian)
	ayah.Abbad = removeHost(ayah.Abbad)
	ayah.Hussarym = removeHost(ayah.Hussarym)
	ayah.HussaryMujawad = removeHost(ayah.HussaryMujawad)
	ayah.Tunaji = removeHost(ayah.Tunaji)
	ayah.Qahdhani = removeHost(ayah.Qahdhani)
	ayah.Albana = removeHost(ayah.Albana)
	ayah.Alqasim = removeHost(ayah.Alqasim)
	ayah.Alqatami = removeHost(ayah.Alqatami)
	ayah.Alajamy = removeHost(ayah.Alajamy)
	return ayah
}

func removeHost(link string) string {
	return strings.Replace(link, "http://www.everyayah.org/data/", "", -1)
}
