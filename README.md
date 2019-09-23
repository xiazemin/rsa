openssl genrsa -out server.key 2048

server.key

openssl req -new -x509 -key server.key -out server.crt -days 365

server.crt

https://127.0.0.1:8081/

我们用http.ListenAndServeTLS替换掉了http.ListenAndServe，就将一个HTTP Server转换为HTTPS Web Server了。不过ListenAndServeTLS 新增了两个参数certFile和keyFile

也可以使用curl工具验证这个HTTPS server：

curl -k https://localhost:8081
Hi, This is an example of http service in golang!
注意如果不加-k，curl会报如下错误：

curl: (60) SSL certificate problem: self signed certificate
More details here: https://curl.haxx.se/docs/sslcerts.html

$ openssl rsa -in server.key -out server.key.public
writing RSA key

$openssl genrsa -out ca.key 2048
ca.key

$ openssl req -x509 -new -nodes -key ca.key -subj "/CN=xzm.com" -days 5000 -out ca.crt
ca.crt

接下来，生成server端的私钥，生成数字证书请求，并用我们的ca私钥签发server的数字证书：
openssl genrsa -out server.key 2048

openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
server.csr
生成Certificate Sign Request，CSR，证书签名请求。

openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
server.crt

Signature ok
subject=/CN=localhost
Getting CA Private Key

$ openssl genrsa -out client.key 2048
client.key

$ openssl req -new -key client.key -subj "/CN=xzm_cn" -out client.csr
client.csr

$ openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000
Signature ok
subject=/CN=xzm_cn
Getting CA Private Key


