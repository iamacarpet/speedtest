package tests

import (
	"strings"
	"fmt"
	"log"
)

import (
	"github.com/zpeters/speedtest/debug"
	"github.com/zpeters/speedtest/print"
	"github.com/zpeters/speedtest/misc"
	"github.com/zpeters/speedtest/sthttp"
)

func DownloadTest(server sthttp.Server, algotype string) float64 {
	var urls []string
	var maxSpeed float64
	var avgSpeed float64

	// http://speedtest1.newbreakcommunications.net/speedtest/speedtest/
	dlsizes := []int{350, 500, 750, 1000, 1500, 2000, 2500, 3000, 3500, 4000}

	for size := range dlsizes {
		url := server.Url
		splits := strings.Split(url, "/")
		baseUrl := strings.Join(splits[1:len(splits)-1], "/")
		randomImage := fmt.Sprintf("random%dx%d.jpg", dlsizes[size], dlsizes[size])
		downloadUrl := "http:/" + baseUrl + "/" + randomImage
		urls = append(urls, downloadUrl)
	}

	if !debug.QUIET {
		log.Printf("Testing download speed")
	}

	for u := range urls {

		if debug.DEBUG {
			fmt.Printf("Download Test Run: %s\n", urls[u])
		}
		dlSpeed := sthttp.DownloadSpeed(urls[u])
		if !debug.QUIET && !debug.DEBUG {
			fmt.Printf(".")
		}
		if debug.DEBUG {
			log.Printf("Dl Speed: %v\n", dlSpeed)
		}

		if algotype == "max" {
			if dlSpeed > maxSpeed {
				maxSpeed = dlSpeed
			}
		} else {
			avgSpeed = avgSpeed + dlSpeed
		}

	}

	if !debug.QUIET {
		fmt.Printf("\n")
	}

	if algotype == "max" {
		return maxSpeed
	} else {
		return avgSpeed / float64(len(urls))
	}
}


func UploadTest(server sthttp.Server, algotype string) float64 {
	// https://github.com/sivel/speedtest-cli/blob/master/speedtest-cli
	var ulsize []int
	var maxSpeed float64
	var avgSpeed float64

	ulsizesizes := []int{
		int(0.25 * 1024 * 1024),
		int(0.5 * 1024 * 1024),
		int(1.0 * 1024 * 1024),
		int(1.5 * 1024 * 1024),
		int(2.0 * 1024 * 1024),
	}

	for size := range ulsizesizes {
		ulsize = append(ulsize, ulsizesizes[size])
	}

	if !debug.QUIET {
		log.Printf("Testing upload speed")
	}

	for i := 0; i < len(ulsize); i++ {
		if debug.DEBUG {
			fmt.Printf("Upload Test Run: %v\n", i)
		}
		r := misc.Urandom(ulsize[i])
		ulSpeed := sthttp.UploadSpeed(server.Url, "text/xml", r)
		if !debug.QUIET && !debug.DEBUG {
			fmt.Printf(".")
		}

		if algotype == "max" {
			if ulSpeed > maxSpeed {
				maxSpeed = ulSpeed
			}
		} else {
			avgSpeed = avgSpeed + ulSpeed
		}

	}

	if !debug.QUIET {
		fmt.Printf("\n")
	}

	if algotype == "max" {
		return maxSpeed
	} else {
		return avgSpeed / float64(len(ulsizesizes))
	}
}


func FindServer(id string, serversList []sthttp.Server) sthttp.Server {
	var foundServer sthttp.Server
	for s := range serversList {
		if serversList[s].Id == id {
			foundServer = serversList[s]
		}
	}
	if foundServer.Id == "" {
		log.Fatalf("Cannot locate server Id '%s' in our list of speedtest servers!\n", id)
	}
	return foundServer
}


func ListServers() {
	if debug.DEBUG {
		fmt.Printf("Loading config from speedtest.net\n")
	}
	sthttp.CONFIG = sthttp.GetConfig()
	if debug.DEBUG {
		fmt.Printf("\n")
	}

	if debug.DEBUG {
		fmt.Printf("Getting servers list...")
	}
	allServers := sthttp.GetServers()
	if debug.DEBUG {
		fmt.Printf("(%d) found\n", len(allServers))
	}
	for s := range allServers {
		server := allServers[s]
		print.PrintServer(server)
	}
}
