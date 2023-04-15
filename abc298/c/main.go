package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    nextInt()
    num := make(map[int] []int)

    hako := make(map[int] []int)


    q := nextInt()

    for i := 0; i < q; i++ {
        query := nextInt()

        if query == 1 {
            query1(num, hako)
        } else if query == 2 {
            query2(hako)
        } else {
            query3(num)
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

func query1(num map[int] []int, hako map[int] []int) {
    i := nextInt()
    j := nextInt()

    if val, ok := hako[j-1]; ok {
        val = val
        hako[j-1] = append(hako[j-1], i)
    } else {
        hako[j-1] = make([]int, 0)
        hako[j-1] = append(hako[j-1], i)
    }

    if val, ok := num[i-1]; ok {
        val = val
        num[i-1] = append(num[i-1], j)
    } else {
        num[i-1] = make([]int, 0)
        num[i-1] = append(num[i-1], j)
    }
}

func query2(hako map[int] []int) {
    i := nextInt()
    sort.Slice(hako[i-1], func(j,k int) bool { return hako[i-1][j] < hako[i-1][k] } )

    length := len(hako[i-1])
    for j:=0; j < length; j++ {
        if j != length -1 {
            fmt.Printf("%d ", hako[i-1][j] )
        } else if j == length -1{
            fmt.Println(hako[i-1][j])
        }
    }
}

func query3(hako map[int] []int) {
    i := nextInt()
    sort.Slice(hako[i-1], func(j,k int) bool { return hako[i-1][j] < hako[i-1][k] } )

    length := len(hako[i-1])
    pre_num := 0
    str := ""
    for j:=0; j < length; j++ {
        if  hako[i-1][j]  > pre_num{
            str += strconv.Itoa(hako[i-1][j]) + " "
        } 
        pre_num = hako[i-1][j]
    }

    fmt.Println(str[:len(str)-1])
}