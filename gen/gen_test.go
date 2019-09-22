package gen

import (
	"fmt"
	"testing"
)

func TestGet(t*testing.T) {
	//rsa 密钥文件产生
	fmt.Println(GenRsaKey(1024))
}