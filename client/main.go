package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	certPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile("./tls/root.crt")
	if err != nil {
		log.Fatal("Could not load CA certificate!")
		return
	}

	if ok := certPool.AppendCertsFromPEM(rootCA); !ok {
		panic("failed to append cert from pem")
	}

	clientCert, err := tls.LoadX509KeyPair("./tls/client.crt", "./tls/client.key")
	if err != nil {
		log.Fatal("Could not load server certificate!")
		return
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName:   "test-server",
			RootCAs:      certPool,
			Certificates: []tls.Certificate{clientCert},
			ClientAuth:   tls.RequireAndVerifyClientCert,
		},
	}
	c := http.Client{
		Transport: transport,
	}
	resp, err := c.Get("https://localhost:8080/healthcheck")
	if err != nil {
		log.Fatal(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(b))
}
