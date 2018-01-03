package main

//import (
//   "fmt"
////    "net"
//	//"strings"

//	"os/exec"
//	//"log"

//)
//const targetIP = "8.8.8.8"

//func main() {
//	// obtained from ping -c 1 stackoverflow.com, should print "stackoverflow.com"
//	addr, err := net.LookupAddr("198.252.206.16")
//	fmt.Println(addr, err)

//	addr, err := net.LookupIP("stackoverflow.com")
//	if err != nil {
//		fmt.Println("Unknown host")
//	} else {
//		fmt.Println("IP address: ", addr)
//	}

//	out, _ := exec.Command("ping", "192.168.0.111", "-c 5", "-i 3", "-w 10").Output()
//	fmt.Println(string(out))
//	if strings.Contains(string(out), "Destination Host Unreachable") {
//		fmt.Println("TANGO DOWN")
//	} else {
//		fmt.Println("IT'S ALIVEEE")
//	}

//}
