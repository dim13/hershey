package main

import "fmt"

var selector = map[string]string{
	"Gothic English Triplex": "gothgbt.hmp",
	"Gothic German Triplex":  "gothgrt.hmp",
	"Gothic Italian Triplex": "gothitt.hmp",
	"Greek Complex":          "greekc.hmp",
	"Greek Complex Small":    "greekcs.hmp",
	"Greek Plain":            "greekp.hmp",
	"Greek Simplex":          "greeks.hmp",
	"Cyrillic Complex":       "cyrilc.hmp",
	"Italic Complex":         "italicc.hmp",
	"Italic Complex Small":   "italiccs.hmp",
	"Italic Triplex":         "italict.hmp",
	"Script Complex":         "scriptc.hmp",
	"Script Simplex":         "scripts.hmp",
	"Roman Complex":          "romanc.hmp",
	"Roman Complex Small":    "romancs.hmp",
	"Roman Duplex":           "romand.hmp",
	"Roman Plain":            "romanp.hmp",
	"Roman Simplex":          "romans.hmp",
	"Roman Triplex":          "romant.hmp",
}

func main() {
	fnt := loadFont("data/hershey")

	for k, v := range selector {
		fmt.Println(k)
		m := getMap("data/" + v)

		for k, gl := range fnt.Select(m) {
			fmt.Println(k, gl)
		}
	}
}
