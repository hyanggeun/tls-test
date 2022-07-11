package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	e := gin.Default()

	e.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	certPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile("./tls/root.crt")
	if err != nil {
		log.Fatal("Could not load CA certificate!")
		return
	}

	if ok := certPool.AppendCertsFromPEM(rootCA); !ok {
		panic("failed to append cert from pem")
	}

	serverCert, err := tls.LoadX509KeyPair("./tls/server.crt", "./tls/server.key")
	if err != nil {
		log.Fatal("Could not load server certificate!")
		return
	}
	config := &tls.Config{
		//InsecureSkipVerify: false,
		ServerName:   "test-server",
		RootCAs:      certPool,
		Certificates: []tls.Certificate{serverCert},
	}

	s := &http.Server{
		TLSConfig: config,

		Handler: e,
		Addr:    ":8080",
	}
	if err := s.ListenAndServeTLS("./tls/server.crt", "./tls/server.key"); err != nil {

		log.Fatal(err)
		return
	}

}
