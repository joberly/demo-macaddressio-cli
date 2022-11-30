package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// SearchResponse contains the data from the search query.
// FUTURE add more fields as needed
type SearchResponse struct {
	VendorDetails SearchResponseVendorDetails
}

// SearchResponseVendorDetails contains the vendorDetails search query data.
// FUTURE add more fields as needed
type SearchResponseVendorDetails struct {
	CompanyName string
}

var (
	// Search query parameter for calling request().
	querySearch = "search=%s"
)

// Search queries macaddress.io for MAC address information.
// See https://macaddress.io/api/documentation/making-requests for more info.
func (c *Client) Search(macaddr net.HardwareAddr) (*SearchResponse, error) {
	// Make the query API call
	respBuf, err := c.request(fmt.Sprintf(querySearch, macaddr.String()))

	// Check if the request failed
	if err != nil {
		return nil, err
	}

	// Return the decoded response
	return decodeSearch(respBuf)
}

// DecodeSearch decodes the search query response JSON to the
// SearchResponse structure.
func decodeSearch(respBuf *bytes.Buffer) (*SearchResponse, error) {
	// Decode the response JSON
	dec := json.NewDecoder(respBuf)
	resp := &SearchResponse{}
	err := dec.Decode(resp)

	// Check for JSON decoding errors
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	// Return the decoded response data
	return resp, nil
}
