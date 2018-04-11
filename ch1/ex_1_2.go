/*
* Exercise 1.2: Modify the echo program to print the index and value of each of its arguments,
* one per line.
*/

package main

import (
    "fmt"
    "os"
)

func EchoEachLine() {
    for idx, arg := range os.Args[:] {
        fmt.Println(idx, " ", arg)
    }
}
