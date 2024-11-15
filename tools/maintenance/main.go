package main

import (
    "fmt"
    "maintenance/cmd"
    "os"
)

func main() {
    if err := cmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
