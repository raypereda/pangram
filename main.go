package main

import "fmt"

type sentence []byte

var s sentence

var prefix = "a b sentence has"

func countString(s string) sentence {
	c := make([]byte, 256)
	var i int
	for i = 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
		// increment count at that letter's index
		// for example, "a" has index 65
		c[s[i]]++
	}
	return c
}

func iter(s sentence) sentence {
	r := make([]byte, 256)

	// maybe r is initialized with prefix string counts, pre-computed

	var i int
	for i = 0; i < len(s); i++ {
		count := s[i]
		switch count {
		case 0:
			// do nothing
		case 1: // "one"
			r[65]++ // "o"
			r[65]++ // "n"
			r[65]++ // "e"
		case 2: // "two"
			r[65]++ // "t"
			r[65]++ // "w"
			r[65]++ // "o"
		case 3: // "three"
			r[65]++ // "t"
			r[65]++ // "w"
			r[65]++ // "o"
		}
	}
	return r
}

func count(sentence) sentence {
	return make([]byte, 2)
}

func main() {
	fmt.Println("Hello from Pangram Machine")

	c := countString(prefix)
	fmt.Println(c)

	fmt.Println(count(c))

	c2 := iter(c)
	fmt.Println("c2 ", c2)

}
