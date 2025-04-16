package ipapi

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var mockIp = "192.34.176.174";

var mockSuccessResponse = map[string]any{
	"ip":                   mockIp,
	"city":                 "Frankfurt am Main",
	"region":               "Hesse",
	"region_code":          "HE",
	"country":              "DE",
	"country_name":         "Germany",
	"continent_code":       "EU",
	"in_eu":                true,
	"postal":               "60313",
	"latitude":             50.1153,
	"longitude":            8.6823,
	"timezone":             "Europe/Berlin",
	"utc_offset":           "+0100",
	"country_calling_code": "+49",
	"currency":             "EUR",
	"languages":            "de",
	"asn":                  "AS3209",
	"org":                  "Vodafone GmbH",
}

var mockErrorResponse = map[string]any{
	"ip": mockIp,
	"error": true,
	"reason": "some reason",
}

func TestGetIpLocation(t *testing.T) {
	defer gock.Off()

	// Test error case
	gock.New(baseUrl).
		Get("json").
		Reply(200).
		JSON(mockErrorResponse)

	_, err := GetIpLocation(mockIp)
	assert.Error(t, err, "it should return error when ipapi returns error")

	_, err = GetIpLocation("some fake ip")
	assert.Error(t, err, "it should return error when incorrect ip is passed")

	// Test success case
	gock.New(baseUrl).
		Get("json").
		Reply(200).
		JSON(mockSuccessResponse)

	location, err := GetIpLocation(mockIp)

	assert.NoError(t, err, "it should not return error when ipapi returns success response")
	assert.Equal(t, mockIp, location.IP, "it should map data correctly")
	assert.IsType(t, &IpapiResponse{}, location, "it should return correct type")
}

func TestGetClientLocation(t *testing.T) {
	defer gock.Off() // flush mocks when test ends

	gock.New(baseUrl).
		Get("json").
		Reply(200).
		JSON(mockErrorResponse)

	_, err := GetClientLocation()

	assert.Error(t, err, "it should return error when ipapi returns error")

	gock.New(baseUrl).
		Get("json").
		Reply(200).
		JSON(mockSuccessResponse)

	location, err := GetClientLocation()

	assert.NoError(t, err, "it should not return error when ipapi returns success response")
	assert.Equal(t, mockIp, location.IP, "it should map data correctly")
	assert.IsType(t, &IpapiResponse{}, location, "it should return correct type")
}