package Services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	_ "strings"
)

type Getipstruct struct {
	Getipfromuser string
}

func Iplookup(Ipaddress string) (resp *http.Response) {
	var client http.Client
	clienturl := "http://api.whatismyip.com/ip-address-lookup.php?key=89cd62c5ed209af9513765d85f690fef&input=" + Ipaddress + "&output=json"
	resp, _ = client.Get(clienturl)
	fmt.Println(&resp.Body)
	return resp

}

func GETIP(w http.ResponseWriter, r *http.Request) {

	Ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	Senddata, err := json.Marshal(Ip)

	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func GETProxy(w http.ResponseWriter, r *http.Request) {
	Ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	var client http.Client

	clienturl := "https://api.whatismyip.com/proxy.php?key=89cd62c5ed209af9513765d85f690fef" + Ip + "&output=json"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}
func GETIPdetails(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	resp := Iplookup(GetIpvalue.Getipfromuser)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}
func Blocklist(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	var client http.Client

	clienturl := "http://api.whatismyip.com/domain-black-list.php?key=89cd62c5ed209af9513765d85f690fef&input=" + GetIpvalue.Getipfromuser + "&output=json"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}
func HostnameLookup(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	var client http.Client

	clienturl := "http://api.whatismyip.com/host-name.php?key=89cd62c5ed209af9513765d85f690fef&input=" + GetIpvalue.Getipfromuser + "&output=json"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}
func IPwhoislookup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside lookup")
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	var client http.Client

	clienturl := "http://api.whatismyip.com/whois.php?key=89cd62c5ed209af9513765d85f690fef&input=" + GetIpvalue.Getipfromuser + "&output=json"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}

func Serverheadercheck(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	var client http.Client

	clienturl := "http://api.whatismyip.com/server-headers.php?key=89cd62c5ed209af9513765d85f690fef&input=http://" + GetIpvalue.Getipfromuser + "&output=json"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}

func Useragent(w http.ResponseWriter, r *http.Request) {

	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	//	var client http.Client

	//	clienturl := "http://api.whatismyip.com/user-agent.php?key=89cd62c5ed209af9513765d85f690fef&input=176.111.105.86&output=json"
	//	resp, _ := client.Get(clienturl)

	//defer resp.Body.Close()
	fmt.Printf("r: %+v\n", r.UserAgent())
	addr := r.UserAgent()

	Senddata, err := json.Marshal(addr)

	//	if err != nil {
	//		log.Error(err)
	//	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Dnslookup(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println("value", GetIpvalue.Getipfromuser)
	var client http.Client

	clienturl := "http://api.whatismyip.com/host-name.php?key=89cd62c5ed209af9513765d85f690fef&input=" + GetIpvalue.Getipfromuser + "&output=json"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}
}
func ReverseDnslookup(w http.ResponseWriter, r *http.Request) {
	var Gethost Getipstruct
	err := json.NewDecoder(r.Body).Decode(&Gethost)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	fmt.Println(Gethost)
	addr, err := net.LookupIP(Gethost.Getipfromuser)
	fmt.Println(addr, err)
	Senddata, err := json.Marshal(addr)

	//	if err != nil {
	//		log.Error(err)
	//	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)
}

func Sslchecker(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Getipstruct
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}

	fmt.Println("value", GetIpvalue.Getipfromuser)
	var client http.Client
	clienturl := "https://www.sslshopper.com/assets/snippets/sslshopper/ajax/ajax_check_ssl.php?hostname=" + GetIpvalue.Getipfromuser + "&recaptcha_challenge_field=&recaptcha_response_field=&rand=371"
	resp, _ := client.Get(clienturl)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")
		w.Write(bodyBytes)
	}

}
