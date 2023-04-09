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
    a := int64(nextInt())
    b := int64(nextInt())

    count := 0
    
    for {
        if a == b || a == 0 || b == 0{
            break
        } else if a > b {
            tmp := a/b
            a = a - b * tmp
            count = count + int(tmp)
        } else {
            tmp := b/a
            b = b - a * tmp
            count = count + int(tmp)
        }
    }

    
    if a == 0 || b == 0{
        fmt.Println(count - 1)
    } else {
        fmt.Println(count)
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
