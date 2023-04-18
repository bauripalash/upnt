package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

var rawInput = `
A,U+0041; 
B,U+0042;
C,U+0043;
D,U+0044;
E,U+0045;
F,U+0046;
G,U+0047;
H,U+0048;
I,U+0049;
অ,U+0985;
প,U+09AA;
`

func TestGetCodePoint(t *testing.T) {
	tc := map[rune]string{}

	x := strings.Split(rawInput, ";")
	for _, item := range x {
		ts := strings.TrimSpace(item)
		x := strings.Split(ts, ",")
		if len(x) == 2 {
			key, _ := utf8.DecodeRuneInString(x[0])
			tc[key] = x[1]
		}
	}

	for k, v := range tc {
		got := GetCodePoint(k)
		if got != v {
			t.Errorf("%s != %s", v, got)
		}
	}
}

func BenchmarkGetCodePoint(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetCodePoint('অ')

	}

}

func ExampleGetCodePoint() {
	fmt.Println(GetCodePoint('A'))
	// Output: U+0041
}
