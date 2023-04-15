package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    nextInt()
    s := nextStr()

	if strings.Index(s, "o") >= 0 && strings.Index(s, "x") == -1 {
		fmt.Println("Yes")
	} else { 
		fmt.Println("No")
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

func nextStr() string {
    sc.Scan()
    return sc.Text()
}
