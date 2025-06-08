package main

import (
	"fmt"
	"os"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s, %s\n", arg, basename(arg))
	}
}
