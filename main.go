package main

import "fmt"

var selector = map[string]string{
	/*
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
	*/
	"Roman Simplex":          "romans.hmp",
}

func main() {
	fnt := loadFont("data/hershey")
	var x, y int

	for _, v := range selector {
		m := getMap("data/" + v)

		f := fnt.Select(m)
		for i := 32; i < 128; i++ {
			gl := f[rune(i)]
			if y + gl.W >= 4000 {
				y = 0
				x += 200
			}
			fmt.Printf("^%d,%d,%s", x, y, gl)
			y += gl.W
		}
	}
}
