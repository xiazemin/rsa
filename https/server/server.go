// gohttps/2-https/server.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	cwd,err:=os.Getwd()
	fmt.Println(cwd)
	if err!=nil{
		return
	}
	cwd+="/src/github.com/xiazemin/rsa"
	//"" ""
	fmt.Println(http.ListenAndServeTLS(":8081", cwd+"/server.crt",
		cwd+"/server.key", nil))
}