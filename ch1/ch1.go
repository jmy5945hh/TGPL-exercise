package main

import (
	"fmt"
	"os"
)

func main() {
	// ex 1.1
	fmt.Println("############## ex 1.1 ##############")
	EchoVer1()

	// ex 1.2
	fmt.Println("############## ex 1.2 ##############")
	EchoEachLine()

	// ex 1.3 is a benchmark test, run as "go test -bench=."
	fmt.Println("############## ex 1.3 ##############")
	fmt.Println("PASS")

	// ex 1.4
	fmt.Println("############## ex 1.4 ##############")
	fileLineCount := dup2()

	for key, val := range fileLineCount {
		if val > 1 {
			fmt.Println(key, val)
		}
	}

	fmt.Println("############## ex 1.5 ##############")
	outGif, err := os.Create("lissajous.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outGif.Close()
	lissajous(outGif)

	fmt.Println("############## ex 1.6 ##############")
	outGif2, err := os.Create("lissajous2.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outGif2.Close()
	lissajous2(outGif2)
}
