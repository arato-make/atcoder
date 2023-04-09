package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    h := nextInt()
    nextInt()

    for i:= 0; i < h ; i++ {
        fmt.Println(strings.Replace(nextStr(), "TT", "PC", -1))
    }
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func nextStr() string {
    sc.Scan()
    return sc.Text()
}
