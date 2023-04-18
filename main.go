package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCodePoint(r rune) string {
	result := fmt.Sprintf("%+q", r)
	result = strings.TrimPrefix(result, "'")
	result = strings.TrimSuffix(result, "'")
	return result
}

func main() {
	var inp []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		for _, item := range sc.Text() {
			inp = append(inp, GetCodePoint(item))
		}
	}
	os.Stdout.WriteString(strings.Join(inp, ""))
}
