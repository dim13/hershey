package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getMap(fname string) []int {
	var n []int
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		r := scanner.Text()
		if strings.Contains(r, "-") {
			rr := strings.Split(r, "-")
			a, _ := strconv.Atoi(rr[0])
			b, _ := strconv.Atoi(rr[1])
			for i := a; i <= b; i++ {
				n = append(n, i)
			}
		} else {
			a, _ := strconv.Atoi(r)
			n = append(n, a)
		}
	}
	return n
}
