package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/joberly/demo-macaddressio-cli/macaddressio/api"
)

// Running the CLI requires the MAC address parameter on the command line,
// and the API key in the MACADDRESSIO_API_KEY environment variable.
//
// > export MACADDRIO_API_KEY="<REDACTED>"
// > macaddrio 00:11:22:33:44:55

const EnvAPIKey = "MACADDRIO_API_KEY"

// Main is the main CLI function.
func main() {
	// Check for MAC address parameter
	if len(os.Args) < 2 {
		msg("no mac address specified")
		os.Exit(1)
	}

	// Lookup the API key in the environment
	apiKey, apiKeyPresent := os.LookupEnv(EnvAPIKey)

	// Print a particular message if env var is not present
	if !apiKeyPresent {
		msg(fmt.Sprintf("%s environment variable not found", EnvAPIKey))
		os.Exit(2)
	}

	// Trim any quotes around the API key value just in case
	apiKey = strings.Trim(apiKey, "\"")

	// Check that the API key is not empty
	if apiKey == "" {
		msg("empty API key")
		os.Exit(3)
	}

	// Validate MAC address parameter
	mac := os.Args[1]
	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		msg(fmt.Sprintf("badly formatted mac address: %s", err.Error()))
		os.Exit(4)
	}

	// Create API client
	c := api.New(api.WithAPIKey(apiKey))

	// Make request
	resp, err := c.Search(hwAddr)
	if err != nil {
		msg(fmt.Sprintf("API error: %s", err.Error()))
		os.Exit(5)
	}

	// Print company name from API
	fmt.Printf(resp.VendorDetails.CompanyName)

	// Exit successfully
	os.Exit(0)
}

// Msg prints a message prefixed by the executable name.
func msg(m string) {
	fmt.Printf("macaddrio: %s", m)
}
