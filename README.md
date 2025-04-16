# ipapi

`ipapi` is a Go package for retrieving IP geolocation data using the [ipapi.co](https://ipapi.co/) API.

## Table of Contents

- [ipapi](#ipapi)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Get Location for a Specific IP](#get-location-for-a-specific-ip)
    - [Get Client IP Location](#get-client-ip-location)
  - [Response Structure](#response-structure)
  - [License](#license)

## Features

- Get location data for a specific IP address
- Get location data for the client making the request
- Built-in error handling

## Installation

```bash
go get github.com/derenko404/ipapi-go
```

## Usage

#### Get Location for a Specific IP

```go
location, err := ipapi.GetIpLocation("192.34.176.174")
if err != nil {
  log.Fatalf("Failed to get IP location: %v", err)
}
fmt.Println(location.City, location.Country)
```

#### Get Client IP Location

```go
location, err := ipapi.GetClientLocation()
if err != nil {
  log.Fatalf("Failed to get client location: %v", err)
}
fmt.Println(location.City, location.Country)
```

## Response Structure

The `IpapiResponse` struct includes the following fields:

- `IP` – The IP address
- `City` – The city name
- `Region` – The region name
- `RegionCode` – Region code (e.g., "HE" for Hesse)
- `Country` – Country code (e.g., "DE" for Germany)
- `CountryName` – Full country name
- `ContinentCode` – Continent code (e.g., "EU")
- `InEU` – Boolean indicating if the country is in the EU
- `Postal` – Postal code
- `Latitude` – Latitude coordinate
- `Longitude` – Longitude coordinate
- `Timezone` – Timezone name (e.g., "Europe/Berlin")
- `UTCOffset` – Offset from UTC (e.g., "+0100")
- `CountryCallingCode` – International calling code (e.g., "+49")
- `Currency` – Currency code (e.g., "EUR")
- `Languages` – Comma-separated list of languages (e.g., "de")
- `ASN` – Autonomous system number
- `Org` – Organization name (e.g., ISP)

In case of an error, the response may include:

- `Error` – Boolean indicating if there was an error
- `Reason` – Description of the error

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

You are free to use, modify, and distribute this software with proper attribution.
