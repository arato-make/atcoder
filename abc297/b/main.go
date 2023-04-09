package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    sc.Scan()
    str := sc.Text()

    check_first := false
    check_second := false

    if (strings.Index(str, "B") + strings.LastIndex(str, "B")) % 2 == 1 {
        check_first = true
    }

    if strings.Index(str, "R") < strings.Index(str, "K") && strings.Index(str, "K") < strings.LastIndex(str, "R") {
        check_second = true
    }

    if check_first && check_second {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}