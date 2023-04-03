package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Printf("%d\n", strings.Count(s, "1"))
} 