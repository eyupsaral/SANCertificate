package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eyupsaral/crypto/acme"
	"github.com/eyupsaral/crypto/acme/autocert"
)

//DefaultACME staging link
const DefaultACME = "https://acme-staging-v02.api.letsencrypt.org/directory"

func main() {
   createHTTPServer()
}

func  createHTTPServer() {
	var certman *autocert.Manager

	certman = &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: func(ctx context.Context, host string) error {
			fmt.Printf("Requested Host %s\n", host)
			return nil
		},
		Client: &acme.Client{DirectoryURL: DefaultACME},
		Cache: autocert.DirCache("./"),
		WWWtoDomain: true,
		SANHosts: func(ctx context.Context, host string) []string {
			return []string{"sancertificate.tk", "www.sancertificate.tk"}
		},
	}

	HTTPSServer := makeHTTPServer()
	HTTPSServer.TLSConfig = &tls.Config{GetCertificate: certman.GetCertificate}
	HTTPSServer.Addr = ":443"

	go func() {
		fmt.Printf("Starting HTTPS server on %s\n", HTTPSServer.Addr)
		err := HTTPSServer.ListenAndServeTLS("", "")
		if err != nil {
			log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
		}
	}()

	httpSrv := makeHTTPServer()

	if certman != nil {
		httpSrv.Handler = certman.HTTPHandler(httpSrv.Handler)
	}

	httpSrv.Addr = ":80"
	fmt.Printf("Starting HTTP server on %s\n", httpSrv.Addr)
	err := httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
	}
}

func makeServerFromMux(mux *http.ServeMux) *http.Server {
	return &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
		Handler:           mux,
	}
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", http.HandlerFunc(replyClientRequest))
	return makeServerFromMux(mux)
}

//ForwardClientRequest ddd
func replyClientRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Working web server"))
}