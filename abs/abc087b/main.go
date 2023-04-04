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
    a := nextInt()
    b := nextInt()
    c := nextInt()
    x := nextInt()

    count := 0

    for i := 0; i <= a; i++ {
        sum := i * 500
        if sum == x {
            count++
        }

        if sum > x {
            break
        }
        for j := 0; j <= b; j++ {
            sum = i * 500 + j * 100
            if sum == x && j != 0 {
                count++
            }

            if sum > x {
                break
            }

            for k := 1; k <= c; k++ {
                sum = i * 500 + j * 100 + k * 50
                if sum == x && k != 0 {
                    count++
                }

                if sum > x {
                    break
                }
            }
        }
    }

    fmt.Println(count)
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}
