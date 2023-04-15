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
    a := make([][]int, n)
    for i:= 0; i < n; i ++ {
        a[i] = make([]int, n)
    }
    b := make([][]int, n)
    for i:= 0; i < n; i ++ {
        b[i] = make([]int, n)
    }

    a = initalSlice(a, n)
    b = initalSlice(b, n)

    output := "No"

    for i := 0; i < 4; i++ {
        if matrixCheck(a, b, n){
            output = "Yes"
            break
        }
        a = rotateSlice(a, n)
    }

    fmt.Println(output)
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func initalSlice(slice [][]int, n int) [][]int {
    for i := 0; i < n; i++ {
        for j:= 0; j < n; j++ {
            slice[i][j] = nextInt()
        }
    }

    return slice
}

func matrixCheck(a [][]int, b [][]int, n int) bool {
    for i := 0; i < n; i++ {
        for j:= 0; j < n; j++ {
            if a[i][j] == 1 && b[i][j] != 1 {
                return false
            }
        }
    }

    return true
}


func rotateSlice(a [][]int, n int) [][]int {
    tmpa := make([][]int, n)
    for i:= 0; i < n; i ++ {
        tmpa[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for j:= 0; j < n; j++ {
            tmpa[i][j] = a[n-1-j][i]
        }
    }

    return tmpa
}
