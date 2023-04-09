package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    n := nextInt()
    d := nextInt()

    previous_click := nextInt()
    printed := false

    for i:= 0; i < n-1 ; i++ {
        now_click := nextInt()

        if now_click - previous_click <= d {
            fmt.Println(now_click)
            printed = true
            break
        }

        previous_click = now_click
    }

    if !printed {
        fmt.Println(-1)
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
