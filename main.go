package main

import (
	"flag"
	"fmt"
	"log"
)

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

func printAll(f Font) {
	var x, y int

	for i := 32; i < 128; i++ {
		gl := f[rune(i)]
		width := gl.Right - gl.Left
		if y+width >= 4000 {
			y = 0
			x += 100
		}
		fmt.Printf("^%v,%v,%s", x, y, gl)
		y += width
	}
}

func printStruct(f Font) {
	fmt.Println("var font = Font{")
	for i := 0; i < len(f); i++ {
		r := rune(i + 32)
		gl := f[r]
		fmt.Printf("%q: Glyph{\n", r)
		fmt.Println("Set: Set{")
		for _, s := range gl.Set {
			fmt.Println("Path{")
			for _, p := range s {
				fmt.Printf("Point{%v, %v},\n", p.X, p.Y)
			}
			fmt.Println("},")
		}
		fmt.Println("},")
		fmt.Printf("Left: %v,\n", gl.Left)
		fmt.Printf("Right: %v,\n", gl.Right)
		fmt.Println("},")
	}
	fmt.Println("}")
}

var font = flag.String("font", "Roman Simplex", "Font to use")

func main() {
	flag.Parse()
	f := loadFont("data/hershey")
	s, ok := selector[*font]
	if !ok {
		log.Fatal("no such font")
	}
	m := getMap("data/" + s)
	fnt := f.Select(m)
	printStruct(fnt)
}
