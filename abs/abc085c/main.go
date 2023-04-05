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
    y := nextInt()

    printed := false

    Loop:
    for i := n; i >= 0; i-- {
        for j := n - i; j >= 0; j-- {
            k := n - i - j
            sum := i * 10000 + j * 5000 + k * 1000
            if sum == y {
                fmt.Printf("%d %d %d\n",i ,j ,k)
                printed = true
                break Loop
            }
        }
    }

    if !printed {
        fmt.Println("-1 -1 -1")
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
