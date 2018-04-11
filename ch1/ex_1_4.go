/*
* Exercis e 1.4: Modify dup2 to print the names of all files in which
* each duplicated line occurs.
 */
package main

import (
    "bufio"
    "fmt"
    "os"
)

type fileLine struct {
	filename string
	line string
}

func dup2() map[fileLine]int {
    counts := make(map[fileLine]int)

    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts, "")
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Println(os.Stderr, "dup2: %v\n", err)
                continue
            }

            countLines(f, counts, arg)

            f.Close()
        }
    }

    return counts
}

func countLines(f *os.File, counts map[fileLine]int, filename string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
    	fl := fileLine{filename, input.Text()}
        counts[fl]++
    }
}
