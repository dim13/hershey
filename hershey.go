package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Unit float64

type Point struct {
	X, Y Unit
}

type Path []Point
type Set []Path

type Glyph struct {
	S Set
	W Unit
}

type Font map[rune]Glyph

func parseInt(s string) (n int) {
	s = strings.TrimSpace(s)
	n, _ = strconv.Atoi(s)
	return
}

func parsePoint(in uint8) Unit {
	return Unit(in) - Unit('R')
}

func parseData(s string, w, h, scale Unit) Set {
	var st Set
	for i, el := range strings.Fields(s) {
		var ph Path
		if i > 0 {
			el = el[1:]
		}
		for n := 0; n < len(el); n += 2 {
			var p Point
			p.Y = scale * (w/2 + parsePoint(el[n]))
			p.X = scale * (h/2 + parsePoint(el[n+1]))
			ph = append(ph, p)
		}
		st = append(st, ph)
	}
	return st
}

func loadFont(fname string) Font {
	fnt := make(Font)

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n := parseInt(line[0:5])
		//k := parseInt(line[5:8])
		l := parsePoint(line[8])
		r := parsePoint(line[9])
		w := r - l
		scale := Unit(2.5)
		fnt[rune(n)] = Glyph{
			S: parseData(line[10:], w, 32, scale),
			W: w * scale,
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fnt
}

func (p Point) String() string {
	return fmt.Sprintf("%v,%v,", p.X, p.Y)
}

func (p Path) String() (s string) {
	//s = fmt.Sprint("Y0,", p[0])
	s = fmt.Sprint("M", p[0], "D")
	for _, pt := range p[1:] {
		s += fmt.Sprint(pt)
	}
	return
}

func (st Set) String() (s string) {
	for _, p := range st {
		s += fmt.Sprint(p)
	}
	return
}

func (g Glyph) String() string {
	return fmt.Sprint(g.S)
}

func (f Font) Select(n []int) Font {
	ret := make(Font)
	for i, p := range n {
		ret[rune(i+32)] = f[rune(p)]
	}
	return ret
}
