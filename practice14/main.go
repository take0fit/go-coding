package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	str := "HELLO WORLD"
	fmt.Println(toLowerCase(str))
}
