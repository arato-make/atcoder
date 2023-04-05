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
    var str string
    fmt.Scan(&str)

    str = strings.ReplaceAll(str, "eraser", "0")
    str = strings.ReplaceAll(str, "erase", "0")
    str = strings.ReplaceAll(str, "dreamer", "0")
    str = strings.ReplaceAll(str, "dream", "0")
    str = strings.ReplaceAll(str, "0", "")

    if str == "" {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
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
