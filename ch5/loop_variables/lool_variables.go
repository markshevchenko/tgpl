package main

import "fmt"

func main() {
	var fs []func()
	for i := range 20 {
		fs = append(fs, func() { fmt.Printf("%d\n", i) })
	}

	for _, f := range fs {
		f()
	}
}
