package Services

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os/exec"

	"time"
)

type Ping struct {
	IP string
}

func GoPing(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Ping
	var response string
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}

	_, err = exec.Command("ping", GetIpvalue.IP).Output()
	fmt.Println(err)
	if err != nil {
		fmt.Println("TANGO DOWN")
		response = "SERVER DOWN"

	} else {
		fmt.Println("IT'S ALIVEEE")
		response = "SERVER  ALIVEEE"

	}
	Senddata, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

func HostPing(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Ping
	var response string

	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}

	portNum := "80"
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second

	_, err = net.DialTimeout("tcp", GetIpvalue.IP+":"+portNum, timeOut)
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}

	if err != nil {

		fmt.Println("HOST DOWN")
		response = "HOST DOWN"
	} else {
		response = "HOST ALIVEEE"

	}

	//	fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", hostName, int64(seconds))
	//fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	//fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())
	Senddata, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
