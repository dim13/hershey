package main

import "fmt"

var selector = map[string]string{
	"Roman Plain":            "romanp.hmp",
	"Roman Simplex":          "romans.hmp",
	"Roman Duplex":           "romand.hmp",
	"Roman Complex":          "romanc.hmp",
	"Roman Complex Small":    "romancs.hmp",
	"Roman Triplex":          "romant.hmp",
	"Script Simplex":         "scripts.hmp",
	"Script Complex":         "scriptc.hmp",
	"Italic Complex":         "italicc.hmp",
	"Italic Complex Small":   "italiccs.hmp",
	"Italic Triplex":         "italict.hmp",
	"Greek Plain":            "greekp.hmp",
	"Greek Simplex":          "greeks.hmp",
	"Greek Complex":          "greekc.hmp",
	"Greek Complex Small":    "greekcs.hmp",
	"Cyrillic Complex":       "cyrilc.hmp",
	"Gothic English Triplex": "gothgbt.hmp",
	"Gothic German Triplex":  "gothgrt.hmp",
	"Gothic Italian Triplex": "gothitt.hmp",
}

func main() {
	var x, y Unit

	fnt := loadFont("data/hershey")
	m := getMap("data/" + selector["Roman Simplex"])

	f := fnt.Select(m)
	for i := 32; i < 128; i++ {
		gl := f[rune(i)]
		if y+gl.W >= 4000 {
			y = 0
			x += 100
		}
		fmt.Printf("^%v,%v,%s", x, y, gl)
		y += gl.W
	}
}
