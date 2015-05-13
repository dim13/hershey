package main

import (
	"io/ioutil"
	"strings"
	"log"
	"strconv"
)

func getMap(fname string) []int {
	var n []int

	f, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range strings.Fields(string(f)) {
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
