package main

import (
  "fmt"
  "strconv"
  "strings"
)

func main() {
    var n int = 1345679
    stringNum := strconv.FormatInt(int64(n), 10)

    for pos, char := range stringNum {
       zeroCount := strings.Repeat("0", len(stringNum)-pos-1)
       fmt.Printf("%c%s\n", char, zeroCount)
    }
}