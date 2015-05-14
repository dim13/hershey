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
	N, K int
	L, R int
	S    Set
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

func parseData(s string) Set {
	var st Set
	for i, el := range strings.Fields(s) {
		var ph Path
		if i > 0 {
			el = el[1:]
		}
		for n := 0; n < len(el); n += 2 {
			var p Point
			p.X = parsePoint(el[n])
			p.Y = parsePoint(el[n+1])
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
		gl := Glyph{
			N: parseInt(line[0:5]),
			K: parseInt(line[5:8]),
			L: parsePoint(line[8]),
			R: parsePoint(line[9]),
			S: parseData(line[10:]),
		}
		fnt[rune(gl.N)] = gl
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fnt
}

func (p Point) String() string {
	return fmt.Sprintf("(%v %v)", p.X, p.Y)
}

func (p Path) String() (s string) {
	for i, pt := range p {
		s += fmt.Sprint(pt)
		if i < len(p)-1 {
			s += fmt.Sprint(",")
		} else {
			s += fmt.Sprint(" ")
		}
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

func (g Glyph) Width() int {
	return g.R - g.L
}

func (g Glyph) Max() (x, y int) {
	var top, bottom int
	var left, right int
	for _, pt := range g.S {
		for _, p := range pt {
			if p.X > right {
				right = p.X
			}
			if p.X < left {
				left = p.X
			}
			if p.Y > top {
				top = p.Y
			}
			if p.Y < bottom {
				bottom = p.Y
			}
		}
	}
	return right - left, top - bottom
}

func (f Font) Select(n []int) Font {
	ret := make(Font)
	for i, p := range n {
		ret[rune(i+32)] = f[rune(p)]
	}
	return ret
}
