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
    k := nextInt()

    takoyakitachi := make([]int, n)
    takoyakitachi = readIntSlice(takoyakitachi, n)

    Loop:
    for {
        for i: = 0; ;i++ {

        }
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

func readIntSlice(slice []int , limit int) []int {
    for i := 0; i < limit; i++ {
        slice[i] = nextInt()
    }

    return slice
}
