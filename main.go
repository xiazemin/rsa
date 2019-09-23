package main

import (
	"fmt"
	"github.com/xiazemin/rsa/gen"
)

func main() {
	fmt.Println(gen.GenRsaKey(1024))
}

