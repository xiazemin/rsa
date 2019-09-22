package decode

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"github.com/golang/go/src/fmt"
	"crypto"
)

//rsa
type rsaUtils struct {
	PrivateKey []byte
	PublicKey  []byte
}

//当前类的指针
var rasUtilsClass *rsaUtils

//同步锁
var rasUtilsClassonce sync.Once

//rsa工具
func RSAUtils() *rsaUtils {
	rasUtilsClassonce.Do(func() {
		rasUtilsClass = new(rsaUtils)
	})
	//返回处理对象
	return rasUtilsClass
}

func (r *rsaUtils) init() bool {
	cwd,err:=os.Getwd()
	if err!=nil{
		fmt.Println(err)
		return false
	}
	//cwd+="/src/github.com/xiazemin/rsa"
	//读取文件
	pubkey, err := r.readFile(cwd+"/public.pem")
	if err != nil {
		fmt.Println(err)
		return false
	}
	r.PublicKey = pubkey

	//获取私钥
	prikey, err := r.readFile(cwd+"/private.pem")
	if err != nil {
		fmt.Println(err)
		return false
	}
	r.PrivateKey = prikey
	return true
}

func (r *rsaUtils) readFile(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// 加密
func (r *rsaUtils) RsaEncrypt(origData []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	//MarshalPKCS1PrivateKey  not ParsePKCS8PrivateKey
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	priKey := priv //.(*rsa.PrivateKey)
	return rsa.SignPKCS1v15(rand.Reader, priKey,crypto.Hash(0), origData)

}

// 解密
func (r *rsaUtils) RsaDecrypt(ciphertext []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//MarshalPKIXPublicKey
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return publicDecrypt(pub, crypto.Hash(0), nil, ciphertext)//rsa.DecryptPKCS1v15(rand.Reader, pub, ciphertext)
}