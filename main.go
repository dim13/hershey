package main

import (
	"flag"
	"fmt"
	"log"
)

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

func printAll(f Font) {
	var x, y Unit

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

func printStruct(f Font) {
	fmt.Println("var height = Unit(72)")
	fmt.Println("var font = Font{")
	for i := 0; i < len(f); i++ {
		r := rune(i + 32)
		gl := f[r]
		fmt.Printf("%q: Glyph{\n", r)
		fmt.Println("S: Set{")
		for _, s := range gl.S {
			fmt.Println("Path{")
			for _, p := range s {
				fmt.Printf("Point{%v, %v},\n", p.X, p.Y)
			}
			fmt.Println("},")
		}
		fmt.Println("},")
		fmt.Printf("W: %v,\n", gl.W)
		fmt.Println("},")
	}
	fmt.Println("}")
}

var font = flag.String("font", "Roman Simplex", "Font to use")

func main() {
	flag.Parse()
	f := loadFont("data/hershey", Unit(2))
	s, ok := selector[*font]
	if !ok {
		log.Fatal("no such font")
	}
	m := getMap("data/" + s)
	fnt := f.Select(m)
	printStruct(fnt)
}
