package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tcnksm/go-httpstat"
)

type Awsout struct {
	Servername       string
	Ping             string
	DNSLookup        int
	TCPConnection    int
	TLSHandshake     int
	ServerProcessing int
	ContentTransfer  int
}

func Awsping(w http.ResponseWriter, r *http.Request) {

	var Awsresponse []Awsout

	t := []string{
		"https://ec2.sa-east-1.amazonaws.com/ping",
		"http://ec2.us-east-2.amazonaws.com/ping",
		"http://ec2.us-west-1.amazonaws.com/ping",
		"http://ec2.us-west-2.amazonaws.com/ping",
		"http://ec2.ca-central-1.amazonaws.com/ping",
		"http://ec2.eu-west-1.amazonaws.com/ping",
		"http://ec2.eu-central-1.amazonaws.com/ping",
		"http://ec2.eu-west-2.amazonaws.com/ping",
		"http://ec2.ap-southeast-1.amazonaws.com/ping",
		"http://ec2.ap-southeast-2.amazonaws.com/ping",
		"http://ec2.ap-northeast-2.amazonaws.com/ping",
		"http://ec2.ap-northeast-1.amazonaws.com/ping",
		"http://ec2.ap-south-1.amazonaws.com/ping",
		"http://ec2.sa-east-1.amazonaws.com/ping",
	}
	// Create a new HTTP request

	for _, server := range t {
		var Response Awsout
		var bodyString string

		req, err := http.NewRequest("GET", server, nil)
		if err != nil {
			log.Fatal(err)
		}

		// Create a httpstat powered context
		var result httpstat.Result
		ctx := httpstat.WithHTTPStat(req.Context(), &result)
		req = req.WithContext(ctx)
		// Send request by default HTTP client
		client := http.DefaultClient
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode == http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			bodyString = string(bodyBytes)
			bodyString = strings.TrimSuffix(bodyString, "\n")
			log.Println("bodystring", bodyString)

		}
		if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
			log.Fatal(err)
		}
		res.Body.Close()

		//end := time.Now()
		// Show the results
		log.Printf("responsebody", res.Body)
		log.Printf("DNS lookup: %d ms", int(result.DNSLookup/time.Millisecond))
		log.Printf("TCP connection: %d ms", int(result.TCPConnection/time.Millisecond))
		log.Printf("TLS handshake: %d ms", int(result.TLSHandshake/time.Millisecond))
		log.Printf("Server processing: %d ms", int(result.ServerProcessing/time.Millisecond))
		log.Printf("Content transfer: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))
		Response.Ping = bodyString
		Response.Servername = server
		Response.DNSLookup = int(result.DNSLookup / time.Millisecond)
		Response.TCPConnection = int(result.TCPConnection / time.Millisecond)
		Response.TLSHandshake = int(result.TLSHandshake / time.Millisecond)
		Response.ServerProcessing = int(result.ServerProcessing / time.Millisecond)
		Response.ContentTransfer = int(result.ContentTransfer(time.Now()) / time.Millisecond)
		Awsresponse = append(Awsresponse, Response)

	}

	Senddata, err := json.Marshal(Awsresponse)

	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
