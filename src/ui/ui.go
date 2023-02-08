package ui

import ("fmt"
        "os"
)

var (
    stdout = os.Stdout
    stderr = os.Stderr
)

func Println(s string) {
    fmt.Fprintln(stdout, s)
}

func ErrPrintln(s string) {
    fmt.Fprintf(stderr, s)
}


