package decode

import (
	"testing"
	"github.com/golang/go/src/fmt"
)

func TestEnDe(t *testing.T){
    ru:=RSAUtils()
	fmt.Println(ru.init())
	fmt.Println(ru)
	//私钥加密(签名) 公钥解密
	eby1,err:=ru.RsaEncrypt([]byte("hello world"),ru.PrivateKey)
	fmt.Println(err)
	dby1,err:=ru.RsaDecrypt(eby1,ru.PublicKey)
	fmt.Println(string(dby1),err)
}
