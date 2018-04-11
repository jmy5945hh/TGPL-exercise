/*
* Exercise 1.3: Experiment to measure the difference in running time between our potentially
* inefficient versions and the one that uses strings.Join.
*/

package main

import (
    "strings"
    "os"
    "testing"
)

func BenchmarkEchoVer2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        EchoVer2()
    }
}

func BenchmarkEchoVer3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        EchoVer3()
    }
}

func EchoVer2() {
    strings.Join(os.Args[1:], " ")
}

func EchoVer3() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
	sep = " "
    }
}
