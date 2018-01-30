package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Path []Point
type Set []Path

type Glyph struct {
	Set    Set
	Width  int
	Height int
}

type Font map[rune]Glyph

func parseInt(s string) (n int) {
	s = strings.TrimSpace(s)
	n, _ = strconv.Atoi(s)
	return
}

func parsePoint(in uint8) int {
	return int(in) - int('R')
}

func parseData(s string, w, h int) Set {
	var st Set
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		el := scanner.Text()
		if len(el)%2 != 0 && el[0] == 'R' {
			el = el[1:]
		}
		var l Path
		for n := 0; n < len(el); n += 2 {
			p := Point{
				Y: w/2 + parsePoint(el[n]),
				X: h/2 + parsePoint(el[n+1]),
			}
			l = append(l, p)
		}
		st = append(st, l)
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
		fnt[rune(n)] = Glyph{
			Set:    parseData(line[10:], w, 32),
			Width:  w,
			Height: 32,
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
	return fmt.Sprint(g.Set)
}

func (f Font) Select(n []int) Font {
	ret := make(Font)
	for i, p := range n {
		ret[rune(i+32)] = f[rune(p)]
	}
	return ret
}
