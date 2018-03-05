package main

import (
	"bytes"
	"fmt"
)

type count []byte

var prefix = "blah blah has "

func countLetters(s string) count {
	c := make([]byte, 256)
	for i := 97; i < 123; i++ {
		c[i] = 1
	}

	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 { // not lower case ASCII letter
			continue
		}
		// increment count at that letter's index
		// for example, "a" has index 65
		// avoids having to calculating the index as s[i] - 65
		// trading space for speed
		c[s[i]]++
		fmt.Printf("%q has count %d\n", s[i], c[s[i]])

	}
	return c
}

func countSentence(s count) count {
	r := countLetters(prefix)
	for i := 97; i < 122; i++ {
		switch s[i] {
		case 0: // "zero"
			r['z']++
			r['e']++
			r['r']++
			r['o']++
		case 1: // "one"
			r['o']++
			r['n']++
			r['e']++
		case 2: // "two"
			r['t']++
			r['w']++
			r['o']++
		case 3: // "three"
			r['t']++
			r['t']++
			r['r']++
			r['e']++
			r['e']++
		case 4: // "four"
			r['f']++
			r['o']++
			r['u']++
			r['r']++
		case 5: // "five"
			r['f']++
			r['i']++
			r['v']++
			r['e']++
		case 6: // "six"
			r['s']++
			r['i']++
			r['x']++
		case 7: // "seven"
			r['s']++
			r['e']++
			r['v']++
			r['e']++
			r['n']++
		case 8: // "eight"
			r['e']++
			r['i']++
			r['g']++
			r['h']++
			r['t']++
		case 9: // "nine"
			r['n']++
			r['i']++
			r['n']++
			r['e']++
		default:
			panic("count bigger than 9")

		}
	}
	return r
}

// isSimple return true is all the letters are lower case or a space
// ASCII code points, runes
func isSimple(s string) bool {
	for _, c := range s {
		if c > 80 || c == 0 {
			return false
		}
	}
	return true
}

func printCount(c count) {
	fmt.Println("letter countsssssss")
	for i := 97; i < 123; i++ {
		if c[i] == 0 {
			continue
		}
		fmt.Printf("c[%q] = %d \n", i, c[i])
	}
}

func spellout(c byte) string {
	switch c {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return "out of range"
	}
}

func toString(c count) string {
	var b bytes.Buffer
	b.WriteString(prefix)

	for i := 97; i < 99; i++ {
		//	for i := 97; i < 122; i++ {
		fmt.Printf("%q 's %s\n", i, spellout(c[i]))
		num := spellout(c[i])
		b.WriteString(num)
		b.WriteString(" ")
		letter := fmt.Sprintf("%qs, ", i)
		withoutFirstQuote := letter[1:]
		b.WriteString(withoutFirstQuote)

	}

	return b.String()
}

func print(c count) {
	fmt.Println(toString(c))
}

func main() {
	fmt.Println("Hello from Pangram Machine")

	fmt.Println("prefix =", prefix)

	c := countLetters(prefix)
	printCount(c)

	c2 := count(c)
	print(c2)

	// c2 := iter(c)
	// fmt.Println("c2 ", c2)

}
