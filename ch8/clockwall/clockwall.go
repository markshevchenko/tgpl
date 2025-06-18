package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for _, arg := range os.Args[1:] {
		pair := strings.SplitN(arg, "=", 2)
		if len(pair) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid arg %s\n", arg)
		} else {
			wg.Add(1)
			go printTime(pair[0], pair[1])
		}
	}

	wg.Wait()
}

func printTime(city, host string) {
	defer wg.Done()
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Printf("%s: %s\n", city, scanner.Text())
	}
}
