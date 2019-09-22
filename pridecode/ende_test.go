package decode

import (
	"testing"
	"github.com/golang/go/src/fmt"
)

func TestEnDe(t *testing.T){
    ru:=RSAUtils()
	fmt.Println(ru.init())
	fmt.Println(ru)
	//公钥加密 私钥解密
	eby,err:=ru.RsaEncrypt([]byte("hello world"),ru.PublicKey)
	fmt.Println(err)
	dby,err:=ru.RsaDecrypt(eby,ru.PrivateKey)
	fmt.Println(string(dby),err)
}
