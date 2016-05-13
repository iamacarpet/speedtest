package tests

import (
	"fmt"
	"log"
	"strings"
)

import (
	"github.com/iamacarpet/speedtest/debug"
	"github.com/iamacarpet/speedtest/misc"
	"github.com/iamacarpet/speedtest/print"
	"github.com/iamacarpet/speedtest/settings"
	"github.com/iamacarpet/speedtest/sthttp"
)

// DownloadTest will perform the "normal" speedtest download test
func DownloadTest(server sthttp.Server) (float64, error) {
	var urls []string
	var maxSpeed float64
	var avgSpeed float64

	// http://speedtest1.newbreakcommunications.net/speedtest/speedtest/
	dlsizes := []int{350, 500, 750, 1000, 1500, 2000, 2500, 3000, 3500, 4000}

	for size := range dlsizes {
		url := server.URL
		splits := strings.Split(url, "/")
		baseURL := strings.Join(splits[1:len(splits)-1], "/")
		randomImage := fmt.Sprintf("random%dx%d.jpg", dlsizes[size], dlsizes[size])
		downloadURL := "http:/" + baseURL + "/" + randomImage
		urls = append(urls, downloadURL)
	}

	if !debug.QUIET && !debug.REPORT {
		log.Printf("Testing download speed")
	}

	for u := range urls {

		if debug.DEBUG {
			log.Printf("Download Test Run: %s\n", urls[u])
		}
		dlSpeed, err := sthttp.DownloadSpeed(urls[u])
		if err != nil {
			return 0, err
		}
		if !debug.QUIET && !debug.DEBUG && !debug.REPORT {
			fmt.Printf(".")
		}
		if debug.DEBUG {
			log.Printf("Dl Speed: %v\n", dlSpeed)
		}

		if settings.ALGOTYPE == "max" {
			if dlSpeed > maxSpeed {
				maxSpeed = dlSpeed
			}
		} else {
			avgSpeed = avgSpeed + dlSpeed
		}

	}

	if !debug.QUIET && !debug.REPORT {
		fmt.Printf("\n")
	}

	if settings.ALGOTYPE != "max" {
		return avgSpeed / float64(len(urls)), nil
	}
	return maxSpeed, nil

}

// UploadTest runs a "normal" speedtest upload test
func UploadTest(server sthttp.Server) (float64, error) {
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

	if !debug.QUIET && !debug.REPORT {
		log.Printf("Testing upload speed")
	}

	for i := 0; i < len(ulsize); i++ {
		if debug.DEBUG {
			log.Printf("Upload Test Run: %v\n", i)
		}
		r := misc.Urandom(ulsize[i])
		ulSpeed, err := sthttp.UploadSpeed(server.URL, "text/xml", r)
		if err != nil {
			return 0, err
		}
		if !debug.QUIET && !debug.DEBUG && !debug.REPORT {
			fmt.Printf(".")
		}
		if debug.DEBUG {
			log.Printf("Ul Amount: %v bytes\n", len(r))
			log.Printf("Ul Speed: %vMbps\n", ulSpeed)
		}

		if settings.ALGOTYPE == "max" {
			if ulSpeed > maxSpeed {
				maxSpeed = ulSpeed
			}
		} else {
			avgSpeed = avgSpeed + ulSpeed
		}

	}

	if !debug.QUIET && !debug.REPORT {
		fmt.Printf("\n")
	}

	if settings.ALGOTYPE != "max" {
		return avgSpeed / float64(len(ulsizesizes)), nil
	}
	return maxSpeed, nil
}

// FindServer will find a specific server in the servers list
func FindServer(id string, serversList []sthttp.Server) (sthttp.Server, error) {
	var foundServer sthttp.Server
	for s := range serversList {
		if serversList[s].ID == id {
			foundServer = serversList[s]
		}
	}
	if foundServer.ID == "" {
		return foundServer, fmt.Errorf("Cannot locate server Id '%s' in our list of speedtest servers!\n", id)
	}
	return foundServer, nil
}

// ListServers prints a list of all "global" servers
func ListServers() (error) {
	if debug.DEBUG {
		fmt.Printf("Loading config from speedtest.net\n")
	}
	var err error
	sthttp.CONFIG, err = sthttp.GetConfig()
	if err != nil {
		return err
	}
	if debug.DEBUG {
		fmt.Printf("\n")
	}

	if debug.DEBUG {
		fmt.Printf("Getting servers list...")
	}
	allServers, err := sthttp.GetServers()
	if err != nil {
		return err
	}
	if debug.DEBUG {
		fmt.Printf("(%d) found\n", len(allServers))
	}
	for s := range allServers {
		server := allServers[s]
		print.Server(server)
	}

	return nil
}
