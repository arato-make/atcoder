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
    cards := make([]int, n)
    cards = readIntSlice(cards, n)

    for limit := len(cards); limit != 0; limit-- {
        for i := 0; i < limit - 1; i++ {
            j := i + 1
            if cards[i] < cards[j] {
                tmp := cards[i]
                cards[i] = cards[j]
                cards[j] = tmp
            }
        }
    }


    fmt.Println(sumSkipSlice(0, cards) - sumSkipSlice(1, cards))
}

func sumSkipSlice(start int, slice []int) int{
    sum := 0
    for i := start; i < len(slice); i += 2 {
        sum += slice[i]
    }
    return sum
}

func readIntSlice(slice []int , limit int) []int {
    for i := 0; i < limit; i++ {
        slice[i] = nextInt()
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
