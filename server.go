package main

import (
	Service "Findip/Services"
	"crypto/tls"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	contenttypeJSON = "application/json; charset=utf-8"
)

func Serve() bool {

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./Ip-tool/"))
	//http.Handle("/", fs)
	router.PathPrefix("/Ip-tool/").Handler(http.StripPrefix("/Ip-tool/", fs))
	fs1 := http.FileServer(http.Dir("./Findipssl/"))
	//http.Handle("/", fs)
	router.PathPrefix("/Findipssl/").Handler(http.StripPrefix("/Findipssl/", fs1))
	fs2 := http.FileServer(http.Dir("./assets/"))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs2))

	fs3 := http.FileServer(http.Dir("./Docker/"))

	router.PathPrefix("/Docker/").Handler(http.StripPrefix("/Docker/", fs3))
	fs4 := http.FileServer(http.Dir("./IP/"))

	router.PathPrefix("/IP/").Handler(http.StripPrefix("/IP/", fs4))
	fs5 := http.FileServer(http.Dir("./Filestorage/"))

	router.PathPrefix("/Filestorage/").Handler(http.StripPrefix("/Filestorage/", fs5))

	router.HandleFunc("/GetIP", Service.GETIP)
	router.HandleFunc("/GetProxy", Service.GETProxy)
	router.HandleFunc("/GetIPdetails", Service.GETIPdetails)
	router.HandleFunc("/Blocklist", Service.Blocklist)
	router.HandleFunc("/HostnameLookup", Service.HostnameLookup)
	router.HandleFunc("/IPwhoislookup", Service.IPwhoislookup)
	router.HandleFunc("/Serverheadercheck", Service.Serverheadercheck)
	router.HandleFunc("/Useragent", Service.Useragent)
	router.HandleFunc("/ReverseDnslookup", Service.ReverseDnslookup)
	router.HandleFunc("/Dnslookup", Service.Dnslookup)
	router.HandleFunc("/SSlchecker", Service.Sslchecker)
	router.HandleFunc("/Speed_Service", Speed_Service)
	router.HandleFunc("/Dockerfile", Service.WriteDockerFile)
	router.HandleFunc("/Awsping", Service.Awsping)
	router.HandleFunc("/Googleping", Service.Googleping)
	router.HandleFunc("/GoPing", Service.GoPing)
	router.HandleFunc("/HostPing", Service.HostPing)
	router.HandleFunc("/Monitorip", Service.Monitorip)
	router.HandleFunc("/Transfer", Service.Transfer)
	router.HandleFunc("/TraceAPI", Service.TraceAPI)
	router.HandleFunc("/Shorterturl", Service.ShorternURL)

	//For HTTPS

	// Default server - non-trusted for debugging

	serverhttp := func() {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*", "http://develop.rentmatics.com", "http://api.msg91.com"},
			AllowCredentials: true,
		})
		handler := c.Handler(router)

		fmt.Println("Server should be available at http", config.Port)
		fmt.Println(http.ListenAndServe(config.Port, handler))
	}

	// Setup TLS parameters for trusted server implementation
	if config.SSL && config.Key != "" && config.Cert != "" {
		// Setup TLS parameters
		tlsConfig := &tls.Config{
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS12,
			Certificates: make([]tls.Certificate, 1),
		}

		var err error
		// Setup API server private key and certificate
		tlsConfig.Certificates[0], err = tls.X509KeyPair([]byte(config.Cert), []byte(config.Key))
		if err != nil {
			fmt.Println("Error during decoding service key and certificate:", err)
			return false
		}

		tlsConfig.BuildNameToCertificate()

		https := &http.Server{
			Addr:      config.Https_port,
			TLSConfig: tlsConfig,
			Handler:   router,
		}

		// Trusted server implementation
		server := func() {
			fmt.Println("Server should be available at https", config.Https_port)
			fmt.Println(https.ListenAndServeTLS("", ""))
		}
		go server()
	}

	// Schedule API server
	go serverhttp()

	return true
}
