package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCodePoint(r rune) string {
	return fmt.Sprintf("%U", r)
}

func main() {
	var inp []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
        for _,item := range sc.Text(){
            inp = append(inp, GetCodePoint(item))
        }
	}
    os.Stdout.WriteString(strings.Join(inp , "\n"))
}
