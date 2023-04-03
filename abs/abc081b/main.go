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
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = nextInt()
    }

	i := 0
	pow := 2
	for isDivide(nums, pow) {
		pow = pow * 2;
		i++;
    }

    fmt.Println(i)
}

func isDivide(nums[] int ,n int) bool{
	for i := 0; i < len(nums); i++ {
        if nums[i] % n != 0 {
			return false
		}
    }
	return true
}


func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}
