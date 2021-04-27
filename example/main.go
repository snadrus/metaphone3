package main

import (
	"fmt"

	"github.com/snadrus/metaphone3"
)

/**
 * @param args
 */
func main() {
	// example code
	m3 := metaphone3.New()

	//m3.SetEncodeVowels(true);
	//m3.SetEncodeExact(true);

	p, s := m3.Encode("ron")

	fmt.Println("iron : " + p)
	fmt.Println("iron : (alt) " + s)

	p, s = m3.Encode("witz")

	fmt.Println("witz : " + p)
	fmt.Println("witz : (alt) " + s)

	p, s = m3.Encode("")

	fmt.Println("BLANK : " + p)
	fmt.Println("BLANK : (alt) " + s)

	// these settings default to false
	m3.SetEncodeExact(true)
	m3.SetEncodeVowels(true)

	for _, test := range []string{"Guillermo", "VILLASENOR", "GUILLERMINA", "PADILLA", "BJORK", "belle", "ERICH", "GLOWACKI", "qing", "tsing"} {
		p, s := m3.Encode(test)
		fmt.Println(test + " : " + p)
		fmt.Println(test + " : (alt) " + s)
	}
}
