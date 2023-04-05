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
    mochi := make(map[string] int)
    mochi = makeMochiMap(mochi, n)

    fmt.Println(len(mochi))
}

func makeMochiMap(mochi map[string] int, limit int) map[string] int {
    for i := 0; i < limit; i++ {
        mochi[nextStr()] = 0
    }

    return mochi
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
