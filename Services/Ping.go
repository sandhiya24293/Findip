package Services

import (
	Db "Findip/Common/DB/Mysql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	//"os"
	"os/exec"
	"strings"
	"time"

	//"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
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
func heartBeatmonthly(ip string, Email string, ticker *time.Ticker) (int, int) {
	log.Println("Pinging------")
	var Monthlypass int = 0
	var Monthlyfail int = 0

	for range ticker.C {
		out, _ := exec.Command("ping", ip, "-c 5", "-i 3", "-w 10").Output()
		//out, _ := exec.Command("ping", ip).Output()
		if strings.Contains(string(out), "Destination Host Unreachable") {
			fmt.Println(ip, "-TANGO DOWN send email to", Email)
			Monthlyfail = Monthlyfail + 1

		} else {
			fmt.Println(ip, "IT'S ALIVEEE")
			Monthlypass = Monthlypass + 1
			fmt.Println(Monthlypass)

		}

	}

	return Monthlypass, Monthlyfail

}
func Monitorip(w http.ResponseWriter, r *http.Request) {
	var GetIpvalue Monitor

	err := json.NewDecoder(r.Body).Decode(&GetIpvalue)
	if err != nil {
		fmt.Println("Error on Get particular details", err)
	}
	Db.Insertip(GetIpvalue.Ip, GetIpvalue.Mailid)
	ticker := time.NewTicker(2 * time.Second)
	var getpass int
	var getfails int

	go func() {
		getpass, getfails = heartBeatmonthly(GetIpvalue.Ip, GetIpvalue.Mailid, ticker)

	}()

	time.Sleep(time.Second * 10)
	fmt.Println("loop finished ", GetIpvalue.Ip, GetIpvalue.Mailid)
	fmt.Println("got pass and fail count  ", getpass, getfails)
	//SendReport("Monthly", GetIpvalue.Ip, GetIpvalue.Mailid)
	ticker.Stop()

}
func SendReport(Monthly string, Monthlypassreport int, monthylyfail int, addr string, number string) {
	var htmlContent string
	//	sendkey := os.Getenv("SENDGRID_API_KEYGO")
	//	from := mail.NewEmail("E3 Shopping", "sandhiyabalakrishnan6@gmail.com")
	//	subject := "E3 NOTIFICATION - New Order Recieved!"
	//	to := mail.NewEmail("Foxutech Tool", "sandhiyabalakrishnan6@gmail.com")
	//	plainTextContent := "sandhiyabalakrishnan6@gmail.com"

	switch Monthly {
	case "Monthly":

		htmlContent = "<div><b style='font-size:15px'>: </b></div><br> " +
			"<div style='font-style:sans-serif'>LOGINID - " +
			"<div style='font-style:sans-serif'>ADDRESS - " +
			"</div><div style='font-style:sans-serif'>DATE -" +
			"</div><div style='font-style:sans-serif'>TOTAL AMOUNT -" +
			"</div><div style='font-style:sans-serif'>NO OF PRODUCTS -" +
			"</div><table class='table' border='1' style='padding:5px;font-style:sans-serif'><tbody >" + "<tr style='border-bottom:1pt solid black;'><th >PRODUCT</th><th>RATE</th><th>WEIGHT</th></tr>" +
			"</tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	case "Failure":

		htmlContent = "<div><b style='font-size:15px'>E3 NEW ORDER : </b></div><br> " +
			"<div style='font-style:sans-serif'>NAME  - " +
			"<div style='font-style:sans-serif'>ADDRESS - " +
			"</div><div style='font-style:sans-serif'>NUMBER -" +
			"</div><div style='font-style:sans-serif'>PRODUCT  -" +
			"</div><div style='font-style:sans-serif'>MESSAGE" +
			"</div></tbody></table><br><div>Please Check E3 Admin Panel for more detail ...!</div>"

	default:
		fmt.Println("default case")

	}
	fmt.Println(htmlContent)
	//message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	//	client := sendgrid.NewSendClient(sendkey)
	//	response, err := client.Send(message)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(response.StatusCode)
	//		fmt.Println(response.Body)
	//		fmt.Println(response.Headers)

	//	}

}
