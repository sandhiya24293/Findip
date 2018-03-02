package Services

import (
	Db "Findip/Common/DB/Mysql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type Ping struct {
	IP string
}
type Monitor struct {
	Ip     string
	Mailid string
}

func GoPing(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Ping
	var response string
	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}

	out, _ := exec.Command("ping", GetIpvalue.IP, "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
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

	Senddata, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
func heartBeatmonthly(ip string, Email string) {
	log.Println("Pinging------")

	for range time.Tick(time.Second * 2) {
		out, _ := exec.Command("ping", ip, "-c 5", "-i 3", "-w 10").Output()
		//out, _ := exec.Command("ping", ip).Output()
		if strings.Contains(string(out), "Destination Host Unreachable") {
			fmt.Println(ip, "-TANGO DOWN send email to", Email)
		} else {
			fmt.Println(ip, "IT'S ALIVEEE")
		}
	}
}
func Monitorip(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Monitor

	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	Db.Insertip(GetIpvalue.Ip, GetIpvalue.Mailid)

	go heartBeatmonthly(GetIpvalue.Ip, GetIpvalue.Mailid)
	time.Sleep(time.Second * 100)

}
