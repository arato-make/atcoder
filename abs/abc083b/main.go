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
    a := nextInt()
    b := nextInt()

    sum := 0

    for i := 1; i <= n; i++ {
        var slice []int
        slice = getDigit(slice, i)

        s := digitSum(slice)

        if a <= s && s <= b {
            sum += i
        }
    }
    
    fmt.Println(sum)
}

func digitSum(slice[] int) int {
    sum := 0
    for _, n := range slice {
        sum += n
    }

    return sum
}

func getDigit(slice[] int, num int) []int {
    if num > 0 {
        slice = append(slice, num % 10)
        slice = getDigit(slice, num / 10)
    }
    return slice
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}
