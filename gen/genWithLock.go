package gen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"flag"
	"log"
	"os"
	"sync"
)

//当前类的指针
var rsaGkey *rsaGenerateKey

//同步锁
var rsaGenerateKeyonce sync.Once

//实现单例
type rsaGenerateKey struct {
}

//生成密钥
func RSAGenerateKeys() *rsaGenerateKey {
	rsaGenerateKeyonce.Do(func() {
		rsaGkey = new(rsaGenerateKey)
	})
	return rsaGkey
}

//生成文件
func (r *rsaGenerateKey) GenerateKey() {
	var bits int
	flag.IntVar(&bits, "b", 1024, "密钥长度，默认为1024位")
	if err := r.GenRsaKey(bits); err != nil {
		log.Fatal("密钥文件生成失败！")
	}
}

//生成公私私钥--return（私钥，公私，错误）
func (r *rsaGenerateKey) Generate() ([]byte, []byte, error) {
	//定义变量并赋值
	var bits int
	bits = 1024 //密钥长度，默认为128位"
	//主线程中只用定义一次
	//flag.IntVar(&bits, "bint", 512, "密钥长度，默认为128位")
	// 生成私钥文件
	privateKey, err := rsa.GenerateMultiPrimeKey(rand.Reader, 3, bits)
	if err != nil {
		return nil, nil, err
	}
	//生成私钥
	derStream := MarshalPKCS8PrivateKey(privateKey)
	blockpri := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	blockpub := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	return pem.EncodeToMemory(blockpri), pem.EncodeToMemory(blockpub), nil
}

//转化成pkcs8
func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)

	k, err := asn1.Marshal(info)
	if err != nil {
		log.Panic(err.Error())
	}
	return k
}

//生成文件
func (r *rsaGenerateKey) GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}