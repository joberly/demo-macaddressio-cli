package api

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO add more tests using a mock server

// TestDecodeSearch tests decoding of a JSON response body to the
// SearchResponse structure.
func TestDecodeSearch(t *testing.T) {
	// Test data from https://macaddress.io/api/documentation/output-format
	testJSON := `{
	"vendorDetails":{
		"oui":"443839",
		"isPrivate":false,
		"companyName":"Cumulus Networks, Inc",
		"companyAddress":"650 Castro Street, suite 120-245 Mountain View  CA  94041 US",
		"countryCode":"US"
	},
	"blockDetails":{
		"blockFound":true,
		"borderLeft":"443839000000",
		"borderRight":"443839FFFFFF",
		"blockSize":16777216,
		"assignmentBlockSize":"MA-L",
		"dateCreated":"2012-04-08",
		"dateUpdated":"2015-09-27"
	},
	"macAddressDetails":{
		"searchTerm":"44:38:39:ff:ef:57",
		"isValid":true,
		"virtualMachine":"Not detected",
		"applications":[
			"Multi-Chassis Link Aggregation (Cumulus Linux)"
		],
		"transmissionType":"unicast",
		"administrationType":"UAA",
		"wiresharkNotes":"No details",
		"comment":""
	}
}`

	// Write test output
	respBuf := &bytes.Buffer{}
	respBuf.WriteString(testJSON)

	// Test decode
	resp, err := decodeSearch(respBuf)

	// Above test should be successful
	assert.Nil(t, err)

	// Above test should return a response structure
	if assert.NotNil(t, resp) {
		// Above test should return valid data
		assert.Equal(t, "Cumulus Networks, Inc", resp.VendorDetails.CompanyName)
	}
}
