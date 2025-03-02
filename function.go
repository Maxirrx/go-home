package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func find_day(jour string) (time.Time, error) {

	joursMap := map[string]time.Weekday{
		"dimanche": time.Sunday,
		"lundi":    time.Monday,
		"mardi":    time.Tuesday,
		"mercredi": time.Wednesday,
		"jeudi":    time.Thursday,
		"vendredi": time.Friday,
		"samedi":   time.Saturday,
	}

	jourSemaine, existe := joursMap[jour]
	if !existe {
		return time.Time{}, errors.New("jour invalide : " + jour)
	}

	aujourdHui := time.Now()
	joursRestants := (int(jourSemaine) - int(aujourdHui.Weekday()) + 7) % 7

	if joursRestants == 0 {
		joursRestants = 7
	}

	prochain := aujourdHui.AddDate(0, 0, joursRestants)

	return prochain, nil
}

func check_date(phrase string) (string, error) {
	jour_semaine := []string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	numero_jour_moi := []string{
		"31", "trente-et-un", "30", "trente", "29", "vingt-neuf", "28", "vingt-huit", "27", "vingt-sept",
		"26", "vingt-six", "25", "vingt-cinq", "24", "vingt-quatre", "23", "vingt-trois", "22", "vingt-deux",
		"21", "vingt-et-un", "20", "vingt", "19", "dix-neuf", "18", "dix-huit", "17", "dix-sept", "16", "seize",
		"15", "quinze", "14", "quatorze", "13", "treize", "12", "douze", "11", "onze", "10", "dix",
		"9", "neuf", "8", "huit", "7", "sept", "6", "six", "5", "cinq", "4", "quatre", "3", "trois",
		"2", "deux", "1", "un",
	}
	mois_annee := []string{
		"janvier", "février", "mars", "avril", "mai", "juin",
		"juillet", "août", "septembre", "octobre", "novembre", "décembre",
	}

	for _, jour := range jour_semaine {
		if strings.Contains(phrase, jour) {
			convert, _ := find_day(jour)
			date := convert.Format("01-01-2006")
			fmt.Println(date)
		}
		for _, numero := range numero_jour_moi {
			if strings.Contains(phrase, numero) {
				for _, moi := range mois_annee {
					if strings.Contains(phrase, moi) {
						return moi + " " + numero, nil
					}
				}
			}
		}

	}
	return "err", nil

}

func gestion_date_longue(mois string) (time.Time, error) {

	moiMap := map[string]time.Month{
		"janvier":   time.January,
		"février":   time.February,
		"mars":      time.March,
		"avril":     time.April,
		"mai":       time.May,
		"juin":      time.June,
		"juillet":   time.July,
		"août":      time.August,
		"septembre": time.September,
		"octobre":   time.October,
		"novembre":  time.November,
		"décembre":  time.December,
	}

	joursmap := map[int]string{
		1: "un", 2: "deux", 3: "trois", 4: "quatre", 5: "cinq",
		6: "six", 7: "sept", 8: "huit", 9: "neuf", 10: "dix",
		11: "onze", 12: "douze", 13: "treize", 14: "quatorze", 15: "quinze",
		16: "seize", 17: "dix-sept", 18: "dix-huit", 19: "dix-neuf", 20: "vingt",
		21: "vingt-et-un", 22: "vingt-deux", 23: "vingt-trois", 24: "vingt-quatre", 25: "vingt-cinq",
		26: "vingt-six", 27: "vingt-sept", 28: "vingt-huit", 29: "vingt-neuf", 30: "trente",
		31: "trente-et-un",
	}

	jour := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	mots := strings.Split(mois, " ")
	sjour := mots[0]

	_ = sjour
	_ = jour
	_ = joursmap

	mois = strings.ToLower(mots[1])

	moisDelAnnee, existe := moiMap[mois]
	if !existe {
		return time.Time{}, errors.New("jour invalide : " + mois)
	}

	aujourdHui := time.Now()
	moisActuel := aujourdHui.Month()

	annee := aujourdHui.Year()
	if moisDelAnnee <= moisActuel {
		annee++
	}

	date := time.Date(annee, moisDelAnnee, 1, 0, 0, 0, 0, aujourdHui.Location())

	return date, nil
}

func gestion_event(phrase string) string {
	synonyme_agenda := []string{"agenda", "calendrier", "emplois du temps"}
	synonyme_ajouter := []string{"ajoutes", "insert", "mets"}
	var evenement string

	for _, synonyme := range synonyme_ajouter {
		if strings.Contains(phrase, synonyme) {
			for _, synonyme := range synonyme_agenda {
				if strings.Contains(phrase, synonyme) {
					event := strings.Split(phrase, synonyme)
					if len(event) > 1 {
						evenement = event[1]

					}
				}
			}
		}
	}
	return evenement
}
