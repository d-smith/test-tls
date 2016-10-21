package main

import "crypto/tls"
import "github.com/d-smith/test-tls"
import "fmt"
import "io"
import "os"
import "log"

//import "net"

func main() {

	config := common.MustGetTlsConfiguration()

	//conn, err := tls.Dial("tcp", "localhost:51000", config)
	println("assuming nginx proxy")
	conn, err := tls.Dial("tcp", "localhost:5000", config)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = conn.Handshake()
	if err != nil {
		fmt.Printf("Failed handshake:%v\n", err)
		return
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Printf("Failed receiving data:%v\n", err)
	}

	conn.Close()
}
