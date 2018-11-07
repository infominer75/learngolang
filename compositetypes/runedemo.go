package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "Hello World"
	fmt.Println("Contains \"world\"", Contains(asciiString, "World"))
	fmt.Println("Contains \"orld\"", Contains(asciiString, "orld"))
	//the below method is not guaraneteed to work on the
	//strings that contain other encodings e.g. East Asian languages like Japanese
	japanseString := "こんにちは世界"
	fmt.Println("Contains \"世界\"", Contains(japanseString, "世界"))
	fmt.Println("Rune count in japanese string : ", utf8.RuneCountInString(japanseString))
	DecodeRunes(asciiString)
	DecodeRunes(japanseString)
}

//Contains determines if s contains substring substr
func Contains(s, substr string) bool {
	if len(substr) > len(s) {
		return false
	}
	//an important note to observe below is that
	//each position in s is a rune (a sequence of bytes)
	//in case of ascii this would be the same as the character
	//however this is not true for other encodings
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

//HasPrefix determines if string s has prefix prefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//HasSuffix determines if string s has suffix suffix
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//DecodeRunes decodes the runes in the string s and prints them to the screen
func DecodeRunes(s string) {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d->%c,%d\n", i, r, size)
		//increment i by the size of the rune
		i = i + size
	}
}
