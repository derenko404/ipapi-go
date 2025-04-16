package ipapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type IpapiResponse struct {
	IP                 string  `json:"ip"`
	Network            string  `json:"network"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeISO3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTLD         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEU               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UTCOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	ASN                string  `json:"asn"`
	Org                string  `json:"org"`
}

type ipapiError struct {
	IP      string  `json:"ip"`
	Error   bool   `json:"error"`
	Reason  string `json:"reason"`
}

const baseUrl = "https://ipapi.co"

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func getResponseFormat(params []string) string {
	responseFormat := "json"
	if len(params) > 0 && params[0] != "" {
		responseFormat = params[0]
	}

	return responseFormat
}

func getRequestUrl(ip string, format string) string {
	if ip == "" {
		if format == "" {
			return fmt.Sprintf("%s/%s/json", baseUrl, ip)
		}

		return fmt.Sprintf("%s/json", baseUrl)
	}

	return fmt.Sprintf("%s/%s/%s", baseUrl, ip, format)
}

func getRequestObject(ip string, format string) (*http.Request, error) {
	req, err := http.NewRequest("GET", getRequestUrl(ip, format), nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func doRequest(req *http.Request) (*IpapiResponse, error) {
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	var raw map[string]any

	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %w", err)
	}

	if isError, ok := raw["error"].(bool); ok && isError {
		// re-decode into ErrorResponse
		var apiErr ipapiError
		data, _ := json.Marshal(raw)
		json.Unmarshal(data, &apiErr)
		return nil, fmt.Errorf("API returns error: %s", apiErr.Reason)
	}
	
	var data IpapiResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal API response %w", err)
	}

	return &data, nil
}

func GetIpLocation(ip string, params ...string) (*IpapiResponse, error) {
	if isValidIP(ip) {
		format := getResponseFormat(params)
		req, err := getRequestObject(ip, format)
		if err != nil {
			return nil, err
		}

		response, err := doRequest(req)

		if err != nil {
			return nil, err
		}

		return response, nil
	}

	return nil, fmt.Errorf("invalid ip address %s", ip)
}

func GetClientLocation(params ...string) (*IpapiResponse, error) {
	req, err := getRequestObject("", getResponseFormat(params))
	if err != nil {
		return nil, err
	}

	return doRequest(req)
}
