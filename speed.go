package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"net/http"
	"os"
	"time"

	"github.com/surol/speedtest-cli/speedtest"
)

type Responsespeed struct {
	ISP           string
	IP            string
	Downloadspeed float64
	Uploadspeed   float64
}

func version() {
	fmt.Print(speedtest.Version)
}

func usage() {
	fmt.Fprint(os.Stderr, "Command line interface for testing internet bandwidth using speedtest.net.\n\n")
	flag.PrintDefaults()
}

var client *speedtest.Client
var config1 *speedtest.Config
var opts *speedtest.Opts

func Init() {
	opts = speedtest.ParseOpts()

	switch {
	case opts.Help:
		usage()
		return
	case opts.Version:
		version()
		return
	}

	client = speedtest.NewClient(opts)

	if opts.List {
		servers, err := client.AllServers()
		if err != nil {
			log.Fatal("Failed to load server list: %v\n", err)
		}
		fmt.Println(servers)
		return
	}
	config1, err := client.Config()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config1)
}
func Speed_Service(w http.ResponseWriter, r *http.Request) {

	client.Log("Testing from %s (%s)...\n", config1.Client.ISP, config1.Client.IP)

	server := selectServer(opts, client)

	downloadSpeed := server.DownloadSpeed()
	getDownload := reportSpeed(opts, "Download", downloadSpeed)
	fmt.Println("getDownload", getDownload)

	uploadSpeed := server.UploadSpeed()
	getupload := reportSpeed(opts, "Upload", uploadSpeed)
	fmt.Println("getupload", getupload)

	Senddata := &Responsespeed{}
	Senddata.ISP = config1.Client.ISP
	Senddata.IP = config1.Client.IP
	Senddata.Downloadspeed = getDownload
	Senddata.Uploadspeed = getupload

	Responsedata, _ := json.Marshal(Senddata)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Responsedata)

}

func reportSpeed(opts *speedtest.Opts, prefix string, speed int) (value float64) {

	if opts.SpeedInBytes {
		fmt.Printf("%s: %.2f MiB/s\n", prefix, float64(speed)/(1<<20))
		value = float64(speed) / (1 << 20)

	} else {
		fmt.Printf("%s: %.2f Mib/s\n", prefix, float64(speed)/(1<<17))
		return float64(speed) / (1 << 17)
		value = float64(speed) / (1 << 17)
	}
	return

}

func selectServer(opts *speedtest.Opts, client *speedtest.Client) (selected *speedtest.Server) {
	if opts.Server != 0 {
		servers, err := client.AllServers()
		if err != nil {
			log.Fatal("Failed to load server list: %v\n", err)
			return nil
		}
		selected = servers.Find(opts.Server)
		if selected == nil {
			log.Fatal("Server not found: %d\n", opts.Server)
			return nil
		}
		selected.MeasureLatency(speedtest.DefaultLatencyMeasureTimes, speedtest.DefaultErrorLatency)
	} else {
		servers, err := client.ClosestServers()
		if err != nil {
			log.Fatal("Failed to load server list: %v\n", err)
			return nil
		}
		selected = servers.MeasureLatencies(
			speedtest.DefaultLatencyMeasureTimes,
			speedtest.DefaultErrorLatency).First()
	}

	if opts.Quiet {
		log.Fatal("Ping: %d ms\n", selected.Latency/time.Millisecond)
	} else {
		client.Log("Hosted by %s (%s) [%.2f km]: %d ms\n",
			selected.Sponsor,
			selected.Name,
			selected.Distance,
			selected.Latency/time.Millisecond)
	}

	return selected
}
