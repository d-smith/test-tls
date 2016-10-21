This page has details about the code: http://www.hydrogen18.com/blog/your-own-pki-tls-golang.html

Note that openssl can be used to produce the keys and certs used in this code. Some raw commands:

<pre>
mkdir openssl
cd openssl/
mkdir ca
openssl genrsa -aes256 -out ca/ca.key 4096 chmod 400 ca/ca.key
(remember your passphrase)

openssl req -new -x509 -sha256 -days 730 -key ca/ca.key -out ca/ca.crt
(I used EZ for the org name)
mkdir server
openssl genrsa -out server/localhost.key 2048
openssl req -new -key server/localhost.key -sha256 -out server/localhost.csr
(I used localhost for the CN)
openssl x509 -req -days 3650 -sha256 -in server/localhost.csr -CA ca/ca.crt -CAkey ca/ca.key -set_serial 1 -out server/localhost.crt
mkdir client
openssl genrsa -out client/okguy.key 2048
openssl req -new -key client/okguy.key -out client/okguy.csr
(I used okguy for the CN)
openssl x509 -req -days 3650 -sha256 -in client/okguy.csr -CA ca/ca.crt -CAkey ca/ca.key -set_serial 2 -out client/okguy.crt
</pre>


With the keys and certs created, run the server from the service dir like this:

<pre>
go run main.go ../openssl/server/localhost.key ../openssl/server/localhost.crt ../openssl/ca/ca.crt 
</pre>

Run the client from the client directory like this:

<pre>
go run main.go ../openssl/client/okguy.key ../openssl/client/okguy.crt ../openssl/ca/ca.crt 
</pre>


